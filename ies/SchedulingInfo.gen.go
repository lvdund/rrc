// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SchedulingInfo (line 15060).

var schedulingInfoConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "si-BroadcastStatus"},
		{Name: "si-Periodicity"},
		{Name: "sib-MappingInfo"},
	},
}

const (
	SchedulingInfo_Si_BroadcastStatus_Broadcasting    = 0
	SchedulingInfo_Si_BroadcastStatus_NotBroadcasting = 1
)

var schedulingInfoSiBroadcastStatusConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SchedulingInfo_Si_BroadcastStatus_Broadcasting, SchedulingInfo_Si_BroadcastStatus_NotBroadcasting},
}

const (
	SchedulingInfo_Si_Periodicity_Rf8   = 0
	SchedulingInfo_Si_Periodicity_Rf16  = 1
	SchedulingInfo_Si_Periodicity_Rf32  = 2
	SchedulingInfo_Si_Periodicity_Rf64  = 3
	SchedulingInfo_Si_Periodicity_Rf128 = 4
	SchedulingInfo_Si_Periodicity_Rf256 = 5
	SchedulingInfo_Si_Periodicity_Rf512 = 6
)

var schedulingInfoSiPeriodicityConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SchedulingInfo_Si_Periodicity_Rf8, SchedulingInfo_Si_Periodicity_Rf16, SchedulingInfo_Si_Periodicity_Rf32, SchedulingInfo_Si_Periodicity_Rf64, SchedulingInfo_Si_Periodicity_Rf128, SchedulingInfo_Si_Periodicity_Rf256, SchedulingInfo_Si_Periodicity_Rf512},
}

type SchedulingInfo struct {
	Si_BroadcastStatus int64
	Si_Periodicity     int64
	Sib_MappingInfo    SIB_Mapping
}

func (ie *SchedulingInfo) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(schedulingInfoConstraints)
	_ = seq

	// 1. si-BroadcastStatus: enumerated
	{
		if err := e.EncodeEnumerated(ie.Si_BroadcastStatus, schedulingInfoSiBroadcastStatusConstraints); err != nil {
			return err
		}
	}

	// 2. si-Periodicity: enumerated
	{
		if err := e.EncodeEnumerated(ie.Si_Periodicity, schedulingInfoSiPeriodicityConstraints); err != nil {
			return err
		}
	}

	// 3. sib-MappingInfo: ref
	{
		if err := ie.Sib_MappingInfo.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SchedulingInfo) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(schedulingInfoConstraints)
	_ = seq

	// 1. si-BroadcastStatus: enumerated
	{
		v0, err := d.DecodeEnumerated(schedulingInfoSiBroadcastStatusConstraints)
		if err != nil {
			return err
		}
		ie.Si_BroadcastStatus = v0
	}

	// 2. si-Periodicity: enumerated
	{
		v1, err := d.DecodeEnumerated(schedulingInfoSiPeriodicityConstraints)
		if err != nil {
			return err
		}
		ie.Si_Periodicity = v1
	}

	// 3. sib-MappingInfo: ref
	{
		if err := ie.Sib_MappingInfo.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
