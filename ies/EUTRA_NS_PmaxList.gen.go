// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: EUTRA-NS-PmaxList (line 26178).
// EUTRA-NS-PmaxList ::=               SEQUENCE (SIZE (1..maxEUTRA-NS-Pmax)) OF EUTRA-NS-PmaxValue

var eUTRANSPmaxListSizeConstraints = per.SizeRange(1, common.MaxEUTRA_NS_Pmax)

type EUTRA_NS_PmaxList []EUTRA_NS_PmaxValue

func (ie *EUTRA_NS_PmaxList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(eUTRANSPmaxListSizeConstraints)
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

func (ie *EUTRA_NS_PmaxList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(eUTRANSPmaxListSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(EUTRA_NS_PmaxList, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
