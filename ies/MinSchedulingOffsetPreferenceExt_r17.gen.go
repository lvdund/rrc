// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MinSchedulingOffsetPreferenceExt-r17 (line 2597).

var minSchedulingOffsetPreferenceExtR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "preferredK0-r17", Optional: true},
		{Name: "preferredK2-r17", Optional: true},
	},
}

var minSchedulingOffsetPreferenceExtR17PreferredK0R17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "preferredK0-SCS-480kHz-r17", Optional: true},
		{Name: "preferredK0-SCS-960kHz-r17", Optional: true},
	},
}

const (
	MinSchedulingOffsetPreferenceExt_r17_PreferredK0_r17_PreferredK0_SCS_480kHz_r17_Sl8  = 0
	MinSchedulingOffsetPreferenceExt_r17_PreferredK0_r17_PreferredK0_SCS_480kHz_r17_Sl16 = 1
	MinSchedulingOffsetPreferenceExt_r17_PreferredK0_r17_PreferredK0_SCS_480kHz_r17_Sl32 = 2
	MinSchedulingOffsetPreferenceExt_r17_PreferredK0_r17_PreferredK0_SCS_480kHz_r17_Sl48 = 3
)

var minSchedulingOffsetPreferenceExtR17PreferredK0R17PreferredK0SCS480kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MinSchedulingOffsetPreferenceExt_r17_PreferredK0_r17_PreferredK0_SCS_480kHz_r17_Sl8, MinSchedulingOffsetPreferenceExt_r17_PreferredK0_r17_PreferredK0_SCS_480kHz_r17_Sl16, MinSchedulingOffsetPreferenceExt_r17_PreferredK0_r17_PreferredK0_SCS_480kHz_r17_Sl32, MinSchedulingOffsetPreferenceExt_r17_PreferredK0_r17_PreferredK0_SCS_480kHz_r17_Sl48},
}

const (
	MinSchedulingOffsetPreferenceExt_r17_PreferredK0_r17_PreferredK0_SCS_960kHz_r17_Sl8  = 0
	MinSchedulingOffsetPreferenceExt_r17_PreferredK0_r17_PreferredK0_SCS_960kHz_r17_Sl16 = 1
	MinSchedulingOffsetPreferenceExt_r17_PreferredK0_r17_PreferredK0_SCS_960kHz_r17_Sl32 = 2
	MinSchedulingOffsetPreferenceExt_r17_PreferredK0_r17_PreferredK0_SCS_960kHz_r17_Sl48 = 3
)

var minSchedulingOffsetPreferenceExtR17PreferredK0R17PreferredK0SCS960kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MinSchedulingOffsetPreferenceExt_r17_PreferredK0_r17_PreferredK0_SCS_960kHz_r17_Sl8, MinSchedulingOffsetPreferenceExt_r17_PreferredK0_r17_PreferredK0_SCS_960kHz_r17_Sl16, MinSchedulingOffsetPreferenceExt_r17_PreferredK0_r17_PreferredK0_SCS_960kHz_r17_Sl32, MinSchedulingOffsetPreferenceExt_r17_PreferredK0_r17_PreferredK0_SCS_960kHz_r17_Sl48},
}

var minSchedulingOffsetPreferenceExtR17PreferredK2R17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "preferredK2-SCS-480kHz-r17", Optional: true},
		{Name: "preferredK2-SCS-960kHz-r17", Optional: true},
	},
}

const (
	MinSchedulingOffsetPreferenceExt_r17_PreferredK2_r17_PreferredK2_SCS_480kHz_r17_Sl8  = 0
	MinSchedulingOffsetPreferenceExt_r17_PreferredK2_r17_PreferredK2_SCS_480kHz_r17_Sl16 = 1
	MinSchedulingOffsetPreferenceExt_r17_PreferredK2_r17_PreferredK2_SCS_480kHz_r17_Sl32 = 2
	MinSchedulingOffsetPreferenceExt_r17_PreferredK2_r17_PreferredK2_SCS_480kHz_r17_Sl48 = 3
)

