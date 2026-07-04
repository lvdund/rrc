// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SFN-Scheme-r18 (line 12517).

var sFNSchemeR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxRankSFN-r18", Optional: true},
		{Name: "maxRankSFN-DCI-0-2-r18", Optional: true},
	},
}

var sFNSchemeR18MaxRankSFNR18Constraints = per.Constrained(1, 2)

var sFNSchemeR18MaxRankSFNDCI02R18Constraints = per.Constrained(1, 2)

type SFN_Scheme_r18 struct {
	MaxRankSFN_r18         *int64
	MaxRankSFN_DCI_0_2_r18 *int64
}

func (ie *SFN_Scheme_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sFNSchemeR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MaxRankSFN_r18 != nil, ie.MaxRankSFN_DCI_0_2_r18 != nil}); err != nil {
		return err
	}

	// 2. maxRankSFN-r18: integer
	{
		if ie.MaxRankSFN_r18 != nil {
			if err := e.EncodeInteger(*ie.MaxRankSFN_r18, sFNSchemeR18MaxRankSFNR18Constraints); err != nil {
				return err
			}
		}
	}

	// 3. maxRankSFN-DCI-0-2-r18: integer
	{
		if ie.MaxRankSFN_DCI_0_2_r18 != nil {
			if err := e.EncodeInteger(*ie.MaxRankSFN_DCI_0_2_r18, sFNSchemeR18MaxRankSFNDCI02R18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SFN_Scheme_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sFNSchemeR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. maxRankSFN-r18: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(sFNSchemeR18MaxRankSFNR18Constraints)
			if err != nil {
				return err
			}
			ie.MaxRankSFN_r18 = &v
		}
	}

	// 3. maxRankSFN-DCI-0-2-r18: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(sFNSchemeR18MaxRankSFNDCI02R18Constraints)
			if err != nil {
				return err
			}
			ie.MaxRankSFN_DCI_0_2_r18 = &v
		}
	}

	return nil
}
