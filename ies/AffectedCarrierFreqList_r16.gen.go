// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: AffectedCarrierFreqList-r16 (line 2504).
// AffectedCarrierFreqList-r16 ::= SEQUENCE (SIZE (1.. maxFreqIDC-r16)) OF AffectedCarrierFreq-r16

var affectedCarrierFreqListR16SizeConstraints = per.SizeRange(1, common.MaxFreqIDC_r16)

type AffectedCarrierFreqList_r16 []AffectedCarrierFreq_r16

func (ie *AffectedCarrierFreqList_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(affectedCarrierFreqListR16SizeConstraints)
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

func (ie *AffectedCarrierFreqList_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(affectedCarrierFreqListR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(AffectedCarrierFreqList_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
