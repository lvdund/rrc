// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: NPN-IdentityInfoList-r16 (line 10581).
// NPN-IdentityInfoList-r16 ::=     SEQUENCE (SIZE (1..maxNPN-r16)) OF NPN-IdentityInfo-r16

var nPNIdentityInfoListR16SizeConstraints = per.SizeRange(1, common.MaxNPN_r16)

type NPN_IdentityInfoList_r16 []NPN_IdentityInfo_r16

func (ie *NPN_IdentityInfoList_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(nPNIdentityInfoListR16SizeConstraints)
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

func (ie *NPN_IdentityInfoList_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(nPNIdentityInfoListR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(NPN_IdentityInfoList_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
