// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PDCCH-BlindDetectionMCG-SCG-r17 (line 18262).

var pDCCHBlindDetectionMCGSCGR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pdcch-BlindDetectionMCG-UE-r17"},
		{Name: "pdcch-BlindDetectionSCG-UE-r17"},
	},
}

var pDCCHBlindDetectionMCGSCGR17PdcchBlindDetectionMCGUER17Constraints = per.Constrained(1, 15)

var pDCCHBlindDetectionMCGSCGR17PdcchBlindDetectionSCGUER17Constraints = per.Constrained(1, 15)

type PDCCH_BlindDetectionMCG_SCG_r17 struct {
	Pdcch_BlindDetectionMCG_UE_r17 int64
	Pdcch_BlindDetectionSCG_UE_r17 int64
}

func (ie *PDCCH_BlindDetectionMCG_SCG_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDCCHBlindDetectionMCGSCGR17Constraints)
	_ = seq

	// 1. pdcch-BlindDetectionMCG-UE-r17: integer
	{
		if err := e.EncodeInteger(ie.Pdcch_BlindDetectionMCG_UE_r17, pDCCHBlindDetectionMCGSCGR17PdcchBlindDetectionMCGUER17Constraints); err != nil {
			return err
		}
	}

	// 2. pdcch-BlindDetectionSCG-UE-r17: integer
	{
		if err := e.EncodeInteger(ie.Pdcch_BlindDetectionSCG_UE_r17, pDCCHBlindDetectionMCGSCGR17PdcchBlindDetectionSCGUER17Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PDCCH_BlindDetectionMCG_SCG_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDCCHBlindDetectionMCGSCGR17Constraints)
	_ = seq

	// 1. pdcch-BlindDetectionMCG-UE-r17: integer
	{
		v0, err := d.DecodeInteger(pDCCHBlindDetectionMCGSCGR17PdcchBlindDetectionMCGUER17Constraints)
		if err != nil {
			return err
		}
		ie.Pdcch_BlindDetectionMCG_UE_r17 = v0
	}

	// 2. pdcch-BlindDetectionSCG-UE-r17: integer
	{
		v1, err := d.DecodeInteger(pDCCHBlindDetectionMCGSCGR17PdcchBlindDetectionSCGUER17Constraints)
		if err != nil {
			return err
		}
		ie.Pdcch_BlindDetectionSCG_UE_r17 = v1
	}

	return nil
}
