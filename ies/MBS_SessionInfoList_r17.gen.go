// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MBS-SessionInfoList-r17 (line 28498).
// MBS-SessionInfoList-r17 ::=      SEQUENCE (SIZE (1..maxNrofMBS-Session-r17)) OF MBS-SessionInfo-r17

var mBSSessionInfoListR17SizeConstraints = per.SizeRange(1, common.MaxNrofMBS_Session_r17)

type MBS_SessionInfoList_r17 []MBS_SessionInfo_r17

func (ie *MBS_SessionInfoList_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(mBSSessionInfoListR17SizeConstraints)
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

func (ie *MBS_SessionInfoList_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(mBSSessionInfoListR17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MBS_SessionInfoList_r17, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
