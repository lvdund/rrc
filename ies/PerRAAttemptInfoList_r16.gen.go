// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PerRAAttemptInfoList-r16 (line 3197).
// PerRAAttemptInfoList-r16 ::=         SEQUENCE (SIZE (1..200)) OF PerRAAttemptInfo-r16

var perRAAttemptInfoListR16SizeConstraints = per.SizeRange(1, 200)

type PerRAAttemptInfoList_r16 []PerRAAttemptInfo_r16

func (ie *PerRAAttemptInfoList_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(perRAAttemptInfoListR16SizeConstraints)
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

func (ie *PerRAAttemptInfoList_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(perRAAttemptInfoListR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(PerRAAttemptInfoList_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
