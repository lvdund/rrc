// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MBS-SessionInfoList-v1900 (line 28500).
// MBS-SessionInfoList-v1900 ::=    SEQUENCE (SIZE (1..maxNrofMBS-Session-r17)) OF MBS-SessionInfo-v1900

var mBSSessionInfoListV1900SizeConstraints = per.SizeRange(1, common.MaxNrofMBS_Session_r17)

type MBS_SessionInfoList_v1900 []MBS_SessionInfo_v1900

func (ie *MBS_SessionInfoList_v1900) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(mBSSessionInfoListV1900SizeConstraints)
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

func (ie *MBS_SessionInfoList_v1900) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(mBSSessionInfoListV1900SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MBS_SessionInfoList_v1900, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
