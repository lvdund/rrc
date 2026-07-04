// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SRS-PosResourceSP-r16 (line 20473).

var sRSPosResourceSPR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxNumberSP-SRS-PosResourcesPerBWP-r16"},
		{Name: "maxNumberSP-SRS-PosResourcesPerBWP-PerSlot-r16"},
	},
}

const (
	SRS_PosResourceSP_r16_MaxNumberSP_SRS_PosResourcesPerBWP_r16_N1  = 0
	SRS_PosResourceSP_r16_MaxNumberSP_SRS_PosResourcesPerBWP_r16_N2  = 1
	SRS_PosResourceSP_r16_MaxNumberSP_SRS_PosResourcesPerBWP_r16_N4  = 2
	SRS_PosResourceSP_r16_MaxNumberSP_SRS_PosResourcesPerBWP_r16_N8  = 3
	SRS_PosResourceSP_r16_MaxNumberSP_SRS_PosResourcesPerBWP_r16_N16 = 4
	SRS_PosResourceSP_r16_MaxNumberSP_SRS_PosResourcesPerBWP_r16_N32 = 5
	SRS_PosResourceSP_r16_MaxNumberSP_SRS_PosResourcesPerBWP_r16_N64 = 6
)

var sRSPosResourceSPR16MaxNumberSPSRSPosResourcesPerBWPR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_PosResourceSP_r16_MaxNumberSP_SRS_PosResourcesPerBWP_r16_N1, SRS_PosResourceSP_r16_MaxNumberSP_SRS_PosResourcesPerBWP_r16_N2, SRS_PosResourceSP_r16_MaxNumberSP_SRS_PosResourcesPerBWP_r16_N4, SRS_PosResourceSP_r16_MaxNumberSP_SRS_PosResourcesPerBWP_r16_N8, SRS_PosResourceSP_r16_MaxNumberSP_SRS_PosResourcesPerBWP_r16_N16, SRS_PosResourceSP_r16_MaxNumberSP_SRS_PosResourcesPerBWP_r16_N32, SRS_PosResourceSP_r16_MaxNumberSP_SRS_PosResourcesPerBWP_r16_N64},
}

const (
	SRS_PosResourceSP_r16_MaxNumberSP_SRS_PosResourcesPerBWP_PerSlot_r16_N1  = 0
	SRS_PosResourceSP_r16_MaxNumberSP_SRS_PosResourcesPerBWP_PerSlot_r16_N2  = 1
	SRS_PosResourceSP_r16_MaxNumberSP_SRS_PosResourcesPerBWP_PerSlot_r16_N3  = 2
	SRS_PosResourceSP_r16_MaxNumberSP_SRS_PosResourcesPerBWP_PerSlot_r16_N4  = 3
	SRS_PosResourceSP_r16_MaxNumberSP_SRS_PosResourcesPerBWP_PerSlot_r16_N5  = 4
	SRS_PosResourceSP_r16_MaxNumberSP_SRS_PosResourcesPerBWP_PerSlot_r16_N6  = 5
	SRS_PosResourceSP_r16_MaxNumberSP_SRS_PosResourcesPerBWP_PerSlot_r16_N8  = 6
	SRS_PosResourceSP_r16_MaxNumberSP_SRS_PosResourcesPerBWP_PerSlot_r16_N10 = 7
	SRS_PosResourceSP_r16_MaxNumberSP_SRS_PosResourcesPerBWP_PerSlot_r16_N12 = 8
	SRS_PosResourceSP_r16_MaxNumberSP_SRS_PosResourcesPerBWP_PerSlot_r16_N14 = 9
)

var sRSPosResourceSPR16MaxNumberSPSRSPosResourcesPerBWPPerSlotR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_PosResourceSP_r16_MaxNumberSP_SRS_PosResourcesPerBWP_PerSlot_r16_N1, SRS_PosResourceSP_r16_MaxNumberSP_SRS_PosResourcesPerBWP_PerSlot_r16_N2, SRS_PosResourceSP_r16_MaxNumberSP_SRS_PosResourcesPerBWP_PerSlot_r16_N3, SRS_PosResourceSP_r16_MaxNumberSP_SRS_PosResourcesPerBWP_PerSlot_r16_N4, SRS_PosResourceSP_r16_MaxNumberSP_SRS_PosResourcesPerBWP_PerSlot_r16_N5, SRS_PosResourceSP_r16_MaxNumberSP_SRS_PosResourcesPerBWP_PerSlot_r16_N6, SRS_PosResourceSP_r16_MaxNumberSP_SRS_PosResourcesPerBWP_PerSlot_r16_N8, SRS_PosResourceSP_r16_MaxNumberSP_SRS_PosResourcesPerBWP_PerSlot_r16_N10, SRS_PosResourceSP_r16_MaxNumberSP_SRS_PosResourcesPerBWP_PerSlot_r16_N12, SRS_PosResourceSP_r16_MaxNumberSP_SRS_PosResourcesPerBWP_PerSlot_r16_N14},
}

type SRS_PosResourceSP_r16 struct {
	MaxNumberSP_SRS_PosResourcesPerBWP_r16         int64
	MaxNumberSP_SRS_PosResourcesPerBWP_PerSlot_r16 int64
}

func (ie *SRS_PosResourceSP_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sRSPosResourceSPR16Constraints)
	_ = seq

	// 1. maxNumberSP-SRS-PosResourcesPerBWP-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxNumberSP_SRS_PosResourcesPerBWP_r16, sRSPosResourceSPR16MaxNumberSPSRSPosResourcesPerBWPR16Constraints); err != nil {
			return err
		}
	}

	// 2. maxNumberSP-SRS-PosResourcesPerBWP-PerSlot-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxNumberSP_SRS_PosResourcesPerBWP_PerSlot_r16, sRSPosResourceSPR16MaxNumberSPSRSPosResourcesPerBWPPerSlotR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SRS_PosResourceSP_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sRSPosResourceSPR16Constraints)
	_ = seq

	// 1. maxNumberSP-SRS-PosResourcesPerBWP-r16: enumerated
	{
		v0, err := d.DecodeEnumerated(sRSPosResourceSPR16MaxNumberSPSRSPosResourcesPerBWPR16Constraints)
		if err != nil {
			return err
		}
		ie.MaxNumberSP_SRS_PosResourcesPerBWP_r16 = v0
	}

	// 2. maxNumberSP-SRS-PosResourcesPerBWP-PerSlot-r16: enumerated
	{
		v1, err := d.DecodeEnumerated(sRSPosResourceSPR16MaxNumberSPSRSPosResourcesPerBWPPerSlotR16Constraints)
		if err != nil {
			return err
		}
		ie.MaxNumberSP_SRS_PosResourcesPerBWP_PerSlot_r16 = v1
	}

	return nil
}
