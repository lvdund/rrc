// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PRB-Id (line 11927).
// PRB-Id ::=                          INTEGER (0..maxNrofPhysicalResourceBlocks-1)

var pRBIdConstraints = per.Constrained(0, common.MaxNrofPhysicalResourceBlocks_1)

type PRB_Id struct {
	Value int64
}

func (v *PRB_Id) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, pRBIdConstraints)
}

func (v *PRB_Id) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(pRBIdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
