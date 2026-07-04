// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: ZP-CSI-RS-ResourceId (line 16452).
// ZP-CSI-RS-ResourceId ::=            INTEGER (0..maxNrofZP-CSI-RS-Resources-1)

var zPCSIRSResourceIdConstraints = per.Constrained(0, common.MaxNrofZP_CSI_RS_Resources_1)

type ZP_CSI_RS_ResourceId struct {
	Value int64
}

func (v *ZP_CSI_RS_ResourceId) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, zPCSIRSResourceIdConstraints)
}

func (v *ZP_CSI_RS_ResourceId) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(zPCSIRSResourceIdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
