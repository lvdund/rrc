package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SBCCH_SL_BCH_MessageType_Choice_nothing uint64 = iota
	SBCCH_SL_BCH_MessageType_Choice_C1
	SBCCH_SL_BCH_MessageType_Choice_MessageClassExtension
)

type SBCCH_SL_BCH_MessageType struct {
	Choice                uint64
	C1                    *SBCCH_SL_BCH_MessageType_C1
	MessageClassExtension interface{} `madatory`
}

func (ie *SBCCH_SL_BCH_MessageType) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SBCCH_SL_BCH_MessageType_Choice_C1:
		if err = ie.C1.Encode(w); err != nil {
			err = utils.WrapError("Encode C1", err)
		}
	case SBCCH_SL_BCH_MessageType_Choice_MessageClassExtension:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SBCCH_SL_BCH_MessageType) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SBCCH_SL_BCH_MessageType_Choice_C1:
		ie.C1 = new(SBCCH_SL_BCH_MessageType_C1)
		if err = ie.C1.Decode(r); err != nil {
			return utils.WrapError("Decode C1", err)
		}
	case SBCCH_SL_BCH_MessageType_Choice_MessageClassExtension:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
