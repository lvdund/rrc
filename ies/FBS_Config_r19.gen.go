// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FBS-Config-r19 (line 9133).

var fBSConfigR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "fbs-ThresholdP-r19", Optional: true},
		{Name: "fbs-ThresholdQ-r19", Optional: true},
	},
}

type FBS_Config_r19 struct {
	Fbs_ThresholdP_r19 *RSRP_Range
	Fbs_ThresholdQ_r19 *RSRQ_Range
}

func (ie *FBS_Config_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(fBSConfigR19Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Fbs_ThresholdP_r19 != nil, ie.Fbs_ThresholdQ_r19 != nil}); err != nil {
		return err
	}

	// 2. fbs-ThresholdP-r19: ref
	{
		if ie.Fbs_ThresholdP_r19 != nil {
			if err := ie.Fbs_ThresholdP_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. fbs-ThresholdQ-r19: ref
	{
		if ie.Fbs_ThresholdQ_r19 != nil {
			if err := ie.Fbs_ThresholdQ_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FBS_Config_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(fBSConfigR19Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. fbs-ThresholdP-r19: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Fbs_ThresholdP_r19 = new(RSRP_Range)
			if err := ie.Fbs_ThresholdP_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. fbs-ThresholdQ-r19: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Fbs_ThresholdQ_r19 = new(RSRQ_Range)
			if err := ie.Fbs_ThresholdQ_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
