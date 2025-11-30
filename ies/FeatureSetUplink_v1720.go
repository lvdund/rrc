package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetUplink_v1720 struct {
	Pucch_Repetition_F0_1_2_3_4_RRC_Config_r17         *FeatureSetUplink_v1720_pucch_Repetition_F0_1_2_3_4_RRC_Config_r17         `optional`
	Pucch_Repetition_F0_1_2_3_4_DynamicIndication_r17  *FeatureSetUplink_v1720_pucch_Repetition_F0_1_2_3_4_DynamicIndication_r17  `optional`
	InterSubslotFreqHopping_PUCCH_r17                  *FeatureSetUplink_v1720_interSubslotFreqHopping_PUCCH_r17                  `optional`
	SemiStaticHARQ_ACK_CodebookSub_SlotPUCCH_r17       *FeatureSetUplink_v1720_semiStaticHARQ_ACK_CodebookSub_SlotPUCCH_r17       `optional`
	Phy_PrioritizationLowPriorityDG_HighPriorityCG_r17 *int64                                                                     `lb:1,ub:16,optional`
	Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17 *FeatureSetUplink_v1720_phy_PrioritizationHighPriorityDG_LowPriorityCG_r17 `lb:1,ub:16,optional`
	ExtendedDC_LocationReport_r17                      *FeatureSetUplink_v1720_extendedDC_LocationReport_r17                      `optional`
}

func (ie *FeatureSetUplink_v1720) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Pucch_Repetition_F0_1_2_3_4_RRC_Config_r17 != nil, ie.Pucch_Repetition_F0_1_2_3_4_DynamicIndication_r17 != nil, ie.InterSubslotFreqHopping_PUCCH_r17 != nil, ie.SemiStaticHARQ_ACK_CodebookSub_SlotPUCCH_r17 != nil, ie.Phy_PrioritizationLowPriorityDG_HighPriorityCG_r17 != nil, ie.Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17 != nil, ie.ExtendedDC_LocationReport_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Pucch_Repetition_F0_1_2_3_4_RRC_Config_r17 != nil {
		if err = ie.Pucch_Repetition_F0_1_2_3_4_RRC_Config_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Pucch_Repetition_F0_1_2_3_4_RRC_Config_r17", err)
		}
	}
	if ie.Pucch_Repetition_F0_1_2_3_4_DynamicIndication_r17 != nil {
		if err = ie.Pucch_Repetition_F0_1_2_3_4_DynamicIndication_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Pucch_Repetition_F0_1_2_3_4_DynamicIndication_r17", err)
		}
	}
	if ie.InterSubslotFreqHopping_PUCCH_r17 != nil {
		if err = ie.InterSubslotFreqHopping_PUCCH_r17.Encode(w); err != nil {
			return utils.WrapError("Encode InterSubslotFreqHopping_PUCCH_r17", err)
		}
	}
	if ie.SemiStaticHARQ_ACK_CodebookSub_SlotPUCCH_r17 != nil {
		if err = ie.SemiStaticHARQ_ACK_CodebookSub_SlotPUCCH_r17.Encode(w); err != nil {
			return utils.WrapError("Encode SemiStaticHARQ_ACK_CodebookSub_SlotPUCCH_r17", err)
		}
	}
	if ie.Phy_PrioritizationLowPriorityDG_HighPriorityCG_r17 != nil {
		if err = w.WriteInteger(*ie.Phy_PrioritizationLowPriorityDG_HighPriorityCG_r17, &aper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
			return utils.WrapError("Encode Phy_PrioritizationLowPriorityDG_HighPriorityCG_r17", err)
		}
	}
	if ie.Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17 != nil {
		if err = ie.Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17", err)
		}
	}
	if ie.ExtendedDC_LocationReport_r17 != nil {
		if err = ie.ExtendedDC_LocationReport_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ExtendedDC_LocationReport_r17", err)
		}
	}
	return nil
}

