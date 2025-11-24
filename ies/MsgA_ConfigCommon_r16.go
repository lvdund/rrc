package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MsgA_ConfigCommon_r16 struct {
	Rach_ConfigCommonTwoStepRA_r16 RACH_ConfigCommonTwoStepRA_r16 `madatory`
	MsgA_PUSCH_Config_r16          *MsgA_PUSCH_Config_r16         `optional`
}

func (ie *MsgA_ConfigCommon_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.MsgA_PUSCH_Config_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Rach_ConfigCommonTwoStepRA_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Rach_ConfigCommonTwoStepRA_r16", err)
	}
	if ie.MsgA_PUSCH_Config_r16 != nil {
		if err = ie.MsgA_PUSCH_Config_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MsgA_PUSCH_Config_r16", err)
		}
	}
	return nil
}

func (ie *MsgA_ConfigCommon_r16) Decode(r *uper.UperReader) error {
	var err error
	var MsgA_PUSCH_Config_r16Present bool
	if MsgA_PUSCH_Config_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Rach_ConfigCommonTwoStepRA_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Rach_ConfigCommonTwoStepRA_r16", err)
	}
	if MsgA_PUSCH_Config_r16Present {
		ie.MsgA_PUSCH_Config_r16 = new(MsgA_PUSCH_Config_r16)
		if err = ie.MsgA_PUSCH_Config_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MsgA_PUSCH_Config_r16", err)
		}
	}
	return nil
}
