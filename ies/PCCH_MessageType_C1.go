package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PCCH_MessageType_C1_Choice_nothing uint64 = iota
	PCCH_MessageType_C1_Choice_Paging
	PCCH_MessageType_C1_Choice_Spare1
)

type PCCH_MessageType_C1 struct {
	Choice uint64
	Paging *Paging
	Spare1 aper.NULL `madatory`
}

func (ie *PCCH_MessageType_C1) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PCCH_MessageType_C1_Choice_Paging:
		if err = ie.Paging.Encode(w); err != nil {
			err = utils.WrapError("Encode Paging", err)
		}
	case PCCH_MessageType_C1_Choice_Spare1:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Spare1", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *PCCH_MessageType_C1) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PCCH_MessageType_C1_Choice_Paging:
		ie.Paging = new(Paging)
		if err = ie.Paging.Decode(r); err != nil {
			return utils.WrapError("Decode Paging", err)
		}
	case PCCH_MessageType_C1_Choice_Spare1:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Spare1", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
