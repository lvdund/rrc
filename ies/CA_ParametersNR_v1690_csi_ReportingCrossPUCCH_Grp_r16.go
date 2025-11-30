package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNR_v1690_csi_ReportingCrossPUCCH_Grp_r16 struct {
	ComputationTimeForA_CSI_r16 CA_ParametersNR_v1690_csi_ReportingCrossPUCCH_Grp_r16_computationTimeForA_CSI_r16  `madatory`
	AdditionalSymbols_r16       *CA_ParametersNR_v1690_csi_ReportingCrossPUCCH_Grp_r16_additionalSymbols_r16       `optional`
	Sp_CSI_ReportingOnPUCCH_r16 *CA_ParametersNR_v1690_csi_ReportingCrossPUCCH_Grp_r16_sp_CSI_ReportingOnPUCCH_r16 `optional`
	Sp_CSI_ReportingOnPUSCH_r16 *CA_ParametersNR_v1690_csi_ReportingCrossPUCCH_Grp_r16_sp_CSI_ReportingOnPUSCH_r16 `optional`
	CarrierTypePairList_r16     []CarrierTypePair_r16                                                              `lb:1,ub:maxCarrierTypePairList_r16,madatory`
}

func (ie *CA_ParametersNR_v1690_csi_ReportingCrossPUCCH_Grp_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.AdditionalSymbols_r16 != nil, ie.Sp_CSI_ReportingOnPUCCH_r16 != nil, ie.Sp_CSI_ReportingOnPUSCH_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.ComputationTimeForA_CSI_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ComputationTimeForA_CSI_r16", err)
	}
	if ie.AdditionalSymbols_r16 != nil {
		if err = ie.AdditionalSymbols_r16.Encode(w); err != nil {
			return utils.WrapError("Encode AdditionalSymbols_r16", err)
		}
	}
	if ie.Sp_CSI_ReportingOnPUCCH_r16 != nil {
		if err = ie.Sp_CSI_ReportingOnPUCCH_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sp_CSI_ReportingOnPUCCH_r16", err)
		}
	}
	if ie.Sp_CSI_ReportingOnPUSCH_r16 != nil {
		if err = ie.Sp_CSI_ReportingOnPUSCH_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sp_CSI_ReportingOnPUSCH_r16", err)
		}
	}
	tmp_CarrierTypePairList_r16 := utils.NewSequence[*CarrierTypePair_r16]([]*CarrierTypePair_r16{}, aper.Constraint{Lb: 1, Ub: maxCarrierTypePairList_r16}, false)
	for _, i := range ie.CarrierTypePairList_r16 {
		tmp_CarrierTypePairList_r16.Value = append(tmp_CarrierTypePairList_r16.Value, &i)
	}
	if err = tmp_CarrierTypePairList_r16.Encode(w); err != nil {
		return utils.WrapError("Encode CarrierTypePairList_r16", err)
	}
	return nil
}

func (ie *CA_ParametersNR_v1690_csi_ReportingCrossPUCCH_Grp_r16) Decode(r *aper.AperReader) error {
	var err error
	var AdditionalSymbols_r16Present bool
	if AdditionalSymbols_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sp_CSI_ReportingOnPUCCH_r16Present bool
	if Sp_CSI_ReportingOnPUCCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sp_CSI_ReportingOnPUSCH_r16Present bool
	if Sp_CSI_ReportingOnPUSCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.ComputationTimeForA_CSI_r16.Decode(r); err != nil {
		return utils.WrapError("Decode ComputationTimeForA_CSI_r16", err)
	}
	if AdditionalSymbols_r16Present {
		ie.AdditionalSymbols_r16 = new(CA_ParametersNR_v1690_csi_ReportingCrossPUCCH_Grp_r16_additionalSymbols_r16)
		if err = ie.AdditionalSymbols_r16.Decode(r); err != nil {
			return utils.WrapError("Decode AdditionalSymbols_r16", err)
		}
	}
	if Sp_CSI_ReportingOnPUCCH_r16Present {
		ie.Sp_CSI_ReportingOnPUCCH_r16 = new(CA_ParametersNR_v1690_csi_ReportingCrossPUCCH_Grp_r16_sp_CSI_ReportingOnPUCCH_r16)
		if err = ie.Sp_CSI_ReportingOnPUCCH_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sp_CSI_ReportingOnPUCCH_r16", err)
		}
	}
	if Sp_CSI_ReportingOnPUSCH_r16Present {
		ie.Sp_CSI_ReportingOnPUSCH_r16 = new(CA_ParametersNR_v1690_csi_ReportingCrossPUCCH_Grp_r16_sp_CSI_ReportingOnPUSCH_r16)
		if err = ie.Sp_CSI_ReportingOnPUSCH_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sp_CSI_ReportingOnPUSCH_r16", err)
		}
	}
	tmp_CarrierTypePairList_r16 := utils.NewSequence[*CarrierTypePair_r16]([]*CarrierTypePair_r16{}, aper.Constraint{Lb: 1, Ub: maxCarrierTypePairList_r16}, false)
	fn_CarrierTypePairList_r16 := func() *CarrierTypePair_r16 {
		return new(CarrierTypePair_r16)
	}
	if err = tmp_CarrierTypePairList_r16.Decode(r, fn_CarrierTypePairList_r16); err != nil {
		return utils.WrapError("Decode CarrierTypePairList_r16", err)
	}
	ie.CarrierTypePairList_r16 = []CarrierTypePair_r16{}
	for _, i := range tmp_CarrierTypePairList_r16.Value {
		ie.CarrierTypePairList_r16 = append(ie.CarrierTypePairList_r16, *i)
	}
	return nil
}
