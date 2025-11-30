package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNR_v1640 struct {
	UplinkTxDC_TwoCarrierReport_r16                            *CA_ParametersNR_v1640_uplinkTxDC_TwoCarrierReport_r16                            `optional`
	MaxUpTo3Diff_NumerologiesConfigSinglePUCCH_grp_r16         *PUCCH_Grp_CarrierTypes_r16                                                       `optional`
	MaxUpTo4Diff_NumerologiesConfigSinglePUCCH_grp_r16         *PUCCH_Grp_CarrierTypes_r16                                                       `optional`
	TwoPUCCH_Grp_ConfigurationsList_r16                        []TwoPUCCH_Grp_Configurations_r16                                                 `lb:1,ub:maxTwoPUCCH_Grp_ConfigList_r16,optional`
	DiffNumerologyAcrossPUCCH_Group_CarrierTypes_r16           *CA_ParametersNR_v1640_diffNumerologyAcrossPUCCH_Group_CarrierTypes_r16           `optional`
	DiffNumerologyWithinPUCCH_GroupSmallerSCS_CarrierTypes_r16 *CA_ParametersNR_v1640_diffNumerologyWithinPUCCH_GroupSmallerSCS_CarrierTypes_r16 `optional`
	DiffNumerologyWithinPUCCH_GroupLargerSCS_CarrierTypes_r16  *CA_ParametersNR_v1640_diffNumerologyWithinPUCCH_GroupLargerSCS_CarrierTypes_r16  `optional`
	Pdcch_MonitoringCA_NonAlignedSpan_r16                      *int64                                                                            `lb:2,ub:16,optional`
	Pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_r16            *CA_ParametersNR_v1640_pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_r16            `lb:1,ub:15,optional`
}

func (ie *CA_ParametersNR_v1640) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.UplinkTxDC_TwoCarrierReport_r16 != nil, ie.MaxUpTo3Diff_NumerologiesConfigSinglePUCCH_grp_r16 != nil, ie.MaxUpTo4Diff_NumerologiesConfigSinglePUCCH_grp_r16 != nil, len(ie.TwoPUCCH_Grp_ConfigurationsList_r16) > 0, ie.DiffNumerologyAcrossPUCCH_Group_CarrierTypes_r16 != nil, ie.DiffNumerologyWithinPUCCH_GroupSmallerSCS_CarrierTypes_r16 != nil, ie.DiffNumerologyWithinPUCCH_GroupLargerSCS_CarrierTypes_r16 != nil, ie.Pdcch_MonitoringCA_NonAlignedSpan_r16 != nil, ie.Pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.UplinkTxDC_TwoCarrierReport_r16 != nil {
		if err = ie.UplinkTxDC_TwoCarrierReport_r16.Encode(w); err != nil {
			return utils.WrapError("Encode UplinkTxDC_TwoCarrierReport_r16", err)
		}
	}
	if ie.MaxUpTo3Diff_NumerologiesConfigSinglePUCCH_grp_r16 != nil {
		if err = ie.MaxUpTo3Diff_NumerologiesConfigSinglePUCCH_grp_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MaxUpTo3Diff_NumerologiesConfigSinglePUCCH_grp_r16", err)
		}
	}
	if ie.MaxUpTo4Diff_NumerologiesConfigSinglePUCCH_grp_r16 != nil {
		if err = ie.MaxUpTo4Diff_NumerologiesConfigSinglePUCCH_grp_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MaxUpTo4Diff_NumerologiesConfigSinglePUCCH_grp_r16", err)
		}
	}
	if len(ie.TwoPUCCH_Grp_ConfigurationsList_r16) > 0 {
		tmp_TwoPUCCH_Grp_ConfigurationsList_r16 := utils.NewSequence[*TwoPUCCH_Grp_Configurations_r16]([]*TwoPUCCH_Grp_Configurations_r16{}, aper.Constraint{Lb: 1, Ub: maxTwoPUCCH_Grp_ConfigList_r16}, false)
		for _, i := range ie.TwoPUCCH_Grp_ConfigurationsList_r16 {
			tmp_TwoPUCCH_Grp_ConfigurationsList_r16.Value = append(tmp_TwoPUCCH_Grp_ConfigurationsList_r16.Value, &i)
		}
		if err = tmp_TwoPUCCH_Grp_ConfigurationsList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode TwoPUCCH_Grp_ConfigurationsList_r16", err)
		}
	}
	if ie.DiffNumerologyAcrossPUCCH_Group_CarrierTypes_r16 != nil {
		if err = ie.DiffNumerologyAcrossPUCCH_Group_CarrierTypes_r16.Encode(w); err != nil {
			return utils.WrapError("Encode DiffNumerologyAcrossPUCCH_Group_CarrierTypes_r16", err)
		}
	}
	if ie.DiffNumerologyWithinPUCCH_GroupSmallerSCS_CarrierTypes_r16 != nil {
		if err = ie.DiffNumerologyWithinPUCCH_GroupSmallerSCS_CarrierTypes_r16.Encode(w); err != nil {
			return utils.WrapError("Encode DiffNumerologyWithinPUCCH_GroupSmallerSCS_CarrierTypes_r16", err)
		}
	}
	if ie.DiffNumerologyWithinPUCCH_GroupLargerSCS_CarrierTypes_r16 != nil {
		if err = ie.DiffNumerologyWithinPUCCH_GroupLargerSCS_CarrierTypes_r16.Encode(w); err != nil {
			return utils.WrapError("Encode DiffNumerologyWithinPUCCH_GroupLargerSCS_CarrierTypes_r16", err)
		}
	}
	if ie.Pdcch_MonitoringCA_NonAlignedSpan_r16 != nil {
		if err = w.WriteInteger(*ie.Pdcch_MonitoringCA_NonAlignedSpan_r16, &aper.Constraint{Lb: 2, Ub: 16}, false); err != nil {
			return utils.WrapError("Encode Pdcch_MonitoringCA_NonAlignedSpan_r16", err)
		}
	}
	if ie.Pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_r16 != nil {
		if err = ie.Pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_r16", err)
		}
	}
	return nil
}

