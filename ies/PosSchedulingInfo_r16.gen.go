// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PosSchedulingInfo-r16 (line 4892).

var posSchedulingInfoR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "offsetToSI-Used-r16", Optional: true},
		{Name: "posSI-Periodicity-r16"},
		{Name: "posSI-BroadcastStatus-r16"},
		{Name: "posSIB-MappingInfo-r16"},
	},
}

const (
	PosSchedulingInfo_r16_OffsetToSI_Used_r16_True = 0
)

var posSchedulingInfoR16OffsetToSIUsedR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSchedulingInfo_r16_OffsetToSI_Used_r16_True},
}

const (
	PosSchedulingInfo_r16_PosSI_Periodicity_r16_Rf8   = 0
	PosSchedulingInfo_r16_PosSI_Periodicity_r16_Rf16  = 1
	PosSchedulingInfo_r16_PosSI_Periodicity_r16_Rf32  = 2
	PosSchedulingInfo_r16_PosSI_Periodicity_r16_Rf64  = 3
	PosSchedulingInfo_r16_PosSI_Periodicity_r16_Rf128 = 4
	PosSchedulingInfo_r16_PosSI_Periodicity_r16_Rf256 = 5
	PosSchedulingInfo_r16_PosSI_Periodicity_r16_Rf512 = 6
)

var posSchedulingInfoR16PosSIPeriodicityR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSchedulingInfo_r16_PosSI_Periodicity_r16_Rf8, PosSchedulingInfo_r16_PosSI_Periodicity_r16_Rf16, PosSchedulingInfo_r16_PosSI_Periodicity_r16_Rf32, PosSchedulingInfo_r16_PosSI_Periodicity_r16_Rf64, PosSchedulingInfo_r16_PosSI_Periodicity_r16_Rf128, PosSchedulingInfo_r16_PosSI_Periodicity_r16_Rf256, PosSchedulingInfo_r16_PosSI_Periodicity_r16_Rf512},
}

const (
	PosSchedulingInfo_r16_PosSI_BroadcastStatus_r16_Broadcasting    = 0
	PosSchedulingInfo_r16_PosSI_BroadcastStatus_r16_NotBroadcasting = 1
)

var posSchedulingInfoR16PosSIBroadcastStatusR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSchedulingInfo_r16_PosSI_BroadcastStatus_r16_Broadcasting, PosSchedulingInfo_r16_PosSI_BroadcastStatus_r16_NotBroadcasting},
}

type PosSchedulingInfo_r16 struct {
	OffsetToSI_Used_r16       *int64
	PosSI_Periodicity_r16     int64
	PosSI_BroadcastStatus_r16 int64
	PosSIB_MappingInfo_r16    PosSIB_MappingInfo_r16
}

func (ie *PosSchedulingInfo_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(posSchedulingInfoR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.OffsetToSI_Used_r16 != nil}); err != nil {
		return err
	}

	// 3. offsetToSI-Used-r16: enumerated
	{
		if ie.OffsetToSI_Used_r16 != nil {
			if err := e.EncodeEnumerated(*ie.OffsetToSI_Used_r16, posSchedulingInfoR16OffsetToSIUsedR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. posSI-Periodicity-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.PosSI_Periodicity_r16, posSchedulingInfoR16PosSIPeriodicityR16Constraints); err != nil {
			return err
		}
	}

	// 5. posSI-BroadcastStatus-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.PosSI_BroadcastStatus_r16, posSchedulingInfoR16PosSIBroadcastStatusR16Constraints); err != nil {
			return err
		}
	}

	// 6. posSIB-MappingInfo-r16: ref
	{
		if err := ie.PosSIB_MappingInfo_r16.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PosSchedulingInfo_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(posSchedulingInfoR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. offsetToSI-Used-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(posSchedulingInfoR16OffsetToSIUsedR16Constraints)
			if err != nil {
				return err
			}
			ie.OffsetToSI_Used_r16 = &idx
		}
	}

	// 4. posSI-Periodicity-r16: enumerated
	{
		v1, err := d.DecodeEnumerated(posSchedulingInfoR16PosSIPeriodicityR16Constraints)
		if err != nil {
			return err
		}
		ie.PosSI_Periodicity_r16 = v1
	}

	// 5. posSI-BroadcastStatus-r16: enumerated
	{
		v2, err := d.DecodeEnumerated(posSchedulingInfoR16PosSIBroadcastStatusR16Constraints)
		if err != nil {
			return err
		}
		ie.PosSI_BroadcastStatus_r16 = v2
	}

	// 6. posSIB-MappingInfo-r16: ref
	{
		if err := ie.PosSIB_MappingInfo_r16.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
