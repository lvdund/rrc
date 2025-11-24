package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MCCH_MessageType_r17_Choice_nothing uint64 = iota
	MCCH_MessageType_r17_Choice_C1
	MCCH_MessageType_r17_Choice_MessageClassExtension
)

type MCCH_MessageType_r17 struct {
	Choice                uint64
	C1                    *MCCH_MessageType_r17_C1
	MessageClassExtension interface{} `madatory`
}

func (ie *MCCH_MessageType_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MCCH_MessageType_r17_Choice_C1:
		if err = ie.C1.Encode(w); err != nil {
			err = utils.WrapError("Encode C1", err)
		}
	case MCCH_MessageType_r17_Choice_MessageClassExtension:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *MCCH_MessageType_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MCCH_MessageType_r17_Choice_C1:
		ie.C1 = new(MCCH_MessageType_r17_C1)
		if err = ie.C1.Decode(r); err != nil {
			return utils.WrapError("Decode C1", err)
		}
	case MCCH_MessageType_r17_Choice_MessageClassExtension:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
