// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RLC-ParametersSidelink-r16 (line 25058).

var rLCParametersSidelinkR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "am-WithLongSN-Sidelink-r16", Optional: true},
		{Name: "um-WithLongSN-Sidelink-r16", Optional: true},
	},
}

const (
	RLC_ParametersSidelink_r16_Am_WithLongSN_Sidelink_r16_Supported = 0
)

var rLCParametersSidelinkR16AmWithLongSNSidelinkR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RLC_ParametersSidelink_r16_Am_WithLongSN_Sidelink_r16_Supported},
}

const (
	RLC_ParametersSidelink_r16_Um_WithLongSN_Sidelink_r16_Supported = 0
)

var rLCParametersSidelinkR16UmWithLongSNSidelinkR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RLC_ParametersSidelink_r16_Um_WithLongSN_Sidelink_r16_Supported},
}

type RLC_ParametersSidelink_r16 struct {
	Am_WithLongSN_Sidelink_r16 *int64
	Um_WithLongSN_Sidelink_r16 *int64
}

func (ie *RLC_ParametersSidelink_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rLCParametersSidelinkR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Am_WithLongSN_Sidelink_r16 != nil, ie.Um_WithLongSN_Sidelink_r16 != nil}); err != nil {
		return err
	}

	// 3. am-WithLongSN-Sidelink-r16: enumerated
	{
		if ie.Am_WithLongSN_Sidelink_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Am_WithLongSN_Sidelink_r16, rLCParametersSidelinkR16AmWithLongSNSidelinkR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. um-WithLongSN-Sidelink-r16: enumerated
	{
		if ie.Um_WithLongSN_Sidelink_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Um_WithLongSN_Sidelink_r16, rLCParametersSidelinkR16UmWithLongSNSidelinkR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *RLC_ParametersSidelink_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rLCParametersSidelinkR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. am-WithLongSN-Sidelink-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(rLCParametersSidelinkR16AmWithLongSNSidelinkR16Constraints)
			if err != nil {
				return err
			}
			ie.Am_WithLongSN_Sidelink_r16 = &idx
		}
	}

	// 4. um-WithLongSN-Sidelink-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(rLCParametersSidelinkR16UmWithLongSNSidelinkR16Constraints)
			if err != nil {
				return err
			}
			ie.Um_WithLongSN_Sidelink_r16 = &idx
		}
	}

	return nil
}
