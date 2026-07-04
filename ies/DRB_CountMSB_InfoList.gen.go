// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: DRB-CountMSB-InfoList (line 221).
// DRB-CountMSB-InfoList ::=       SEQUENCE (SIZE (1..maxDRB)) OF DRB-CountMSB-Info

var dRBCountMSBInfoListSizeConstraints = per.SizeRange(1, common.MaxDRB)

type DRB_CountMSB_InfoList []DRB_CountMSB_Info

func (ie *DRB_CountMSB_InfoList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(dRBCountMSBInfoListSizeConstraints)
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

func (ie *DRB_CountMSB_InfoList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(dRBCountMSBInfoListSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(DRB_CountMSB_InfoList, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
