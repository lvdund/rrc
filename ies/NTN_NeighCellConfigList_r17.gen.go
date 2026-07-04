// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: NTN-NeighCellConfigList-r17 (line 4529).
// NTN-NeighCellConfigList-r17 ::=          SEQUENCE (SIZE(1..maxCellNTN-r17))  OF NTN-NeighCellConfig-r17

var nTNNeighCellConfigListR17SizeConstraints = per.SizeRange(1, common.MaxCellNTN_r17)

type NTN_NeighCellConfigList_r17 []NTN_NeighCellConfig_r17

func (ie *NTN_NeighCellConfigList_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(nTNNeighCellConfigListR17SizeConstraints)
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

func (ie *NTN_NeighCellConfigList_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(nTNNeighCellConfigListR17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(NTN_NeighCellConfigList_r17, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
