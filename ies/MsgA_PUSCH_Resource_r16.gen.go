// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MsgA-PUSCH-Resource-r16 (line 10129).

var msgAPUSCHResourceR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "msgA-MCS-r16"},
		{Name: "nrofSlotsMsgA-PUSCH-r16"},
		{Name: "nrofMsgA-PO-PerSlot-r16"},
		{Name: "msgA-PUSCH-TimeDomainOffset-r16"},
		{Name: "msgA-PUSCH-TimeDomainAllocation-r16", Optional: true},
		{Name: "startSymbolAndLengthMsgA-PO-r16", Optional: true},
		{Name: "mappingTypeMsgA-PUSCH-r16", Optional: true},
		{Name: "guardPeriodMsgA-PUSCH-r16", Optional: true},
		{Name: "guardBandMsgA-PUSCH-r16"},
		{Name: "frequencyStartMsgA-PUSCH-r16"},
		{Name: "nrofPRBs-PerMsgA-PO-r16"},
		{Name: "nrofMsgA-PO-FDM-r16"},
		{Name: "msgA-IntraSlotFrequencyHopping-r16", Optional: true},
		{Name: "msgA-HoppingBits-r16", Optional: true},
		{Name: "msgA-DMRS-Config-r16"},
		{Name: "nrofDMRS-Sequences-r16"},
		{Name: "msgA-Alpha-r16", Optional: true},
		{Name: "interlaceIndexFirstPO-MsgA-PUSCH-r16", Optional: true},
		{Name: "nrofInterlacesPerMsgA-PO-r16", Optional: true},
	},
}

var msgAPUSCHResourceR16MsgAMCSR16Constraints = per.Constrained(0, 15)

var msgAPUSCHResourceR16NrofSlotsMsgAPUSCHR16Constraints = per.Constrained(1, 4)

const (
	MsgA_PUSCH_Resource_r16_NrofMsgA_PO_PerSlot_r16_One   = 0
	MsgA_PUSCH_Resource_r16_NrofMsgA_PO_PerSlot_r16_Two   = 1
	MsgA_PUSCH_Resource_r16_NrofMsgA_PO_PerSlot_r16_Three = 2
	MsgA_PUSCH_Resource_r16_NrofMsgA_PO_PerSlot_r16_Six   = 3
)

var msgAPUSCHResourceR16NrofMsgAPOPerSlotR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MsgA_PUSCH_Resource_r16_NrofMsgA_PO_PerSlot_r16_One, MsgA_PUSCH_Resource_r16_NrofMsgA_PO_PerSlot_r16_Two, MsgA_PUSCH_Resource_r16_NrofMsgA_PO_PerSlot_r16_Three, MsgA_PUSCH_Resource_r16_NrofMsgA_PO_PerSlot_r16_Six},
}

var msgAPUSCHResourceR16MsgAPUSCHTimeDomainOffsetR16Constraints = per.Constrained(1, 32)

var msgAPUSCHResourceR16MsgAPUSCHTimeDomainAllocationR16Constraints = per.Constrained(1, common.MaxNrofUL_Allocations)

var msgAPUSCHResourceR16StartSymbolAndLengthMsgAPOR16Constraints = per.Constrained(0, 127)

const (
	MsgA_PUSCH_Resource_r16_MappingTypeMsgA_PUSCH_r16_TypeA = 0
	MsgA_PUSCH_Resource_r16_MappingTypeMsgA_PUSCH_r16_TypeB = 1
)

var msgAPUSCHResourceR16MappingTypeMsgAPUSCHR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MsgA_PUSCH_Resource_r16_MappingTypeMsgA_PUSCH_r16_TypeA, MsgA_PUSCH_Resource_r16_MappingTypeMsgA_PUSCH_r16_TypeB},
}

var msgAPUSCHResourceR16GuardPeriodMsgAPUSCHR16Constraints = per.Constrained(0, 3)

var msgAPUSCHResourceR16GuardBandMsgAPUSCHR16Constraints = per.Constrained(0, 1)

var msgAPUSCHResourceR16FrequencyStartMsgAPUSCHR16Constraints = per.Constrained(0, common.MaxNrofPhysicalResourceBlocks_1)

var msgAPUSCHResourceR16NrofPRBsPerMsgAPOR16Constraints = per.Constrained(1, 32)

