// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SemiStaticChannelAccessConfigUE-r17 (line 14586).

var semiStaticChannelAccessConfigUER17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "periodUE-r17"},
		{Name: "offsetUE-r17"},
	},
}

const (
	SemiStaticChannelAccessConfigUE_r17_PeriodUE_r17_Ms1     = 0
	SemiStaticChannelAccessConfigUE_r17_PeriodUE_r17_Ms2     = 1
	SemiStaticChannelAccessConfigUE_r17_PeriodUE_r17_Ms2dot5 = 2
	SemiStaticChannelAccessConfigUE_r17_PeriodUE_r17_Ms4     = 3
	SemiStaticChannelAccessConfigUE_r17_PeriodUE_r17_Ms5     = 4
	SemiStaticChannelAccessConfigUE_r17_PeriodUE_r17_Ms10    = 5
	SemiStaticChannelAccessConfigUE_r17_PeriodUE_r17_Spare2  = 6
	SemiStaticChannelAccessConfigUE_r17_PeriodUE_r17_Spare1  = 7
)

var semiStaticChannelAccessConfigUER17PeriodUER17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SemiStaticChannelAccessConfigUE_r17_PeriodUE_r17_Ms1, SemiStaticChannelAccessConfigUE_r17_PeriodUE_r17_Ms2, SemiStaticChannelAccessConfigUE_r17_PeriodUE_r17_Ms2dot5, SemiStaticChannelAccessConfigUE_r17_PeriodUE_r17_Ms4, SemiStaticChannelAccessConfigUE_r17_PeriodUE_r17_Ms5, SemiStaticChannelAccessConfigUE_r17_PeriodUE_r17_Ms10, SemiStaticChannelAccessConfigUE_r17_PeriodUE_r17_Spare2, SemiStaticChannelAccessConfigUE_r17_PeriodUE_r17_Spare1},
}

var semiStaticChannelAccessConfigUER17OffsetUER17Constraints = per.Constrained(0, 559)

type SemiStaticChannelAccessConfigUE_r17 struct {
	PeriodUE_r17 int64
	OffsetUE_r17 int64
}

func (ie *SemiStaticChannelAccessConfigUE_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(semiStaticChannelAccessConfigUER17Constraints)
	_ = seq

	// 1. periodUE-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.PeriodUE_r17, semiStaticChannelAccessConfigUER17PeriodUER17Constraints); err != nil {
			return err
		}
	}

	// 2. offsetUE-r17: integer
	{
		if err := e.EncodeInteger(ie.OffsetUE_r17, semiStaticChannelAccessConfigUER17OffsetUER17Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SemiStaticChannelAccessConfigUE_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(semiStaticChannelAccessConfigUER17Constraints)
	_ = seq

	// 1. periodUE-r17: enumerated
	{
		v0, err := d.DecodeEnumerated(semiStaticChannelAccessConfigUER17PeriodUER17Constraints)
		if err != nil {
			return err
		}
		ie.PeriodUE_r17 = v0
	}

	// 2. offsetUE-r17: integer
	{
		v1, err := d.DecodeInteger(semiStaticChannelAccessConfigUER17OffsetUER17Constraints)
		if err != nil {
			return err
		}
		ie.OffsetUE_r17 = v1
	}

	return nil
}
