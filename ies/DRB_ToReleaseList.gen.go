// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: DRB-ToReleaseList (line 13170).
// DRB-ToReleaseList ::=                   SEQUENCE (SIZE (1..maxDRB)) OF DRB-Identity

var dRBToReleaseListSizeConstraints = per.SizeRange(1, common.MaxDRB)

type DRB_ToReleaseList []DRB_Identity

func (ie *DRB_ToReleaseList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(dRBToReleaseListSizeConstraints)
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

func (ie *DRB_ToReleaseList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(dRBToReleaseListSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(DRB_ToReleaseList, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
