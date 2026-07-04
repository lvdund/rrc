// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SupportedBandwidth-v1900 (line 25425).
// SupportedBandwidth-v1900 ::=  ENUMERATED {mhz7}

const (
	SupportedBandwidth_v1900_Mhz7 = 0
)

var supportedBandwidthV1900Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SupportedBandwidth_v1900_Mhz7},
}

type SupportedBandwidth_v1900 struct {
	Value int64
}

func (v *SupportedBandwidth_v1900) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, supportedBandwidthV1900Constraints)
}

func (v *SupportedBandwidth_v1900) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(supportedBandwidthV1900Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
