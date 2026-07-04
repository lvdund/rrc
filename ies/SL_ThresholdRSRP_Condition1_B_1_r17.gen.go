// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-ThresholdRSRP-Condition1-B-1-r17 (line 27390).

var sLThresholdRSRPCondition1B1R17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-Priority-r17"},
		{Name: "sl-ThresholdRSRP-Condition1-B-1-r17"},
	},
}

var sLThresholdRSRPCondition1B1R17SlPriorityR17Constraints = per.Constrained(1, 8)

var sLThresholdRSRPCondition1B1R17SlThresholdRSRPCondition1B1R17Constraints = per.Constrained(0, 66)

type SL_ThresholdRSRP_Condition1_B_1_r17 struct {
	Sl_Priority_r17                     int64
	Sl_ThresholdRSRP_Condition1_B_1_r17 int64
}

func (ie *SL_ThresholdRSRP_Condition1_B_1_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLThresholdRSRPCondition1B1R17Constraints)
	_ = seq

	// 1. sl-Priority-r17: integer
	{
		if err := e.EncodeInteger(ie.Sl_Priority_r17, sLThresholdRSRPCondition1B1R17SlPriorityR17Constraints); err != nil {
			return err
		}
	}

	// 2. sl-ThresholdRSRP-Condition1-B-1-r17: integer
	{
		if err := e.EncodeInteger(ie.Sl_ThresholdRSRP_Condition1_B_1_r17, sLThresholdRSRPCondition1B1R17SlThresholdRSRPCondition1B1R17Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SL_ThresholdRSRP_Condition1_B_1_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLThresholdRSRPCondition1B1R17Constraints)
	_ = seq

	// 1. sl-Priority-r17: integer
	{
		v0, err := d.DecodeInteger(sLThresholdRSRPCondition1B1R17SlPriorityR17Constraints)
		if err != nil {
			return err
		}
		ie.Sl_Priority_r17 = v0
	}

	// 2. sl-ThresholdRSRP-Condition1-B-1-r17: integer
	{
		v1, err := d.DecodeInteger(sLThresholdRSRPCondition1B1R17SlThresholdRSRPCondition1B1R17Constraints)
		if err != nil {
			return err
		}
		ie.Sl_ThresholdRSRP_Condition1_B_1_r17 = v1
	}

	return nil
}
