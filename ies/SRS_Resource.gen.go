// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SRS-Resource (line 15434).

var sRSResourceConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "srs-ResourceId"},
		{Name: "nrofSRS-Ports"},
		{Name: "ptrs-PortIndex", Optional: true},
		{Name: "transmissionComb"},
		{Name: "resourceMapping"},
		{Name: "freqDomainPosition"},
		{Name: "freqDomainShift"},
		{Name: "freqHopping"},
		{Name: "groupOrSequenceHopping"},
		{Name: "resourceType"},
		{Name: "sequenceId"},
		{Name: "spatialRelationInfo", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
		{Name: "extension-addition-3", Optional: true},
	},
}

const (
	SRS_Resource_NrofSRS_Ports_Port1  = 0
	SRS_Resource_NrofSRS_Ports_Ports2 = 1
	SRS_Resource_NrofSRS_Ports_Ports4 = 2
)

var sRSResourceNrofSRSPortsConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_Resource_NrofSRS_Ports_Port1, SRS_Resource_NrofSRS_Ports_Ports2, SRS_Resource_NrofSRS_Ports_Ports4},
}

const (
	SRS_Resource_Ptrs_PortIndex_N0 = 0
	SRS_Resource_Ptrs_PortIndex_N1 = 1
)

var sRSResourcePtrsPortIndexConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_Resource_Ptrs_PortIndex_N0, SRS_Resource_Ptrs_PortIndex_N1},
}

var sRS_ResourceTransmissionCombConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "n2"},
		{Name: "n4"},
	},
}

const (
	SRS_Resource_TransmissionComb_N2 = 0
	SRS_Resource_TransmissionComb_N4 = 1
)

var sRSResourceFreqDomainPositionConstraints = per.Constrained(0, 67)

var sRSResourceFreqDomainShiftConstraints = per.Constrained(0, 268)

const (
	SRS_Resource_GroupOrSequenceHopping_Neither         = 0
	SRS_Resource_GroupOrSequenceHopping_GroupHopping    = 1
	SRS_Resource_GroupOrSequenceHopping_SequenceHopping = 2
)

var sRSResourceGroupOrSequenceHoppingConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_Resource_GroupOrSequenceHopping_Neither, SRS_Resource_GroupOrSequenceHopping_GroupHopping, SRS_Resource_GroupOrSequenceHopping_SequenceHopping},
}

var sRS_ResourceResourceTypeConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "aperiodic"},
		{Name: "semi-persistent"},
		{Name: "periodic"},
	},
}

const (
	SRS_Resource_ResourceType_Aperiodic       = 0
	SRS_Resource_ResourceType_Semi_Persistent = 1
	SRS_Resource_ResourceType_Periodic        = 2
)

var sRSResourceSequenceIdConstraints = per.Constrained(0, 1023)

var sRSResourceExtSpatialRelationInfoPDCR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SRS_Resource_Ext_SpatialRelationInfo_PDC_r17_Release = 0
	SRS_Resource_Ext_SpatialRelationInfo_PDC_r17_Setup   = 1
)

var sRSResourceExtSrsTCIStateR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "srs-UL-TCI-State"},
		{Name: "srs-DLorJointTCI-State"},
	},
}

const (
	SRS_Resource_Ext_Srs_TCI_State_r17_Srs_UL_TCI_State       = 0
	SRS_Resource_Ext_Srs_TCI_State_r17_Srs_DLorJointTCI_State = 1
)

const (
	SRS_Resource_Ext_RepetitionFactor_v1730_N3 = 0
)

var sRSResourceExtRepetitionFactorV1730Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_Resource_Ext_RepetitionFactor_v1730_N3},
}

const (
	SRS_Resource_Ext_NrofSRS_Ports_N8_r18_Ports8    = 0
	SRS_Resource_Ext_NrofSRS_Ports_N8_r18_Ports8tdm = 1
)

var sRSResourceExtNrofSRSPortsN8R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_Resource_Ext_NrofSRS_Ports_N8_r18_Ports8, SRS_Resource_Ext_NrofSRS_Ports_N8_r18_Ports8tdm},
}

const (
	SRS_Resource_ResourceMapping_NrofSymbols_N1 = 0
	SRS_Resource_ResourceMapping_NrofSymbols_N2 = 1
	SRS_Resource_ResourceMapping_NrofSymbols_N4 = 2
)

var sRSResourceResourceMappingNrofSymbolsConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_Resource_ResourceMapping_NrofSymbols_N1, SRS_Resource_ResourceMapping_NrofSymbols_N2, SRS_Resource_ResourceMapping_NrofSymbols_N4},
}

const (
	SRS_Resource_ResourceMapping_RepetitionFactor_N1 = 0
	SRS_Resource_ResourceMapping_RepetitionFactor_N2 = 1
	SRS_Resource_ResourceMapping_RepetitionFactor_N4 = 2
)

var sRSResourceResourceMappingRepetitionFactorConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_Resource_ResourceMapping_RepetitionFactor_N1, SRS_Resource_ResourceMapping_RepetitionFactor_N2, SRS_Resource_ResourceMapping_RepetitionFactor_N4},
}

var sRSResourceResourceTypeAperiodicConstraints = per.SequenceConstraints{
	Extensible: true,
}

var sRSResourceResourceTypeSemiPersistentConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "periodicityAndOffset-sp"},
	},
}

var sRSResourceResourceTypePeriodicConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "periodicityAndOffset-p"},
	},
}

const (
	SRS_Resource_Ext_ResourceMapping_r16_NrofSymbols_r16_N1 = 0
	SRS_Resource_Ext_ResourceMapping_r16_NrofSymbols_r16_N2 = 1
	SRS_Resource_Ext_ResourceMapping_r16_NrofSymbols_r16_N4 = 2
)

var sRSResourceExtResourceMappingR16NrofSymbolsR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_Resource_Ext_ResourceMapping_r16_NrofSymbols_r16_N1, SRS_Resource_Ext_ResourceMapping_r16_NrofSymbols_r16_N2, SRS_Resource_Ext_ResourceMapping_r16_NrofSymbols_r16_N4},
}

const (
	SRS_Resource_Ext_ResourceMapping_r16_RepetitionFactor_r16_N1 = 0
	SRS_Resource_Ext_ResourceMapping_r16_RepetitionFactor_r16_N2 = 1
	SRS_Resource_Ext_ResourceMapping_r16_RepetitionFactor_r16_N4 = 2
)

