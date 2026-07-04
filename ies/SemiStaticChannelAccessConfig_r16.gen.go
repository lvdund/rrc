// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SemiStaticChannelAccessConfig-r16 (line 14579).

var semiStaticChannelAccessConfigR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "period-r16"},
	},
}

const (
	SemiStaticChannelAccessConfig_r16_Period_r16_Ms1     = 0
	SemiStaticChannelAccessConfig_r16_Period_r16_Ms2     = 1
	SemiStaticChannelAccessConfig_r16_Period_r16_Ms2dot5 = 2
	SemiStaticChannelAccessConfig_r16_Period_r16_Ms4     = 3
	SemiStaticChannelAccessConfig_r16_Period_r16_Ms5     = 4
	SemiStaticChannelAccessConfig_r16_Period_r16_Ms10    = 5
)

var semiStaticChannelAccessConfigR16PeriodR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SemiStaticChannelAccessConfig_r16_Period_r16_Ms1, SemiStaticChannelAccessConfig_r16_Period_r16_Ms2, SemiStaticChannelAccessConfig_r16_Period_r16_Ms2dot5, SemiStaticChannelAccessConfig_r16_Period_r16_Ms4, SemiStaticChannelAccessConfig_r16_Period_r16_Ms5, SemiStaticChannelAccessConfig_r16_Period_r16_Ms10},
}

type SemiStaticChannelAccessConfig_r16 struct {
	Period_r16 int64
}

func (ie *SemiStaticChannelAccessConfig_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(semiStaticChannelAccessConfigR16Constraints)
	_ = seq

	// 1. period-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.Period_r16, semiStaticChannelAccessConfigR16PeriodR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SemiStaticChannelAccessConfig_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(semiStaticChannelAccessConfigR16Constraints)
	_ = seq

	// 1. period-r16: enumerated
	{
		v0, err := d.DecodeEnumerated(semiStaticChannelAccessConfigR16PeriodR16Constraints)
		if err != nil {
			return err
		}
		ie.Period_r16 = v0
	}

	return nil
}
