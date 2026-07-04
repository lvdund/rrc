// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: ATG-NeighCellConfigList-r18 (line 4620).
// ATG-NeighCellConfigList-r18 ::=       SEQUENCE (SIZE(1..maxCellATG-r18))  OF ATG-NeighCellConfig-r18

var aTGNeighCellConfigListR18SizeConstraints = per.SizeRange(1, common.MaxCellATG_r18)

type ATG_NeighCellConfigList_r18 []ATG_NeighCellConfig_r18

func (ie *ATG_NeighCellConfigList_r18) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(aTGNeighCellConfigListR18SizeConstraints)
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

func (ie *ATG_NeighCellConfigList_r18) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(aTGNeighCellConfigListR18SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(ATG_NeighCellConfigList_r18, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
