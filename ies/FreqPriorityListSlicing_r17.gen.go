// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: FreqPriorityListSlicing-r17 (line 8374).
// FreqPriorityListSlicing-r17 ::= SEQUENCE (SIZE (1..maxFreqPlus1)) OF FreqPrioritySlicing-r17

var freqPriorityListSlicingR17SizeConstraints = per.SizeRange(1, common.MaxFreqPlus1)

type FreqPriorityListSlicing_r17 []FreqPrioritySlicing_r17

func (ie *FreqPriorityListSlicing_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(freqPriorityListSlicingR17SizeConstraints)
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

func (ie *FreqPriorityListSlicing_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(freqPriorityListSlicingR17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(FreqPriorityListSlicing_r17, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
