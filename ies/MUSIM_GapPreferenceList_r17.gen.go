// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MUSIM-GapPreferenceList-r17 (line 2613).
// MUSIM-GapPreferenceList-r17 ::= SEQUENCE (SIZE (1..4)) OF MUSIM-GapInfo-r17

var mUSIMGapPreferenceListR17SizeConstraints = per.SizeRange(1, 4)

type MUSIM_GapPreferenceList_r17 []MUSIM_GapInfo_r17

func (ie *MUSIM_GapPreferenceList_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(mUSIMGapPreferenceListR17SizeConstraints)
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

func (ie *MUSIM_GapPreferenceList_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(mUSIMGapPreferenceListR17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MUSIM_GapPreferenceList_r17, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
