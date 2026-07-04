// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RSRP-ChangeThreshold-r17 (line 1443).
// RSRP-ChangeThreshold-r17 ::= ENUMERATED {dB4, dB6, dB8, dB10, dB14, dB18, dB22, dB26, dB30, dB34, spare6, spare5, spare4, spare3, spare2, spare1}

const (
	RSRP_ChangeThreshold_r17_DB4    = 0
	RSRP_ChangeThreshold_r17_DB6    = 1
	RSRP_ChangeThreshold_r17_DB8    = 2
	RSRP_ChangeThreshold_r17_DB10   = 3
	RSRP_ChangeThreshold_r17_DB14   = 4
	RSRP_ChangeThreshold_r17_DB18   = 5
	RSRP_ChangeThreshold_r17_DB22   = 6
	RSRP_ChangeThreshold_r17_DB26   = 7
	RSRP_ChangeThreshold_r17_DB30   = 8
	RSRP_ChangeThreshold_r17_DB34   = 9
	RSRP_ChangeThreshold_r17_Spare6 = 10
	RSRP_ChangeThreshold_r17_Spare5 = 11
	RSRP_ChangeThreshold_r17_Spare4 = 12
	RSRP_ChangeThreshold_r17_Spare3 = 13
	RSRP_ChangeThreshold_r17_Spare2 = 14
	RSRP_ChangeThreshold_r17_Spare1 = 15
)

var rSRPChangeThresholdR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RSRP_ChangeThreshold_r17_DB4, RSRP_ChangeThreshold_r17_DB6, RSRP_ChangeThreshold_r17_DB8, RSRP_ChangeThreshold_r17_DB10, RSRP_ChangeThreshold_r17_DB14, RSRP_ChangeThreshold_r17_DB18, RSRP_ChangeThreshold_r17_DB22, RSRP_ChangeThreshold_r17_DB26, RSRP_ChangeThreshold_r17_DB30, RSRP_ChangeThreshold_r17_DB34, RSRP_ChangeThreshold_r17_Spare6, RSRP_ChangeThreshold_r17_Spare5, RSRP_ChangeThreshold_r17_Spare4, RSRP_ChangeThreshold_r17_Spare3, RSRP_ChangeThreshold_r17_Spare2, RSRP_ChangeThreshold_r17_Spare1},
}

type RSRP_ChangeThreshold_r17 struct {
	Value int64
}

func (v *RSRP_ChangeThreshold_r17) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, rSRPChangeThresholdR17Constraints)
}

func (v *RSRP_ChangeThreshold_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(rSRPChangeThresholdR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
