// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PUCCH-ResourceSetId (line 12070).
// PUCCH-ResourceSetId ::=                 INTEGER (0..maxNrofPUCCH-ResourceSets-1)

var pUCCHResourceSetIdConstraints = per.Constrained(0, common.MaxNrofPUCCH_ResourceSets_1)

type PUCCH_ResourceSetId struct {
	Value int64
}

func (v *PUCCH_ResourceSetId) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, pUCCHResourceSetIdConstraints)
}

func (v *PUCCH_ResourceSetId) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(pUCCHResourceSetIdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
