package wire

import (
	"bytes"
	"fmt"
	"io"
)

type SNACError struct {
	Code uint16
	TLVRestBlock
}

const (
	FLAPFrameSignon    uint8 = 0x01
	FLAPFrameData      uint8 = 0x02
	FLAPFrameError     uint8 = 0x03
	FLAPFrameSignoff   uint8 = 0x04
	FLAPFrameKeepAlive uint8 = 0x05
)

type FLAPFrame struct {
	StartMarker uint8
	FrameType   uint8
	Sequence    uint16
	Payload     []byte `oscar:"len_prefix=uint16"`
}

// FLAPFrameDisconnect is the last FLAP frame sent to a client before
// disconnection. It differs from FLAPFrame in that there is no payload length
// prefix at the end, which causes Windows AIM clients to improperly handle
// server disconnections, as when the regular FLAPFrame type is used.
type FLAPFrameDisconnect struct {
	StartMarker uint8
	FrameType   uint8
	Sequence    uint16
}

type SNACFrame struct {
	FoodGroup uint16
	SubGroup  uint16
	Flags     uint16
	RequestID uint32
}

// ReqIDFromServer is the SNAC frame Request ID value that indicates the SNAC
// is initiated by the server. Some clients, such as the Java AIM 1.1.19,
// completely fail to process some server SNACs if the high bit is not set on
// request ID.
const ReqIDFromServer = 1 << 31

type FLAPSignonFrame struct {
	FLAPVersion uint32
	TLVRestBlock
}

type SNACMessage struct {
	Frame SNACFrame
	Body  any
}

// NewFlapClient creates a new FLAP client instance. startSeq is the initial
// sequence value, which is typically 0. r receives FLAP messages, w writes
// FLAP messages.
func NewFlapClient(startSeq uint32, r io.Reader, w io.Writer) *FlapClient {
	return &FlapClient{
		sequence: startSeq,
		r:        r,
		w:        w,
	}
}

// FlapClient sends and receive FLAP frames to and from the server. It ensures
// that the message sequence numbers are properly incremented after sending
// each successive message. It is not safe to use with multiple goroutines
// without synchronization.
type FlapClient struct {
	sequence uint32
	r        io.Reader
	w        io.Writer
}

// Fixes a race condition caused by testify. Yup...
// https://github.com/stretchr/testify/issues/625
func (f *FlapClient) String() string {
	return ""
}

// SendSignonFrame sends a signon FLAP frame containing a list of TLVs to
// authenticate or initiate a session.
func (f *FlapClient) SendSignonFrame(tlvs []TLV) error {
	signonFrame := FLAPSignonFrame{
		FLAPVersion: 1,
	}
	if len(tlvs) > 0 {
		signonFrame.AppendList(tlvs)
	}
	buf := &bytes.Buffer{}
	if err := MarshalBE(signonFrame, buf); err != nil {
		return err
	}

	flap := FLAPFrame{
		StartMarker: 42,
		FrameType:   FLAPFrameSignon,
		Sequence:    uint16(f.sequence),
		Payload:     buf.Bytes(),
	}
	if err := MarshalBE(flap, f.w); err != nil {
		return err
	}

	f.sequence++

	return nil
}

// ReceiveSignonFrame receives a signon FLAP response message.
func (f *FlapClient) ReceiveSignonFrame() (FLAPSignonFrame, error) {
	flap := FLAPFrame{}
	if err := UnmarshalBE(&flap, f.r); err != nil {
		return FLAPSignonFrame{}, err
	}

	signonFrame := FLAPSignonFrame{}
	if err := UnmarshalBE(&signonFrame, bytes.NewBuffer(flap.Payload)); err != nil {
		return FLAPSignonFrame{}, err
	}

	return signonFrame, nil
}

// ReceiveFLAP receives a FLAP frame and body. It only returns a body if the
// FLAP frame is a data frame.
func (f *FlapClient) ReceiveFLAP() (FLAPFrame, error) {
	flap := FLAPFrame{}
	err := UnmarshalBE(&flap, f.r)
	if err != nil {
		err = fmt.Errorf("unable to unmarshal FLAP frame: %w", err)
	}
	return flap, err
}

// SendSignoffFrame sends a sign-off FLAP frame with attached TLVs as the last
// request sent in the FLAP auth flow. This is unrelated to the Disconnect()
// method, which sends a sign-off frame to terminate a BOS connection.
// todo: combine this method with Disconnect()
func (f *FlapClient) SendSignoffFrame(tlvs TLVRestBlock) error {
	tlvBuf := &bytes.Buffer{}
	if err := MarshalBE(tlvs, tlvBuf); err != nil {
		return err
	}

	flap := FLAPFrame{
		StartMarker: 42,
		FrameType:   FLAPFrameSignoff,
		Sequence:    uint16(f.sequence),
		Payload:     tlvBuf.Bytes(),
	}

	if err := MarshalBE(flap, f.w); err != nil {
		return err
	}

	f.sequence++
	return nil
}

// SendSNAC sends a SNAC message wrapped in a FLAP frame.
func (f *FlapClient) SendSNAC(frame SNACFrame, body any) error {
	snacBuf := &bytes.Buffer{}
	if err := MarshalBE(frame, snacBuf); err != nil {
		return err
	}
	if err := MarshalBE(body, snacBuf); err != nil {
		return err
	}

	flap := FLAPFrame{
		StartMarker: 42,
		FrameType:   FLAPFrameData,
		Sequence:    uint16(f.sequence),
		Payload:     snacBuf.Bytes(),
	}
	if err := MarshalBE(flap, f.w); err != nil {
		return err
	}

	f.sequence++
	return nil
}

func (f *FlapClient) SendDataFrame(payload []byte) error {
	flap := FLAPFrame{
		StartMarker: 42,
		FrameType:   FLAPFrameData,
		Sequence:    uint16(f.sequence),
		Payload:     payload,
	}
	if err := MarshalBE(flap, f.w); err != nil {
		return err
	}

	f.sequence++
	return nil
}

// ReceiveSNAC receives a SNAC message wrapped in a FLAP frame.
func (f *FlapClient) ReceiveSNAC(frame *SNACFrame, body any) error {
	flap := FLAPFrame{}
	if err := UnmarshalBE(&flap, f.r); err != nil {
		return err
	}
	buf := bytes.NewBuffer(flap.Payload)
	if err := UnmarshalBE(frame, buf); err != nil {
		return err
	}
	return UnmarshalBE(body, buf)
}

// Disconnect sends a signoff FLAP frame.
func (f *FlapClient) Disconnect() error {
	flap := FLAPFrameDisconnect{
		StartMarker: 42,
		FrameType:   FLAPFrameSignoff,
		Sequence:    uint16(f.sequence),
	}
	return MarshalBE(flap, f.w)
}
