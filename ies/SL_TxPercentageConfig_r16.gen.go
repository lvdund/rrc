// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-TxPercentageConfig-r16 (line 28091).

var sLTxPercentageConfigR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-Priority-r16"},
		{Name: "sl-TxPercentage-r16"},
	},
}

var sLTxPercentageConfigR16SlPriorityR16Constraints = per.Constrained(1, 8)

const (
	SL_TxPercentageConfig_r16_Sl_TxPercentage_r16_P20 = 0
	SL_TxPercentageConfig_r16_Sl_TxPercentage_r16_P35 = 1
	SL_TxPercentageConfig_r16_Sl_TxPercentage_r16_P50 = 2
)

var sLTxPercentageConfigR16SlTxPercentageR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_TxPercentageConfig_r16_Sl_TxPercentage_r16_P20, SL_TxPercentageConfig_r16_Sl_TxPercentage_r16_P35, SL_TxPercentageConfig_r16_Sl_TxPercentage_r16_P50},
}

type SL_TxPercentageConfig_r16 struct {
	Sl_Priority_r16     int64
	Sl_TxPercentage_r16 int64
}

func (ie *SL_TxPercentageConfig_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLTxPercentageConfigR16Constraints)
	_ = seq

	// 1. sl-Priority-r16: integer
	{
		if err := e.EncodeInteger(ie.Sl_Priority_r16, sLTxPercentageConfigR16SlPriorityR16Constraints); err != nil {
			return err
		}
	}

	// 2. sl-TxPercentage-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.Sl_TxPercentage_r16, sLTxPercentageConfigR16SlTxPercentageR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SL_TxPercentageConfig_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLTxPercentageConfigR16Constraints)
	_ = seq

	// 1. sl-Priority-r16: integer
	{
		v0, err := d.DecodeInteger(sLTxPercentageConfigR16SlPriorityR16Constraints)
		if err != nil {
			return err
		}
		ie.Sl_Priority_r16 = v0
	}

	// 2. sl-TxPercentage-r16: enumerated
	{
		v1, err := d.DecodeEnumerated(sLTxPercentageConfigR16SlTxPercentageR16Constraints)
		if err != nil {
			return err
		}
		ie.Sl_TxPercentage_r16 = v1
	}

	return nil
}
