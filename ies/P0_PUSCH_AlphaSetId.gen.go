// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: P0-PUSCH-AlphaSetId (line 12590).
// P0-PUSCH-AlphaSetId ::=             INTEGER (0..maxNrofP0-PUSCH-AlphaSets-1)

var p0PUSCHAlphaSetIdConstraints = per.Constrained(0, common.MaxNrofP0_PUSCH_AlphaSets_1)

type P0_PUSCH_AlphaSetId struct {
	Value int64
}

func (v *P0_PUSCH_AlphaSetId) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, p0PUSCHAlphaSetIdConstraints)
}

func (v *P0_PUSCH_AlphaSetId) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(p0PUSCHAlphaSetIdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
