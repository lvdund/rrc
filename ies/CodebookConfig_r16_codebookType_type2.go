package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CodebookConfig_r16_codebookType_type2 struct {
	SubType                                CodebookConfig_r16_codebookType_type2_subType `lb:16,ub:16,madatory`
	NumberOfPMI_SubbandsPerCQI_Subband_r16 int64                                         `lb:1,ub:2,madatory`
	ParamCombination_r16                   int64                                         `lb:1,ub:8,madatory`
}

func (ie *CodebookConfig_r16_codebookType_type2) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.SubType.Encode(w); err != nil {
		return utils.WrapError("Encode SubType", err)
	}
	if err = w.WriteInteger(ie.NumberOfPMI_SubbandsPerCQI_Subband_r16, &aper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
		return utils.WrapError("WriteInteger NumberOfPMI_SubbandsPerCQI_Subband_r16", err)
	}
	if err = w.WriteInteger(ie.ParamCombination_r16, &aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger ParamCombination_r16", err)
	}
	return nil
}

func (ie *CodebookConfig_r16_codebookType_type2) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.SubType.Decode(r); err != nil {
		return utils.WrapError("Decode SubType", err)
	}
	var tmp_int_NumberOfPMI_SubbandsPerCQI_Subband_r16 int64
	if tmp_int_NumberOfPMI_SubbandsPerCQI_Subband_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
		return utils.WrapError("ReadInteger NumberOfPMI_SubbandsPerCQI_Subband_r16", err)
	}
	ie.NumberOfPMI_SubbandsPerCQI_Subband_r16 = tmp_int_NumberOfPMI_SubbandsPerCQI_Subband_r16
	var tmp_int_ParamCombination_r16 int64
	if tmp_int_ParamCombination_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger ParamCombination_r16", err)
	}
	ie.ParamCombination_r16 = tmp_int_ParamCombination_r16
	return nil
}
