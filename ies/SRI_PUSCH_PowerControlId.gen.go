// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SRI-PUSCH-PowerControlId (line 12626).
// SRI-PUSCH-PowerControlId ::=        INTEGER (0..maxNrofSRI-PUSCH-Mappings-1)

var sRIPUSCHPowerControlIdConstraints = per.Constrained(0, common.MaxNrofSRI_PUSCH_Mappings_1)

type SRI_PUSCH_PowerControlId struct {
	Value int64
}

func (v *SRI_PUSCH_PowerControlId) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, sRIPUSCHPowerControlIdConstraints)
}

func (v *SRI_PUSCH_PowerControlId) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(sRIPUSCHPowerControlIdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
