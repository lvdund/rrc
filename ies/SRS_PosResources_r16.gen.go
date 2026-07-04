// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SRS-PosResources-r16 (line 20460).

var sRSPosResourcesR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxNumberSRS-PosResourceSetPerBWP-r16"},
		{Name: "maxNumberSRS-PosResourcesPerBWP-r16"},
		{Name: "maxNumberSRS-ResourcesPerBWP-PerSlot-r16"},
		{Name: "maxNumberPeriodicSRS-PosResourcesPerBWP-r16"},
		{Name: "maxNumberPeriodicSRS-PosResourcesPerBWP-PerSlot-r16"},
	},
}

const (
	SRS_PosResources_r16_MaxNumberSRS_PosResourceSetPerBWP_r16_N1  = 0
	SRS_PosResources_r16_MaxNumberSRS_PosResourceSetPerBWP_r16_N2  = 1
	SRS_PosResources_r16_MaxNumberSRS_PosResourceSetPerBWP_r16_N4  = 2
	SRS_PosResources_r16_MaxNumberSRS_PosResourceSetPerBWP_r16_N8  = 3
	SRS_PosResources_r16_MaxNumberSRS_PosResourceSetPerBWP_r16_N12 = 4
	SRS_PosResources_r16_MaxNumberSRS_PosResourceSetPerBWP_r16_N16 = 5
)

var sRSPosResourcesR16MaxNumberSRSPosResourceSetPerBWPR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_PosResources_r16_MaxNumberSRS_PosResourceSetPerBWP_r16_N1, SRS_PosResources_r16_MaxNumberSRS_PosResourceSetPerBWP_r16_N2, SRS_PosResources_r16_MaxNumberSRS_PosResourceSetPerBWP_r16_N4, SRS_PosResources_r16_MaxNumberSRS_PosResourceSetPerBWP_r16_N8, SRS_PosResources_r16_MaxNumberSRS_PosResourceSetPerBWP_r16_N12, SRS_PosResources_r16_MaxNumberSRS_PosResourceSetPerBWP_r16_N16},
}

const (
	SRS_PosResources_r16_MaxNumberSRS_PosResourcesPerBWP_r16_N1  = 0
	SRS_PosResources_r16_MaxNumberSRS_PosResourcesPerBWP_r16_N2  = 1
	SRS_PosResources_r16_MaxNumberSRS_PosResourcesPerBWP_r16_N4  = 2
	SRS_PosResources_r16_MaxNumberSRS_PosResourcesPerBWP_r16_N8  = 3
	SRS_PosResources_r16_MaxNumberSRS_PosResourcesPerBWP_r16_N16 = 4
	SRS_PosResources_r16_MaxNumberSRS_PosResourcesPerBWP_r16_N32 = 5
	SRS_PosResources_r16_MaxNumberSRS_PosResourcesPerBWP_r16_N64 = 6
)

var sRSPosResourcesR16MaxNumberSRSPosResourcesPerBWPR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_PosResources_r16_MaxNumberSRS_PosResourcesPerBWP_r16_N1, SRS_PosResources_r16_MaxNumberSRS_PosResourcesPerBWP_r16_N2, SRS_PosResources_r16_MaxNumberSRS_PosResourcesPerBWP_r16_N4, SRS_PosResources_r16_MaxNumberSRS_PosResourcesPerBWP_r16_N8, SRS_PosResources_r16_MaxNumberSRS_PosResourcesPerBWP_r16_N16, SRS_PosResources_r16_MaxNumberSRS_PosResourcesPerBWP_r16_N32, SRS_PosResources_r16_MaxNumberSRS_PosResourcesPerBWP_r16_N64},
}

const (
	SRS_PosResources_r16_MaxNumberSRS_ResourcesPerBWP_PerSlot_r16_N1  = 0
	SRS_PosResources_r16_MaxNumberSRS_ResourcesPerBWP_PerSlot_r16_N2  = 1
	SRS_PosResources_r16_MaxNumberSRS_ResourcesPerBWP_PerSlot_r16_N3  = 2
	SRS_PosResources_r16_MaxNumberSRS_ResourcesPerBWP_PerSlot_r16_N4  = 3
	SRS_PosResources_r16_MaxNumberSRS_ResourcesPerBWP_PerSlot_r16_N5  = 4
	SRS_PosResources_r16_MaxNumberSRS_ResourcesPerBWP_PerSlot_r16_N6  = 5
	SRS_PosResources_r16_MaxNumberSRS_ResourcesPerBWP_PerSlot_r16_N8  = 6
	SRS_PosResources_r16_MaxNumberSRS_ResourcesPerBWP_PerSlot_r16_N10 = 7
	SRS_PosResources_r16_MaxNumberSRS_ResourcesPerBWP_PerSlot_r16_N12 = 8
	SRS_PosResources_r16_MaxNumberSRS_ResourcesPerBWP_PerSlot_r16_N14 = 9
)

