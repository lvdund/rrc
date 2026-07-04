// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: AdditionalPCIIndex-r17 (line 4938).
// AdditionalPCIIndex-r17  ::=  INTEGER(1..maxNrofAdditionalPCI-r17)

var additionalPCIIndexR17Constraints = per.Constrained(1, common.MaxNrofAdditionalPCI_r17)

type AdditionalPCIIndex_r17 struct {
	Value int64
}

func (v *AdditionalPCIIndex_r17) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, additionalPCIIndexR17Constraints)
}

func (v *AdditionalPCIIndex_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(additionalPCIIndexR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
