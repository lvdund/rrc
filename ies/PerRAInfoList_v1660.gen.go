// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PerRAInfoList-v1660 (line 3156).
// PerRAInfoList-v1660 ::= SEQUENCE (SIZE (1..200)) OF PerRACSI-RSInfo-v1660

var perRAInfoListV1660SizeConstraints = per.SizeRange(1, 200)

type PerRAInfoList_v1660 []PerRACSI_RSInfo_v1660

func (ie *PerRAInfoList_v1660) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(perRAInfoListV1660SizeConstraints)
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

func (ie *PerRAInfoList_v1660) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(perRAInfoListV1660SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(PerRAInfoList_v1660, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
