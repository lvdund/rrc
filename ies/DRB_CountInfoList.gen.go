// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: DRB-CountInfoList (line 247).
// DRB-CountInfoList ::=           SEQUENCE (SIZE (0..maxDRB)) OF DRB-CountInfo

var dRBCountInfoListSizeConstraints = per.SizeRange(0, common.MaxDRB)

type DRB_CountInfoList []DRB_CountInfo

func (ie *DRB_CountInfoList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(dRBCountInfoListSizeConstraints)
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

func (ie *DRB_CountInfoList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(dRBCountInfoListSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(DRB_CountInfoList, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
