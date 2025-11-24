package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetDownlink_v1700 struct {
	ScalingFactor_1024QAM_FR1_r17    *FeatureSetDownlink_v1700_scalingFactor_1024QAM_FR1_r17    `optional`
	TimeDurationForQCL_v1710         *FeatureSetDownlink_v1700_timeDurationForQCL_v1710         `optional`
	Sfn_SchemeA_r17                  *FeatureSetDownlink_v1700_sfn_SchemeA_r17                  `optional`
	Sfn_SchemeA_PDCCH_only_r17       *FeatureSetDownlink_v1700_sfn_SchemeA_PDCCH_only_r17       `optional`
	Sfn_SchemeA_DynamicSwitching_r17 *FeatureSetDownlink_v1700_sfn_SchemeA_DynamicSwitching_r17 `optional`
	Sfn_SchemeA_PDSCH_only_r17       *FeatureSetDownlink_v1700_sfn_SchemeA_PDSCH_only_r17       `optional`
	Sfn_SchemeB_r17                  *FeatureSetDownlink_v1700_sfn_SchemeB_r17                  `optional`
	Sfn_SchemeB_DynamicSwitching_r17 *FeatureSetDownlink_v1700_sfn_SchemeB_DynamicSwitching_r17 `optional`
	Sfn_SchemeB_PDSCH_only_r17       *FeatureSetDownlink_v1700_sfn_SchemeB_PDSCH_only_r17       `optional`
	MTRP_PDCCH_Case2_1SpanGap_r17    *FeatureSetDownlink_v1700_mTRP_PDCCH_Case2_1SpanGap_r17    `optional`
	MTRP_PDCCH_legacyMonitoring_r17  *FeatureSetDownlink_v1700_mTRP_PDCCH_legacyMonitoring_r17  `optional`
	MTRP_PDCCH_multiDCI_multiTRP_r17 *FeatureSetDownlink_v1700_mTRP_PDCCH_multiDCI_multiTRP_r17 `optional`
	DynamicMulticastPCell_r17        *FeatureSetDownlink_v1700_dynamicMulticastPCell_r17        `optional`
	MTRP_PDCCH_Repetition_r17        *FeatureSetDownlink_v1700_mTRP_PDCCH_Repetition_r17        `lb:2,ub:3,optional`
}

func (ie *FeatureSetDownlink_v1700) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ScalingFactor_1024QAM_FR1_r17 != nil, ie.TimeDurationForQCL_v1710 != nil, ie.Sfn_SchemeA_r17 != nil, ie.Sfn_SchemeA_PDCCH_only_r17 != nil, ie.Sfn_SchemeA_DynamicSwitching_r17 != nil, ie.Sfn_SchemeA_PDSCH_only_r17 != nil, ie.Sfn_SchemeB_r17 != nil, ie.Sfn_SchemeB_DynamicSwitching_r17 != nil, ie.Sfn_SchemeB_PDSCH_only_r17 != nil, ie.MTRP_PDCCH_Case2_1SpanGap_r17 != nil, ie.MTRP_PDCCH_legacyMonitoring_r17 != nil, ie.MTRP_PDCCH_multiDCI_multiTRP_r17 != nil, ie.DynamicMulticastPCell_r17 != nil, ie.MTRP_PDCCH_Repetition_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ScalingFactor_1024QAM_FR1_r17 != nil {
		if err = ie.ScalingFactor_1024QAM_FR1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ScalingFactor_1024QAM_FR1_r17", err)
		}
	}
	if ie.TimeDurationForQCL_v1710 != nil {
		if err = ie.TimeDurationForQCL_v1710.Encode(w); err != nil {
			return utils.WrapError("Encode TimeDurationForQCL_v1710", err)
		}
	}
	if ie.Sfn_SchemeA_r17 != nil {
		if err = ie.Sfn_SchemeA_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sfn_SchemeA_r17", err)
		}
	}
	if ie.Sfn_SchemeA_PDCCH_only_r17 != nil {
		if err = ie.Sfn_SchemeA_PDCCH_only_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sfn_SchemeA_PDCCH_only_r17", err)
		}
	}
	if ie.Sfn_SchemeA_DynamicSwitching_r17 != nil {
		if err = ie.Sfn_SchemeA_DynamicSwitching_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sfn_SchemeA_DynamicSwitching_r17", err)
		}
	}
	if ie.Sfn_SchemeA_PDSCH_only_r17 != nil {
		if err = ie.Sfn_SchemeA_PDSCH_only_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sfn_SchemeA_PDSCH_only_r17", err)
		}
	}
	if ie.Sfn_SchemeB_r17 != nil {
		if err = ie.Sfn_SchemeB_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sfn_SchemeB_r17", err)
		}
	}
	if ie.Sfn_SchemeB_DynamicSwitching_r17 != nil {
		if err = ie.Sfn_SchemeB_DynamicSwitching_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sfn_SchemeB_DynamicSwitching_r17", err)
		}
	}
	if ie.Sfn_SchemeB_PDSCH_only_r17 != nil {
		if err = ie.Sfn_SchemeB_PDSCH_only_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sfn_SchemeB_PDSCH_only_r17", err)
		}
	}
	if ie.MTRP_PDCCH_Case2_1SpanGap_r17 != nil {
		if err = ie.MTRP_PDCCH_Case2_1SpanGap_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MTRP_PDCCH_Case2_1SpanGap_r17", err)
		}
	}
	if ie.MTRP_PDCCH_legacyMonitoring_r17 != nil {
		if err = ie.MTRP_PDCCH_legacyMonitoring_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MTRP_PDCCH_legacyMonitoring_r17", err)
		}
	}
	if ie.MTRP_PDCCH_multiDCI_multiTRP_r17 != nil {
		if err = ie.MTRP_PDCCH_multiDCI_multiTRP_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MTRP_PDCCH_multiDCI_multiTRP_r17", err)
		}
	}
	if ie.DynamicMulticastPCell_r17 != nil {
		if err = ie.DynamicMulticastPCell_r17.Encode(w); err != nil {
			return utils.WrapError("Encode DynamicMulticastPCell_r17", err)
		}
	}
	if ie.MTRP_PDCCH_Repetition_r17 != nil {
		if err = ie.MTRP_PDCCH_Repetition_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MTRP_PDCCH_Repetition_r17", err)
		}
	}
	return nil
}

