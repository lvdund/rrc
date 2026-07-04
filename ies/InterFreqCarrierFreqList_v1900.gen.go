// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: InterFreqCarrierFreqList-v1900 (line 3972).
// InterFreqCarrierFreqList-v1900 ::=  SEQUENCE (SIZE (1..maxFreq)) OF InterFreqCarrierFreqInfo-v1900

var interFreqCarrierFreqListV1900SizeConstraints = per.SizeRange(1, common.MaxFreq)

type InterFreqCarrierFreqList_v1900 []InterFreqCarrierFreqInfo_v1900

func (ie *InterFreqCarrierFreqList_v1900) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(interFreqCarrierFreqListV1900SizeConstraints)
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

func (ie *InterFreqCarrierFreqList_v1900) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(interFreqCarrierFreqListV1900SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(InterFreqCarrierFreqList_v1900, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
