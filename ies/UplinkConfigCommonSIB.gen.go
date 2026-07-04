// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UplinkConfigCommonSIB (line 16312).

var uplinkConfigCommonSIBConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "frequencyInfoUL"},
		{Name: "initialUplinkBWP"},
		{Name: "timeAlignmentTimerCommon"},
	},
}

type UplinkConfigCommonSIB struct {
	FrequencyInfoUL          FrequencyInfoUL_SIB
	InitialUplinkBWP         BWP_UplinkCommon
	TimeAlignmentTimerCommon TimeAlignmentTimer
}

func (ie *UplinkConfigCommonSIB) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uplinkConfigCommonSIBConstraints)
	_ = seq

	// 1. frequencyInfoUL: ref
	{
		if err := ie.FrequencyInfoUL.Encode(e); err != nil {
			return err
		}
	}

	// 2. initialUplinkBWP: ref
	{
		if err := ie.InitialUplinkBWP.Encode(e); err != nil {
			return err
		}
	}

	// 3. timeAlignmentTimerCommon: ref
	{
		if err := ie.TimeAlignmentTimerCommon.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *UplinkConfigCommonSIB) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uplinkConfigCommonSIBConstraints)
	_ = seq

	// 1. frequencyInfoUL: ref
	{
		if err := ie.FrequencyInfoUL.Decode(d); err != nil {
			return err
		}
	}

	// 2. initialUplinkBWP: ref
	{
		if err := ie.InitialUplinkBWP.Decode(d); err != nil {
			return err
		}
	}

	// 3. timeAlignmentTimerCommon: ref
	{
		if err := ie.TimeAlignmentTimerCommon.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
