// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MUSIM-AvoidedBandsList-r18 (line 2665).
// MUSIM-AvoidedBandsList-r18 ::=          SEQUENCE (SIZE (1..maxBandComb-MUSIM-r18)) OF MUSIM-AvoidedBands-r18

var mUSIMAvoidedBandsListR18SizeConstraints = per.SizeRange(1, common.MaxBandComb_MUSIM_r18)

type MUSIM_AvoidedBandsList_r18 []MUSIM_AvoidedBands_r18

func (ie *MUSIM_AvoidedBandsList_r18) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(mUSIMAvoidedBandsListR18SizeConstraints)
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

func (ie *MUSIM_AvoidedBandsList_r18) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(mUSIMAvoidedBandsListR18SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MUSIM_AvoidedBandsList_r18, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
