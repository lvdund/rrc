// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SRS-PosResourceAP-r16 (line 20468).

var sRSPosResourceAPR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxNumberAP-SRS-PosResourcesPerBWP-r16"},
		{Name: "maxNumberAP-SRS-PosResourcesPerBWP-PerSlot-r16"},
	},
}

const (
	SRS_PosResourceAP_r16_MaxNumberAP_SRS_PosResourcesPerBWP_r16_N1  = 0
	SRS_PosResourceAP_r16_MaxNumberAP_SRS_PosResourcesPerBWP_r16_N2  = 1
	SRS_PosResourceAP_r16_MaxNumberAP_SRS_PosResourcesPerBWP_r16_N4  = 2
	SRS_PosResourceAP_r16_MaxNumberAP_SRS_PosResourcesPerBWP_r16_N8  = 3
	SRS_PosResourceAP_r16_MaxNumberAP_SRS_PosResourcesPerBWP_r16_N16 = 4
	SRS_PosResourceAP_r16_MaxNumberAP_SRS_PosResourcesPerBWP_r16_N32 = 5
	SRS_PosResourceAP_r16_MaxNumberAP_SRS_PosResourcesPerBWP_r16_N64 = 6
)

var sRSPosResourceAPR16MaxNumberAPSRSPosResourcesPerBWPR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_PosResourceAP_r16_MaxNumberAP_SRS_PosResourcesPerBWP_r16_N1, SRS_PosResourceAP_r16_MaxNumberAP_SRS_PosResourcesPerBWP_r16_N2, SRS_PosResourceAP_r16_MaxNumberAP_SRS_PosResourcesPerBWP_r16_N4, SRS_PosResourceAP_r16_MaxNumberAP_SRS_PosResourcesPerBWP_r16_N8, SRS_PosResourceAP_r16_MaxNumberAP_SRS_PosResourcesPerBWP_r16_N16, SRS_PosResourceAP_r16_MaxNumberAP_SRS_PosResourcesPerBWP_r16_N32, SRS_PosResourceAP_r16_MaxNumberAP_SRS_PosResourcesPerBWP_r16_N64},
}

const (
	SRS_PosResourceAP_r16_MaxNumberAP_SRS_PosResourcesPerBWP_PerSlot_r16_N1  = 0
	SRS_PosResourceAP_r16_MaxNumberAP_SRS_PosResourcesPerBWP_PerSlot_r16_N2  = 1
	SRS_PosResourceAP_r16_MaxNumberAP_SRS_PosResourcesPerBWP_PerSlot_r16_N3  = 2
	SRS_PosResourceAP_r16_MaxNumberAP_SRS_PosResourcesPerBWP_PerSlot_r16_N4  = 3
	SRS_PosResourceAP_r16_MaxNumberAP_SRS_PosResourcesPerBWP_PerSlot_r16_N5  = 4
	SRS_PosResourceAP_r16_MaxNumberAP_SRS_PosResourcesPerBWP_PerSlot_r16_N6  = 5
	SRS_PosResourceAP_r16_MaxNumberAP_SRS_PosResourcesPerBWP_PerSlot_r16_N8  = 6
	SRS_PosResourceAP_r16_MaxNumberAP_SRS_PosResourcesPerBWP_PerSlot_r16_N10 = 7
	SRS_PosResourceAP_r16_MaxNumberAP_SRS_PosResourcesPerBWP_PerSlot_r16_N12 = 8
	SRS_PosResourceAP_r16_MaxNumberAP_SRS_PosResourcesPerBWP_PerSlot_r16_N14 = 9
)

var sRSPosResourceAPR16MaxNumberAPSRSPosResourcesPerBWPPerSlotR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_PosResourceAP_r16_MaxNumberAP_SRS_PosResourcesPerBWP_PerSlot_r16_N1, SRS_PosResourceAP_r16_MaxNumberAP_SRS_PosResourcesPerBWP_PerSlot_r16_N2, SRS_PosResourceAP_r16_MaxNumberAP_SRS_PosResourcesPerBWP_PerSlot_r16_N3, SRS_PosResourceAP_r16_MaxNumberAP_SRS_PosResourcesPerBWP_PerSlot_r16_N4, SRS_PosResourceAP_r16_MaxNumberAP_SRS_PosResourcesPerBWP_PerSlot_r16_N5, SRS_PosResourceAP_r16_MaxNumberAP_SRS_PosResourcesPerBWP_PerSlot_r16_N6, SRS_PosResourceAP_r16_MaxNumberAP_SRS_PosResourcesPerBWP_PerSlot_r16_N8, SRS_PosResourceAP_r16_MaxNumberAP_SRS_PosResourcesPerBWP_PerSlot_r16_N10, SRS_PosResourceAP_r16_MaxNumberAP_SRS_PosResourcesPerBWP_PerSlot_r16_N12, SRS_PosResourceAP_r16_MaxNumberAP_SRS_PosResourcesPerBWP_PerSlot_r16_N14},
}

type SRS_PosResourceAP_r16 struct {
	MaxNumberAP_SRS_PosResourcesPerBWP_r16         int64
	MaxNumberAP_SRS_PosResourcesPerBWP_PerSlot_r16 int64
}

func (ie *SRS_PosResourceAP_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sRSPosResourceAPR16Constraints)
	_ = seq

	// 1. maxNumberAP-SRS-PosResourcesPerBWP-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxNumberAP_SRS_PosResourcesPerBWP_r16, sRSPosResourceAPR16MaxNumberAPSRSPosResourcesPerBWPR16Constraints); err != nil {
			return err
		}
	}

	// 2. maxNumberAP-SRS-PosResourcesPerBWP-PerSlot-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxNumberAP_SRS_PosResourcesPerBWP_PerSlot_r16, sRSPosResourceAPR16MaxNumberAPSRSPosResourcesPerBWPPerSlotR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SRS_PosResourceAP_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sRSPosResourceAPR16Constraints)
	_ = seq

	// 1. maxNumberAP-SRS-PosResourcesPerBWP-r16: enumerated
	{
		v0, err := d.DecodeEnumerated(sRSPosResourceAPR16MaxNumberAPSRSPosResourcesPerBWPR16Constraints)
		if err != nil {
			return err
		}
		ie.MaxNumberAP_SRS_PosResourcesPerBWP_r16 = v0
	}

	// 2. maxNumberAP-SRS-PosResourcesPerBWP-PerSlot-r16: enumerated
	{
		v1, err := d.DecodeEnumerated(sRSPosResourceAPR16MaxNumberAPSRSPosResourcesPerBWPPerSlotR16Constraints)
		if err != nil {
			return err
		}
		ie.MaxNumberAP_SRS_PosResourcesPerBWP_PerSlot_r16 = v1
	}

	return nil
}
