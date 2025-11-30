package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CG_ConfigInfo_CriticalExtensions_C1_Choice_nothing uint64 = iota
	CG_ConfigInfo_CriticalExtensions_C1_Choice_Cg_ConfigInfo
	CG_ConfigInfo_CriticalExtensions_C1_Choice_Spare3
	CG_ConfigInfo_CriticalExtensions_C1_Choice_Spare2
	CG_ConfigInfo_CriticalExtensions_C1_Choice_Spare1
)

type CG_ConfigInfo_CriticalExtensions_C1 struct {
	Choice        uint64
	Cg_ConfigInfo *CG_ConfigInfo_IEs
	Spare3        aper.NULL `madatory`
	Spare2        aper.NULL `madatory`
	Spare1        aper.NULL `madatory`
}

func (ie *CG_ConfigInfo_CriticalExtensions_C1) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CG_ConfigInfo_CriticalExtensions_C1_Choice_Cg_ConfigInfo:
		if err = ie.Cg_ConfigInfo.Encode(w); err != nil {
			err = utils.WrapError("Encode Cg_ConfigInfo", err)
		}
	case CG_ConfigInfo_CriticalExtensions_C1_Choice_Spare3:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Spare3", err)
		}
	case CG_ConfigInfo_CriticalExtensions_C1_Choice_Spare2:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Spare2", err)
		}
	case CG_ConfigInfo_CriticalExtensions_C1_Choice_Spare1:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Spare1", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CG_ConfigInfo_CriticalExtensions_C1) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CG_ConfigInfo_CriticalExtensions_C1_Choice_Cg_ConfigInfo:
		ie.Cg_ConfigInfo = new(CG_ConfigInfo_IEs)
		if err = ie.Cg_ConfigInfo.Decode(r); err != nil {
			return utils.WrapError("Decode Cg_ConfigInfo", err)
		}
	case CG_ConfigInfo_CriticalExtensions_C1_Choice_Spare3:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Spare3", err)
		}
	case CG_ConfigInfo_CriticalExtensions_C1_Choice_Spare2:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Spare2", err)
		}
	case CG_ConfigInfo_CriticalExtensions_C1_Choice_Spare1:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Spare1", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
