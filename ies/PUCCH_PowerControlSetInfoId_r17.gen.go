// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PUCCH-PowerControlSetInfoId-r17 (line 12282).
// PUCCH-PowerControlSetInfoId-r17 ::=     INTEGER (1.. maxNrofPowerControlSetInfos-r17)

var pUCCHPowerControlSetInfoIdR17Constraints = per.Constrained(1, common.MaxNrofPowerControlSetInfos_r17)

type PUCCH_PowerControlSetInfoId_r17 struct {
	Value int64
}

func (v *PUCCH_PowerControlSetInfoId_r17) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, pUCCHPowerControlSetInfoIdR17Constraints)
}

func (v *PUCCH_PowerControlSetInfoId_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(pUCCHPowerControlSetInfoIdR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
