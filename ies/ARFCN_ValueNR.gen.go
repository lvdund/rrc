// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: ARFCN-ValueNR (line 5024).
// ARFCN-ValueNR ::=               INTEGER (0..maxNARFCN)

var aRFCNValueNRConstraints = per.Constrained(0, common.MaxNARFCN)

type ARFCN_ValueNR struct {
	Value int64
}

func (v *ARFCN_ValueNR) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, aRFCNValueNRConstraints)
}

func (v *ARFCN_ValueNR) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(aRFCNValueNRConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
