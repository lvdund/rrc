// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SI-RequestResources (line 15024).

var sIRequestResourcesConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ra-PreambleStartIndex"},
		{Name: "ra-AssociationPeriodIndex", Optional: true},
		{Name: "ra-ssb-OccasionMaskIndex", Optional: true},
	},
}

var sIRequestResourcesRaPreambleStartIndexConstraints = per.Constrained(0, 63)

var sIRequestResourcesRaAssociationPeriodIndexConstraints = per.Constrained(0, 15)

var sIRequestResourcesRaSsbOccasionMaskIndexConstraints = per.Constrained(0, 15)

type SI_RequestResources struct {
	Ra_PreambleStartIndex     int64
	Ra_AssociationPeriodIndex *int64
	Ra_Ssb_OccasionMaskIndex  *int64
}

func (ie *SI_RequestResources) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIRequestResourcesConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ra_AssociationPeriodIndex != nil, ie.Ra_Ssb_OccasionMaskIndex != nil}); err != nil {
		return err
	}

	// 2. ra-PreambleStartIndex: integer
	{
		if err := e.EncodeInteger(ie.Ra_PreambleStartIndex, sIRequestResourcesRaPreambleStartIndexConstraints); err != nil {
			return err
		}
	}

	// 3. ra-AssociationPeriodIndex: integer
	{
		if ie.Ra_AssociationPeriodIndex != nil {
			if err := e.EncodeInteger(*ie.Ra_AssociationPeriodIndex, sIRequestResourcesRaAssociationPeriodIndexConstraints); err != nil {
				return err
			}
		}
	}

	// 4. ra-ssb-OccasionMaskIndex: integer
	{
		if ie.Ra_Ssb_OccasionMaskIndex != nil {
			if err := e.EncodeInteger(*ie.Ra_Ssb_OccasionMaskIndex, sIRequestResourcesRaSsbOccasionMaskIndexConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SI_RequestResources) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIRequestResourcesConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ra-PreambleStartIndex: integer
	{
		v0, err := d.DecodeInteger(sIRequestResourcesRaPreambleStartIndexConstraints)
		if err != nil {
			return err
		}
		ie.Ra_PreambleStartIndex = v0
	}

	// 3. ra-AssociationPeriodIndex: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(sIRequestResourcesRaAssociationPeriodIndexConstraints)
			if err != nil {
				return err
			}
			ie.Ra_AssociationPeriodIndex = &v
		}
	}

	// 4. ra-ssb-OccasionMaskIndex: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(sIRequestResourcesRaSsbOccasionMaskIndexConstraints)
			if err != nil {
				return err
			}
			ie.Ra_Ssb_OccasionMaskIndex = &v
		}
	}

	return nil
}
