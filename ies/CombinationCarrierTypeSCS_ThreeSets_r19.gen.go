// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CombinationCarrierTypeSCS-ThreeSets-r19 (line 18360).

var combinationCarrierTypeSCSThreeSetsR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "carrierTypeSCS-CellSet1-r19"},
		{Name: "carrierTypeSCS-CellSet2-r19"},
		{Name: "carrierTypeSCS-CellSet3-r19"},
		{Name: "carrierTypeSCS-SchedulingCell-r19"},
	},
}

const (
	CombinationCarrierTypeSCS_ThreeSets_r19_CarrierTypeSCS_CellSet1_r19_Licensed_Fdd_Fr1_15kHz = 0
	CombinationCarrierTypeSCS_ThreeSets_r19_CarrierTypeSCS_CellSet1_r19_Licensed_Tdd_Fr1_30kHz = 1
	CombinationCarrierTypeSCS_ThreeSets_r19_CarrierTypeSCS_CellSet1_r19_Fr2_1_60kHz            = 2
	CombinationCarrierTypeSCS_ThreeSets_r19_CarrierTypeSCS_CellSet1_r19_Fr2_1_120kHz           = 3
)

var combinationCarrierTypeSCSThreeSetsR19CarrierTypeSCSCellSet1R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CombinationCarrierTypeSCS_ThreeSets_r19_CarrierTypeSCS_CellSet1_r19_Licensed_Fdd_Fr1_15kHz, CombinationCarrierTypeSCS_ThreeSets_r19_CarrierTypeSCS_CellSet1_r19_Licensed_Tdd_Fr1_30kHz, CombinationCarrierTypeSCS_ThreeSets_r19_CarrierTypeSCS_CellSet1_r19_Fr2_1_60kHz, CombinationCarrierTypeSCS_ThreeSets_r19_CarrierTypeSCS_CellSet1_r19_Fr2_1_120kHz},
}

const (
	CombinationCarrierTypeSCS_ThreeSets_r19_CarrierTypeSCS_CellSet2_r19_Licensed_Fdd_Fr1_15kHz = 0
	CombinationCarrierTypeSCS_ThreeSets_r19_CarrierTypeSCS_CellSet2_r19_Licensed_Tdd_Fr1_30kHz = 1
	CombinationCarrierTypeSCS_ThreeSets_r19_CarrierTypeSCS_CellSet2_r19_Fr2_1_60kHz            = 2
	CombinationCarrierTypeSCS_ThreeSets_r19_CarrierTypeSCS_CellSet2_r19_Fr2_1_120kHz           = 3
)

var combinationCarrierTypeSCSThreeSetsR19CarrierTypeSCSCellSet2R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CombinationCarrierTypeSCS_ThreeSets_r19_CarrierTypeSCS_CellSet2_r19_Licensed_Fdd_Fr1_15kHz, CombinationCarrierTypeSCS_ThreeSets_r19_CarrierTypeSCS_CellSet2_r19_Licensed_Tdd_Fr1_30kHz, CombinationCarrierTypeSCS_ThreeSets_r19_CarrierTypeSCS_CellSet2_r19_Fr2_1_60kHz, CombinationCarrierTypeSCS_ThreeSets_r19_CarrierTypeSCS_CellSet2_r19_Fr2_1_120kHz},
}

const (
	CombinationCarrierTypeSCS_ThreeSets_r19_CarrierTypeSCS_CellSet3_r19_Licensed_Fdd_Fr1_15kHz = 0
	CombinationCarrierTypeSCS_ThreeSets_r19_CarrierTypeSCS_CellSet3_r19_Licensed_Tdd_Fr1_30kHz = 1
	CombinationCarrierTypeSCS_ThreeSets_r19_CarrierTypeSCS_CellSet3_r19_Fr2_1_60kHz            = 2
	CombinationCarrierTypeSCS_ThreeSets_r19_CarrierTypeSCS_CellSet3_r19_Fr2_1_120kHz           = 3
)

var combinationCarrierTypeSCSThreeSetsR19CarrierTypeSCSCellSet3R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CombinationCarrierTypeSCS_ThreeSets_r19_CarrierTypeSCS_CellSet3_r19_Licensed_Fdd_Fr1_15kHz, CombinationCarrierTypeSCS_ThreeSets_r19_CarrierTypeSCS_CellSet3_r19_Licensed_Tdd_Fr1_30kHz, CombinationCarrierTypeSCS_ThreeSets_r19_CarrierTypeSCS_CellSet3_r19_Fr2_1_60kHz, CombinationCarrierTypeSCS_ThreeSets_r19_CarrierTypeSCS_CellSet3_r19_Fr2_1_120kHz},
}

