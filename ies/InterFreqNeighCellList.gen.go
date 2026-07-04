// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: InterFreqNeighCellList (line 4067).
// InterFreqNeighCellList ::=          SEQUENCE (SIZE (1..maxCellInter)) OF InterFreqNeighCellInfo

var interFreqNeighCellListSizeConstraints = per.SizeRange(1, common.MaxCellInter)

type InterFreqNeighCellList []InterFreqNeighCellInfo

func (ie *InterFreqNeighCellList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(interFreqNeighCellListSizeConstraints)
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

func (ie *InterFreqNeighCellList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(interFreqNeighCellListSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(InterFreqNeighCellList, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
