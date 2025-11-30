package rrc

import (
	"bytes"
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/ies"
	"github.com/lvdund/rrc/utils"
)

type RRCMessage interface {
	Encode(w *aper.AperWriter) error
	Decode(r *aper.AperReader) error
}

// MessageContainerType represents the type of RRC message container
type MessageContainerType int

const (
	MessageContainerTypeUnknown MessageContainerType = iota
	MessageContainerTypeUL_CCCH
	MessageContainerTypeUL_DCCH
	MessageContainerTypeDL_CCCH
	MessageContainerTypeDL_DCCH
	MessageContainerTypeBCCH_BCH
	MessageContainerTypeBCCH_DL_SCH
	MessageContainerTypePCCH
)

func (t MessageContainerType) String() string {
	switch t {
	case MessageContainerTypeUL_CCCH:
		return "UL-CCCH"
	case MessageContainerTypeUL_DCCH:
		return "UL-DCCH"
	case MessageContainerTypeDL_CCCH:
		return "DL-CCCH"
	case MessageContainerTypeDL_DCCH:
		return "DL-DCCH"
	case MessageContainerTypeBCCH_BCH:
		return "BCCH-BCH"
	case MessageContainerTypeBCCH_DL_SCH:
		return "BCCH-DL-SCH"
	case MessageContainerTypePCCH:
		return "PCCH"
	default:
		return "Unknown"
	}
}

// DecodedMessage wraps a decoded RRC message with its container type
type DecodedMessage struct {
	Type    MessageContainerType
	Message RRCMessage
}

func Encode(msg RRCMessage) ([]byte, error) {
	var buf bytes.Buffer
	writer := aper.NewWriter(&buf)

	if err := msg.Encode(writer); err != nil {
		return nil, utils.WrapError("Encode RRCMessage", err)
	}

	if err := writer.Close(); err != nil {
		return nil, utils.WrapError("Close writer", err)
	}

	return buf.Bytes(), nil
}

// Decode decodes a known RRC message type (when you know the container type)
func Decode(buf []byte, msg RRCMessage) error {
	reader := aper.NewReader(bytes.NewBuffer(buf))

	if msg == nil {
		return utils.WrapError("Decode RRCMessage", fmt.Errorf("message cannot be nil"))
	}

	if err := msg.Decode(reader); err != nil {
		return utils.WrapError("Decode RRCMessage", err)
	}

	return nil
}

// DecodeAny attempts to decode the byte array by trying all possible container types.
// Returns the decoded message along with its container type, or an error if none match.
// This is useful when you receive bytes but don't know which container type it is.
func DecodeAny(buf []byte) (*DecodedMessage, error) {
	if len(buf) == 0 {
		return nil, fmt.Errorf("empty buffer")
	}

	// Try each container type in order
	containerTypes := []struct {
		Type    MessageContainerType
		Factory func() RRCMessage
		Name    string
	}{
		{MessageContainerTypeUL_CCCH, func() RRCMessage { return &ies.UL_CCCH_Message{} }, "UL-CCCH"},
		{MessageContainerTypeUL_DCCH, func() RRCMessage { return &ies.UL_DCCH_Message{} }, "UL-DCCH"},
		{MessageContainerTypeDL_CCCH, func() RRCMessage { return &ies.DL_CCCH_Message{} }, "DL-CCCH"},
		{MessageContainerTypeDL_DCCH, func() RRCMessage { return &ies.DL_DCCH_Message{} }, "DL-DCCH"},
		{MessageContainerTypeBCCH_BCH, func() RRCMessage { return &ies.BCCH_BCH_Message{} }, "BCCH-BCH"},
		{MessageContainerTypeBCCH_DL_SCH, func() RRCMessage { return &ies.BCCH_DL_SCH_Message{} }, "BCCH-DL-SCH"},
		{MessageContainerTypePCCH, func() RRCMessage { return &ies.PCCH_Message{} }, "PCCH"},
	}

	var lastErr error
	for _, ct := range containerTypes {
		msg := ct.Factory()
		reader := aper.NewReader(bytes.NewReader(buf))

		if err := msg.Decode(reader); err == nil {
			// Check if we consumed all bytes (successful decode)
			// Note: APER reader might not have a way to check remaining bytes,
			// so we rely on successful decode without error
			return &DecodedMessage{
				Type:    ct.Type,
				Message: msg,
			}, nil
		} else {
			lastErr = err
		}
	}

	return nil, fmt.Errorf("failed to decode as any known container type, last error: %w", lastErr)
}

// DecodeWithChannel decodes a message when you know the channel/container type.
// This is more efficient than DecodeAny when you know the context.
func DecodeWithChannel(buf []byte, containerType MessageContainerType) (RRCMessage, error) {
	var msg RRCMessage

	switch containerType {
	case MessageContainerTypeUL_CCCH:
		msg = &ies.UL_CCCH_Message{}
	case MessageContainerTypeUL_DCCH:
		msg = &ies.UL_DCCH_Message{}
	case MessageContainerTypeDL_CCCH:
		msg = &ies.DL_CCCH_Message{}
	case MessageContainerTypeDL_DCCH:
		msg = &ies.DL_DCCH_Message{}
	case MessageContainerTypeBCCH_BCH:
		msg = &ies.BCCH_BCH_Message{}
	case MessageContainerTypeBCCH_DL_SCH:
		msg = &ies.BCCH_DL_SCH_Message{}
	case MessageContainerTypePCCH:
		msg = &ies.PCCH_Message{}
	default:
		return nil, fmt.Errorf("unknown container type: %v", containerType)
	}

	if err := Decode(buf, msg); err != nil {
		return nil, err
	}

	return msg, nil
}
