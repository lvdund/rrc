// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: ServingAdditionalPCIIndex-r17 (line 7623).
// ServingAdditionalPCIIndex-r17  ::=  INTEGER(0..maxNrofAdditionalPCI-r17)

var servingAdditionalPCIIndexR17Constraints = per.Constrained(0, common.MaxNrofAdditionalPCI_r17)

type ServingAdditionalPCIIndex_r17 struct {
	Value int64
}

func (v *ServingAdditionalPCIIndex_r17) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, servingAdditionalPCIIndexR17Constraints)
}

func (v *ServingAdditionalPCIIndex_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(servingAdditionalPCIIndexR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
