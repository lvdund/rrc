// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MAC-ParametersSidelinkXDD-Diff-r16 (line 25086).

var mACParametersSidelinkXDDDiffR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "multipleSR-ConfigurationsSidelink-r16", Optional: true},
		{Name: "logicalChannelSR-DelayTimerSidelink-r16", Optional: true},
	},
}

const (
	MAC_ParametersSidelinkXDD_Diff_r16_MultipleSR_ConfigurationsSidelink_r16_Supported = 0
)

var mACParametersSidelinkXDDDiffR16MultipleSRConfigurationsSidelinkR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersSidelinkXDD_Diff_r16_MultipleSR_ConfigurationsSidelink_r16_Supported},
}

const (
	MAC_ParametersSidelinkXDD_Diff_r16_LogicalChannelSR_DelayTimerSidelink_r16_Supported = 0
)

var mACParametersSidelinkXDDDiffR16LogicalChannelSRDelayTimerSidelinkR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersSidelinkXDD_Diff_r16_LogicalChannelSR_DelayTimerSidelink_r16_Supported},
}

type MAC_ParametersSidelinkXDD_Diff_r16 struct {
	MultipleSR_ConfigurationsSidelink_r16   *int64
	LogicalChannelSR_DelayTimerSidelink_r16 *int64
}

func (ie *MAC_ParametersSidelinkXDD_Diff_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mACParametersSidelinkXDDDiffR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MultipleSR_ConfigurationsSidelink_r16 != nil, ie.LogicalChannelSR_DelayTimerSidelink_r16 != nil}); err != nil {
		return err
	}

	// 3. multipleSR-ConfigurationsSidelink-r16: enumerated
	{
		if ie.MultipleSR_ConfigurationsSidelink_r16 != nil {
			if err := e.EncodeEnumerated(*ie.MultipleSR_ConfigurationsSidelink_r16, mACParametersSidelinkXDDDiffR16MultipleSRConfigurationsSidelinkR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. logicalChannelSR-DelayTimerSidelink-r16: enumerated
	{
		if ie.LogicalChannelSR_DelayTimerSidelink_r16 != nil {
			if err := e.EncodeEnumerated(*ie.LogicalChannelSR_DelayTimerSidelink_r16, mACParametersSidelinkXDDDiffR16LogicalChannelSRDelayTimerSidelinkR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MAC_ParametersSidelinkXDD_Diff_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mACParametersSidelinkXDDDiffR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. multipleSR-ConfigurationsSidelink-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(mACParametersSidelinkXDDDiffR16MultipleSRConfigurationsSidelinkR16Constraints)
			if err != nil {
				return err
			}
			ie.MultipleSR_ConfigurationsSidelink_r16 = &idx
		}
	}

	// 4. logicalChannelSR-DelayTimerSidelink-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(mACParametersSidelinkXDDDiffR16LogicalChannelSRDelayTimerSidelinkR16Constraints)
			if err != nil {
				return err
			}
			ie.LogicalChannelSR_DelayTimerSidelink_r16 = &idx
		}
	}

	return nil
}
