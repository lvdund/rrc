// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PDCCH-BlindDetectionCG-UE-Mixed-r17 (line 18275).

var pDCCHBlindDetectionCGUEMixedR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pdcch-BlindDetectionCG-UE1-r17"},
		{Name: "pdcch-BlindDetectionCG-UE2-r17"},
	},
}

var pDCCHBlindDetectionCGUEMixedR17PdcchBlindDetectionCGUE1R17Constraints = per.Constrained(0, 15)

var pDCCHBlindDetectionCGUEMixedR17PdcchBlindDetectionCGUE2R17Constraints = per.Constrained(0, 15)

type PDCCH_BlindDetectionCG_UE_Mixed_r17 struct {
	Pdcch_BlindDetectionCG_UE1_r17 int64
	Pdcch_BlindDetectionCG_UE2_r17 int64
}

func (ie *PDCCH_BlindDetectionCG_UE_Mixed_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDCCHBlindDetectionCGUEMixedR17Constraints)
	_ = seq

	// 1. pdcch-BlindDetectionCG-UE1-r17: integer
	{
		if err := e.EncodeInteger(ie.Pdcch_BlindDetectionCG_UE1_r17, pDCCHBlindDetectionCGUEMixedR17PdcchBlindDetectionCGUE1R17Constraints); err != nil {
			return err
		}
	}

	// 2. pdcch-BlindDetectionCG-UE2-r17: integer
	{
		if err := e.EncodeInteger(ie.Pdcch_BlindDetectionCG_UE2_r17, pDCCHBlindDetectionCGUEMixedR17PdcchBlindDetectionCGUE2R17Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PDCCH_BlindDetectionCG_UE_Mixed_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDCCHBlindDetectionCGUEMixedR17Constraints)
	_ = seq

	// 1. pdcch-BlindDetectionCG-UE1-r17: integer
	{
		v0, err := d.DecodeInteger(pDCCHBlindDetectionCGUEMixedR17PdcchBlindDetectionCGUE1R17Constraints)
		if err != nil {
			return err
		}
		ie.Pdcch_BlindDetectionCG_UE1_r17 = v0
	}

	// 2. pdcch-BlindDetectionCG-UE2-r17: integer
	{
		v1, err := d.DecodeInteger(pDCCHBlindDetectionCGUEMixedR17PdcchBlindDetectionCGUE2R17Constraints)
		if err != nil {
			return err
		}
		ie.Pdcch_BlindDetectionCG_UE2_r17 = v1
	}

	return nil
}
