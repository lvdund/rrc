// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: ZP-CSI-RS-ResourceSetId (line 16466).
// ZP-CSI-RS-ResourceSetId ::=                     INTEGER (0..maxNrofZP-CSI-RS-ResourceSets-1)

var zPCSIRSResourceSetIdConstraints = per.Constrained(0, common.MaxNrofZP_CSI_RS_ResourceSets_1)

type ZP_CSI_RS_ResourceSetId struct {
	Value int64
}

func (v *ZP_CSI_RS_ResourceSetId) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, zPCSIRSResourceSetIdConstraints)
}

func (v *ZP_CSI_RS_ResourceSetId) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(zPCSIRSResourceSetIdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
