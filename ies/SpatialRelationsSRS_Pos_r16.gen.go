// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SpatialRelationsSRS-Pos-r16 (line 25356).

var spatialRelationsSRSPosR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "spatialRelation-SRS-PosBasedOnSSB-Serving-r16", Optional: true},
		{Name: "spatialRelation-SRS-PosBasedOnCSI-RS-Serving-r16", Optional: true},
		{Name: "spatialRelation-SRS-PosBasedOnPRS-Serving-r16", Optional: true},
		{Name: "spatialRelation-SRS-PosBasedOnSRS-r16", Optional: true},
		{Name: "spatialRelation-SRS-PosBasedOnSSB-Neigh-r16", Optional: true},
		{Name: "spatialRelation-SRS-PosBasedOnPRS-Neigh-r16", Optional: true},
	},
}

const (
	SpatialRelationsSRS_Pos_r16_SpatialRelation_SRS_PosBasedOnSSB_Serving_r16_Supported = 0
)

var spatialRelationsSRSPosR16SpatialRelationSRSPosBasedOnSSBServingR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SpatialRelationsSRS_Pos_r16_SpatialRelation_SRS_PosBasedOnSSB_Serving_r16_Supported},
}

const (
	SpatialRelationsSRS_Pos_r16_SpatialRelation_SRS_PosBasedOnCSI_RS_Serving_r16_Supported = 0
)

var spatialRelationsSRSPosR16SpatialRelationSRSPosBasedOnCSIRSServingR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SpatialRelationsSRS_Pos_r16_SpatialRelation_SRS_PosBasedOnCSI_RS_Serving_r16_Supported},
}

const (
	SpatialRelationsSRS_Pos_r16_SpatialRelation_SRS_PosBasedOnPRS_Serving_r16_Supported = 0
)

var spatialRelationsSRSPosR16SpatialRelationSRSPosBasedOnPRSServingR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SpatialRelationsSRS_Pos_r16_SpatialRelation_SRS_PosBasedOnPRS_Serving_r16_Supported},
}

const (
	SpatialRelationsSRS_Pos_r16_SpatialRelation_SRS_PosBasedOnSRS_r16_Supported = 0
)

var spatialRelationsSRSPosR16SpatialRelationSRSPosBasedOnSRSR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SpatialRelationsSRS_Pos_r16_SpatialRelation_SRS_PosBasedOnSRS_r16_Supported},
}

const (
	SpatialRelationsSRS_Pos_r16_SpatialRelation_SRS_PosBasedOnSSB_Neigh_r16_Supported = 0
)

var spatialRelationsSRSPosR16SpatialRelationSRSPosBasedOnSSBNeighR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SpatialRelationsSRS_Pos_r16_SpatialRelation_SRS_PosBasedOnSSB_Neigh_r16_Supported},
}

const (
	SpatialRelationsSRS_Pos_r16_SpatialRelation_SRS_PosBasedOnPRS_Neigh_r16_Supported = 0
)

var spatialRelationsSRSPosR16SpatialRelationSRSPosBasedOnPRSNeighR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SpatialRelationsSRS_Pos_r16_SpatialRelation_SRS_PosBasedOnPRS_Neigh_r16_Supported},
}

type SpatialRelationsSRS_Pos_r16 struct {
	SpatialRelation_SRS_PosBasedOnSSB_Serving_r16    *int64
	SpatialRelation_SRS_PosBasedOnCSI_RS_Serving_r16 *int64
	SpatialRelation_SRS_PosBasedOnPRS_Serving_r16    *int64
	SpatialRelation_SRS_PosBasedOnSRS_r16            *int64
	SpatialRelation_SRS_PosBasedOnSSB_Neigh_r16      *int64
	SpatialRelation_SRS_PosBasedOnPRS_Neigh_r16      *int64
}

