// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PagingRecordList-v1800 (line 860).
// PagingRecordList-v1800 ::=          SEQUENCE (SIZE(1..maxNrofPageRec)) OF PagingRecord-v1800

var pagingRecordListV1800SizeConstraints = per.SizeRange(1, common.MaxNrofPageRec)

type PagingRecordList_v1800 []PagingRecord_v1800

func (ie *PagingRecordList_v1800) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(pagingRecordListV1800SizeConstraints)
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

func (ie *PagingRecordList_v1800) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(pagingRecordListV1800SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(PagingRecordList_v1800, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
