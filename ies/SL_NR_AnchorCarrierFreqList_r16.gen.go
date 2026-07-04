// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-NR-AnchorCarrierFreqList-r16 (line 4335).
// SL-NR-AnchorCarrierFreqList-r16 ::=  SEQUENCE (SIZE (1..maxFreqSL-NR-r16)) OF ARFCN-ValueNR

var sLNRAnchorCarrierFreqListR16SizeConstraints = per.SizeRange(1, common.MaxFreqSL_NR_r16)

type SL_NR_AnchorCarrierFreqList_r16 []ARFCN_ValueNR

func (ie *SL_NR_AnchorCarrierFreqList_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sLNRAnchorCarrierFreqListR16SizeConstraints)
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

func (ie *SL_NR_AnchorCarrierFreqList_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sLNRAnchorCarrierFreqListR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SL_NR_AnchorCarrierFreqList_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
