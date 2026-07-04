// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: InterFreqNeighCellList-v1710 (line 4071).
// InterFreqNeighCellList-v1710 ::=    SEQUENCE (SIZE (1..maxCellInter)) OF InterFreqNeighCellInfo-v1710

var interFreqNeighCellListV1710SizeConstraints = per.SizeRange(1, common.MaxCellInter)

type InterFreqNeighCellList_v1710 []InterFreqNeighCellInfo_v1710

func (ie *InterFreqNeighCellList_v1710) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(interFreqNeighCellListV1710SizeConstraints)
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

func (ie *InterFreqNeighCellList_v1710) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(interFreqNeighCellListV1710SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(InterFreqNeighCellList_v1710, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
