// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: Uplink-powerControlId-r17 (line 16342).
// Uplink-powerControlId-r17 ::= INTEGER(1.. maxUL-TCI-r17)

var uplinkPowerControlIdR17Constraints = per.Constrained(1, common.MaxUL_TCI_r17)

type Uplink_PowerControlId_r17 struct {
	Value int64
}

func (v *Uplink_PowerControlId_r17) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, uplinkPowerControlIdR17Constraints)
}

func (v *Uplink_PowerControlId_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(uplinkPowerControlIdR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
