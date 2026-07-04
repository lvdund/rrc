// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PosMeasGapPreConfigToAddModList-r17 (line 9194).
// PosMeasGapPreConfigToAddModList-r17 ::= SEQUENCE (SIZE (1..maxNrofPreConfigPosGapId-r17)) OF PosGapConfig-r17

var posMeasGapPreConfigToAddModListR17SizeConstraints = per.SizeRange(1, common.MaxNrofPreConfigPosGapId_r17)

type PosMeasGapPreConfigToAddModList_r17 []PosGapConfig_r17

func (ie *PosMeasGapPreConfigToAddModList_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(posMeasGapPreConfigToAddModListR17SizeConstraints)
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

func (ie *PosMeasGapPreConfigToAddModList_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(posMeasGapPreConfigToAddModListR17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(PosMeasGapPreConfigToAddModList_r17, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
