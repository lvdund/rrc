// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: BT-NameList-r16 (line 26122).
// BT-NameList-r16 ::=                SEQUENCE (SIZE (1..maxBT-Name-r16)) OF BT-Name-r16

var bTNameListR16SizeConstraints = per.SizeRange(1, common.MaxBT_Name_r16)

type BT_NameList_r16 []BT_Name_r16

func (ie *BT_NameList_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(bTNameListR16SizeConstraints)
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

func (ie *BT_NameList_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(bTNameListR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(BT_NameList_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
