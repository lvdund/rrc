// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: P0-PUSCH-SetId-r16 (line 12657).
// P0-PUSCH-SetId-r16 ::=              INTEGER (0..maxNrofSRI-PUSCH-Mappings-1)

var p0PUSCHSetIdR16Constraints = per.Constrained(0, common.MaxNrofSRI_PUSCH_Mappings_1)

type P0_PUSCH_SetId_r16 struct {
	Value int64
}

func (v *P0_PUSCH_SetId_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, p0PUSCHSetIdR16Constraints)
}

func (v *P0_PUSCH_SetId_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(p0PUSCHSetIdR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
