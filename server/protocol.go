package server

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/mkaminski/goaim/oscar"
	"io"
)

const (
	ErrorCodeInvalidSnac          uint16 = 0x01
	ErrorCodeRateToHost           uint16 = 0x02
	ErrorCodeRateToClient         uint16 = 0x03
	ErrorCodeNotLoggedOn          uint16 = 0x04
	ErrorCodeServiceUnavailable   uint16 = 0x05
	ErrorCodeServiceNotDefined    uint16 = 0x06
	ErrorCodeObsoleteSnac         uint16 = 0x07
	ErrorCodeNotSupportedByHost   uint16 = 0x08
	ErrorCodeNotSupportedByClient uint16 = 0x09
	ErrorCodeRefusedByClient      uint16 = 0x0A
	ErrorCodeReplyTooBig          uint16 = 0x0B
	ErrorCodeResponsesLost        uint16 = 0x0C
	ErrorCodeRequestDenied        uint16 = 0x0D
	ErrorCodeBustedSnacPayload    uint16 = 0x0E
	ErrorCodeInsufficientRights   uint16 = 0x0F
	ErrorCodeInLocalPermitDeny    uint16 = 0x10
	ErrorCodeTooEvilSender        uint16 = 0x11
	ErrorCodeTooEvilReceiver      uint16 = 0x12
	ErrorCodeUserTempUnavail      uint16 = 0x13
	ErrorCodeNoMatch              uint16 = 0x14
	ErrorCodeListOverflow         uint16 = 0x15
	ErrorCodeRequestAmbigous      uint16 = 0x16
	ErrorCodeQueueFull            uint16 = 0x17
	ErrorCodeNotWhileOnAol        uint16 = 0x18
	ErrorCodeQueryFail            uint16 = 0x19
	ErrorCodeTimeout              uint16 = 0x1A
	ErrorCodeErrorText            uint16 = 0x1B
	ErrorCodeGeneralFailure       uint16 = 0x1C
	ErrorCodeProgress             uint16 = 0x1D
	ErrorCodeInFreeArea           uint16 = 0x1E
	ErrorCodeRestrictedByPc       uint16 = 0x1F
	ErrorCodeRemoteRestrictedByPc uint16 = 0x20
)

const (
	ErrorTagsFailUrl        = 0x04
	ErrorTagsErrorSubcode   = 0x08
	ErrorTagsErrorText      = 0x1B
	ErrorTagsErrorInfoClsid = 0x29
	ErrorTagsErrorInfoData  = 0x2A
)

var (
	CapChat, _ = uuid.MustParse("748F2420-6287-11D1-8222-444553540000").MarshalBinary()
)

type Config struct {
	OSCARHost string `envconfig:"OSCAR_HOST" required:"true"`
	OSCARPort int    `envconfig:"OSCAR_PORT" default:"5190"`
	BOSPort   int    `envconfig:"BOS_PORT" default:"5191"`
	ChatPort  int    `envconfig:"CHAT_PORT" default:"5192"`
	DBPath    string `envconfig:"DB_PATH" required:"true"`
}

func Address(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}

func SendAndReceiveSignonFrame(rw io.ReadWriter, sequence *uint32) (*oscar.FlapSignonFrame, error) {
	flapOut := oscar.FlapSignonFrame{
		FlapFrame: oscar.FlapFrame{
			StartMarker:   42,
			FrameType:     1,
			Sequence:      uint16(*sequence),
			PayloadLength: 4, // size of flapVersion
		},
		FlapVersion: 1,
	}

	if err := flapOut.Write(rw); err != nil {
		return nil, err
	}

	fmt.Printf("SendAndReceiveSignonFrame read FLAP: %+v\n", flapOut)

	// receive
	flapIn := oscar.FlapSignonFrame{}
	if err := flapIn.Read(rw); err != nil {
		return nil, err
	}

	fmt.Printf("SendAndReceiveSignonFrame write FLAP: %+v\n", flapIn)

	*sequence++

	return &flapIn, nil
}

