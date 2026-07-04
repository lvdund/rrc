// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: BandCombinationList-v1560 (line 16548).
// BandCombinationList-v1560 ::=       SEQUENCE (SIZE (1..maxBandComb)) OF BandCombination-v1560

var bandCombinationListV1560SizeConstraints = per.SizeRange(1, common.MaxBandComb)

type BandCombinationList_v1560 []BandCombination_v1560

func (ie *BandCombinationList_v1560) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(bandCombinationListV1560SizeConstraints)
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

func (ie *BandCombinationList_v1560) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(bandCombinationListV1560SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(BandCombinationList_v1560, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
