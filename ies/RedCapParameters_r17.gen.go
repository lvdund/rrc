// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RedCapParameters-r17 (line 23528).

var redCapParametersR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "supportOfRedCap-r17", Optional: true},
		{Name: "supportOf16DRB-RedCap-r17", Optional: true},
	},
}

const (
	RedCapParameters_r17_SupportOfRedCap_r17_Supported = 0
)

var redCapParametersR17SupportOfRedCapR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RedCapParameters_r17_SupportOfRedCap_r17_Supported},
}

const (
	RedCapParameters_r17_SupportOf16DRB_RedCap_r17_Supported = 0
)

var redCapParametersR17SupportOf16DRBRedCapR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RedCapParameters_r17_SupportOf16DRB_RedCap_r17_Supported},
}

type RedCapParameters_r17 struct {
	SupportOfRedCap_r17       *int64
	SupportOf16DRB_RedCap_r17 *int64
}

func (ie *RedCapParameters_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(redCapParametersR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SupportOfRedCap_r17 != nil, ie.SupportOf16DRB_RedCap_r17 != nil}); err != nil {
		return err
	}

	// 2. supportOfRedCap-r17: enumerated
	{
		if ie.SupportOfRedCap_r17 != nil {
			if err := e.EncodeEnumerated(*ie.SupportOfRedCap_r17, redCapParametersR17SupportOfRedCapR17Constraints); err != nil {
				return err
			}
		}
	}

	// 3. supportOf16DRB-RedCap-r17: enumerated
	{
		if ie.SupportOf16DRB_RedCap_r17 != nil {
			if err := e.EncodeEnumerated(*ie.SupportOf16DRB_RedCap_r17, redCapParametersR17SupportOf16DRBRedCapR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *RedCapParameters_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(redCapParametersR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. supportOfRedCap-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(redCapParametersR17SupportOfRedCapR17Constraints)
			if err != nil {
				return err
			}
			ie.SupportOfRedCap_r17 = &idx
		}
	}

	// 3. supportOf16DRB-RedCap-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(redCapParametersR17SupportOf16DRBRedCapR17Constraints)
			if err != nil {
				return err
			}
			ie.SupportOf16DRB_RedCap_r17 = &idx
		}
	}

	return nil
}
