// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: IntendedServiceAreaInfo-r19 (line 4768).

var intendedServiceAreaInfoR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "intendedServiceAreaId-r19"},
		{Name: "areaCoordinates-r19"},
	},
}

type IntendedServiceAreaInfo_r19 struct {
	IntendedServiceAreaId_r19 MBS_IntendedAreaID_r19
	AreaCoordinates_r19       IntendedServiceAreaInfo_AreaCoordinates_r19
}

func (ie *IntendedServiceAreaInfo_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(intendedServiceAreaInfoR19Constraints)
	_ = seq

	// 1. intendedServiceAreaId-r19: ref
	{
		if err := ie.IntendedServiceAreaId_r19.Encode(e); err != nil {
			return err
		}
	}

	// 2. areaCoordinates-r19: ref
	{
		if err := ie.AreaCoordinates_r19.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *IntendedServiceAreaInfo_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(intendedServiceAreaInfoR19Constraints)
	_ = seq

	// 1. intendedServiceAreaId-r19: ref
	{
		if err := ie.IntendedServiceAreaId_r19.Decode(d); err != nil {
			return err
		}
	}

	// 2. areaCoordinates-r19: ref
	{
		if err := ie.AreaCoordinates_r19.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
