// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PerRASSBInfo-r16 (line 3170).

var perRASSBInfoR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ssb-Index-r16"},
		{Name: "numberOfPreamblesSentOnSSB-r16"},
		{Name: "perRAAttemptInfoList-r16"},
	},
}

var perRASSBInfoR16NumberOfPreamblesSentOnSSBR16Constraints = per.Constrained(1, 200)

type PerRASSBInfo_r16 struct {
	Ssb_Index_r16                  SSB_Index
	NumberOfPreamblesSentOnSSB_r16 int64
	PerRAAttemptInfoList_r16       PerRAAttemptInfoList_r16
}

func (ie *PerRASSBInfo_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(perRASSBInfoR16Constraints)
	_ = seq

	// 1. ssb-Index-r16: ref
	{
		if err := ie.Ssb_Index_r16.Encode(e); err != nil {
			return err
		}
	}

	// 2. numberOfPreamblesSentOnSSB-r16: integer
	{
		if err := e.EncodeInteger(ie.NumberOfPreamblesSentOnSSB_r16, perRASSBInfoR16NumberOfPreamblesSentOnSSBR16Constraints); err != nil {
			return err
		}
	}

	// 3. perRAAttemptInfoList-r16: ref
	{
		if err := ie.PerRAAttemptInfoList_r16.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PerRASSBInfo_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(perRASSBInfoR16Constraints)
	_ = seq

	// 1. ssb-Index-r16: ref
	{
		if err := ie.Ssb_Index_r16.Decode(d); err != nil {
			return err
		}
	}

	// 2. numberOfPreamblesSentOnSSB-r16: integer
	{
		v1, err := d.DecodeInteger(perRASSBInfoR16NumberOfPreamblesSentOnSSBR16Constraints)
		if err != nil {
			return err
		}
		ie.NumberOfPreamblesSentOnSSB_r16 = v1
	}

	// 3. perRAAttemptInfoList-r16: ref
	{
		if err := ie.PerRAAttemptInfoList_r16.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
