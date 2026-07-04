// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CHO-WithCandidateSCGInfoList-r19 (line 3546).
// CHO-WithCandidateSCGInfoList-r19 ::=             SEQUENCE (SIZE (1..maxNrofCondCells-r16)) OF CHO-WithCandidateSCGInfo-r19

var cHOWithCandidateSCGInfoListR19SizeConstraints = per.SizeRange(1, common.MaxNrofCondCells_r16)

type CHO_WithCandidateSCGInfoList_r19 []CHO_WithCandidateSCGInfo_r19

func (ie *CHO_WithCandidateSCGInfoList_r19) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(cHOWithCandidateSCGInfoListR19SizeConstraints)
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

func (ie *CHO_WithCandidateSCGInfoList_r19) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(cHOWithCandidateSCGInfoListR19SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(CHO_WithCandidateSCGInfoList_r19, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