const (
	MsgA_PUSCH_Resource_r16_NrofMsgA_PO_FDM_r16_One   = 0
	MsgA_PUSCH_Resource_r16_NrofMsgA_PO_FDM_r16_Two   = 1
	MsgA_PUSCH_Resource_r16_NrofMsgA_PO_FDM_r16_Four  = 2
	MsgA_PUSCH_Resource_r16_NrofMsgA_PO_FDM_r16_Eight = 3
)

var msgAPUSCHResourceR16NrofMsgAPOFDMR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MsgA_PUSCH_Resource_r16_NrofMsgA_PO_FDM_r16_One, MsgA_PUSCH_Resource_r16_NrofMsgA_PO_FDM_r16_Two, MsgA_PUSCH_Resource_r16_NrofMsgA_PO_FDM_r16_Four, MsgA_PUSCH_Resource_r16_NrofMsgA_PO_FDM_r16_Eight},
}

const (
	MsgA_PUSCH_Resource_r16_MsgA_IntraSlotFrequencyHopping_r16_Enabled = 0
)

var msgAPUSCHResourceR16MsgAIntraSlotFrequencyHoppingR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MsgA_PUSCH_Resource_r16_MsgA_IntraSlotFrequencyHopping_r16_Enabled},
}

var msgAPUSCHResourceR16MsgAHoppingBitsR16Constraints = per.FixedSize(2)

var msgAPUSCHResourceR16NrofDMRSSequencesR16Constraints = per.Constrained(1, 2)

const (
	MsgA_PUSCH_Resource_r16_MsgA_Alpha_r16_Alpha0  = 0
	MsgA_PUSCH_Resource_r16_MsgA_Alpha_r16_Alpha04 = 1
	MsgA_PUSCH_Resource_r16_MsgA_Alpha_r16_Alpha05 = 2
	MsgA_PUSCH_Resource_r16_MsgA_Alpha_r16_Alpha06 = 3
	MsgA_PUSCH_Resource_r16_MsgA_Alpha_r16_Alpha07 = 4
	MsgA_PUSCH_Resource_r16_MsgA_Alpha_r16_Alpha08 = 5
	MsgA_PUSCH_Resource_r16_MsgA_Alpha_r16_Alpha09 = 6
	MsgA_PUSCH_Resource_r16_MsgA_Alpha_r16_Alpha1  = 7
)

var msgAPUSCHResourceR16MsgAAlphaR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MsgA_PUSCH_Resource_r16_MsgA_Alpha_r16_Alpha0, MsgA_PUSCH_Resource_r16_MsgA_Alpha_r16_Alpha04, MsgA_PUSCH_Resource_r16_MsgA_Alpha_r16_Alpha05, MsgA_PUSCH_Resource_r16_MsgA_Alpha_r16_Alpha06, MsgA_PUSCH_Resource_r16_MsgA_Alpha_r16_Alpha07, MsgA_PUSCH_Resource_r16_MsgA_Alpha_r16_Alpha08, MsgA_PUSCH_Resource_r16_MsgA_Alpha_r16_Alpha09, MsgA_PUSCH_Resource_r16_MsgA_Alpha_r16_Alpha1},
}

var msgAPUSCHResourceR16InterlaceIndexFirstPOMsgAPUSCHR16Constraints = per.Constrained(1, 10)

var msgAPUSCHResourceR16NrofInterlacesPerMsgAPOR16Constraints = per.Constrained(1, 10)

type MsgA_PUSCH_Resource_r16 struct {
	MsgA_MCS_r16                         int64
	NrofSlotsMsgA_PUSCH_r16              int64
	NrofMsgA_PO_PerSlot_r16              int64
	MsgA_PUSCH_TimeDomainOffset_r16      int64
	MsgA_PUSCH_TimeDomainAllocation_r16  *int64
	StartSymbolAndLengthMsgA_PO_r16      *int64
	MappingTypeMsgA_PUSCH_r16            *int64
	GuardPeriodMsgA_PUSCH_r16            *int64
	GuardBandMsgA_PUSCH_r16              int64
	FrequencyStartMsgA_PUSCH_r16         int64
	NrofPRBs_PerMsgA_PO_r16              int64
	NrofMsgA_PO_FDM_r16                  int64
	MsgA_IntraSlotFrequencyHopping_r16   *int64
	MsgA_HoppingBits_r16                 *per.BitString
	MsgA_DMRS_Config_r16                 MsgA_DMRS_Config_r16
	NrofDMRS_Sequences_r16               int64
	MsgA_Alpha_r16                       *int64
	InterlaceIndexFirstPO_MsgA_PUSCH_r16 *int64
	NrofInterlacesPerMsgA_PO_r16         *int64
}

