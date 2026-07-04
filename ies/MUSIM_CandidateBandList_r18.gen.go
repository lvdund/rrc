// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MUSIM-CandidateBandList-r18 (line 26404).
// MUSIM-CandidateBandList-r18::= SEQUENCE (SIZE (1..maxCandidateBandIndex-r18)) OF FreqBandIndicatorNR

var mUSIMCandidateBandListR18SizeConstraints = per.SizeRange(1, common.MaxCandidateBandIndex_r18)

type MUSIM_CandidateBandList_r18 []FreqBandIndicatorNR

func (ie *MUSIM_CandidateBandList_r18) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(mUSIMCandidateBandListR18SizeConstraints)
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

func (ie *MUSIM_CandidateBandList_r18) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(mUSIMCandidateBandListR18SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MUSIM_CandidateBandList_r18, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
