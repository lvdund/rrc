// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PDCCH-BlindDetectionMixed2-r18 (line 18304).

var pDCCHBlindDetectionMixed2R18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pdcch-BlindDetectionMCG-UE-Mixed-r18"},
		{Name: "pdcch-BlindDetectionSCG-UE-Mixed-r18"},
	},
}

type PDCCH_BlindDetectionMixed2_r18 struct {
	Pdcch_BlindDetectionMCG_UE_Mixed_r18 PDCCH_BlindDetectionCG_UE_MixedExt_r16
	Pdcch_BlindDetectionSCG_UE_Mixed_r18 PDCCH_BlindDetectionCG_UE_MixedExt_r16
}

func (ie *PDCCH_BlindDetectionMixed2_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDCCHBlindDetectionMixed2R18Constraints)
	_ = seq

	// 1. pdcch-BlindDetectionMCG-UE-Mixed-r18: ref
	{
		if err := ie.Pdcch_BlindDetectionMCG_UE_Mixed_r18.Encode(e); err != nil {
			return err
		}
	}

	// 2. pdcch-BlindDetectionSCG-UE-Mixed-r18: ref
	{
		if err := ie.Pdcch_BlindDetectionSCG_UE_Mixed_r18.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PDCCH_BlindDetectionMixed2_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDCCHBlindDetectionMixed2R18Constraints)
	_ = seq

	// 1. pdcch-BlindDetectionMCG-UE-Mixed-r18: ref
	{
		if err := ie.Pdcch_BlindDetectionMCG_UE_Mixed_r18.Decode(d); err != nil {
			return err
		}
	}

	// 2. pdcch-BlindDetectionSCG-UE-Mixed-r18: ref
	{
		if err := ie.Pdcch_BlindDetectionSCG_UE_Mixed_r18.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
