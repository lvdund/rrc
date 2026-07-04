// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: LTE-NeighCellsCRS-AssistInfoList-r17 (line 8644).
// LTE-NeighCellsCRS-AssistInfoList-r17 ::= SEQUENCE (SIZE (1..maxNrofCRS-IM-InterfCell-r17)) OF LTE-NeighCellsCRS-AssistInfo-r17

var lTENeighCellsCRSAssistInfoListR17SizeConstraints = per.SizeRange(1, common.MaxNrofCRS_IM_InterfCell_r17)

type LTE_NeighCellsCRS_AssistInfoList_r17 []LTE_NeighCellsCRS_AssistInfo_r17

func (ie *LTE_NeighCellsCRS_AssistInfoList_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(lTENeighCellsCRSAssistInfoListR17SizeConstraints)
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

func (ie *LTE_NeighCellsCRS_AssistInfoList_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(lTENeighCellsCRSAssistInfoListR17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(LTE_NeighCellsCRS_AssistInfoList_r17, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
