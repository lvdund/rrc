// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FeatureCombinationPreambles-r17 (line 8321).

var featureCombinationPreamblesR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "featureCombination-r17"},
		{Name: "startPreambleForThisPartition-r17"},
		{Name: "numberOfPreamblesPerSSB-ForThisPartition-r17"},
		{Name: "ssb-SharedRO-MaskIndex-r17", Optional: true},
		{Name: "groupBconfigured-r17", Optional: true},
		{Name: "separateMsgA-PUSCH-Config-r17", Optional: true},
		{Name: "msgA-RSRP-Threshold-r17", Optional: true},
		{Name: "rsrp-ThresholdSSB-r17", Optional: true},
		{Name: "deltaPreamble-r17", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var featureCombinationPreamblesR17StartPreambleForThisPartitionR17Constraints = per.Constrained(0, 63)

var featureCombinationPreamblesR17NumberOfPreamblesPerSSBForThisPartitionR17Constraints = per.Constrained(1, 64)

var featureCombinationPreamblesR17SsbSharedROMaskIndexR17Constraints = per.Constrained(1, 15)

var featureCombinationPreamblesR17DeltaPreambleR17Constraints = per.Constrained(-1, 6)

const (
	FeatureCombinationPreambles_r17_Ext_Msg1_RepetitionNum_r18_N2     = 0
	FeatureCombinationPreambles_r17_Ext_Msg1_RepetitionNum_r18_N4     = 1
	FeatureCombinationPreambles_r17_Ext_Msg1_RepetitionNum_r18_N8     = 2
	FeatureCombinationPreambles_r17_Ext_Msg1_RepetitionNum_r18_Spare1 = 3
)

var featureCombinationPreamblesR17ExtMsg1RepetitionNumR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureCombinationPreambles_r17_Ext_Msg1_RepetitionNum_r18_N2, FeatureCombinationPreambles_r17_Ext_Msg1_RepetitionNum_r18_N4, FeatureCombinationPreambles_r17_Ext_Msg1_RepetitionNum_r18_N8, FeatureCombinationPreambles_r17_Ext_Msg1_RepetitionNum_r18_Spare1},
}

const (
	FeatureCombinationPreambles_r17_Ext_Msg1_RepetitionTimeOffsetROGroup_r18_N4     = 0
	FeatureCombinationPreambles_r17_Ext_Msg1_RepetitionTimeOffsetROGroup_r18_N8     = 1
	FeatureCombinationPreambles_r17_Ext_Msg1_RepetitionTimeOffsetROGroup_r18_N16    = 2
	FeatureCombinationPreambles_r17_Ext_Msg1_RepetitionTimeOffsetROGroup_r18_Spare1 = 3
)

var featureCombinationPreamblesR17ExtMsg1RepetitionTimeOffsetROGroupR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureCombinationPreambles_r17_Ext_Msg1_RepetitionTimeOffsetROGroup_r18_N4, FeatureCombinationPreambles_r17_Ext_Msg1_RepetitionTimeOffsetROGroup_r18_N8, FeatureCombinationPreambles_r17_Ext_Msg1_RepetitionTimeOffsetROGroup_r18_N16, FeatureCombinationPreambles_r17_Ext_Msg1_RepetitionTimeOffsetROGroup_r18_Spare1},
}

const (
	FeatureCombinationPreambles_r17_GroupBconfigured_r17_Ra_SizeGroupA_r17_B56    = 0
	FeatureCombinationPreambles_r17_GroupBconfigured_r17_Ra_SizeGroupA_r17_B144   = 1
	FeatureCombinationPreambles_r17_GroupBconfigured_r17_Ra_SizeGroupA_r17_B208   = 2
	FeatureCombinationPreambles_r17_GroupBconfigured_r17_Ra_SizeGroupA_r17_B256   = 3
	FeatureCombinationPreambles_r17_GroupBconfigured_r17_Ra_SizeGroupA_r17_B282   = 4
	FeatureCombinationPreambles_r17_GroupBconfigured_r17_Ra_SizeGroupA_r17_B480   = 5
	FeatureCombinationPreambles_r17_GroupBconfigured_r17_Ra_SizeGroupA_r17_B640   = 6
	FeatureCombinationPreambles_r17_GroupBconfigured_r17_Ra_SizeGroupA_r17_B800   = 7
	FeatureCombinationPreambles_r17_GroupBconfigured_r17_Ra_SizeGroupA_r17_B1000  = 8
	FeatureCombinationPreambles_r17_GroupBconfigured_r17_Ra_SizeGroupA_r17_B72    = 9
	FeatureCombinationPreambles_r17_GroupBconfigured_r17_Ra_SizeGroupA_r17_Spare6 = 10
	FeatureCombinationPreambles_r17_GroupBconfigured_r17_Ra_SizeGroupA_r17_Spare5 = 11
	FeatureCombinationPreambles_r17_GroupBconfigured_r17_Ra_SizeGroupA_r17_Spare4 = 12
	FeatureCombinationPreambles_r17_GroupBconfigured_r17_Ra_SizeGroupA_r17_Spare3 = 13
	FeatureCombinationPreambles_r17_GroupBconfigured_r17_Ra_SizeGroupA_r17_Spare2 = 14
	FeatureCombinationPreambles_r17_GroupBconfigured_r17_Ra_SizeGroupA_r17_Spare1 = 15
)

