// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: FreqPriorityListNR (line 1330).
// FreqPriorityListNR ::=              SEQUENCE (SIZE (1..maxFreq)) OF FreqPriorityNR

var freqPriorityListNRSizeConstraints = per.SizeRange(1, common.MaxFreq)

type FreqPriorityListNR []FreqPriorityNR

func (ie *FreqPriorityListNR) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(freqPriorityListNRSizeConstraints)
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

func (ie *FreqPriorityListNR) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(freqPriorityListNRSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(FreqPriorityListNR, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
