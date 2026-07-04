// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BFD-RelaxationReportingConfig-r17 (line 26474).

var bFDRelaxationReportingConfigR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "bfd-RelaxtionReportingProhibitTimer-r17"},
	},
}

const (
	BFD_RelaxationReportingConfig_r17_Bfd_RelaxtionReportingProhibitTimer_r17_S0       = 0
	BFD_RelaxationReportingConfig_r17_Bfd_RelaxtionReportingProhibitTimer_r17_S0dot5   = 1
	BFD_RelaxationReportingConfig_r17_Bfd_RelaxtionReportingProhibitTimer_r17_S1       = 2
	BFD_RelaxationReportingConfig_r17_Bfd_RelaxtionReportingProhibitTimer_r17_S2       = 3
	BFD_RelaxationReportingConfig_r17_Bfd_RelaxtionReportingProhibitTimer_r17_S5       = 4
	BFD_RelaxationReportingConfig_r17_Bfd_RelaxtionReportingProhibitTimer_r17_S10      = 5
	BFD_RelaxationReportingConfig_r17_Bfd_RelaxtionReportingProhibitTimer_r17_S20      = 6
	BFD_RelaxationReportingConfig_r17_Bfd_RelaxtionReportingProhibitTimer_r17_S30      = 7
	BFD_RelaxationReportingConfig_r17_Bfd_RelaxtionReportingProhibitTimer_r17_S60      = 8
	BFD_RelaxationReportingConfig_r17_Bfd_RelaxtionReportingProhibitTimer_r17_S90      = 9
	BFD_RelaxationReportingConfig_r17_Bfd_RelaxtionReportingProhibitTimer_r17_S120     = 10
	BFD_RelaxationReportingConfig_r17_Bfd_RelaxtionReportingProhibitTimer_r17_S300     = 11
	BFD_RelaxationReportingConfig_r17_Bfd_RelaxtionReportingProhibitTimer_r17_S600     = 12
	BFD_RelaxationReportingConfig_r17_Bfd_RelaxtionReportingProhibitTimer_r17_Infinity = 13
	BFD_RelaxationReportingConfig_r17_Bfd_RelaxtionReportingProhibitTimer_r17_Spare2   = 14
	BFD_RelaxationReportingConfig_r17_Bfd_RelaxtionReportingProhibitTimer_r17_Spare1   = 15
)

var bFDRelaxationReportingConfigR17BfdRelaxtionReportingProhibitTimerR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BFD_RelaxationReportingConfig_r17_Bfd_RelaxtionReportingProhibitTimer_r17_S0, BFD_RelaxationReportingConfig_r17_Bfd_RelaxtionReportingProhibitTimer_r17_S0dot5, BFD_RelaxationReportingConfig_r17_Bfd_RelaxtionReportingProhibitTimer_r17_S1, BFD_RelaxationReportingConfig_r17_Bfd_RelaxtionReportingProhibitTimer_r17_S2, BFD_RelaxationReportingConfig_r17_Bfd_RelaxtionReportingProhibitTimer_r17_S5, BFD_RelaxationReportingConfig_r17_Bfd_RelaxtionReportingProhibitTimer_r17_S10, BFD_RelaxationReportingConfig_r17_Bfd_RelaxtionReportingProhibitTimer_r17_S20, BFD_RelaxationReportingConfig_r17_Bfd_RelaxtionReportingProhibitTimer_r17_S30, BFD_RelaxationReportingConfig_r17_Bfd_RelaxtionReportingProhibitTimer_r17_S60, BFD_RelaxationReportingConfig_r17_Bfd_RelaxtionReportingProhibitTimer_r17_S90, BFD_RelaxationReportingConfig_r17_Bfd_RelaxtionReportingProhibitTimer_r17_S120, BFD_RelaxationReportingConfig_r17_Bfd_RelaxtionReportingProhibitTimer_r17_S300, BFD_RelaxationReportingConfig_r17_Bfd_RelaxtionReportingProhibitTimer_r17_S600, BFD_RelaxationReportingConfig_r17_Bfd_RelaxtionReportingProhibitTimer_r17_Infinity, BFD_RelaxationReportingConfig_r17_Bfd_RelaxtionReportingProhibitTimer_r17_Spare2, BFD_RelaxationReportingConfig_r17_Bfd_RelaxtionReportingProhibitTimer_r17_Spare1},
}

type BFD_RelaxationReportingConfig_r17 struct {
	Bfd_RelaxtionReportingProhibitTimer_r17 int64
}

func (ie *BFD_RelaxationReportingConfig_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bFDRelaxationReportingConfigR17Constraints)
	_ = seq

	// 1. bfd-RelaxtionReportingProhibitTimer-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.Bfd_RelaxtionReportingProhibitTimer_r17, bFDRelaxationReportingConfigR17BfdRelaxtionReportingProhibitTimerR17Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *BFD_RelaxationReportingConfig_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bFDRelaxationReportingConfigR17Constraints)
	_ = seq

	// 1. bfd-RelaxtionReportingProhibitTimer-r17: enumerated
	{
		v0, err := d.DecodeEnumerated(bFDRelaxationReportingConfigR17BfdRelaxtionReportingProhibitTimerR17Constraints)
		if err != nil {
			return err
		}
		ie.Bfd_RelaxtionReportingProhibitTimer_r17 = v0
	}

	return nil
}
