// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: Altitude-r18 (line 4983).
// Altitude-r18 ::=              INTEGER (minAltitude-r18..maxAltitude-r18)

var altitudeR18Constraints = per.Constrained(common.MinAltitude_r18, common.MaxAltitude_r18)

type Altitude_r18 struct {
	Value int64
}

func (v *Altitude_r18) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, altitudeR18Constraints)
}

func (v *Altitude_r18) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(altitudeR18Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
