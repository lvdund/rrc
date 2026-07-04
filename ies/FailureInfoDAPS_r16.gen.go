// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FailureInfoDAPS-r16 (line 421).

var failureInfoDAPSR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "failureType-r16"},
	},
}

const (
	FailureInfoDAPS_r16_FailureType_r16_Daps_Failure = 0
	FailureInfoDAPS_r16_FailureType_r16_Spare3       = 1
	FailureInfoDAPS_r16_FailureType_r16_Spare2       = 2
	FailureInfoDAPS_r16_FailureType_r16_Spare1       = 3
)

var failureInfoDAPSR16FailureTypeR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FailureInfoDAPS_r16_FailureType_r16_Daps_Failure, FailureInfoDAPS_r16_FailureType_r16_Spare3, FailureInfoDAPS_r16_FailureType_r16_Spare2, FailureInfoDAPS_r16_FailureType_r16_Spare1},
}

type FailureInfoDAPS_r16 struct {
	FailureType_r16 int64
}

func (ie *FailureInfoDAPS_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(failureInfoDAPSR16Constraints)
	_ = seq

	// 1. failureType-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.FailureType_r16, failureInfoDAPSR16FailureTypeR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *FailureInfoDAPS_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(failureInfoDAPSR16Constraints)
	_ = seq

	// 1. failureType-r16: enumerated
	{
		v0, err := d.DecodeEnumerated(failureInfoDAPSR16FailureTypeR16Constraints)
		if err != nil {
			return err
		}
		ie.FailureType_r16 = v0
	}

	return nil
}
