// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: IntendedServiceAreaList-r19 (line 4757).
// IntendedServiceAreaList-r19 ::=   SEQUENCE (SIZE (1..maxNrofMBS-Area-r19)) OF IntendedServiceAreaInfo-r19

var intendedServiceAreaListR19SizeConstraints = per.SizeRange(1, common.MaxNrofMBS_Area_r19)

type IntendedServiceAreaList_r19 []IntendedServiceAreaInfo_r19

func (ie *IntendedServiceAreaList_r19) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(intendedServiceAreaListR19SizeConstraints)
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

func (ie *IntendedServiceAreaList_r19) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(intendedServiceAreaListR19SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(IntendedServiceAreaList_r19, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
