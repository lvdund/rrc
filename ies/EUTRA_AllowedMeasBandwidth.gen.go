// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: EUTRA-AllowedMeasBandwidth (line 26134).
// EUTRA-AllowedMeasBandwidth ::=              ENUMERATED {mbw6, mbw15, mbw25, mbw50, mbw75, mbw100}

const (
	EUTRA_AllowedMeasBandwidth_Mbw6   = 0
	EUTRA_AllowedMeasBandwidth_Mbw15  = 1
	EUTRA_AllowedMeasBandwidth_Mbw25  = 2
	EUTRA_AllowedMeasBandwidth_Mbw50  = 3
	EUTRA_AllowedMeasBandwidth_Mbw75  = 4
	EUTRA_AllowedMeasBandwidth_Mbw100 = 5
)

var eUTRAAllowedMeasBandwidthConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{EUTRA_AllowedMeasBandwidth_Mbw6, EUTRA_AllowedMeasBandwidth_Mbw15, EUTRA_AllowedMeasBandwidth_Mbw25, EUTRA_AllowedMeasBandwidth_Mbw50, EUTRA_AllowedMeasBandwidth_Mbw75, EUTRA_AllowedMeasBandwidth_Mbw100},
}

type EUTRA_AllowedMeasBandwidth struct {
	Value int64
}

func (v *EUTRA_AllowedMeasBandwidth) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, eUTRAAllowedMeasBandwidthConstraints)
}

func (v *EUTRA_AllowedMeasBandwidth) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(eUTRAAllowedMeasBandwidthConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
