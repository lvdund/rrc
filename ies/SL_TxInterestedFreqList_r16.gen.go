// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-TxInterestedFreqList-r16 (line 2275).
// SL-TxInterestedFreqList-r16 ::=        SEQUENCE (SIZE (1..maxNrofFreqSL-r16)) OF INTEGER (1..maxNrofFreqSL-r16)

var sLTxInterestedFreqListR16SizeConstraints = per.SizeRange(1, common.MaxNrofFreqSL_r16)

type SL_TxInterestedFreqList_r16 []int64

func (ie *SL_TxInterestedFreqList_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sLTxInterestedFreqListR16SizeConstraints)
	if err := seqOf.EncodeLength(int64(len(*ie))); err != nil {
		return err
	}
	for i := range *ie {
		if err := e.EncodeInteger((*ie)[i], per.Constrained(1, common.MaxNrofFreqSL_r16)); err != nil {
			return err
		}
	}
	return nil
}

func (ie *SL_TxInterestedFreqList_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sLTxInterestedFreqListR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SL_TxInterestedFreqList_r16, n)
	for i := int64(0); i < n; i++ {
		v, err := d.DecodeInteger(per.Constrained(1, common.MaxNrofFreqSL_r16))
		if err != nil {
			return err
		}
		(*ie)[i] = v
	}
	return nil
}
