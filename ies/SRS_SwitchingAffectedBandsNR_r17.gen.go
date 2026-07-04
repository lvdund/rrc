// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SRS-SwitchingAffectedBandsNR-r17 (line 17127).
// SRS-SwitchingAffectedBandsNR-r17 ::= BIT STRING (SIZE (1..maxSimultaneousBands))

var sRSSwitchingAffectedBandsNRR17Constraints = per.SizeRange(1, common.MaxSimultaneousBands)

type SRS_SwitchingAffectedBandsNR_r17 struct {
	Value per.BitString
}

func (v *SRS_SwitchingAffectedBandsNR_r17) Encode(e *per.Encoder) error {
	return e.EncodeBitString(v.Value, sRSSwitchingAffectedBandsNRR17Constraints)
}

func (v *SRS_SwitchingAffectedBandsNR_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeBitString(sRSSwitchingAffectedBandsNRR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
