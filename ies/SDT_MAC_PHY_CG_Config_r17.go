package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SDT_MAC_PHY_CG_Config_r17 struct {
	Cg_SDT_ConfigLCH_RestrictionToAddModList_r17  []CG_SDT_ConfigLCH_Restriction_r17 `lb:1,ub:maxLC_ID,optional`
	Cg_SDT_ConfigLCH_RestrictionToReleaseList_r17 []LogicalChannelIdentity           `lb:1,ub:maxLC_ID,optional`
	Cg_SDT_ConfigInitialBWP_NUL_r17               *BWP_UplinkDedicatedSDT_r17        `optional,setuprelease`
	Cg_SDT_ConfigInitialBWP_SUL_r17               *BWP_UplinkDedicatedSDT_r17        `optional,setuprelease`
	Cg_SDT_ConfigInitialBWP_DL_r17                *BWP_DownlinkDedicatedSDT_r17      `optional`
	Cg_SDT_TimeAlignmentTimer_r17                 *TimeAlignmentTimer                `optional`
	Cg_SDT_RSRP_ThresholdSSB_r17                  *RSRP_Range                        `optional`
	Cg_SDT_TA_ValidationConfig_r17                *CG_SDT_TA_ValidationConfig_r17    `optional,setuprelease`
	Cg_SDT_CS_RNTI_r17                            *RNTI_Value                        `optional`
}

func (ie *SDT_MAC_PHY_CG_Config_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{len(ie.Cg_SDT_ConfigLCH_RestrictionToAddModList_r17) > 0, len(ie.Cg_SDT_ConfigLCH_RestrictionToReleaseList_r17) > 0, ie.Cg_SDT_ConfigInitialBWP_NUL_r17 != nil, ie.Cg_SDT_ConfigInitialBWP_SUL_r17 != nil, ie.Cg_SDT_ConfigInitialBWP_DL_r17 != nil, ie.Cg_SDT_TimeAlignmentTimer_r17 != nil, ie.Cg_SDT_RSRP_ThresholdSSB_r17 != nil, ie.Cg_SDT_TA_ValidationConfig_r17 != nil, ie.Cg_SDT_CS_RNTI_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.Cg_SDT_ConfigLCH_RestrictionToAddModList_r17) > 0 {
		tmp_Cg_SDT_ConfigLCH_RestrictionToAddModList_r17 := utils.NewSequence[*CG_SDT_ConfigLCH_Restriction_r17]([]*CG_SDT_ConfigLCH_Restriction_r17{}, aper.Constraint{Lb: 1, Ub: maxLC_ID}, false)
		for _, i := range ie.Cg_SDT_ConfigLCH_RestrictionToAddModList_r17 {
			tmp_Cg_SDT_ConfigLCH_RestrictionToAddModList_r17.Value = append(tmp_Cg_SDT_ConfigLCH_RestrictionToAddModList_r17.Value, &i)
		}
		if err = tmp_Cg_SDT_ConfigLCH_RestrictionToAddModList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Cg_SDT_ConfigLCH_RestrictionToAddModList_r17", err)
		}
	}
	if len(ie.Cg_SDT_ConfigLCH_RestrictionToReleaseList_r17) > 0 {
		tmp_Cg_SDT_ConfigLCH_RestrictionToReleaseList_r17 := utils.NewSequence[*LogicalChannelIdentity]([]*LogicalChannelIdentity{}, aper.Constraint{Lb: 1, Ub: maxLC_ID}, false)
		for _, i := range ie.Cg_SDT_ConfigLCH_RestrictionToReleaseList_r17 {
			tmp_Cg_SDT_ConfigLCH_RestrictionToReleaseList_r17.Value = append(tmp_Cg_SDT_ConfigLCH_RestrictionToReleaseList_r17.Value, &i)
		}
		if err = tmp_Cg_SDT_ConfigLCH_RestrictionToReleaseList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Cg_SDT_ConfigLCH_RestrictionToReleaseList_r17", err)
		}
	}
	if ie.Cg_SDT_ConfigInitialBWP_NUL_r17 != nil {
		tmp_Cg_SDT_ConfigInitialBWP_NUL_r17 := utils.SetupRelease[*BWP_UplinkDedicatedSDT_r17]{
			Setup: ie.Cg_SDT_ConfigInitialBWP_NUL_r17,
		}
		if err = tmp_Cg_SDT_ConfigInitialBWP_NUL_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Cg_SDT_ConfigInitialBWP_NUL_r17", err)
		}
	}
	if ie.Cg_SDT_ConfigInitialBWP_SUL_r17 != nil {
		tmp_Cg_SDT_ConfigInitialBWP_SUL_r17 := utils.SetupRelease[*BWP_UplinkDedicatedSDT_r17]{
			Setup: ie.Cg_SDT_ConfigInitialBWP_SUL_r17,
		}
		if err = tmp_Cg_SDT_ConfigInitialBWP_SUL_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Cg_SDT_ConfigInitialBWP_SUL_r17", err)
		}
	}
	if ie.Cg_SDT_ConfigInitialBWP_DL_r17 != nil {
		if err = ie.Cg_SDT_ConfigInitialBWP_DL_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Cg_SDT_ConfigInitialBWP_DL_r17", err)
		}
	}
	if ie.Cg_SDT_TimeAlignmentTimer_r17 != nil {
		if err = ie.Cg_SDT_TimeAlignmentTimer_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Cg_SDT_TimeAlignmentTimer_r17", err)
		}
	}
	if ie.Cg_SDT_RSRP_ThresholdSSB_r17 != nil {
		if err = ie.Cg_SDT_RSRP_ThresholdSSB_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Cg_SDT_RSRP_ThresholdSSB_r17", err)
		}
	}
	if ie.Cg_SDT_TA_ValidationConfig_r17 != nil {
		tmp_Cg_SDT_TA_ValidationConfig_r17 := utils.SetupRelease[*CG_SDT_TA_ValidationConfig_r17]{
			Setup: ie.Cg_SDT_TA_ValidationConfig_r17,
		}
		if err = tmp_Cg_SDT_TA_ValidationConfig_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Cg_SDT_TA_ValidationConfig_r17", err)
		}
	}
	if ie.Cg_SDT_CS_RNTI_r17 != nil {
		if err = ie.Cg_SDT_CS_RNTI_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Cg_SDT_CS_RNTI_r17", err)
		}
	}
	return nil
}

