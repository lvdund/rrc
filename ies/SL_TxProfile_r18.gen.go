// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-TxProfile-r18 (line 2288).
// SL-TxProfile-r18 ::=                   ENUMERATED {backwardsCompatible, backwardsIncompatible}

const (
	SL_TxProfile_r18_BackwardsCompatible   = 0
	SL_TxProfile_r18_BackwardsIncompatible = 1
)

var sLTxProfileR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_TxProfile_r18_BackwardsCompatible, SL_TxProfile_r18_BackwardsIncompatible},
}

type SL_TxProfile_r18 struct {
	Value int64
}

func (v *SL_TxProfile_r18) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, sLTxProfileR18Constraints)
}

func (v *SL_TxProfile_r18) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(sLTxProfileR18Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
