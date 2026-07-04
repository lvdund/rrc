// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: OD-SSB-ConfigId-r19 (line 10913).
// OD-SSB-ConfigId-r19   ::= INTEGER (0.. maxNrofOD-SSB-1-r19)

var oDSSBConfigIdR19Constraints = per.Constrained(0, common.MaxNrofOD_SSB_1_r19)

type OD_SSB_ConfigId_r19 struct {
	Value int64
}

func (v *OD_SSB_ConfigId_r19) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, oDSSBConfigIdR19Constraints)
}

func (v *OD_SSB_ConfigId_r19) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(oDSSBConfigIdR19Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
