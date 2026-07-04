// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: AffectedCarrierFreqRangeCombList-r18 (line 2745).
// AffectedCarrierFreqRangeCombList-r18 ::= SEQUENCE (SIZE (1..maxCombIDC-r16)) OF AffectedCarrierFreqRangeComb-r18

var affectedCarrierFreqRangeCombListR18SizeConstraints = per.SizeRange(1, common.MaxCombIDC_r16)

type AffectedCarrierFreqRangeCombList_r18 []AffectedCarrierFreqRangeComb_r18

func (ie *AffectedCarrierFreqRangeCombList_r18) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(affectedCarrierFreqRangeCombListR18SizeConstraints)
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

func (ie *AffectedCarrierFreqRangeCombList_r18) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(affectedCarrierFreqRangeCombListR18SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(AffectedCarrierFreqRangeCombList_r18, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
