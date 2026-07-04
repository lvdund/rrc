// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SRS-PosResourceSet-r16 (line 15402).

var sRSPosResourceSetR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "srs-PosResourceSetId-r16"},
		{Name: "srs-PosResourceIdList-r16", Optional: true},
		{Name: "resourceType-r16"},
		{Name: "alpha-r16", Optional: true},
		{Name: "p0-r16", Optional: true},
		{Name: "pathlossReferenceRS-Pos-r16", Optional: true},
	},
}

var sRSPosResourceSetR16SrsPosResourceIdListR16Constraints = per.SizeRange(1, common.MaxNrofSRS_ResourcesPerSet)

var sRS_PosResourceSet_r16ResourceTypeR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "aperiodic-r16"},
		{Name: "semi-persistent-r16"},
		{Name: "periodic-r16"},
	},
}

const (
	SRS_PosResourceSet_r16_ResourceType_r16_Aperiodic_r16       = 0
	SRS_PosResourceSet_r16_ResourceType_r16_Semi_Persistent_r16 = 1
	SRS_PosResourceSet_r16_ResourceType_r16_Periodic_r16        = 2
)

var sRSPosResourceSetR16P0R16Constraints = per.Constrained(-202, 24)

var sRS_PosResourceSet_r16PathlossReferenceRSPosR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "ssb-IndexServing-r16"},
		{Name: "ssb-Ncell-r16"},
		{Name: "dl-PRS-r16"},
	},
}

const (
	SRS_PosResourceSet_r16_PathlossReferenceRS_Pos_r16_Ssb_IndexServing_r16 = 0
	SRS_PosResourceSet_r16_PathlossReferenceRS_Pos_r16_Ssb_Ncell_r16        = 1
	SRS_PosResourceSet_r16_PathlossReferenceRS_Pos_r16_Dl_PRS_r16           = 2
)

var sRSPosResourceSetR16ResourceTypeR16AperiodicR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "aperiodicSRS-ResourceTriggerList-r16", Optional: true},
	},
}

var sRSPosResourceSetR16ResourceTypeR16AperiodicR16AperiodicSRSResourceTriggerListR16Constraints = per.SizeRange(1, common.MaxNrofSRS_TriggerStates_1)

var sRSPosResourceSetR16ResourceTypeR16SemiPersistentR16Constraints = per.SequenceConstraints{
	Extensible: true,
}

var sRSPosResourceSetR16ResourceTypeR16PeriodicR16Constraints = per.SequenceConstraints{
	Extensible: true,
}

type SRS_PosResourceSet_r16 struct {
	Srs_PosResourceSetId_r16  SRS_PosResourceSetId_r16
	Srs_PosResourceIdList_r16 []SRS_PosResourceId_r16
	ResourceType_r16          struct {
		Choice              int
		Aperiodic_r16       *struct{ AperiodicSRS_ResourceTriggerList_r16 []int64 }
		Semi_Persistent_r16 *struct{}
		Periodic_r16        *struct{}
	}
	Alpha_r16                   *Alpha
	P0_r16                      *int64
	PathlossReferenceRS_Pos_r16 *struct {
		Choice               int
		Ssb_IndexServing_r16 *SSB_Index
		Ssb_Ncell_r16        *SSB_InfoNcell_r16
		Dl_PRS_r16           *DL_PRS_Info_r16
	}
}

