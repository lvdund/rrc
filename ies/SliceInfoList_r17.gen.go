// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SliceInfoList-r17 (line 8381).
// SliceInfoList-r17 ::=             SEQUENCE (SIZE (1..maxSliceInfo-r17)) OF SliceInfo-r17

var sliceInfoListR17SizeConstraints = per.SizeRange(1, common.MaxSliceInfo_r17)

type SliceInfoList_r17 []SliceInfo_r17

func (ie *SliceInfoList_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sliceInfoListR17SizeConstraints)
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

func (ie *SliceInfoList_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sliceInfoListR17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SliceInfoList_r17, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
