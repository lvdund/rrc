// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PerRAInfoList-v1800 (line 3163).
// PerRAInfoList-v1800 ::= SEQUENCE (SIZE (1..200)) OF PerRAInfo-v1800

var perRAInfoListV1800SizeConstraints = per.SizeRange(1, 200)

type PerRAInfoList_v1800 []PerRAInfo_v1800

func (ie *PerRAInfoList_v1800) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(perRAInfoListV1800SizeConstraints)
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

func (ie *PerRAInfoList_v1800) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(perRAInfoListV1800SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(PerRAInfoList_v1800, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
