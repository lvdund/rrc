// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SRS-PosResourceSetLinkedForAggBW-List-r18 (line 1129).
// SRS-PosResourceSetLinkedForAggBW-List-r18 ::= SEQUENCE (SIZE(2..maxNrOfLinkedSRS-PosResourceSet-r18)) OF SRS-PosResourceSetLinkedForAggBW-r18

var sRSPosResourceSetLinkedForAggBWListR18SizeConstraints = per.SizeRange(2, common.MaxNrOfLinkedSRS_PosResourceSet_r18)

type SRS_PosResourceSetLinkedForAggBW_List_r18 []SRS_PosResourceSetLinkedForAggBW_r18

func (ie *SRS_PosResourceSetLinkedForAggBW_List_r18) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sRSPosResourceSetLinkedForAggBWListR18SizeConstraints)
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

func (ie *SRS_PosResourceSetLinkedForAggBW_List_r18) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sRSPosResourceSetLinkedForAggBWListR18SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SRS_PosResourceSetLinkedForAggBW_List_r18, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
