// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: AffectedFreqRange-r18 (line 2753).

var affectedFreqRangeR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "centerFreq-r18"},
		{Name: "affectedBandwidth-r18"},
	},
}

const (
	AffectedFreqRange_r18_AffectedBandwidth_r18_Khz200  = 0
	AffectedFreqRange_r18_AffectedBandwidth_r18_Khz400  = 1
	AffectedFreqRange_r18_AffectedBandwidth_r18_Khz600  = 2
	AffectedFreqRange_r18_AffectedBandwidth_r18_Khz800  = 3
	AffectedFreqRange_r18_AffectedBandwidth_r18_Mhz1    = 4
	AffectedFreqRange_r18_AffectedBandwidth_r18_Mhz2    = 5
	AffectedFreqRange_r18_AffectedBandwidth_r18_Mhz3    = 6
	AffectedFreqRange_r18_AffectedBandwidth_r18_Mhz4    = 7
	AffectedFreqRange_r18_AffectedBandwidth_r18_Mhz5    = 8
	AffectedFreqRange_r18_AffectedBandwidth_r18_Mhz6    = 9
	AffectedFreqRange_r18_AffectedBandwidth_r18_Mhz8    = 10
	AffectedFreqRange_r18_AffectedBandwidth_r18_Mhz10   = 11
	AffectedFreqRange_r18_AffectedBandwidth_r18_Mhz20   = 12
	AffectedFreqRange_r18_AffectedBandwidth_r18_Mhz30   = 13
	AffectedFreqRange_r18_AffectedBandwidth_r18_Mhz40   = 14
	AffectedFreqRange_r18_AffectedBandwidth_r18_Mhz50   = 15
	AffectedFreqRange_r18_AffectedBandwidth_r18_Mhz60   = 16
	AffectedFreqRange_r18_AffectedBandwidth_r18_Mhz80   = 17
	AffectedFreqRange_r18_AffectedBandwidth_r18_Mhz100  = 18
	AffectedFreqRange_r18_AffectedBandwidth_r18_Mhz200  = 19
	AffectedFreqRange_r18_AffectedBandwidth_r18_Mhz300  = 20
	AffectedFreqRange_r18_AffectedBandwidth_r18_Mhz400  = 21
	AffectedFreqRange_r18_AffectedBandwidth_r18_Spare10 = 22
	AffectedFreqRange_r18_AffectedBandwidth_r18_Spare9  = 23
	AffectedFreqRange_r18_AffectedBandwidth_r18_Spare8  = 24
	AffectedFreqRange_r18_AffectedBandwidth_r18_Spare7  = 25
	AffectedFreqRange_r18_AffectedBandwidth_r18_Spare6  = 26
	AffectedFreqRange_r18_AffectedBandwidth_r18_Spare5  = 27
	AffectedFreqRange_r18_AffectedBandwidth_r18_Spare4  = 28
	AffectedFreqRange_r18_AffectedBandwidth_r18_Spare3  = 29
	AffectedFreqRange_r18_AffectedBandwidth_r18_Spare2  = 30
	AffectedFreqRange_r18_AffectedBandwidth_r18_Spare1  = 31
)

var affectedFreqRangeR18AffectedBandwidthR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{AffectedFreqRange_r18_AffectedBandwidth_r18_Khz200, AffectedFreqRange_r18_AffectedBandwidth_r18_Khz400, AffectedFreqRange_r18_AffectedBandwidth_r18_Khz600, AffectedFreqRange_r18_AffectedBandwidth_r18_Khz800, AffectedFreqRange_r18_AffectedBandwidth_r18_Mhz1, AffectedFreqRange_r18_AffectedBandwidth_r18_Mhz2, AffectedFreqRange_r18_AffectedBandwidth_r18_Mhz3, AffectedFreqRange_r18_AffectedBandwidth_r18_Mhz4, AffectedFreqRange_r18_AffectedBandwidth_r18_Mhz5, AffectedFreqRange_r18_AffectedBandwidth_r18_Mhz6, AffectedFreqRange_r18_AffectedBandwidth_r18_Mhz8, AffectedFreqRange_r18_AffectedBandwidth_r18_Mhz10, AffectedFreqRange_r18_AffectedBandwidth_r18_Mhz20, AffectedFreqRange_r18_AffectedBandwidth_r18_Mhz30, AffectedFreqRange_r18_AffectedBandwidth_r18_Mhz40, AffectedFreqRange_r18_AffectedBandwidth_r18_Mhz50, AffectedFreqRange_r18_AffectedBandwidth_r18_Mhz60, AffectedFreqRange_r18_AffectedBandwidth_r18_Mhz80, AffectedFreqRange_r18_AffectedBandwidth_r18_Mhz100, AffectedFreqRange_r18_AffectedBandwidth_r18_Mhz200, AffectedFreqRange_r18_AffectedBandwidth_r18_Mhz300, AffectedFreqRange_r18_AffectedBandwidth_r18_Mhz400, AffectedFreqRange_r18_AffectedBandwidth_r18_Spare10, AffectedFreqRange_r18_AffectedBandwidth_r18_Spare9, AffectedFreqRange_r18_AffectedBandwidth_r18_Spare8, AffectedFreqRange_r18_AffectedBandwidth_r18_Spare7, AffectedFreqRange_r18_AffectedBandwidth_r18_Spare6, AffectedFreqRange_r18_AffectedBandwidth_r18_Spare5, AffectedFreqRange_r18_AffectedBandwidth_r18_Spare4, AffectedFreqRange_r18_AffectedBandwidth_r18_Spare3, AffectedFreqRange_r18_AffectedBandwidth_r18_Spare2, AffectedFreqRange_r18_AffectedBandwidth_r18_Spare1},
}

type AffectedFreqRange_r18 struct {
	CenterFreq_r18        ARFCN_ValueNR
	AffectedBandwidth_r18 int64
}

func (ie *AffectedFreqRange_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(affectedFreqRangeR18Constraints)
	_ = seq

	// 1. centerFreq-r18: ref
	{
		if err := ie.CenterFreq_r18.Encode(e); err != nil {
			return err
		}
	}

	// 2. affectedBandwidth-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.AffectedBandwidth_r18, affectedFreqRangeR18AffectedBandwidthR18Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *AffectedFreqRange_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(affectedFreqRangeR18Constraints)
	_ = seq

	// 1. centerFreq-r18: ref
	{
		if err := ie.CenterFreq_r18.Decode(d); err != nil {
			return err
		}
	}

	// 2. affectedBandwidth-r18: enumerated
	{
		v1, err := d.DecodeEnumerated(affectedFreqRangeR18AffectedBandwidthR18Constraints)
		if err != nil {
			return err
		}
		ie.AffectedBandwidth_r18 = v1
	}

	return nil
}
