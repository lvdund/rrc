// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: ApplicabilitySetConfigId-r19 (line 4993).
// ApplicabilitySetConfigId-r19 ::=            INTEGER (0..maxNrofApplicabilitySetCSI-Configs-1-r19)

var applicabilitySetConfigIdR19Constraints = per.Constrained(0, common.MaxNrofApplicabilitySetCSI_Configs_1_r19)

type ApplicabilitySetConfigId_r19 struct {
	Value int64
}

func (v *ApplicabilitySetConfigId_r19) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, applicabilitySetConfigIdR19Constraints)
}

func (v *ApplicabilitySetConfigId_r19) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(applicabilitySetConfigIdR19Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
