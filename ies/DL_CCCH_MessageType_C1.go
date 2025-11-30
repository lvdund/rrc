package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	DL_CCCH_MessageType_C1_Choice_nothing uint64 = iota
	DL_CCCH_MessageType_C1_Choice_RrcReject
	DL_CCCH_MessageType_C1_Choice_RrcSetup
	DL_CCCH_MessageType_C1_Choice_Spare2
	DL_CCCH_MessageType_C1_Choice_Spare1
)

type DL_CCCH_MessageType_C1 struct {
	Choice    uint64
	RrcReject *RRCReject
	RrcSetup  *RRCSetup
	Spare2    aper.NULL `madatory`
	Spare1    aper.NULL `madatory`
}

func (ie *DL_CCCH_MessageType_C1) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case DL_CCCH_MessageType_C1_Choice_RrcReject:
		if err = ie.RrcReject.Encode(w); err != nil {
			err = utils.WrapError("Encode RrcReject", err)
		}
	case DL_CCCH_MessageType_C1_Choice_RrcSetup:
		if err = ie.RrcSetup.Encode(w); err != nil {
			err = utils.WrapError("Encode RrcSetup", err)
		}
	case DL_CCCH_MessageType_C1_Choice_Spare2:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Spare2", err)
		}
	case DL_CCCH_MessageType_C1_Choice_Spare1:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Spare1", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *DL_CCCH_MessageType_C1) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case DL_CCCH_MessageType_C1_Choice_RrcReject:
		ie.RrcReject = new(RRCReject)
		if err = ie.RrcReject.Decode(r); err != nil {
			return utils.WrapError("Decode RrcReject", err)
		}
	case DL_CCCH_MessageType_C1_Choice_RrcSetup:
		ie.RrcSetup = new(RRCSetup)
		if err = ie.RrcSetup.Decode(r); err != nil {
			return utils.WrapError("Decode RrcSetup", err)
		}
	case DL_CCCH_MessageType_C1_Choice_Spare2:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Spare2", err)
		}
	case DL_CCCH_MessageType_C1_Choice_Spare1:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Spare1", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
