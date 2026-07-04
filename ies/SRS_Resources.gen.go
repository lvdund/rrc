// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SRS-Resources (line 20478).

var sRSResourcesConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxNumberAperiodicSRS-PerBWP"},
		{Name: "maxNumberAperiodicSRS-PerBWP-PerSlot"},
		{Name: "maxNumberPeriodicSRS-PerBWP"},
		{Name: "maxNumberPeriodicSRS-PerBWP-PerSlot"},
		{Name: "maxNumberSemiPersistentSRS-PerBWP"},
		{Name: "maxNumberSemiPersistentSRS-PerBWP-PerSlot"},
		{Name: "maxNumberSRS-Ports-PerResource"},
	},
}

const (
	SRS_Resources_MaxNumberAperiodicSRS_PerBWP_N1  = 0
	SRS_Resources_MaxNumberAperiodicSRS_PerBWP_N2  = 1
	SRS_Resources_MaxNumberAperiodicSRS_PerBWP_N4  = 2
	SRS_Resources_MaxNumberAperiodicSRS_PerBWP_N8  = 3
	SRS_Resources_MaxNumberAperiodicSRS_PerBWP_N16 = 4
)

var sRSResourcesMaxNumberAperiodicSRSPerBWPConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_Resources_MaxNumberAperiodicSRS_PerBWP_N1, SRS_Resources_MaxNumberAperiodicSRS_PerBWP_N2, SRS_Resources_MaxNumberAperiodicSRS_PerBWP_N4, SRS_Resources_MaxNumberAperiodicSRS_PerBWP_N8, SRS_Resources_MaxNumberAperiodicSRS_PerBWP_N16},
}

var sRSResourcesMaxNumberAperiodicSRSPerBWPPerSlotConstraints = per.Constrained(1, 6)

const (
	SRS_Resources_MaxNumberPeriodicSRS_PerBWP_N1  = 0
	SRS_Resources_MaxNumberPeriodicSRS_PerBWP_N2  = 1
	SRS_Resources_MaxNumberPeriodicSRS_PerBWP_N4  = 2
	SRS_Resources_MaxNumberPeriodicSRS_PerBWP_N8  = 3
	SRS_Resources_MaxNumberPeriodicSRS_PerBWP_N16 = 4
)

var sRSResourcesMaxNumberPeriodicSRSPerBWPConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_Resources_MaxNumberPeriodicSRS_PerBWP_N1, SRS_Resources_MaxNumberPeriodicSRS_PerBWP_N2, SRS_Resources_MaxNumberPeriodicSRS_PerBWP_N4, SRS_Resources_MaxNumberPeriodicSRS_PerBWP_N8, SRS_Resources_MaxNumberPeriodicSRS_PerBWP_N16},
}

var sRSResourcesMaxNumberPeriodicSRSPerBWPPerSlotConstraints = per.Constrained(1, 6)

const (
	SRS_Resources_MaxNumberSemiPersistentSRS_PerBWP_N1  = 0
	SRS_Resources_MaxNumberSemiPersistentSRS_PerBWP_N2  = 1
	SRS_Resources_MaxNumberSemiPersistentSRS_PerBWP_N4  = 2
	SRS_Resources_MaxNumberSemiPersistentSRS_PerBWP_N8  = 3
	SRS_Resources_MaxNumberSemiPersistentSRS_PerBWP_N16 = 4
)

var sRSResourcesMaxNumberSemiPersistentSRSPerBWPConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_Resources_MaxNumberSemiPersistentSRS_PerBWP_N1, SRS_Resources_MaxNumberSemiPersistentSRS_PerBWP_N2, SRS_Resources_MaxNumberSemiPersistentSRS_PerBWP_N4, SRS_Resources_MaxNumberSemiPersistentSRS_PerBWP_N8, SRS_Resources_MaxNumberSemiPersistentSRS_PerBWP_N16},
}

var sRSResourcesMaxNumberSemiPersistentSRSPerBWPPerSlotConstraints = per.Constrained(1, 6)

const (
	SRS_Resources_MaxNumberSRS_Ports_PerResource_N1 = 0
	SRS_Resources_MaxNumberSRS_Ports_PerResource_N2 = 1
	SRS_Resources_MaxNumberSRS_Ports_PerResource_N4 = 2
)

var sRSResourcesMaxNumberSRSPortsPerResourceConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_Resources_MaxNumberSRS_Ports_PerResource_N1, SRS_Resources_MaxNumberSRS_Ports_PerResource_N2, SRS_Resources_MaxNumberSRS_Ports_PerResource_N4},
}

