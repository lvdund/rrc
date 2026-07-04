// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NeedForGapsInfoNR-r16 (line 10453).

var needForGapsInfoNRR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "intraFreq-needForGap-r16"},
		{Name: "interFreq-needForGap-r16"},
	},
}

type NeedForGapsInfoNR_r16 struct {
	IntraFreq_NeedForGap_r16 NeedForGapsIntraFreqList_r16
	InterFreq_NeedForGap_r16 NeedForGapsBandListNR_r16
}

func (ie *NeedForGapsInfoNR_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(needForGapsInfoNRR16Constraints)
	_ = seq

	// 1. intraFreq-needForGap-r16: ref
	{
		if err := ie.IntraFreq_NeedForGap_r16.Encode(e); err != nil {
			return err
		}
	}

	// 2. interFreq-needForGap-r16: ref
	{
		if err := ie.InterFreq_NeedForGap_r16.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *NeedForGapsInfoNR_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(needForGapsInfoNRR16Constraints)
	_ = seq

	// 1. intraFreq-needForGap-r16: ref
	{
		if err := ie.IntraFreq_NeedForGap_r16.Decode(d); err != nil {
			return err
		}
	}

	// 2. interFreq-needForGap-r16: ref
	{
		if err := ie.InterFreq_NeedForGap_r16.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
