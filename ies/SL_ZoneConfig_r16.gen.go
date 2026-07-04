// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-ZoneConfig-r16 (line 28374).

var sLZoneConfigR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-ZoneLength-r16"},
	},
}

const (
	SL_ZoneConfig_r16_Sl_ZoneLength_r16_M5     = 0
	SL_ZoneConfig_r16_Sl_ZoneLength_r16_M10    = 1
	SL_ZoneConfig_r16_Sl_ZoneLength_r16_M20    = 2
	SL_ZoneConfig_r16_Sl_ZoneLength_r16_M30    = 3
	SL_ZoneConfig_r16_Sl_ZoneLength_r16_M40    = 4
	SL_ZoneConfig_r16_Sl_ZoneLength_r16_M50    = 5
	SL_ZoneConfig_r16_Sl_ZoneLength_r16_Spare2 = 6
	SL_ZoneConfig_r16_Sl_ZoneLength_r16_Spare1 = 7
)

var sLZoneConfigR16SlZoneLengthR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_ZoneConfig_r16_Sl_ZoneLength_r16_M5, SL_ZoneConfig_r16_Sl_ZoneLength_r16_M10, SL_ZoneConfig_r16_Sl_ZoneLength_r16_M20, SL_ZoneConfig_r16_Sl_ZoneLength_r16_M30, SL_ZoneConfig_r16_Sl_ZoneLength_r16_M40, SL_ZoneConfig_r16_Sl_ZoneLength_r16_M50, SL_ZoneConfig_r16_Sl_ZoneLength_r16_Spare2, SL_ZoneConfig_r16_Sl_ZoneLength_r16_Spare1},
}

type SL_ZoneConfig_r16 struct {
	Sl_ZoneLength_r16 int64
}

func (ie *SL_ZoneConfig_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLZoneConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. sl-ZoneLength-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.Sl_ZoneLength_r16, sLZoneConfigR16SlZoneLengthR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SL_ZoneConfig_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLZoneConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. sl-ZoneLength-r16: enumerated
	{
		v0, err := d.DecodeEnumerated(sLZoneConfigR16SlZoneLengthR16Constraints)
		if err != nil {
			return err
		}
		ie.Sl_ZoneLength_r16 = v0
	}

	return nil
}
