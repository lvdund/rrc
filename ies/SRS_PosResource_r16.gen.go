// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SRS-PosResource-r16 (line 15535).

var sRSPosResourceR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "srs-PosResourceId-r16"},
		{Name: "transmissionComb-r16"},
		{Name: "resourceMapping-r16"},
		{Name: "freqDomainShift-r16"},
		{Name: "freqHopping-r16"},
		{Name: "groupOrSequenceHopping-r16"},
		{Name: "resourceType-r16"},
		{Name: "sequenceId-r16"},
		{Name: "spatialRelationInfoPos-r16", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var sRS_PosResource_r16TransmissionCombR16Constraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "n2-r16"},
		{Name: "n4-r16"},
		{Name: "n8-r16"},
	},
}

const (
	SRS_PosResource_r16_TransmissionComb_r16_N2_r16 = 0
	SRS_PosResource_r16_TransmissionComb_r16_N4_r16 = 1
	SRS_PosResource_r16_TransmissionComb_r16_N8_r16 = 2
)

var sRSPosResourceR16FreqDomainShiftR16Constraints = per.Constrained(0, 268)

const (
	SRS_PosResource_r16_GroupOrSequenceHopping_r16_Neither         = 0
	SRS_PosResource_r16_GroupOrSequenceHopping_r16_GroupHopping    = 1
	SRS_PosResource_r16_GroupOrSequenceHopping_r16_SequenceHopping = 2
)

var sRSPosResourceR16GroupOrSequenceHoppingR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_PosResource_r16_GroupOrSequenceHopping_r16_Neither, SRS_PosResource_r16_GroupOrSequenceHopping_r16_GroupHopping, SRS_PosResource_r16_GroupOrSequenceHopping_r16_SequenceHopping},
}

var sRS_PosResource_r16ResourceTypeR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "aperiodic-r16"},
		{Name: "semi-persistent-r16"},
		{Name: "periodic-r16"},
	},
}

const (
	SRS_PosResource_r16_ResourceType_r16_Aperiodic_r16       = 0
	SRS_PosResource_r16_ResourceType_r16_Semi_Persistent_r16 = 1
	SRS_PosResource_r16_ResourceType_r16_Periodic_r16        = 2
)

var sRSPosResourceR16SequenceIdR16Constraints = per.Constrained(0, 65535)

const (
	SRS_PosResource_r16_ResourceMapping_r16_NrofSymbols_r16_N1  = 0
	SRS_PosResource_r16_ResourceMapping_r16_NrofSymbols_r16_N2  = 1
	SRS_PosResource_r16_ResourceMapping_r16_NrofSymbols_r16_N4  = 2
	SRS_PosResource_r16_ResourceMapping_r16_NrofSymbols_r16_N8  = 3
	SRS_PosResource_r16_ResourceMapping_r16_NrofSymbols_r16_N12 = 4
)

var sRSPosResourceR16ResourceMappingR16NrofSymbolsR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_PosResource_r16_ResourceMapping_r16_NrofSymbols_r16_N1, SRS_PosResource_r16_ResourceMapping_r16_NrofSymbols_r16_N2, SRS_PosResource_r16_ResourceMapping_r16_NrofSymbols_r16_N4, SRS_PosResource_r16_ResourceMapping_r16_NrofSymbols_r16_N8, SRS_PosResource_r16_ResourceMapping_r16_NrofSymbols_r16_N12},
}

var sRSPosResourceR16FreqHoppingR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "c-SRS-r16"},
	},
}

var sRSPosResourceR16ResourceTypeR16AperiodicR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "slotOffset-r16", Optional: true},
	},
}

var sRSPosResourceR16ResourceTypeR16SemiPersistentR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "periodicityAndOffset-sp-r16"},
	},
}

var sRSPosResourceR16ResourceTypeR16PeriodicR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "periodicityAndOffset-p-r16"},
	},
}

type SRS_PosResource_r16 struct {
	Srs_PosResourceId_r16 SRS_PosResourceId_r16
	TransmissionComb_r16  struct {
		Choice int
		N2_r16 *struct {
			CombOffset_N2_r16  int64
			CyclicShift_N2_r16 int64
		}
		N4_r16 *struct {
			CombOffset_N4_r16  int64
			CyclicShift_N4_r16 int64
		}
		N8_r16 *struct {
			CombOffset_N8_r16  int64
			CyclicShift_N8_r16 int64
		}
	}
	ResourceMapping_r16 struct {
		StartPosition_r16 int64
		NrofSymbols_r16   int64
	}
	FreqDomainShift_r16        int64
	FreqHopping_r16            struct{ C_SRS_r16 int64 }
	GroupOrSequenceHopping_r16 int64
	ResourceType_r16           struct {
		Choice              int
		Aperiodic_r16       *struct{ SlotOffset_r16 *int64 }
		Semi_Persistent_r16 *struct{ PeriodicityAndOffset_Sp_r16 SRS_PeriodicityAndOffset_r16 }
		Periodic_r16        *struct{ PeriodicityAndOffset_P_r16 SRS_PeriodicityAndOffset_r16 }
	}
	SequenceId_r16             int64
	SpatialRelationInfoPos_r16 *SRS_SpatialRelationInfoPos_r16
	TxHoppingConfig_r18        *TxHoppingConfig_r18
}

