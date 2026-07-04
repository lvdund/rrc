// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PLMN-IdentityList-EUTRA-EPC (line 5552).
// PLMN-IdentityList-EUTRA-EPC::=          SEQUENCE (SIZE (1..maxPLMN)) OF PLMN-Identity

var pLMNIdentityListEUTRAEPCSizeConstraints = per.SizeRange(1, common.MaxPLMN)

type PLMN_IdentityList_EUTRA_EPC []PLMN_Identity

func (ie *PLMN_IdentityList_EUTRA_EPC) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(pLMNIdentityListEUTRAEPCSizeConstraints)
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

func (ie *PLMN_IdentityList_EUTRA_EPC) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(pLMNIdentityListEUTRAEPCSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(PLMN_IdentityList_EUTRA_EPC, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
