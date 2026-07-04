// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PUCCH-ResourceId (line 12121).
// PUCCH-ResourceId ::=                    INTEGER (0..maxNrofPUCCH-Resources-1)

var pUCCHResourceIdConstraints = per.Constrained(0, common.MaxNrofPUCCH_Resources_1)

type PUCCH_ResourceId struct {
	Value int64
}

func (v *PUCCH_ResourceId) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, pUCCHResourceIdConstraints)
}

func (v *PUCCH_ResourceId) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(pUCCHResourceIdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
