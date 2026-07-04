// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SSB-Index (line 15846).
// SSB-Index ::=                       INTEGER (0..maxNrofSSBs-1)

var sSBIndexConstraints = per.Constrained(0, common.MaxNrofSSBs_1)

type SSB_Index struct {
	Value int64
}

func (v *SSB_Index) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, sSBIndexConstraints)
}

func (v *SSB_Index) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(sSBIndexConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
