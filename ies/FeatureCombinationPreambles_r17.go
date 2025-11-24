package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureCombinationPreambles_r17 struct {
	FeatureCombination_r17                       FeatureCombination_r17                                `madatory`
	StartPreambleForThisPartition_r17            int64                                                 `lb:0,ub:63,madatory`
	NumberOfPreamblesPerSSB_ForThisPartition_r17 int64                                                 `lb:1,ub:64,madatory`
	Ssb_SharedRO_MaskIndex_r17                   *int64                                                `lb:1,ub:15,optional`
	GroupBconfigured_r17                         *FeatureCombinationPreambles_r17_groupBconfigured_r17 `lb:1,ub:64,optional`
	SeparateMsgA_PUSCH_Config_r17                *MsgA_PUSCH_Config_r16                                `optional`
	MsgA_RSRP_Threshold_r17                      *RSRP_Range                                           `optional`
	Rsrp_ThresholdSSB_r17                        *RSRP_Range                                           `optional`
	DeltaPreamble_r17                            *int64                                                `lb:-1,ub:6,optional`
}

func (ie *FeatureCombinationPreambles_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Ssb_SharedRO_MaskIndex_r17 != nil, ie.GroupBconfigured_r17 != nil, ie.SeparateMsgA_PUSCH_Config_r17 != nil, ie.MsgA_RSRP_Threshold_r17 != nil, ie.Rsrp_ThresholdSSB_r17 != nil, ie.DeltaPreamble_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.FeatureCombination_r17.Encode(w); err != nil {
		return utils.WrapError("Encode FeatureCombination_r17", err)
	}
	if err = w.WriteInteger(ie.StartPreambleForThisPartition_r17, &uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("WriteInteger StartPreambleForThisPartition_r17", err)
	}
	if err = w.WriteInteger(ie.NumberOfPreamblesPerSSB_ForThisPartition_r17, &uper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return utils.WrapError("WriteInteger NumberOfPreamblesPerSSB_ForThisPartition_r17", err)
	}
	if ie.Ssb_SharedRO_MaskIndex_r17 != nil {
		if err = w.WriteInteger(*ie.Ssb_SharedRO_MaskIndex_r17, &uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode Ssb_SharedRO_MaskIndex_r17", err)
		}
	}
	if ie.GroupBconfigured_r17 != nil {
		if err = ie.GroupBconfigured_r17.Encode(w); err != nil {
			return utils.WrapError("Encode GroupBconfigured_r17", err)
		}
	}
	if ie.SeparateMsgA_PUSCH_Config_r17 != nil {
		if err = ie.SeparateMsgA_PUSCH_Config_r17.Encode(w); err != nil {
			return utils.WrapError("Encode SeparateMsgA_PUSCH_Config_r17", err)
		}
	}
	if ie.MsgA_RSRP_Threshold_r17 != nil {
		if err = ie.MsgA_RSRP_Threshold_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MsgA_RSRP_Threshold_r17", err)
		}
	}
	if ie.Rsrp_ThresholdSSB_r17 != nil {
		if err = ie.Rsrp_ThresholdSSB_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Rsrp_ThresholdSSB_r17", err)
		}
	}
	if ie.DeltaPreamble_r17 != nil {
		if err = w.WriteInteger(*ie.DeltaPreamble_r17, &uper.Constraint{Lb: -1, Ub: 6}, false); err != nil {
			return utils.WrapError("Encode DeltaPreamble_r17", err)
		}
	}
	return nil
}

func (ie *FeatureCombinationPreambles_r17) Decode(r *uper.UperReader) error {
	var err error
	var Ssb_SharedRO_MaskIndex_r17Present bool
	if Ssb_SharedRO_MaskIndex_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var GroupBconfigured_r17Present bool
	if GroupBconfigured_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SeparateMsgA_PUSCH_Config_r17Present bool
	if SeparateMsgA_PUSCH_Config_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MsgA_RSRP_Threshold_r17Present bool
	if MsgA_RSRP_Threshold_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Rsrp_ThresholdSSB_r17Present bool
	if Rsrp_ThresholdSSB_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var DeltaPreamble_r17Present bool
	if DeltaPreamble_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.FeatureCombination_r17.Decode(r); err != nil {
		return utils.WrapError("Decode FeatureCombination_r17", err)
	}
	var tmp_int_StartPreambleForThisPartition_r17 int64
	if tmp_int_StartPreambleForThisPartition_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("ReadInteger StartPreambleForThisPartition_r17", err)
	}
	ie.StartPreambleForThisPartition_r17 = tmp_int_StartPreambleForThisPartition_r17
	var tmp_int_NumberOfPreamblesPerSSB_ForThisPartition_r17 int64
	if tmp_int_NumberOfPreamblesPerSSB_ForThisPartition_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return utils.WrapError("ReadInteger NumberOfPreamblesPerSSB_ForThisPartition_r17", err)
	}
	ie.NumberOfPreamblesPerSSB_ForThisPartition_r17 = tmp_int_NumberOfPreamblesPerSSB_ForThisPartition_r17
	if Ssb_SharedRO_MaskIndex_r17Present {
		var tmp_int_Ssb_SharedRO_MaskIndex_r17 int64
		if tmp_int_Ssb_SharedRO_MaskIndex_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode Ssb_SharedRO_MaskIndex_r17", err)
		}
		ie.Ssb_SharedRO_MaskIndex_r17 = &tmp_int_Ssb_SharedRO_MaskIndex_r17
	}
	if GroupBconfigured_r17Present {
		ie.GroupBconfigured_r17 = new(FeatureCombinationPreambles_r17_groupBconfigured_r17)
		if err = ie.GroupBconfigured_r17.Decode(r); err != nil {
			return utils.WrapError("Decode GroupBconfigured_r17", err)
		}
	}
	if SeparateMsgA_PUSCH_Config_r17Present {
		ie.SeparateMsgA_PUSCH_Config_r17 = new(MsgA_PUSCH_Config_r16)
		if err = ie.SeparateMsgA_PUSCH_Config_r17.Decode(r); err != nil {
			return utils.WrapError("Decode SeparateMsgA_PUSCH_Config_r17", err)
		}
	}
	if MsgA_RSRP_Threshold_r17Present {
		ie.MsgA_RSRP_Threshold_r17 = new(RSRP_Range)
		if err = ie.MsgA_RSRP_Threshold_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MsgA_RSRP_Threshold_r17", err)
		}
	}
	if Rsrp_ThresholdSSB_r17Present {
		ie.Rsrp_ThresholdSSB_r17 = new(RSRP_Range)
		if err = ie.Rsrp_ThresholdSSB_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Rsrp_ThresholdSSB_r17", err)
		}
	}
	if DeltaPreamble_r17Present {
		var tmp_int_DeltaPreamble_r17 int64
		if tmp_int_DeltaPreamble_r17, err = r.ReadInteger(&uper.Constraint{Lb: -1, Ub: 6}, false); err != nil {
			return utils.WrapError("Decode DeltaPreamble_r17", err)
		}
		ie.DeltaPreamble_r17 = &tmp_int_DeltaPreamble_r17
	}
	return nil
}
