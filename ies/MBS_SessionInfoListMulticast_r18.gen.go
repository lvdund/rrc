// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MBS-SessionInfoListMulticast-r18 (line 28555).
// MBS-SessionInfoListMulticast-r18 ::=      SEQUENCE (SIZE (1..maxNrofMBS-Session-r17)) OF MBS-SessionInfoMulticast-r18

var mBSSessionInfoListMulticastR18SizeConstraints = per.SizeRange(1, common.MaxNrofMBS_Session_r17)

type MBS_SessionInfoListMulticast_r18 []MBS_SessionInfoMulticast_r18

func (ie *MBS_SessionInfoListMulticast_r18) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(mBSSessionInfoListMulticastR18SizeConstraints)
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

func (ie *MBS_SessionInfoListMulticast_r18) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(mBSSessionInfoListMulticastR18SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MBS_SessionInfoListMulticast_r18, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