var featureCombinationPreamblesR17GroupBconfiguredR17RaSizeGroupAR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureCombinationPreambles_r17_GroupBconfigured_r17_Ra_SizeGroupA_r17_B56, FeatureCombinationPreambles_r17_GroupBconfigured_r17_Ra_SizeGroupA_r17_B144, FeatureCombinationPreambles_r17_GroupBconfigured_r17_Ra_SizeGroupA_r17_B208, FeatureCombinationPreambles_r17_GroupBconfigured_r17_Ra_SizeGroupA_r17_B256, FeatureCombinationPreambles_r17_GroupBconfigured_r17_Ra_SizeGroupA_r17_B282, FeatureCombinationPreambles_r17_GroupBconfigured_r17_Ra_SizeGroupA_r17_B480, FeatureCombinationPreambles_r17_GroupBconfigured_r17_Ra_SizeGroupA_r17_B640, FeatureCombinationPreambles_r17_GroupBconfigured_r17_Ra_SizeGroupA_r17_B800, FeatureCombinationPreambles_r17_GroupBconfigured_r17_Ra_SizeGroupA_r17_B1000, FeatureCombinationPreambles_r17_GroupBconfigured_r17_Ra_SizeGroupA_r17_B72, FeatureCombinationPreambles_r17_GroupBconfigured_r17_Ra_SizeGroupA_r17_Spare6, FeatureCombinationPreambles_r17_GroupBconfigured_r17_Ra_SizeGroupA_r17_Spare5, FeatureCombinationPreambles_r17_GroupBconfigured_r17_Ra_SizeGroupA_r17_Spare4, FeatureCombinationPreambles_r17_GroupBconfigured_r17_Ra_SizeGroupA_r17_Spare3, FeatureCombinationPreambles_r17_GroupBconfigured_r17_Ra_SizeGroupA_r17_Spare2, FeatureCombinationPreambles_r17_GroupBconfigured_r17_Ra_SizeGroupA_r17_Spare1},
}

const (
	FeatureCombinationPreambles_r17_GroupBconfigured_r17_MessagePowerOffsetGroupB_r17_Minusinfinity = 0
	FeatureCombinationPreambles_r17_GroupBconfigured_r17_MessagePowerOffsetGroupB_r17_DB0           = 1
	FeatureCombinationPreambles_r17_GroupBconfigured_r17_MessagePowerOffsetGroupB_r17_DB5           = 2
	FeatureCombinationPreambles_r17_GroupBconfigured_r17_MessagePowerOffsetGroupB_r17_DB8           = 3
	FeatureCombinationPreambles_r17_GroupBconfigured_r17_MessagePowerOffsetGroupB_r17_DB10          = 4
	FeatureCombinationPreambles_r17_GroupBconfigured_r17_MessagePowerOffsetGroupB_r17_DB12          = 5
	FeatureCombinationPreambles_r17_GroupBconfigured_r17_MessagePowerOffsetGroupB_r17_DB15          = 6
	FeatureCombinationPreambles_r17_GroupBconfigured_r17_MessagePowerOffsetGroupB_r17_DB18          = 7
)

var featureCombinationPreamblesR17GroupBconfiguredR17MessagePowerOffsetGroupBR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureCombinationPreambles_r17_GroupBconfigured_r17_MessagePowerOffsetGroupB_r17_Minusinfinity, FeatureCombinationPreambles_r17_GroupBconfigured_r17_MessagePowerOffsetGroupB_r17_DB0, FeatureCombinationPreambles_r17_GroupBconfigured_r17_MessagePowerOffsetGroupB_r17_DB5, FeatureCombinationPreambles_r17_GroupBconfigured_r17_MessagePowerOffsetGroupB_r17_DB8, FeatureCombinationPreambles_r17_GroupBconfigured_r17_MessagePowerOffsetGroupB_r17_DB10, FeatureCombinationPreambles_r17_GroupBconfigured_r17_MessagePowerOffsetGroupB_r17_DB12, FeatureCombinationPreambles_r17_GroupBconfigured_r17_MessagePowerOffsetGroupB_r17_DB15, FeatureCombinationPreambles_r17_GroupBconfigured_r17_MessagePowerOffsetGroupB_r17_DB18},
}

