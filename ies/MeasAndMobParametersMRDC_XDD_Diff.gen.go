// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasAndMobParametersMRDC-XDD-Diff (line 21456).

var measAndMobParametersMRDCXDDDiffConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sftd-MeasPSCell", Optional: true},
		{Name: "sftd-MeasNR-Cell", Optional: true},
	},
}

const (
	MeasAndMobParametersMRDC_XDD_Diff_Sftd_MeasPSCell_Supported = 0
)

var measAndMobParametersMRDCXDDDiffSftdMeasPSCellConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersMRDC_XDD_Diff_Sftd_MeasPSCell_Supported},
}

const (
	MeasAndMobParametersMRDC_XDD_Diff_Sftd_MeasNR_Cell_Supported = 0
)

var measAndMobParametersMRDCXDDDiffSftdMeasNRCellConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersMRDC_XDD_Diff_Sftd_MeasNR_Cell_Supported},
}

type MeasAndMobParametersMRDC_XDD_Diff struct {
	Sftd_MeasPSCell  *int64
	Sftd_MeasNR_Cell *int64
}

func (ie *MeasAndMobParametersMRDC_XDD_Diff) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measAndMobParametersMRDCXDDDiffConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sftd_MeasPSCell != nil, ie.Sftd_MeasNR_Cell != nil}); err != nil {
		return err
	}

	// 2. sftd-MeasPSCell: enumerated
	{
		if ie.Sftd_MeasPSCell != nil {
			if err := e.EncodeEnumerated(*ie.Sftd_MeasPSCell, measAndMobParametersMRDCXDDDiffSftdMeasPSCellConstraints); err != nil {
				return err
			}
		}
	}

	// 3. sftd-MeasNR-Cell: enumerated
	{
		if ie.Sftd_MeasNR_Cell != nil {
			if err := e.EncodeEnumerated(*ie.Sftd_MeasNR_Cell, measAndMobParametersMRDCXDDDiffSftdMeasNRCellConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MeasAndMobParametersMRDC_XDD_Diff) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measAndMobParametersMRDCXDDDiffConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sftd-MeasPSCell: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(measAndMobParametersMRDCXDDDiffSftdMeasPSCellConstraints)
			if err != nil {
				return err
			}
			ie.Sftd_MeasPSCell = &idx
		}
	}

	// 3. sftd-MeasNR-Cell: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(measAndMobParametersMRDCXDDDiffSftdMeasNRCellConstraints)
			if err != nil {
				return err
			}
			ie.Sftd_MeasNR_Cell = &idx
		}
	}

	return nil
}