func VerifyLogin(sm *SessionManager, rw io.ReadWriter) (*Session, uint32, error) {
	seq := uint32(100)
	fmt.Println("VerifyLogin...")

	flap, err := SendAndReceiveSignonFrame(rw, &seq)
	if err != nil {
		return nil, 0, err
	}

	var ok bool
	ID, ok := flap.GetSlice(OserviceTlvTagsLoginCookie)
	if !ok {
		return nil, 0, errors.New("unable to get session ID from payload")
	}

	sess, ok := sm.Retrieve(string(ID))
	if !ok {
		return nil, 0, fmt.Errorf("unable to find session by ID %s", ID)
	}

	return sess, seq, nil
}

func VerifyChatLogin(rw io.ReadWriter) (*ChatCookie, uint32, error) {
	seq := uint32(100)
	fmt.Println("VerifyChatLogin...")

	flap, err := SendAndReceiveSignonFrame(rw, &seq)
	if err != nil {
		return nil, 0, err
	}

	var ok bool
	buf, ok := flap.GetSlice(OserviceTlvTagsLoginCookie)
	if !ok {
		return nil, 0, errors.New("unable to get session ID from payload")
	}

	cookie := ChatCookie{}
	err = cookie.Read(bytes.NewBuffer(buf))

	return &cookie, seq, err
}

const (
	OSERVICE      uint16 = 0x0001
	LOCATE               = 0x0002
	BUDDY                = 0x0003
	ICBM                 = 0x0004
	ADVERT               = 0x0005
	INVITE               = 0x0006
	ADMIN                = 0x0007
	POPUP                = 0x0008
	PD                   = 0x0009
	USER_LOOKUP          = 0x000A
	STATS                = 0x000B
	TRANSLATE            = 0x000C
	CHAT_NAV             = 0x000D
	CHAT                 = 0x000E
	ODIR                 = 0x000F
	BART                 = 0x0010
	FEEDBAG              = 0x0013
	ICQ                  = 0x0015
	BUCP                 = 0x0017
	ALERT                = 0x0018
	PLUGIN               = 0x0022
	UNNAMED_FG_24        = 0x0024
	MDIR                 = 0x0025
	ARS                  = 0x044A
)

type IncomingMessage struct {
	flap oscar.FlapFrame
	snac oscar.SnacFrame
	buf  io.Reader
}

type XMessage struct {
	snacFrame oscar.SnacFrame
	snacOut   any
}

const (
	FlapFrameSignon    uint8 = 0x01
	FlapFrameData            = 0x02
	FlapFrameError           = 0x03
	FlapFrameSignoff         = 0x04
	FlapFrameKeepAlive       = 0x05
)

func readIncomingRequests(rw io.Reader, msCh chan IncomingMessage, errCh chan error) {
	defer close(msCh)
	defer close(errCh)

	for {
		flap := oscar.FlapFrame{}
		if err := oscar.Unmarshal(&flap, rw); err != nil {
			errCh <- err
			return
		}

		switch flap.FrameType {
		case FlapFrameSignon:
			errCh <- errors.New("shouldn't get FlapFrameSignon")
			return
		case FlapFrameData:
			b := make([]byte, flap.PayloadLength)
			if _, err := rw.Read(b); err != nil {
				errCh <- err
				return
			}

			snac := oscar.SnacFrame{}
			buf := bytes.NewBuffer(b)
			if err := oscar.Unmarshal(&snac, buf); err != nil {
				errCh <- err
				return
			}

			msCh <- IncomingMessage{
				flap: flap,
				snac: snac,
				buf:  buf,
			}
		case FlapFrameError:
			errCh <- fmt.Errorf("got FlapFrameError: %v", flap)
			return
		case FlapFrameSignoff:
			errCh <- fmt.Errorf("got signoff: %v", flap)
			return
		case FlapFrameKeepAlive:
			fmt.Println("keepalive heartbeat")
		default:
			errCh <- fmt.Errorf("unknown frame type: %v", flap)
			return
		}
	}
}

