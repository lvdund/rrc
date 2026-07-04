// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasAndMobParametersMRDC-XDD-Diff-v1560 (line 21461).

var measAndMobParametersMRDCXDDDiffV1560Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sftd-MeasPSCell-NEDC", Optional: true},
	},
}

const (
	MeasAndMobParametersMRDC_XDD_Diff_v1560_Sftd_MeasPSCell_NEDC_Supported = 0
)

var measAndMobParametersMRDCXDDDiffV1560SftdMeasPSCellNEDCConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersMRDC_XDD_Diff_v1560_Sftd_MeasPSCell_NEDC_Supported},
}

type MeasAndMobParametersMRDC_XDD_Diff_v1560 struct {
	Sftd_MeasPSCell_NEDC *int64
}

func (ie *MeasAndMobParametersMRDC_XDD_Diff_v1560) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measAndMobParametersMRDCXDDDiffV1560Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sftd_MeasPSCell_NEDC != nil}); err != nil {
		return err
	}

	// 2. sftd-MeasPSCell-NEDC: enumerated
	{
		if ie.Sftd_MeasPSCell_NEDC != nil {
			if err := e.EncodeEnumerated(*ie.Sftd_MeasPSCell_NEDC, measAndMobParametersMRDCXDDDiffV1560SftdMeasPSCellNEDCConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MeasAndMobParametersMRDC_XDD_Diff_v1560) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measAndMobParametersMRDCXDDDiffV1560Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sftd-MeasPSCell-NEDC: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(measAndMobParametersMRDCXDDDiffV1560SftdMeasPSCellNEDCConstraints)
			if err != nil {
				return err
			}
			ie.Sftd_MeasPSCell_NEDC = &idx
		}
	}

	return nil
}
