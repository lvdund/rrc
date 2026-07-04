// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: UplinkTxDirectCurrentList (line 16370).
// UplinkTxDirectCurrentList ::=           SEQUENCE (SIZE (1..maxNrofServingCells)) OF UplinkTxDirectCurrentCell

var uplinkTxDirectCurrentListSizeConstraints = per.SizeRange(1, common.MaxNrofServingCells)

type UplinkTxDirectCurrentList []UplinkTxDirectCurrentCell

func (ie *UplinkTxDirectCurrentList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(uplinkTxDirectCurrentListSizeConstraints)
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

func (ie *UplinkTxDirectCurrentList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(uplinkTxDirectCurrentListSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(UplinkTxDirectCurrentList, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
