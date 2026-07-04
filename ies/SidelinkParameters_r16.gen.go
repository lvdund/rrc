// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SidelinkParameters-r16 (line 25015).

var sidelinkParametersR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sidelinkParametersNR-r16", Optional: true},
		{Name: "sidelinkParametersEUTRA-r16", Optional: true},
	},
}

type SidelinkParameters_r16 struct {
	SidelinkParametersNR_r16    *SidelinkParametersNR_r16
	SidelinkParametersEUTRA_r16 *SidelinkParametersEUTRA_r16
}

func (ie *SidelinkParameters_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sidelinkParametersR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SidelinkParametersNR_r16 != nil, ie.SidelinkParametersEUTRA_r16 != nil}); err != nil {
		return err
	}

	// 2. sidelinkParametersNR-r16: ref
	{
		if ie.SidelinkParametersNR_r16 != nil {
			if err := ie.SidelinkParametersNR_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. sidelinkParametersEUTRA-r16: ref
	{
		if ie.SidelinkParametersEUTRA_r16 != nil {
			if err := ie.SidelinkParametersEUTRA_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SidelinkParameters_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sidelinkParametersR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sidelinkParametersNR-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.SidelinkParametersNR_r16 = new(SidelinkParametersNR_r16)
			if err := ie.SidelinkParametersNR_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. sidelinkParametersEUTRA-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.SidelinkParametersEUTRA_r16 = new(SidelinkParametersEUTRA_r16)
			if err := ie.SidelinkParametersEUTRA_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
