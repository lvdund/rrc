// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ReferenceTimeInfo-r16 (line 13321).

var referenceTimeInfoR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "time-r16"},
		{Name: "uncertainty-r16", Optional: true},
		{Name: "timeInfoType-r16", Optional: true},
		{Name: "referenceSFN-r16", Optional: true},
	},
}

var referenceTimeInfoR16UncertaintyR16Constraints = per.Constrained(0, 32767)

const (
	ReferenceTimeInfo_r16_TimeInfoType_r16_LocalClock = 0
)

var referenceTimeInfoR16TimeInfoTypeR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ReferenceTimeInfo_r16_TimeInfoType_r16_LocalClock},
}

var referenceTimeInfoR16ReferenceSFNR16Constraints = per.Constrained(0, 1023)

type ReferenceTimeInfo_r16 struct {
	Time_r16         ReferenceTime_r16
	Uncertainty_r16  *int64
	TimeInfoType_r16 *int64
	ReferenceSFN_r16 *int64
}

func (ie *ReferenceTimeInfo_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(referenceTimeInfoR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Uncertainty_r16 != nil, ie.TimeInfoType_r16 != nil, ie.ReferenceSFN_r16 != nil}); err != nil {
		return err
	}

	// 2. time-r16: ref
	{
		if err := ie.Time_r16.Encode(e); err != nil {
			return err
		}
	}

	// 3. uncertainty-r16: integer
	{
		if ie.Uncertainty_r16 != nil {
			if err := e.EncodeInteger(*ie.Uncertainty_r16, referenceTimeInfoR16UncertaintyR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. timeInfoType-r16: enumerated
	{
		if ie.TimeInfoType_r16 != nil {
			if err := e.EncodeEnumerated(*ie.TimeInfoType_r16, referenceTimeInfoR16TimeInfoTypeR16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. referenceSFN-r16: integer
	{
		if ie.ReferenceSFN_r16 != nil {
			if err := e.EncodeInteger(*ie.ReferenceSFN_r16, referenceTimeInfoR16ReferenceSFNR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *ReferenceTimeInfo_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(referenceTimeInfoR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. time-r16: ref
	{
		if err := ie.Time_r16.Decode(d); err != nil {
			return err
		}
	}

	// 3. uncertainty-r16: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(referenceTimeInfoR16UncertaintyR16Constraints)
			if err != nil {
				return err
			}
			ie.Uncertainty_r16 = &v
		}
	}

	// 4. timeInfoType-r16: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(referenceTimeInfoR16TimeInfoTypeR16Constraints)
			if err != nil {
				return err
			}
			ie.TimeInfoType_r16 = &idx
		}
	}

	// 5. referenceSFN-r16: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(referenceTimeInfoR16ReferenceSFNR16Constraints)
			if err != nil {
				return err
			}
			ie.ReferenceSFN_r16 = &v
		}
	}

	return nil
}
