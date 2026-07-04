// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PCI-Range (line 10962).

var pCIRangeConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "start"},
		{Name: "range", Optional: true},
	},
}

const (
	PCI_Range_Range_N4     = 0
	PCI_Range_Range_N8     = 1
	PCI_Range_Range_N12    = 2
	PCI_Range_Range_N16    = 3
	PCI_Range_Range_N24    = 4
	PCI_Range_Range_N32    = 5
	PCI_Range_Range_N48    = 6
	PCI_Range_Range_N64    = 7
	PCI_Range_Range_N84    = 8
	PCI_Range_Range_N96    = 9
	PCI_Range_Range_N128   = 10
	PCI_Range_Range_N168   = 11
	PCI_Range_Range_N252   = 12
	PCI_Range_Range_N504   = 13
	PCI_Range_Range_N1008  = 14
	PCI_Range_Range_Spare1 = 15
)

var pCIRangeRangeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PCI_Range_Range_N4, PCI_Range_Range_N8, PCI_Range_Range_N12, PCI_Range_Range_N16, PCI_Range_Range_N24, PCI_Range_Range_N32, PCI_Range_Range_N48, PCI_Range_Range_N64, PCI_Range_Range_N84, PCI_Range_Range_N96, PCI_Range_Range_N128, PCI_Range_Range_N168, PCI_Range_Range_N252, PCI_Range_Range_N504, PCI_Range_Range_N1008, PCI_Range_Range_Spare1},
}

type PCI_Range struct {
	Start PhysCellId
	Range *int64
}

func (ie *PCI_Range) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pCIRangeConstraints)

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
			if err := e.EncodeEnumerated(*ie.Range, pCIRangeRangeConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PCI_Range) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pCIRangeConstraints)

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
			idx, err := d.DecodeEnumerated(pCIRangeRangeConstraints)
			if err != nil {
				return err
			}
			ie.Range = &idx
		}
	}

	return nil
}
