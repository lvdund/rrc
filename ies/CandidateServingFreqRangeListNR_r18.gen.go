// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CandidateServingFreqRangeListNR-r18 (line 26507).
// CandidateServingFreqRangeListNR-r18 ::= SEQUENCE (SIZE (1..maxFreqIDC-r16)) OF CandidateServingFreqRangeNR-r18

var candidateServingFreqRangeListNRR18SizeConstraints = per.SizeRange(1, common.MaxFreqIDC_r16)

type CandidateServingFreqRangeListNR_r18 []CandidateServingFreqRangeNR_r18

func (ie *CandidateServingFreqRangeListNR_r18) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(candidateServingFreqRangeListNRR18SizeConstraints)
	if err := seqOf.EncodeLength(int64(len(*ie))); err != nil {
		return err
	}
	for i := range *ie {
		if err := (*ie)[i].Encode(e); err != nil {
			return err
		}
	}
	return nil
}

func (ie *CandidateServingFreqRangeListNR_r18) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(candidateServingFreqRangeListNRR18SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(CandidateServingFreqRangeListNR_r18, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
