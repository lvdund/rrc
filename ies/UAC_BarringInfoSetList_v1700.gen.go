// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: UAC-BarringInfoSetList-v1700 (line 16180).
// UAC-BarringInfoSetList-v1700 ::= SEQUENCE (SIZE(1..maxBarringInfoSet)) OF UAC-BarringInfoSet-v1700

var uACBarringInfoSetListV1700SizeConstraints = per.SizeRange(1, common.MaxBarringInfoSet)

type UAC_BarringInfoSetList_v1700 []UAC_BarringInfoSet_v1700

func (ie *UAC_BarringInfoSetList_v1700) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(uACBarringInfoSetListV1700SizeConstraints)
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

func (ie *UAC_BarringInfoSetList_v1700) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(uACBarringInfoSetListV1700SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(UAC_BarringInfoSetList_v1700, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