func (ie *FeatureSetUplink_v1720) Decode(r *aper.AperReader) error {
	var err error
	var Pucch_Repetition_F0_1_2_3_4_RRC_Config_r17Present bool
	if Pucch_Repetition_F0_1_2_3_4_RRC_Config_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pucch_Repetition_F0_1_2_3_4_DynamicIndication_r17Present bool
	if Pucch_Repetition_F0_1_2_3_4_DynamicIndication_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var InterSubslotFreqHopping_PUCCH_r17Present bool
	if InterSubslotFreqHopping_PUCCH_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SemiStaticHARQ_ACK_CodebookSub_SlotPUCCH_r17Present bool
	if SemiStaticHARQ_ACK_CodebookSub_SlotPUCCH_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Phy_PrioritizationLowPriorityDG_HighPriorityCG_r17Present bool
	if Phy_PrioritizationLowPriorityDG_HighPriorityCG_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17Present bool
	if Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ExtendedDC_LocationReport_r17Present bool
	if ExtendedDC_LocationReport_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Pucch_Repetition_F0_1_2_3_4_RRC_Config_r17Present {
		ie.Pucch_Repetition_F0_1_2_3_4_RRC_Config_r17 = new(FeatureSetUplink_v1720_pucch_Repetition_F0_1_2_3_4_RRC_Config_r17)
		if err = ie.Pucch_Repetition_F0_1_2_3_4_RRC_Config_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Pucch_Repetition_F0_1_2_3_4_RRC_Config_r17", err)
		}
	}
	if Pucch_Repetition_F0_1_2_3_4_DynamicIndication_r17Present {
		ie.Pucch_Repetition_F0_1_2_3_4_DynamicIndication_r17 = new(FeatureSetUplink_v1720_pucch_Repetition_F0_1_2_3_4_DynamicIndication_r17)
		if err = ie.Pucch_Repetition_F0_1_2_3_4_DynamicIndication_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Pucch_Repetition_F0_1_2_3_4_DynamicIndication_r17", err)
		}
	}
	if InterSubslotFreqHopping_PUCCH_r17Present {
		ie.InterSubslotFreqHopping_PUCCH_r17 = new(FeatureSetUplink_v1720_interSubslotFreqHopping_PUCCH_r17)
		if err = ie.InterSubslotFreqHopping_PUCCH_r17.Decode(r); err != nil {
			return utils.WrapError("Decode InterSubslotFreqHopping_PUCCH_r17", err)
		}
	}
	if SemiStaticHARQ_ACK_CodebookSub_SlotPUCCH_r17Present {
		ie.SemiStaticHARQ_ACK_CodebookSub_SlotPUCCH_r17 = new(FeatureSetUplink_v1720_semiStaticHARQ_ACK_CodebookSub_SlotPUCCH_r17)
		if err = ie.SemiStaticHARQ_ACK_CodebookSub_SlotPUCCH_r17.Decode(r); err != nil {
			return utils.WrapError("Decode SemiStaticHARQ_ACK_CodebookSub_SlotPUCCH_r17", err)
		}
	}
	if Phy_PrioritizationLowPriorityDG_HighPriorityCG_r17Present {
		var tmp_int_Phy_PrioritizationLowPriorityDG_HighPriorityCG_r17 int64
		if tmp_int_Phy_PrioritizationLowPriorityDG_HighPriorityCG_r17, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode Phy_PrioritizationLowPriorityDG_HighPriorityCG_r17", err)
		}
		ie.Phy_PrioritizationLowPriorityDG_HighPriorityCG_r17 = &tmp_int_Phy_PrioritizationLowPriorityDG_HighPriorityCG_r17
	}
	if Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17Present {
		ie.Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17 = new(FeatureSetUplink_v1720_phy_PrioritizationHighPriorityDG_LowPriorityCG_r17)
		if err = ie.Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17", err)
		}
	}
	if ExtendedDC_LocationReport_r17Present {
		ie.ExtendedDC_LocationReport_r17 = new(FeatureSetUplink_v1720_extendedDC_LocationReport_r17)
		if err = ie.ExtendedDC_LocationReport_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ExtendedDC_LocationReport_r17", err)
		}
	}
	return nil
}
