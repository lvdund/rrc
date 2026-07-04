// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: FreqPriorityListEUTRA (line 1328).
// FreqPriorityListEUTRA ::=           SEQUENCE (SIZE (1..maxFreq)) OF FreqPriorityEUTRA

var freqPriorityListEUTRASizeConstraints = per.SizeRange(1, common.MaxFreq)

type FreqPriorityListEUTRA []FreqPriorityEUTRA

func (ie *FreqPriorityListEUTRA) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(freqPriorityListEUTRASizeConstraints)
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

func (ie *FreqPriorityListEUTRA) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(freqPriorityListEUTRASizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(FreqPriorityListEUTRA, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
