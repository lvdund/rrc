// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SliceInfoListDedicated-r17 (line 8363).
// SliceInfoListDedicated-r17 ::=           SEQUENCE (SIZE (1..maxSliceInfo-r17)) OF SliceInfoDedicated-r17

var sliceInfoListDedicatedR17SizeConstraints = per.SizeRange(1, common.MaxSliceInfo_r17)

type SliceInfoListDedicated_r17 []SliceInfoDedicated_r17

func (ie *SliceInfoListDedicated_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sliceInfoListDedicatedR17SizeConstraints)
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

func (ie *SliceInfoListDedicated_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sliceInfoListDedicatedR17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SliceInfoListDedicated_r17, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