func (ie *SDT_MAC_PHY_CG_Config_r17) Decode(r *aper.AperReader) error {
	var err error
	var Cg_SDT_ConfigLCH_RestrictionToAddModList_r17Present bool
	if Cg_SDT_ConfigLCH_RestrictionToAddModList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Cg_SDT_ConfigLCH_RestrictionToReleaseList_r17Present bool
	if Cg_SDT_ConfigLCH_RestrictionToReleaseList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Cg_SDT_ConfigInitialBWP_NUL_r17Present bool
	if Cg_SDT_ConfigInitialBWP_NUL_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Cg_SDT_ConfigInitialBWP_SUL_r17Present bool
	if Cg_SDT_ConfigInitialBWP_SUL_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Cg_SDT_ConfigInitialBWP_DL_r17Present bool
	if Cg_SDT_ConfigInitialBWP_DL_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Cg_SDT_TimeAlignmentTimer_r17Present bool
	if Cg_SDT_TimeAlignmentTimer_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Cg_SDT_RSRP_ThresholdSSB_r17Present bool
	if Cg_SDT_RSRP_ThresholdSSB_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Cg_SDT_TA_ValidationConfig_r17Present bool
	if Cg_SDT_TA_ValidationConfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Cg_SDT_CS_RNTI_r17Present bool
	if Cg_SDT_CS_RNTI_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Cg_SDT_ConfigLCH_RestrictionToAddModList_r17Present {
		tmp_Cg_SDT_ConfigLCH_RestrictionToAddModList_r17 := utils.NewSequence[*CG_SDT_ConfigLCH_Restriction_r17]([]*CG_SDT_ConfigLCH_Restriction_r17{}, aper.Constraint{Lb: 1, Ub: maxLC_ID}, false)
		fn_Cg_SDT_ConfigLCH_RestrictionToAddModList_r17 := func() *CG_SDT_ConfigLCH_Restriction_r17 {
			return new(CG_SDT_ConfigLCH_Restriction_r17)
		}
		if err = tmp_Cg_SDT_ConfigLCH_RestrictionToAddModList_r17.Decode(r, fn_Cg_SDT_ConfigLCH_RestrictionToAddModList_r17); err != nil {
			return utils.WrapError("Decode Cg_SDT_ConfigLCH_RestrictionToAddModList_r17", err)
		}
		ie.Cg_SDT_ConfigLCH_RestrictionToAddModList_r17 = []CG_SDT_ConfigLCH_Restriction_r17{}
		for _, i := range tmp_Cg_SDT_ConfigLCH_RestrictionToAddModList_r17.Value {
			ie.Cg_SDT_ConfigLCH_RestrictionToAddModList_r17 = append(ie.Cg_SDT_ConfigLCH_RestrictionToAddModList_r17, *i)
		}
	}
	if Cg_SDT_ConfigLCH_RestrictionToReleaseList_r17Present {
		tmp_Cg_SDT_ConfigLCH_RestrictionToReleaseList_r17 := utils.NewSequence[*LogicalChannelIdentity]([]*LogicalChannelIdentity{}, aper.Constraint{Lb: 1, Ub: maxLC_ID}, false)
		fn_Cg_SDT_ConfigLCH_RestrictionToReleaseList_r17 := func() *LogicalChannelIdentity {
			return new(LogicalChannelIdentity)
		}
		if err = tmp_Cg_SDT_ConfigLCH_RestrictionToReleaseList_r17.Decode(r, fn_Cg_SDT_ConfigLCH_RestrictionToReleaseList_r17); err != nil {
			return utils.WrapError("Decode Cg_SDT_ConfigLCH_RestrictionToReleaseList_r17", err)
		}
		ie.Cg_SDT_ConfigLCH_RestrictionToReleaseList_r17 = []LogicalChannelIdentity{}
		for _, i := range tmp_Cg_SDT_ConfigLCH_RestrictionToReleaseList_r17.Value {
			ie.Cg_SDT_ConfigLCH_RestrictionToReleaseList_r17 = append(ie.Cg_SDT_ConfigLCH_RestrictionToReleaseList_r17, *i)
		}
	}
	if Cg_SDT_ConfigInitialBWP_NUL_r17Present {
		tmp_Cg_SDT_ConfigInitialBWP_NUL_r17 := utils.SetupRelease[*BWP_UplinkDedicatedSDT_r17]{}
		if err = tmp_Cg_SDT_ConfigInitialBWP_NUL_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Cg_SDT_ConfigInitialBWP_NUL_r17", err)
		}
		ie.Cg_SDT_ConfigInitialBWP_NUL_r17 = tmp_Cg_SDT_ConfigInitialBWP_NUL_r17.Setup
	}
	if Cg_SDT_ConfigInitialBWP_SUL_r17Present {
		tmp_Cg_SDT_ConfigInitialBWP_SUL_r17 := utils.SetupRelease[*BWP_UplinkDedicatedSDT_r17]{}
		if err = tmp_Cg_SDT_ConfigInitialBWP_SUL_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Cg_SDT_ConfigInitialBWP_SUL_r17", err)
		}
		ie.Cg_SDT_ConfigInitialBWP_SUL_r17 = tmp_Cg_SDT_ConfigInitialBWP_SUL_r17.Setup
	}
	if Cg_SDT_ConfigInitialBWP_DL_r17Present {
		ie.Cg_SDT_ConfigInitialBWP_DL_r17 = new(BWP_DownlinkDedicatedSDT_r17)
		if err = ie.Cg_SDT_ConfigInitialBWP_DL_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Cg_SDT_ConfigInitialBWP_DL_r17", err)
		}
	}
	if Cg_SDT_TimeAlignmentTimer_r17Present {
		ie.Cg_SDT_TimeAlignmentTimer_r17 = new(TimeAlignmentTimer)
		if err = ie.Cg_SDT_TimeAlignmentTimer_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Cg_SDT_TimeAlignmentTimer_r17", err)
		}
	}
	if Cg_SDT_RSRP_ThresholdSSB_r17Present {
		ie.Cg_SDT_RSRP_ThresholdSSB_r17 = new(RSRP_Range)
		if err = ie.Cg_SDT_RSRP_ThresholdSSB_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Cg_SDT_RSRP_ThresholdSSB_r17", err)
		}
	}
	if Cg_SDT_TA_ValidationConfig_r17Present {
		tmp_Cg_SDT_TA_ValidationConfig_r17 := utils.SetupRelease[*CG_SDT_TA_ValidationConfig_r17]{}
		if err = tmp_Cg_SDT_TA_ValidationConfig_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Cg_SDT_TA_ValidationConfig_r17", err)
		}
		ie.Cg_SDT_TA_ValidationConfig_r17 = tmp_Cg_SDT_TA_ValidationConfig_r17.Setup
	}
	if Cg_SDT_CS_RNTI_r17Present {
		ie.Cg_SDT_CS_RNTI_r17 = new(RNTI_Value)
		if err = ie.Cg_SDT_CS_RNTI_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Cg_SDT_CS_RNTI_r17", err)
		}
	}
	return nil
}
