// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CSI-ResourceConfigId (line 7492).
// CSI-ResourceConfigId ::=            INTEGER (0..maxNrofCSI-ResourceConfigurations-1)

var cSIResourceConfigIdConstraints = per.Constrained(0, common.MaxNrofCSI_ResourceConfigurations_1)

type CSI_ResourceConfigId struct {
	Value int64
}

func (v *CSI_ResourceConfigId) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, cSIResourceConfigIdConstraints)
}

func (v *CSI_ResourceConfigId) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(cSIResourceConfigIdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
