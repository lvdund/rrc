// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ThresholdNR (line 9535).

var thresholdNRConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "thresholdRSRP", Optional: true},
		{Name: "thresholdRSRQ", Optional: true},
		{Name: "thresholdSINR", Optional: true},
	},
}

type ThresholdNR struct {
	ThresholdRSRP *RSRP_Range
	ThresholdRSRQ *RSRQ_Range
	ThresholdSINR *SINR_Range
}

func (ie *ThresholdNR) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(thresholdNRConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ThresholdRSRP != nil, ie.ThresholdRSRQ != nil, ie.ThresholdSINR != nil}); err != nil {
		return err
	}

	// 2. thresholdRSRP: ref
	{
		if ie.ThresholdRSRP != nil {
			if err := ie.ThresholdRSRP.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. thresholdRSRQ: ref
	{
		if ie.ThresholdRSRQ != nil {
			if err := ie.ThresholdRSRQ.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. thresholdSINR: ref
	{
		if ie.ThresholdSINR != nil {
			if err := ie.ThresholdSINR.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *ThresholdNR) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(thresholdNRConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. thresholdRSRP: ref
	{
		if seq.IsComponentPresent(0) {
			ie.ThresholdRSRP = new(RSRP_Range)
			if err := ie.ThresholdRSRP.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. thresholdRSRQ: ref
	{
		if seq.IsComponentPresent(1) {
			ie.ThresholdRSRQ = new(RSRQ_Range)
			if err := ie.ThresholdRSRQ.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. thresholdSINR: ref
	{
		if seq.IsComponentPresent(2) {
			ie.ThresholdSINR = new(SINR_Range)
			if err := ie.ThresholdSINR.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
