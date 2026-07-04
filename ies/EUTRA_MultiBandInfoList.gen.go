// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: EUTRA-MultiBandInfoList (line 26158).
// EUTRA-MultiBandInfoList ::=     SEQUENCE (SIZE (1..maxMultiBands)) OF EUTRA-MultiBandInfo

var eUTRAMultiBandInfoListSizeConstraints = per.SizeRange(1, common.MaxMultiBands)

type EUTRA_MultiBandInfoList []EUTRA_MultiBandInfo

func (ie *EUTRA_MultiBandInfoList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(eUTRAMultiBandInfoListSizeConstraints)
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

func (ie *EUTRA_MultiBandInfoList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(eUTRAMultiBandInfoListSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(EUTRA_MultiBandInfoList, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
