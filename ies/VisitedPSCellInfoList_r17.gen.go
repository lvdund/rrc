// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: VisitedPSCellInfoList-r17 (line 26717).
// VisitedPSCellInfoList-r17 ::= SEQUENCE (SIZE (1..maxPSCellHistory-r17)) OF VisitedPSCellInfo-r17

var visitedPSCellInfoListR17SizeConstraints = per.SizeRange(1, common.MaxPSCellHistory_r17)

type VisitedPSCellInfoList_r17 []VisitedPSCellInfo_r17

func (ie *VisitedPSCellInfoList_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(visitedPSCellInfoListR17SizeConstraints)
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

func (ie *VisitedPSCellInfoList_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(visitedPSCellInfoListR17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(VisitedPSCellInfoList_r17, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
