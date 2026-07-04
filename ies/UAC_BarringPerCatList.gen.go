// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: UAC-BarringPerCatList (line 16197).
// UAC-BarringPerCatList ::=           SEQUENCE (SIZE (1..maxAccessCat-1)) OF UAC-BarringPerCat

var uACBarringPerCatListSizeConstraints = per.SizeRange(1, common.MaxAccessCat_1)

type UAC_BarringPerCatList []UAC_BarringPerCat

func (ie *UAC_BarringPerCatList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(uACBarringPerCatListSizeConstraints)
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

func (ie *UAC_BarringPerCatList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(uACBarringPerCatListSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(UAC_BarringPerCatList, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
