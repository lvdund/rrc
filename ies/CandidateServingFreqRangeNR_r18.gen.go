// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CandidateServingFreqRangeNR-r18 (line 26509).

var candidateServingFreqRangeNRR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "candidateCenterFreq-r18"},
		{Name: "candidateBandwidth-r18", Optional: true},
	},
}

const (
	CandidateServingFreqRangeNR_r18_CandidateBandwidth_r18_Khz200 = 0
	CandidateServingFreqRangeNR_r18_CandidateBandwidth_r18_Khz400 = 1
	CandidateServingFreqRangeNR_r18_CandidateBandwidth_r18_Khz600 = 2
	CandidateServingFreqRangeNR_r18_CandidateBandwidth_r18_Khz800 = 3
	CandidateServingFreqRangeNR_r18_CandidateBandwidth_r18_Mhz1   = 4
	CandidateServingFreqRangeNR_r18_CandidateBandwidth_r18_Mhz2   = 5
	CandidateServingFreqRangeNR_r18_CandidateBandwidth_r18_Mhz3   = 6
	CandidateServingFreqRangeNR_r18_CandidateBandwidth_r18_Mhz4   = 7
	CandidateServingFreqRangeNR_r18_CandidateBandwidth_r18_Mhz5   = 8
	CandidateServingFreqRangeNR_r18_CandidateBandwidth_r18_Mhz6   = 9
	CandidateServingFreqRangeNR_r18_CandidateBandwidth_r18_Mhz8   = 10
	CandidateServingFreqRangeNR_r18_CandidateBandwidth_r18_Mhz10  = 11
	CandidateServingFreqRangeNR_r18_CandidateBandwidth_r18_Mhz20  = 12
	CandidateServingFreqRangeNR_r18_CandidateBandwidth_r18_Mhz30  = 13
	CandidateServingFreqRangeNR_r18_CandidateBandwidth_r18_Mhz40  = 14
	CandidateServingFreqRangeNR_r18_CandidateBandwidth_r18_Mhz50  = 15
	CandidateServingFreqRangeNR_r18_CandidateBandwidth_r18_Mhz60  = 16
	CandidateServingFreqRangeNR_r18_CandidateBandwidth_r18_Mhz80  = 17
	CandidateServingFreqRangeNR_r18_CandidateBandwidth_r18_Mhz100 = 18
	CandidateServingFreqRangeNR_r18_CandidateBandwidth_r18_Mhz200 = 19
	CandidateServingFreqRangeNR_r18_CandidateBandwidth_r18_Mhz300 = 20
	CandidateServingFreqRangeNR_r18_CandidateBandwidth_r18_Mhz400 = 21
)

var candidateServingFreqRangeNRR18CandidateBandwidthR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CandidateServingFreqRangeNR_r18_CandidateBandwidth_r18_Khz200, CandidateServingFreqRangeNR_r18_CandidateBandwidth_r18_Khz400, CandidateServingFreqRangeNR_r18_CandidateBandwidth_r18_Khz600, CandidateServingFreqRangeNR_r18_CandidateBandwidth_r18_Khz800, CandidateServingFreqRangeNR_r18_CandidateBandwidth_r18_Mhz1, CandidateServingFreqRangeNR_r18_CandidateBandwidth_r18_Mhz2, CandidateServingFreqRangeNR_r18_CandidateBandwidth_r18_Mhz3, CandidateServingFreqRangeNR_r18_CandidateBandwidth_r18_Mhz4, CandidateServingFreqRangeNR_r18_CandidateBandwidth_r18_Mhz5, CandidateServingFreqRangeNR_r18_CandidateBandwidth_r18_Mhz6, CandidateServingFreqRangeNR_r18_CandidateBandwidth_r18_Mhz8, CandidateServingFreqRangeNR_r18_CandidateBandwidth_r18_Mhz10, CandidateServingFreqRangeNR_r18_CandidateBandwidth_r18_Mhz20, CandidateServingFreqRangeNR_r18_CandidateBandwidth_r18_Mhz30, CandidateServingFreqRangeNR_r18_CandidateBandwidth_r18_Mhz40, CandidateServingFreqRangeNR_r18_CandidateBandwidth_r18_Mhz50, CandidateServingFreqRangeNR_r18_CandidateBandwidth_r18_Mhz60, CandidateServingFreqRangeNR_r18_CandidateBandwidth_r18_Mhz80, CandidateServingFreqRangeNR_r18_CandidateBandwidth_r18_Mhz100, CandidateServingFreqRangeNR_r18_CandidateBandwidth_r18_Mhz200, CandidateServingFreqRangeNR_r18_CandidateBandwidth_r18_Mhz300, CandidateServingFreqRangeNR_r18_CandidateBandwidth_r18_Mhz400},
}

type CandidateServingFreqRangeNR_r18 struct {
	CandidateCenterFreq_r18 ARFCN_ValueNR
	CandidateBandwidth_r18  *int64
}

func (ie *CandidateServingFreqRangeNR_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(candidateServingFreqRangeNRR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.CandidateBandwidth_r18 != nil}); err != nil {
		return err
	}

	// 2. candidateCenterFreq-r18: ref
	{
		if err := ie.CandidateCenterFreq_r18.Encode(e); err != nil {
			return err
		}
	}

	// 3. candidateBandwidth-r18: enumerated
	{
		if ie.CandidateBandwidth_r18 != nil {
			if err := e.EncodeEnumerated(*ie.CandidateBandwidth_r18, candidateServingFreqRangeNRR18CandidateBandwidthR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CandidateServingFreqRangeNR_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(candidateServingFreqRangeNRR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. candidateCenterFreq-r18: ref
	{
		if err := ie.CandidateCenterFreq_r18.Decode(d); err != nil {
			return err
		}
	}

	// 3. candidateBandwidth-r18: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(candidateServingFreqRangeNRR18CandidateBandwidthR18Constraints)
			if err != nil {
				return err
			}
			ie.CandidateBandwidth_r18 = &idx
		}
	}

	return nil
}
