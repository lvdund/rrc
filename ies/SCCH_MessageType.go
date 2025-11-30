package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SCCH_MessageType_Choice_nothing uint64 = iota
	SCCH_MessageType_Choice_C1
	SCCH_MessageType_Choice_MessageClassExtension
)

type SCCH_MessageType struct {
	Choice                uint64
	C1                    *SCCH_MessageType_C1
	MessageClassExtension *SCCH_MessageType_MessageClassExtension
}

func (ie *SCCH_MessageType) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SCCH_MessageType_Choice_C1:
		if err = ie.C1.Encode(w); err != nil {
			err = utils.WrapError("Encode C1", err)
		}
	case SCCH_MessageType_Choice_MessageClassExtension:
		if err = ie.MessageClassExtension.Encode(w); err != nil {
			err = utils.WrapError("Encode MessageClassExtension", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SCCH_MessageType) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SCCH_MessageType_Choice_C1:
		ie.C1 = new(SCCH_MessageType_C1)
		if err = ie.C1.Decode(r); err != nil {
			return utils.WrapError("Decode C1", err)
		}
	case SCCH_MessageType_Choice_MessageClassExtension:
		ie.MessageClassExtension = new(SCCH_MessageType_MessageClassExtension)
		if err = ie.MessageClassExtension.Decode(r); err != nil {
			return utils.WrapError("Decode MessageClassExtension", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
