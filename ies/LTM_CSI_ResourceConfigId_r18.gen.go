// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: LTM-CSI-ResourceConfigId-r18 (line 8946).
// LTM-CSI-ResourceConfigId-r18 ::=            INTEGER (0..maxNrofLTM-CSI-ResourceConfigurations-1-r18)

var lTMCSIResourceConfigIdR18Constraints = per.Constrained(0, common.MaxNrofLTM_CSI_ResourceConfigurations_1_r18)

type LTM_CSI_ResourceConfigId_r18 struct {
	Value int64
}

func (v *LTM_CSI_ResourceConfigId_r18) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, lTMCSIResourceConfigIdR18Constraints)
}

func (v *LTM_CSI_ResourceConfigId_r18) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(lTMCSIResourceConfigIdR18Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
