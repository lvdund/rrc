// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: InterFreqCarrierFreqList-v1800 (line 3970).
// InterFreqCarrierFreqList-v1800 ::=  SEQUENCE (SIZE (1..maxFreq)) OF InterFreqCarrierFreqInfo-v1800

var interFreqCarrierFreqListV1800SizeConstraints = per.SizeRange(1, common.MaxFreq)

type InterFreqCarrierFreqList_v1800 []InterFreqCarrierFreqInfo_v1800

func (ie *InterFreqCarrierFreqList_v1800) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(interFreqCarrierFreqListV1800SizeConstraints)
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

func (ie *InterFreqCarrierFreqList_v1800) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(interFreqCarrierFreqListV1800SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(InterFreqCarrierFreqList_v1800, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
