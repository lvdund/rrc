// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SIB1-RequestResources-r19 (line 4742).

var sIB1RequestResourcesR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sib1-RA-PreambleStartIndex-r19"},
		{Name: "sib1-RA-AssociationPeriodIndex-r19", Optional: true},
		{Name: "sib1-RA-SSB-OccasionMaskIndex-r19", Optional: true},
	},
}

var sIB1RequestResourcesR19Sib1RAPreambleStartIndexR19Constraints = per.Constrained(0, 63)

var sIB1RequestResourcesR19Sib1RAAssociationPeriodIndexR19Constraints = per.Constrained(0, 15)

var sIB1RequestResourcesR19Sib1RASSBOccasionMaskIndexR19Constraints = per.Constrained(0, 15)

type SIB1_RequestResources_r19 struct {
	Sib1_RA_PreambleStartIndex_r19     int64
	Sib1_RA_AssociationPeriodIndex_r19 *int64
	Sib1_RA_SSB_OccasionMaskIndex_r19  *int64
}

func (ie *SIB1_RequestResources_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIB1RequestResourcesR19Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sib1_RA_AssociationPeriodIndex_r19 != nil, ie.Sib1_RA_SSB_OccasionMaskIndex_r19 != nil}); err != nil {
		return err
	}

	// 2. sib1-RA-PreambleStartIndex-r19: integer
	{
		if err := e.EncodeInteger(ie.Sib1_RA_PreambleStartIndex_r19, sIB1RequestResourcesR19Sib1RAPreambleStartIndexR19Constraints); err != nil {
			return err
		}
	}

	// 3. sib1-RA-AssociationPeriodIndex-r19: integer
	{
		if ie.Sib1_RA_AssociationPeriodIndex_r19 != nil {
			if err := e.EncodeInteger(*ie.Sib1_RA_AssociationPeriodIndex_r19, sIB1RequestResourcesR19Sib1RAAssociationPeriodIndexR19Constraints); err != nil {
				return err
			}
		}
	}

	// 4. sib1-RA-SSB-OccasionMaskIndex-r19: integer
	{
		if ie.Sib1_RA_SSB_OccasionMaskIndex_r19 != nil {
			if err := e.EncodeInteger(*ie.Sib1_RA_SSB_OccasionMaskIndex_r19, sIB1RequestResourcesR19Sib1RASSBOccasionMaskIndexR19Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SIB1_RequestResources_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIB1RequestResourcesR19Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sib1-RA-PreambleStartIndex-r19: integer
	{
		v0, err := d.DecodeInteger(sIB1RequestResourcesR19Sib1RAPreambleStartIndexR19Constraints)
		if err != nil {
			return err
		}
		ie.Sib1_RA_PreambleStartIndex_r19 = v0
	}

	// 3. sib1-RA-AssociationPeriodIndex-r19: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(sIB1RequestResourcesR19Sib1RAAssociationPeriodIndexR19Constraints)
			if err != nil {
				return err
			}
			ie.Sib1_RA_AssociationPeriodIndex_r19 = &v
		}
	}

	// 4. sib1-RA-SSB-OccasionMaskIndex-r19: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(sIB1RequestResourcesR19Sib1RASSBOccasionMaskIndexR19Constraints)
			if err != nil {
				return err
			}
			ie.Sib1_RA_SSB_OccasionMaskIndex_r19 = &v
		}
	}

	return nil
}
