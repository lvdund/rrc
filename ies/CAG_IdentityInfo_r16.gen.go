// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CAG-IdentityInfo-r16 (line 10571).

var cAGIdentityInfoR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "cag-Identity-r16"},
		{Name: "manualCAGselectionAllowed-r16", Optional: true},
	},
}

var cAGIdentityInfoR16CagIdentityR16Constraints = per.FixedSize(32)

const (
	CAG_IdentityInfo_r16_ManualCAGselectionAllowed_r16_True = 0
)

var cAGIdentityInfoR16ManualCAGselectionAllowedR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CAG_IdentityInfo_r16_ManualCAGselectionAllowed_r16_True},
}

type CAG_IdentityInfo_r16 struct {
	Cag_Identity_r16              per.BitString
	ManualCAGselectionAllowed_r16 *int64
}

func (ie *CAG_IdentityInfo_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cAGIdentityInfoR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ManualCAGselectionAllowed_r16 != nil}); err != nil {
		return err
	}

	// 2. cag-Identity-r16: bit-string
	{
		if err := e.EncodeBitString(ie.Cag_Identity_r16, cAGIdentityInfoR16CagIdentityR16Constraints); err != nil {
			return err
		}
	}

	// 3. manualCAGselectionAllowed-r16: enumerated
	{
		if ie.ManualCAGselectionAllowed_r16 != nil {
			if err := e.EncodeEnumerated(*ie.ManualCAGselectionAllowed_r16, cAGIdentityInfoR16ManualCAGselectionAllowedR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CAG_IdentityInfo_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cAGIdentityInfoR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. cag-Identity-r16: bit-string
	{
		v0, err := d.DecodeBitString(cAGIdentityInfoR16CagIdentityR16Constraints)
		if err != nil {
			return err
		}
		ie.Cag_Identity_r16 = v0
	}

	// 3. manualCAGselectionAllowed-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(cAGIdentityInfoR16ManualCAGselectionAllowedR16Constraints)
			if err != nil {
				return err
			}
			ie.ManualCAGselectionAllowed_r16 = &idx
		}
	}

	return nil
}
