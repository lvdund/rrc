// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-RelayIndicationMP-r18 (line 27803).
// SL-RelayIndicationMP-r18 ::=   ENUMERATED {support}

const (
	SL_RelayIndicationMP_r18_Support = 0
)

var sLRelayIndicationMPR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_RelayIndicationMP_r18_Support},
}

type SL_RelayIndicationMP_r18 struct {
	Value int64
}

func (v *SL_RelayIndicationMP_r18) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, sLRelayIndicationMPR18Constraints)
}

func (v *SL_RelayIndicationMP_r18) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(sLRelayIndicationMPR18Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