func (ie *SRS_PosResource_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sRSPosResourceR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.TxHoppingConfig_r18 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SpatialRelationInfoPos_r16 != nil}); err != nil {
		return err
	}

	// 3. srs-PosResourceId-r16: ref
	{
		if err := ie.Srs_PosResourceId_r16.Encode(e); err != nil {
			return err
		}
	}

	// 4. transmissionComb-r16: choice
	{
		choiceEnc := e.NewChoiceEncoder(sRS_PosResource_r16TransmissionCombR16Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.TransmissionComb_r16.Choice), false, nil); err != nil {
			return err
		}
		switch ie.TransmissionComb_r16.Choice {
		case SRS_PosResource_r16_TransmissionComb_r16_N2_r16:
			c := (*ie.TransmissionComb_r16.N2_r16)
			if err := e.EncodeInteger(c.CombOffset_N2_r16, per.Constrained(0, 1)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.CyclicShift_N2_r16, per.Constrained(0, 7)); err != nil {
				return err
			}
		case SRS_PosResource_r16_TransmissionComb_r16_N4_r16:
			c := (*ie.TransmissionComb_r16.N4_r16)
			if err := e.EncodeInteger(c.CombOffset_N4_r16, per.Constrained(0, 3)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.CyclicShift_N4_r16, per.Constrained(0, 11)); err != nil {
				return err
			}
		case SRS_PosResource_r16_TransmissionComb_r16_N8_r16:
			c := (*ie.TransmissionComb_r16.N8_r16)
			if err := e.EncodeInteger(c.CombOffset_N8_r16, per.Constrained(0, 7)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.CyclicShift_N8_r16, per.Constrained(0, 5)); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.TransmissionComb_r16.Choice), Constraint: "index out of range"}
		}
	}

	// 5. resourceMapping-r16: sequence
	{
		{
			c := &ie.ResourceMapping_r16
			if err := e.EncodeInteger(c.StartPosition_r16, per.Constrained(0, 13)); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.NrofSymbols_r16, sRSPosResourceR16ResourceMappingR16NrofSymbolsR16Constraints); err != nil {
				return err
			}
		}
	}

	// 6. freqDomainShift-r16: integer
	{
		if err := e.EncodeInteger(ie.FreqDomainShift_r16, sRSPosResourceR16FreqDomainShiftR16Constraints); err != nil {
			return err
		}
	}

	// 7. freqHopping-r16: sequence
	{
		{
			c := &ie.FreqHopping_r16
			sRSPosResourceR16FreqHoppingR16Seq := e.NewSequenceEncoder(sRSPosResourceR16FreqHoppingR16Constraints)
			if err := sRSPosResourceR16FreqHoppingR16Seq.EncodeExtensionBit(false); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.C_SRS_r16, per.Constrained(0, 63)); err != nil {
				return err
			}
		}
	}

	// 8. groupOrSequenceHopping-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.GroupOrSequenceHopping_r16, sRSPosResourceR16GroupOrSequenceHoppingR16Constraints); err != nil {
			return err
		}
	}

	// 9. resourceType-r16: choice
	{
		choiceEnc := e.NewChoiceEncoder(sRS_PosResource_r16ResourceTypeR16Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.ResourceType_r16.Choice), false, nil); err != nil {
			return err
		}
		switch ie.ResourceType_r16.Choice {
		case SRS_PosResource_r16_ResourceType_r16_Aperiodic_r16:
			c := (*ie.ResourceType_r16.Aperiodic_r16)
			sRSPosResourceR16ResourceTypeR16AperiodicR16Seq := e.NewSequenceEncoder(sRSPosResourceR16ResourceTypeR16AperiodicR16Constraints)
			if err := sRSPosResourceR16ResourceTypeR16AperiodicR16Seq.EncodeExtensionBit(false); err != nil {
				return err
			}
			if err := sRSPosResourceR16ResourceTypeR16AperiodicR16Seq.EncodePreamble([]bool{c.SlotOffset_r16 != nil}); err != nil {
				return err
			}
			if c.SlotOffset_r16 != nil {
				if err := e.EncodeInteger((*c.SlotOffset_r16), per.Constrained(1, 32)); err != nil {
					return err
				}
			}
		case SRS_PosResource_r16_ResourceType_r16_Semi_Persistent_r16:
			c := (*ie.ResourceType_r16.Semi_Persistent_r16)
			sRSPosResourceR16ResourceTypeR16SemiPersistentR16Seq := e.NewSequenceEncoder(sRSPosResourceR16ResourceTypeR16SemiPersistentR16Constraints)
			if err := sRSPosResourceR16ResourceTypeR16SemiPersistentR16Seq.EncodeExtensionBit(false); err != nil {
				return err
			}
			if err := c.PeriodicityAndOffset_Sp_r16.Encode(e); err != nil {
				return err
			}
		case SRS_PosResource_r16_ResourceType_r16_Periodic_r16:
			c := (*ie.ResourceType_r16.Periodic_r16)
			sRSPosResourceR16ResourceTypeR16PeriodicR16Seq := e.NewSequenceEncoder(sRSPosResourceR16ResourceTypeR16PeriodicR16Constraints)
			if err := sRSPosResourceR16ResourceTypeR16PeriodicR16Seq.EncodeExtensionBit(false); err != nil {
				return err
			}
			if err := c.PeriodicityAndOffset_P_r16.Encode(e); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.ResourceType_r16.Choice), Constraint: "index out of range"}
		}
	}

	// 10. sequenceId-r16: integer
	{
		if err := e.EncodeInteger(ie.SequenceId_r16, sRSPosResourceR16SequenceIdR16Constraints); err != nil {
			return err
		}
	}

	// 11. spatialRelationInfoPos-r16: ref
	{
		if ie.SpatialRelationInfoPos_r16 != nil {
			if err := ie.SpatialRelationInfoPos_r16.Encode(e); err != nil {
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
					{Name: "txHoppingConfig-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.TxHoppingConfig_r18 != nil}); err != nil {
				return err
			}

			if ie.TxHoppingConfig_r18 != nil {
				if err := ie.TxHoppingConfig_r18.Encode(ex); err != nil {
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

func (ie *SRS_PosResource_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sRSPosResourceR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. srs-PosResourceId-r16: ref
	{
		if err := ie.Srs_PosResourceId_r16.Decode(d); err != nil {
			return err
		}
	}

	// 4. transmissionComb-r16: choice
	{
		choiceDec := d.NewChoiceDecoder(sRS_PosResource_r16TransmissionCombR16Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.TransmissionComb_r16.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case SRS_PosResource_r16_TransmissionComb_r16_N2_r16:
			ie.TransmissionComb_r16.N2_r16 = &struct {
				CombOffset_N2_r16  int64
				CyclicShift_N2_r16 int64
			}{}
			c := (*ie.TransmissionComb_r16.N2_r16)
			{
				v, err := d.DecodeInteger(per.Constrained(0, 1))
				if err != nil {
					return err
				}
				c.CombOffset_N2_r16 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(0, 7))
				if err != nil {
					return err
				}
				c.CyclicShift_N2_r16 = v
			}
		case SRS_PosResource_r16_TransmissionComb_r16_N4_r16:
			ie.TransmissionComb_r16.N4_r16 = &struct {
				CombOffset_N4_r16  int64
				CyclicShift_N4_r16 int64
			}{}
			c := (*ie.TransmissionComb_r16.N4_r16)
			{
				v, err := d.DecodeInteger(per.Constrained(0, 3))
				if err != nil {
					return err
				}
				c.CombOffset_N4_r16 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(0, 11))
				if err != nil {
					return err
				}
				c.CyclicShift_N4_r16 = v
			}
		case SRS_PosResource_r16_TransmissionComb_r16_N8_r16:
			ie.TransmissionComb_r16.N8_r16 = &struct {
				CombOffset_N8_r16  int64
				CyclicShift_N8_r16 int64
			}{}
			c := (*ie.TransmissionComb_r16.N8_r16)
			{
				v, err := d.DecodeInteger(per.Constrained(0, 7))
				if err != nil {
					return err
				}
				c.CombOffset_N8_r16 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(0, 5))
				if err != nil {
					return err
				}
				c.CyclicShift_N8_r16 = v
			}
		}
	}

	// 5. resourceMapping-r16: sequence
	{
		{
			c := &ie.ResourceMapping_r16
			{
				v, err := d.DecodeInteger(per.Constrained(0, 13))
				if err != nil {
					return err
				}
				c.StartPosition_r16 = v
			}
			{
				v, err := d.DecodeEnumerated(sRSPosResourceR16ResourceMappingR16NrofSymbolsR16Constraints)
				if err != nil {
					return err
				}
				c.NrofSymbols_r16 = v
			}
		}
	}

	// 6. freqDomainShift-r16: integer
	{
		v3, err := d.DecodeInteger(sRSPosResourceR16FreqDomainShiftR16Constraints)
		if err != nil {
			return err
		}
		ie.FreqDomainShift_r16 = v3
	}

	// 7. freqHopping-r16: sequence
	{
		{
			c := &ie.FreqHopping_r16
			sRSPosResourceR16FreqHoppingR16Seq := d.NewSequenceDecoder(sRSPosResourceR16FreqHoppingR16Constraints)
			if err := sRSPosResourceR16FreqHoppingR16Seq.DecodeExtensionBit(); err != nil {
				return err
			}
			{
				v, err := d.DecodeInteger(per.Constrained(0, 63))
				if err != nil {
					return err
				}
				c.C_SRS_r16 = v
			}
		}
	}

	// 8. groupOrSequenceHopping-r16: enumerated
	{
		v5, err := d.DecodeEnumerated(sRSPosResourceR16GroupOrSequenceHoppingR16Constraints)
		if err != nil {
			return err
		}
		ie.GroupOrSequenceHopping_r16 = v5
	}

	// 9. resourceType-r16: choice
	{
		choiceDec := d.NewChoiceDecoder(sRS_PosResource_r16ResourceTypeR16Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.ResourceType_r16.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case SRS_PosResource_r16_ResourceType_r16_Aperiodic_r16:
			ie.ResourceType_r16.Aperiodic_r16 = &struct{ SlotOffset_r16 *int64 }{}
			c := (*ie.ResourceType_r16.Aperiodic_r16)
			sRSPosResourceR16ResourceTypeR16AperiodicR16Seq := d.NewSequenceDecoder(sRSPosResourceR16ResourceTypeR16AperiodicR16Constraints)
			if err := sRSPosResourceR16ResourceTypeR16AperiodicR16Seq.DecodeExtensionBit(); err != nil {
				return err
			}
			if err := sRSPosResourceR16ResourceTypeR16AperiodicR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if sRSPosResourceR16ResourceTypeR16AperiodicR16Seq.IsComponentPresent(0) {
				c.SlotOffset_r16 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 32))
				if err != nil {
					return err
				}
				(*c.SlotOffset_r16) = v
			}
		case SRS_PosResource_r16_ResourceType_r16_Semi_Persistent_r16:
			ie.ResourceType_r16.Semi_Persistent_r16 = &struct{ PeriodicityAndOffset_Sp_r16 SRS_PeriodicityAndOffset_r16 }{}
			c := (*ie.ResourceType_r16.Semi_Persistent_r16)
			sRSPosResourceR16ResourceTypeR16SemiPersistentR16Seq := d.NewSequenceDecoder(sRSPosResourceR16ResourceTypeR16SemiPersistentR16Constraints)
			if err := sRSPosResourceR16ResourceTypeR16SemiPersistentR16Seq.DecodeExtensionBit(); err != nil {
				return err
			}
			{
				if err := c.PeriodicityAndOffset_Sp_r16.Decode(d); err != nil {
					return err
				}
			}
		case SRS_PosResource_r16_ResourceType_r16_Periodic_r16:
			ie.ResourceType_r16.Periodic_r16 = &struct{ PeriodicityAndOffset_P_r16 SRS_PeriodicityAndOffset_r16 }{}
			c := (*ie.ResourceType_r16.Periodic_r16)
			sRSPosResourceR16ResourceTypeR16PeriodicR16Seq := d.NewSequenceDecoder(sRSPosResourceR16ResourceTypeR16PeriodicR16Constraints)
			if err := sRSPosResourceR16ResourceTypeR16PeriodicR16Seq.DecodeExtensionBit(); err != nil {
				return err
			}
			{
				if err := c.PeriodicityAndOffset_P_r16.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 10. sequenceId-r16: integer
	{
		v7, err := d.DecodeInteger(sRSPosResourceR16SequenceIdR16Constraints)
		if err != nil {
			return err
		}
		ie.SequenceId_r16 = v7
	}

	// 11. spatialRelationInfoPos-r16: ref
	{
		if seq.IsComponentPresent(8) {
			ie.SpatialRelationInfoPos_r16 = new(SRS_SpatialRelationInfoPos_r16)
			if err := ie.SpatialRelationInfoPos_r16.Decode(d); err != nil {
				return err
			}
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
				{Name: "txHoppingConfig-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.TxHoppingConfig_r18 = new(TxHoppingConfig_r18)
			if err := ie.TxHoppingConfig_r18.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
