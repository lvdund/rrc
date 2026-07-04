// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PDCCH-BlindDetectionMixed-r17 (line 18267).

var pDCCHBlindDetectionMixedR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pdcch-BlindDetectionCA-Mixed-r17", Optional: true},
		{Name: "pdcch-BlindDetectionCG-UE-Mixed-r17", Optional: true},
	},
}

type PDCCH_BlindDetectionMixed_r17 struct {
	Pdcch_BlindDetectionCA_Mixed_r17    *PDCCH_BlindDetectionCA_Mixed_r17
	Pdcch_BlindDetectionCG_UE_Mixed_r17 *struct {
		Pdcch_BlindDetectionMCG_UE_Mixed_v17 PDCCH_BlindDetectionCG_UE_Mixed_r17
		Pdcch_BlindDetectionSCG_UE_Mixed_v17 PDCCH_BlindDetectionCG_UE_Mixed_r17
	}
}

func (ie *PDCCH_BlindDetectionMixed_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDCCHBlindDetectionMixedR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Pdcch_BlindDetectionCA_Mixed_r17 != nil, ie.Pdcch_BlindDetectionCG_UE_Mixed_r17 != nil}); err != nil {
		return err
	}

	// 2. pdcch-BlindDetectionCA-Mixed-r17: ref
	{
		if ie.Pdcch_BlindDetectionCA_Mixed_r17 != nil {
			if err := ie.Pdcch_BlindDetectionCA_Mixed_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. pdcch-BlindDetectionCG-UE-Mixed-r17: sequence
	{
		if ie.Pdcch_BlindDetectionCG_UE_Mixed_r17 != nil {
			c := ie.Pdcch_BlindDetectionCG_UE_Mixed_r17
			if err := c.Pdcch_BlindDetectionMCG_UE_Mixed_v17.Encode(e); err != nil {
				return err
			}
			if err := c.Pdcch_BlindDetectionSCG_UE_Mixed_v17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PDCCH_BlindDetectionMixed_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDCCHBlindDetectionMixedR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. pdcch-BlindDetectionCA-Mixed-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Pdcch_BlindDetectionCA_Mixed_r17 = new(PDCCH_BlindDetectionCA_Mixed_r17)
			if err := ie.Pdcch_BlindDetectionCA_Mixed_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. pdcch-BlindDetectionCG-UE-Mixed-r17: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.Pdcch_BlindDetectionCG_UE_Mixed_r17 = &struct {
				Pdcch_BlindDetectionMCG_UE_Mixed_v17 PDCCH_BlindDetectionCG_UE_Mixed_r17
				Pdcch_BlindDetectionSCG_UE_Mixed_v17 PDCCH_BlindDetectionCG_UE_Mixed_r17
			}{}
			c := ie.Pdcch_BlindDetectionCG_UE_Mixed_r17
			{
				if err := c.Pdcch_BlindDetectionMCG_UE_Mixed_v17.Decode(d); err != nil {
					return err
				}
			}
			{
				if err := c.Pdcch_BlindDetectionSCG_UE_Mixed_v17.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
