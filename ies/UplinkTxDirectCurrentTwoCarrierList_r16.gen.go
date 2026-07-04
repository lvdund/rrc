// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: UplinkTxDirectCurrentTwoCarrierList-r16 (line 16419).
// UplinkTxDirectCurrentTwoCarrierList-r16 ::=   SEQUENCE (SIZE (1..maxNrofTxDC-TwoCarrier-r16)) OF UplinkTxDirectCurrentTwoCarrier-r16

var uplinkTxDirectCurrentTwoCarrierListR16SizeConstraints = per.SizeRange(1, common.MaxNrofTxDC_TwoCarrier_r16)

type UplinkTxDirectCurrentTwoCarrierList_r16 []UplinkTxDirectCurrentTwoCarrier_r16

func (ie *UplinkTxDirectCurrentTwoCarrierList_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(uplinkTxDirectCurrentTwoCarrierListR16SizeConstraints)
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

func (ie *UplinkTxDirectCurrentTwoCarrierList_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(uplinkTxDirectCurrentTwoCarrierListR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(UplinkTxDirectCurrentTwoCarrierList_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