func (ie *MsgA_PUSCH_Resource_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(msgAPUSCHResourceR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MsgA_PUSCH_TimeDomainAllocation_r16 != nil, ie.StartSymbolAndLengthMsgA_PO_r16 != nil, ie.MappingTypeMsgA_PUSCH_r16 != nil, ie.GuardPeriodMsgA_PUSCH_r16 != nil, ie.MsgA_IntraSlotFrequencyHopping_r16 != nil, ie.MsgA_HoppingBits_r16 != nil, ie.MsgA_Alpha_r16 != nil, ie.InterlaceIndexFirstPO_MsgA_PUSCH_r16 != nil, ie.NrofInterlacesPerMsgA_PO_r16 != nil}); err != nil {
		return err
	}

	// 3. msgA-MCS-r16: integer
	{
		if err := e.EncodeInteger(ie.MsgA_MCS_r16, msgAPUSCHResourceR16MsgAMCSR16Constraints); err != nil {
			return err
		}
	}

	// 4. nrofSlotsMsgA-PUSCH-r16: integer
	{
		if err := e.EncodeInteger(ie.NrofSlotsMsgA_PUSCH_r16, msgAPUSCHResourceR16NrofSlotsMsgAPUSCHR16Constraints); err != nil {
			return err
		}
	}

	// 5. nrofMsgA-PO-PerSlot-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.NrofMsgA_PO_PerSlot_r16, msgAPUSCHResourceR16NrofMsgAPOPerSlotR16Constraints); err != nil {
			return err
		}
	}

	// 6. msgA-PUSCH-TimeDomainOffset-r16: integer
	{
		if err := e.EncodeInteger(ie.MsgA_PUSCH_TimeDomainOffset_r16, msgAPUSCHResourceR16MsgAPUSCHTimeDomainOffsetR16Constraints); err != nil {
			return err
		}
	}

	// 7. msgA-PUSCH-TimeDomainAllocation-r16: integer
	{
		if ie.MsgA_PUSCH_TimeDomainAllocation_r16 != nil {
			if err := e.EncodeInteger(*ie.MsgA_PUSCH_TimeDomainAllocation_r16, msgAPUSCHResourceR16MsgAPUSCHTimeDomainAllocationR16Constraints); err != nil {
				return err
			}
		}
	}

	// 8. startSymbolAndLengthMsgA-PO-r16: integer
	{
		if ie.StartSymbolAndLengthMsgA_PO_r16 != nil {
			if err := e.EncodeInteger(*ie.StartSymbolAndLengthMsgA_PO_r16, msgAPUSCHResourceR16StartSymbolAndLengthMsgAPOR16Constraints); err != nil {
				return err
			}
		}
	}

	// 9. mappingTypeMsgA-PUSCH-r16: enumerated
	{
		if ie.MappingTypeMsgA_PUSCH_r16 != nil {
			if err := e.EncodeEnumerated(*ie.MappingTypeMsgA_PUSCH_r16, msgAPUSCHResourceR16MappingTypeMsgAPUSCHR16Constraints); err != nil {
				return err
			}
		}
	}

	// 10. guardPeriodMsgA-PUSCH-r16: integer
	{
		if ie.GuardPeriodMsgA_PUSCH_r16 != nil {
			if err := e.EncodeInteger(*ie.GuardPeriodMsgA_PUSCH_r16, msgAPUSCHResourceR16GuardPeriodMsgAPUSCHR16Constraints); err != nil {
				return err
			}
		}
	}

	// 11. guardBandMsgA-PUSCH-r16: integer
	{
		if err := e.EncodeInteger(ie.GuardBandMsgA_PUSCH_r16, msgAPUSCHResourceR16GuardBandMsgAPUSCHR16Constraints); err != nil {
			return err
		}
	}

	// 12. frequencyStartMsgA-PUSCH-r16: integer
	{
		if err := e.EncodeInteger(ie.FrequencyStartMsgA_PUSCH_r16, msgAPUSCHResourceR16FrequencyStartMsgAPUSCHR16Constraints); err != nil {
			return err
		}
	}

	// 13. nrofPRBs-PerMsgA-PO-r16: integer
	{
		if err := e.EncodeInteger(ie.NrofPRBs_PerMsgA_PO_r16, msgAPUSCHResourceR16NrofPRBsPerMsgAPOR16Constraints); err != nil {
			return err
		}
	}

	// 14. nrofMsgA-PO-FDM-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.NrofMsgA_PO_FDM_r16, msgAPUSCHResourceR16NrofMsgAPOFDMR16Constraints); err != nil {
			return err
		}
	}

	// 15. msgA-IntraSlotFrequencyHopping-r16: enumerated
	{
		if ie.MsgA_IntraSlotFrequencyHopping_r16 != nil {
			if err := e.EncodeEnumerated(*ie.MsgA_IntraSlotFrequencyHopping_r16, msgAPUSCHResourceR16MsgAIntraSlotFrequencyHoppingR16Constraints); err != nil {
				return err
			}
		}
	}

	// 16. msgA-HoppingBits-r16: bit-string
	{
		if ie.MsgA_HoppingBits_r16 != nil {
			if err := e.EncodeBitString(*ie.MsgA_HoppingBits_r16, msgAPUSCHResourceR16MsgAHoppingBitsR16Constraints); err != nil {
				return err
			}
		}
	}

	// 17. msgA-DMRS-Config-r16: ref
	{
		if err := ie.MsgA_DMRS_Config_r16.Encode(e); err != nil {
			return err
		}
	}

	// 18. nrofDMRS-Sequences-r16: integer
	{
		if err := e.EncodeInteger(ie.NrofDMRS_Sequences_r16, msgAPUSCHResourceR16NrofDMRSSequencesR16Constraints); err != nil {
			return err
		}
	}

	// 19. msgA-Alpha-r16: enumerated
	{
		if ie.MsgA_Alpha_r16 != nil {
			if err := e.EncodeEnumerated(*ie.MsgA_Alpha_r16, msgAPUSCHResourceR16MsgAAlphaR16Constraints); err != nil {
				return err
			}
		}
	}

	// 20. interlaceIndexFirstPO-MsgA-PUSCH-r16: integer
	{
		if ie.InterlaceIndexFirstPO_MsgA_PUSCH_r16 != nil {
			if err := e.EncodeInteger(*ie.InterlaceIndexFirstPO_MsgA_PUSCH_r16, msgAPUSCHResourceR16InterlaceIndexFirstPOMsgAPUSCHR16Constraints); err != nil {
				return err
			}
		}
	}

	// 21. nrofInterlacesPerMsgA-PO-r16: integer
	{
		if ie.NrofInterlacesPerMsgA_PO_r16 != nil {
			if err := e.EncodeInteger(*ie.NrofInterlacesPerMsgA_PO_r16, msgAPUSCHResourceR16NrofInterlacesPerMsgAPOR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MsgA_PUSCH_Resource_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(msgAPUSCHResourceR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. msgA-MCS-r16: integer
	{
		v0, err := d.DecodeInteger(msgAPUSCHResourceR16MsgAMCSR16Constraints)
		if err != nil {
			return err
		}
		ie.MsgA_MCS_r16 = v0
	}

	// 4. nrofSlotsMsgA-PUSCH-r16: integer
	{
		v1, err := d.DecodeInteger(msgAPUSCHResourceR16NrofSlotsMsgAPUSCHR16Constraints)
		if err != nil {
			return err
		}
		ie.NrofSlotsMsgA_PUSCH_r16 = v1
	}

	// 5. nrofMsgA-PO-PerSlot-r16: enumerated
	{
		v2, err := d.DecodeEnumerated(msgAPUSCHResourceR16NrofMsgAPOPerSlotR16Constraints)
		if err != nil {
			return err
		}
		ie.NrofMsgA_PO_PerSlot_r16 = v2
	}

	// 6. msgA-PUSCH-TimeDomainOffset-r16: integer
	{
		v3, err := d.DecodeInteger(msgAPUSCHResourceR16MsgAPUSCHTimeDomainOffsetR16Constraints)
		if err != nil {
			return err
		}
		ie.MsgA_PUSCH_TimeDomainOffset_r16 = v3
	}

	// 7. msgA-PUSCH-TimeDomainAllocation-r16: integer
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeInteger(msgAPUSCHResourceR16MsgAPUSCHTimeDomainAllocationR16Constraints)
			if err != nil {
				return err
			}
			ie.MsgA_PUSCH_TimeDomainAllocation_r16 = &v
		}
	}

	// 8. startSymbolAndLengthMsgA-PO-r16: integer
	{
		if seq.IsComponentPresent(5) {
			v, err := d.DecodeInteger(msgAPUSCHResourceR16StartSymbolAndLengthMsgAPOR16Constraints)
			if err != nil {
				return err
			}
			ie.StartSymbolAndLengthMsgA_PO_r16 = &v
		}
	}

	// 9. mappingTypeMsgA-PUSCH-r16: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(msgAPUSCHResourceR16MappingTypeMsgAPUSCHR16Constraints)
			if err != nil {
				return err
			}
			ie.MappingTypeMsgA_PUSCH_r16 = &idx
		}
	}

	// 10. guardPeriodMsgA-PUSCH-r16: integer
	{
		if seq.IsComponentPresent(7) {
			v, err := d.DecodeInteger(msgAPUSCHResourceR16GuardPeriodMsgAPUSCHR16Constraints)
			if err != nil {
				return err
			}
			ie.GuardPeriodMsgA_PUSCH_r16 = &v
		}
	}

	// 11. guardBandMsgA-PUSCH-r16: integer
	{
		v8, err := d.DecodeInteger(msgAPUSCHResourceR16GuardBandMsgAPUSCHR16Constraints)
		if err != nil {
			return err
		}
		ie.GuardBandMsgA_PUSCH_r16 = v8
	}

	// 12. frequencyStartMsgA-PUSCH-r16: integer
	{
		v9, err := d.DecodeInteger(msgAPUSCHResourceR16FrequencyStartMsgAPUSCHR16Constraints)
		if err != nil {
			return err
		}
		ie.FrequencyStartMsgA_PUSCH_r16 = v9
	}

	// 13. nrofPRBs-PerMsgA-PO-r16: integer
	{
		v10, err := d.DecodeInteger(msgAPUSCHResourceR16NrofPRBsPerMsgAPOR16Constraints)
		if err != nil {
			return err
		}
		ie.NrofPRBs_PerMsgA_PO_r16 = v10
	}

	// 14. nrofMsgA-PO-FDM-r16: enumerated
	{
		v11, err := d.DecodeEnumerated(msgAPUSCHResourceR16NrofMsgAPOFDMR16Constraints)
		if err != nil {
			return err
		}
		ie.NrofMsgA_PO_FDM_r16 = v11
	}

	// 15. msgA-IntraSlotFrequencyHopping-r16: enumerated
	{
		if seq.IsComponentPresent(12) {
			idx, err := d.DecodeEnumerated(msgAPUSCHResourceR16MsgAIntraSlotFrequencyHoppingR16Constraints)
			if err != nil {
				return err
			}
			ie.MsgA_IntraSlotFrequencyHopping_r16 = &idx
		}
	}

	// 16. msgA-HoppingBits-r16: bit-string
	{
		if seq.IsComponentPresent(13) {
			v, err := d.DecodeBitString(msgAPUSCHResourceR16MsgAHoppingBitsR16Constraints)
			if err != nil {
				return err
			}
			ie.MsgA_HoppingBits_r16 = &v
		}
	}

	// 17. msgA-DMRS-Config-r16: ref
	{
		if err := ie.MsgA_DMRS_Config_r16.Decode(d); err != nil {
			return err
		}
	}

	// 18. nrofDMRS-Sequences-r16: integer
	{
		v15, err := d.DecodeInteger(msgAPUSCHResourceR16NrofDMRSSequencesR16Constraints)
		if err != nil {
			return err
		}
		ie.NrofDMRS_Sequences_r16 = v15
	}

	// 19. msgA-Alpha-r16: enumerated
	{
		if seq.IsComponentPresent(16) {
			idx, err := d.DecodeEnumerated(msgAPUSCHResourceR16MsgAAlphaR16Constraints)
			if err != nil {
				return err
			}
			ie.MsgA_Alpha_r16 = &idx
		}
	}

	// 20. interlaceIndexFirstPO-MsgA-PUSCH-r16: integer
	{
		if seq.IsComponentPresent(17) {
			v, err := d.DecodeInteger(msgAPUSCHResourceR16InterlaceIndexFirstPOMsgAPUSCHR16Constraints)
			if err != nil {
				return err
			}
			ie.InterlaceIndexFirstPO_MsgA_PUSCH_r16 = &v
		}
	}

	// 21. nrofInterlacesPerMsgA-PO-r16: integer
	{
		if seq.IsComponentPresent(18) {
			v, err := d.DecodeInteger(msgAPUSCHResourceR16NrofInterlacesPerMsgAPOR16Constraints)
			if err != nil {
				return err
			}
			ie.NrofInterlacesPerMsgA_PO_r16 = &v
		}
	}

	return nil
}
