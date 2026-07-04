// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SRS-SwitchingSimultaneousBandsNR-r19 (line 17129).
// SRS-SwitchingSimultaneousBandsNR-r19 ::= BIT STRING (SIZE (1..maxSimultaneousBands))

var sRSSwitchingSimultaneousBandsNRR19Constraints = per.SizeRange(1, common.MaxSimultaneousBands)

type SRS_SwitchingSimultaneousBandsNR_r19 struct {
	Value per.BitString
}

func (v *SRS_SwitchingSimultaneousBandsNR_r19) Encode(e *per.Encoder) error {
	return e.EncodeBitString(v.Value, sRSSwitchingSimultaneousBandsNRR19Constraints)
}

func (v *SRS_SwitchingSimultaneousBandsNR_r19) Decode(d *per.Decoder) error {
	x, err := d.DecodeBitString(sRSSwitchingSimultaneousBandsNRR19Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
