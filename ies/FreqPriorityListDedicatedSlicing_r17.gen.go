// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: FreqPriorityListDedicatedSlicing-r17 (line 8356).
// FreqPriorityListDedicatedSlicing-r17 ::= SEQUENCE (SIZE (1.. maxFreq)) OF FreqPriorityDedicatedSlicing-r17

var freqPriorityListDedicatedSlicingR17SizeConstraints = per.SizeRange(1, common.MaxFreq)

type FreqPriorityListDedicatedSlicing_r17 []FreqPriorityDedicatedSlicing_r17

func (ie *FreqPriorityListDedicatedSlicing_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(freqPriorityListDedicatedSlicingR17SizeConstraints)
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

func (ie *FreqPriorityListDedicatedSlicing_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(freqPriorityListDedicatedSlicingR17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(FreqPriorityListDedicatedSlicing_r17, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
