// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MinSchedulingOffsetPreference-r16 (line 2582).

var minSchedulingOffsetPreferenceR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "preferredK0-r16", Optional: true},
		{Name: "preferredK2-r16", Optional: true},
	},
}

var minSchedulingOffsetPreferenceR16PreferredK0R16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "preferredK0-SCS-15kHz-r16", Optional: true},
		{Name: "preferredK0-SCS-30kHz-r16", Optional: true},
		{Name: "preferredK0-SCS-60kHz-r16", Optional: true},
		{Name: "preferredK0-SCS-120kHz-r16", Optional: true},
	},
}

const (
	MinSchedulingOffsetPreference_r16_PreferredK0_r16_PreferredK0_SCS_15kHz_r16_Sl1 = 0
	MinSchedulingOffsetPreference_r16_PreferredK0_r16_PreferredK0_SCS_15kHz_r16_Sl2 = 1
	MinSchedulingOffsetPreference_r16_PreferredK0_r16_PreferredK0_SCS_15kHz_r16_Sl4 = 2
	MinSchedulingOffsetPreference_r16_PreferredK0_r16_PreferredK0_SCS_15kHz_r16_Sl6 = 3
)

var minSchedulingOffsetPreferenceR16PreferredK0R16PreferredK0SCS15kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MinSchedulingOffsetPreference_r16_PreferredK0_r16_PreferredK0_SCS_15kHz_r16_Sl1, MinSchedulingOffsetPreference_r16_PreferredK0_r16_PreferredK0_SCS_15kHz_r16_Sl2, MinSchedulingOffsetPreference_r16_PreferredK0_r16_PreferredK0_SCS_15kHz_r16_Sl4, MinSchedulingOffsetPreference_r16_PreferredK0_r16_PreferredK0_SCS_15kHz_r16_Sl6},
}

const (
	MinSchedulingOffsetPreference_r16_PreferredK0_r16_PreferredK0_SCS_30kHz_r16_Sl1 = 0
	MinSchedulingOffsetPreference_r16_PreferredK0_r16_PreferredK0_SCS_30kHz_r16_Sl2 = 1
	MinSchedulingOffsetPreference_r16_PreferredK0_r16_PreferredK0_SCS_30kHz_r16_Sl4 = 2
	MinSchedulingOffsetPreference_r16_PreferredK0_r16_PreferredK0_SCS_30kHz_r16_Sl6 = 3
)

var minSchedulingOffsetPreferenceR16PreferredK0R16PreferredK0SCS30kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MinSchedulingOffsetPreference_r16_PreferredK0_r16_PreferredK0_SCS_30kHz_r16_Sl1, MinSchedulingOffsetPreference_r16_PreferredK0_r16_PreferredK0_SCS_30kHz_r16_Sl2, MinSchedulingOffsetPreference_r16_PreferredK0_r16_PreferredK0_SCS_30kHz_r16_Sl4, MinSchedulingOffsetPreference_r16_PreferredK0_r16_PreferredK0_SCS_30kHz_r16_Sl6},
}

const (
	MinSchedulingOffsetPreference_r16_PreferredK0_r16_PreferredK0_SCS_60kHz_r16_Sl2  = 0
	MinSchedulingOffsetPreference_r16_PreferredK0_r16_PreferredK0_SCS_60kHz_r16_Sl4  = 1
	MinSchedulingOffsetPreference_r16_PreferredK0_r16_PreferredK0_SCS_60kHz_r16_Sl8  = 2
	MinSchedulingOffsetPreference_r16_PreferredK0_r16_PreferredK0_SCS_60kHz_r16_Sl12 = 3
)

var minSchedulingOffsetPreferenceR16PreferredK0R16PreferredK0SCS60kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MinSchedulingOffsetPreference_r16_PreferredK0_r16_PreferredK0_SCS_60kHz_r16_Sl2, MinSchedulingOffsetPreference_r16_PreferredK0_r16_PreferredK0_SCS_60kHz_r16_Sl4, MinSchedulingOffsetPreference_r16_PreferredK0_r16_PreferredK0_SCS_60kHz_r16_Sl8, MinSchedulingOffsetPreference_r16_PreferredK0_r16_PreferredK0_SCS_60kHz_r16_Sl12},
}

