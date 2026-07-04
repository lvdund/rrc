// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UplinkTxSwitching-r16 (line 14800).

var uplinkTxSwitchingR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "uplinkTxSwitchingPeriodLocation-r16"},
		{Name: "uplinkTxSwitchingCarrier-r16"},
	},
}

const (
	UplinkTxSwitching_r16_UplinkTxSwitchingCarrier_r16_Carrier1 = 0
	UplinkTxSwitching_r16_UplinkTxSwitchingCarrier_r16_Carrier2 = 1
)

var uplinkTxSwitchingR16UplinkTxSwitchingCarrierR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UplinkTxSwitching_r16_UplinkTxSwitchingCarrier_r16_Carrier1, UplinkTxSwitching_r16_UplinkTxSwitchingCarrier_r16_Carrier2},
}

type UplinkTxSwitching_r16 struct {
	UplinkTxSwitchingPeriodLocation_r16 bool
	UplinkTxSwitchingCarrier_r16        int64
}

func (ie *UplinkTxSwitching_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uplinkTxSwitchingR16Constraints)
	_ = seq

	// 1. uplinkTxSwitchingPeriodLocation-r16: boolean
	{
		if err := e.EncodeBoolean(ie.UplinkTxSwitchingPeriodLocation_r16); err != nil {
			return err
		}
	}

	// 2. uplinkTxSwitchingCarrier-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.UplinkTxSwitchingCarrier_r16, uplinkTxSwitchingR16UplinkTxSwitchingCarrierR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *UplinkTxSwitching_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uplinkTxSwitchingR16Constraints)
	_ = seq

	// 1. uplinkTxSwitchingPeriodLocation-r16: boolean
	{
		v0, err := d.DecodeBoolean()
		if err != nil {
			return err
		}
		ie.UplinkTxSwitchingPeriodLocation_r16 = v0
	}

	// 2. uplinkTxSwitchingCarrier-r16: enumerated
	{
		v1, err := d.DecodeEnumerated(uplinkTxSwitchingR16UplinkTxSwitchingCarrierR16Constraints)
		if err != nil {
			return err
		}
		ie.UplinkTxSwitchingCarrier_r16 = v1
	}

	return nil
}
