// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MultiFrequencyBandListNR-SIB (line 10170).
// MultiFrequencyBandListNR-SIB ::=            SEQUENCE (SIZE (1.. maxNrofMultiBands)) OF NR-MultiBandInfo

var multiFrequencyBandListNRSIBSizeConstraints = per.SizeRange(1, common.MaxNrofMultiBands)

type MultiFrequencyBandListNR_SIB []NR_MultiBandInfo

func (ie *MultiFrequencyBandListNR_SIB) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(multiFrequencyBandListNRSIBSizeConstraints)
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

func (ie *MultiFrequencyBandListNR_SIB) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(multiFrequencyBandListNRSIBSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MultiFrequencyBandListNR_SIB, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
