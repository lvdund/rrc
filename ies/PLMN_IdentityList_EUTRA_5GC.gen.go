// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PLMN-IdentityList-EUTRA-5GC (line 5531).
// PLMN-IdentityList-EUTRA-5GC::=          SEQUENCE (SIZE (1..maxPLMN)) OF PLMN-Identity-EUTRA-5GC

var pLMNIdentityListEUTRA5GCSizeConstraints = per.SizeRange(1, common.MaxPLMN)

type PLMN_IdentityList_EUTRA_5GC []PLMN_Identity_EUTRA_5GC

func (ie *PLMN_IdentityList_EUTRA_5GC) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(pLMNIdentityListEUTRA5GCSizeConstraints)
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

func (ie *PLMN_IdentityList_EUTRA_5GC) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(pLMNIdentityListEUTRA5GCSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(PLMN_IdentityList_EUTRA_5GC, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
