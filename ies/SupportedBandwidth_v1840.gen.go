// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SupportedBandwidth-v1840 (line 25423).
// SupportedBandwidth-v1840 ::=  ENUMERATED {mhz3}

const (
	SupportedBandwidth_v1840_Mhz3 = 0
)

var supportedBandwidthV1840Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SupportedBandwidth_v1840_Mhz3},
}

type SupportedBandwidth_v1840 struct {
	Value int64
}

func (v *SupportedBandwidth_v1840) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, supportedBandwidthV1840Constraints)
}

func (v *SupportedBandwidth_v1840) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(supportedBandwidthV1840Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
