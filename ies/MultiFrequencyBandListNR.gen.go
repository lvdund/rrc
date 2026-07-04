// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MultiFrequencyBandListNR (line 10165).
// MultiFrequencyBandListNR ::=        SEQUENCE (SIZE (1..maxNrofMultiBands)) OF FreqBandIndicatorNR

var multiFrequencyBandListNRSizeConstraints = per.SizeRange(1, common.MaxNrofMultiBands)

type MultiFrequencyBandListNR []FreqBandIndicatorNR

func (ie *MultiFrequencyBandListNR) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(multiFrequencyBandListNRSizeConstraints)
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

func (ie *MultiFrequencyBandListNR) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(multiFrequencyBandListNRSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MultiFrequencyBandListNR, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
