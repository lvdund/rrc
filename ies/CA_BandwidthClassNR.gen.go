// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CA-BandwidthClassNR (line 17242).
// CA-BandwidthClassNR ::=             ENUMERATED {a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, ...,r2-v1730, r3-v1730, r4-v1730, r5-v1730, r6-v1730, r7-v1730, r8-v1730, r9-v1730, r10-v1730, r11-v1730, r12-v1730,v-v1770, w-v1770 }

const (
	CA_BandwidthClassNR_A = 0
	CA_BandwidthClassNR_B = 1
	CA_BandwidthClassNR_C = 2
	CA_BandwidthClassNR_D = 3
	CA_BandwidthClassNR_E = 4
	CA_BandwidthClassNR_F = 5
	CA_BandwidthClassNR_G = 6
	CA_BandwidthClassNR_H = 7
	CA_BandwidthClassNR_I = 8
	CA_BandwidthClassNR_J = 9
	CA_BandwidthClassNR_K = 10
	CA_BandwidthClassNR_L = 11
	CA_BandwidthClassNR_M = 12
	CA_BandwidthClassNR_N = 13
	CA_BandwidthClassNR_O = 14
	CA_BandwidthClassNR_P = 15
	CA_BandwidthClassNR_Q = 16
)

var cABandwidthClassNRConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{CA_BandwidthClassNR_A, CA_BandwidthClassNR_B, CA_BandwidthClassNR_C, CA_BandwidthClassNR_D, CA_BandwidthClassNR_E, CA_BandwidthClassNR_F, CA_BandwidthClassNR_G, CA_BandwidthClassNR_H, CA_BandwidthClassNR_I, CA_BandwidthClassNR_J, CA_BandwidthClassNR_K, CA_BandwidthClassNR_L, CA_BandwidthClassNR_M, CA_BandwidthClassNR_N, CA_BandwidthClassNR_O, CA_BandwidthClassNR_P, CA_BandwidthClassNR_Q},
	ExtValues:  []int64{17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29},
}

type CA_BandwidthClassNR struct {
	Value int64
}

func (v *CA_BandwidthClassNR) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, cABandwidthClassNRConstraints)
}

func (v *CA_BandwidthClassNR) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(cABandwidthClassNRConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
