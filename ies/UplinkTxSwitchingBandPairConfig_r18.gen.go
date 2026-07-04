// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UplinkTxSwitchingBandPairConfig-r18 (line 5825).

var uplinkTxSwitchingBandPairConfigR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "bandInfoUL1-r18"},
		{Name: "bandInfoUL2-r18"},
		{Name: "switchingOptionConfigForBandPair-r18"},
		{Name: "switching2T-Mode-r18", Optional: true},
		{Name: "switchingPeriodConfigForBandPair-r18", Optional: true},
	},
}

const (
	UplinkTxSwitchingBandPairConfig_r18_SwitchingOptionConfigForBandPair_r18_SwitchedUL = 0
	UplinkTxSwitchingBandPairConfig_r18_SwitchingOptionConfigForBandPair_r18_DualUL     = 1
)

var uplinkTxSwitchingBandPairConfigR18SwitchingOptionConfigForBandPairR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UplinkTxSwitchingBandPairConfig_r18_SwitchingOptionConfigForBandPair_r18_SwitchedUL, UplinkTxSwitchingBandPairConfig_r18_SwitchingOptionConfigForBandPair_r18_DualUL},
}

const (
	UplinkTxSwitchingBandPairConfig_r18_Switching2T_Mode_r18_Enabled = 0
)

var uplinkTxSwitchingBandPairConfigR18Switching2TModeR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UplinkTxSwitchingBandPairConfig_r18_Switching2T_Mode_r18_Enabled},
}

const (
	UplinkTxSwitchingBandPairConfig_r18_SwitchingPeriodConfigForBandPair_r18_N35us  = 0
	UplinkTxSwitchingBandPairConfig_r18_SwitchingPeriodConfigForBandPair_r18_N140us = 1
)

var uplinkTxSwitchingBandPairConfigR18SwitchingPeriodConfigForBandPairR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UplinkTxSwitchingBandPairConfig_r18_SwitchingPeriodConfigForBandPair_r18_N35us, UplinkTxSwitchingBandPairConfig_r18_SwitchingPeriodConfigForBandPair_r18_N140us},
}

type UplinkTxSwitchingBandPairConfig_r18 struct {
	BandInfoUL1_r18                      UplinkTxSwitchingBandIndex_r18
	BandInfoUL2_r18                      UplinkTxSwitchingBandIndex_r18
	SwitchingOptionConfigForBandPair_r18 int64
	Switching2T_Mode_r18                 *int64
	SwitchingPeriodConfigForBandPair_r18 *int64
}

func (ie *UplinkTxSwitchingBandPairConfig_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uplinkTxSwitchingBandPairConfigR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Switching2T_Mode_r18 != nil, ie.SwitchingPeriodConfigForBandPair_r18 != nil}); err != nil {
		return err
	}

	// 3. bandInfoUL1-r18: ref
	{
		if err := ie.BandInfoUL1_r18.Encode(e); err != nil {
			return err
		}
	}

	// 4. bandInfoUL2-r18: ref
	{
		if err := ie.BandInfoUL2_r18.Encode(e); err != nil {
			return err
		}
	}

	// 5. switchingOptionConfigForBandPair-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.SwitchingOptionConfigForBandPair_r18, uplinkTxSwitchingBandPairConfigR18SwitchingOptionConfigForBandPairR18Constraints); err != nil {
			return err
		}
	}

	// 6. switching2T-Mode-r18: enumerated
	{
		if ie.Switching2T_Mode_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Switching2T_Mode_r18, uplinkTxSwitchingBandPairConfigR18Switching2TModeR18Constraints); err != nil {
				return err
			}
		}
	}

	// 7. switchingPeriodConfigForBandPair-r18: enumerated
	{
		if ie.SwitchingPeriodConfigForBandPair_r18 != nil {
			if err := e.EncodeEnumerated(*ie.SwitchingPeriodConfigForBandPair_r18, uplinkTxSwitchingBandPairConfigR18SwitchingPeriodConfigForBandPairR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UplinkTxSwitchingBandPairConfig_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uplinkTxSwitchingBandPairConfigR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. bandInfoUL1-r18: ref
	{
		if err := ie.BandInfoUL1_r18.Decode(d); err != nil {
			return err
		}
	}

	// 4. bandInfoUL2-r18: ref
	{
		if err := ie.BandInfoUL2_r18.Decode(d); err != nil {
			return err
		}
	}

	// 5. switchingOptionConfigForBandPair-r18: enumerated
	{
		v2, err := d.DecodeEnumerated(uplinkTxSwitchingBandPairConfigR18SwitchingOptionConfigForBandPairR18Constraints)
		if err != nil {
			return err
		}
		ie.SwitchingOptionConfigForBandPair_r18 = v2
	}

	// 6. switching2T-Mode-r18: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(uplinkTxSwitchingBandPairConfigR18Switching2TModeR18Constraints)
			if err != nil {
				return err
			}
			ie.Switching2T_Mode_r18 = &idx
		}
	}

	// 7. switchingPeriodConfigForBandPair-r18: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(uplinkTxSwitchingBandPairConfigR18SwitchingPeriodConfigForBandPairR18Constraints)
			if err != nil {
				return err
			}
			ie.SwitchingPeriodConfigForBandPair_r18 = &idx
		}
	}

	return nil
}
