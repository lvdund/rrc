// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: IAB-IP-Usage-r16 (line 26232).
// IAB-IP-Usage-r16 ::= ENUMERATED {f1-C, f1-U, non-F1, spare}

const (
	IAB_IP_Usage_r16_F1_C   = 0
	IAB_IP_Usage_r16_F1_U   = 1
	IAB_IP_Usage_r16_Non_F1 = 2
	IAB_IP_Usage_r16_Spare  = 3
)

var iABIPUsageR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{IAB_IP_Usage_r16_F1_C, IAB_IP_Usage_r16_F1_U, IAB_IP_Usage_r16_Non_F1, IAB_IP_Usage_r16_Spare},
}

type IAB_IP_Usage_r16 struct {
	Value int64
}

func (v *IAB_IP_Usage_r16) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, iABIPUsageR16Constraints)
}

func (v *IAB_IP_Usage_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(iABIPUsageR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
