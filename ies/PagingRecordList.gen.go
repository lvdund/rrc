// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PagingRecordList (line 854).
// PagingRecordList ::=                SEQUENCE (SIZE(1..maxNrofPageRec)) OF PagingRecord

var pagingRecordListSizeConstraints = per.SizeRange(1, common.MaxNrofPageRec)

type PagingRecordList []PagingRecord

func (ie *PagingRecordList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(pagingRecordListSizeConstraints)
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

func (ie *PagingRecordList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(pagingRecordListSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(PagingRecordList, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
