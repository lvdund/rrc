// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CoverageAreaInfo-r18 (line 4678).

var coverageAreaInfoR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "tn-AreaId-r18"},
		{Name: "tn-ReferenceLocation-r18"},
		{Name: "tn-DistanceRadius-r18"},
	},
}

var coverageAreaInfoR18TnDistanceRadiusR18Constraints = per.Constrained(0, 65535)

type CoverageAreaInfo_r18 struct {
	Tn_AreaId_r18            TN_AreaId_r18
	Tn_ReferenceLocation_r18 ReferenceLocation_r17
	Tn_DistanceRadius_r18    int64
}

func (ie *CoverageAreaInfo_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(coverageAreaInfoR18Constraints)
	_ = seq

	// 1. tn-AreaId-r18: ref
	{
		if err := ie.Tn_AreaId_r18.Encode(e); err != nil {
			return err
		}
	}

	// 2. tn-ReferenceLocation-r18: ref
	{
		if err := ie.Tn_ReferenceLocation_r18.Encode(e); err != nil {
			return err
		}
	}

	// 3. tn-DistanceRadius-r18: integer
	{
		if err := e.EncodeInteger(ie.Tn_DistanceRadius_r18, coverageAreaInfoR18TnDistanceRadiusR18Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *CoverageAreaInfo_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(coverageAreaInfoR18Constraints)
	_ = seq

	// 1. tn-AreaId-r18: ref
	{
		if err := ie.Tn_AreaId_r18.Decode(d); err != nil {
			return err
		}
	}

	// 2. tn-ReferenceLocation-r18: ref
	{
		if err := ie.Tn_ReferenceLocation_r18.Decode(d); err != nil {
			return err
		}
	}

	// 3. tn-DistanceRadius-r18: integer
	{
		v2, err := d.DecodeInteger(coverageAreaInfoR18TnDistanceRadiusR18Constraints)
		if err != nil {
			return err
		}
		ie.Tn_DistanceRadius_r18 = v2
	}

	return nil
}
