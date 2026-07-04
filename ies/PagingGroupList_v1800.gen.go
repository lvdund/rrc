// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PagingGroupList-v1800 (line 862).
// PagingGroupList-v1800 ::=           SEQUENCE (SIZE(1..maxNrofPageGroup-r17)) OF GroupPaging-r18

var pagingGroupListV1800SizeConstraints = per.SizeRange(1, common.MaxNrofPageGroup_r17)

type PagingGroupList_v1800 []GroupPaging_r18

func (ie *PagingGroupList_v1800) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(pagingGroupListV1800SizeConstraints)
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

func (ie *PagingGroupList_v1800) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(pagingGroupListV1800SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(PagingGroupList_v1800, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
