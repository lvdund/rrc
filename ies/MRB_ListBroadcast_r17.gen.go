// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MRB-ListBroadcast-r17 (line 28522).
// MRB-ListBroadcast-r17 ::=            SEQUENCE (SIZE (1..maxNrofMRB-Broadcast-r17)) OF MRB-InfoBroadcast-r17

var mRBListBroadcastR17SizeConstraints = per.SizeRange(1, common.MaxNrofMRB_Broadcast_r17)

type MRB_ListBroadcast_r17 []MRB_InfoBroadcast_r17

func (ie *MRB_ListBroadcast_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(mRBListBroadcastR17SizeConstraints)
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

func (ie *MRB_ListBroadcast_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(mRBListBroadcastR17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MRB_ListBroadcast_r17, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
