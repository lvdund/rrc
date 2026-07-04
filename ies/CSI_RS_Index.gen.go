// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CSI-RS-Index (line 7568).
// CSI-RS-Index ::=                    INTEGER (0..maxNrofCSI-RS-ResourcesRRM-1)

var cSIRSIndexConstraints = per.Constrained(0, common.MaxNrofCSI_RS_ResourcesRRM_1)

type CSI_RS_Index struct {
	Value int64
}

func (v *CSI_RS_Index) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, cSIRSIndexConstraints)
}

func (v *CSI_RS_Index) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(cSIRSIndexConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
