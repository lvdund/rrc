// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MinTimeGapFR2-2-r17 (line 21083).

var minTimeGapFR22R17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-120kHz-r17", Optional: true},
		{Name: "scs-480kHz-r17", Optional: true},
		{Name: "scs-960kHz-r17", Optional: true},
	},
}

const (
	MinTimeGapFR2_2_r17_Scs_120kHz_r17_Sl2  = 0
	MinTimeGapFR2_2_r17_Scs_120kHz_r17_Sl24 = 1
)

var minTimeGapFR22R17Scs120kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MinTimeGapFR2_2_r17_Scs_120kHz_r17_Sl2, MinTimeGapFR2_2_r17_Scs_120kHz_r17_Sl24},
}

const (
	MinTimeGapFR2_2_r17_Scs_480kHz_r17_Sl8  = 0
	MinTimeGapFR2_2_r17_Scs_480kHz_r17_Sl96 = 1
)

var minTimeGapFR22R17Scs480kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MinTimeGapFR2_2_r17_Scs_480kHz_r17_Sl8, MinTimeGapFR2_2_r17_Scs_480kHz_r17_Sl96},
}

const (
	MinTimeGapFR2_2_r17_Scs_960kHz_r17_Sl16  = 0
	MinTimeGapFR2_2_r17_Scs_960kHz_r17_Sl192 = 1
)

var minTimeGapFR22R17Scs960kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MinTimeGapFR2_2_r17_Scs_960kHz_r17_Sl16, MinTimeGapFR2_2_r17_Scs_960kHz_r17_Sl192},
}

type MinTimeGapFR2_2_r17 struct {
	Scs_120kHz_r17 *int64
	Scs_480kHz_r17 *int64
	Scs_960kHz_r17 *int64
}

func (ie *MinTimeGapFR2_2_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(minTimeGapFR22R17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Scs_120kHz_r17 != nil, ie.Scs_480kHz_r17 != nil, ie.Scs_960kHz_r17 != nil}); err != nil {
		return err
	}

	// 2. scs-120kHz-r17: enumerated
	{
		if ie.Scs_120kHz_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Scs_120kHz_r17, minTimeGapFR22R17Scs120kHzR17Constraints); err != nil {
				return err
			}
		}
	}

	// 3. scs-480kHz-r17: enumerated
	{
		if ie.Scs_480kHz_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Scs_480kHz_r17, minTimeGapFR22R17Scs480kHzR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. scs-960kHz-r17: enumerated
	{
		if ie.Scs_960kHz_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Scs_960kHz_r17, minTimeGapFR22R17Scs960kHzR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MinTimeGapFR2_2_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(minTimeGapFR22R17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. scs-120kHz-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(minTimeGapFR22R17Scs120kHzR17Constraints)
			if err != nil {
				return err
			}
			ie.Scs_120kHz_r17 = &idx
		}
	}

	// 3. scs-480kHz-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(minTimeGapFR22R17Scs480kHzR17Constraints)
			if err != nil {
				return err
			}
			ie.Scs_480kHz_r17 = &idx
		}
	}

	// 4. scs-960kHz-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(minTimeGapFR22R17Scs960kHzR17Constraints)
			if err != nil {
				return err
			}
			ie.Scs_960kHz_r17 = &idx
		}
	}

	return nil
}
