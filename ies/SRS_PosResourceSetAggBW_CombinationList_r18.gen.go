// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SRS-PosResourceSetAggBW-CombinationList-r18 (line 1127).
// SRS-PosResourceSetAggBW-CombinationList-r18 ::= SEQUENCE (SIZE(1.. maxNrOfLinkedSRS-PosResSetComb-r18)) OF SRS-PosResourceSetLinkedForAggBW-List-r18

var sRSPosResourceSetAggBWCombinationListR18SizeConstraints = per.SizeRange(1, common.MaxNrOfLinkedSRS_PosResSetComb_r18)

type SRS_PosResourceSetAggBW_CombinationList_r18 []SRS_PosResourceSetLinkedForAggBW_List_r18

func (ie *SRS_PosResourceSetAggBW_CombinationList_r18) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sRSPosResourceSetAggBWCombinationListR18SizeConstraints)
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

func (ie *SRS_PosResourceSetAggBW_CombinationList_r18) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sRSPosResourceSetAggBWCombinationListR18SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SRS_PosResourceSetAggBW_CombinationList_r18, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
