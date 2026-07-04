// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PagingGroupList-r17 (line 858).
// PagingGroupList-r17 ::=             SEQUENCE (SIZE(1..maxNrofPageGroup-r17)) OF TMGI-r17

var pagingGroupListR17SizeConstraints = per.SizeRange(1, common.MaxNrofPageGroup_r17)

type PagingGroupList_r17 []TMGI_r17

func (ie *PagingGroupList_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(pagingGroupListR17SizeConstraints)
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

func (ie *PagingGroupList_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(pagingGroupListR17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(PagingGroupList_r17, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
