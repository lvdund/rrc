// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-Failure-r16 (line 2300).

var sLFailureR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-DestinationIdentity-r16"},
		{Name: "sl-Failure-r16"},
	},
}

const (
	SL_Failure_r16_Sl_Failure_r16_Rlf             = 0
	SL_Failure_r16_Sl_Failure_r16_ConfigFailure   = 1
	SL_Failure_r16_Sl_Failure_r16_DrxReject_v1710 = 2
	SL_Failure_r16_Sl_Failure_r16_Spare5          = 3
	SL_Failure_r16_Sl_Failure_r16_Spare4          = 4
	SL_Failure_r16_Sl_Failure_r16_Spare3          = 5
	SL_Failure_r16_Sl_Failure_r16_Spare2          = 6
	SL_Failure_r16_Sl_Failure_r16_Spare1          = 7
)

var sLFailureR16SlFailureR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_Failure_r16_Sl_Failure_r16_Rlf, SL_Failure_r16_Sl_Failure_r16_ConfigFailure, SL_Failure_r16_Sl_Failure_r16_DrxReject_v1710, SL_Failure_r16_Sl_Failure_r16_Spare5, SL_Failure_r16_Sl_Failure_r16_Spare4, SL_Failure_r16_Sl_Failure_r16_Spare3, SL_Failure_r16_Sl_Failure_r16_Spare2, SL_Failure_r16_Sl_Failure_r16_Spare1},
}

type SL_Failure_r16 struct {
	Sl_DestinationIdentity_r16 SL_DestinationIdentity_r16
	Sl_Failure_r16             int64
}

func (ie *SL_Failure_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLFailureR16Constraints)
	_ = seq

	// 1. sl-DestinationIdentity-r16: ref
	{
		if err := ie.Sl_DestinationIdentity_r16.Encode(e); err != nil {
			return err
		}
	}

	// 2. sl-Failure-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.Sl_Failure_r16, sLFailureR16SlFailureR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SL_Failure_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLFailureR16Constraints)
	_ = seq

	// 1. sl-DestinationIdentity-r16: ref
	{
		if err := ie.Sl_DestinationIdentity_r16.Decode(d); err != nil {
			return err
		}
	}

	// 2. sl-Failure-r16: enumerated
	{
		v1, err := d.DecodeEnumerated(sLFailureR16SlFailureR16Constraints)
		if err != nil {
			return err
		}
		ie.Sl_Failure_r16 = v1
	}

	return nil
}
