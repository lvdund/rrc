package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SIB1_v1630_IEs_uac_BarringInfo_v1630 struct {
	Uac_AC1_SelectAssistInfo_r16 []UAC_AC1_SelectAssistInfo_r16 `lb:2,ub:maxPLMN,madatory`
}

func (ie *SIB1_v1630_IEs_uac_BarringInfo_v1630) Encode(w *aper.AperWriter) error {
	var err error
	tmp_Uac_AC1_SelectAssistInfo_r16 := utils.NewSequence[*UAC_AC1_SelectAssistInfo_r16]([]*UAC_AC1_SelectAssistInfo_r16{}, aper.Constraint{Lb: 2, Ub: maxPLMN}, false)
	for _, i := range ie.Uac_AC1_SelectAssistInfo_r16 {
		tmp_Uac_AC1_SelectAssistInfo_r16.Value = append(tmp_Uac_AC1_SelectAssistInfo_r16.Value, &i)
	}
	if err = tmp_Uac_AC1_SelectAssistInfo_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Uac_AC1_SelectAssistInfo_r16", err)
	}
	return nil
}

func (ie *SIB1_v1630_IEs_uac_BarringInfo_v1630) Decode(r *aper.AperReader) error {
	var err error
	tmp_Uac_AC1_SelectAssistInfo_r16 := utils.NewSequence[*UAC_AC1_SelectAssistInfo_r16]([]*UAC_AC1_SelectAssistInfo_r16{}, aper.Constraint{Lb: 2, Ub: maxPLMN}, false)
	fn_Uac_AC1_SelectAssistInfo_r16 := func() *UAC_AC1_SelectAssistInfo_r16 {
		return new(UAC_AC1_SelectAssistInfo_r16)
	}
	if err = tmp_Uac_AC1_SelectAssistInfo_r16.Decode(r, fn_Uac_AC1_SelectAssistInfo_r16); err != nil {
		return utils.WrapError("Decode Uac_AC1_SelectAssistInfo_r16", err)
	}
	ie.Uac_AC1_SelectAssistInfo_r16 = []UAC_AC1_SelectAssistInfo_r16{}
	for _, i := range tmp_Uac_AC1_SelectAssistInfo_r16.Value {
		ie.Uac_AC1_SelectAssistInfo_r16 = append(ie.Uac_AC1_SelectAssistInfo_r16, *i)
	}
	return nil
}
