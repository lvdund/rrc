package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SCCH_MessageType_MessageClassExtension_Choice_nothing uint64 = iota
	SCCH_MessageType_MessageClassExtension_Choice_C2
	SCCH_MessageType_MessageClassExtension_Choice_MessageClassExtensionFuture_r17
)

type SCCH_MessageType_MessageClassExtension struct {
	Choice                          uint64
	C2                              *SCCH_MessageType_MessageClassExtension_C2
	MessageClassExtensionFuture_r17 interface{} `madatory`
}

func (ie *SCCH_MessageType_MessageClassExtension) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SCCH_MessageType_MessageClassExtension_Choice_C2:
		if err = ie.C2.Encode(w); err != nil {
			err = utils.WrapError("Encode C2", err)
		}
	case SCCH_MessageType_MessageClassExtension_Choice_MessageClassExtensionFuture_r17:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SCCH_MessageType_MessageClassExtension) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SCCH_MessageType_MessageClassExtension_Choice_C2:
		ie.C2 = new(SCCH_MessageType_MessageClassExtension_C2)
		if err = ie.C2.Decode(r); err != nil {
			return utils.WrapError("Decode C2", err)
		}
	case SCCH_MessageType_MessageClassExtension_Choice_MessageClassExtensionFuture_r17:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
