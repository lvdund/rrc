// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CombinationCarrierTypeSCS-r19 (line 18354).

var combinationCarrierTypeSCSR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "carrierTypeSCS-CellSet1-r19"},
		{Name: "carrierTypeSCS-CellSet2-r19"},
		{Name: "carrierTypeSCS-SchedulingCell-r19"},
	},
}

const (
	CombinationCarrierTypeSCS_r19_CarrierTypeSCS_CellSet1_r19_Licensed_Fdd_Fr1_15kHz = 0
	CombinationCarrierTypeSCS_r19_CarrierTypeSCS_CellSet1_r19_Licensed_Tdd_Fr1_30kHz = 1
	CombinationCarrierTypeSCS_r19_CarrierTypeSCS_CellSet1_r19_Fr2_1_60kHz            = 2
	CombinationCarrierTypeSCS_r19_CarrierTypeSCS_CellSet1_r19_Fr2_1_120kHz           = 3
)

var combinationCarrierTypeSCSR19CarrierTypeSCSCellSet1R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CombinationCarrierTypeSCS_r19_CarrierTypeSCS_CellSet1_r19_Licensed_Fdd_Fr1_15kHz, CombinationCarrierTypeSCS_r19_CarrierTypeSCS_CellSet1_r19_Licensed_Tdd_Fr1_30kHz, CombinationCarrierTypeSCS_r19_CarrierTypeSCS_CellSet1_r19_Fr2_1_60kHz, CombinationCarrierTypeSCS_r19_CarrierTypeSCS_CellSet1_r19_Fr2_1_120kHz},
}

const (
	CombinationCarrierTypeSCS_r19_CarrierTypeSCS_CellSet2_r19_Licensed_Fdd_Fr1_15kHz = 0
	CombinationCarrierTypeSCS_r19_CarrierTypeSCS_CellSet2_r19_Licensed_Tdd_Fr1_30kHz = 1
	CombinationCarrierTypeSCS_r19_CarrierTypeSCS_CellSet2_r19_Fr2_1_60kHz            = 2
	CombinationCarrierTypeSCS_r19_CarrierTypeSCS_CellSet2_r19_Fr2_1_120kHz           = 3
)

var combinationCarrierTypeSCSR19CarrierTypeSCSCellSet2R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CombinationCarrierTypeSCS_r19_CarrierTypeSCS_CellSet2_r19_Licensed_Fdd_Fr1_15kHz, CombinationCarrierTypeSCS_r19_CarrierTypeSCS_CellSet2_r19_Licensed_Tdd_Fr1_30kHz, CombinationCarrierTypeSCS_r19_CarrierTypeSCS_CellSet2_r19_Fr2_1_60kHz, CombinationCarrierTypeSCS_r19_CarrierTypeSCS_CellSet2_r19_Fr2_1_120kHz},
}

const (
	CombinationCarrierTypeSCS_r19_CarrierTypeSCS_SchedulingCell_r19_Licensed_Fdd_Fr1_15kHz = 0
	CombinationCarrierTypeSCS_r19_CarrierTypeSCS_SchedulingCell_r19_Licensed_Tdd_Fr1_30kHz = 1
)

var combinationCarrierTypeSCSR19CarrierTypeSCSSchedulingCellR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CombinationCarrierTypeSCS_r19_CarrierTypeSCS_SchedulingCell_r19_Licensed_Fdd_Fr1_15kHz, CombinationCarrierTypeSCS_r19_CarrierTypeSCS_SchedulingCell_r19_Licensed_Tdd_Fr1_30kHz},
}

type CombinationCarrierTypeSCS_r19 struct {
	CarrierTypeSCS_CellSet1_r19       int64
	CarrierTypeSCS_CellSet2_r19       int64
	CarrierTypeSCS_SchedulingCell_r19 int64
}

func (ie *CombinationCarrierTypeSCS_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(combinationCarrierTypeSCSR19Constraints)
	_ = seq

	// 1. carrierTypeSCS-CellSet1-r19: enumerated
	{
		if err := e.EncodeEnumerated(ie.CarrierTypeSCS_CellSet1_r19, combinationCarrierTypeSCSR19CarrierTypeSCSCellSet1R19Constraints); err != nil {
			return err
		}
	}

	// 2. carrierTypeSCS-CellSet2-r19: enumerated
	{
		if err := e.EncodeEnumerated(ie.CarrierTypeSCS_CellSet2_r19, combinationCarrierTypeSCSR19CarrierTypeSCSCellSet2R19Constraints); err != nil {
			return err
		}
	}

	// 3. carrierTypeSCS-SchedulingCell-r19: enumerated
	{
		if err := e.EncodeEnumerated(ie.CarrierTypeSCS_SchedulingCell_r19, combinationCarrierTypeSCSR19CarrierTypeSCSSchedulingCellR19Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *CombinationCarrierTypeSCS_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(combinationCarrierTypeSCSR19Constraints)
	_ = seq

	// 1. carrierTypeSCS-CellSet1-r19: enumerated
	{
		v0, err := d.DecodeEnumerated(combinationCarrierTypeSCSR19CarrierTypeSCSCellSet1R19Constraints)
		if err != nil {
			return err
		}
		ie.CarrierTypeSCS_CellSet1_r19 = v0
	}

	// 2. carrierTypeSCS-CellSet2-r19: enumerated
	{
		v1, err := d.DecodeEnumerated(combinationCarrierTypeSCSR19CarrierTypeSCSCellSet2R19Constraints)
		if err != nil {
			return err
		}
		ie.CarrierTypeSCS_CellSet2_r19 = v1
	}

	// 3. carrierTypeSCS-SchedulingCell-r19: enumerated
	{
		v2, err := d.DecodeEnumerated(combinationCarrierTypeSCSR19CarrierTypeSCSSchedulingCellR19Constraints)
		if err != nil {
			return err
		}
		ie.CarrierTypeSCS_SchedulingCell_r19 = v2
	}

	return nil
}
