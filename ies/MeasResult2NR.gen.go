// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasResult2NR (line 9903).

var measResult2NRConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ssbFrequency", Optional: true},
		{Name: "refFreqCSI-RS", Optional: true},
		{Name: "measResultServingCell", Optional: true},
		{Name: "measResultNeighCellListNR", Optional: true},
	},
}

type MeasResult2NR struct {
	SsbFrequency              *ARFCN_ValueNR
	RefFreqCSI_RS             *ARFCN_ValueNR
	MeasResultServingCell     *MeasResultNR
	MeasResultNeighCellListNR *MeasResultListNR
}

func (ie *MeasResult2NR) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measResult2NRConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SsbFrequency != nil, ie.RefFreqCSI_RS != nil, ie.MeasResultServingCell != nil, ie.MeasResultNeighCellListNR != nil}); err != nil {
		return err
	}

	// 3. ssbFrequency: ref
	{
		if ie.SsbFrequency != nil {
			if err := ie.SsbFrequency.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. refFreqCSI-RS: ref
	{
		if ie.RefFreqCSI_RS != nil {
			if err := ie.RefFreqCSI_RS.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. measResultServingCell: ref
	{
		if ie.MeasResultServingCell != nil {
			if err := ie.MeasResultServingCell.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. measResultNeighCellListNR: ref
	{
		if ie.MeasResultNeighCellListNR != nil {
			if err := ie.MeasResultNeighCellListNR.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MeasResult2NR) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measResult2NRConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. ssbFrequency: ref
	{
		if seq.IsComponentPresent(0) {
			ie.SsbFrequency = new(ARFCN_ValueNR)
			if err := ie.SsbFrequency.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. refFreqCSI-RS: ref
	{
		if seq.IsComponentPresent(1) {
			ie.RefFreqCSI_RS = new(ARFCN_ValueNR)
			if err := ie.RefFreqCSI_RS.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. measResultServingCell: ref
	{
		if seq.IsComponentPresent(2) {
			ie.MeasResultServingCell = new(MeasResultNR)
			if err := ie.MeasResultServingCell.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. measResultNeighCellListNR: ref
	{
		if seq.IsComponentPresent(3) {
			ie.MeasResultNeighCellListNR = new(MeasResultListNR)
			if err := ie.MeasResultNeighCellListNR.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
