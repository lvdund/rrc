// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-CBR-PriorityTxConfigList-r16 (line 26888).
// SL-CBR-PriorityTxConfigList-r16 ::= SEQUENCE (SIZE (1..8)) OF SL-PriorityTxConfigIndex-r16

var sLCBRPriorityTxConfigListR16SizeConstraints = per.SizeRange(1, 8)

type SL_CBR_PriorityTxConfigList_r16 []SL_PriorityTxConfigIndex_r16

func (ie *SL_CBR_PriorityTxConfigList_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sLCBRPriorityTxConfigListR16SizeConstraints)
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

func (ie *SL_CBR_PriorityTxConfigList_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sLCBRPriorityTxConfigListR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SL_CBR_PriorityTxConfigList_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
