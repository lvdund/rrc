// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CombinationCarrierType-r18 (line 18349).

var combinationCarrierTypeR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "schedulingCellCarrierType-r18"},
		{Name: "scheduledCellCarrierType-r18"},
	},
}

const (
	CombinationCarrierType_r18_SchedulingCellCarrierType_r18_Licensed_Fdd_Fr1   = 0
	CombinationCarrierType_r18_SchedulingCellCarrierType_r18_Licensed_Tdd_Fr1   = 1
	CombinationCarrierType_r18_SchedulingCellCarrierType_r18_Unlicensed_Tdd_Fr1 = 2
	CombinationCarrierType_r18_SchedulingCellCarrierType_r18_Fr2_1              = 3
	CombinationCarrierType_r18_SchedulingCellCarrierType_r18_Fr2_2              = 4
)

var combinationCarrierTypeR18SchedulingCellCarrierTypeR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CombinationCarrierType_r18_SchedulingCellCarrierType_r18_Licensed_Fdd_Fr1, CombinationCarrierType_r18_SchedulingCellCarrierType_r18_Licensed_Tdd_Fr1, CombinationCarrierType_r18_SchedulingCellCarrierType_r18_Unlicensed_Tdd_Fr1, CombinationCarrierType_r18_SchedulingCellCarrierType_r18_Fr2_1, CombinationCarrierType_r18_SchedulingCellCarrierType_r18_Fr2_2},
}

const (
	CombinationCarrierType_r18_ScheduledCellCarrierType_r18_Licensed_Fdd_Fr1   = 0
	CombinationCarrierType_r18_ScheduledCellCarrierType_r18_Licensed_Tdd_Fr1   = 1
	CombinationCarrierType_r18_ScheduledCellCarrierType_r18_Unlicensed_Tdd_Fr1 = 2
	CombinationCarrierType_r18_ScheduledCellCarrierType_r18_Fr2_1              = 3
	CombinationCarrierType_r18_ScheduledCellCarrierType_r18_Fr2_2              = 4
)

var combinationCarrierTypeR18ScheduledCellCarrierTypeR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CombinationCarrierType_r18_ScheduledCellCarrierType_r18_Licensed_Fdd_Fr1, CombinationCarrierType_r18_ScheduledCellCarrierType_r18_Licensed_Tdd_Fr1, CombinationCarrierType_r18_ScheduledCellCarrierType_r18_Unlicensed_Tdd_Fr1, CombinationCarrierType_r18_ScheduledCellCarrierType_r18_Fr2_1, CombinationCarrierType_r18_ScheduledCellCarrierType_r18_Fr2_2},
}

type CombinationCarrierType_r18 struct {
	SchedulingCellCarrierType_r18 int64
	ScheduledCellCarrierType_r18  int64
}

func (ie *CombinationCarrierType_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(combinationCarrierTypeR18Constraints)
	_ = seq

	// 1. schedulingCellCarrierType-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.SchedulingCellCarrierType_r18, combinationCarrierTypeR18SchedulingCellCarrierTypeR18Constraints); err != nil {
			return err
		}
	}

	// 2. scheduledCellCarrierType-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.ScheduledCellCarrierType_r18, combinationCarrierTypeR18ScheduledCellCarrierTypeR18Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *CombinationCarrierType_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(combinationCarrierTypeR18Constraints)
	_ = seq

	// 1. schedulingCellCarrierType-r18: enumerated
	{
		v0, err := d.DecodeEnumerated(combinationCarrierTypeR18SchedulingCellCarrierTypeR18Constraints)
		if err != nil {
			return err
		}
		ie.SchedulingCellCarrierType_r18 = v0
	}

	// 2. scheduledCellCarrierType-r18: enumerated
	{
		v1, err := d.DecodeEnumerated(combinationCarrierTypeR18ScheduledCellCarrierTypeR18Constraints)
		if err != nil {
			return err
		}
		ie.ScheduledCellCarrierType_r18 = v1
	}

	return nil
}
