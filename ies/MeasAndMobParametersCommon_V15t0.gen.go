// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasAndMobParametersCommon-v15t0 (line 21301).

var measAndMobParametersCommonV15t0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "intraF-NeighMeasForSCellWithoutSSB", Optional: true},
	},
}

const (
	MeasAndMobParametersCommon_V15t0_IntraF_NeighMeasForSCellWithoutSSB_Supported = 0
)

var measAndMobParametersCommonV15t0IntraFNeighMeasForSCellWithoutSSBConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_V15t0_IntraF_NeighMeasForSCellWithoutSSB_Supported},
}

type MeasAndMobParametersCommon_V15t0 struct {
	IntraF_NeighMeasForSCellWithoutSSB *int64
}

func (ie *MeasAndMobParametersCommon_V15t0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measAndMobParametersCommonV15t0Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.IntraF_NeighMeasForSCellWithoutSSB != nil}); err != nil {
		return err
	}

	// 2. intraF-NeighMeasForSCellWithoutSSB: enumerated
	{
		if ie.IntraF_NeighMeasForSCellWithoutSSB != nil {
			if err := e.EncodeEnumerated(*ie.IntraF_NeighMeasForSCellWithoutSSB, measAndMobParametersCommonV15t0IntraFNeighMeasForSCellWithoutSSBConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MeasAndMobParametersCommon_V15t0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measAndMobParametersCommonV15t0Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. intraF-NeighMeasForSCellWithoutSSB: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(measAndMobParametersCommonV15t0IntraFNeighMeasForSCellWithoutSSBConstraints)
			if err != nil {
				return err
			}
			ie.IntraF_NeighMeasForSCellWithoutSSB = &idx
		}
	}

	return nil
}
