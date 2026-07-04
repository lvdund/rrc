// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MinTimeGap-r16 (line 21076).

var minTimeGapR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz-r16", Optional: true},
		{Name: "scs-30kHz-r16", Optional: true},
		{Name: "scs-60kHz-r16", Optional: true},
		{Name: "scs-120kHz-r16", Optional: true},
	},
}

const (
	MinTimeGap_r16_Scs_15kHz_r16_Sl1 = 0
	MinTimeGap_r16_Scs_15kHz_r16_Sl3 = 1
)

var minTimeGapR16Scs15kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MinTimeGap_r16_Scs_15kHz_r16_Sl1, MinTimeGap_r16_Scs_15kHz_r16_Sl3},
}

const (
	MinTimeGap_r16_Scs_30kHz_r16_Sl1 = 0
	MinTimeGap_r16_Scs_30kHz_r16_Sl6 = 1
)

var minTimeGapR16Scs30kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MinTimeGap_r16_Scs_30kHz_r16_Sl1, MinTimeGap_r16_Scs_30kHz_r16_Sl6},
}

const (
	MinTimeGap_r16_Scs_60kHz_r16_Sl1  = 0
	MinTimeGap_r16_Scs_60kHz_r16_Sl12 = 1
)

var minTimeGapR16Scs60kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MinTimeGap_r16_Scs_60kHz_r16_Sl1, MinTimeGap_r16_Scs_60kHz_r16_Sl12},
}

const (
	MinTimeGap_r16_Scs_120kHz_r16_Sl2  = 0
	MinTimeGap_r16_Scs_120kHz_r16_Sl24 = 1
)

var minTimeGapR16Scs120kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MinTimeGap_r16_Scs_120kHz_r16_Sl2, MinTimeGap_r16_Scs_120kHz_r16_Sl24},
}

type MinTimeGap_r16 struct {
	Scs_15kHz_r16  *int64
	Scs_30kHz_r16  *int64
	Scs_60kHz_r16  *int64
	Scs_120kHz_r16 *int64
}

func (ie *MinTimeGap_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(minTimeGapR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Scs_15kHz_r16 != nil, ie.Scs_30kHz_r16 != nil, ie.Scs_60kHz_r16 != nil, ie.Scs_120kHz_r16 != nil}); err != nil {
		return err
	}

	// 2. scs-15kHz-r16: enumerated
	{
		if ie.Scs_15kHz_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Scs_15kHz_r16, minTimeGapR16Scs15kHzR16Constraints); err != nil {
				return err
			}
		}
	}

	// 3. scs-30kHz-r16: enumerated
	{
		if ie.Scs_30kHz_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Scs_30kHz_r16, minTimeGapR16Scs30kHzR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. scs-60kHz-r16: enumerated
	{
		if ie.Scs_60kHz_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Scs_60kHz_r16, minTimeGapR16Scs60kHzR16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. scs-120kHz-r16: enumerated
	{
		if ie.Scs_120kHz_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Scs_120kHz_r16, minTimeGapR16Scs120kHzR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MinTimeGap_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(minTimeGapR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. scs-15kHz-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(minTimeGapR16Scs15kHzR16Constraints)
			if err != nil {
				return err
			}
			ie.Scs_15kHz_r16 = &idx
		}
	}

	// 3. scs-30kHz-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(minTimeGapR16Scs30kHzR16Constraints)
			if err != nil {
				return err
			}
			ie.Scs_30kHz_r16 = &idx
		}
	}

	// 4. scs-60kHz-r16: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(minTimeGapR16Scs60kHzR16Constraints)
			if err != nil {
				return err
			}
			ie.Scs_60kHz_r16 = &idx
		}
	}

	// 5. scs-120kHz-r16: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(minTimeGapR16Scs120kHzR16Constraints)
			if err != nil {
				return err
			}
			ie.Scs_120kHz_r16 = &idx
		}
	}

	return nil
}
