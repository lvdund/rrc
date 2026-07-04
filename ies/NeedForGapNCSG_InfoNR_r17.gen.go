// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NeedForGapNCSG-InfoNR-r17 (line 10501).

var needForGapNCSGInfoNRR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "intraFreq-needForNCSG-r17"},
		{Name: "interFreq-needForNCSG-r17"},
	},
}

type NeedForGapNCSG_InfoNR_r17 struct {
	IntraFreq_NeedForNCSG_r17 NeedForNCSG_IntraFreqList_r17
	InterFreq_NeedForNCSG_r17 NeedForNCSG_BandListNR_r17
}

func (ie *NeedForGapNCSG_InfoNR_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(needForGapNCSGInfoNRR17Constraints)
	_ = seq

	// 1. intraFreq-needForNCSG-r17: ref
	{
		if err := ie.IntraFreq_NeedForNCSG_r17.Encode(e); err != nil {
			return err
		}
	}

	// 2. interFreq-needForNCSG-r17: ref
	{
		if err := ie.InterFreq_NeedForNCSG_r17.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *NeedForGapNCSG_InfoNR_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(needForGapNCSGInfoNRR17Constraints)
	_ = seq

	// 1. intraFreq-needForNCSG-r17: ref
	{
		if err := ie.IntraFreq_NeedForNCSG_r17.Decode(d); err != nil {
			return err
		}
	}

	// 2. interFreq-needForNCSG-r17: ref
	{
		if err := ie.InterFreq_NeedForNCSG_r17.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