func (ie *SRS_PosResourceSet_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sRSPosResourceSetR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Srs_PosResourceIdList_r16 != nil, ie.Alpha_r16 != nil, ie.P0_r16 != nil, ie.PathlossReferenceRS_Pos_r16 != nil}); err != nil {
		return err
	}

	// 3. srs-PosResourceSetId-r16: ref
	{
		if err := ie.Srs_PosResourceSetId_r16.Encode(e); err != nil {
			return err
		}
	}

	// 4. srs-PosResourceIdList-r16: sequence-of
	{
		if ie.Srs_PosResourceIdList_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(sRSPosResourceSetR16SrsPosResourceIdListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Srs_PosResourceIdList_r16))); err != nil {
				return err
			}
			for i := range ie.Srs_PosResourceIdList_r16 {
				if err := ie.Srs_PosResourceIdList_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. resourceType-r16: choice
	{
		choiceEnc := e.NewChoiceEncoder(sRS_PosResourceSet_r16ResourceTypeR16Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.ResourceType_r16.Choice), false, nil); err != nil {
			return err
		}
		switch ie.ResourceType_r16.Choice {
		case SRS_PosResourceSet_r16_ResourceType_r16_Aperiodic_r16:
			c := (*ie.ResourceType_r16.Aperiodic_r16)
			sRSPosResourceSetR16ResourceTypeR16AperiodicR16Seq := e.NewSequenceEncoder(sRSPosResourceSetR16ResourceTypeR16AperiodicR16Constraints)
			if err := sRSPosResourceSetR16ResourceTypeR16AperiodicR16Seq.EncodeExtensionBit(false); err != nil {
				return err
			}
			if err := sRSPosResourceSetR16ResourceTypeR16AperiodicR16Seq.EncodePreamble([]bool{c.AperiodicSRS_ResourceTriggerList_r16 != nil}); err != nil {
				return err
			}
			if c.AperiodicSRS_ResourceTriggerList_r16 != nil {
				seqOf := e.NewSequenceOfEncoder(sRSPosResourceSetR16ResourceTypeR16AperiodicR16AperiodicSRSResourceTriggerListR16Constraints)
				if err := seqOf.EncodeLength(int64(len(c.AperiodicSRS_ResourceTriggerList_r16))); err != nil {
					return err
				}
				for i := range c.AperiodicSRS_ResourceTriggerList_r16 {
					if err := e.EncodeInteger(c.AperiodicSRS_ResourceTriggerList_r16[i], per.Constrained(1, common.MaxNrofSRS_TriggerStates_1)); err != nil {
						return err
					}
				}
			}
		case SRS_PosResourceSet_r16_ResourceType_r16_Semi_Persistent_r16:
			sRSPosResourceSetR16ResourceTypeR16SemiPersistentR16Seq := e.NewSequenceEncoder(sRSPosResourceSetR16ResourceTypeR16SemiPersistentR16Constraints)
			if err := sRSPosResourceSetR16ResourceTypeR16SemiPersistentR16Seq.EncodeExtensionBit(false); err != nil {
				return err
			}
		case SRS_PosResourceSet_r16_ResourceType_r16_Periodic_r16:
			sRSPosResourceSetR16ResourceTypeR16PeriodicR16Seq := e.NewSequenceEncoder(sRSPosResourceSetR16ResourceTypeR16PeriodicR16Constraints)
			if err := sRSPosResourceSetR16ResourceTypeR16PeriodicR16Seq.EncodeExtensionBit(false); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.ResourceType_r16.Choice), Constraint: "index out of range"}
		}
	}

	// 6. alpha-r16: ref
	{
		if ie.Alpha_r16 != nil {
			if err := ie.Alpha_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. p0-r16: integer
	{
		if ie.P0_r16 != nil {
			if err := e.EncodeInteger(*ie.P0_r16, sRSPosResourceSetR16P0R16Constraints); err != nil {
				return err
			}
		}
	}

	// 8. pathlossReferenceRS-Pos-r16: choice
	{
		if ie.PathlossReferenceRS_Pos_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(sRS_PosResourceSet_r16PathlossReferenceRSPosR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.PathlossReferenceRS_Pos_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.PathlossReferenceRS_Pos_r16).Choice {
			case SRS_PosResourceSet_r16_PathlossReferenceRS_Pos_r16_Ssb_IndexServing_r16:
				if err := (*ie.PathlossReferenceRS_Pos_r16).Ssb_IndexServing_r16.Encode(e); err != nil {
					return err
				}
			case SRS_PosResourceSet_r16_PathlossReferenceRS_Pos_r16_Ssb_Ncell_r16:
				if err := (*ie.PathlossReferenceRS_Pos_r16).Ssb_Ncell_r16.Encode(e); err != nil {
					return err
				}
			case SRS_PosResourceSet_r16_PathlossReferenceRS_Pos_r16_Dl_PRS_r16:
				if err := (*ie.PathlossReferenceRS_Pos_r16).Dl_PRS_r16.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.PathlossReferenceRS_Pos_r16).Choice), Constraint: "index out of range"}
			}
		}
	}

	return nil
}

