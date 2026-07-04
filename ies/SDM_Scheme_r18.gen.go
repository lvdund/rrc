// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SDM-Scheme-r18 (line 12511).

var sDMSchemeR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxRankSDM-r18", Optional: true},
		{Name: "maxRankSDM-DCI-0-2-r18", Optional: true},
	},
}

var sDMSchemeR18MaxRankSDMR18Constraints = per.Constrained(1, 2)

var sDMSchemeR18MaxRankSDMDCI02R18Constraints = per.Constrained(1, 2)

type SDM_Scheme_r18 struct {
	MaxRankSDM_r18         *int64
	MaxRankSDM_DCI_0_2_r18 *int64
}

func (ie *SDM_Scheme_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sDMSchemeR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MaxRankSDM_r18 != nil, ie.MaxRankSDM_DCI_0_2_r18 != nil}); err != nil {
		return err
	}

	// 2. maxRankSDM-r18: integer
	{
		if ie.MaxRankSDM_r18 != nil {
			if err := e.EncodeInteger(*ie.MaxRankSDM_r18, sDMSchemeR18MaxRankSDMR18Constraints); err != nil {
				return err
			}
		}
	}

	// 3. maxRankSDM-DCI-0-2-r18: integer
	{
		if ie.MaxRankSDM_DCI_0_2_r18 != nil {
			if err := e.EncodeInteger(*ie.MaxRankSDM_DCI_0_2_r18, sDMSchemeR18MaxRankSDMDCI02R18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SDM_Scheme_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sDMSchemeR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. maxRankSDM-r18: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(sDMSchemeR18MaxRankSDMR18Constraints)
			if err != nil {
				return err
			}
			ie.MaxRankSDM_r18 = &v
		}
	}

	// 3. maxRankSDM-DCI-0-2-r18: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(sDMSchemeR18MaxRankSDMDCI02R18Constraints)
			if err != nil {
				return err
			}
			ie.MaxRankSDM_DCI_0_2_r18 = &v
		}
	}

	return nil
}
