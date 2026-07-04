// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: NZP-CSI-RS-ResourceSetId (line 10893).
// NZP-CSI-RS-ResourceSetId ::=        INTEGER (0..maxNrofNZP-CSI-RS-ResourceSets-1)

var nZPCSIRSResourceSetIdConstraints = per.Constrained(0, common.MaxNrofNZP_CSI_RS_ResourceSets_1)

type NZP_CSI_RS_ResourceSetId struct {
	Value int64
}

func (v *NZP_CSI_RS_ResourceSetId) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, nZPCSIRSResourceSetIdConstraints)
}

func (v *NZP_CSI_RS_ResourceSetId) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(nZPCSIRSResourceSetIdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
