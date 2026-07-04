// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: AffectedCarrierFreqRangeList-r18 (line 2737).
// AffectedCarrierFreqRangeList-r18 ::=  SEQUENCE (SIZE (1..maxFreqIDC-r16)) OF AffectedCarrierFreqRange-r18

var affectedCarrierFreqRangeListR18SizeConstraints = per.SizeRange(1, common.MaxFreqIDC_r16)

type AffectedCarrierFreqRangeList_r18 []AffectedCarrierFreqRange_r18

func (ie *AffectedCarrierFreqRangeList_r18) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(affectedCarrierFreqRangeListR18SizeConstraints)
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

func (ie *AffectedCarrierFreqRangeList_r18) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(affectedCarrierFreqRangeListR18SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(AffectedCarrierFreqRangeList_r18, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
