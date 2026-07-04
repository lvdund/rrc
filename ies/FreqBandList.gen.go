// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: FreqBandList (line 20730).
// FreqBandList ::=                SEQUENCE (SIZE (1..maxBandsMRDC)) OF FreqBandInformation

var freqBandListSizeConstraints = per.SizeRange(1, common.MaxBandsMRDC)

type FreqBandList []FreqBandInformation

func (ie *FreqBandList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(freqBandListSizeConstraints)
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

func (ie *FreqBandList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(freqBandListSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(FreqBandList, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
