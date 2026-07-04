// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: NZP-CSI-RS-ResourceId (line 10842).
// NZP-CSI-RS-ResourceId ::=           INTEGER (0..maxNrofNZP-CSI-RS-Resources-1)

var nZPCSIRSResourceIdConstraints = per.Constrained(0, common.MaxNrofNZP_CSI_RS_Resources_1)

type NZP_CSI_RS_ResourceId struct {
	Value int64
}

func (v *NZP_CSI_RS_ResourceId) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, nZPCSIRSResourceIdConstraints)
}

func (v *NZP_CSI_RS_ResourceId) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(nZPCSIRSResourceIdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