var sRSResourceExtResourceMappingR16RepetitionFactorR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_Resource_Ext_ResourceMapping_r16_RepetitionFactor_r16_N1, SRS_Resource_Ext_ResourceMapping_r16_RepetitionFactor_r16_N2, SRS_Resource_Ext_ResourceMapping_r16_RepetitionFactor_r16_N4},
}

const (
	SRS_Resource_Ext_ResourceMapping_r17_NrofSymbols_r17_N1  = 0
	SRS_Resource_Ext_ResourceMapping_r17_NrofSymbols_r17_N2  = 1
	SRS_Resource_Ext_ResourceMapping_r17_NrofSymbols_r17_N4  = 2
	SRS_Resource_Ext_ResourceMapping_r17_NrofSymbols_r17_N8  = 3
	SRS_Resource_Ext_ResourceMapping_r17_NrofSymbols_r17_N10 = 4
	SRS_Resource_Ext_ResourceMapping_r17_NrofSymbols_r17_N12 = 5
	SRS_Resource_Ext_ResourceMapping_r17_NrofSymbols_r17_N14 = 6
)

var sRSResourceExtResourceMappingR17NrofSymbolsR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_Resource_Ext_ResourceMapping_r17_NrofSymbols_r17_N1, SRS_Resource_Ext_ResourceMapping_r17_NrofSymbols_r17_N2, SRS_Resource_Ext_ResourceMapping_r17_NrofSymbols_r17_N4, SRS_Resource_Ext_ResourceMapping_r17_NrofSymbols_r17_N8, SRS_Resource_Ext_ResourceMapping_r17_NrofSymbols_r17_N10, SRS_Resource_Ext_ResourceMapping_r17_NrofSymbols_r17_N12, SRS_Resource_Ext_ResourceMapping_r17_NrofSymbols_r17_N14},
}

const (
	SRS_Resource_Ext_ResourceMapping_r17_RepetitionFactor_r17_N1  = 0
	SRS_Resource_Ext_ResourceMapping_r17_RepetitionFactor_r17_N2  = 1
	SRS_Resource_Ext_ResourceMapping_r17_RepetitionFactor_r17_N4  = 2
	SRS_Resource_Ext_ResourceMapping_r17_RepetitionFactor_r17_N5  = 3
	SRS_Resource_Ext_ResourceMapping_r17_RepetitionFactor_r17_N6  = 4
	SRS_Resource_Ext_ResourceMapping_r17_RepetitionFactor_r17_N7  = 5
	SRS_Resource_Ext_ResourceMapping_r17_RepetitionFactor_r17_N8  = 6
	SRS_Resource_Ext_ResourceMapping_r17_RepetitionFactor_r17_N10 = 7
	SRS_Resource_Ext_ResourceMapping_r17_RepetitionFactor_r17_N12 = 8
	SRS_Resource_Ext_ResourceMapping_r17_RepetitionFactor_r17_N14 = 9
)

var sRSResourceExtResourceMappingR17RepetitionFactorR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_Resource_Ext_ResourceMapping_r17_RepetitionFactor_r17_N1, SRS_Resource_Ext_ResourceMapping_r17_RepetitionFactor_r17_N2, SRS_Resource_Ext_ResourceMapping_r17_RepetitionFactor_r17_N4, SRS_Resource_Ext_ResourceMapping_r17_RepetitionFactor_r17_N5, SRS_Resource_Ext_ResourceMapping_r17_RepetitionFactor_r17_N6, SRS_Resource_Ext_ResourceMapping_r17_RepetitionFactor_r17_N7, SRS_Resource_Ext_ResourceMapping_r17_RepetitionFactor_r17_N8, SRS_Resource_Ext_ResourceMapping_r17_RepetitionFactor_r17_N10, SRS_Resource_Ext_ResourceMapping_r17_RepetitionFactor_r17_N12, SRS_Resource_Ext_ResourceMapping_r17_RepetitionFactor_r17_N14},
}

var sRSResourceExtPartialFreqSoundingR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "startRBIndexFScaling-r17"},
		{Name: "enableStartRBHopping-r17", Optional: true},
	},
}

var sRSResourceExtPartialFreqSoundingR17StartRBIndexFScalingR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "startRBIndexAndFreqScalingFactor2-r17"},
		{Name: "startRBIndexAndFreqScalingFactor4-r17"},
	},
}

const (
	SRS_Resource_Ext_PartialFreqSounding_r17_StartRBIndexFScaling_r17_StartRBIndexAndFreqScalingFactor2_r17 = 0
	SRS_Resource_Ext_PartialFreqSounding_r17_StartRBIndexFScaling_r17_StartRBIndexAndFreqScalingFactor4_r17 = 1
)

const (
	SRS_Resource_Ext_PartialFreqSounding_r17_EnableStartRBHopping_r17_Enable = 0
)

var sRSResourceExtPartialFreqSoundingR17EnableStartRBHoppingR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_Resource_Ext_PartialFreqSounding_r17_EnableStartRBHopping_r17_Enable},
}

var sRSResourceExtCombOffsetHoppingR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "hoppingId-r18", Optional: true},
		{Name: "hoppingSubset-r18", Optional: true},
		{Name: "hoppingWithRepetition-r18", Optional: true},
	},
}

var sRSResourceExtCombOffsetHoppingR18HoppingSubsetR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "transmissionComb-n4"},
		{Name: "transmissionComb-n8"},
	},
}

const (
	SRS_Resource_Ext_CombOffsetHopping_r18_HoppingSubset_r18_TransmissionComb_N4 = 0
	SRS_Resource_Ext_CombOffsetHopping_r18_HoppingSubset_r18_TransmissionComb_N8 = 1
)

const (
	SRS_Resource_Ext_CombOffsetHopping_r18_HoppingWithRepetition_r18_Symbol     = 0
	SRS_Resource_Ext_CombOffsetHopping_r18_HoppingWithRepetition_r18_Repetition = 1
)

var sRSResourceExtCombOffsetHoppingR18HoppingWithRepetitionR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_Resource_Ext_CombOffsetHopping_r18_HoppingWithRepetition_r18_Symbol, SRS_Resource_Ext_CombOffsetHopping_r18_HoppingWithRepetition_r18_Repetition},
}

var sRSResourceExtCyclicShiftHoppingR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "hoppingId-r18", Optional: true},
		{Name: "hoppingSubset-r18", Optional: true},
		{Name: "hoppingFinerGranularity-r18", Optional: true},
	},
}

var sRSResourceExtCyclicShiftHoppingR18HoppingSubsetR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "transmissionComb-n2"},
		{Name: "transmissionComb-n4"},
		{Name: "transmissionComb-n8"},
	},
}

