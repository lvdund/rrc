// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NeedForInterruptionInfoNR-r18 (line 10523).

var needForInterruptionInfoNRR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "intraFreq-needForInterruption-r18"},
		{Name: "interFreq-needForInterruption-r18"},
	},
}

type NeedForInterruptionInfoNR_r18 struct {
	IntraFreq_NeedForInterruption_r18 NeedForInterruptionIntraFreqList_r18
	InterFreq_NeedForInterruption_r18 NeedForInterruptionBandListNR_r18
}

func (ie *NeedForInterruptionInfoNR_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(needForInterruptionInfoNRR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. intraFreq-needForInterruption-r18: ref
	{
		if err := ie.IntraFreq_NeedForInterruption_r18.Encode(e); err != nil {
			return err
		}
	}

	// 3. interFreq-needForInterruption-r18: ref
	{
		if err := ie.InterFreq_NeedForInterruption_r18.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *NeedForInterruptionInfoNR_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(needForInterruptionInfoNRR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. intraFreq-needForInterruption-r18: ref
	{
		if err := ie.IntraFreq_NeedForInterruption_r18.Decode(d); err != nil {
			return err
		}
	}

	// 3. interFreq-needForInterruption-r18: ref
	{
		if err := ie.InterFreq_NeedForInterruption_r18.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
