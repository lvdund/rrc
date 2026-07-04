// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-SRAP-ConfigId-r19 (line 28289).
// SL-SRAP-ConfigId-r19 ::=    INTEGER (1.. maxNrofRemoteUE-r17)

var sLSRAPConfigIdR19Constraints = per.Constrained(1, common.MaxNrofRemoteUE_r17)

type SL_SRAP_ConfigId_r19 struct {
	Value int64
}

func (v *SL_SRAP_ConfigId_r19) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, sLSRAPConfigIdR19Constraints)
}

func (v *SL_SRAP_ConfigId_r19) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(sLSRAPConfigIdR19Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
