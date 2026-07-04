// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CSSF-MeasMO-List-r19 (line 9131).
// CSSF-MeasMO-List-r19 ::=                SEQUENCE (SIZE (1..maxNrofMeasMO-r19)) OF MeasObjectId

var cSSFMeasMOListR19SizeConstraints = per.SizeRange(1, common.MaxNrofMeasMO_r19)

type CSSF_MeasMO_List_r19 []MeasObjectId

func (ie *CSSF_MeasMO_List_r19) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(cSSFMeasMOListR19SizeConstraints)
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

func (ie *CSSF_MeasMO_List_r19) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(cSSFMeasMOListR19SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(CSSF_MeasMO_List_r19, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
