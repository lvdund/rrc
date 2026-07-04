// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: IAB-ResourceConfig-r17 (line 5774).

var iABResourceConfigR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "iab-ResourceConfigID-r17"},
		{Name: "slotList-r17", Optional: true},
		{Name: "periodicitySlotList-r17", Optional: true},
		{Name: "slotListSubcarrierSpacing-r17", Optional: true},
	},
}

var iABResourceConfigR17SlotListR17Constraints = per.SizeRange(1, 5120)

const (
	IAB_ResourceConfig_r17_PeriodicitySlotList_r17_Ms0p5   = 0
	IAB_ResourceConfig_r17_PeriodicitySlotList_r17_Ms0p625 = 1
	IAB_ResourceConfig_r17_PeriodicitySlotList_r17_Ms1     = 2
	IAB_ResourceConfig_r17_PeriodicitySlotList_r17_Ms1p25  = 3
	IAB_ResourceConfig_r17_PeriodicitySlotList_r17_Ms2     = 4
	IAB_ResourceConfig_r17_PeriodicitySlotList_r17_Ms2p5   = 5
	IAB_ResourceConfig_r17_PeriodicitySlotList_r17_Ms5     = 6
	IAB_ResourceConfig_r17_PeriodicitySlotList_r17_Ms10    = 7
	IAB_ResourceConfig_r17_PeriodicitySlotList_r17_Ms20    = 8
	IAB_ResourceConfig_r17_PeriodicitySlotList_r17_Ms40    = 9
	IAB_ResourceConfig_r17_PeriodicitySlotList_r17_Ms80    = 10
	IAB_ResourceConfig_r17_PeriodicitySlotList_r17_Ms160   = 11
)

var iABResourceConfigR17PeriodicitySlotListR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{IAB_ResourceConfig_r17_PeriodicitySlotList_r17_Ms0p5, IAB_ResourceConfig_r17_PeriodicitySlotList_r17_Ms0p625, IAB_ResourceConfig_r17_PeriodicitySlotList_r17_Ms1, IAB_ResourceConfig_r17_PeriodicitySlotList_r17_Ms1p25, IAB_ResourceConfig_r17_PeriodicitySlotList_r17_Ms2, IAB_ResourceConfig_r17_PeriodicitySlotList_r17_Ms2p5, IAB_ResourceConfig_r17_PeriodicitySlotList_r17_Ms5, IAB_ResourceConfig_r17_PeriodicitySlotList_r17_Ms10, IAB_ResourceConfig_r17_PeriodicitySlotList_r17_Ms20, IAB_ResourceConfig_r17_PeriodicitySlotList_r17_Ms40, IAB_ResourceConfig_r17_PeriodicitySlotList_r17_Ms80, IAB_ResourceConfig_r17_PeriodicitySlotList_r17_Ms160},
}

type IAB_ResourceConfig_r17 struct {
	Iab_ResourceConfigID_r17      IAB_ResourceConfigID_r17
	SlotList_r17                  []int64
	PeriodicitySlotList_r17       *int64
	SlotListSubcarrierSpacing_r17 *SubcarrierSpacing
}

func (ie *IAB_ResourceConfig_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(iABResourceConfigR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SlotList_r17 != nil, ie.PeriodicitySlotList_r17 != nil, ie.SlotListSubcarrierSpacing_r17 != nil}); err != nil {
		return err
	}

	// 3. iab-ResourceConfigID-r17: ref
	{
		if err := ie.Iab_ResourceConfigID_r17.Encode(e); err != nil {
			return err
		}
	}

	// 4. slotList-r17: sequence-of
	{
		if ie.SlotList_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(iABResourceConfigR17SlotListR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.SlotList_r17))); err != nil {
				return err
			}
			for i := range ie.SlotList_r17 {
				if err := e.EncodeInteger(ie.SlotList_r17[i], per.Constrained(0, 5119)); err != nil {
					return err
				}
			}
		}
	}

	// 5. periodicitySlotList-r17: enumerated
	{
		if ie.PeriodicitySlotList_r17 != nil {
			if err := e.EncodeEnumerated(*ie.PeriodicitySlotList_r17, iABResourceConfigR17PeriodicitySlotListR17Constraints); err != nil {
				return err
			}
		}
	}

	// 6. slotListSubcarrierSpacing-r17: ref
	{
		if ie.SlotListSubcarrierSpacing_r17 != nil {
			if err := ie.SlotListSubcarrierSpacing_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *IAB_ResourceConfig_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(iABResourceConfigR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. iab-ResourceConfigID-r17: ref
	{
		if err := ie.Iab_ResourceConfigID_r17.Decode(d); err != nil {
			return err
		}
	}

	// 4. slotList-r17: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(iABResourceConfigR17SlotListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SlotList_r17 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, 5119))
				if err != nil {
					return err
				}
				ie.SlotList_r17[i] = v
			}
		}
	}

	// 5. periodicitySlotList-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(iABResourceConfigR17PeriodicitySlotListR17Constraints)
			if err != nil {
				return err
			}
			ie.PeriodicitySlotList_r17 = &idx
		}
	}

	// 6. slotListSubcarrierSpacing-r17: ref
	{
		if seq.IsComponentPresent(3) {
			ie.SlotListSubcarrierSpacing_r17 = new(SubcarrierSpacing)
			if err := ie.SlotListSubcarrierSpacing_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
