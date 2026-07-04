// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: EUTRA-MultiBandInfoListAerial-r18 (line 26168).
// EUTRA-MultiBandInfoListAerial-r18 ::=     SEQUENCE (SIZE (1..maxMultiBands)) OF EUTRA-MultiBandInfoAerial-r18

var eUTRAMultiBandInfoListAerialR18SizeConstraints = per.SizeRange(1, common.MaxMultiBands)

type EUTRA_MultiBandInfoListAerial_r18 []EUTRA_MultiBandInfoAerial_r18

func (ie *EUTRA_MultiBandInfoListAerial_r18) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(eUTRAMultiBandInfoListAerialR18SizeConstraints)
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

func (ie *EUTRA_MultiBandInfoListAerial_r18) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(eUTRAMultiBandInfoListAerialR18SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(EUTRA_MultiBandInfoListAerial_r18, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
