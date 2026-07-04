// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SCellActivationRS-ConfigId-r17 (line 14240).
// SCellActivationRS-ConfigId-r17 ::=        INTEGER (1.. maxNrofSCellActRS-r17)

var sCellActivationRSConfigIdR17Constraints = per.Constrained(1, common.MaxNrofSCellActRS_r17)

type SCellActivationRS_ConfigId_r17 struct {
	Value int64
}

func (v *SCellActivationRS_ConfigId_r17) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, sCellActivationRSConfigIdR17Constraints)
}

func (v *SCellActivationRS_ConfigId_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(sCellActivationRSConfigIdR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