func (ie *SRS_PosResourceSet_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sRSPosResourceSetR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. srs-PosResourceSetId-r16: ref
	{
		if err := ie.Srs_PosResourceSetId_r16.Decode(d); err != nil {
			return err
		}
	}

	// 4. srs-PosResourceIdList-r16: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(sRSPosResourceSetR16SrsPosResourceIdListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Srs_PosResourceIdList_r16 = make([]SRS_PosResourceId_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Srs_PosResourceIdList_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. resourceType-r16: choice
	{
		choiceDec := d.NewChoiceDecoder(sRS_PosResourceSet_r16ResourceTypeR16Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.ResourceType_r16.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case SRS_PosResourceSet_r16_ResourceType_r16_Aperiodic_r16:
			ie.ResourceType_r16.Aperiodic_r16 = &struct{ AperiodicSRS_ResourceTriggerList_r16 []int64 }{}
			c := (*ie.ResourceType_r16.Aperiodic_r16)
			sRSPosResourceSetR16ResourceTypeR16AperiodicR16Seq := d.NewSequenceDecoder(sRSPosResourceSetR16ResourceTypeR16AperiodicR16Constraints)
			if err := sRSPosResourceSetR16ResourceTypeR16AperiodicR16Seq.DecodeExtensionBit(); err != nil {
				return err
			}
			if err := sRSPosResourceSetR16ResourceTypeR16AperiodicR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if sRSPosResourceSetR16ResourceTypeR16AperiodicR16Seq.IsComponentPresent(0) {
				seqOf := d.NewSequenceOfDecoder(sRSPosResourceSetR16ResourceTypeR16AperiodicR16AperiodicSRSResourceTriggerListR16Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.AperiodicSRS_ResourceTriggerList_r16 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(1, common.MaxNrofSRS_TriggerStates_1))
					if err != nil {
						return err
					}
					c.AperiodicSRS_ResourceTriggerList_r16[i] = v
				}
			}
		case SRS_PosResourceSet_r16_ResourceType_r16_Semi_Persistent_r16:
			ie.ResourceType_r16.Semi_Persistent_r16 = &struct{}{}
			sRSPosResourceSetR16ResourceTypeR16SemiPersistentR16Seq := d.NewSequenceDecoder(sRSPosResourceSetR16ResourceTypeR16SemiPersistentR16Constraints)
			if err := sRSPosResourceSetR16ResourceTypeR16SemiPersistentR16Seq.DecodeExtensionBit(); err != nil {
				return err
			}
		case SRS_PosResourceSet_r16_ResourceType_r16_Periodic_r16:
			ie.ResourceType_r16.Periodic_r16 = &struct{}{}
			sRSPosResourceSetR16ResourceTypeR16PeriodicR16Seq := d.NewSequenceDecoder(sRSPosResourceSetR16ResourceTypeR16PeriodicR16Constraints)
			if err := sRSPosResourceSetR16ResourceTypeR16PeriodicR16Seq.DecodeExtensionBit(); err != nil {
				return err
			}
		}
	}

	// 6. alpha-r16: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Alpha_r16 = new(Alpha)
			if err := ie.Alpha_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. p0-r16: integer
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeInteger(sRSPosResourceSetR16P0R16Constraints)
			if err != nil {
				return err
			}
			ie.P0_r16 = &v
		}
	}

	// 8. pathlossReferenceRS-Pos-r16: choice
	{
		if seq.IsComponentPresent(5) {
			ie.PathlossReferenceRS_Pos_r16 = &struct {
				Choice               int
				Ssb_IndexServing_r16 *SSB_Index
				Ssb_Ncell_r16        *SSB_InfoNcell_r16
				Dl_PRS_r16           *DL_PRS_Info_r16
			}{}
			choiceDec := d.NewChoiceDecoder(sRS_PosResourceSet_r16PathlossReferenceRSPosR16Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.PathlossReferenceRS_Pos_r16).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case SRS_PosResourceSet_r16_PathlossReferenceRS_Pos_r16_Ssb_IndexServing_r16:
				(*ie.PathlossReferenceRS_Pos_r16).Ssb_IndexServing_r16 = new(SSB_Index)
				if err := (*ie.PathlossReferenceRS_Pos_r16).Ssb_IndexServing_r16.Decode(d); err != nil {
					return err
				}
			case SRS_PosResourceSet_r16_PathlossReferenceRS_Pos_r16_Ssb_Ncell_r16:
				(*ie.PathlossReferenceRS_Pos_r16).Ssb_Ncell_r16 = new(SSB_InfoNcell_r16)
				if err := (*ie.PathlossReferenceRS_Pos_r16).Ssb_Ncell_r16.Decode(d); err != nil {
					return err
				}
			case SRS_PosResourceSet_r16_PathlossReferenceRS_Pos_r16_Dl_PRS_r16:
				(*ie.PathlossReferenceRS_Pos_r16).Dl_PRS_r16 = new(DL_PRS_Info_r16)
				if err := (*ie.PathlossReferenceRS_Pos_r16).Dl_PRS_r16.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
