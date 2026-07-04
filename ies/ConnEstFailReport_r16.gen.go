// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ConnEstFailReport-r16 (line 3023).

var connEstFailReportR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "measResultFailedCell-r16"},
		{Name: "locationInfo-r16", Optional: true},
		{Name: "measResultNeighCells-r16"},
		{Name: "numberOfConnFail-r16"},
		{Name: "perRAInfoList-r16"},
		{Name: "timeSinceFailure-r16"},
	},
}

var connEstFailReportR16NumberOfConnFailR16Constraints = per.Constrained(1, 8)

var connEstFailReportR16MeasResultNeighCellsR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "measResultNeighCellListNR", Optional: true},
		{Name: "measResultNeighCellListEUTRA", Optional: true},
	},
}

type ConnEstFailReport_r16 struct {
	MeasResultFailedCell_r16 MeasResultFailedCell_r16
	LocationInfo_r16         *LocationInfo_r16
	MeasResultNeighCells_r16 struct {
		MeasResultNeighCellListNR    *MeasResultList2NR_r16
		MeasResultNeighCellListEUTRA *MeasResultList2EUTRA_r16
	}
	NumberOfConnFail_r16 int64
	PerRAInfoList_r16    PerRAInfoList_r16
	TimeSinceFailure_r16 TimeSinceFailure_r16
}

func (ie *ConnEstFailReport_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(connEstFailReportR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.LocationInfo_r16 != nil}); err != nil {
		return err
	}

	// 3. measResultFailedCell-r16: ref
	{
		if err := ie.MeasResultFailedCell_r16.Encode(e); err != nil {
			return err
		}
	}

	// 4. locationInfo-r16: ref
	{
		if ie.LocationInfo_r16 != nil {
			if err := ie.LocationInfo_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. measResultNeighCells-r16: sequence
	{
		{
			c := &ie.MeasResultNeighCells_r16
			connEstFailReportR16MeasResultNeighCellsR16Seq := e.NewSequenceEncoder(connEstFailReportR16MeasResultNeighCellsR16Constraints)
			if err := connEstFailReportR16MeasResultNeighCellsR16Seq.EncodePreamble([]bool{c.MeasResultNeighCellListNR != nil, c.MeasResultNeighCellListEUTRA != nil}); err != nil {
				return err
			}
			if c.MeasResultNeighCellListNR != nil {
				if err := c.MeasResultNeighCellListNR.Encode(e); err != nil {
					return err
				}
			}
			if c.MeasResultNeighCellListEUTRA != nil {
				if err := c.MeasResultNeighCellListEUTRA.Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 6. numberOfConnFail-r16: integer
	{
		if err := e.EncodeInteger(ie.NumberOfConnFail_r16, connEstFailReportR16NumberOfConnFailR16Constraints); err != nil {
			return err
		}
	}

	// 7. perRAInfoList-r16: ref
	{
		if err := ie.PerRAInfoList_r16.Encode(e); err != nil {
			return err
		}
	}

	// 8. timeSinceFailure-r16: ref
	{
		if err := ie.TimeSinceFailure_r16.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *ConnEstFailReport_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(connEstFailReportR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. measResultFailedCell-r16: ref
	{
		if err := ie.MeasResultFailedCell_r16.Decode(d); err != nil {
			return err
		}
	}

	// 4. locationInfo-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.LocationInfo_r16 = new(LocationInfo_r16)
			if err := ie.LocationInfo_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. measResultNeighCells-r16: sequence
	{
		{
			c := &ie.MeasResultNeighCells_r16
			connEstFailReportR16MeasResultNeighCellsR16Seq := d.NewSequenceDecoder(connEstFailReportR16MeasResultNeighCellsR16Constraints)
			if err := connEstFailReportR16MeasResultNeighCellsR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if connEstFailReportR16MeasResultNeighCellsR16Seq.IsComponentPresent(0) {
				c.MeasResultNeighCellListNR = new(MeasResultList2NR_r16)
				if err := (*c.MeasResultNeighCellListNR).Decode(d); err != nil {
					return err
				}
			}
			if connEstFailReportR16MeasResultNeighCellsR16Seq.IsComponentPresent(1) {
				c.MeasResultNeighCellListEUTRA = new(MeasResultList2EUTRA_r16)
				if err := (*c.MeasResultNeighCellListEUTRA).Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. numberOfConnFail-r16: integer
	{
		v3, err := d.DecodeInteger(connEstFailReportR16NumberOfConnFailR16Constraints)
		if err != nil {
			return err
		}
		ie.NumberOfConnFail_r16 = v3
	}

	// 7. perRAInfoList-r16: ref
	{
		if err := ie.PerRAInfoList_r16.Decode(d); err != nil {
			return err
		}
	}

	// 8. timeSinceFailure-r16: ref
	{
		if err := ie.TimeSinceFailure_r16.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
