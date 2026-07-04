// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CSI-IM-ResourceSetId (line 6973).
// CSI-IM-ResourceSetId ::=            INTEGER (0..maxNrofCSI-IM-ResourceSets-1)

var cSIIMResourceSetIdConstraints = per.Constrained(0, common.MaxNrofCSI_IM_ResourceSets_1)

type CSI_IM_ResourceSetId struct {
	Value int64
}

func (v *CSI_IM_ResourceSetId) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, cSIIMResourceSetIdConstraints)
}

func (v *CSI_IM_ResourceSetId) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(cSIIMResourceSetIdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
