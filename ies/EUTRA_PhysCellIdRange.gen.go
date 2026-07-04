// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: EUTRA-PhysCellIdRange (line 26193).

var eUTRAPhysCellIdRangeConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "start"},
		{Name: "range", Optional: true},
	},
}

const (
	EUTRA_PhysCellIdRange_Range_N4     = 0
	EUTRA_PhysCellIdRange_Range_N8     = 1
	EUTRA_PhysCellIdRange_Range_N12    = 2
	EUTRA_PhysCellIdRange_Range_N16    = 3
	EUTRA_PhysCellIdRange_Range_N24    = 4
	EUTRA_PhysCellIdRange_Range_N32    = 5
	EUTRA_PhysCellIdRange_Range_N48    = 6
	EUTRA_PhysCellIdRange_Range_N64    = 7
	EUTRA_PhysCellIdRange_Range_N84    = 8
	EUTRA_PhysCellIdRange_Range_N96    = 9
	EUTRA_PhysCellIdRange_Range_N128   = 10
	EUTRA_PhysCellIdRange_Range_N168   = 11
	EUTRA_PhysCellIdRange_Range_N252   = 12
	EUTRA_PhysCellIdRange_Range_N504   = 13
	EUTRA_PhysCellIdRange_Range_Spare2 = 14
	EUTRA_PhysCellIdRange_Range_Spare1 = 15
)

var eUTRAPhysCellIdRangeRangeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{EUTRA_PhysCellIdRange_Range_N4, EUTRA_PhysCellIdRange_Range_N8, EUTRA_PhysCellIdRange_Range_N12, EUTRA_PhysCellIdRange_Range_N16, EUTRA_PhysCellIdRange_Range_N24, EUTRA_PhysCellIdRange_Range_N32, EUTRA_PhysCellIdRange_Range_N48, EUTRA_PhysCellIdRange_Range_N64, EUTRA_PhysCellIdRange_Range_N84, EUTRA_PhysCellIdRange_Range_N96, EUTRA_PhysCellIdRange_Range_N128, EUTRA_PhysCellIdRange_Range_N168, EUTRA_PhysCellIdRange_Range_N252, EUTRA_PhysCellIdRange_Range_N504, EUTRA_PhysCellIdRange_Range_Spare2, EUTRA_PhysCellIdRange_Range_Spare1},
}

type EUTRA_PhysCellIdRange struct {
	Start EUTRA_PhysCellId
	Range *int64
}

func (ie *EUTRA_PhysCellIdRange) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(eUTRAPhysCellIdRangeConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Range != nil}); err != nil {
		return err
	}

	// 2. start: ref
	{
		if err := ie.Start.Encode(e); err != nil {
			return err
		}
	}

	// 3. range: enumerated
	{
		if ie.Range != nil {
			if err := e.EncodeEnumerated(*ie.Range, eUTRAPhysCellIdRangeRangeConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *EUTRA_PhysCellIdRange) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(eUTRAPhysCellIdRangeConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. start: ref
	{
		if err := ie.Start.Decode(d); err != nil {
			return err
		}
	}

	// 3. range: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(eUTRAPhysCellIdRangeRangeConstraints)
			if err != nil {
				return err
			}
			ie.Range = &idx
		}
	}

	return nil
}
