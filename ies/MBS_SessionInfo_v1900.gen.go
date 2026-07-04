// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MBS-SessionInfo-v1900 (line 28512).

var mBSSessionInfoV1900Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "mbs-SessionAreaList-r19", Optional: true},
	},
}

var mBSSessionInfoV1900MbsSessionAreaListR19Constraints = per.SizeRange(1, common.MaxNrofMBS_AreaPerSession_r19)

type MBS_SessionInfo_v1900 struct {
	Mbs_SessionAreaList_r19 []MBS_IntendedAreaID_r19
}

func (ie *MBS_SessionInfo_v1900) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mBSSessionInfoV1900Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Mbs_SessionAreaList_r19 != nil}); err != nil {
		return err
	}

	// 2. mbs-SessionAreaList-r19: sequence-of
	{
		if ie.Mbs_SessionAreaList_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(mBSSessionInfoV1900MbsSessionAreaListR19Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Mbs_SessionAreaList_r19))); err != nil {
				return err
			}
			for i := range ie.Mbs_SessionAreaList_r19 {
				if err := ie.Mbs_SessionAreaList_r19[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *MBS_SessionInfo_v1900) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mBSSessionInfoV1900Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. mbs-SessionAreaList-r19: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(mBSSessionInfoV1900MbsSessionAreaListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Mbs_SessionAreaList_r19 = make([]MBS_IntendedAreaID_r19, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Mbs_SessionAreaList_r19[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
