package rrc

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCMessage interface {
	Encode(w *uper.UperWriter) error
	Decode(r *uper.UperReader) error
}

func Encode(msg RRCMessage) ([]byte, error) {
	var buf []byte
	writer := uper.NewWriter(bytes.NewBuffer(buf))

	if err := msg.Encode(writer); err != nil {
		return nil, utils.WrapError("Encode RRCMessage", err)
	}

	if err := writer.Close(); err != nil {
		return nil, utils.WrapError("Close writer", err)
	}

	return buf, nil
}

func Decode(buf []byte) (RRCMessage, error) {
	reader := uper.NewReader(bytes.NewBuffer(buf))

	var msg RRCMessage
	if err := msg.Decode(reader); err != nil {
		return nil, utils.WrapError("Decode RRCMessage", err)
	}

	return msg, nil
}
