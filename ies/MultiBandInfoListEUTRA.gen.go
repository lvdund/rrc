// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MultiBandInfoListEUTRA (line 9807).
// MultiBandInfoListEUTRA ::=              SEQUENCE (SIZE (1..maxMultiBands)) OF FreqBandIndicatorEUTRA

var multiBandInfoListEUTRASizeConstraints = per.SizeRange(1, common.MaxMultiBands)

type MultiBandInfoListEUTRA []FreqBandIndicatorEUTRA

func (ie *MultiBandInfoListEUTRA) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(multiBandInfoListEUTRASizeConstraints)
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

func (ie *MultiBandInfoListEUTRA) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(multiBandInfoListEUTRASizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MultiBandInfoListEUTRA, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