func Signout(sess *Session, sm *SessionManager, fm *FeedbagStore) {
	if err := NotifyDeparture(sess, sm, fm); err != nil {
		fmt.Printf("error notifying departure: %s", err.Error())
	}
	sm.Remove(sess)
}

func ReadBos(cfg Config, ready OnReadyCB, sess *Session, seq uint32, sm *SessionManager, fm *FeedbagStore, cr *ChatRegistry, rwc io.ReadWriter, foodGroups []uint16) error {
	if err := WriteOServiceHostOnline(foodGroups, rwc, &seq); err != nil {
		return err
	}

	// buffered so that the go routine has room to exit
	msgCh := make(chan IncomingMessage, 1)
	errCh := make(chan error, 1)
	go readIncomingRequests(rwc, msgCh, errCh)

	for {
		select {
		case m := <-msgCh:
			if err := routeIncomingRequests(cfg, ready, sm, sess, fm, cr, rwc, &seq, m.snac, m.buf); err != nil {
				return err
			}
		case m := <-sess.RecvMessage():
			if err := writeOutSNAC(oscar.SnacFrame{}, m.snacFrame, m.snacOut, &seq, rwc); err != nil {
				return err
			}
		case err := <-errCh:
			return err
		}
	}
}

func routeIncomingRequests(cfg Config, ready OnReadyCB, sm *SessionManager, sess *Session, fm *FeedbagStore, cr *ChatRegistry, rw io.ReadWriter, sequence *uint32, snac oscar.SnacFrame, buf io.Reader) error {
	switch snac.FoodGroup {
	case OSERVICE:
		if err := routeOService(cfg, ready, cr, sm, fm, sess, snac, buf, rw, sequence); err != nil {
			return err
		}
	case LOCATE:
		if err := routeLocate(sess, sm, fm, snac, buf, rw, sequence); err != nil {
			return err
		}
	case BUDDY:
		if err := routeBuddy(snac, buf, rw, sequence); err != nil {
			return err
		}
	case ICBM:
		if err := routeICBM(sm, fm, sess, snac, buf, rw, sequence); err != nil {
			return err
		}
	case PD:
		if err := routePD(snac, buf, rw, sequence); err != nil {
			return err
		}
	case CHAT_NAV:
		if err := routeChatNav(sess, cr, snac, buf, rw, sequence); err != nil {
			return err
		}
	case FEEDBAG:
		if err := routeFeedbag(sm, sess, fm, snac, buf, rw, sequence); err != nil {
			return err
		}
	case BUCP:
		if err := routeBUCP(snac); err != nil {
			return err
		}
	case CHAT:
		if err := routeChat(sess, sm, snac, buf, rw, sequence); err != nil {
			return err
		}
	default:
		panic(fmt.Sprintf("unsupported food group: %d", snac.FoodGroup))
	}

	return nil
}

func writeOutSNAC(originsnac oscar.SnacFrame, snacFrame oscar.SnacFrame, snacOut any, sequence *uint32, w io.Writer) error {
	if originsnac.RequestID != 0 {
		snacFrame.RequestID = originsnac.RequestID
	}

	snacBuf := &bytes.Buffer{}
	if err := oscar.Marshal(snacFrame, snacBuf); err != nil {
		return err
	}
	if err := oscar.Marshal(snacOut, snacBuf); err != nil {
		return err
	}

	flap := oscar.FlapFrame{
		StartMarker:   42,
		FrameType:     2,
		Sequence:      uint16(*sequence),
		PayloadLength: uint16(snacBuf.Len()),
	}

	fmt.Printf(" write FLAP: %+v\n", flap)

	if err := oscar.Marshal(flap, w); err != nil {
		return err
	}

	fmt.Printf(" write SNAC: %+v\n", snacOut)

	expectLen := snacBuf.Len()
	c, err := w.Write(snacBuf.Bytes())
	if err != nil {
		return err
	}
	if c != expectLen {
		panic("did not write the expected # of bytes")
	}

	*sequence++
	return nil
}