var minSchedulingOffsetPreferenceExtR17PreferredK2R17PreferredK2SCS480kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MinSchedulingOffsetPreferenceExt_r17_PreferredK2_r17_PreferredK2_SCS_480kHz_r17_Sl8, MinSchedulingOffsetPreferenceExt_r17_PreferredK2_r17_PreferredK2_SCS_480kHz_r17_Sl16, MinSchedulingOffsetPreferenceExt_r17_PreferredK2_r17_PreferredK2_SCS_480kHz_r17_Sl32, MinSchedulingOffsetPreferenceExt_r17_PreferredK2_r17_PreferredK2_SCS_480kHz_r17_Sl48},
}

const (
	MinSchedulingOffsetPreferenceExt_r17_PreferredK2_r17_PreferredK2_SCS_960kHz_r17_Sl8  = 0
	MinSchedulingOffsetPreferenceExt_r17_PreferredK2_r17_PreferredK2_SCS_960kHz_r17_Sl16 = 1
	MinSchedulingOffsetPreferenceExt_r17_PreferredK2_r17_PreferredK2_SCS_960kHz_r17_Sl32 = 2
	MinSchedulingOffsetPreferenceExt_r17_PreferredK2_r17_PreferredK2_SCS_960kHz_r17_Sl48 = 3
)

var minSchedulingOffsetPreferenceExtR17PreferredK2R17PreferredK2SCS960kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MinSchedulingOffsetPreferenceExt_r17_PreferredK2_r17_PreferredK2_SCS_960kHz_r17_Sl8, MinSchedulingOffsetPreferenceExt_r17_PreferredK2_r17_PreferredK2_SCS_960kHz_r17_Sl16, MinSchedulingOffsetPreferenceExt_r17_PreferredK2_r17_PreferredK2_SCS_960kHz_r17_Sl32, MinSchedulingOffsetPreferenceExt_r17_PreferredK2_r17_PreferredK2_SCS_960kHz_r17_Sl48},
}

type MinSchedulingOffsetPreferenceExt_r17 struct {
	PreferredK0_r17 *struct {
		PreferredK0_SCS_480kHz_r17 *int64
		PreferredK0_SCS_960kHz_r17 *int64
	}
	PreferredK2_r17 *struct {
		PreferredK2_SCS_480kHz_r17 *int64
		PreferredK2_SCS_960kHz_r17 *int64
	}
}

