// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CandidateBeamRSListExt-r16 (line 5153).
// CandidateBeamRSListExt-r16::=       SEQUENCE (SIZE(1.. maxNrofCandidateBeamsExt-r16)) OF PRACH-ResourceDedicatedBFR

var candidateBeamRSListExtR16SizeConstraints = per.SizeRange(1, common.MaxNrofCandidateBeamsExt_r16)

type CandidateBeamRSListExt_r16 []PRACH_ResourceDedicatedBFR

func (ie *CandidateBeamRSListExt_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(candidateBeamRSListExtR16SizeConstraints)
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

func (ie *CandidateBeamRSListExt_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(candidateBeamRSListExtR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(CandidateBeamRSListExt_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
