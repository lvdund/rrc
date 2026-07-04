// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MBS-FSAI-InterFreqList-r17 (line 4600).
// MBS-FSAI-InterFreqList-r17 ::= SEQUENCE (SIZE (1..maxFreq)) OF MBS-FSAI-InterFreq-r17

var mBSFSAIInterFreqListR17SizeConstraints = per.SizeRange(1, common.MaxFreq)

type MBS_FSAI_InterFreqList_r17 []MBS_FSAI_InterFreq_r17

func (ie *MBS_FSAI_InterFreqList_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(mBSFSAIInterFreqListR17SizeConstraints)
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

func (ie *MBS_FSAI_InterFreqList_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(mBSFSAIInterFreqListR17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MBS_FSAI_InterFreqList_r17, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
