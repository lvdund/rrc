// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: InterFreqCarrierFreqInfo-v1900 (line 4046).

var interFreqCarrierFreqInfoV1900Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "uav-PrioritizedFrequency-r19", Optional: true},
		{Name: "uav-PrioritizedFrequencyAltitudeRange-r19", Optional: true},
		{Name: "ssb-ToMeasureAltitudeBasedList-r19", Optional: true},
		{Name: "interFreqOD-SIB1-ExcludedCellList-r19", Optional: true},
		{Name: "od-SIB1-cellReselectionPriority-r19", Optional: true},
		{Name: "od-SIB1-cellReselectionSubPriority-r19", Optional: true},
		{Name: "refLocMapList-r19", Optional: true},
		{Name: "refLocOffset-r19", Optional: true},
		{Name: "smtc5list-r19", Optional: true},
	},
}

const (
	InterFreqCarrierFreqInfo_v1900_Uav_PrioritizedFrequency_r19_True = 0
)

var interFreqCarrierFreqInfoV1900UavPrioritizedFrequencyR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{InterFreqCarrierFreqInfo_v1900_Uav_PrioritizedFrequency_r19_True},
}

var interFreqCarrierFreqInfo_v1900InterFreqODSIB1ExcludedCellListR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "emptyList-r19"},
		{Name: "excludedCells-r19"},
	},
}

const (
	InterFreqCarrierFreqInfo_v1900_InterFreqOD_SIB1_ExcludedCellList_r19_EmptyList_r19     = 0
	InterFreqCarrierFreqInfo_v1900_InterFreqOD_SIB1_ExcludedCellList_r19_ExcludedCells_r19 = 1
)

var interFreqCarrierFreqInfoV1900RefLocMapListR19Constraints = per.SizeRange(1, 7)

var interFreqCarrierFreqInfoV1900Smtc5listR19Constraints = per.SizeRange(1, 6)

var interFreqCarrierFreqInfoV1900UavPrioritizedFrequencyAltitudeRangeR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "altitudeMin-r19", Optional: true},
		{Name: "altitudeMax-r19", Optional: true},
		{Name: "altitudeHyst-r19", Optional: true},
	},
}

type InterFreqCarrierFreqInfo_v1900 struct {
	Uav_PrioritizedFrequency_r19              *int64
	Uav_PrioritizedFrequencyAltitudeRange_r19 *struct {
		AltitudeMin_r19  *Altitude_r18
		AltitudeMax_r19  *Altitude_r18
		AltitudeHyst_r19 *HysteresisAltitude_r18
	}
	Ssb_ToMeasureAltitudeBasedList_r19    *SSB_ToMeasureAltitudeBasedList_r18
	InterFreqOD_SIB1_ExcludedCellList_r19 *struct {
		Choice            int
		EmptyList_r19     *struct{}
		ExcludedCells_r19 *InterFreqExcludedCellList
	}
	Od_SIB1_CellReselectionPriority_r19    *CellReselectionPriority
	Od_SIB1_CellReselectionSubPriority_r19 *CellReselectionSubPriority
	RefLocMapList_r19                      []RefLocMap_r19
	RefLocOffset_r19                       *RefLocOffset_r19
	Smtc5list_r19                          []SSB_MTC5_r19
}

