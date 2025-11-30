package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MBS_RNTI_SpecificConfig_r17 struct {
	Mbs_RNTI_SpecificConfigId_r17     MBS_RNTI_SpecificConfigId_r17                                  `madatory`
	GroupCommon_RNTI_r17              MBS_RNTI_SpecificConfig_r17_groupCommon_RNTI_r17               `madatory`
	Drx_ConfigPTM_r17                 *DRX_ConfigPTM_r17                                             `optional,setuprelease`
	Harq_FeedbackEnablerMulticast_r17 *MBS_RNTI_SpecificConfig_r17_harq_FeedbackEnablerMulticast_r17 `optional`
	Harq_FeedbackOptionMulticast_r17  *MBS_RNTI_SpecificConfig_r17_harq_FeedbackOptionMulticast_r17  `optional`
	Pdsch_AggregationFactor_r17       *MBS_RNTI_SpecificConfig_r17_pdsch_AggregationFactor_r17       `optional`
}

func (ie *MBS_RNTI_SpecificConfig_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Drx_ConfigPTM_r17 != nil, ie.Harq_FeedbackEnablerMulticast_r17 != nil, ie.Harq_FeedbackOptionMulticast_r17 != nil, ie.Pdsch_AggregationFactor_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Mbs_RNTI_SpecificConfigId_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Mbs_RNTI_SpecificConfigId_r17", err)
	}
	if err = ie.GroupCommon_RNTI_r17.Encode(w); err != nil {
		return utils.WrapError("Encode GroupCommon_RNTI_r17", err)
	}
	if ie.Drx_ConfigPTM_r17 != nil {
		tmp_Drx_ConfigPTM_r17 := utils.SetupRelease[*DRX_ConfigPTM_r17]{
			Setup: ie.Drx_ConfigPTM_r17,
		}
		if err = tmp_Drx_ConfigPTM_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Drx_ConfigPTM_r17", err)
		}
	}
	if ie.Harq_FeedbackEnablerMulticast_r17 != nil {
		if err = ie.Harq_FeedbackEnablerMulticast_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Harq_FeedbackEnablerMulticast_r17", err)
		}
	}
	if ie.Harq_FeedbackOptionMulticast_r17 != nil {
		if err = ie.Harq_FeedbackOptionMulticast_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Harq_FeedbackOptionMulticast_r17", err)
		}
	}
	if ie.Pdsch_AggregationFactor_r17 != nil {
		if err = ie.Pdsch_AggregationFactor_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Pdsch_AggregationFactor_r17", err)
		}
	}
	return nil
}

func (ie *MBS_RNTI_SpecificConfig_r17) Decode(r *aper.AperReader) error {
	var err error
	var Drx_ConfigPTM_r17Present bool
	if Drx_ConfigPTM_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Harq_FeedbackEnablerMulticast_r17Present bool
	if Harq_FeedbackEnablerMulticast_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Harq_FeedbackOptionMulticast_r17Present bool
	if Harq_FeedbackOptionMulticast_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdsch_AggregationFactor_r17Present bool
	if Pdsch_AggregationFactor_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Mbs_RNTI_SpecificConfigId_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Mbs_RNTI_SpecificConfigId_r17", err)
	}
	if err = ie.GroupCommon_RNTI_r17.Decode(r); err != nil {
		return utils.WrapError("Decode GroupCommon_RNTI_r17", err)
	}
	if Drx_ConfigPTM_r17Present {
		tmp_Drx_ConfigPTM_r17 := utils.SetupRelease[*DRX_ConfigPTM_r17]{}
		if err = tmp_Drx_ConfigPTM_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Drx_ConfigPTM_r17", err)
		}
		ie.Drx_ConfigPTM_r17 = tmp_Drx_ConfigPTM_r17.Setup
	}
	if Harq_FeedbackEnablerMulticast_r17Present {
		ie.Harq_FeedbackEnablerMulticast_r17 = new(MBS_RNTI_SpecificConfig_r17_harq_FeedbackEnablerMulticast_r17)
		if err = ie.Harq_FeedbackEnablerMulticast_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Harq_FeedbackEnablerMulticast_r17", err)
		}
	}
	if Harq_FeedbackOptionMulticast_r17Present {
		ie.Harq_FeedbackOptionMulticast_r17 = new(MBS_RNTI_SpecificConfig_r17_harq_FeedbackOptionMulticast_r17)
		if err = ie.Harq_FeedbackOptionMulticast_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Harq_FeedbackOptionMulticast_r17", err)
		}
	}
	if Pdsch_AggregationFactor_r17Present {
		ie.Pdsch_AggregationFactor_r17 = new(MBS_RNTI_SpecificConfig_r17_pdsch_AggregationFactor_r17)
		if err = ie.Pdsch_AggregationFactor_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Pdsch_AggregationFactor_r17", err)
		}
	}
	return nil
}
