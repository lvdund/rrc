// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FeatureCombination-r17 (line 8305).

var featureCombinationR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "redCap-r17", Optional: true},
		{Name: "smallData-r17", Optional: true},
		{Name: "nsag-r17", Optional: true},
		{Name: "msg3-Repetitions-r17", Optional: true},
		{Name: "msg1-Repetitions-r18", Optional: true},
		{Name: "eRedCap-r18", Optional: true},
		{Name: "spare2", Optional: true},
		{Name: "spare1", Optional: true},
	},
}

const (
	FeatureCombination_r17_RedCap_r17_True = 0
)

var featureCombinationR17RedCapR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureCombination_r17_RedCap_r17_True},
}

const (
	FeatureCombination_r17_SmallData_r17_True = 0
)

var featureCombinationR17SmallDataR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureCombination_r17_SmallData_r17_True},
}

const (
	FeatureCombination_r17_Msg3_Repetitions_r17_True = 0
)

var featureCombinationR17Msg3RepetitionsR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureCombination_r17_Msg3_Repetitions_r17_True},
}

const (
	FeatureCombination_r17_Msg1_Repetitions_r18_True = 0
)

var featureCombinationR17Msg1RepetitionsR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureCombination_r17_Msg1_Repetitions_r18_True},
}

const (
	FeatureCombination_r17_ERedCap_r18_True = 0
)

var featureCombinationR17ERedCapR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureCombination_r17_ERedCap_r18_True},
}

const (
	FeatureCombination_r17_Spare2_True = 0
)

var featureCombinationR17Spare2Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureCombination_r17_Spare2_True},
}

const (
	FeatureCombination_r17_Spare1_True = 0
)

var featureCombinationR17Spare1Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureCombination_r17_Spare1_True},
}

type FeatureCombination_r17 struct {
	RedCap_r17           *int64
	SmallData_r17        *int64
	Nsag_r17             *NSAG_List_r17
	Msg3_Repetitions_r17 *int64
	Msg1_Repetitions_r18 *int64
	ERedCap_r18          *int64
	Spare2               *int64
	Spare1               *int64
}

func (ie *FeatureCombination_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureCombinationR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.RedCap_r17 != nil, ie.SmallData_r17 != nil, ie.Nsag_r17 != nil, ie.Msg3_Repetitions_r17 != nil, ie.Msg1_Repetitions_r18 != nil, ie.ERedCap_r18 != nil, ie.Spare2 != nil, ie.Spare1 != nil}); err != nil {
		return err
	}

	// 2. redCap-r17: enumerated
	{
		if ie.RedCap_r17 != nil {
			if err := e.EncodeEnumerated(*ie.RedCap_r17, featureCombinationR17RedCapR17Constraints); err != nil {
				return err
			}
		}
	}

	// 3. smallData-r17: enumerated
	{
		if ie.SmallData_r17 != nil {
			if err := e.EncodeEnumerated(*ie.SmallData_r17, featureCombinationR17SmallDataR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. nsag-r17: ref
	{
		if ie.Nsag_r17 != nil {
			if err := ie.Nsag_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. msg3-Repetitions-r17: enumerated
	{
		if ie.Msg3_Repetitions_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Msg3_Repetitions_r17, featureCombinationR17Msg3RepetitionsR17Constraints); err != nil {
				return err
			}
		}
	}

	// 6. msg1-Repetitions-r18: enumerated
	{
		if ie.Msg1_Repetitions_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Msg1_Repetitions_r18, featureCombinationR17Msg1RepetitionsR18Constraints); err != nil {
				return err
			}
		}
	}

	// 7. eRedCap-r18: enumerated
	{
		if ie.ERedCap_r18 != nil {
			if err := e.EncodeEnumerated(*ie.ERedCap_r18, featureCombinationR17ERedCapR18Constraints); err != nil {
				return err
			}
		}
	}

	// 8. spare2: enumerated
	{
		if ie.Spare2 != nil {
			if err := e.EncodeEnumerated(*ie.Spare2, featureCombinationR17Spare2Constraints); err != nil {
				return err
			}
		}
	}

	// 9. spare1: enumerated
	{
		if ie.Spare1 != nil {
			if err := e.EncodeEnumerated(*ie.Spare1, featureCombinationR17Spare1Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FeatureCombination_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureCombinationR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. redCap-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(featureCombinationR17RedCapR17Constraints)
			if err != nil {
				return err
			}
			ie.RedCap_r17 = &idx
		}
	}

	// 3. smallData-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(featureCombinationR17SmallDataR17Constraints)
			if err != nil {
				return err
			}
			ie.SmallData_r17 = &idx
		}
	}

	// 4. nsag-r17: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Nsag_r17 = new(NSAG_List_r17)
			if err := ie.Nsag_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. msg3-Repetitions-r17: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(featureCombinationR17Msg3RepetitionsR17Constraints)
			if err != nil {
				return err
			}
			ie.Msg3_Repetitions_r17 = &idx
		}
	}

	// 6. msg1-Repetitions-r18: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(featureCombinationR17Msg1RepetitionsR18Constraints)
			if err != nil {
				return err
			}
			ie.Msg1_Repetitions_r18 = &idx
		}
	}

	// 7. eRedCap-r18: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(featureCombinationR17ERedCapR18Constraints)
			if err != nil {
				return err
			}
			ie.ERedCap_r18 = &idx
		}
	}

	// 8. spare2: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(featureCombinationR17Spare2Constraints)
			if err != nil {
				return err
			}
			ie.Spare2 = &idx
		}
	}

	// 9. spare1: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(featureCombinationR17Spare1Constraints)
			if err != nil {
				return err
			}
			ie.Spare1 = &idx
		}
	}

	return nil
}
