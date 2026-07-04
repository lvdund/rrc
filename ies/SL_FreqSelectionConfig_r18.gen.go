// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-FreqSelectionConfig-r18 (line 27328).

var sLFreqSelectionConfigR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-priorityList-r18"},
		{Name: "sl-threshCBR-FreqReselection-r18"},
		{Name: "sl-threshCBR-FreqKeeping-r18"},
	},
}

var sLFreqSelectionConfigR18SlPriorityListR18Constraints = per.SizeRange(1, 8)

type SL_FreqSelectionConfig_r18 struct {
	Sl_PriorityList_r18              []int64
	Sl_ThreshCBR_FreqReselection_r18 SL_CBR_r16
	Sl_ThreshCBR_FreqKeeping_r18     SL_CBR_r16
}

func (ie *SL_FreqSelectionConfig_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLFreqSelectionConfigR18Constraints)
	_ = seq

	// 1. sl-priorityList-r18: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(sLFreqSelectionConfigR18SlPriorityListR18Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Sl_PriorityList_r18))); err != nil {
			return err
		}
		for i := range ie.Sl_PriorityList_r18 {
			if err := e.EncodeInteger(ie.Sl_PriorityList_r18[i], per.Constrained(1, 8)); err != nil {
				return err
			}
		}
	}

	// 2. sl-threshCBR-FreqReselection-r18: ref
	{
		if err := ie.Sl_ThreshCBR_FreqReselection_r18.Encode(e); err != nil {
			return err
		}
	}

	// 3. sl-threshCBR-FreqKeeping-r18: ref
	{
		if err := ie.Sl_ThreshCBR_FreqKeeping_r18.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SL_FreqSelectionConfig_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLFreqSelectionConfigR18Constraints)
	_ = seq

	// 1. sl-priorityList-r18: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(sLFreqSelectionConfigR18SlPriorityListR18Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Sl_PriorityList_r18 = make([]int64, n)
		for i := int64(0); i < n; i++ {
			v, err := d.DecodeInteger(per.Constrained(1, 8))
			if err != nil {
				return err
			}
			ie.Sl_PriorityList_r18[i] = v
		}
	}

	// 2. sl-threshCBR-FreqReselection-r18: ref
	{
		if err := ie.Sl_ThreshCBR_FreqReselection_r18.Decode(d); err != nil {
			return err
		}
	}

	// 3. sl-threshCBR-FreqKeeping-r18: ref
	{
		if err := ie.Sl_ThreshCBR_FreqKeeping_r18.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
