// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: TD-DD-Config-r18 (line 6394).

var tDDDConfigR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "vectorLengthDD-r18"},
		{Name: "unitDurationDD-r18", Optional: true},
		{Name: "aperiodicResourceOffset-r18", Optional: true},
		{Name: "tdCQI-r18", Optional: true},
	},
}

const (
	TD_DD_Config_r18_VectorLengthDD_r18_N1 = 0
	TD_DD_Config_r18_VectorLengthDD_r18_N2 = 1
	TD_DD_Config_r18_VectorLengthDD_r18_N4 = 2
	TD_DD_Config_r18_VectorLengthDD_r18_N8 = 3
)

var tDDDConfigR18VectorLengthDDR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{TD_DD_Config_r18_VectorLengthDD_r18_N1, TD_DD_Config_r18_VectorLengthDD_r18_N2, TD_DD_Config_r18_VectorLengthDD_r18_N4, TD_DD_Config_r18_VectorLengthDD_r18_N8},
}

const (
	TD_DD_Config_r18_UnitDurationDD_r18_M1 = 0
	TD_DD_Config_r18_UnitDurationDD_r18_M2 = 1
)

var tDDDConfigR18UnitDurationDDR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{TD_DD_Config_r18_UnitDurationDD_r18_M1, TD_DD_Config_r18_UnitDurationDD_r18_M2},
}

var tDDDConfigR18AperiodicResourceOffsetR18Constraints = per.Constrained(1, 2)

const (
	TD_DD_Config_r18_TdCQI_r18_N11    = 0
	TD_DD_Config_r18_TdCQI_r18_N12    = 1
	TD_DD_Config_r18_TdCQI_r18_N2     = 2
	TD_DD_Config_r18_TdCQI_r18_Spare1 = 3
)

var tDDDConfigR18TdCQIR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{TD_DD_Config_r18_TdCQI_r18_N11, TD_DD_Config_r18_TdCQI_r18_N12, TD_DD_Config_r18_TdCQI_r18_N2, TD_DD_Config_r18_TdCQI_r18_Spare1},
}

type TD_DD_Config_r18 struct {
	VectorLengthDD_r18          int64
	UnitDurationDD_r18          *int64
	AperiodicResourceOffset_r18 *int64
	TdCQI_r18                   *int64
}

func (ie *TD_DD_Config_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(tDDDConfigR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.UnitDurationDD_r18 != nil, ie.AperiodicResourceOffset_r18 != nil, ie.TdCQI_r18 != nil}); err != nil {
		return err
	}

	// 2. vectorLengthDD-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.VectorLengthDD_r18, tDDDConfigR18VectorLengthDDR18Constraints); err != nil {
			return err
		}
	}

	// 3. unitDurationDD-r18: enumerated
	{
		if ie.UnitDurationDD_r18 != nil {
			if err := e.EncodeEnumerated(*ie.UnitDurationDD_r18, tDDDConfigR18UnitDurationDDR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. aperiodicResourceOffset-r18: integer
	{
		if ie.AperiodicResourceOffset_r18 != nil {
			if err := e.EncodeInteger(*ie.AperiodicResourceOffset_r18, tDDDConfigR18AperiodicResourceOffsetR18Constraints); err != nil {
				return err
			}
		}
	}

	// 5. tdCQI-r18: enumerated
	{
		if ie.TdCQI_r18 != nil {
			if err := e.EncodeEnumerated(*ie.TdCQI_r18, tDDDConfigR18TdCQIR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *TD_DD_Config_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(tDDDConfigR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. vectorLengthDD-r18: enumerated
	{
		v0, err := d.DecodeEnumerated(tDDDConfigR18VectorLengthDDR18Constraints)
		if err != nil {
			return err
		}
		ie.VectorLengthDD_r18 = v0
	}

	// 3. unitDurationDD-r18: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(tDDDConfigR18UnitDurationDDR18Constraints)
			if err != nil {
				return err
			}
			ie.UnitDurationDD_r18 = &idx
		}
	}

	// 4. aperiodicResourceOffset-r18: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(tDDDConfigR18AperiodicResourceOffsetR18Constraints)
			if err != nil {
				return err
			}
			ie.AperiodicResourceOffset_r18 = &v
		}
	}

	// 5. tdCQI-r18: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(tDDDConfigR18TdCQIR18Constraints)
			if err != nil {
				return err
			}
			ie.TdCQI_r18 = &idx
		}
	}

	return nil
}
