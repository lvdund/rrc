// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-TxPercentageList-r16 (line 28089).
// SL-TxPercentageList-r16 ::=            SEQUENCE (SIZE (8)) OF SL-TxPercentageConfig-r16

var sLTxPercentageListR16SizeConstraints = per.FixedSize(8)

type SL_TxPercentageList_r16 []SL_TxPercentageConfig_r16

func (ie *SL_TxPercentageList_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sLTxPercentageListR16SizeConstraints)
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

func (ie *SL_TxPercentageList_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sLTxPercentageListR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SL_TxPercentageList_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
