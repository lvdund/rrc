package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ReportConfigToAddMod_reportConfig_Choice_nothing uint64 = iota
	ReportConfigToAddMod_reportConfig_Choice_ReportConfigNR
	ReportConfigToAddMod_reportConfig_Choice_ReportConfigInterRAT
	ReportConfigToAddMod_reportConfig_Choice_ReportConfigNR_SL_r16
)

type ReportConfigToAddMod_reportConfig struct {
	Choice                uint64
	ReportConfigNR        *ReportConfigNR
	ReportConfigInterRAT  *ReportConfigInterRAT
	ReportConfigNR_SL_r16 *ReportConfigNR_SL_r16
}

func (ie *ReportConfigToAddMod_reportConfig) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ReportConfigToAddMod_reportConfig_Choice_ReportConfigNR:
		if err = ie.ReportConfigNR.Encode(w); err != nil {
			err = utils.WrapError("Encode ReportConfigNR", err)
		}
	case ReportConfigToAddMod_reportConfig_Choice_ReportConfigInterRAT:
		if err = ie.ReportConfigInterRAT.Encode(w); err != nil {
			err = utils.WrapError("Encode ReportConfigInterRAT", err)
		}
	case ReportConfigToAddMod_reportConfig_Choice_ReportConfigNR_SL_r16:
		if err = ie.ReportConfigNR_SL_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode ReportConfigNR_SL_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *ReportConfigToAddMod_reportConfig) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ReportConfigToAddMod_reportConfig_Choice_ReportConfigNR:
		ie.ReportConfigNR = new(ReportConfigNR)
		if err = ie.ReportConfigNR.Decode(r); err != nil {
			return utils.WrapError("Decode ReportConfigNR", err)
		}
	case ReportConfigToAddMod_reportConfig_Choice_ReportConfigInterRAT:
		ie.ReportConfigInterRAT = new(ReportConfigInterRAT)
		if err = ie.ReportConfigInterRAT.Decode(r); err != nil {
			return utils.WrapError("Decode ReportConfigInterRAT", err)
		}
	case ReportConfigToAddMod_reportConfig_Choice_ReportConfigNR_SL_r16:
		ie.ReportConfigNR_SL_r16 = new(ReportConfigNR_SL_r16)
		if err = ie.ReportConfigNR_SL_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ReportConfigNR_SL_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
