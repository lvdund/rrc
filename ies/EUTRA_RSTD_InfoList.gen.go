// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: EUTRA-RSTD-InfoList (line 8567).
// EUTRA-RSTD-InfoList ::= SEQUENCE (SIZE (1..maxInterRAT-RSTD-Freq)) OF EUTRA-RSTD-Info

var eUTRARSTDInfoListSizeConstraints = per.SizeRange(1, common.MaxInterRAT_RSTD_Freq)

type EUTRA_RSTD_InfoList []EUTRA_RSTD_Info

func (ie *EUTRA_RSTD_InfoList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(eUTRARSTDInfoListSizeConstraints)
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

func (ie *EUTRA_RSTD_InfoList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(eUTRARSTDInfoListSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(EUTRA_RSTD_InfoList, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
