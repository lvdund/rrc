// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CA-BandwidthClassNR-r17 (line 17244).
// CA-BandwidthClassNR-r17 ::=         ENUMERATED {r, s, t, u, ...}

const (
	CA_BandwidthClassNR_r17_R = 0
	CA_BandwidthClassNR_r17_S = 1
	CA_BandwidthClassNR_r17_T = 2
	CA_BandwidthClassNR_r17_U = 3
)

var cABandwidthClassNRR17Constraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{CA_BandwidthClassNR_r17_R, CA_BandwidthClassNR_r17_S, CA_BandwidthClassNR_r17_T, CA_BandwidthClassNR_r17_U},
}

type CA_BandwidthClassNR_r17 struct {
	Value int64
}

func (v *CA_BandwidthClassNR_r17) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, cABandwidthClassNRR17Constraints)
}

func (v *CA_BandwidthClassNR_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(cABandwidthClassNRR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
