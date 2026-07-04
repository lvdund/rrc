// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: NTN-NeighCellConfigListIoT-r19 (line 4793).
// NTN-NeighCellConfigListIoT-r19 ::=       SEQUENCE (SIZE(1..maxCellNTN-r17))  OF NTN-NeighCellConfigIoT-r19

var nTNNeighCellConfigListIoTR19SizeConstraints = per.SizeRange(1, common.MaxCellNTN_r17)

type NTN_NeighCellConfigListIoT_r19 []NTN_NeighCellConfigIoT_r19

func (ie *NTN_NeighCellConfigListIoT_r19) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(nTNNeighCellConfigListIoTR19SizeConstraints)
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

func (ie *NTN_NeighCellConfigListIoT_r19) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(nTNNeighCellConfigListIoTR19SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(NTN_NeighCellConfigListIoT_r19, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
