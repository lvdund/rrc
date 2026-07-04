// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MaxMIMO-LayerPreferenceConfig-r16 (line 26450).

var maxMIMOLayerPreferenceConfigR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxMIMO-LayerPreferenceProhibitTimer-r16"},
	},
}

const (
	MaxMIMO_LayerPreferenceConfig_r16_MaxMIMO_LayerPreferenceProhibitTimer_r16_S0     = 0
	MaxMIMO_LayerPreferenceConfig_r16_MaxMIMO_LayerPreferenceProhibitTimer_r16_S0dot5 = 1
	MaxMIMO_LayerPreferenceConfig_r16_MaxMIMO_LayerPreferenceProhibitTimer_r16_S1     = 2
	MaxMIMO_LayerPreferenceConfig_r16_MaxMIMO_LayerPreferenceProhibitTimer_r16_S2     = 3
	MaxMIMO_LayerPreferenceConfig_r16_MaxMIMO_LayerPreferenceProhibitTimer_r16_S3     = 4
	MaxMIMO_LayerPreferenceConfig_r16_MaxMIMO_LayerPreferenceProhibitTimer_r16_S4     = 5
	MaxMIMO_LayerPreferenceConfig_r16_MaxMIMO_LayerPreferenceProhibitTimer_r16_S5     = 6
	MaxMIMO_LayerPreferenceConfig_r16_MaxMIMO_LayerPreferenceProhibitTimer_r16_S6     = 7
	MaxMIMO_LayerPreferenceConfig_r16_MaxMIMO_LayerPreferenceProhibitTimer_r16_S7     = 8
	MaxMIMO_LayerPreferenceConfig_r16_MaxMIMO_LayerPreferenceProhibitTimer_r16_S8     = 9
	MaxMIMO_LayerPreferenceConfig_r16_MaxMIMO_LayerPreferenceProhibitTimer_r16_S9     = 10
	MaxMIMO_LayerPreferenceConfig_r16_MaxMIMO_LayerPreferenceProhibitTimer_r16_S10    = 11
	MaxMIMO_LayerPreferenceConfig_r16_MaxMIMO_LayerPreferenceProhibitTimer_r16_S20    = 12
	MaxMIMO_LayerPreferenceConfig_r16_MaxMIMO_LayerPreferenceProhibitTimer_r16_S30    = 13
	MaxMIMO_LayerPreferenceConfig_r16_MaxMIMO_LayerPreferenceProhibitTimer_r16_Spare2 = 14
	MaxMIMO_LayerPreferenceConfig_r16_MaxMIMO_LayerPreferenceProhibitTimer_r16_Spare1 = 15
)

var maxMIMOLayerPreferenceConfigR16MaxMIMOLayerPreferenceProhibitTimerR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MaxMIMO_LayerPreferenceConfig_r16_MaxMIMO_LayerPreferenceProhibitTimer_r16_S0, MaxMIMO_LayerPreferenceConfig_r16_MaxMIMO_LayerPreferenceProhibitTimer_r16_S0dot5, MaxMIMO_LayerPreferenceConfig_r16_MaxMIMO_LayerPreferenceProhibitTimer_r16_S1, MaxMIMO_LayerPreferenceConfig_r16_MaxMIMO_LayerPreferenceProhibitTimer_r16_S2, MaxMIMO_LayerPreferenceConfig_r16_MaxMIMO_LayerPreferenceProhibitTimer_r16_S3, MaxMIMO_LayerPreferenceConfig_r16_MaxMIMO_LayerPreferenceProhibitTimer_r16_S4, MaxMIMO_LayerPreferenceConfig_r16_MaxMIMO_LayerPreferenceProhibitTimer_r16_S5, MaxMIMO_LayerPreferenceConfig_r16_MaxMIMO_LayerPreferenceProhibitTimer_r16_S6, MaxMIMO_LayerPreferenceConfig_r16_MaxMIMO_LayerPreferenceProhibitTimer_r16_S7, MaxMIMO_LayerPreferenceConfig_r16_MaxMIMO_LayerPreferenceProhibitTimer_r16_S8, MaxMIMO_LayerPreferenceConfig_r16_MaxMIMO_LayerPreferenceProhibitTimer_r16_S9, MaxMIMO_LayerPreferenceConfig_r16_MaxMIMO_LayerPreferenceProhibitTimer_r16_S10, MaxMIMO_LayerPreferenceConfig_r16_MaxMIMO_LayerPreferenceProhibitTimer_r16_S20, MaxMIMO_LayerPreferenceConfig_r16_MaxMIMO_LayerPreferenceProhibitTimer_r16_S30, MaxMIMO_LayerPreferenceConfig_r16_MaxMIMO_LayerPreferenceProhibitTimer_r16_Spare2, MaxMIMO_LayerPreferenceConfig_r16_MaxMIMO_LayerPreferenceProhibitTimer_r16_Spare1},
}

type MaxMIMO_LayerPreferenceConfig_r16 struct {
	MaxMIMO_LayerPreferenceProhibitTimer_r16 int64
}

func (ie *MaxMIMO_LayerPreferenceConfig_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(maxMIMOLayerPreferenceConfigR16Constraints)
	_ = seq

	// 1. maxMIMO-LayerPreferenceProhibitTimer-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxMIMO_LayerPreferenceProhibitTimer_r16, maxMIMOLayerPreferenceConfigR16MaxMIMOLayerPreferenceProhibitTimerR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *MaxMIMO_LayerPreferenceConfig_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(maxMIMOLayerPreferenceConfigR16Constraints)
	_ = seq

	// 1. maxMIMO-LayerPreferenceProhibitTimer-r16: enumerated
	{
		v0, err := d.DecodeEnumerated(maxMIMOLayerPreferenceConfigR16MaxMIMOLayerPreferenceProhibitTimerR16Constraints)
		if err != nil {
			return err
		}
		ie.MaxMIMO_LayerPreferenceProhibitTimer_r16 = v0
	}

	return nil
}
