// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasAndMobParametersMRDC (line 21371).

var measAndMobParametersMRDCConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "measAndMobParametersMRDC-Common", Optional: true},
		{Name: "measAndMobParametersMRDC-XDD-Diff", Optional: true},
		{Name: "measAndMobParametersMRDC-FRX-Diff", Optional: true},
	},
}

type MeasAndMobParametersMRDC struct {
	MeasAndMobParametersMRDC_Common   *MeasAndMobParametersMRDC_Common
	MeasAndMobParametersMRDC_XDD_Diff *MeasAndMobParametersMRDC_XDD_Diff
	MeasAndMobParametersMRDC_FRX_Diff *MeasAndMobParametersMRDC_FRX_Diff
}

func (ie *MeasAndMobParametersMRDC) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measAndMobParametersMRDCConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MeasAndMobParametersMRDC_Common != nil, ie.MeasAndMobParametersMRDC_XDD_Diff != nil, ie.MeasAndMobParametersMRDC_FRX_Diff != nil}); err != nil {
		return err
	}

	// 2. measAndMobParametersMRDC-Common: ref
	{
		if ie.MeasAndMobParametersMRDC_Common != nil {
			if err := ie.MeasAndMobParametersMRDC_Common.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. measAndMobParametersMRDC-XDD-Diff: ref
	{
		if ie.MeasAndMobParametersMRDC_XDD_Diff != nil {
			if err := ie.MeasAndMobParametersMRDC_XDD_Diff.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. measAndMobParametersMRDC-FRX-Diff: ref
	{
		if ie.MeasAndMobParametersMRDC_FRX_Diff != nil {
			if err := ie.MeasAndMobParametersMRDC_FRX_Diff.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MeasAndMobParametersMRDC) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measAndMobParametersMRDCConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. measAndMobParametersMRDC-Common: ref
	{
		if seq.IsComponentPresent(0) {
			ie.MeasAndMobParametersMRDC_Common = new(MeasAndMobParametersMRDC_Common)
			if err := ie.MeasAndMobParametersMRDC_Common.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. measAndMobParametersMRDC-XDD-Diff: ref
	{
		if seq.IsComponentPresent(1) {
			ie.MeasAndMobParametersMRDC_XDD_Diff = new(MeasAndMobParametersMRDC_XDD_Diff)
			if err := ie.MeasAndMobParametersMRDC_XDD_Diff.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. measAndMobParametersMRDC-FRX-Diff: ref
	{
		if seq.IsComponentPresent(2) {
			ie.MeasAndMobParametersMRDC_FRX_Diff = new(MeasAndMobParametersMRDC_FRX_Diff)
			if err := ie.MeasAndMobParametersMRDC_FRX_Diff.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
