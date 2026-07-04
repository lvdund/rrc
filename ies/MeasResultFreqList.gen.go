// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasResultFreqList (line 1885).
// MeasResultFreqList ::=                   SEQUENCE (SIZE (1..maxFreq)) OF MeasResult2NR

var measResultFreqListSizeConstraints = per.SizeRange(1, common.MaxFreq)

type MeasResultFreqList []MeasResult2NR

func (ie *MeasResultFreqList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(measResultFreqListSizeConstraints)
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

func (ie *MeasResultFreqList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(measResultFreqListSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MeasResultFreqList, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