const (
	SRS_Resource_Ext_CyclicShiftHopping_r18_HoppingSubset_r18_TransmissionComb_N2 = 0
	SRS_Resource_Ext_CyclicShiftHopping_r18_HoppingSubset_r18_TransmissionComb_N4 = 1
	SRS_Resource_Ext_CyclicShiftHopping_r18_HoppingSubset_r18_TransmissionComb_N8 = 2
)

const (
	SRS_Resource_Ext_CyclicShiftHopping_r18_HoppingFinerGranularity_r18_Enable = 0
)

var sRSResourceExtCyclicShiftHoppingR18HoppingFinerGranularityR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_Resource_Ext_CyclicShiftHopping_r18_HoppingFinerGranularity_r18_Enable},
}

type SRS_Resource struct {
	Srs_ResourceId   SRS_ResourceId
	NrofSRS_Ports    int64
	Ptrs_PortIndex   *int64
	TransmissionComb struct {
		Choice int
		N2     *struct {
			CombOffset_N2  int64
			CyclicShift_N2 int64
		}
		N4 *struct {
			CombOffset_N4  int64
			CyclicShift_N4 int64
		}
	}
	ResourceMapping struct {
		StartPosition    int64
		NrofSymbols      int64
		RepetitionFactor int64
	}
	FreqDomainPosition int64
	FreqDomainShift    int64
	FreqHopping        struct {
		C_SRS int64
		B_SRS int64
		B_Hop int64
	}
	GroupOrSequenceHopping int64
	ResourceType           struct {
		Choice          int
		Aperiodic       *struct{}
		Semi_Persistent *struct{ PeriodicityAndOffset_Sp SRS_PeriodicityAndOffset }
		Periodic        *struct{ PeriodicityAndOffset_P SRS_PeriodicityAndOffset }
	}
	SequenceId          int64
	SpatialRelationInfo *SRS_SpatialRelationInfo
	ResourceMapping_r16 *struct {
		StartPosition_r16    int64
		NrofSymbols_r16      int64
		RepetitionFactor_r16 int64
	}
	SpatialRelationInfo_PDC_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *SpatialRelationInfo_PDC_r17
	}
	ResourceMapping_r17 *struct {
		StartPosition_r17    int64
		NrofSymbols_r17      int64
		RepetitionFactor_r17 int64
	}
	PartialFreqSounding_r17 *struct {
		StartRBIndexFScaling_r17 struct {
			Choice                                int
			StartRBIndexAndFreqScalingFactor2_r17 *int64
			StartRBIndexAndFreqScalingFactor4_r17 *int64
		}
		EnableStartRBHopping_r17 *int64
	}
	TransmissionComb_N8_r17 *struct {
		CombOffset_N8_r17  int64
		CyclicShift_N8_r17 int64
	}
	Srs_TCI_State_r17 *struct {
		Choice                 int
		Srs_UL_TCI_State       *TCI_UL_StateId_r17
		Srs_DLorJointTCI_State *TCI_StateId
	}
	RepetitionFactor_v1730       *int64
	Srs_DLorJointTCI_State_v1730 *struct{ CellAndBWP_r17 ServingCellAndBWP_Id_r17 }
	NrofSRS_Ports_N8_r18         *int64
	CombOffsetHopping_r18        *struct {
		HoppingId_r18     *int64
		HoppingSubset_r18 *struct {
			Choice              int
			TransmissionComb_N4 *per.BitString
			TransmissionComb_N8 *per.BitString
		}
		HoppingWithRepetition_r18 *int64
	}
	CyclicShiftHopping_r18 *struct {
		HoppingId_r18     *int64
		HoppingSubset_r18 *struct {
			Choice              int
			TransmissionComb_N2 *per.BitString
			TransmissionComb_N4 *per.BitString
			TransmissionComb_N8 *per.BitString
		}
		HoppingFinerGranularity_r18 *int64
	}
}

