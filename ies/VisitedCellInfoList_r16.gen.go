// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: VisitedCellInfoList-r16 (line 26697).
// VisitedCellInfoList-r16 ::= SEQUENCE (SIZE (1..maxCellHistory-r16)) OF VisitedCellInfo-r16

var visitedCellInfoListR16SizeConstraints = per.SizeRange(1, common.MaxCellHistory_r16)

type VisitedCellInfoList_r16 []VisitedCellInfo_r16

func (ie *VisitedCellInfoList_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(visitedCellInfoListR16SizeConstraints)
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

func (ie *VisitedCellInfoList_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(visitedCellInfoListR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(VisitedCellInfoList_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
