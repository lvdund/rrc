// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SecurityCellSetId-r18 (line 6527).
// SecurityCellSetId-r18 ::= INTEGER (1.. maxSecurityCellSet-r18)

var securityCellSetIdR18Constraints = per.Constrained(1, common.MaxSecurityCellSet_r18)

type SecurityCellSetId_r18 struct {
	Value int64
}

func (v *SecurityCellSetId_r18) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, securityCellSetIdR18Constraints)
}

func (v *SecurityCellSetId_r18) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(securityCellSetIdR18Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
