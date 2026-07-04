// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: NeedForNCSG-BandListNR-r17 (line 10508).
// NeedForNCSG-BandListNR-r17 ::=    SEQUENCE (SIZE (1..maxBands)) OF NeedForNCSG-NR-r17

var needForNCSGBandListNRR17SizeConstraints = per.SizeRange(1, common.MaxBands)

type NeedForNCSG_BandListNR_r17 []NeedForNCSG_NR_r17

func (ie *NeedForNCSG_BandListNR_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(needForNCSGBandListNRR17SizeConstraints)
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

func (ie *NeedForNCSG_BandListNR_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(needForNCSGBandListNRR17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(NeedForNCSG_BandListNR_r17, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
