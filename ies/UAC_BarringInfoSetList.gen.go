// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: UAC-BarringInfoSetList (line 16178).
// UAC-BarringInfoSetList ::=          SEQUENCE (SIZE(1..maxBarringInfoSet)) OF UAC-BarringInfoSet

var uACBarringInfoSetListSizeConstraints = per.SizeRange(1, common.MaxBarringInfoSet)

type UAC_BarringInfoSetList []UAC_BarringInfoSet

func (ie *UAC_BarringInfoSetList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(uACBarringInfoSetListSizeConstraints)
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

func (ie *UAC_BarringInfoSetList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(uACBarringInfoSetListSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(UAC_BarringInfoSetList, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
