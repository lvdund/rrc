// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MBS-ServiceList-r17 (line 28489).
// MBS-ServiceList-r17 ::=         SEQUENCE (SIZE (1..maxNrofMBS-ServiceListPerUE-r17)) OF MBS-ServiceInfo-r17

var mBSServiceListR17SizeConstraints = per.SizeRange(1, common.MaxNrofMBS_ServiceListPerUE_r17)

type MBS_ServiceList_r17 []MBS_ServiceInfo_r17

func (ie *MBS_ServiceList_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(mBSServiceListR17SizeConstraints)
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

func (ie *MBS_ServiceList_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(mBSServiceListR17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MBS_ServiceList_r17, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
