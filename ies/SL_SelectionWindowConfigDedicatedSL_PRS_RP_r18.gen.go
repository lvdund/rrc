// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-SelectionWindowConfigDedicatedSL-PRS-RP-r18 (line 27694).

var sLSelectionWindowConfigDedicatedSLPRSRPR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-PRS-Priority-r18"},
		{Name: "sl-PRS-SelectionWindow-r18"},
	},
}

var sLSelectionWindowConfigDedicatedSLPRSRPR18SlPRSPriorityR18Constraints = per.Constrained(1, 8)

const (
	SL_SelectionWindowConfigDedicatedSL_PRS_RP_r18_Sl_PRS_SelectionWindow_r18_N1  = 0
	SL_SelectionWindowConfigDedicatedSL_PRS_RP_r18_Sl_PRS_SelectionWindow_r18_N5  = 1
	SL_SelectionWindowConfigDedicatedSL_PRS_RP_r18_Sl_PRS_SelectionWindow_r18_N10 = 2
	SL_SelectionWindowConfigDedicatedSL_PRS_RP_r18_Sl_PRS_SelectionWindow_r18_N20 = 3
)

var sLSelectionWindowConfigDedicatedSLPRSRPR18SlPRSSelectionWindowR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_SelectionWindowConfigDedicatedSL_PRS_RP_r18_Sl_PRS_SelectionWindow_r18_N1, SL_SelectionWindowConfigDedicatedSL_PRS_RP_r18_Sl_PRS_SelectionWindow_r18_N5, SL_SelectionWindowConfigDedicatedSL_PRS_RP_r18_Sl_PRS_SelectionWindow_r18_N10, SL_SelectionWindowConfigDedicatedSL_PRS_RP_r18_Sl_PRS_SelectionWindow_r18_N20},
}

type SL_SelectionWindowConfigDedicatedSL_PRS_RP_r18 struct {
	Sl_PRS_Priority_r18        int64
	Sl_PRS_SelectionWindow_r18 int64
}

func (ie *SL_SelectionWindowConfigDedicatedSL_PRS_RP_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLSelectionWindowConfigDedicatedSLPRSRPR18Constraints)
	_ = seq

	// 1. sl-PRS-Priority-r18: integer
	{
		if err := e.EncodeInteger(ie.Sl_PRS_Priority_r18, sLSelectionWindowConfigDedicatedSLPRSRPR18SlPRSPriorityR18Constraints); err != nil {
			return err
		}
	}

	// 2. sl-PRS-SelectionWindow-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.Sl_PRS_SelectionWindow_r18, sLSelectionWindowConfigDedicatedSLPRSRPR18SlPRSSelectionWindowR18Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SL_SelectionWindowConfigDedicatedSL_PRS_RP_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLSelectionWindowConfigDedicatedSLPRSRPR18Constraints)
	_ = seq

	// 1. sl-PRS-Priority-r18: integer
	{
		v0, err := d.DecodeInteger(sLSelectionWindowConfigDedicatedSLPRSRPR18SlPRSPriorityR18Constraints)
		if err != nil {
			return err
		}
		ie.Sl_PRS_Priority_r18 = v0
	}

	// 2. sl-PRS-SelectionWindow-r18: enumerated
	{
		v1, err := d.DecodeEnumerated(sLSelectionWindowConfigDedicatedSLPRSRPR18SlPRSSelectionWindowR18Constraints)
		if err != nil {
			return err
		}
		ie.Sl_PRS_SelectionWindow_r18 = v1
	}

	return nil
}
