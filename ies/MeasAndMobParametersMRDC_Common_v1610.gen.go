// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasAndMobParametersMRDC-Common-v1610 (line 21406).

var measAndMobParametersMRDCCommonV1610Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "condPSCellChangeParametersCommon-r16", Optional: true},
		{Name: "pscellT312-r16", Optional: true},
	},
}

const (
	MeasAndMobParametersMRDC_Common_v1610_PscellT312_r16_Supported = 0
)

var measAndMobParametersMRDCCommonV1610PscellT312R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersMRDC_Common_v1610_PscellT312_r16_Supported},
}

var measAndMobParametersMRDCCommonV1610CondPSCellChangeParametersCommonR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "condPSCellChangeFDD-TDD-r16", Optional: true},
		{Name: "condPSCellChangeFR1-FR2-r16", Optional: true},
	},
}

const (
	MeasAndMobParametersMRDC_Common_v1610_CondPSCellChangeParametersCommon_r16_CondPSCellChangeFDD_TDD_r16_Supported = 0
)

var measAndMobParametersMRDCCommonV1610CondPSCellChangeParametersCommonR16CondPSCellChangeFDDTDDR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersMRDC_Common_v1610_CondPSCellChangeParametersCommon_r16_CondPSCellChangeFDD_TDD_r16_Supported},
}

const (
	MeasAndMobParametersMRDC_Common_v1610_CondPSCellChangeParametersCommon_r16_CondPSCellChangeFR1_FR2_r16_Supported = 0
)

var measAndMobParametersMRDCCommonV1610CondPSCellChangeParametersCommonR16CondPSCellChangeFR1FR2R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersMRDC_Common_v1610_CondPSCellChangeParametersCommon_r16_CondPSCellChangeFR1_FR2_r16_Supported},
}

type MeasAndMobParametersMRDC_Common_v1610 struct {
	CondPSCellChangeParametersCommon_r16 *struct {
		CondPSCellChangeFDD_TDD_r16 *int64
		CondPSCellChangeFR1_FR2_r16 *int64
	}
	PscellT312_r16 *int64
}

func (ie *MeasAndMobParametersMRDC_Common_v1610) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measAndMobParametersMRDCCommonV1610Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.CondPSCellChangeParametersCommon_r16 != nil, ie.PscellT312_r16 != nil}); err != nil {
		return err
	}

	// 2. condPSCellChangeParametersCommon-r16: sequence
	{
		if ie.CondPSCellChangeParametersCommon_r16 != nil {
			c := ie.CondPSCellChangeParametersCommon_r16
			measAndMobParametersMRDCCommonV1610CondPSCellChangeParametersCommonR16Seq := e.NewSequenceEncoder(measAndMobParametersMRDCCommonV1610CondPSCellChangeParametersCommonR16Constraints)
			if err := measAndMobParametersMRDCCommonV1610CondPSCellChangeParametersCommonR16Seq.EncodePreamble([]bool{c.CondPSCellChangeFDD_TDD_r16 != nil, c.CondPSCellChangeFR1_FR2_r16 != nil}); err != nil {
				return err
			}
			if c.CondPSCellChangeFDD_TDD_r16 != nil {
				if err := e.EncodeEnumerated((*c.CondPSCellChangeFDD_TDD_r16), measAndMobParametersMRDCCommonV1610CondPSCellChangeParametersCommonR16CondPSCellChangeFDDTDDR16Constraints); err != nil {
					return err
				}
			}
			if c.CondPSCellChangeFR1_FR2_r16 != nil {
				if err := e.EncodeEnumerated((*c.CondPSCellChangeFR1_FR2_r16), measAndMobParametersMRDCCommonV1610CondPSCellChangeParametersCommonR16CondPSCellChangeFR1FR2R16Constraints); err != nil {
					return err
				}
			}
		}
	}

	// 3. pscellT312-r16: enumerated
	{
		if ie.PscellT312_r16 != nil {
			if err := e.EncodeEnumerated(*ie.PscellT312_r16, measAndMobParametersMRDCCommonV1610PscellT312R16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MeasAndMobParametersMRDC_Common_v1610) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measAndMobParametersMRDCCommonV1610Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. condPSCellChangeParametersCommon-r16: sequence
	{
		if seq.IsComponentPresent(0) {
			ie.CondPSCellChangeParametersCommon_r16 = &struct {
				CondPSCellChangeFDD_TDD_r16 *int64
				CondPSCellChangeFR1_FR2_r16 *int64
			}{}
			c := ie.CondPSCellChangeParametersCommon_r16
			measAndMobParametersMRDCCommonV1610CondPSCellChangeParametersCommonR16Seq := d.NewSequenceDecoder(measAndMobParametersMRDCCommonV1610CondPSCellChangeParametersCommonR16Constraints)
			if err := measAndMobParametersMRDCCommonV1610CondPSCellChangeParametersCommonR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if measAndMobParametersMRDCCommonV1610CondPSCellChangeParametersCommonR16Seq.IsComponentPresent(0) {
				c.CondPSCellChangeFDD_TDD_r16 = new(int64)
				v, err := d.DecodeEnumerated(measAndMobParametersMRDCCommonV1610CondPSCellChangeParametersCommonR16CondPSCellChangeFDDTDDR16Constraints)
				if err != nil {
					return err
				}
				(*c.CondPSCellChangeFDD_TDD_r16) = v
			}
			if measAndMobParametersMRDCCommonV1610CondPSCellChangeParametersCommonR16Seq.IsComponentPresent(1) {
				c.CondPSCellChangeFR1_FR2_r16 = new(int64)
				v, err := d.DecodeEnumerated(measAndMobParametersMRDCCommonV1610CondPSCellChangeParametersCommonR16CondPSCellChangeFR1FR2R16Constraints)
				if err != nil {
					return err
				}
				(*c.CondPSCellChangeFR1_FR2_r16) = v
			}
		}
	}

	// 3. pscellT312-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(measAndMobParametersMRDCCommonV1610PscellT312R16Constraints)
			if err != nil {
				return err
			}
			ie.PscellT312_r16 = &idx
		}
	}

	return nil
}
