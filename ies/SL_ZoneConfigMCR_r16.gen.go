// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-ZoneConfigMCR-r16 (line 28011).

var sLZoneConfigMCRR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-ZoneConfigMCR-Index-r16"},
		{Name: "sl-TransRange-r16", Optional: true},
		{Name: "sl-ZoneConfig-r16", Optional: true},
	},
}

var sLZoneConfigMCRR16SlZoneConfigMCRIndexR16Constraints = per.Constrained(0, 15)

const (
	SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M20    = 0
	SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M50    = 1
	SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M80    = 2
	SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M100   = 3
	SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M120   = 4
	SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M150   = 5
	SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M180   = 6
	SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M200   = 7
	SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M220   = 8
	SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M250   = 9
	SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M270   = 10
	SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M300   = 11
	SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M350   = 12
	SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M370   = 13
	SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M400   = 14
	SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M420   = 15
	SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M450   = 16
	SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M480   = 17
	SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M500   = 18
	SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M550   = 19
	SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M600   = 20
	SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M700   = 21
	SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M1000  = 22
	SL_ZoneConfigMCR_r16_Sl_TransRange_r16_Spare9 = 23
	SL_ZoneConfigMCR_r16_Sl_TransRange_r16_Spare8 = 24
	SL_ZoneConfigMCR_r16_Sl_TransRange_r16_Spare7 = 25
	SL_ZoneConfigMCR_r16_Sl_TransRange_r16_Spare6 = 26
	SL_ZoneConfigMCR_r16_Sl_TransRange_r16_Spare5 = 27
	SL_ZoneConfigMCR_r16_Sl_TransRange_r16_Spare4 = 28
	SL_ZoneConfigMCR_r16_Sl_TransRange_r16_Spare3 = 29
	SL_ZoneConfigMCR_r16_Sl_TransRange_r16_Spare2 = 30
	SL_ZoneConfigMCR_r16_Sl_TransRange_r16_Spare1 = 31
)

var sLZoneConfigMCRR16SlTransRangeR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M20, SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M50, SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M80, SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M100, SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M120, SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M150, SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M180, SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M200, SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M220, SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M250, SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M270, SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M300, SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M350, SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M370, SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M400, SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M420, SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M450, SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M480, SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M500, SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M550, SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M600, SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M700, SL_ZoneConfigMCR_r16_Sl_TransRange_r16_M1000, SL_ZoneConfigMCR_r16_Sl_TransRange_r16_Spare9, SL_ZoneConfigMCR_r16_Sl_TransRange_r16_Spare8, SL_ZoneConfigMCR_r16_Sl_TransRange_r16_Spare7, SL_ZoneConfigMCR_r16_Sl_TransRange_r16_Spare6, SL_ZoneConfigMCR_r16_Sl_TransRange_r16_Spare5, SL_ZoneConfigMCR_r16_Sl_TransRange_r16_Spare4, SL_ZoneConfigMCR_r16_Sl_TransRange_r16_Spare3, SL_ZoneConfigMCR_r16_Sl_TransRange_r16_Spare2, SL_ZoneConfigMCR_r16_Sl_TransRange_r16_Spare1},
}

type SL_ZoneConfigMCR_r16 struct {
	Sl_ZoneConfigMCR_Index_r16 int64
	Sl_TransRange_r16          *int64
	Sl_ZoneConfig_r16          *SL_ZoneConfig_r16
}

func (ie *SL_ZoneConfigMCR_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLZoneConfigMCRR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_TransRange_r16 != nil, ie.Sl_ZoneConfig_r16 != nil}); err != nil {
		return err
	}

	// 3. sl-ZoneConfigMCR-Index-r16: integer
	{
		if err := e.EncodeInteger(ie.Sl_ZoneConfigMCR_Index_r16, sLZoneConfigMCRR16SlZoneConfigMCRIndexR16Constraints); err != nil {
			return err
		}
	}

	// 4. sl-TransRange-r16: enumerated
	{
		if ie.Sl_TransRange_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_TransRange_r16, sLZoneConfigMCRR16SlTransRangeR16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. sl-ZoneConfig-r16: ref
	{
		if ie.Sl_ZoneConfig_r16 != nil {
			if err := ie.Sl_ZoneConfig_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_ZoneConfigMCR_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLZoneConfigMCRR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-ZoneConfigMCR-Index-r16: integer
	{
		v0, err := d.DecodeInteger(sLZoneConfigMCRR16SlZoneConfigMCRIndexR16Constraints)
		if err != nil {
			return err
		}
		ie.Sl_ZoneConfigMCR_Index_r16 = v0
	}

	// 4. sl-TransRange-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(sLZoneConfigMCRR16SlTransRangeR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_TransRange_r16 = &idx
		}
	}

	// 5. sl-ZoneConfig-r16: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Sl_ZoneConfig_r16 = new(SL_ZoneConfig_r16)
			if err := ie.Sl_ZoneConfig_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