var sRSPosResourcesR16MaxNumberSRSResourcesPerBWPPerSlotR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_PosResources_r16_MaxNumberSRS_ResourcesPerBWP_PerSlot_r16_N1, SRS_PosResources_r16_MaxNumberSRS_ResourcesPerBWP_PerSlot_r16_N2, SRS_PosResources_r16_MaxNumberSRS_ResourcesPerBWP_PerSlot_r16_N3, SRS_PosResources_r16_MaxNumberSRS_ResourcesPerBWP_PerSlot_r16_N4, SRS_PosResources_r16_MaxNumberSRS_ResourcesPerBWP_PerSlot_r16_N5, SRS_PosResources_r16_MaxNumberSRS_ResourcesPerBWP_PerSlot_r16_N6, SRS_PosResources_r16_MaxNumberSRS_ResourcesPerBWP_PerSlot_r16_N8, SRS_PosResources_r16_MaxNumberSRS_ResourcesPerBWP_PerSlot_r16_N10, SRS_PosResources_r16_MaxNumberSRS_ResourcesPerBWP_PerSlot_r16_N12, SRS_PosResources_r16_MaxNumberSRS_ResourcesPerBWP_PerSlot_r16_N14},
}

const (
	SRS_PosResources_r16_MaxNumberPeriodicSRS_PosResourcesPerBWP_r16_N1  = 0
	SRS_PosResources_r16_MaxNumberPeriodicSRS_PosResourcesPerBWP_r16_N2  = 1
	SRS_PosResources_r16_MaxNumberPeriodicSRS_PosResourcesPerBWP_r16_N4  = 2
	SRS_PosResources_r16_MaxNumberPeriodicSRS_PosResourcesPerBWP_r16_N8  = 3
	SRS_PosResources_r16_MaxNumberPeriodicSRS_PosResourcesPerBWP_r16_N16 = 4
	SRS_PosResources_r16_MaxNumberPeriodicSRS_PosResourcesPerBWP_r16_N32 = 5
	SRS_PosResources_r16_MaxNumberPeriodicSRS_PosResourcesPerBWP_r16_N64 = 6
)

var sRSPosResourcesR16MaxNumberPeriodicSRSPosResourcesPerBWPR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_PosResources_r16_MaxNumberPeriodicSRS_PosResourcesPerBWP_r16_N1, SRS_PosResources_r16_MaxNumberPeriodicSRS_PosResourcesPerBWP_r16_N2, SRS_PosResources_r16_MaxNumberPeriodicSRS_PosResourcesPerBWP_r16_N4, SRS_PosResources_r16_MaxNumberPeriodicSRS_PosResourcesPerBWP_r16_N8, SRS_PosResources_r16_MaxNumberPeriodicSRS_PosResourcesPerBWP_r16_N16, SRS_PosResources_r16_MaxNumberPeriodicSRS_PosResourcesPerBWP_r16_N32, SRS_PosResources_r16_MaxNumberPeriodicSRS_PosResourcesPerBWP_r16_N64},
}

const (
	SRS_PosResources_r16_MaxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16_N1  = 0
	SRS_PosResources_r16_MaxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16_N2  = 1
	SRS_PosResources_r16_MaxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16_N3  = 2
	SRS_PosResources_r16_MaxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16_N4  = 3
	SRS_PosResources_r16_MaxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16_N5  = 4
	SRS_PosResources_r16_MaxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16_N6  = 5
	SRS_PosResources_r16_MaxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16_N8  = 6
	SRS_PosResources_r16_MaxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16_N10 = 7
	SRS_PosResources_r16_MaxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16_N12 = 8
	SRS_PosResources_r16_MaxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16_N14 = 9
)

