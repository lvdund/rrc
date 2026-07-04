// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ThresholdMBS-r18 (line 684).

var thresholdMBSR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rsrp-r18", Optional: true},
		{Name: "rsrq-r18", Optional: true},
	},
}

type ThresholdMBS_r18 struct {
	Rsrp_r18 *RSRP_Range
	Rsrq_r18 *RSRQ_Range
}

func (ie *ThresholdMBS_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(thresholdMBSR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Rsrp_r18 != nil, ie.Rsrq_r18 != nil}); err != nil {
		return err
	}

	// 2. rsrp-r18: ref
	{
		if ie.Rsrp_r18 != nil {
			if err := ie.Rsrp_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. rsrq-r18: ref
	{
		if ie.Rsrq_r18 != nil {
			if err := ie.Rsrq_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *ThresholdMBS_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(thresholdMBSR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. rsrp-r18: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Rsrp_r18 = new(RSRP_Range)
			if err := ie.Rsrp_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. rsrq-r18: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Rsrq_r18 = new(RSRQ_Range)
			if err := ie.Rsrq_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
