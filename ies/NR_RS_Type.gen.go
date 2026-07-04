// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NR-RS-Type (line 13863).
// NR-RS-Type ::=                              ENUMERATED {ssb, csi-rs}

const (
	NR_RS_Type_Ssb    = 0
	NR_RS_Type_Csi_Rs = 1
)

var nRRSTypeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{NR_RS_Type_Ssb, NR_RS_Type_Csi_Rs},
}

type NR_RS_Type struct {
	Value int64
}

func (v *NR_RS_Type) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, nRRSTypeConstraints)
}

func (v *NR_RS_Type) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(nRRSTypeConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
