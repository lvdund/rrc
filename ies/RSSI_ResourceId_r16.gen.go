// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: RSSI-ResourceId-r16 (line 9371).
// RSSI-ResourceId-r16 ::=             INTEGER (0.. maxNrofCLI-RSSI-Resources-1-r16)

var rSSIResourceIdR16Constraints = per.Constrained(0, common.MaxNrofCLI_RSSI_Resources_1_r16)

type RSSI_ResourceId_r16 struct {
	Value int64
}

func (v *RSSI_ResourceId_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, rSSIResourceIdR16Constraints)
}

func (v *RSSI_ResourceId_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(rSSIResourceIdR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
