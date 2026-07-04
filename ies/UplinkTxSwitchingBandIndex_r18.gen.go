// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: UplinkTxSwitchingBandIndex-r18 (line 5841).
// UplinkTxSwitchingBandIndex-r18::=  INTEGER (1..maxSimultaneousBands)

var uplinkTxSwitchingBandIndexR18Constraints = per.Constrained(1, common.MaxSimultaneousBands)

type UplinkTxSwitchingBandIndex_r18 struct {
	Value int64
}

func (v *UplinkTxSwitchingBandIndex_r18) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, uplinkTxSwitchingBandIndexR18Constraints)
}

func (v *UplinkTxSwitchingBandIndex_r18) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(uplinkTxSwitchingBandIndexR18Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