func (ie *MinSchedulingOffsetPreferenceExt_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(minSchedulingOffsetPreferenceExtR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.PreferredK0_r17 != nil, ie.PreferredK2_r17 != nil}); err != nil {
		return err
	}

	// 2. preferredK0-r17: sequence
	{
		if ie.PreferredK0_r17 != nil {
			c := ie.PreferredK0_r17
			minSchedulingOffsetPreferenceExtR17PreferredK0R17Seq := e.NewSequenceEncoder(minSchedulingOffsetPreferenceExtR17PreferredK0R17Constraints)
			if err := minSchedulingOffsetPreferenceExtR17PreferredK0R17Seq.EncodePreamble([]bool{c.PreferredK0_SCS_480kHz_r17 != nil, c.PreferredK0_SCS_960kHz_r17 != nil}); err != nil {
				return err
			}
			if c.PreferredK0_SCS_480kHz_r17 != nil {
				if err := e.EncodeEnumerated((*c.PreferredK0_SCS_480kHz_r17), minSchedulingOffsetPreferenceExtR17PreferredK0R17PreferredK0SCS480kHzR17Constraints); err != nil {
					return err
				}
			}
			if c.PreferredK0_SCS_960kHz_r17 != nil {
				if err := e.EncodeEnumerated((*c.PreferredK0_SCS_960kHz_r17), minSchedulingOffsetPreferenceExtR17PreferredK0R17PreferredK0SCS960kHzR17Constraints); err != nil {
					return err
				}
			}
		}
	}

	// 3. preferredK2-r17: sequence
	{
		if ie.PreferredK2_r17 != nil {
			c := ie.PreferredK2_r17
			minSchedulingOffsetPreferenceExtR17PreferredK2R17Seq := e.NewSequenceEncoder(minSchedulingOffsetPreferenceExtR17PreferredK2R17Constraints)
			if err := minSchedulingOffsetPreferenceExtR17PreferredK2R17Seq.EncodePreamble([]bool{c.PreferredK2_SCS_480kHz_r17 != nil, c.PreferredK2_SCS_960kHz_r17 != nil}); err != nil {
				return err
			}
			if c.PreferredK2_SCS_480kHz_r17 != nil {
				if err := e.EncodeEnumerated((*c.PreferredK2_SCS_480kHz_r17), minSchedulingOffsetPreferenceExtR17PreferredK2R17PreferredK2SCS480kHzR17Constraints); err != nil {
					return err
				}
			}
			if c.PreferredK2_SCS_960kHz_r17 != nil {
				if err := e.EncodeEnumerated((*c.PreferredK2_SCS_960kHz_r17), minSchedulingOffsetPreferenceExtR17PreferredK2R17PreferredK2SCS960kHzR17Constraints); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *MinSchedulingOffsetPreferenceExt_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(minSchedulingOffsetPreferenceExtR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. preferredK0-r17: sequence
	{
		if seq.IsComponentPresent(0) {
			ie.PreferredK0_r17 = &struct {
				PreferredK0_SCS_480kHz_r17 *int64
				PreferredK0_SCS_960kHz_r17 *int64
			}{}
			c := ie.PreferredK0_r17
			minSchedulingOffsetPreferenceExtR17PreferredK0R17Seq := d.NewSequenceDecoder(minSchedulingOffsetPreferenceExtR17PreferredK0R17Constraints)
			if err := minSchedulingOffsetPreferenceExtR17PreferredK0R17Seq.DecodePreamble(); err != nil {
				return err
			}
			if minSchedulingOffsetPreferenceExtR17PreferredK0R17Seq.IsComponentPresent(0) {
				c.PreferredK0_SCS_480kHz_r17 = new(int64)
				v, err := d.DecodeEnumerated(minSchedulingOffsetPreferenceExtR17PreferredK0R17PreferredK0SCS480kHzR17Constraints)
				if err != nil {
					return err
				}
				(*c.PreferredK0_SCS_480kHz_r17) = v
			}
			if minSchedulingOffsetPreferenceExtR17PreferredK0R17Seq.IsComponentPresent(1) {
				c.PreferredK0_SCS_960kHz_r17 = new(int64)
				v, err := d.DecodeEnumerated(minSchedulingOffsetPreferenceExtR17PreferredK0R17PreferredK0SCS960kHzR17Constraints)
				if err != nil {
					return err
				}
				(*c.PreferredK0_SCS_960kHz_r17) = v
			}
		}
	}

	// 3. preferredK2-r17: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.PreferredK2_r17 = &struct {
				PreferredK2_SCS_480kHz_r17 *int64
				PreferredK2_SCS_960kHz_r17 *int64
			}{}
			c := ie.PreferredK2_r17
			minSchedulingOffsetPreferenceExtR17PreferredK2R17Seq := d.NewSequenceDecoder(minSchedulingOffsetPreferenceExtR17PreferredK2R17Constraints)
			if err := minSchedulingOffsetPreferenceExtR17PreferredK2R17Seq.DecodePreamble(); err != nil {
				return err
			}
			if minSchedulingOffsetPreferenceExtR17PreferredK2R17Seq.IsComponentPresent(0) {
				c.PreferredK2_SCS_480kHz_r17 = new(int64)
				v, err := d.DecodeEnumerated(minSchedulingOffsetPreferenceExtR17PreferredK2R17PreferredK2SCS480kHzR17Constraints)
				if err != nil {
					return err
				}
				(*c.PreferredK2_SCS_480kHz_r17) = v
			}
			if minSchedulingOffsetPreferenceExtR17PreferredK2R17Seq.IsComponentPresent(1) {
				c.PreferredK2_SCS_960kHz_r17 = new(int64)
				v, err := d.DecodeEnumerated(minSchedulingOffsetPreferenceExtR17PreferredK2R17PreferredK2SCS960kHzR17Constraints)
				if err != nil {
					return err
				}
				(*c.PreferredK2_SCS_960kHz_r17) = v
			}
		}
	}

	return nil
}
