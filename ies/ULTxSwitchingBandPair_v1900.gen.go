// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: ULTxSwitchingBandPair-v1900 (line 16987).

var uLTxSwitchingBandPairV1900Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "uplink3TxSwitchingPeriodUpTo2TPerBandDualUL-v1900"},
		{Name: "uplinkTxSwitching-DL-Interruption-DualUL-v1900", Optional: true},
	},
}

const (
	ULTxSwitchingBandPair_v1900_Uplink3TxSwitchingPeriodUpTo2TPerBandDualUL_v1900_N35us  = 0
	ULTxSwitchingBandPair_v1900_Uplink3TxSwitchingPeriodUpTo2TPerBandDualUL_v1900_N140us = 1
	ULTxSwitchingBandPair_v1900_Uplink3TxSwitchingPeriodUpTo2TPerBandDualUL_v1900_N210us = 2
)

var uLTxSwitchingBandPairV1900Uplink3TxSwitchingPeriodUpTo2TPerBandDualULV1900Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ULTxSwitchingBandPair_v1900_Uplink3TxSwitchingPeriodUpTo2TPerBandDualUL_v1900_N35us, ULTxSwitchingBandPair_v1900_Uplink3TxSwitchingPeriodUpTo2TPerBandDualUL_v1900_N140us, ULTxSwitchingBandPair_v1900_Uplink3TxSwitchingPeriodUpTo2TPerBandDualUL_v1900_N210us},
}

var uLTxSwitchingBandPairV1900UplinkTxSwitchingDLInterruptionDualULV1900Constraints = per.SizeRange(1, common.MaxSimultaneousBands)

type ULTxSwitchingBandPair_v1900 struct {
	Uplink3TxSwitchingPeriodUpTo2TPerBandDualUL_v1900 int64
	UplinkTxSwitching_DL_Interruption_DualUL_v1900    *per.BitString
}

func (ie *ULTxSwitchingBandPair_v1900) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uLTxSwitchingBandPairV1900Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.UplinkTxSwitching_DL_Interruption_DualUL_v1900 != nil}); err != nil {
		return err
	}

	// 2. uplink3TxSwitchingPeriodUpTo2TPerBandDualUL-v1900: enumerated
	{
		if err := e.EncodeEnumerated(ie.Uplink3TxSwitchingPeriodUpTo2TPerBandDualUL_v1900, uLTxSwitchingBandPairV1900Uplink3TxSwitchingPeriodUpTo2TPerBandDualULV1900Constraints); err != nil {
			return err
		}
	}

	// 3. uplinkTxSwitching-DL-Interruption-DualUL-v1900: bit-string
	{
		if ie.UplinkTxSwitching_DL_Interruption_DualUL_v1900 != nil {
			if err := e.EncodeBitString(*ie.UplinkTxSwitching_DL_Interruption_DualUL_v1900, uLTxSwitchingBandPairV1900UplinkTxSwitchingDLInterruptionDualULV1900Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *ULTxSwitchingBandPair_v1900) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uLTxSwitchingBandPairV1900Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. uplink3TxSwitchingPeriodUpTo2TPerBandDualUL-v1900: enumerated
	{
		v0, err := d.DecodeEnumerated(uLTxSwitchingBandPairV1900Uplink3TxSwitchingPeriodUpTo2TPerBandDualULV1900Constraints)
		if err != nil {
			return err
		}
		ie.Uplink3TxSwitchingPeriodUpTo2TPerBandDualUL_v1900 = v0
	}

	// 3. uplinkTxSwitching-DL-Interruption-DualUL-v1900: bit-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeBitString(uLTxSwitchingBandPairV1900UplinkTxSwitchingDLInterruptionDualULV1900Constraints)
			if err != nil {
				return err
			}
			ie.UplinkTxSwitching_DL_Interruption_DualUL_v1900 = &v
		}
	}

	return nil
}
