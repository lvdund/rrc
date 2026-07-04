// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: UplinkTxDirectCurrentMoreCarrierList-r17 (line 16390).
// UplinkTxDirectCurrentMoreCarrierList-r17 ::=   SEQUENCE (SIZE (1..maxNrofCC-Group-r17)) OF CC-Group-r17

var uplinkTxDirectCurrentMoreCarrierListR17SizeConstraints = per.SizeRange(1, common.MaxNrofCC_Group_r17)

type UplinkTxDirectCurrentMoreCarrierList_r17 []CC_Group_r17

func (ie *UplinkTxDirectCurrentMoreCarrierList_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(uplinkTxDirectCurrentMoreCarrierListR17SizeConstraints)
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

func (ie *UplinkTxDirectCurrentMoreCarrierList_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(uplinkTxDirectCurrentMoreCarrierListR17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(UplinkTxDirectCurrentMoreCarrierList_r17, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