type FeatureCombinationPreambles_r17 struct {
	FeatureCombination_r17                       FeatureCombination_r17
	StartPreambleForThisPartition_r17            int64
	NumberOfPreamblesPerSSB_ForThisPartition_r17 int64
	Ssb_SharedRO_MaskIndex_r17                   *int64
	GroupBconfigured_r17                         *struct {
		Ra_SizeGroupA_r17              int64
		MessagePowerOffsetGroupB_r17   int64
		NumberOfRA_PreamblesGroupA_r17 int64
	}
	SeparateMsgA_PUSCH_Config_r17        *MsgA_PUSCH_Config_r16
	MsgA_RSRP_Threshold_r17              *RSRP_Range
	Rsrp_ThresholdSSB_r17                *RSRP_Range
	DeltaPreamble_r17                    *int64
	Msg1_RepetitionNum_r18               *int64
	Msg1_RepetitionTimeOffsetROGroup_r18 *int64
}

func (ie *FeatureCombinationPreambles_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureCombinationPreamblesR17Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Msg1_RepetitionNum_r18 != nil || ie.Msg1_RepetitionTimeOffsetROGroup_r18 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ssb_SharedRO_MaskIndex_r17 != nil, ie.GroupBconfigured_r17 != nil, ie.SeparateMsgA_PUSCH_Config_r17 != nil, ie.MsgA_RSRP_Threshold_r17 != nil, ie.Rsrp_ThresholdSSB_r17 != nil, ie.DeltaPreamble_r17 != nil}); err != nil {
		return err
	}

	// 3. featureCombination-r17: ref
	{
		if err := ie.FeatureCombination_r17.Encode(e); err != nil {
			return err
		}
	}

	// 4. startPreambleForThisPartition-r17: integer
	{
		if err := e.EncodeInteger(ie.StartPreambleForThisPartition_r17, featureCombinationPreamblesR17StartPreambleForThisPartitionR17Constraints); err != nil {
			return err
		}
	}

	// 5. numberOfPreamblesPerSSB-ForThisPartition-r17: integer
	{
		if err := e.EncodeInteger(ie.NumberOfPreamblesPerSSB_ForThisPartition_r17, featureCombinationPreamblesR17NumberOfPreamblesPerSSBForThisPartitionR17Constraints); err != nil {
			return err
		}
	}

	// 6. ssb-SharedRO-MaskIndex-r17: integer
	{
		if ie.Ssb_SharedRO_MaskIndex_r17 != nil {
			if err := e.EncodeInteger(*ie.Ssb_SharedRO_MaskIndex_r17, featureCombinationPreamblesR17SsbSharedROMaskIndexR17Constraints); err != nil {
				return err
			}
		}
	}

	// 7. groupBconfigured-r17: sequence
	{
		if ie.GroupBconfigured_r17 != nil {
			c := ie.GroupBconfigured_r17
			if err := e.EncodeEnumerated(c.Ra_SizeGroupA_r17, featureCombinationPreamblesR17GroupBconfiguredR17RaSizeGroupAR17Constraints); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.MessagePowerOffsetGroupB_r17, featureCombinationPreamblesR17GroupBconfiguredR17MessagePowerOffsetGroupBR17Constraints); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.NumberOfRA_PreamblesGroupA_r17, per.Constrained(1, 64)); err != nil {
				return err
			}
		}
	}

	// 8. separateMsgA-PUSCH-Config-r17: ref
	{
		if ie.SeparateMsgA_PUSCH_Config_r17 != nil {
			if err := ie.SeparateMsgA_PUSCH_Config_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 9. msgA-RSRP-Threshold-r17: ref
	{
		if ie.MsgA_RSRP_Threshold_r17 != nil {
			if err := ie.MsgA_RSRP_Threshold_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 10. rsrp-ThresholdSSB-r17: ref
	{
		if ie.Rsrp_ThresholdSSB_r17 != nil {
			if err := ie.Rsrp_ThresholdSSB_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 11. deltaPreamble-r17: integer
	{
		if ie.DeltaPreamble_r17 != nil {
			if err := e.EncodeInteger(*ie.DeltaPreamble_r17, featureCombinationPreamblesR17DeltaPreambleR17Constraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "msg1-RepetitionNum-r18", Optional: true},
					{Name: "msg1-RepetitionTimeOffsetROGroup-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Msg1_RepetitionNum_r18 != nil, ie.Msg1_RepetitionTimeOffsetROGroup_r18 != nil}); err != nil {
				return err
			}

			if ie.Msg1_RepetitionNum_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Msg1_RepetitionNum_r18, featureCombinationPreamblesR17ExtMsg1RepetitionNumR18Constraints); err != nil {
					return err
				}
			}

			if ie.Msg1_RepetitionTimeOffsetROGroup_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Msg1_RepetitionTimeOffsetROGroup_r18, featureCombinationPreamblesR17ExtMsg1RepetitionTimeOffsetROGroupR18Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		if err := seq.EncodeExtensionAdditions(extPresence, extValues); err != nil {
			return err
		}
	}

	return nil
}

func (ie *FeatureCombinationPreambles_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureCombinationPreamblesR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. featureCombination-r17: ref
	{
		if err := ie.FeatureCombination_r17.Decode(d); err != nil {
			return err
		}
	}

	// 4. startPreambleForThisPartition-r17: integer
	{
		v1, err := d.DecodeInteger(featureCombinationPreamblesR17StartPreambleForThisPartitionR17Constraints)
		if err != nil {
			return err
		}
		ie.StartPreambleForThisPartition_r17 = v1
	}

	// 5. numberOfPreamblesPerSSB-ForThisPartition-r17: integer
	{
		v2, err := d.DecodeInteger(featureCombinationPreamblesR17NumberOfPreamblesPerSSBForThisPartitionR17Constraints)
		if err != nil {
			return err
		}
		ie.NumberOfPreamblesPerSSB_ForThisPartition_r17 = v2
	}

	// 6. ssb-SharedRO-MaskIndex-r17: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(featureCombinationPreamblesR17SsbSharedROMaskIndexR17Constraints)
			if err != nil {
				return err
			}
			ie.Ssb_SharedRO_MaskIndex_r17 = &v
		}
	}

	// 7. groupBconfigured-r17: sequence
	{
		if seq.IsComponentPresent(4) {
			ie.GroupBconfigured_r17 = &struct {
				Ra_SizeGroupA_r17              int64
				MessagePowerOffsetGroupB_r17   int64
				NumberOfRA_PreamblesGroupA_r17 int64
			}{}
			c := ie.GroupBconfigured_r17
			{
				v, err := d.DecodeEnumerated(featureCombinationPreamblesR17GroupBconfiguredR17RaSizeGroupAR17Constraints)
				if err != nil {
					return err
				}
				c.Ra_SizeGroupA_r17 = v
			}
			{
				v, err := d.DecodeEnumerated(featureCombinationPreamblesR17GroupBconfiguredR17MessagePowerOffsetGroupBR17Constraints)
				if err != nil {
					return err
				}
				c.MessagePowerOffsetGroupB_r17 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 64))
				if err != nil {
					return err
				}
				c.NumberOfRA_PreamblesGroupA_r17 = v
			}
		}
	}

	// 8. separateMsgA-PUSCH-Config-r17: ref
	{
		if seq.IsComponentPresent(5) {
			ie.SeparateMsgA_PUSCH_Config_r17 = new(MsgA_PUSCH_Config_r16)
			if err := ie.SeparateMsgA_PUSCH_Config_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 9. msgA-RSRP-Threshold-r17: ref
	{
		if seq.IsComponentPresent(6) {
			ie.MsgA_RSRP_Threshold_r17 = new(RSRP_Range)
			if err := ie.MsgA_RSRP_Threshold_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 10. rsrp-ThresholdSSB-r17: ref
	{
		if seq.IsComponentPresent(7) {
			ie.Rsrp_ThresholdSSB_r17 = new(RSRP_Range)
			if err := ie.Rsrp_ThresholdSSB_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 11. deltaPreamble-r17: integer
	{
		if seq.IsComponentPresent(8) {
			v, err := d.DecodeInteger(featureCombinationPreamblesR17DeltaPreambleR17Constraints)
			if err != nil {
				return err
			}
			ie.DeltaPreamble_r17 = &v
		}
	}

	// Extension additions as open types.
	extValues, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	extValueIndex := 0

	// Extension group 0:
	if seq.IsExtensionPresent(0) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "msg1-RepetitionNum-r18", Optional: true},
				{Name: "msg1-RepetitionTimeOffsetROGroup-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(featureCombinationPreamblesR17ExtMsg1RepetitionNumR18Constraints)
			if err != nil {
				return err
			}
			ie.Msg1_RepetitionNum_r18 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(featureCombinationPreamblesR17ExtMsg1RepetitionTimeOffsetROGroupR18Constraints)
			if err != nil {
				return err
			}
			ie.Msg1_RepetitionTimeOffsetROGroup_r18 = &v
		}
	}

	return nil
}
