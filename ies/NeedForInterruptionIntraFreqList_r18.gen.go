// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: NeedForInterruptionIntraFreqList-r18 (line 10529).
// NeedForInterruptionIntraFreqList-r18 ::=          SEQUENCE (SIZE (1.. maxNrofServingCells)) OF NeedForInterruptionNR-r18

var needForInterruptionIntraFreqListR18SizeConstraints = per.SizeRange(1, common.MaxNrofServingCells)

type NeedForInterruptionIntraFreqList_r18 []NeedForInterruptionNR_r18

func (ie *NeedForInterruptionIntraFreqList_r18) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(needForInterruptionIntraFreqListR18SizeConstraints)
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

func (ie *NeedForInterruptionIntraFreqList_r18) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(needForInterruptionIntraFreqListR18SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(NeedForInterruptionIntraFreqList_r18, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