func (ie *FeatureSetDownlink_v1700) Decode(r *uper.UperReader) error {
	var err error
	var ScalingFactor_1024QAM_FR1_r17Present bool
	if ScalingFactor_1024QAM_FR1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var TimeDurationForQCL_v1710Present bool
	if TimeDurationForQCL_v1710Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sfn_SchemeA_r17Present bool
	if Sfn_SchemeA_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sfn_SchemeA_PDCCH_only_r17Present bool
	if Sfn_SchemeA_PDCCH_only_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sfn_SchemeA_DynamicSwitching_r17Present bool
	if Sfn_SchemeA_DynamicSwitching_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sfn_SchemeA_PDSCH_only_r17Present bool
	if Sfn_SchemeA_PDSCH_only_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sfn_SchemeB_r17Present bool
	if Sfn_SchemeB_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sfn_SchemeB_DynamicSwitching_r17Present bool
	if Sfn_SchemeB_DynamicSwitching_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sfn_SchemeB_PDSCH_only_r17Present bool
	if Sfn_SchemeB_PDSCH_only_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MTRP_PDCCH_Case2_1SpanGap_r17Present bool
	if MTRP_PDCCH_Case2_1SpanGap_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MTRP_PDCCH_legacyMonitoring_r17Present bool
	if MTRP_PDCCH_legacyMonitoring_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MTRP_PDCCH_multiDCI_multiTRP_r17Present bool
	if MTRP_PDCCH_multiDCI_multiTRP_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var DynamicMulticastPCell_r17Present bool
	if DynamicMulticastPCell_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MTRP_PDCCH_Repetition_r17Present bool
	if MTRP_PDCCH_Repetition_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ScalingFactor_1024QAM_FR1_r17Present {
		ie.ScalingFactor_1024QAM_FR1_r17 = new(FeatureSetDownlink_v1700_scalingFactor_1024QAM_FR1_r17)
		if err = ie.ScalingFactor_1024QAM_FR1_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ScalingFactor_1024QAM_FR1_r17", err)
		}
	}
	if TimeDurationForQCL_v1710Present {
		ie.TimeDurationForQCL_v1710 = new(FeatureSetDownlink_v1700_timeDurationForQCL_v1710)
		if err = ie.TimeDurationForQCL_v1710.Decode(r); err != nil {
			return utils.WrapError("Decode TimeDurationForQCL_v1710", err)
		}
	}
	if Sfn_SchemeA_r17Present {
		ie.Sfn_SchemeA_r17 = new(FeatureSetDownlink_v1700_sfn_SchemeA_r17)
		if err = ie.Sfn_SchemeA_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sfn_SchemeA_r17", err)
		}
	}
	if Sfn_SchemeA_PDCCH_only_r17Present {
		ie.Sfn_SchemeA_PDCCH_only_r17 = new(FeatureSetDownlink_v1700_sfn_SchemeA_PDCCH_only_r17)
		if err = ie.Sfn_SchemeA_PDCCH_only_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sfn_SchemeA_PDCCH_only_r17", err)
		}
	}
	if Sfn_SchemeA_DynamicSwitching_r17Present {
		ie.Sfn_SchemeA_DynamicSwitching_r17 = new(FeatureSetDownlink_v1700_sfn_SchemeA_DynamicSwitching_r17)
		if err = ie.Sfn_SchemeA_DynamicSwitching_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sfn_SchemeA_DynamicSwitching_r17", err)
		}
	}
	if Sfn_SchemeA_PDSCH_only_r17Present {
		ie.Sfn_SchemeA_PDSCH_only_r17 = new(FeatureSetDownlink_v1700_sfn_SchemeA_PDSCH_only_r17)
		if err = ie.Sfn_SchemeA_PDSCH_only_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sfn_SchemeA_PDSCH_only_r17", err)
		}
	}
	if Sfn_SchemeB_r17Present {
		ie.Sfn_SchemeB_r17 = new(FeatureSetDownlink_v1700_sfn_SchemeB_r17)
		if err = ie.Sfn_SchemeB_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sfn_SchemeB_r17", err)
		}
	}
	if Sfn_SchemeB_DynamicSwitching_r17Present {
		ie.Sfn_SchemeB_DynamicSwitching_r17 = new(FeatureSetDownlink_v1700_sfn_SchemeB_DynamicSwitching_r17)
		if err = ie.Sfn_SchemeB_DynamicSwitching_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sfn_SchemeB_DynamicSwitching_r17", err)
		}
	}
	if Sfn_SchemeB_PDSCH_only_r17Present {
		ie.Sfn_SchemeB_PDSCH_only_r17 = new(FeatureSetDownlink_v1700_sfn_SchemeB_PDSCH_only_r17)
		if err = ie.Sfn_SchemeB_PDSCH_only_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sfn_SchemeB_PDSCH_only_r17", err)
		}
	}
	if MTRP_PDCCH_Case2_1SpanGap_r17Present {
		ie.MTRP_PDCCH_Case2_1SpanGap_r17 = new(FeatureSetDownlink_v1700_mTRP_PDCCH_Case2_1SpanGap_r17)
		if err = ie.MTRP_PDCCH_Case2_1SpanGap_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MTRP_PDCCH_Case2_1SpanGap_r17", err)
		}
	}
	if MTRP_PDCCH_legacyMonitoring_r17Present {
		ie.MTRP_PDCCH_legacyMonitoring_r17 = new(FeatureSetDownlink_v1700_mTRP_PDCCH_legacyMonitoring_r17)
		if err = ie.MTRP_PDCCH_legacyMonitoring_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MTRP_PDCCH_legacyMonitoring_r17", err)
		}
	}
	if MTRP_PDCCH_multiDCI_multiTRP_r17Present {
		ie.MTRP_PDCCH_multiDCI_multiTRP_r17 = new(FeatureSetDownlink_v1700_mTRP_PDCCH_multiDCI_multiTRP_r17)
		if err = ie.MTRP_PDCCH_multiDCI_multiTRP_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MTRP_PDCCH_multiDCI_multiTRP_r17", err)
		}
	}
	if DynamicMulticastPCell_r17Present {
		ie.DynamicMulticastPCell_r17 = new(FeatureSetDownlink_v1700_dynamicMulticastPCell_r17)
		if err = ie.DynamicMulticastPCell_r17.Decode(r); err != nil {
			return utils.WrapError("Decode DynamicMulticastPCell_r17", err)
		}
	}
	if MTRP_PDCCH_Repetition_r17Present {
		ie.MTRP_PDCCH_Repetition_r17 = new(FeatureSetDownlink_v1700_mTRP_PDCCH_Repetition_r17)
		if err = ie.MTRP_PDCCH_Repetition_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MTRP_PDCCH_Repetition_r17", err)
		}
	}
	return nil
}