const (
	MinSchedulingOffsetPreference_r16_PreferredK0_r16_PreferredK0_SCS_120kHz_r16_Sl2  = 0
	MinSchedulingOffsetPreference_r16_PreferredK0_r16_PreferredK0_SCS_120kHz_r16_Sl4  = 1
	MinSchedulingOffsetPreference_r16_PreferredK0_r16_PreferredK0_SCS_120kHz_r16_Sl8  = 2
	MinSchedulingOffsetPreference_r16_PreferredK0_r16_PreferredK0_SCS_120kHz_r16_Sl12 = 3
)

var minSchedulingOffsetPreferenceR16PreferredK0R16PreferredK0SCS120kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MinSchedulingOffsetPreference_r16_PreferredK0_r16_PreferredK0_SCS_120kHz_r16_Sl2, MinSchedulingOffsetPreference_r16_PreferredK0_r16_PreferredK0_SCS_120kHz_r16_Sl4, MinSchedulingOffsetPreference_r16_PreferredK0_r16_PreferredK0_SCS_120kHz_r16_Sl8, MinSchedulingOffsetPreference_r16_PreferredK0_r16_PreferredK0_SCS_120kHz_r16_Sl12},
}

var minSchedulingOffsetPreferenceR16PreferredK2R16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "preferredK2-SCS-15kHz-r16", Optional: true},
		{Name: "preferredK2-SCS-30kHz-r16", Optional: true},
		{Name: "preferredK2-SCS-60kHz-r16", Optional: true},
		{Name: "preferredK2-SCS-120kHz-r16", Optional: true},
	},
}

const (
	MinSchedulingOffsetPreference_r16_PreferredK2_r16_PreferredK2_SCS_15kHz_r16_Sl1 = 0
	MinSchedulingOffsetPreference_r16_PreferredK2_r16_PreferredK2_SCS_15kHz_r16_Sl2 = 1
	MinSchedulingOffsetPreference_r16_PreferredK2_r16_PreferredK2_SCS_15kHz_r16_Sl4 = 2
	MinSchedulingOffsetPreference_r16_PreferredK2_r16_PreferredK2_SCS_15kHz_r16_Sl6 = 3
)

var minSchedulingOffsetPreferenceR16PreferredK2R16PreferredK2SCS15kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MinSchedulingOffsetPreference_r16_PreferredK2_r16_PreferredK2_SCS_15kHz_r16_Sl1, MinSchedulingOffsetPreference_r16_PreferredK2_r16_PreferredK2_SCS_15kHz_r16_Sl2, MinSchedulingOffsetPreference_r16_PreferredK2_r16_PreferredK2_SCS_15kHz_r16_Sl4, MinSchedulingOffsetPreference_r16_PreferredK2_r16_PreferredK2_SCS_15kHz_r16_Sl6},
}

const (
	MinSchedulingOffsetPreference_r16_PreferredK2_r16_PreferredK2_SCS_30kHz_r16_Sl1 = 0
	MinSchedulingOffsetPreference_r16_PreferredK2_r16_PreferredK2_SCS_30kHz_r16_Sl2 = 1
	MinSchedulingOffsetPreference_r16_PreferredK2_r16_PreferredK2_SCS_30kHz_r16_Sl4 = 2
	MinSchedulingOffsetPreference_r16_PreferredK2_r16_PreferredK2_SCS_30kHz_r16_Sl6 = 3
)

var minSchedulingOffsetPreferenceR16PreferredK2R16PreferredK2SCS30kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MinSchedulingOffsetPreference_r16_PreferredK2_r16_PreferredK2_SCS_30kHz_r16_Sl1, MinSchedulingOffsetPreference_r16_PreferredK2_r16_PreferredK2_SCS_30kHz_r16_Sl2, MinSchedulingOffsetPreference_r16_PreferredK2_r16_PreferredK2_SCS_30kHz_r16_Sl4, MinSchedulingOffsetPreference_r16_PreferredK2_r16_PreferredK2_SCS_30kHz_r16_Sl6},
}

