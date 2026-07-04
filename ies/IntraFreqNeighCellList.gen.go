// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: IntraFreqNeighCellList (line 3894).
// IntraFreqNeighCellList ::=          SEQUENCE (SIZE (1..maxCellIntra)) OF IntraFreqNeighCellInfo

var intraFreqNeighCellListSizeConstraints = per.SizeRange(1, common.MaxCellIntra)

type IntraFreqNeighCellList []IntraFreqNeighCellInfo

func (ie *IntraFreqNeighCellList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(intraFreqNeighCellListSizeConstraints)
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

func (ie *IntraFreqNeighCellList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(intraFreqNeighCellListSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(IntraFreqNeighCellList, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
