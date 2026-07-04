// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MBS-NonServingInfoList-r18 (line 28464).
// MBS-NonServingInfoList-r18 ::=    SEQUENCE (SIZE (1..maxFreqMBS-r17)) OF NonServingInfo-r18

var mBSNonServingInfoListR18SizeConstraints = per.SizeRange(1, common.MaxFreqMBS_r17)

type MBS_NonServingInfoList_r18 []NonServingInfo_r18

func (ie *MBS_NonServingInfoList_r18) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(mBSNonServingInfoListR18SizeConstraints)
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

func (ie *MBS_NonServingInfoList_r18) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(mBSNonServingInfoListR18SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MBS_NonServingInfoList_r18, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