func (ie *SRS_Resource) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sRSResourceConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.ResourceMapping_r16 != nil
	hasExtGroup1 := ie.SpatialRelationInfo_PDC_r17 != nil || ie.ResourceMapping_r17 != nil || ie.PartialFreqSounding_r17 != nil || ie.TransmissionComb_N8_r17 != nil || ie.Srs_TCI_State_r17 != nil
	hasExtGroup2 := ie.RepetitionFactor_v1730 != nil || ie.Srs_DLorJointTCI_State_v1730 != nil
	hasExtGroup3 := ie.NrofSRS_Ports_N8_r18 != nil || ie.CombOffsetHopping_r18 != nil || ie.CyclicShiftHopping_r18 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ptrs_PortIndex != nil, ie.SpatialRelationInfo != nil}); err != nil {
		return err
	}

	// 3. srs-ResourceId: ref
	{
		if err := ie.Srs_ResourceId.Encode(e); err != nil {
			return err
		}
	}

	// 4. nrofSRS-Ports: enumerated
	{
		if err := e.EncodeEnumerated(ie.NrofSRS_Ports, sRSResourceNrofSRSPortsConstraints); err != nil {
			return err
		}
	}

	// 5. ptrs-PortIndex: enumerated
	{
		if ie.Ptrs_PortIndex != nil {
			if err := e.EncodeEnumerated(*ie.Ptrs_PortIndex, sRSResourcePtrsPortIndexConstraints); err != nil {
				return err
			}
		}
	}

	// 6. transmissionComb: choice
	{
		choiceEnc := e.NewChoiceEncoder(sRS_ResourceTransmissionCombConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.TransmissionComb.Choice), false, nil); err != nil {
			return err
		}
		switch ie.TransmissionComb.Choice {
		case SRS_Resource_TransmissionComb_N2:
			c := (*ie.TransmissionComb.N2)
			if err := e.EncodeInteger(c.CombOffset_N2, per.Constrained(0, 1)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.CyclicShift_N2, per.Constrained(0, 7)); err != nil {
				return err
			}
		case SRS_Resource_TransmissionComb_N4:
			c := (*ie.TransmissionComb.N4)
			if err := e.EncodeInteger(c.CombOffset_N4, per.Constrained(0, 3)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.CyclicShift_N4, per.Constrained(0, 11)); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.TransmissionComb.Choice), Constraint: "index out of range"}
		}
	}

	// 7. resourceMapping: sequence
	{
		{
			c := &ie.ResourceMapping
			if err := e.EncodeInteger(c.StartPosition, per.Constrained(0, 5)); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.NrofSymbols, sRSResourceResourceMappingNrofSymbolsConstraints); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.RepetitionFactor, sRSResourceResourceMappingRepetitionFactorConstraints); err != nil {
				return err
			}
		}
	}

	// 8. freqDomainPosition: integer
	{
		if err := e.EncodeInteger(ie.FreqDomainPosition, sRSResourceFreqDomainPositionConstraints); err != nil {
			return err
		}
	}

	// 9. freqDomainShift: integer
	{
		if err := e.EncodeInteger(ie.FreqDomainShift, sRSResourceFreqDomainShiftConstraints); err != nil {
			return err
		}
	}

	// 10. freqHopping: sequence
	{
		{
			c := &ie.FreqHopping
			if err := e.EncodeInteger(c.C_SRS, per.Constrained(0, 63)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.B_SRS, per.Constrained(0, 3)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.B_Hop, per.Constrained(0, 3)); err != nil {
				return err
			}
		}
	}

	// 11. groupOrSequenceHopping: enumerated
	{
		if err := e.EncodeEnumerated(ie.GroupOrSequenceHopping, sRSResourceGroupOrSequenceHoppingConstraints); err != nil {
			return err
		}
	}

	// 12. resourceType: choice
	{
		choiceEnc := e.NewChoiceEncoder(sRS_ResourceResourceTypeConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.ResourceType.Choice), false, nil); err != nil {
			return err
		}
		switch ie.ResourceType.Choice {
		case SRS_Resource_ResourceType_Aperiodic:
			sRSResourceResourceTypeAperiodicSeq := e.NewSequenceEncoder(sRSResourceResourceTypeAperiodicConstraints)
			if err := sRSResourceResourceTypeAperiodicSeq.EncodeExtensionBit(false); err != nil {
				return err
			}
		case SRS_Resource_ResourceType_Semi_Persistent:
			c := (*ie.ResourceType.Semi_Persistent)
			sRSResourceResourceTypeSemiPersistentSeq := e.NewSequenceEncoder(sRSResourceResourceTypeSemiPersistentConstraints)
			if err := sRSResourceResourceTypeSemiPersistentSeq.EncodeExtensionBit(false); err != nil {
				return err
			}
			if err := c.PeriodicityAndOffset_Sp.Encode(e); err != nil {
				return err
			}
		case SRS_Resource_ResourceType_Periodic:
			c := (*ie.ResourceType.Periodic)
			sRSResourceResourceTypePeriodicSeq := e.NewSequenceEncoder(sRSResourceResourceTypePeriodicConstraints)
			if err := sRSResourceResourceTypePeriodicSeq.EncodeExtensionBit(false); err != nil {
				return err
			}
			if err := c.PeriodicityAndOffset_P.Encode(e); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.ResourceType.Choice), Constraint: "index out of range"}
		}
	}

	// 13. sequenceId: integer
	{
		if err := e.EncodeInteger(ie.SequenceId, sRSResourceSequenceIdConstraints); err != nil {
			return err
		}
	}

	// 14. spatialRelationInfo: ref
	{
		if ie.SpatialRelationInfo != nil {
			if err := ie.SpatialRelationInfo.Encode(e); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2, hasExtGroup3}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "resourceMapping-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.ResourceMapping_r16 != nil}); err != nil {
				return err
			}

			if ie.ResourceMapping_r16 != nil {
				c := ie.ResourceMapping_r16
				if err := ex.EncodeInteger(c.StartPosition_r16, per.Constrained(0, 13)); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.NrofSymbols_r16, sRSResourceExtResourceMappingR16NrofSymbolsR16Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.RepetitionFactor_r16, sRSResourceExtResourceMappingR16RepetitionFactorR16Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 1:
		if hasExtGroup1 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "spatialRelationInfo-PDC-r17", Optional: true},
					{Name: "resourceMapping-r17", Optional: true},
					{Name: "partialFreqSounding-r17", Optional: true},
					{Name: "transmissionComb-n8-r17", Optional: true},
					{Name: "srs-TCI-State-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SpatialRelationInfo_PDC_r17 != nil, ie.ResourceMapping_r17 != nil, ie.PartialFreqSounding_r17 != nil, ie.TransmissionComb_N8_r17 != nil, ie.Srs_TCI_State_r17 != nil}); err != nil {
				return err
			}

			if ie.SpatialRelationInfo_PDC_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(sRSResourceExtSpatialRelationInfoPDCR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.SpatialRelationInfo_PDC_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.SpatialRelationInfo_PDC_r17).Choice {
				case SRS_Resource_Ext_SpatialRelationInfo_PDC_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case SRS_Resource_Ext_SpatialRelationInfo_PDC_r17_Setup:
					if err := (*ie.SpatialRelationInfo_PDC_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.ResourceMapping_r17 != nil {
				c := ie.ResourceMapping_r17
				if err := ex.EncodeInteger(c.StartPosition_r17, per.Constrained(0, 13)); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.NrofSymbols_r17, sRSResourceExtResourceMappingR17NrofSymbolsR17Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.RepetitionFactor_r17, sRSResourceExtResourceMappingR17RepetitionFactorR17Constraints); err != nil {
					return err
				}
			}

			if ie.PartialFreqSounding_r17 != nil {
				c := ie.PartialFreqSounding_r17
				sRSResourceExtPartialFreqSoundingR17Seq := ex.NewSequenceEncoder(sRSResourceExtPartialFreqSoundingR17Constraints)
				if err := sRSResourceExtPartialFreqSoundingR17Seq.EncodePreamble([]bool{c.EnableStartRBHopping_r17 != nil}); err != nil {
					return err
				}
				{
					choiceEnc := ex.NewChoiceEncoder(sRSResourceExtPartialFreqSoundingR17StartRBIndexFScalingR17Constraints)
					if err := choiceEnc.EncodeChoice(int64(c.StartRBIndexFScaling_r17.Choice), false, nil); err != nil {
						return err
					}
					switch c.StartRBIndexFScaling_r17.Choice {
					case SRS_Resource_Ext_PartialFreqSounding_r17_StartRBIndexFScaling_r17_StartRBIndexAndFreqScalingFactor2_r17:
						if err := ex.EncodeInteger((*c.StartRBIndexFScaling_r17.StartRBIndexAndFreqScalingFactor2_r17), per.Constrained(0, 1)); err != nil {
							return err
						}
					case SRS_Resource_Ext_PartialFreqSounding_r17_StartRBIndexFScaling_r17_StartRBIndexAndFreqScalingFactor4_r17:
						if err := ex.EncodeInteger((*c.StartRBIndexFScaling_r17.StartRBIndexAndFreqScalingFactor4_r17), per.Constrained(0, 3)); err != nil {
							return err
						}
					}
				}
				if c.EnableStartRBHopping_r17 != nil {
					if err := ex.EncodeEnumerated((*c.EnableStartRBHopping_r17), sRSResourceExtPartialFreqSoundingR17EnableStartRBHoppingR17Constraints); err != nil {
						return err
					}
				}
			}

			if ie.TransmissionComb_N8_r17 != nil {
				c := ie.TransmissionComb_N8_r17
				if err := ex.EncodeInteger(c.CombOffset_N8_r17, per.Constrained(0, 7)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.CyclicShift_N8_r17, per.Constrained(0, 5)); err != nil {
					return err
				}
			}

			if ie.Srs_TCI_State_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(sRSResourceExtSrsTCIStateR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Srs_TCI_State_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Srs_TCI_State_r17).Choice {
				case SRS_Resource_Ext_Srs_TCI_State_r17_Srs_UL_TCI_State:
					if err := (*ie.Srs_TCI_State_r17).Srs_UL_TCI_State.Encode(ex); err != nil {
						return err
					}
				case SRS_Resource_Ext_Srs_TCI_State_r17_Srs_DLorJointTCI_State:
					if err := (*ie.Srs_TCI_State_r17).Srs_DLorJointTCI_State.Encode(ex); err != nil {
						return err
					}
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 2:
		if hasExtGroup2 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "repetitionFactor-v1730", Optional: true},
					{Name: "srs-DLorJointTCI-State-v1730", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.RepetitionFactor_v1730 != nil, ie.Srs_DLorJointTCI_State_v1730 != nil}); err != nil {
				return err
			}

			if ie.RepetitionFactor_v1730 != nil {
				if err := ex.EncodeEnumerated(*ie.RepetitionFactor_v1730, sRSResourceExtRepetitionFactorV1730Constraints); err != nil {
					return err
				}
			}

			if ie.Srs_DLorJointTCI_State_v1730 != nil {
				c := ie.Srs_DLorJointTCI_State_v1730
				if err := c.CellAndBWP_r17.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 3:
		if hasExtGroup3 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "nrofSRS-Ports-n8-r18", Optional: true},
					{Name: "combOffsetHopping-r18", Optional: true},
					{Name: "cyclicShiftHopping-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.NrofSRS_Ports_N8_r18 != nil, ie.CombOffsetHopping_r18 != nil, ie.CyclicShiftHopping_r18 != nil}); err != nil {
				return err
			}

			if ie.NrofSRS_Ports_N8_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.NrofSRS_Ports_N8_r18, sRSResourceExtNrofSRSPortsN8R18Constraints); err != nil {
					return err
				}
			}

			if ie.CombOffsetHopping_r18 != nil {
				c := ie.CombOffsetHopping_r18
				sRSResourceExtCombOffsetHoppingR18Seq := ex.NewSequenceEncoder(sRSResourceExtCombOffsetHoppingR18Constraints)
				if err := sRSResourceExtCombOffsetHoppingR18Seq.EncodePreamble([]bool{c.HoppingId_r18 != nil, c.HoppingSubset_r18 != nil, c.HoppingWithRepetition_r18 != nil}); err != nil {
					return err
				}
				if c.HoppingId_r18 != nil {
					if err := ex.EncodeInteger((*c.HoppingId_r18), per.Constrained(0, 1023)); err != nil {
						return err
					}
				}
				if c.HoppingSubset_r18 != nil {
					choiceEnc := ex.NewChoiceEncoder(sRSResourceExtCombOffsetHoppingR18HoppingSubsetR18Constraints)
					if err := choiceEnc.EncodeChoice(int64((*c.HoppingSubset_r18).Choice), false, nil); err != nil {
						return err
					}
					switch (*c.HoppingSubset_r18).Choice {
					case SRS_Resource_Ext_CombOffsetHopping_r18_HoppingSubset_r18_TransmissionComb_N4:
						if err := ex.EncodeBitString((*(*c.HoppingSubset_r18).TransmissionComb_N4), per.FixedSize(4)); err != nil {
							return err
						}
					case SRS_Resource_Ext_CombOffsetHopping_r18_HoppingSubset_r18_TransmissionComb_N8:
						if err := ex.EncodeBitString((*(*c.HoppingSubset_r18).TransmissionComb_N8), per.FixedSize(8)); err != nil {
							return err
						}
					}
				}
				if c.HoppingWithRepetition_r18 != nil {
					if err := ex.EncodeEnumerated((*c.HoppingWithRepetition_r18), sRSResourceExtCombOffsetHoppingR18HoppingWithRepetitionR18Constraints); err != nil {
						return err
					}
				}
			}

			if ie.CyclicShiftHopping_r18 != nil {
				c := ie.CyclicShiftHopping_r18
				sRSResourceExtCyclicShiftHoppingR18Seq := ex.NewSequenceEncoder(sRSResourceExtCyclicShiftHoppingR18Constraints)
				if err := sRSResourceExtCyclicShiftHoppingR18Seq.EncodePreamble([]bool{c.HoppingId_r18 != nil, c.HoppingSubset_r18 != nil, c.HoppingFinerGranularity_r18 != nil}); err != nil {
					return err
				}
				if c.HoppingId_r18 != nil {
					if err := ex.EncodeInteger((*c.HoppingId_r18), per.Constrained(0, 1023)); err != nil {
						return err
					}
				}
				if c.HoppingSubset_r18 != nil {
					choiceEnc := ex.NewChoiceEncoder(sRSResourceExtCyclicShiftHoppingR18HoppingSubsetR18Constraints)
					if err := choiceEnc.EncodeChoice(int64((*c.HoppingSubset_r18).Choice), false, nil); err != nil {
						return err
					}
					switch (*c.HoppingSubset_r18).Choice {
					case SRS_Resource_Ext_CyclicShiftHopping_r18_HoppingSubset_r18_TransmissionComb_N2:
						if err := ex.EncodeBitString((*(*c.HoppingSubset_r18).TransmissionComb_N2), per.FixedSize(8)); err != nil {
							return err
						}
					case SRS_Resource_Ext_CyclicShiftHopping_r18_HoppingSubset_r18_TransmissionComb_N4:
						if err := ex.EncodeBitString((*(*c.HoppingSubset_r18).TransmissionComb_N4), per.FixedSize(12)); err != nil {
							return err
						}
					case SRS_Resource_Ext_CyclicShiftHopping_r18_HoppingSubset_r18_TransmissionComb_N8:
						if err := ex.EncodeBitString((*(*c.HoppingSubset_r18).TransmissionComb_N8), per.FixedSize(6)); err != nil {
							return err
						}
					}
				}
				if c.HoppingFinerGranularity_r18 != nil {
					if err := ex.EncodeEnumerated((*c.HoppingFinerGranularity_r18), sRSResourceExtCyclicShiftHoppingR18HoppingFinerGranularityR18Constraints); err != nil {
						return err
					}
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

func (ie *SRS_Resource) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sRSResourceConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. srs-ResourceId: ref
	{
		if err := ie.Srs_ResourceId.Decode(d); err != nil {
			return err
		}
	}

	// 4. nrofSRS-Ports: enumerated
	{
		v1, err := d.DecodeEnumerated(sRSResourceNrofSRSPortsConstraints)
		if err != nil {
			return err
		}
		ie.NrofSRS_Ports = v1
	}

	// 5. ptrs-PortIndex: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(sRSResourcePtrsPortIndexConstraints)
			if err != nil {
				return err
			}
			ie.Ptrs_PortIndex = &idx
		}
	}

	// 6. transmissionComb: choice
	{
		choiceDec := d.NewChoiceDecoder(sRS_ResourceTransmissionCombConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.TransmissionComb.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case SRS_Resource_TransmissionComb_N2:
			ie.TransmissionComb.N2 = &struct {
				CombOffset_N2  int64
				CyclicShift_N2 int64
			}{}
			c := (*ie.TransmissionComb.N2)
			{
				v, err := d.DecodeInteger(per.Constrained(0, 1))
				if err != nil {
					return err
				}
				c.CombOffset_N2 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(0, 7))
				if err != nil {
					return err
				}
				c.CyclicShift_N2 = v
			}
		case SRS_Resource_TransmissionComb_N4:
			ie.TransmissionComb.N4 = &struct {
				CombOffset_N4  int64
				CyclicShift_N4 int64
			}{}
			c := (*ie.TransmissionComb.N4)
			{
				v, err := d.DecodeInteger(per.Constrained(0, 3))
				if err != nil {
					return err
				}
				c.CombOffset_N4 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(0, 11))
				if err != nil {
					return err
				}
				c.CyclicShift_N4 = v
			}
		}
	}

	// 7. resourceMapping: sequence
	{
		{
			c := &ie.ResourceMapping
			{
				v, err := d.DecodeInteger(per.Constrained(0, 5))
				if err != nil {
					return err
				}
				c.StartPosition = v
			}
			{
				v, err := d.DecodeEnumerated(sRSResourceResourceMappingNrofSymbolsConstraints)
				if err != nil {
					return err
				}
				c.NrofSymbols = v
			}
			{
				v, err := d.DecodeEnumerated(sRSResourceResourceMappingRepetitionFactorConstraints)
				if err != nil {
					return err
				}
				c.RepetitionFactor = v
			}
		}
	}

	// 8. freqDomainPosition: integer
	{
		v5, err := d.DecodeInteger(sRSResourceFreqDomainPositionConstraints)
		if err != nil {
			return err
		}
		ie.FreqDomainPosition = v5
	}

	// 9. freqDomainShift: integer
	{
		v6, err := d.DecodeInteger(sRSResourceFreqDomainShiftConstraints)
		if err != nil {
			return err
		}
		ie.FreqDomainShift = v6
	}

	// 10. freqHopping: sequence
	{
		{
			c := &ie.FreqHopping
			{
				v, err := d.DecodeInteger(per.Constrained(0, 63))
				if err != nil {
					return err
				}
				c.C_SRS = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(0, 3))
				if err != nil {
					return err
				}
				c.B_SRS = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(0, 3))
				if err != nil {
					return err
				}
				c.B_Hop = v
			}
		}
	}

	// 11. groupOrSequenceHopping: enumerated
	{
		v8, err := d.DecodeEnumerated(sRSResourceGroupOrSequenceHoppingConstraints)
		if err != nil {
			return err
		}
		ie.GroupOrSequenceHopping = v8
	}

	// 12. resourceType: choice
	{
		choiceDec := d.NewChoiceDecoder(sRS_ResourceResourceTypeConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.ResourceType.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case SRS_Resource_ResourceType_Aperiodic:
			ie.ResourceType.Aperiodic = &struct{}{}
			sRSResourceResourceTypeAperiodicSeq := d.NewSequenceDecoder(sRSResourceResourceTypeAperiodicConstraints)
			if err := sRSResourceResourceTypeAperiodicSeq.DecodeExtensionBit(); err != nil {
				return err
			}
		case SRS_Resource_ResourceType_Semi_Persistent:
			ie.ResourceType.Semi_Persistent = &struct{ PeriodicityAndOffset_Sp SRS_PeriodicityAndOffset }{}
			c := (*ie.ResourceType.Semi_Persistent)
			sRSResourceResourceTypeSemiPersistentSeq := d.NewSequenceDecoder(sRSResourceResourceTypeSemiPersistentConstraints)
			if err := sRSResourceResourceTypeSemiPersistentSeq.DecodeExtensionBit(); err != nil {
				return err
			}
			{
				if err := c.PeriodicityAndOffset_Sp.Decode(d); err != nil {
					return err
				}
			}
		case SRS_Resource_ResourceType_Periodic:
			ie.ResourceType.Periodic = &struct{ PeriodicityAndOffset_P SRS_PeriodicityAndOffset }{}
			c := (*ie.ResourceType.Periodic)
			sRSResourceResourceTypePeriodicSeq := d.NewSequenceDecoder(sRSResourceResourceTypePeriodicConstraints)
			if err := sRSResourceResourceTypePeriodicSeq.DecodeExtensionBit(); err != nil {
				return err
			}
			{
				if err := c.PeriodicityAndOffset_P.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 13. sequenceId: integer
	{
		v10, err := d.DecodeInteger(sRSResourceSequenceIdConstraints)
		if err != nil {
			return err
		}
		ie.SequenceId = v10
	}

	// 14. spatialRelationInfo: ref
	{
		if seq.IsComponentPresent(11) {
			ie.SpatialRelationInfo = new(SRS_SpatialRelationInfo)
			if err := ie.SpatialRelationInfo.Decode(d); err != nil {
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
				{Name: "resourceMapping-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.ResourceMapping_r16 = &struct {
				StartPosition_r16    int64
				NrofSymbols_r16      int64
				RepetitionFactor_r16 int64
			}{}
			c := ie.ResourceMapping_r16
			{
				v, err := dx.DecodeInteger(per.Constrained(0, 13))
				if err != nil {
					return err
				}
				c.StartPosition_r16 = v
			}
			{
				v, err := dx.DecodeEnumerated(sRSResourceExtResourceMappingR16NrofSymbolsR16Constraints)
				if err != nil {
					return err
				}
				c.NrofSymbols_r16 = v
			}
			{
				v, err := dx.DecodeEnumerated(sRSResourceExtResourceMappingR16RepetitionFactorR16Constraints)
				if err != nil {
					return err
				}
				c.RepetitionFactor_r16 = v
			}
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "spatialRelationInfo-PDC-r17", Optional: true},
				{Name: "resourceMapping-r17", Optional: true},
				{Name: "partialFreqSounding-r17", Optional: true},
				{Name: "transmissionComb-n8-r17", Optional: true},
				{Name: "srs-TCI-State-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.SpatialRelationInfo_PDC_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SpatialRelationInfo_PDC_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(sRSResourceExtSpatialRelationInfoPDCR17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.SpatialRelationInfo_PDC_r17).Choice = int(index)
			switch index {
			case SRS_Resource_Ext_SpatialRelationInfo_PDC_r17_Release:
				(*ie.SpatialRelationInfo_PDC_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case SRS_Resource_Ext_SpatialRelationInfo_PDC_r17_Setup:
				(*ie.SpatialRelationInfo_PDC_r17).Setup = new(SpatialRelationInfo_PDC_r17)
				if err := (*ie.SpatialRelationInfo_PDC_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.ResourceMapping_r17 = &struct {
				StartPosition_r17    int64
				NrofSymbols_r17      int64
				RepetitionFactor_r17 int64
			}{}
			c := ie.ResourceMapping_r17
			{
				v, err := dx.DecodeInteger(per.Constrained(0, 13))
				if err != nil {
					return err
				}
				c.StartPosition_r17 = v
			}
			{
				v, err := dx.DecodeEnumerated(sRSResourceExtResourceMappingR17NrofSymbolsR17Constraints)
				if err != nil {
					return err
				}
				c.NrofSymbols_r17 = v
			}
			{
				v, err := dx.DecodeEnumerated(sRSResourceExtResourceMappingR17RepetitionFactorR17Constraints)
				if err != nil {
					return err
				}
				c.RepetitionFactor_r17 = v
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.PartialFreqSounding_r17 = &struct {
				StartRBIndexFScaling_r17 struct {
					Choice                                int
					StartRBIndexAndFreqScalingFactor2_r17 *int64
					StartRBIndexAndFreqScalingFactor4_r17 *int64
				}
				EnableStartRBHopping_r17 *int64
			}{}
			c := ie.PartialFreqSounding_r17
			sRSResourceExtPartialFreqSoundingR17Seq := dx.NewSequenceDecoder(sRSResourceExtPartialFreqSoundingR17Constraints)
			if err := sRSResourceExtPartialFreqSoundingR17Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				choiceDec := dx.NewChoiceDecoder(sRSResourceExtPartialFreqSoundingR17StartRBIndexFScalingR17Constraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				c.StartRBIndexFScaling_r17.Choice = int(index)
				switch index {
				case SRS_Resource_Ext_PartialFreqSounding_r17_StartRBIndexFScaling_r17_StartRBIndexAndFreqScalingFactor2_r17:
					c.StartRBIndexFScaling_r17.StartRBIndexAndFreqScalingFactor2_r17 = new(int64)
					v, err := dx.DecodeInteger(per.Constrained(0, 1))
					if err != nil {
						return err
					}
					(*c.StartRBIndexFScaling_r17.StartRBIndexAndFreqScalingFactor2_r17) = v
				case SRS_Resource_Ext_PartialFreqSounding_r17_StartRBIndexFScaling_r17_StartRBIndexAndFreqScalingFactor4_r17:
					c.StartRBIndexFScaling_r17.StartRBIndexAndFreqScalingFactor4_r17 = new(int64)
					v, err := dx.DecodeInteger(per.Constrained(0, 3))
					if err != nil {
						return err
					}
					(*c.StartRBIndexFScaling_r17.StartRBIndexAndFreqScalingFactor4_r17) = v
				}
			}
			if sRSResourceExtPartialFreqSoundingR17Seq.IsComponentPresent(1) {
				c.EnableStartRBHopping_r17 = new(int64)
				v, err := dx.DecodeEnumerated(sRSResourceExtPartialFreqSoundingR17EnableStartRBHoppingR17Constraints)
				if err != nil {
					return err
				}
				(*c.EnableStartRBHopping_r17) = v
			}
		}

		if groupSeq.IsComponentPresent(3) {
			ie.TransmissionComb_N8_r17 = &struct {
				CombOffset_N8_r17  int64
				CyclicShift_N8_r17 int64
			}{}
			c := ie.TransmissionComb_N8_r17
			{
				v, err := dx.DecodeInteger(per.Constrained(0, 7))
				if err != nil {
					return err
				}
				c.CombOffset_N8_r17 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(0, 5))
				if err != nil {
					return err
				}
				c.CyclicShift_N8_r17 = v
			}
		}

		if groupSeq.IsComponentPresent(4) {
			ie.Srs_TCI_State_r17 = &struct {
				Choice                 int
				Srs_UL_TCI_State       *TCI_UL_StateId_r17
				Srs_DLorJointTCI_State *TCI_StateId
			}{}
			choiceDec := dx.NewChoiceDecoder(sRSResourceExtSrsTCIStateR17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Srs_TCI_State_r17).Choice = int(index)
			switch index {
			case SRS_Resource_Ext_Srs_TCI_State_r17_Srs_UL_TCI_State:
				(*ie.Srs_TCI_State_r17).Srs_UL_TCI_State = new(TCI_UL_StateId_r17)
				if err := (*ie.Srs_TCI_State_r17).Srs_UL_TCI_State.Decode(dx); err != nil {
					return err
				}
			case SRS_Resource_Ext_Srs_TCI_State_r17_Srs_DLorJointTCI_State:
				(*ie.Srs_TCI_State_r17).Srs_DLorJointTCI_State = new(TCI_StateId)
				if err := (*ie.Srs_TCI_State_r17).Srs_DLorJointTCI_State.Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "repetitionFactor-v1730", Optional: true},
				{Name: "srs-DLorJointTCI-State-v1730", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(sRSResourceExtRepetitionFactorV1730Constraints)
			if err != nil {
				return err
			}
			ie.RepetitionFactor_v1730 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Srs_DLorJointTCI_State_v1730 = &struct{ CellAndBWP_r17 ServingCellAndBWP_Id_r17 }{}
			c := ie.Srs_DLorJointTCI_State_v1730
			{
				if err := c.CellAndBWP_r17.Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	// Extension group 3:
	if seq.IsExtensionPresent(3) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "nrofSRS-Ports-n8-r18", Optional: true},
				{Name: "combOffsetHopping-r18", Optional: true},
				{Name: "cyclicShiftHopping-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(sRSResourceExtNrofSRSPortsN8R18Constraints)
			if err != nil {
				return err
			}
			ie.NrofSRS_Ports_N8_r18 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			ie.CombOffsetHopping_r18 = &struct {
				HoppingId_r18     *int64
				HoppingSubset_r18 *struct {
					Choice              int
					TransmissionComb_N4 *per.BitString
					TransmissionComb_N8 *per.BitString
				}
				HoppingWithRepetition_r18 *int64
			}{}
			c := ie.CombOffsetHopping_r18
			sRSResourceExtCombOffsetHoppingR18Seq := dx.NewSequenceDecoder(sRSResourceExtCombOffsetHoppingR18Constraints)
			if err := sRSResourceExtCombOffsetHoppingR18Seq.DecodePreamble(); err != nil {
				return err
			}
			if sRSResourceExtCombOffsetHoppingR18Seq.IsComponentPresent(0) {
				c.HoppingId_r18 = new(int64)
				v, err := dx.DecodeInteger(per.Constrained(0, 1023))
				if err != nil {
					return err
				}
				(*c.HoppingId_r18) = v
			}
			if sRSResourceExtCombOffsetHoppingR18Seq.IsComponentPresent(1) {
				c.HoppingSubset_r18 = &struct {
					Choice              int
					TransmissionComb_N4 *per.BitString
					TransmissionComb_N8 *per.BitString
				}{}
				choiceDec := dx.NewChoiceDecoder(sRSResourceExtCombOffsetHoppingR18HoppingSubsetR18Constraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				(*c.HoppingSubset_r18).Choice = int(index)
				switch index {
				case SRS_Resource_Ext_CombOffsetHopping_r18_HoppingSubset_r18_TransmissionComb_N4:
					(*c.HoppingSubset_r18).TransmissionComb_N4 = new(per.BitString)
					v, err := dx.DecodeBitString(per.FixedSize(4))
					if err != nil {
						return err
					}
					(*(*c.HoppingSubset_r18).TransmissionComb_N4) = v
				case SRS_Resource_Ext_CombOffsetHopping_r18_HoppingSubset_r18_TransmissionComb_N8:
					(*c.HoppingSubset_r18).TransmissionComb_N8 = new(per.BitString)
					v, err := dx.DecodeBitString(per.FixedSize(8))
					if err != nil {
						return err
					}
					(*(*c.HoppingSubset_r18).TransmissionComb_N8) = v
				}
			}
			if sRSResourceExtCombOffsetHoppingR18Seq.IsComponentPresent(2) {
				c.HoppingWithRepetition_r18 = new(int64)
				v, err := dx.DecodeEnumerated(sRSResourceExtCombOffsetHoppingR18HoppingWithRepetitionR18Constraints)
				if err != nil {
					return err
				}
				(*c.HoppingWithRepetition_r18) = v
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.CyclicShiftHopping_r18 = &struct {
				HoppingId_r18     *int64
				HoppingSubset_r18 *struct {
					Choice              int
					TransmissionComb_N2 *per.BitString
					TransmissionComb_N4 *per.BitString
					TransmissionComb_N8 *per.BitString
				}
				HoppingFinerGranularity_r18 *int64
			}{}
			c := ie.CyclicShiftHopping_r18
			sRSResourceExtCyclicShiftHoppingR18Seq := dx.NewSequenceDecoder(sRSResourceExtCyclicShiftHoppingR18Constraints)
			if err := sRSResourceExtCyclicShiftHoppingR18Seq.DecodePreamble(); err != nil {
				return err
			}
			if sRSResourceExtCyclicShiftHoppingR18Seq.IsComponentPresent(0) {
				c.HoppingId_r18 = new(int64)
				v, err := dx.DecodeInteger(per.Constrained(0, 1023))
				if err != nil {
					return err
				}
				(*c.HoppingId_r18) = v
			}
			if sRSResourceExtCyclicShiftHoppingR18Seq.IsComponentPresent(1) {
				c.HoppingSubset_r18 = &struct {
					Choice              int
					TransmissionComb_N2 *per.BitString
					TransmissionComb_N4 *per.BitString
					TransmissionComb_N8 *per.BitString
				}{}
				choiceDec := dx.NewChoiceDecoder(sRSResourceExtCyclicShiftHoppingR18HoppingSubsetR18Constraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				(*c.HoppingSubset_r18).Choice = int(index)
				switch index {
				case SRS_Resource_Ext_CyclicShiftHopping_r18_HoppingSubset_r18_TransmissionComb_N2:
					(*c.HoppingSubset_r18).TransmissionComb_N2 = new(per.BitString)
					v, err := dx.DecodeBitString(per.FixedSize(8))
					if err != nil {
						return err
					}
					(*(*c.HoppingSubset_r18).TransmissionComb_N2) = v
				case SRS_Resource_Ext_CyclicShiftHopping_r18_HoppingSubset_r18_TransmissionComb_N4:
					(*c.HoppingSubset_r18).TransmissionComb_N4 = new(per.BitString)
					v, err := dx.DecodeBitString(per.FixedSize(12))
					if err != nil {
						return err
					}
					(*(*c.HoppingSubset_r18).TransmissionComb_N4) = v
				case SRS_Resource_Ext_CyclicShiftHopping_r18_HoppingSubset_r18_TransmissionComb_N8:
					(*c.HoppingSubset_r18).TransmissionComb_N8 = new(per.BitString)
					v, err := dx.DecodeBitString(per.FixedSize(6))
					if err != nil {
						return err
					}
					(*(*c.HoppingSubset_r18).TransmissionComb_N8) = v
				}
			}
			if sRSResourceExtCyclicShiftHoppingR18Seq.IsComponentPresent(2) {
				c.HoppingFinerGranularity_r18 = new(int64)
				v, err := dx.DecodeEnumerated(sRSResourceExtCyclicShiftHoppingR18HoppingFinerGranularityR18Constraints)
				if err != nil {
					return err
				}
				(*c.HoppingFinerGranularity_r18) = v
			}
		}
	}

	return nil
}
