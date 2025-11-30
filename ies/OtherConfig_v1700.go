package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type OtherConfig_v1700 struct {
	Ul_GapFR2_PreferenceConfig_r17             *OtherConfig_v1700_ul_GapFR2_PreferenceConfig_r17             `optional`
	Musim_GapAssistanceConfig_r17              *MUSIM_GapAssistanceConfig_r17                                `optional,setuprelease`
	Musim_LeaveAssistanceConfig_r17            *MUSIM_LeaveAssistanceConfig_r17                              `optional,setuprelease`
	SuccessHO_Config_r17                       *SuccessHO_Config_r17                                         `optional,setuprelease`
	MaxBW_PreferenceConfigFR2_2_r17            *OtherConfig_v1700_maxBW_PreferenceConfigFR2_2_r17            `optional`
	MaxMIMO_LayerPreferenceConfigFR2_2_r17     *OtherConfig_v1700_maxMIMO_LayerPreferenceConfigFR2_2_r17     `optional`
	MinSchedulingOffsetPreferenceConfigExt_r17 *OtherConfig_v1700_minSchedulingOffsetPreferenceConfigExt_r17 `optional`
	Rlm_RelaxationReportingConfig_r17          *RLM_RelaxationReportingConfig_r17                            `optional,setuprelease`
	Bfd_RelaxationReportingConfig_r17          *BFD_RelaxationReportingConfig_r17                            `optional,setuprelease`
	Scg_DeactivationPreferenceConfig_r17       *SCG_DeactivationPreferenceConfig_r17                         `optional,setuprelease`
	Rrm_MeasRelaxationReportingConfig_r17      *RRM_MeasRelaxationReportingConfig_r17                        `optional,setuprelease`
	PropDelayDiffReportConfig_r17              *PropDelayDiffReportConfig_r17                                `optional,setuprelease`
}

func (ie *OtherConfig_v1700) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Ul_GapFR2_PreferenceConfig_r17 != nil, ie.Musim_GapAssistanceConfig_r17 != nil, ie.Musim_LeaveAssistanceConfig_r17 != nil, ie.SuccessHO_Config_r17 != nil, ie.MaxBW_PreferenceConfigFR2_2_r17 != nil, ie.MaxMIMO_LayerPreferenceConfigFR2_2_r17 != nil, ie.MinSchedulingOffsetPreferenceConfigExt_r17 != nil, ie.Rlm_RelaxationReportingConfig_r17 != nil, ie.Bfd_RelaxationReportingConfig_r17 != nil, ie.Scg_DeactivationPreferenceConfig_r17 != nil, ie.Rrm_MeasRelaxationReportingConfig_r17 != nil, ie.PropDelayDiffReportConfig_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ul_GapFR2_PreferenceConfig_r17 != nil {
		if err = ie.Ul_GapFR2_PreferenceConfig_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ul_GapFR2_PreferenceConfig_r17", err)
		}
	}
	if ie.Musim_GapAssistanceConfig_r17 != nil {
		tmp_Musim_GapAssistanceConfig_r17 := utils.SetupRelease[*MUSIM_GapAssistanceConfig_r17]{
			Setup: ie.Musim_GapAssistanceConfig_r17,
		}
		if err = tmp_Musim_GapAssistanceConfig_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Musim_GapAssistanceConfig_r17", err)
		}
	}
	if ie.Musim_LeaveAssistanceConfig_r17 != nil {
		tmp_Musim_LeaveAssistanceConfig_r17 := utils.SetupRelease[*MUSIM_LeaveAssistanceConfig_r17]{
			Setup: ie.Musim_LeaveAssistanceConfig_r17,
		}
		if err = tmp_Musim_LeaveAssistanceConfig_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Musim_LeaveAssistanceConfig_r17", err)
		}
	}
	if ie.SuccessHO_Config_r17 != nil {
		tmp_SuccessHO_Config_r17 := utils.SetupRelease[*SuccessHO_Config_r17]{
			Setup: ie.SuccessHO_Config_r17,
		}
		if err = tmp_SuccessHO_Config_r17.Encode(w); err != nil {
			return utils.WrapError("Encode SuccessHO_Config_r17", err)
		}
	}
	if ie.MaxBW_PreferenceConfigFR2_2_r17 != nil {
		if err = ie.MaxBW_PreferenceConfigFR2_2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MaxBW_PreferenceConfigFR2_2_r17", err)
		}
	}
	if ie.MaxMIMO_LayerPreferenceConfigFR2_2_r17 != nil {
		if err = ie.MaxMIMO_LayerPreferenceConfigFR2_2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MaxMIMO_LayerPreferenceConfigFR2_2_r17", err)
		}
	}
	if ie.MinSchedulingOffsetPreferenceConfigExt_r17 != nil {
		if err = ie.MinSchedulingOffsetPreferenceConfigExt_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MinSchedulingOffsetPreferenceConfigExt_r17", err)
		}
	}
	if ie.Rlm_RelaxationReportingConfig_r17 != nil {
		tmp_Rlm_RelaxationReportingConfig_r17 := utils.SetupRelease[*RLM_RelaxationReportingConfig_r17]{
			Setup: ie.Rlm_RelaxationReportingConfig_r17,
		}
		if err = tmp_Rlm_RelaxationReportingConfig_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Rlm_RelaxationReportingConfig_r17", err)
		}
	}
	if ie.Bfd_RelaxationReportingConfig_r17 != nil {
		tmp_Bfd_RelaxationReportingConfig_r17 := utils.SetupRelease[*BFD_RelaxationReportingConfig_r17]{
			Setup: ie.Bfd_RelaxationReportingConfig_r17,
		}
		if err = tmp_Bfd_RelaxationReportingConfig_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Bfd_RelaxationReportingConfig_r17", err)
		}
	}
	if ie.Scg_DeactivationPreferenceConfig_r17 != nil {
		tmp_Scg_DeactivationPreferenceConfig_r17 := utils.SetupRelease[*SCG_DeactivationPreferenceConfig_r17]{
			Setup: ie.Scg_DeactivationPreferenceConfig_r17,
		}
		if err = tmp_Scg_DeactivationPreferenceConfig_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Scg_DeactivationPreferenceConfig_r17", err)
		}
	}
	if ie.Rrm_MeasRelaxationReportingConfig_r17 != nil {
		tmp_Rrm_MeasRelaxationReportingConfig_r17 := utils.SetupRelease[*RRM_MeasRelaxationReportingConfig_r17]{
			Setup: ie.Rrm_MeasRelaxationReportingConfig_r17,
		}
		if err = tmp_Rrm_MeasRelaxationReportingConfig_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Rrm_MeasRelaxationReportingConfig_r17", err)
		}
	}
	if ie.PropDelayDiffReportConfig_r17 != nil {
		tmp_PropDelayDiffReportConfig_r17 := utils.SetupRelease[*PropDelayDiffReportConfig_r17]{
			Setup: ie.PropDelayDiffReportConfig_r17,
		}
		if err = tmp_PropDelayDiffReportConfig_r17.Encode(w); err != nil {
			return utils.WrapError("Encode PropDelayDiffReportConfig_r17", err)
		}
	}
	return nil
}

