// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: InterFreqCarrierFreqList-v1760 (line 3968).
// InterFreqCarrierFreqList-v1760 ::=  SEQUENCE (SIZE (1..maxFreq)) OF InterFreqCarrierFreqInfo-v1760

var interFreqCarrierFreqListV1760SizeConstraints = per.SizeRange(1, common.MaxFreq)

type InterFreqCarrierFreqList_v1760 []InterFreqCarrierFreqInfo_v1760

func (ie *InterFreqCarrierFreqList_v1760) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(interFreqCarrierFreqListV1760SizeConstraints)
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

func (ie *InterFreqCarrierFreqList_v1760) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(interFreqCarrierFreqListV1760SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(InterFreqCarrierFreqList_v1760, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