const (
	CombinationCarrierTypeSCS_ThreeSets_r19_CarrierTypeSCS_SchedulingCell_r19_Licensed_Fdd_Fr1_15kHz = 0
	CombinationCarrierTypeSCS_ThreeSets_r19_CarrierTypeSCS_SchedulingCell_r19_Licensed_Tdd_Fr1_30kHz = 1
)

var combinationCarrierTypeSCSThreeSetsR19CarrierTypeSCSSchedulingCellR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CombinationCarrierTypeSCS_ThreeSets_r19_CarrierTypeSCS_SchedulingCell_r19_Licensed_Fdd_Fr1_15kHz, CombinationCarrierTypeSCS_ThreeSets_r19_CarrierTypeSCS_SchedulingCell_r19_Licensed_Tdd_Fr1_30kHz},
}

type CombinationCarrierTypeSCS_ThreeSets_r19 struct {
	CarrierTypeSCS_CellSet1_r19       int64
	CarrierTypeSCS_CellSet2_r19       int64
	CarrierTypeSCS_CellSet3_r19       int64
	CarrierTypeSCS_SchedulingCell_r19 int64
}

func (ie *CombinationCarrierTypeSCS_ThreeSets_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(combinationCarrierTypeSCSThreeSetsR19Constraints)
	_ = seq

	// 1. carrierTypeSCS-CellSet1-r19: enumerated
	{
		if err := e.EncodeEnumerated(ie.CarrierTypeSCS_CellSet1_r19, combinationCarrierTypeSCSThreeSetsR19CarrierTypeSCSCellSet1R19Constraints); err != nil {
			return err
		}
	}

	// 2. carrierTypeSCS-CellSet2-r19: enumerated
	{
		if err := e.EncodeEnumerated(ie.CarrierTypeSCS_CellSet2_r19, combinationCarrierTypeSCSThreeSetsR19CarrierTypeSCSCellSet2R19Constraints); err != nil {
			return err
		}
	}

	// 3. carrierTypeSCS-CellSet3-r19: enumerated
	{
		if err := e.EncodeEnumerated(ie.CarrierTypeSCS_CellSet3_r19, combinationCarrierTypeSCSThreeSetsR19CarrierTypeSCSCellSet3R19Constraints); err != nil {
			return err
		}
	}

	// 4. carrierTypeSCS-SchedulingCell-r19: enumerated
	{
		if err := e.EncodeEnumerated(ie.CarrierTypeSCS_SchedulingCell_r19, combinationCarrierTypeSCSThreeSetsR19CarrierTypeSCSSchedulingCellR19Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *CombinationCarrierTypeSCS_ThreeSets_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(combinationCarrierTypeSCSThreeSetsR19Constraints)
	_ = seq

	// 1. carrierTypeSCS-CellSet1-r19: enumerated
	{
		v0, err := d.DecodeEnumerated(combinationCarrierTypeSCSThreeSetsR19CarrierTypeSCSCellSet1R19Constraints)
		if err != nil {
			return err
		}
		ie.CarrierTypeSCS_CellSet1_r19 = v0
	}

	// 2. carrierTypeSCS-CellSet2-r19: enumerated
	{
		v1, err := d.DecodeEnumerated(combinationCarrierTypeSCSThreeSetsR19CarrierTypeSCSCellSet2R19Constraints)
		if err != nil {
			return err
		}
		ie.CarrierTypeSCS_CellSet2_r19 = v1
	}

	// 3. carrierTypeSCS-CellSet3-r19: enumerated
	{
		v2, err := d.DecodeEnumerated(combinationCarrierTypeSCSThreeSetsR19CarrierTypeSCSCellSet3R19Constraints)
		if err != nil {
			return err
		}
		ie.CarrierTypeSCS_CellSet3_r19 = v2
	}

	// 4. carrierTypeSCS-SchedulingCell-r19: enumerated
	{
		v3, err := d.DecodeEnumerated(combinationCarrierTypeSCSThreeSetsR19CarrierTypeSCSSchedulingCellR19Constraints)
		if err != nil {
			return err
		}
		ie.CarrierTypeSCS_SchedulingCell_r19 = v3
	}

	return nil
}