const (
	MinSchedulingOffsetPreference_r16_PreferredK2_r16_PreferredK2_SCS_60kHz_r16_Sl2  = 0
	MinSchedulingOffsetPreference_r16_PreferredK2_r16_PreferredK2_SCS_60kHz_r16_Sl4  = 1
	MinSchedulingOffsetPreference_r16_PreferredK2_r16_PreferredK2_SCS_60kHz_r16_Sl8  = 2
	MinSchedulingOffsetPreference_r16_PreferredK2_r16_PreferredK2_SCS_60kHz_r16_Sl12 = 3
)

var minSchedulingOffsetPreferenceR16PreferredK2R16PreferredK2SCS60kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MinSchedulingOffsetPreference_r16_PreferredK2_r16_PreferredK2_SCS_60kHz_r16_Sl2, MinSchedulingOffsetPreference_r16_PreferredK2_r16_PreferredK2_SCS_60kHz_r16_Sl4, MinSchedulingOffsetPreference_r16_PreferredK2_r16_PreferredK2_SCS_60kHz_r16_Sl8, MinSchedulingOffsetPreference_r16_PreferredK2_r16_PreferredK2_SCS_60kHz_r16_Sl12},
}

const (
	MinSchedulingOffsetPreference_r16_PreferredK2_r16_PreferredK2_SCS_120kHz_r16_Sl2  = 0
	MinSchedulingOffsetPreference_r16_PreferredK2_r16_PreferredK2_SCS_120kHz_r16_Sl4  = 1
	MinSchedulingOffsetPreference_r16_PreferredK2_r16_PreferredK2_SCS_120kHz_r16_Sl8  = 2
	MinSchedulingOffsetPreference_r16_PreferredK2_r16_PreferredK2_SCS_120kHz_r16_Sl12 = 3
)

var minSchedulingOffsetPreferenceR16PreferredK2R16PreferredK2SCS120kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MinSchedulingOffsetPreference_r16_PreferredK2_r16_PreferredK2_SCS_120kHz_r16_Sl2, MinSchedulingOffsetPreference_r16_PreferredK2_r16_PreferredK2_SCS_120kHz_r16_Sl4, MinSchedulingOffsetPreference_r16_PreferredK2_r16_PreferredK2_SCS_120kHz_r16_Sl8, MinSchedulingOffsetPreference_r16_PreferredK2_r16_PreferredK2_SCS_120kHz_r16_Sl12},
}

type MinSchedulingOffsetPreference_r16 struct {
	PreferredK0_r16 *struct {
		PreferredK0_SCS_15kHz_r16  *int64
		PreferredK0_SCS_30kHz_r16  *int64
		PreferredK0_SCS_60kHz_r16  *int64
		PreferredK0_SCS_120kHz_r16 *int64
	}
	PreferredK2_r16 *struct {
		PreferredK2_SCS_15kHz_r16  *int64
		PreferredK2_SCS_30kHz_r16  *int64
		PreferredK2_SCS_60kHz_r16  *int64
		PreferredK2_SCS_120kHz_r16 *int64
	}
}