func (ie *OtherConfig_v1700) Decode(r *aper.AperReader) error {
	var err error
	var Ul_GapFR2_PreferenceConfig_r17Present bool
	if Ul_GapFR2_PreferenceConfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Musim_GapAssistanceConfig_r17Present bool
	if Musim_GapAssistanceConfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Musim_LeaveAssistanceConfig_r17Present bool
	if Musim_LeaveAssistanceConfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SuccessHO_Config_r17Present bool
	if SuccessHO_Config_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxBW_PreferenceConfigFR2_2_r17Present bool
	if MaxBW_PreferenceConfigFR2_2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxMIMO_LayerPreferenceConfigFR2_2_r17Present bool
	if MaxMIMO_LayerPreferenceConfigFR2_2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MinSchedulingOffsetPreferenceConfigExt_r17Present bool
	if MinSchedulingOffsetPreferenceConfigExt_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Rlm_RelaxationReportingConfig_r17Present bool
	if Rlm_RelaxationReportingConfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Bfd_RelaxationReportingConfig_r17Present bool
	if Bfd_RelaxationReportingConfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Scg_DeactivationPreferenceConfig_r17Present bool
	if Scg_DeactivationPreferenceConfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Rrm_MeasRelaxationReportingConfig_r17Present bool
	if Rrm_MeasRelaxationReportingConfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var PropDelayDiffReportConfig_r17Present bool
	if PropDelayDiffReportConfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Ul_GapFR2_PreferenceConfig_r17Present {
		ie.Ul_GapFR2_PreferenceConfig_r17 = new(OtherConfig_v1700_ul_GapFR2_PreferenceConfig_r17)
		if err = ie.Ul_GapFR2_PreferenceConfig_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ul_GapFR2_PreferenceConfig_r17", err)
		}
	}
	if Musim_GapAssistanceConfig_r17Present {
		tmp_Musim_GapAssistanceConfig_r17 := utils.SetupRelease[*MUSIM_GapAssistanceConfig_r17]{}
		if err = tmp_Musim_GapAssistanceConfig_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Musim_GapAssistanceConfig_r17", err)
		}
		ie.Musim_GapAssistanceConfig_r17 = tmp_Musim_GapAssistanceConfig_r17.Setup
	}
	if Musim_LeaveAssistanceConfig_r17Present {
		tmp_Musim_LeaveAssistanceConfig_r17 := utils.SetupRelease[*MUSIM_LeaveAssistanceConfig_r17]{}
		if err = tmp_Musim_LeaveAssistanceConfig_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Musim_LeaveAssistanceConfig_r17", err)
		}
		ie.Musim_LeaveAssistanceConfig_r17 = tmp_Musim_LeaveAssistanceConfig_r17.Setup
	}
	if SuccessHO_Config_r17Present {
		tmp_SuccessHO_Config_r17 := utils.SetupRelease[*SuccessHO_Config_r17]{}
		if err = tmp_SuccessHO_Config_r17.Decode(r); err != nil {
			return utils.WrapError("Decode SuccessHO_Config_r17", err)
		}
		ie.SuccessHO_Config_r17 = tmp_SuccessHO_Config_r17.Setup
	}
	if MaxBW_PreferenceConfigFR2_2_r17Present {
		ie.MaxBW_PreferenceConfigFR2_2_r17 = new(OtherConfig_v1700_maxBW_PreferenceConfigFR2_2_r17)
		if err = ie.MaxBW_PreferenceConfigFR2_2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MaxBW_PreferenceConfigFR2_2_r17", err)
		}
	}
	if MaxMIMO_LayerPreferenceConfigFR2_2_r17Present {
		ie.MaxMIMO_LayerPreferenceConfigFR2_2_r17 = new(OtherConfig_v1700_maxMIMO_LayerPreferenceConfigFR2_2_r17)
		if err = ie.MaxMIMO_LayerPreferenceConfigFR2_2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MaxMIMO_LayerPreferenceConfigFR2_2_r17", err)
		}
	}
	if MinSchedulingOffsetPreferenceConfigExt_r17Present {
		ie.MinSchedulingOffsetPreferenceConfigExt_r17 = new(OtherConfig_v1700_minSchedulingOffsetPreferenceConfigExt_r17)
		if err = ie.MinSchedulingOffsetPreferenceConfigExt_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MinSchedulingOffsetPreferenceConfigExt_r17", err)
		}
	}
	if Rlm_RelaxationReportingConfig_r17Present {
		tmp_Rlm_RelaxationReportingConfig_r17 := utils.SetupRelease[*RLM_RelaxationReportingConfig_r17]{}
		if err = tmp_Rlm_RelaxationReportingConfig_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Rlm_RelaxationReportingConfig_r17", err)
		}
		ie.Rlm_RelaxationReportingConfig_r17 = tmp_Rlm_RelaxationReportingConfig_r17.Setup
	}
	if Bfd_RelaxationReportingConfig_r17Present {
		tmp_Bfd_RelaxationReportingConfig_r17 := utils.SetupRelease[*BFD_RelaxationReportingConfig_r17]{}
		if err = tmp_Bfd_RelaxationReportingConfig_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Bfd_RelaxationReportingConfig_r17", err)
		}
		ie.Bfd_RelaxationReportingConfig_r17 = tmp_Bfd_RelaxationReportingConfig_r17.Setup
	}
	if Scg_DeactivationPreferenceConfig_r17Present {
		tmp_Scg_DeactivationPreferenceConfig_r17 := utils.SetupRelease[*SCG_DeactivationPreferenceConfig_r17]{}
		if err = tmp_Scg_DeactivationPreferenceConfig_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Scg_DeactivationPreferenceConfig_r17", err)
		}
		ie.Scg_DeactivationPreferenceConfig_r17 = tmp_Scg_DeactivationPreferenceConfig_r17.Setup
	}
	if Rrm_MeasRelaxationReportingConfig_r17Present {
		tmp_Rrm_MeasRelaxationReportingConfig_r17 := utils.SetupRelease[*RRM_MeasRelaxationReportingConfig_r17]{}
		if err = tmp_Rrm_MeasRelaxationReportingConfig_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Rrm_MeasRelaxationReportingConfig_r17", err)
		}
		ie.Rrm_MeasRelaxationReportingConfig_r17 = tmp_Rrm_MeasRelaxationReportingConfig_r17.Setup
	}
	if PropDelayDiffReportConfig_r17Present {
		tmp_PropDelayDiffReportConfig_r17 := utils.SetupRelease[*PropDelayDiffReportConfig_r17]{}
		if err = tmp_PropDelayDiffReportConfig_r17.Decode(r); err != nil {
			return utils.WrapError("Decode PropDelayDiffReportConfig_r17", err)
		}
		ie.PropDelayDiffReportConfig_r17 = tmp_PropDelayDiffReportConfig_r17.Setup
	}
	return nil
}
