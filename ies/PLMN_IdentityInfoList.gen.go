// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PLMN-IdentityInfoList (line 11899).
// PLMN-IdentityInfoList ::=               SEQUENCE (SIZE (1..maxPLMN)) OF PLMN-IdentityInfo

var pLMNIdentityInfoListSizeConstraints = per.SizeRange(1, common.MaxPLMN)

type PLMN_IdentityInfoList []PLMN_IdentityInfo

func (ie *PLMN_IdentityInfoList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(pLMNIdentityInfoListSizeConstraints)
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

func (ie *PLMN_IdentityInfoList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(pLMNIdentityInfoListSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(PLMN_IdentityInfoList, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
