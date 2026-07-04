// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasResultSFTD-EUTRA (line 9701).

var measResultSFTDEUTRAConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "eutra-PhysCellId"},
		{Name: "sfn-OffsetResult"},
		{Name: "frameBoundaryOffsetResult"},
		{Name: "rsrp-Result", Optional: true},
	},
}

var measResultSFTDEUTRASfnOffsetResultConstraints = per.Constrained(0, 1023)

var measResultSFTDEUTRAFrameBoundaryOffsetResultConstraints = per.Constrained(-30720, 30719)

type MeasResultSFTD_EUTRA struct {
	Eutra_PhysCellId          EUTRA_PhysCellId
	Sfn_OffsetResult          int64
	FrameBoundaryOffsetResult int64
	Rsrp_Result               *RSRP_Range
}

func (ie *MeasResultSFTD_EUTRA) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measResultSFTDEUTRAConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Rsrp_Result != nil}); err != nil {
		return err
	}

	// 2. eutra-PhysCellId: ref
	{
		if err := ie.Eutra_PhysCellId.Encode(e); err != nil {
			return err
		}
	}

	// 3. sfn-OffsetResult: integer
	{
		if err := e.EncodeInteger(ie.Sfn_OffsetResult, measResultSFTDEUTRASfnOffsetResultConstraints); err != nil {
			return err
		}
	}

	// 4. frameBoundaryOffsetResult: integer
	{
		if err := e.EncodeInteger(ie.FrameBoundaryOffsetResult, measResultSFTDEUTRAFrameBoundaryOffsetResultConstraints); err != nil {
			return err
		}
	}

	// 5. rsrp-Result: ref
	{
		if ie.Rsrp_Result != nil {
			if err := ie.Rsrp_Result.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MeasResultSFTD_EUTRA) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measResultSFTDEUTRAConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. eutra-PhysCellId: ref
	{
		if err := ie.Eutra_PhysCellId.Decode(d); err != nil {
			return err
		}
	}

	// 3. sfn-OffsetResult: integer
	{
		v1, err := d.DecodeInteger(measResultSFTDEUTRASfnOffsetResultConstraints)
		if err != nil {
			return err
		}
		ie.Sfn_OffsetResult = v1
	}

	// 4. frameBoundaryOffsetResult: integer
	{
		v2, err := d.DecodeInteger(measResultSFTDEUTRAFrameBoundaryOffsetResultConstraints)
		if err != nil {
			return err
		}
		ie.FrameBoundaryOffsetResult = v2
	}

	// 5. rsrp-Result: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Rsrp_Result = new(RSRP_Range)
			if err := ie.Rsrp_Result.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