func (ie *MinSchedulingOffsetPreference_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(minSchedulingOffsetPreferenceR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.PreferredK0_r16 != nil, ie.PreferredK2_r16 != nil}); err != nil {
		return err
	}

	// 2. preferredK0-r16: sequence
	{
		if ie.PreferredK0_r16 != nil {
			c := ie.PreferredK0_r16
			minSchedulingOffsetPreferenceR16PreferredK0R16Seq := e.NewSequenceEncoder(minSchedulingOffsetPreferenceR16PreferredK0R16Constraints)
			if err := minSchedulingOffsetPreferenceR16PreferredK0R16Seq.EncodePreamble([]bool{c.PreferredK0_SCS_15kHz_r16 != nil, c.PreferredK0_SCS_30kHz_r16 != nil, c.PreferredK0_SCS_60kHz_r16 != nil, c.PreferredK0_SCS_120kHz_r16 != nil}); err != nil {
				return err
			}
			if c.PreferredK0_SCS_15kHz_r16 != nil {
				if err := e.EncodeEnumerated((*c.PreferredK0_SCS_15kHz_r16), minSchedulingOffsetPreferenceR16PreferredK0R16PreferredK0SCS15kHzR16Constraints); err != nil {
					return err
				}
			}
			if c.PreferredK0_SCS_30kHz_r16 != nil {
				if err := e.EncodeEnumerated((*c.PreferredK0_SCS_30kHz_r16), minSchedulingOffsetPreferenceR16PreferredK0R16PreferredK0SCS30kHzR16Constraints); err != nil {
					return err
				}
			}
			if c.PreferredK0_SCS_60kHz_r16 != nil {
				if err := e.EncodeEnumerated((*c.PreferredK0_SCS_60kHz_r16), minSchedulingOffsetPreferenceR16PreferredK0R16PreferredK0SCS60kHzR16Constraints); err != nil {
					return err
				}
			}
			if c.PreferredK0_SCS_120kHz_r16 != nil {
				if err := e.EncodeEnumerated((*c.PreferredK0_SCS_120kHz_r16), minSchedulingOffsetPreferenceR16PreferredK0R16PreferredK0SCS120kHzR16Constraints); err != nil {
					return err
				}
			}
		}
	}

	// 3. preferredK2-r16: sequence
	{
		if ie.PreferredK2_r16 != nil {
			c := ie.PreferredK2_r16
			minSchedulingOffsetPreferenceR16PreferredK2R16Seq := e.NewSequenceEncoder(minSchedulingOffsetPreferenceR16PreferredK2R16Constraints)
			if err := minSchedulingOffsetPreferenceR16PreferredK2R16Seq.EncodePreamble([]bool{c.PreferredK2_SCS_15kHz_r16 != nil, c.PreferredK2_SCS_30kHz_r16 != nil, c.PreferredK2_SCS_60kHz_r16 != nil, c.PreferredK2_SCS_120kHz_r16 != nil}); err != nil {
				return err
			}
			if c.PreferredK2_SCS_15kHz_r16 != nil {
				if err := e.EncodeEnumerated((*c.PreferredK2_SCS_15kHz_r16), minSchedulingOffsetPreferenceR16PreferredK2R16PreferredK2SCS15kHzR16Constraints); err != nil {
					return err
				}
			}
			if c.PreferredK2_SCS_30kHz_r16 != nil {
				if err := e.EncodeEnumerated((*c.PreferredK2_SCS_30kHz_r16), minSchedulingOffsetPreferenceR16PreferredK2R16PreferredK2SCS30kHzR16Constraints); err != nil {
					return err
				}
			}
			if c.PreferredK2_SCS_60kHz_r16 != nil {
				if err := e.EncodeEnumerated((*c.PreferredK2_SCS_60kHz_r16), minSchedulingOffsetPreferenceR16PreferredK2R16PreferredK2SCS60kHzR16Constraints); err != nil {
					return err
				}
			}
			if c.PreferredK2_SCS_120kHz_r16 != nil {
				if err := e.EncodeEnumerated((*c.PreferredK2_SCS_120kHz_r16), minSchedulingOffsetPreferenceR16PreferredK2R16PreferredK2SCS120kHzR16Constraints); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *MinSchedulingOffsetPreference_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(minSchedulingOffsetPreferenceR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. preferredK0-r16: sequence
	{
		if seq.IsComponentPresent(0) {
			ie.PreferredK0_r16 = &struct {
				PreferredK0_SCS_15kHz_r16  *int64
				PreferredK0_SCS_30kHz_r16  *int64
				PreferredK0_SCS_60kHz_r16  *int64
				PreferredK0_SCS_120kHz_r16 *int64
			}{}
			c := ie.PreferredK0_r16
			minSchedulingOffsetPreferenceR16PreferredK0R16Seq := d.NewSequenceDecoder(minSchedulingOffsetPreferenceR16PreferredK0R16Constraints)
			if err := minSchedulingOffsetPreferenceR16PreferredK0R16Seq.DecodePreamble(); err != nil {
				return err
			}
			if minSchedulingOffsetPreferenceR16PreferredK0R16Seq.IsComponentPresent(0) {
				c.PreferredK0_SCS_15kHz_r16 = new(int64)
				v, err := d.DecodeEnumerated(minSchedulingOffsetPreferenceR16PreferredK0R16PreferredK0SCS15kHzR16Constraints)
				if err != nil {
					return err
				}
				(*c.PreferredK0_SCS_15kHz_r16) = v
			}
			if minSchedulingOffsetPreferenceR16PreferredK0R16Seq.IsComponentPresent(1) {
				c.PreferredK0_SCS_30kHz_r16 = new(int64)
				v, err := d.DecodeEnumerated(minSchedulingOffsetPreferenceR16PreferredK0R16PreferredK0SCS30kHzR16Constraints)
				if err != nil {
					return err
				}
				(*c.PreferredK0_SCS_30kHz_r16) = v
			}
			if minSchedulingOffsetPreferenceR16PreferredK0R16Seq.IsComponentPresent(2) {
				c.PreferredK0_SCS_60kHz_r16 = new(int64)
				v, err := d.DecodeEnumerated(minSchedulingOffsetPreferenceR16PreferredK0R16PreferredK0SCS60kHzR16Constraints)
				if err != nil {
					return err
				}
				(*c.PreferredK0_SCS_60kHz_r16) = v
			}
			if minSchedulingOffsetPreferenceR16PreferredK0R16Seq.IsComponentPresent(3) {
				c.PreferredK0_SCS_120kHz_r16 = new(int64)
				v, err := d.DecodeEnumerated(minSchedulingOffsetPreferenceR16PreferredK0R16PreferredK0SCS120kHzR16Constraints)
				if err != nil {
					return err
				}
				(*c.PreferredK0_SCS_120kHz_r16) = v
			}
		}
	}

	// 3. preferredK2-r16: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.PreferredK2_r16 = &struct {
				PreferredK2_SCS_15kHz_r16  *int64
				PreferredK2_SCS_30kHz_r16  *int64
				PreferredK2_SCS_60kHz_r16  *int64
				PreferredK2_SCS_120kHz_r16 *int64
			}{}
			c := ie.PreferredK2_r16
			minSchedulingOffsetPreferenceR16PreferredK2R16Seq := d.NewSequenceDecoder(minSchedulingOffsetPreferenceR16PreferredK2R16Constraints)
			if err := minSchedulingOffsetPreferenceR16PreferredK2R16Seq.DecodePreamble(); err != nil {
				return err
			}
			if minSchedulingOffsetPreferenceR16PreferredK2R16Seq.IsComponentPresent(0) {
				c.PreferredK2_SCS_15kHz_r16 = new(int64)
				v, err := d.DecodeEnumerated(minSchedulingOffsetPreferenceR16PreferredK2R16PreferredK2SCS15kHzR16Constraints)
				if err != nil {
					return err
				}
				(*c.PreferredK2_SCS_15kHz_r16) = v
			}
			if minSchedulingOffsetPreferenceR16PreferredK2R16Seq.IsComponentPresent(1) {
				c.PreferredK2_SCS_30kHz_r16 = new(int64)
				v, err := d.DecodeEnumerated(minSchedulingOffsetPreferenceR16PreferredK2R16PreferredK2SCS30kHzR16Constraints)
				if err != nil {
					return err
				}
				(*c.PreferredK2_SCS_30kHz_r16) = v
			}
			if minSchedulingOffsetPreferenceR16PreferredK2R16Seq.IsComponentPresent(2) {
				c.PreferredK2_SCS_60kHz_r16 = new(int64)
				v, err := d.DecodeEnumerated(minSchedulingOffsetPreferenceR16PreferredK2R16PreferredK2SCS60kHzR16Constraints)
				if err != nil {
					return err
				}
				(*c.PreferredK2_SCS_60kHz_r16) = v
			}
			if minSchedulingOffsetPreferenceR16PreferredK2R16Seq.IsComponentPresent(3) {
				c.PreferredK2_SCS_120kHz_r16 = new(int64)
				v, err := d.DecodeEnumerated(minSchedulingOffsetPreferenceR16PreferredK2R16PreferredK2SCS120kHzR16Constraints)
				if err != nil {
					return err
				}
				(*c.PreferredK2_SCS_120kHz_r16) = v
			}
		}
	}

	return nil
}
