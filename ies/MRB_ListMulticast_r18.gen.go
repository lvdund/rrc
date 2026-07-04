// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MRB-ListMulticast-r18 (line 28576).
// MRB-ListMulticast-r18 ::=          SEQUENCE (SIZE (1.. maxMRB-r17)) OF MRB-InfoMulticast-r18

var mRBListMulticastR18SizeConstraints = per.SizeRange(1, common.MaxMRB_r17)

type MRB_ListMulticast_r18 []MRB_InfoMulticast_r18

func (ie *MRB_ListMulticast_r18) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(mRBListMulticastR18SizeConstraints)
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

func (ie *MRB_ListMulticast_r18) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(mRBListMulticastR18SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MRB_ListMulticast_r18, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
