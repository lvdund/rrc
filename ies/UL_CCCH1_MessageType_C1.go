package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	UL_CCCH1_MessageType_C1_Choice_nothing uint64 = iota
	UL_CCCH1_MessageType_C1_Choice_RrcResumeRequest1
	UL_CCCH1_MessageType_C1_Choice_Spare3
	UL_CCCH1_MessageType_C1_Choice_Spare2
	UL_CCCH1_MessageType_C1_Choice_Spare1
)

type UL_CCCH1_MessageType_C1 struct {
	Choice            uint64
	RrcResumeRequest1 *RRCResumeRequest1
	Spare3            aper.NULL `madatory`
	Spare2            aper.NULL `madatory`
	Spare1            aper.NULL `madatory`
}

func (ie *UL_CCCH1_MessageType_C1) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UL_CCCH1_MessageType_C1_Choice_RrcResumeRequest1:
		if err = ie.RrcResumeRequest1.Encode(w); err != nil {
			err = utils.WrapError("Encode RrcResumeRequest1", err)
		}
	case UL_CCCH1_MessageType_C1_Choice_Spare3:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Spare3", err)
		}
	case UL_CCCH1_MessageType_C1_Choice_Spare2:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Spare2", err)
		}
	case UL_CCCH1_MessageType_C1_Choice_Spare1:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Spare1", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *UL_CCCH1_MessageType_C1) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UL_CCCH1_MessageType_C1_Choice_RrcResumeRequest1:
		ie.RrcResumeRequest1 = new(RRCResumeRequest1)
		if err = ie.RrcResumeRequest1.Decode(r); err != nil {
			return utils.WrapError("Decode RrcResumeRequest1", err)
		}
	case UL_CCCH1_MessageType_C1_Choice_Spare3:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Spare3", err)
		}
	case UL_CCCH1_MessageType_C1_Choice_Spare2:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Spare2", err)
		}
	case UL_CCCH1_MessageType_C1_Choice_Spare1:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Spare1", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
