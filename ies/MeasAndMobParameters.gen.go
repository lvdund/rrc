// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasAndMobParameters (line 21098).

var measAndMobParametersConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "measAndMobParametersCommon", Optional: true},
		{Name: "measAndMobParametersXDD-Diff", Optional: true},
		{Name: "measAndMobParametersFRX-Diff", Optional: true},
	},
}

type MeasAndMobParameters struct {
	MeasAndMobParametersCommon   *MeasAndMobParametersCommon
	MeasAndMobParametersXDD_Diff *MeasAndMobParametersXDD_Diff
	MeasAndMobParametersFRX_Diff *MeasAndMobParametersFRX_Diff
}

func (ie *MeasAndMobParameters) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measAndMobParametersConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MeasAndMobParametersCommon != nil, ie.MeasAndMobParametersXDD_Diff != nil, ie.MeasAndMobParametersFRX_Diff != nil}); err != nil {
		return err
	}

	// 2. measAndMobParametersCommon: ref
	{
		if ie.MeasAndMobParametersCommon != nil {
			if err := ie.MeasAndMobParametersCommon.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. measAndMobParametersXDD-Diff: ref
	{
		if ie.MeasAndMobParametersXDD_Diff != nil {
			if err := ie.MeasAndMobParametersXDD_Diff.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. measAndMobParametersFRX-Diff: ref
	{
		if ie.MeasAndMobParametersFRX_Diff != nil {
			if err := ie.MeasAndMobParametersFRX_Diff.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MeasAndMobParameters) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measAndMobParametersConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. measAndMobParametersCommon: ref
	{
		if seq.IsComponentPresent(0) {
			ie.MeasAndMobParametersCommon = new(MeasAndMobParametersCommon)
			if err := ie.MeasAndMobParametersCommon.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. measAndMobParametersXDD-Diff: ref
	{
		if seq.IsComponentPresent(1) {
			ie.MeasAndMobParametersXDD_Diff = new(MeasAndMobParametersXDD_Diff)
			if err := ie.MeasAndMobParametersXDD_Diff.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. measAndMobParametersFRX-Diff: ref
	{
		if seq.IsComponentPresent(2) {
			ie.MeasAndMobParametersFRX_Diff = new(MeasAndMobParametersFRX_Diff)
			if err := ie.MeasAndMobParametersFRX_Diff.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
