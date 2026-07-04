// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MUSIM-GapPriorityPreferenceList-r18 (line 2624).
// MUSIM-GapPriorityPreferenceList-r18 ::= SEQUENCE (SIZE (1..3)) OF GapPriority-r17

var mUSIMGapPriorityPreferenceListR18SizeConstraints = per.SizeRange(1, 3)

type MUSIM_GapPriorityPreferenceList_r18 []GapPriority_r17

func (ie *MUSIM_GapPriorityPreferenceList_r18) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(mUSIMGapPriorityPreferenceListR18SizeConstraints)
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

func (ie *MUSIM_GapPriorityPreferenceList_r18) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(mUSIMGapPriorityPreferenceListR18SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MUSIM_GapPriorityPreferenceList_r18, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
