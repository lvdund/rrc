// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: LTM-ReportContent-r18 (line 8863).

var lTMReportContentR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "nrOfReportedCells-r18"},
		{Name: "nrOfReportedRS-PerCell-r18"},
		{Name: "spCellInclusion-r18", Optional: true},
	},
}

const (
	LTM_ReportContent_r18_NrOfReportedCells_r18_N1 = 0
	LTM_ReportContent_r18_NrOfReportedCells_r18_N2 = 1
	LTM_ReportContent_r18_NrOfReportedCells_r18_N3 = 2
	LTM_ReportContent_r18_NrOfReportedCells_r18_N4 = 3
)

var lTMReportContentR18NrOfReportedCellsR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LTM_ReportContent_r18_NrOfReportedCells_r18_N1, LTM_ReportContent_r18_NrOfReportedCells_r18_N2, LTM_ReportContent_r18_NrOfReportedCells_r18_N3, LTM_ReportContent_r18_NrOfReportedCells_r18_N4},
}

const (
	LTM_ReportContent_r18_NrOfReportedRS_PerCell_r18_N1 = 0
	LTM_ReportContent_r18_NrOfReportedRS_PerCell_r18_N2 = 1
	LTM_ReportContent_r18_NrOfReportedRS_PerCell_r18_N3 = 2
	LTM_ReportContent_r18_NrOfReportedRS_PerCell_r18_N4 = 3
)

var lTMReportContentR18NrOfReportedRSPerCellR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LTM_ReportContent_r18_NrOfReportedRS_PerCell_r18_N1, LTM_ReportContent_r18_NrOfReportedRS_PerCell_r18_N2, LTM_ReportContent_r18_NrOfReportedRS_PerCell_r18_N3, LTM_ReportContent_r18_NrOfReportedRS_PerCell_r18_N4},
}

const (
	LTM_ReportContent_r18_SpCellInclusion_r18_True = 0
)

var lTMReportContentR18SpCellInclusionR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LTM_ReportContent_r18_SpCellInclusion_r18_True},
}

type LTM_ReportContent_r18 struct {
	NrOfReportedCells_r18      int64
	NrOfReportedRS_PerCell_r18 int64
	SpCellInclusion_r18        *int64
}

func (ie *LTM_ReportContent_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(lTMReportContentR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SpCellInclusion_r18 != nil}); err != nil {
		return err
	}

	// 2. nrOfReportedCells-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.NrOfReportedCells_r18, lTMReportContentR18NrOfReportedCellsR18Constraints); err != nil {
			return err
		}
	}

	// 3. nrOfReportedRS-PerCell-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.NrOfReportedRS_PerCell_r18, lTMReportContentR18NrOfReportedRSPerCellR18Constraints); err != nil {
			return err
		}
	}

	// 4. spCellInclusion-r18: enumerated
	{
		if ie.SpCellInclusion_r18 != nil {
			if err := e.EncodeEnumerated(*ie.SpCellInclusion_r18, lTMReportContentR18SpCellInclusionR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *LTM_ReportContent_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(lTMReportContentR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. nrOfReportedCells-r18: enumerated
	{
		v0, err := d.DecodeEnumerated(lTMReportContentR18NrOfReportedCellsR18Constraints)
		if err != nil {
			return err
		}
		ie.NrOfReportedCells_r18 = v0
	}

	// 3. nrOfReportedRS-PerCell-r18: enumerated
	{
		v1, err := d.DecodeEnumerated(lTMReportContentR18NrOfReportedRSPerCellR18Constraints)
		if err != nil {
			return err
		}
		ie.NrOfReportedRS_PerCell_r18 = v1
	}

	// 4. spCellInclusion-r18: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(lTMReportContentR18SpCellInclusionR18Constraints)
			if err != nil {
				return err
			}
			ie.SpCellInclusion_r18 = &idx
		}
	}

	return nil
}