func (ie *InterFreqCarrierFreqInfo_v1900) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(interFreqCarrierFreqInfoV1900Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Uav_PrioritizedFrequency_r19 != nil, ie.Uav_PrioritizedFrequencyAltitudeRange_r19 != nil, ie.Ssb_ToMeasureAltitudeBasedList_r19 != nil, ie.InterFreqOD_SIB1_ExcludedCellList_r19 != nil, ie.Od_SIB1_CellReselectionPriority_r19 != nil, ie.Od_SIB1_CellReselectionSubPriority_r19 != nil, ie.RefLocMapList_r19 != nil, ie.RefLocOffset_r19 != nil, ie.Smtc5list_r19 != nil}); err != nil {
		return err
	}

	// 2. uav-PrioritizedFrequency-r19: enumerated
	{
		if ie.Uav_PrioritizedFrequency_r19 != nil {
			if err := e.EncodeEnumerated(*ie.Uav_PrioritizedFrequency_r19, interFreqCarrierFreqInfoV1900UavPrioritizedFrequencyR19Constraints); err != nil {
				return err
			}
		}
	}

	// 3. uav-PrioritizedFrequencyAltitudeRange-r19: sequence
	{
		if ie.Uav_PrioritizedFrequencyAltitudeRange_r19 != nil {
			c := ie.Uav_PrioritizedFrequencyAltitudeRange_r19
			interFreqCarrierFreqInfoV1900UavPrioritizedFrequencyAltitudeRangeR19Seq := e.NewSequenceEncoder(interFreqCarrierFreqInfoV1900UavPrioritizedFrequencyAltitudeRangeR19Constraints)
			if err := interFreqCarrierFreqInfoV1900UavPrioritizedFrequencyAltitudeRangeR19Seq.EncodePreamble([]bool{c.AltitudeMin_r19 != nil, c.AltitudeMax_r19 != nil, c.AltitudeHyst_r19 != nil}); err != nil {
				return err
			}
			if c.AltitudeMin_r19 != nil {
				if err := c.AltitudeMin_r19.Encode(e); err != nil {
					return err
				}
			}
			if c.AltitudeMax_r19 != nil {
				if err := c.AltitudeMax_r19.Encode(e); err != nil {
					return err
				}
			}
			if c.AltitudeHyst_r19 != nil {
				if err := c.AltitudeHyst_r19.Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. ssb-ToMeasureAltitudeBasedList-r19: ref
	{
		if ie.Ssb_ToMeasureAltitudeBasedList_r19 != nil {
			if err := ie.Ssb_ToMeasureAltitudeBasedList_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. interFreqOD-SIB1-ExcludedCellList-r19: choice
	{
		if ie.InterFreqOD_SIB1_ExcludedCellList_r19 != nil {
			choiceEnc := e.NewChoiceEncoder(interFreqCarrierFreqInfo_v1900InterFreqODSIB1ExcludedCellListR19Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.InterFreqOD_SIB1_ExcludedCellList_r19).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.InterFreqOD_SIB1_ExcludedCellList_r19).Choice {
			case InterFreqCarrierFreqInfo_v1900_InterFreqOD_SIB1_ExcludedCellList_r19_EmptyList_r19:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case InterFreqCarrierFreqInfo_v1900_InterFreqOD_SIB1_ExcludedCellList_r19_ExcludedCells_r19:
				if err := (*ie.InterFreqOD_SIB1_ExcludedCellList_r19).ExcludedCells_r19.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.InterFreqOD_SIB1_ExcludedCellList_r19).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 6. od-SIB1-cellReselectionPriority-r19: ref
	{
		if ie.Od_SIB1_CellReselectionPriority_r19 != nil {
			if err := ie.Od_SIB1_CellReselectionPriority_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. od-SIB1-cellReselectionSubPriority-r19: ref
	{
		if ie.Od_SIB1_CellReselectionSubPriority_r19 != nil {
			if err := ie.Od_SIB1_CellReselectionSubPriority_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. refLocMapList-r19: sequence-of
	{
		if ie.RefLocMapList_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(interFreqCarrierFreqInfoV1900RefLocMapListR19Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.RefLocMapList_r19))); err != nil {
				return err
			}
			for i := range ie.RefLocMapList_r19 {
				if err := ie.RefLocMapList_r19[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 9. refLocOffset-r19: ref
	{
		if ie.RefLocOffset_r19 != nil {
			if err := ie.RefLocOffset_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 10. smtc5list-r19: sequence-of
	{
		if ie.Smtc5list_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(interFreqCarrierFreqInfoV1900Smtc5listR19Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Smtc5list_r19))); err != nil {
				return err
			}
			for i := range ie.Smtc5list_r19 {
				if err := ie.Smtc5list_r19[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *InterFreqCarrierFreqInfo_v1900) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(interFreqCarrierFreqInfoV1900Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. uav-PrioritizedFrequency-r19: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(interFreqCarrierFreqInfoV1900UavPrioritizedFrequencyR19Constraints)
			if err != nil {
				return err
			}
			ie.Uav_PrioritizedFrequency_r19 = &idx
		}
	}

	// 3. uav-PrioritizedFrequencyAltitudeRange-r19: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.Uav_PrioritizedFrequencyAltitudeRange_r19 = &struct {
				AltitudeMin_r19  *Altitude_r18
				AltitudeMax_r19  *Altitude_r18
				AltitudeHyst_r19 *HysteresisAltitude_r18
			}{}
			c := ie.Uav_PrioritizedFrequencyAltitudeRange_r19
			interFreqCarrierFreqInfoV1900UavPrioritizedFrequencyAltitudeRangeR19Seq := d.NewSequenceDecoder(interFreqCarrierFreqInfoV1900UavPrioritizedFrequencyAltitudeRangeR19Constraints)
			if err := interFreqCarrierFreqInfoV1900UavPrioritizedFrequencyAltitudeRangeR19Seq.DecodePreamble(); err != nil {
				return err
			}
			if interFreqCarrierFreqInfoV1900UavPrioritizedFrequencyAltitudeRangeR19Seq.IsComponentPresent(0) {
				c.AltitudeMin_r19 = new(Altitude_r18)
				if err := (*c.AltitudeMin_r19).Decode(d); err != nil {
					return err
				}
			}
			if interFreqCarrierFreqInfoV1900UavPrioritizedFrequencyAltitudeRangeR19Seq.IsComponentPresent(1) {
				c.AltitudeMax_r19 = new(Altitude_r18)
				if err := (*c.AltitudeMax_r19).Decode(d); err != nil {
					return err
				}
			}
			if interFreqCarrierFreqInfoV1900UavPrioritizedFrequencyAltitudeRangeR19Seq.IsComponentPresent(2) {
				c.AltitudeHyst_r19 = new(HysteresisAltitude_r18)
				if err := (*c.AltitudeHyst_r19).Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. ssb-ToMeasureAltitudeBasedList-r19: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Ssb_ToMeasureAltitudeBasedList_r19 = new(SSB_ToMeasureAltitudeBasedList_r18)
			if err := ie.Ssb_ToMeasureAltitudeBasedList_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. interFreqOD-SIB1-ExcludedCellList-r19: choice
	{
		if seq.IsComponentPresent(3) {
			ie.InterFreqOD_SIB1_ExcludedCellList_r19 = &struct {
				Choice            int
				EmptyList_r19     *struct{}
				ExcludedCells_r19 *InterFreqExcludedCellList
			}{}
			choiceDec := d.NewChoiceDecoder(interFreqCarrierFreqInfo_v1900InterFreqODSIB1ExcludedCellListR19Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.InterFreqOD_SIB1_ExcludedCellList_r19).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case InterFreqCarrierFreqInfo_v1900_InterFreqOD_SIB1_ExcludedCellList_r19_EmptyList_r19:
				(*ie.InterFreqOD_SIB1_ExcludedCellList_r19).EmptyList_r19 = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case InterFreqCarrierFreqInfo_v1900_InterFreqOD_SIB1_ExcludedCellList_r19_ExcludedCells_r19:
				(*ie.InterFreqOD_SIB1_ExcludedCellList_r19).ExcludedCells_r19 = new(InterFreqExcludedCellList)
				if err := (*ie.InterFreqOD_SIB1_ExcludedCellList_r19).ExcludedCells_r19.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. od-SIB1-cellReselectionPriority-r19: ref
	{
		if seq.IsComponentPresent(4) {
			ie.Od_SIB1_CellReselectionPriority_r19 = new(CellReselectionPriority)
			if err := ie.Od_SIB1_CellReselectionPriority_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. od-SIB1-cellReselectionSubPriority-r19: ref
	{
		if seq.IsComponentPresent(5) {
			ie.Od_SIB1_CellReselectionSubPriority_r19 = new(CellReselectionSubPriority)
			if err := ie.Od_SIB1_CellReselectionSubPriority_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. refLocMapList-r19: sequence-of
	{
		if seq.IsComponentPresent(6) {
			seqOf := d.NewSequenceOfDecoder(interFreqCarrierFreqInfoV1900RefLocMapListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.RefLocMapList_r19 = make([]RefLocMap_r19, n)
			for i := int64(0); i < n; i++ {
				if err := ie.RefLocMapList_r19[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 9. refLocOffset-r19: ref
	{
		if seq.IsComponentPresent(7) {
			ie.RefLocOffset_r19 = new(RefLocOffset_r19)
			if err := ie.RefLocOffset_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 10. smtc5list-r19: sequence-of
	{
		if seq.IsComponentPresent(8) {
			seqOf := d.NewSequenceOfDecoder(interFreqCarrierFreqInfoV1900Smtc5listR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Smtc5list_r19 = make([]SSB_MTC5_r19, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Smtc5list_r19[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
