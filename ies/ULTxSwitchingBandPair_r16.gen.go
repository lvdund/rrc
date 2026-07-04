// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: ULTxSwitchingBandPair-r16 (line 16954).

var uLTxSwitchingBandPairR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "bandIndexUL1-r16"},
		{Name: "bandIndexUL2-r16"},
		{Name: "uplinkTxSwitchingPeriod-r16"},
		{Name: "uplinkTxSwitching-DL-Interruption-r16", Optional: true},
	},
}

var uLTxSwitchingBandPairR16BandIndexUL1R16Constraints = per.Constrained(1, common.MaxSimultaneousBands)

var uLTxSwitchingBandPairR16BandIndexUL2R16Constraints = per.Constrained(1, common.MaxSimultaneousBands)

const (
	ULTxSwitchingBandPair_r16_UplinkTxSwitchingPeriod_r16_N35us  = 0
	ULTxSwitchingBandPair_r16_UplinkTxSwitchingPeriod_r16_N140us = 1
	ULTxSwitchingBandPair_r16_UplinkTxSwitchingPeriod_r16_N210us = 2
)

var uLTxSwitchingBandPairR16UplinkTxSwitchingPeriodR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ULTxSwitchingBandPair_r16_UplinkTxSwitchingPeriod_r16_N35us, ULTxSwitchingBandPair_r16_UplinkTxSwitchingPeriod_r16_N140us, ULTxSwitchingBandPair_r16_UplinkTxSwitchingPeriod_r16_N210us},
}

var uLTxSwitchingBandPairR16UplinkTxSwitchingDLInterruptionR16Constraints = per.SizeRange(1, common.MaxSimultaneousBands)

type ULTxSwitchingBandPair_r16 struct {
	BandIndexUL1_r16                      int64
	BandIndexUL2_r16                      int64
	UplinkTxSwitchingPeriod_r16           int64
	UplinkTxSwitching_DL_Interruption_r16 *per.BitString
}

func (ie *ULTxSwitchingBandPair_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uLTxSwitchingBandPairR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.UplinkTxSwitching_DL_Interruption_r16 != nil}); err != nil {
		return err
	}

	// 2. bandIndexUL1-r16: integer
	{
		if err := e.EncodeInteger(ie.BandIndexUL1_r16, uLTxSwitchingBandPairR16BandIndexUL1R16Constraints); err != nil {
			return err
		}
	}

	// 3. bandIndexUL2-r16: integer
	{
		if err := e.EncodeInteger(ie.BandIndexUL2_r16, uLTxSwitchingBandPairR16BandIndexUL2R16Constraints); err != nil {
			return err
		}
	}

	// 4. uplinkTxSwitchingPeriod-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.UplinkTxSwitchingPeriod_r16, uLTxSwitchingBandPairR16UplinkTxSwitchingPeriodR16Constraints); err != nil {
			return err
		}
	}

	// 5. uplinkTxSwitching-DL-Interruption-r16: bit-string
	{
		if ie.UplinkTxSwitching_DL_Interruption_r16 != nil {
			if err := e.EncodeBitString(*ie.UplinkTxSwitching_DL_Interruption_r16, uLTxSwitchingBandPairR16UplinkTxSwitchingDLInterruptionR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *ULTxSwitchingBandPair_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uLTxSwitchingBandPairR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. bandIndexUL1-r16: integer
	{
		v0, err := d.DecodeInteger(uLTxSwitchingBandPairR16BandIndexUL1R16Constraints)
		if err != nil {
			return err
		}
		ie.BandIndexUL1_r16 = v0
	}

	// 3. bandIndexUL2-r16: integer
	{
		v1, err := d.DecodeInteger(uLTxSwitchingBandPairR16BandIndexUL2R16Constraints)
		if err != nil {
			return err
		}
		ie.BandIndexUL2_r16 = v1
	}

	// 4. uplinkTxSwitchingPeriod-r16: enumerated
	{
		v2, err := d.DecodeEnumerated(uLTxSwitchingBandPairR16UplinkTxSwitchingPeriodR16Constraints)
		if err != nil {
			return err
		}
		ie.UplinkTxSwitchingPeriod_r16 = v2
	}

	// 5. uplinkTxSwitching-DL-Interruption-r16: bit-string
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeBitString(uLTxSwitchingBandPairR16UplinkTxSwitchingDLInterruptionR16Constraints)
			if err != nil {
				return err
			}
			ie.UplinkTxSwitching_DL_Interruption_r16 = &v
		}
	}

	return nil
}
