// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CSI-IM-ResourceId (line 6960).
// CSI-IM-ResourceId ::=               INTEGER (0..maxNrofCSI-IM-Resources-1)

var cSIIMResourceIdConstraints = per.Constrained(0, common.MaxNrofCSI_IM_Resources_1)

type CSI_IM_ResourceId struct {
	Value int64
}

func (v *CSI_IM_ResourceId) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, cSIIMResourceIdConstraints)
}

func (v *CSI_IM_ResourceId) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(cSIIMResourceIdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
