// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: AffectedCarrierFreqCombList-r16 (line 2511).
// AffectedCarrierFreqCombList-r16 ::= SEQUENCE (SIZE (1..maxCombIDC-r16)) OF AffectedCarrierFreqComb-r16

var affectedCarrierFreqCombListR16SizeConstraints = per.SizeRange(1, common.MaxCombIDC_r16)

type AffectedCarrierFreqCombList_r16 []AffectedCarrierFreqComb_r16

func (ie *AffectedCarrierFreqCombList_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(affectedCarrierFreqCombListR16SizeConstraints)
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

func (ie *AffectedCarrierFreqCombList_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(affectedCarrierFreqCombListR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(AffectedCarrierFreqCombList_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
