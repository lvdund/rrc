package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	DL_DCCH_MessageType_Choice_nothing uint64 = iota
	DL_DCCH_MessageType_Choice_C1
	DL_DCCH_MessageType_Choice_MessageClassExtension
)

type DL_DCCH_MessageType struct {
	Choice                uint64
	C1                    *DL_DCCH_MessageType_C1
	MessageClassExtension interface{} `madatory`
}

func (ie *DL_DCCH_MessageType) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case DL_DCCH_MessageType_Choice_C1:
		if err = ie.C1.Encode(w); err != nil {
			err = utils.WrapError("Encode C1", err)
		}
	case DL_DCCH_MessageType_Choice_MessageClassExtension:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *DL_DCCH_MessageType) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case DL_DCCH_MessageType_Choice_C1:
		ie.C1 = new(DL_DCCH_MessageType_C1)
		if err = ie.C1.Decode(r); err != nil {
			return utils.WrapError("Decode C1", err)
		}
	case DL_DCCH_MessageType_Choice_MessageClassExtension:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