type SRS_Resources struct {
	MaxNumberAperiodicSRS_PerBWP              int64
	MaxNumberAperiodicSRS_PerBWP_PerSlot      int64
	MaxNumberPeriodicSRS_PerBWP               int64
	MaxNumberPeriodicSRS_PerBWP_PerSlot       int64
	MaxNumberSemiPersistentSRS_PerBWP         int64
	MaxNumberSemiPersistentSRS_PerBWP_PerSlot int64
	MaxNumberSRS_Ports_PerResource            int64
}

func (ie *SRS_Resources) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sRSResourcesConstraints)
	_ = seq

	// 1. maxNumberAperiodicSRS-PerBWP: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxNumberAperiodicSRS_PerBWP, sRSResourcesMaxNumberAperiodicSRSPerBWPConstraints); err != nil {
			return err
		}
	}

	// 2. maxNumberAperiodicSRS-PerBWP-PerSlot: integer
	{
		if err := e.EncodeInteger(ie.MaxNumberAperiodicSRS_PerBWP_PerSlot, sRSResourcesMaxNumberAperiodicSRSPerBWPPerSlotConstraints); err != nil {
			return err
		}
	}

	// 3. maxNumberPeriodicSRS-PerBWP: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxNumberPeriodicSRS_PerBWP, sRSResourcesMaxNumberPeriodicSRSPerBWPConstraints); err != nil {
			return err
		}
	}

	// 4. maxNumberPeriodicSRS-PerBWP-PerSlot: integer
	{
		if err := e.EncodeInteger(ie.MaxNumberPeriodicSRS_PerBWP_PerSlot, sRSResourcesMaxNumberPeriodicSRSPerBWPPerSlotConstraints); err != nil {
			return err
		}
	}

	// 5. maxNumberSemiPersistentSRS-PerBWP: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxNumberSemiPersistentSRS_PerBWP, sRSResourcesMaxNumberSemiPersistentSRSPerBWPConstraints); err != nil {
			return err
		}
	}

	// 6. maxNumberSemiPersistentSRS-PerBWP-PerSlot: integer
	{
		if err := e.EncodeInteger(ie.MaxNumberSemiPersistentSRS_PerBWP_PerSlot, sRSResourcesMaxNumberSemiPersistentSRSPerBWPPerSlotConstraints); err != nil {
			return err
		}
	}

	// 7. maxNumberSRS-Ports-PerResource: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxNumberSRS_Ports_PerResource, sRSResourcesMaxNumberSRSPortsPerResourceConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SRS_Resources) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sRSResourcesConstraints)
	_ = seq

	// 1. maxNumberAperiodicSRS-PerBWP: enumerated
	{
		v0, err := d.DecodeEnumerated(sRSResourcesMaxNumberAperiodicSRSPerBWPConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberAperiodicSRS_PerBWP = v0
	}

	// 2. maxNumberAperiodicSRS-PerBWP-PerSlot: integer
	{
		v1, err := d.DecodeInteger(sRSResourcesMaxNumberAperiodicSRSPerBWPPerSlotConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberAperiodicSRS_PerBWP_PerSlot = v1
	}

	// 3. maxNumberPeriodicSRS-PerBWP: enumerated
	{
		v2, err := d.DecodeEnumerated(sRSResourcesMaxNumberPeriodicSRSPerBWPConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberPeriodicSRS_PerBWP = v2
	}

	// 4. maxNumberPeriodicSRS-PerBWP-PerSlot: integer
	{
		v3, err := d.DecodeInteger(sRSResourcesMaxNumberPeriodicSRSPerBWPPerSlotConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberPeriodicSRS_PerBWP_PerSlot = v3
	}

	// 5. maxNumberSemiPersistentSRS-PerBWP: enumerated
	{
		v4, err := d.DecodeEnumerated(sRSResourcesMaxNumberSemiPersistentSRSPerBWPConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberSemiPersistentSRS_PerBWP = v4
	}

	// 6. maxNumberSemiPersistentSRS-PerBWP-PerSlot: integer
	{
		v5, err := d.DecodeInteger(sRSResourcesMaxNumberSemiPersistentSRSPerBWPPerSlotConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberSemiPersistentSRS_PerBWP_PerSlot = v5
	}

	// 7. maxNumberSRS-Ports-PerResource: enumerated
	{
		v6, err := d.DecodeEnumerated(sRSResourcesMaxNumberSRSPortsPerResourceConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberSRS_Ports_PerResource = v6
	}

	return nil
}
