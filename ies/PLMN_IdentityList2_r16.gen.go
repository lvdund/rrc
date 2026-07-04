// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PLMN-IdentityList2-r16 (line 11922).
// PLMN-IdentityList2-r16 ::= SEQUENCE (SIZE (1..16)) OF PLMN-Identity

var pLMNIdentityList2R16SizeConstraints = per.SizeRange(1, 16)

type PLMN_IdentityList2_r16 []PLMN_Identity

func (ie *PLMN_IdentityList2_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(pLMNIdentityList2R16SizeConstraints)
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

func (ie *PLMN_IdentityList2_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(pLMNIdentityList2R16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(PLMN_IdentityList2_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
