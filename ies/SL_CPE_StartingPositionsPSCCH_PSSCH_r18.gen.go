// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-CPE-StartingPositionsPSCCH-PSSCH-r18 (line 28006).

var sLCPEStartingPositionsPSCCHPSSCHR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-Priority-r18"},
		{Name: "sl-CPE-StartingPositions-r18"},
	},
}

var sLCPEStartingPositionsPSCCHPSSCHR18SlPriorityR18Constraints = per.Constrained(1, 8)

var sLCPEStartingPositionsPSCCHPSSCHR18SlCPEStartingPositionsR18Constraints = per.SizeRange(1, 9)

type SL_CPE_StartingPositionsPSCCH_PSSCH_r18 struct {
	Sl_Priority_r18              int64
	Sl_CPE_StartingPositions_r18 []int64
}

func (ie *SL_CPE_StartingPositionsPSCCH_PSSCH_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLCPEStartingPositionsPSCCHPSSCHR18Constraints)
	_ = seq

	// 1. sl-Priority-r18: integer
	{
		if err := e.EncodeInteger(ie.Sl_Priority_r18, sLCPEStartingPositionsPSCCHPSSCHR18SlPriorityR18Constraints); err != nil {
			return err
		}
	}

	// 2. sl-CPE-StartingPositions-r18: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(sLCPEStartingPositionsPSCCHPSSCHR18SlCPEStartingPositionsR18Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Sl_CPE_StartingPositions_r18))); err != nil {
			return err
		}
		for i := range ie.Sl_CPE_StartingPositions_r18 {
			if err := e.EncodeInteger(ie.Sl_CPE_StartingPositions_r18[i], per.Constrained(1, 9)); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_CPE_StartingPositionsPSCCH_PSSCH_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLCPEStartingPositionsPSCCHPSSCHR18Constraints)
	_ = seq

	// 1. sl-Priority-r18: integer
	{
		v0, err := d.DecodeInteger(sLCPEStartingPositionsPSCCHPSSCHR18SlPriorityR18Constraints)
		if err != nil {
			return err
		}
		ie.Sl_Priority_r18 = v0
	}

	// 2. sl-CPE-StartingPositions-r18: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(sLCPEStartingPositionsPSCCHPSSCHR18SlCPEStartingPositionsR18Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Sl_CPE_StartingPositions_r18 = make([]int64, n)
		for i := int64(0); i < n; i++ {
			v, err := d.DecodeInteger(per.Constrained(1, 9))
			if err != nil {
				return err
			}
			ie.Sl_CPE_StartingPositions_r18[i] = v
		}
	}

	return nil
}
