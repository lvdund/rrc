// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RefLocList-r19 (line 4546).
// RefLocList-r19 ::=                       SEQUENCE (SIZE(1..7)) OF Assisted-SSB-MTC-MG-RefLoc-r19

var refLocListR19SizeConstraints = per.SizeRange(1, 7)

type RefLocList_r19 []Assisted_SSB_MTC_MG_RefLoc_r19

func (ie *RefLocList_r19) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(refLocListR19SizeConstraints)
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

func (ie *RefLocList_r19) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(refLocListR19SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(RefLocList_r19, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
