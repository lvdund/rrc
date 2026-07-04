// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: L1-MeasResultList-r19 (line 3538).
// L1-MeasResultList-r19 ::=            SEQUENCE (SIZE (1..maxCellReport)) OF L1-MeasResultPerCell-r19

var l1MeasResultListR19SizeConstraints = per.SizeRange(1, common.MaxCellReport)

type L1_MeasResultList_r19 []L1_MeasResultPerCell_r19

func (ie *L1_MeasResultList_r19) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(l1MeasResultListR19SizeConstraints)
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

func (ie *L1_MeasResultList_r19) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(l1MeasResultListR19SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(L1_MeasResultList_r19, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
