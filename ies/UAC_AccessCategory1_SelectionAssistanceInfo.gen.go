// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UAC-AccessCategory1-SelectionAssistanceInfo (line 2078).
// UAC-AccessCategory1-SelectionAssistanceInfo ::=    ENUMERATED {a, b, c}

const (
	UAC_AccessCategory1_SelectionAssistanceInfo_A = 0
	UAC_AccessCategory1_SelectionAssistanceInfo_B = 1
	UAC_AccessCategory1_SelectionAssistanceInfo_C = 2
)

var uACAccessCategory1SelectionAssistanceInfoConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UAC_AccessCategory1_SelectionAssistanceInfo_A, UAC_AccessCategory1_SelectionAssistanceInfo_B, UAC_AccessCategory1_SelectionAssistanceInfo_C},
}

type UAC_AccessCategory1_SelectionAssistanceInfo struct {
	Value int64
}

func (v *UAC_AccessCategory1_SelectionAssistanceInfo) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, uACAccessCategory1SelectionAssistanceInfoConstraints)
}

func (v *UAC_AccessCategory1_SelectionAssistanceInfo) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(uACAccessCategory1SelectionAssistanceInfoConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
