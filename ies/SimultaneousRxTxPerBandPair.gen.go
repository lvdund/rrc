// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SimultaneousRxTxPerBandPair (line 25323).
// SimultaneousRxTxPerBandPair ::=             BIT STRING (SIZE (3..496))

var simultaneousRxTxPerBandPairConstraints = per.SizeRange(3, 496)

type SimultaneousRxTxPerBandPair struct {
	Value per.BitString
}

func (v *SimultaneousRxTxPerBandPair) Encode(e *per.Encoder) error {
	return e.EncodeBitString(v.Value, simultaneousRxTxPerBandPairConstraints)
}

func (v *SimultaneousRxTxPerBandPair) Decode(d *per.Decoder) error {
	x, err := d.DecodeBitString(simultaneousRxTxPerBandPairConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
