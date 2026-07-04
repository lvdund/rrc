// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: BandCombinationList-v15g0 (line 16556).
// BandCombinationList-v15g0 ::=       SEQUENCE (SIZE (1..maxBandComb)) OF BandCombination-v15g0

var bandCombinationListV15g0SizeConstraints = per.SizeRange(1, common.MaxBandComb)

type BandCombinationList_V15g0 []BandCombination_V15g0

func (ie *BandCombinationList_V15g0) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(bandCombinationListV15g0SizeConstraints)
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

func (ie *BandCombinationList_V15g0) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(bandCombinationListV15g0SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(BandCombinationList_V15g0, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