func (ie *CA_ParametersNR_v1640) Decode(r *aper.AperReader) error {
	var err error
	var UplinkTxDC_TwoCarrierReport_r16Present bool
	if UplinkTxDC_TwoCarrierReport_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxUpTo3Diff_NumerologiesConfigSinglePUCCH_grp_r16Present bool
	if MaxUpTo3Diff_NumerologiesConfigSinglePUCCH_grp_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxUpTo4Diff_NumerologiesConfigSinglePUCCH_grp_r16Present bool
	if MaxUpTo4Diff_NumerologiesConfigSinglePUCCH_grp_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var TwoPUCCH_Grp_ConfigurationsList_r16Present bool
	if TwoPUCCH_Grp_ConfigurationsList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var DiffNumerologyAcrossPUCCH_Group_CarrierTypes_r16Present bool
	if DiffNumerologyAcrossPUCCH_Group_CarrierTypes_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var DiffNumerologyWithinPUCCH_GroupSmallerSCS_CarrierTypes_r16Present bool
	if DiffNumerologyWithinPUCCH_GroupSmallerSCS_CarrierTypes_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var DiffNumerologyWithinPUCCH_GroupLargerSCS_CarrierTypes_r16Present bool
	if DiffNumerologyWithinPUCCH_GroupLargerSCS_CarrierTypes_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdcch_MonitoringCA_NonAlignedSpan_r16Present bool
	if Pdcch_MonitoringCA_NonAlignedSpan_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_r16Present bool
	if Pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if UplinkTxDC_TwoCarrierReport_r16Present {
		ie.UplinkTxDC_TwoCarrierReport_r16 = new(CA_ParametersNR_v1640_uplinkTxDC_TwoCarrierReport_r16)
		if err = ie.UplinkTxDC_TwoCarrierReport_r16.Decode(r); err != nil {
			return utils.WrapError("Decode UplinkTxDC_TwoCarrierReport_r16", err)
		}
	}
	if MaxUpTo3Diff_NumerologiesConfigSinglePUCCH_grp_r16Present {
		ie.MaxUpTo3Diff_NumerologiesConfigSinglePUCCH_grp_r16 = new(PUCCH_Grp_CarrierTypes_r16)
		if err = ie.MaxUpTo3Diff_NumerologiesConfigSinglePUCCH_grp_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MaxUpTo3Diff_NumerologiesConfigSinglePUCCH_grp_r16", err)
		}
	}
	if MaxUpTo4Diff_NumerologiesConfigSinglePUCCH_grp_r16Present {
		ie.MaxUpTo4Diff_NumerologiesConfigSinglePUCCH_grp_r16 = new(PUCCH_Grp_CarrierTypes_r16)
		if err = ie.MaxUpTo4Diff_NumerologiesConfigSinglePUCCH_grp_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MaxUpTo4Diff_NumerologiesConfigSinglePUCCH_grp_r16", err)
		}
	}
	if TwoPUCCH_Grp_ConfigurationsList_r16Present {
		tmp_TwoPUCCH_Grp_ConfigurationsList_r16 := utils.NewSequence[*TwoPUCCH_Grp_Configurations_r16]([]*TwoPUCCH_Grp_Configurations_r16{}, aper.Constraint{Lb: 1, Ub: maxTwoPUCCH_Grp_ConfigList_r16}, false)
		fn_TwoPUCCH_Grp_ConfigurationsList_r16 := func() *TwoPUCCH_Grp_Configurations_r16 {
			return new(TwoPUCCH_Grp_Configurations_r16)
		}
		if err = tmp_TwoPUCCH_Grp_ConfigurationsList_r16.Decode(r, fn_TwoPUCCH_Grp_ConfigurationsList_r16); err != nil {
			return utils.WrapError("Decode TwoPUCCH_Grp_ConfigurationsList_r16", err)
		}
		ie.TwoPUCCH_Grp_ConfigurationsList_r16 = []TwoPUCCH_Grp_Configurations_r16{}
		for _, i := range tmp_TwoPUCCH_Grp_ConfigurationsList_r16.Value {
			ie.TwoPUCCH_Grp_ConfigurationsList_r16 = append(ie.TwoPUCCH_Grp_ConfigurationsList_r16, *i)
		}
	}
	if DiffNumerologyAcrossPUCCH_Group_CarrierTypes_r16Present {
		ie.DiffNumerologyAcrossPUCCH_Group_CarrierTypes_r16 = new(CA_ParametersNR_v1640_diffNumerologyAcrossPUCCH_Group_CarrierTypes_r16)
		if err = ie.DiffNumerologyAcrossPUCCH_Group_CarrierTypes_r16.Decode(r); err != nil {
			return utils.WrapError("Decode DiffNumerologyAcrossPUCCH_Group_CarrierTypes_r16", err)
		}
	}
	if DiffNumerologyWithinPUCCH_GroupSmallerSCS_CarrierTypes_r16Present {
		ie.DiffNumerologyWithinPUCCH_GroupSmallerSCS_CarrierTypes_r16 = new(CA_ParametersNR_v1640_diffNumerologyWithinPUCCH_GroupSmallerSCS_CarrierTypes_r16)
		if err = ie.DiffNumerologyWithinPUCCH_GroupSmallerSCS_CarrierTypes_r16.Decode(r); err != nil {
			return utils.WrapError("Decode DiffNumerologyWithinPUCCH_GroupSmallerSCS_CarrierTypes_r16", err)
		}
	}
	if DiffNumerologyWithinPUCCH_GroupLargerSCS_CarrierTypes_r16Present {
		ie.DiffNumerologyWithinPUCCH_GroupLargerSCS_CarrierTypes_r16 = new(CA_ParametersNR_v1640_diffNumerologyWithinPUCCH_GroupLargerSCS_CarrierTypes_r16)
		if err = ie.DiffNumerologyWithinPUCCH_GroupLargerSCS_CarrierTypes_r16.Decode(r); err != nil {
			return utils.WrapError("Decode DiffNumerologyWithinPUCCH_GroupLargerSCS_CarrierTypes_r16", err)
		}
	}
	if Pdcch_MonitoringCA_NonAlignedSpan_r16Present {
		var tmp_int_Pdcch_MonitoringCA_NonAlignedSpan_r16 int64
		if tmp_int_Pdcch_MonitoringCA_NonAlignedSpan_r16, err = r.ReadInteger(&aper.Constraint{Lb: 2, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode Pdcch_MonitoringCA_NonAlignedSpan_r16", err)
		}
		ie.Pdcch_MonitoringCA_NonAlignedSpan_r16 = &tmp_int_Pdcch_MonitoringCA_NonAlignedSpan_r16
	}
	if Pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_r16Present {
		ie.Pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_r16 = new(CA_ParametersNR_v1640_pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_r16)
		if err = ie.Pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_r16", err)
		}
	}
	return nil
}
