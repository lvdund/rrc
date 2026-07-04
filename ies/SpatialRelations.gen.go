// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SpatialRelations (line 22519).

var spatialRelationsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxNumberConfiguredSpatialRelations"},
		{Name: "maxNumberActiveSpatialRelations"},
		{Name: "additionalActiveSpatialRelationPUCCH", Optional: true},
		{Name: "maxNumberDL-RS-QCL-TypeD"},
	},
}

const (
	SpatialRelations_MaxNumberConfiguredSpatialRelations_N4  = 0
	SpatialRelations_MaxNumberConfiguredSpatialRelations_N8  = 1
	SpatialRelations_MaxNumberConfiguredSpatialRelations_N16 = 2
	SpatialRelations_MaxNumberConfiguredSpatialRelations_N32 = 3
	SpatialRelations_MaxNumberConfiguredSpatialRelations_N64 = 4
	SpatialRelations_MaxNumberConfiguredSpatialRelations_N96 = 5
)

var spatialRelationsMaxNumberConfiguredSpatialRelationsConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SpatialRelations_MaxNumberConfiguredSpatialRelations_N4, SpatialRelations_MaxNumberConfiguredSpatialRelations_N8, SpatialRelations_MaxNumberConfiguredSpatialRelations_N16, SpatialRelations_MaxNumberConfiguredSpatialRelations_N32, SpatialRelations_MaxNumberConfiguredSpatialRelations_N64, SpatialRelations_MaxNumberConfiguredSpatialRelations_N96},
}

const (
	SpatialRelations_MaxNumberActiveSpatialRelations_N1  = 0
	SpatialRelations_MaxNumberActiveSpatialRelations_N2  = 1
	SpatialRelations_MaxNumberActiveSpatialRelations_N4  = 2
	SpatialRelations_MaxNumberActiveSpatialRelations_N8  = 3
	SpatialRelations_MaxNumberActiveSpatialRelations_N14 = 4
)

var spatialRelationsMaxNumberActiveSpatialRelationsConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SpatialRelations_MaxNumberActiveSpatialRelations_N1, SpatialRelations_MaxNumberActiveSpatialRelations_N2, SpatialRelations_MaxNumberActiveSpatialRelations_N4, SpatialRelations_MaxNumberActiveSpatialRelations_N8, SpatialRelations_MaxNumberActiveSpatialRelations_N14},
}

const (
	SpatialRelations_AdditionalActiveSpatialRelationPUCCH_Supported = 0
)

var spatialRelationsAdditionalActiveSpatialRelationPUCCHConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SpatialRelations_AdditionalActiveSpatialRelationPUCCH_Supported},
}

const (
	SpatialRelations_MaxNumberDL_RS_QCL_TypeD_N1  = 0
	SpatialRelations_MaxNumberDL_RS_QCL_TypeD_N2  = 1
	SpatialRelations_MaxNumberDL_RS_QCL_TypeD_N4  = 2
	SpatialRelations_MaxNumberDL_RS_QCL_TypeD_N8  = 3
	SpatialRelations_MaxNumberDL_RS_QCL_TypeD_N14 = 4
)

var spatialRelationsMaxNumberDLRSQCLTypeDConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SpatialRelations_MaxNumberDL_RS_QCL_TypeD_N1, SpatialRelations_MaxNumberDL_RS_QCL_TypeD_N2, SpatialRelations_MaxNumberDL_RS_QCL_TypeD_N4, SpatialRelations_MaxNumberDL_RS_QCL_TypeD_N8, SpatialRelations_MaxNumberDL_RS_QCL_TypeD_N14},
}

type SpatialRelations struct {
	MaxNumberConfiguredSpatialRelations  int64
	MaxNumberActiveSpatialRelations      int64
	AdditionalActiveSpatialRelationPUCCH *int64
	MaxNumberDL_RS_QCL_TypeD             int64
}

func (ie *SpatialRelations) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(spatialRelationsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.AdditionalActiveSpatialRelationPUCCH != nil}); err != nil {
		return err
	}

	// 2. maxNumberConfiguredSpatialRelations: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxNumberConfiguredSpatialRelations, spatialRelationsMaxNumberConfiguredSpatialRelationsConstraints); err != nil {
			return err
		}
	}

	// 3. maxNumberActiveSpatialRelations: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxNumberActiveSpatialRelations, spatialRelationsMaxNumberActiveSpatialRelationsConstraints); err != nil {
			return err
		}
	}

	// 4. additionalActiveSpatialRelationPUCCH: enumerated
	{
		if ie.AdditionalActiveSpatialRelationPUCCH != nil {
			if err := e.EncodeEnumerated(*ie.AdditionalActiveSpatialRelationPUCCH, spatialRelationsAdditionalActiveSpatialRelationPUCCHConstraints); err != nil {
				return err
			}
		}
	}

	// 5. maxNumberDL-RS-QCL-TypeD: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxNumberDL_RS_QCL_TypeD, spatialRelationsMaxNumberDLRSQCLTypeDConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SpatialRelations) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(spatialRelationsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. maxNumberConfiguredSpatialRelations: enumerated
	{
		v0, err := d.DecodeEnumerated(spatialRelationsMaxNumberConfiguredSpatialRelationsConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberConfiguredSpatialRelations = v0
	}

	// 3. maxNumberActiveSpatialRelations: enumerated
	{
		v1, err := d.DecodeEnumerated(spatialRelationsMaxNumberActiveSpatialRelationsConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberActiveSpatialRelations = v1
	}

	// 4. additionalActiveSpatialRelationPUCCH: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(spatialRelationsAdditionalActiveSpatialRelationPUCCHConstraints)
			if err != nil {
				return err
			}
			ie.AdditionalActiveSpatialRelationPUCCH = &idx
		}
	}

	// 5. maxNumberDL-RS-QCL-TypeD: enumerated
	{
		v3, err := d.DecodeEnumerated(spatialRelationsMaxNumberDLRSQCLTypeDConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberDL_RS_QCL_TypeD = v3
	}

	return nil
}
