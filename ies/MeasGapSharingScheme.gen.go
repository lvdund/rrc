// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasGapSharingScheme (line 9227).
// MeasGapSharingScheme::=         ENUMERATED {scheme00, scheme01, scheme10, scheme11}

const (
	MeasGapSharingScheme_Scheme00 = 0
	MeasGapSharingScheme_Scheme01 = 1
	MeasGapSharingScheme_Scheme10 = 2
	MeasGapSharingScheme_Scheme11 = 3
)

var measGapSharingSchemeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasGapSharingScheme_Scheme00, MeasGapSharingScheme_Scheme01, MeasGapSharingScheme_Scheme10, MeasGapSharingScheme_Scheme11},
}

type MeasGapSharingScheme struct {
	Value int64
}

func (v *MeasGapSharingScheme) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, measGapSharingSchemeConstraints)
}

func (v *MeasGapSharingScheme) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(measGapSharingSchemeConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
