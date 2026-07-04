// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NTN-CovEnh-r18 (line 4541).

var nTNCovEnhR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "numberOfMsg4HARQ-ACK-Repetitions-r18"},
		{Name: "rsrp-ThresholdMsg4HARQ-ACK-r18", Optional: true},
	},
}

var nTNCovEnhR18NumberOfMsg4HARQACKRepetitionsR18Constraints = per.FixedSize(4)

type NTN_CovEnh_r18 struct {
	NumberOfMsg4HARQ_ACK_Repetitions_r18 per.BitString
	Rsrp_ThresholdMsg4HARQ_ACK_r18       *RSRP_Range
}

func (ie *NTN_CovEnh_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nTNCovEnhR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Rsrp_ThresholdMsg4HARQ_ACK_r18 != nil}); err != nil {
		return err
	}

	// 2. numberOfMsg4HARQ-ACK-Repetitions-r18: bit-string
	{
		if err := e.EncodeBitString(ie.NumberOfMsg4HARQ_ACK_Repetitions_r18, nTNCovEnhR18NumberOfMsg4HARQACKRepetitionsR18Constraints); err != nil {
			return err
		}
	}

	// 3. rsrp-ThresholdMsg4HARQ-ACK-r18: ref
	{
		if ie.Rsrp_ThresholdMsg4HARQ_ACK_r18 != nil {
			if err := ie.Rsrp_ThresholdMsg4HARQ_ACK_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *NTN_CovEnh_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nTNCovEnhR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. numberOfMsg4HARQ-ACK-Repetitions-r18: bit-string
	{
		v0, err := d.DecodeBitString(nTNCovEnhR18NumberOfMsg4HARQACKRepetitionsR18Constraints)
		if err != nil {
			return err
		}
		ie.NumberOfMsg4HARQ_ACK_Repetitions_r18 = v0
	}

	// 3. rsrp-ThresholdMsg4HARQ-ACK-r18: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Rsrp_ThresholdMsg4HARQ_ACK_r18 = new(RSRP_Range)
			if err := ie.Rsrp_ThresholdMsg4HARQ_ACK_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
