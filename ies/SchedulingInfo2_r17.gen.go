// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SchedulingInfo2-r17 (line 15081).

var schedulingInfo2R17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "si-BroadcastStatus-r17"},
		{Name: "si-WindowPosition-r17"},
		{Name: "si-Periodicity-r17"},
		{Name: "sib-MappingInfo-r17"},
	},
}

const (
	SchedulingInfo2_r17_Si_BroadcastStatus_r17_Broadcasting    = 0
	SchedulingInfo2_r17_Si_BroadcastStatus_r17_NotBroadcasting = 1
)

var schedulingInfo2R17SiBroadcastStatusR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SchedulingInfo2_r17_Si_BroadcastStatus_r17_Broadcasting, SchedulingInfo2_r17_Si_BroadcastStatus_r17_NotBroadcasting},
}

var schedulingInfo2R17SiWindowPositionR17Constraints = per.Constrained(1, 256)

const (
	SchedulingInfo2_r17_Si_Periodicity_r17_Rf8   = 0
	SchedulingInfo2_r17_Si_Periodicity_r17_Rf16  = 1
	SchedulingInfo2_r17_Si_Periodicity_r17_Rf32  = 2
	SchedulingInfo2_r17_Si_Periodicity_r17_Rf64  = 3
	SchedulingInfo2_r17_Si_Periodicity_r17_Rf128 = 4
	SchedulingInfo2_r17_Si_Periodicity_r17_Rf256 = 5
	SchedulingInfo2_r17_Si_Periodicity_r17_Rf512 = 6
)

var schedulingInfo2R17SiPeriodicityR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SchedulingInfo2_r17_Si_Periodicity_r17_Rf8, SchedulingInfo2_r17_Si_Periodicity_r17_Rf16, SchedulingInfo2_r17_Si_Periodicity_r17_Rf32, SchedulingInfo2_r17_Si_Periodicity_r17_Rf64, SchedulingInfo2_r17_Si_Periodicity_r17_Rf128, SchedulingInfo2_r17_Si_Periodicity_r17_Rf256, SchedulingInfo2_r17_Si_Periodicity_r17_Rf512},
}

type SchedulingInfo2_r17 struct {
	Si_BroadcastStatus_r17 int64
	Si_WindowPosition_r17  int64
	Si_Periodicity_r17     int64
	Sib_MappingInfo_r17    SIB_Mapping_v1700
}

func (ie *SchedulingInfo2_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(schedulingInfo2R17Constraints)
	_ = seq

	// 1. si-BroadcastStatus-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.Si_BroadcastStatus_r17, schedulingInfo2R17SiBroadcastStatusR17Constraints); err != nil {
			return err
		}
	}

	// 2. si-WindowPosition-r17: integer
	{
		if err := e.EncodeInteger(ie.Si_WindowPosition_r17, schedulingInfo2R17SiWindowPositionR17Constraints); err != nil {
			return err
		}
	}

	// 3. si-Periodicity-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.Si_Periodicity_r17, schedulingInfo2R17SiPeriodicityR17Constraints); err != nil {
			return err
		}
	}

	// 4. sib-MappingInfo-r17: ref
	{
		if err := ie.Sib_MappingInfo_r17.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SchedulingInfo2_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(schedulingInfo2R17Constraints)
	_ = seq

	// 1. si-BroadcastStatus-r17: enumerated
	{
		v0, err := d.DecodeEnumerated(schedulingInfo2R17SiBroadcastStatusR17Constraints)
		if err != nil {
			return err
		}
		ie.Si_BroadcastStatus_r17 = v0
	}

	// 2. si-WindowPosition-r17: integer
	{
		v1, err := d.DecodeInteger(schedulingInfo2R17SiWindowPositionR17Constraints)
		if err != nil {
			return err
		}
		ie.Si_WindowPosition_r17 = v1
	}

	// 3. si-Periodicity-r17: enumerated
	{
		v2, err := d.DecodeEnumerated(schedulingInfo2R17SiPeriodicityR17Constraints)
		if err != nil {
			return err
		}
		ie.Si_Periodicity_r17 = v2
	}

	// 4. sib-MappingInfo-r17: ref
	{
		if err := ie.Sib_MappingInfo_r17.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