var sRSPosResourcesR16MaxNumberPeriodicSRSPosResourcesPerBWPPerSlotR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_PosResources_r16_MaxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16_N1, SRS_PosResources_r16_MaxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16_N2, SRS_PosResources_r16_MaxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16_N3, SRS_PosResources_r16_MaxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16_N4, SRS_PosResources_r16_MaxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16_N5, SRS_PosResources_r16_MaxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16_N6, SRS_PosResources_r16_MaxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16_N8, SRS_PosResources_r16_MaxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16_N10, SRS_PosResources_r16_MaxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16_N12, SRS_PosResources_r16_MaxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16_N14},
}

type SRS_PosResources_r16 struct {
	MaxNumberSRS_PosResourceSetPerBWP_r16               int64
	MaxNumberSRS_PosResourcesPerBWP_r16                 int64
	MaxNumberSRS_ResourcesPerBWP_PerSlot_r16            int64
	MaxNumberPeriodicSRS_PosResourcesPerBWP_r16         int64
	MaxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16 int64
}

func (ie *SRS_PosResources_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sRSPosResourcesR16Constraints)
	_ = seq

	// 1. maxNumberSRS-PosResourceSetPerBWP-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxNumberSRS_PosResourceSetPerBWP_r16, sRSPosResourcesR16MaxNumberSRSPosResourceSetPerBWPR16Constraints); err != nil {
			return err
		}
	}

	// 2. maxNumberSRS-PosResourcesPerBWP-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxNumberSRS_PosResourcesPerBWP_r16, sRSPosResourcesR16MaxNumberSRSPosResourcesPerBWPR16Constraints); err != nil {
			return err
		}
	}

	// 3. maxNumberSRS-ResourcesPerBWP-PerSlot-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxNumberSRS_ResourcesPerBWP_PerSlot_r16, sRSPosResourcesR16MaxNumberSRSResourcesPerBWPPerSlotR16Constraints); err != nil {
			return err
		}
	}

	// 4. maxNumberPeriodicSRS-PosResourcesPerBWP-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxNumberPeriodicSRS_PosResourcesPerBWP_r16, sRSPosResourcesR16MaxNumberPeriodicSRSPosResourcesPerBWPR16Constraints); err != nil {
			return err
		}
	}

	// 5. maxNumberPeriodicSRS-PosResourcesPerBWP-PerSlot-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16, sRSPosResourcesR16MaxNumberPeriodicSRSPosResourcesPerBWPPerSlotR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SRS_PosResources_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sRSPosResourcesR16Constraints)
	_ = seq

	// 1. maxNumberSRS-PosResourceSetPerBWP-r16: enumerated
	{
		v0, err := d.DecodeEnumerated(sRSPosResourcesR16MaxNumberSRSPosResourceSetPerBWPR16Constraints)
		if err != nil {
			return err
		}
		ie.MaxNumberSRS_PosResourceSetPerBWP_r16 = v0
	}

	// 2. maxNumberSRS-PosResourcesPerBWP-r16: enumerated
	{
		v1, err := d.DecodeEnumerated(sRSPosResourcesR16MaxNumberSRSPosResourcesPerBWPR16Constraints)
		if err != nil {
			return err
		}
		ie.MaxNumberSRS_PosResourcesPerBWP_r16 = v1
	}

	// 3. maxNumberSRS-ResourcesPerBWP-PerSlot-r16: enumerated
	{
		v2, err := d.DecodeEnumerated(sRSPosResourcesR16MaxNumberSRSResourcesPerBWPPerSlotR16Constraints)
		if err != nil {
			return err
		}
		ie.MaxNumberSRS_ResourcesPerBWP_PerSlot_r16 = v2
	}

	// 4. maxNumberPeriodicSRS-PosResourcesPerBWP-r16: enumerated
	{
		v3, err := d.DecodeEnumerated(sRSPosResourcesR16MaxNumberPeriodicSRSPosResourcesPerBWPR16Constraints)
		if err != nil {
			return err
		}
		ie.MaxNumberPeriodicSRS_PosResourcesPerBWP_r16 = v3
	}

	// 5. maxNumberPeriodicSRS-PosResourcesPerBWP-PerSlot-r16: enumerated
	{
		v4, err := d.DecodeEnumerated(sRSPosResourcesR16MaxNumberPeriodicSRSPosResourcesPerBWPPerSlotR16Constraints)
		if err != nil {
			return err
		}
		ie.MaxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16 = v4
	}

	return nil
}
