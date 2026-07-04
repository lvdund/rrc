// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SRS-InactivePosResourceSetLinkedForAggBW-List-r18 (line 1507).

var sRSInactivePosResourceSetLinkedForAggBWListR18SizeConstraints = per.SizeRange(2, common.MaxNrOfLinkedSRS_PosResourceSet_r18)

type SRS_InactivePosResourceSetLinkedForAggBW_List_r18 []SRS_PosResourceSetLinkedForAggBW_r18

func (ie *SRS_InactivePosResourceSetLinkedForAggBW_List_r18) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sRSInactivePosResourceSetLinkedForAggBWListR18SizeConstraints)
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

func (ie *SRS_InactivePosResourceSetLinkedForAggBW_List_r18) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sRSInactivePosResourceSetLinkedForAggBWListR18SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SRS_InactivePosResourceSetLinkedForAggBW_List_r18, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
