// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UAC-AC1-SelectAssistInfo-r16 (line 2080).
// UAC-AC1-SelectAssistInfo-r16 ::=     ENUMERATED {a, b, c, notConfigured}

const (
	UAC_AC1_SelectAssistInfo_r16_A             = 0
	UAC_AC1_SelectAssistInfo_r16_B             = 1
	UAC_AC1_SelectAssistInfo_r16_C             = 2
	UAC_AC1_SelectAssistInfo_r16_NotConfigured = 3
)

var uACAC1SelectAssistInfoR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UAC_AC1_SelectAssistInfo_r16_A, UAC_AC1_SelectAssistInfo_r16_B, UAC_AC1_SelectAssistInfo_r16_C, UAC_AC1_SelectAssistInfo_r16_NotConfigured},
}

type UAC_AC1_SelectAssistInfo_r16 struct {
	Value int64
}

func (v *UAC_AC1_SelectAssistInfo_r16) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, uACAC1SelectAssistInfoR16Constraints)
}

func (v *UAC_AC1_SelectAssistInfo_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(uACAC1SelectAssistInfoR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