func (ie *SpatialRelationsSRS_Pos_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(spatialRelationsSRSPosR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SpatialRelation_SRS_PosBasedOnSSB_Serving_r16 != nil, ie.SpatialRelation_SRS_PosBasedOnCSI_RS_Serving_r16 != nil, ie.SpatialRelation_SRS_PosBasedOnPRS_Serving_r16 != nil, ie.SpatialRelation_SRS_PosBasedOnSRS_r16 != nil, ie.SpatialRelation_SRS_PosBasedOnSSB_Neigh_r16 != nil, ie.SpatialRelation_SRS_PosBasedOnPRS_Neigh_r16 != nil}); err != nil {
		return err
	}

	// 2. spatialRelation-SRS-PosBasedOnSSB-Serving-r16: enumerated
	{
		if ie.SpatialRelation_SRS_PosBasedOnSSB_Serving_r16 != nil {
			if err := e.EncodeEnumerated(*ie.SpatialRelation_SRS_PosBasedOnSSB_Serving_r16, spatialRelationsSRSPosR16SpatialRelationSRSPosBasedOnSSBServingR16Constraints); err != nil {
				return err
			}
		}
	}

	// 3. spatialRelation-SRS-PosBasedOnCSI-RS-Serving-r16: enumerated
	{
		if ie.SpatialRelation_SRS_PosBasedOnCSI_RS_Serving_r16 != nil {
			if err := e.EncodeEnumerated(*ie.SpatialRelation_SRS_PosBasedOnCSI_RS_Serving_r16, spatialRelationsSRSPosR16SpatialRelationSRSPosBasedOnCSIRSServingR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. spatialRelation-SRS-PosBasedOnPRS-Serving-r16: enumerated
	{
		if ie.SpatialRelation_SRS_PosBasedOnPRS_Serving_r16 != nil {
			if err := e.EncodeEnumerated(*ie.SpatialRelation_SRS_PosBasedOnPRS_Serving_r16, spatialRelationsSRSPosR16SpatialRelationSRSPosBasedOnPRSServingR16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. spatialRelation-SRS-PosBasedOnSRS-r16: enumerated
	{
		if ie.SpatialRelation_SRS_PosBasedOnSRS_r16 != nil {
			if err := e.EncodeEnumerated(*ie.SpatialRelation_SRS_PosBasedOnSRS_r16, spatialRelationsSRSPosR16SpatialRelationSRSPosBasedOnSRSR16Constraints); err != nil {
				return err
			}
		}
	}

	// 6. spatialRelation-SRS-PosBasedOnSSB-Neigh-r16: enumerated
	{
		if ie.SpatialRelation_SRS_PosBasedOnSSB_Neigh_r16 != nil {
			if err := e.EncodeEnumerated(*ie.SpatialRelation_SRS_PosBasedOnSSB_Neigh_r16, spatialRelationsSRSPosR16SpatialRelationSRSPosBasedOnSSBNeighR16Constraints); err != nil {
				return err
			}
		}
	}

	// 7. spatialRelation-SRS-PosBasedOnPRS-Neigh-r16: enumerated
	{
		if ie.SpatialRelation_SRS_PosBasedOnPRS_Neigh_r16 != nil {
			if err := e.EncodeEnumerated(*ie.SpatialRelation_SRS_PosBasedOnPRS_Neigh_r16, spatialRelationsSRSPosR16SpatialRelationSRSPosBasedOnPRSNeighR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SpatialRelationsSRS_Pos_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(spatialRelationsSRSPosR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. spatialRelation-SRS-PosBasedOnSSB-Serving-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(spatialRelationsSRSPosR16SpatialRelationSRSPosBasedOnSSBServingR16Constraints)
			if err != nil {
				return err
			}
			ie.SpatialRelation_SRS_PosBasedOnSSB_Serving_r16 = &idx
		}
	}

	// 3. spatialRelation-SRS-PosBasedOnCSI-RS-Serving-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(spatialRelationsSRSPosR16SpatialRelationSRSPosBasedOnCSIRSServingR16Constraints)
			if err != nil {
				return err
			}
			ie.SpatialRelation_SRS_PosBasedOnCSI_RS_Serving_r16 = &idx
		}
	}

	// 4. spatialRelation-SRS-PosBasedOnPRS-Serving-r16: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(spatialRelationsSRSPosR16SpatialRelationSRSPosBasedOnPRSServingR16Constraints)
			if err != nil {
				return err
			}
			ie.SpatialRelation_SRS_PosBasedOnPRS_Serving_r16 = &idx
		}
	}

	// 5. spatialRelation-SRS-PosBasedOnSRS-r16: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(spatialRelationsSRSPosR16SpatialRelationSRSPosBasedOnSRSR16Constraints)
			if err != nil {
				return err
			}
			ie.SpatialRelation_SRS_PosBasedOnSRS_r16 = &idx
		}
	}

	// 6. spatialRelation-SRS-PosBasedOnSSB-Neigh-r16: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(spatialRelationsSRSPosR16SpatialRelationSRSPosBasedOnSSBNeighR16Constraints)
			if err != nil {
				return err
			}
			ie.SpatialRelation_SRS_PosBasedOnSSB_Neigh_r16 = &idx
		}
	}

	// 7. spatialRelation-SRS-PosBasedOnPRS-Neigh-r16: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(spatialRelationsSRSPosR16SpatialRelationSRSPosBasedOnPRSNeighR16Constraints)
			if err != nil {
				return err
			}
			ie.SpatialRelation_SRS_PosBasedOnPRS_Neigh_r16 = &idx
		}
	}

	return nil
}
