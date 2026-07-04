// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: NeedForGapsIntraFreqList-r16 (line 10458).
// NeedForGapsIntraFreqList-r16 ::=          SEQUENCE (SIZE (1.. maxNrofServingCells)) OF NeedForGapsIntraFreq-r16

var needForGapsIntraFreqListR16SizeConstraints = per.SizeRange(1, common.MaxNrofServingCells)

type NeedForGapsIntraFreqList_r16 []NeedForGapsIntraFreq_r16

func (ie *NeedForGapsIntraFreqList_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(needForGapsIntraFreqListR16SizeConstraints)
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

func (ie *NeedForGapsIntraFreqList_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(needForGapsIntraFreqListR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(NeedForGapsIntraFreqList_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
