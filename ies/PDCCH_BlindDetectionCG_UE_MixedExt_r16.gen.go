// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PDCCH-BlindDetectionCG-UE-MixedExt-r16 (line 18257).

var pDCCHBlindDetectionCGUEMixedExtR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pdcch-BlindDetectionCG-UE1-r16"},
		{Name: "pdcch-BlindDetectionCG-UE2-r16"},
	},
}

var pDCCHBlindDetectionCGUEMixedExtR16PdcchBlindDetectionCGUE1R16Constraints = per.Constrained(0, 15)

var pDCCHBlindDetectionCGUEMixedExtR16PdcchBlindDetectionCGUE2R16Constraints = per.Constrained(0, 15)

type PDCCH_BlindDetectionCG_UE_MixedExt_r16 struct {
	Pdcch_BlindDetectionCG_UE1_r16 int64
	Pdcch_BlindDetectionCG_UE2_r16 int64
}

func (ie *PDCCH_BlindDetectionCG_UE_MixedExt_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDCCHBlindDetectionCGUEMixedExtR16Constraints)
	_ = seq

	// 1. pdcch-BlindDetectionCG-UE1-r16: integer
	{
		if err := e.EncodeInteger(ie.Pdcch_BlindDetectionCG_UE1_r16, pDCCHBlindDetectionCGUEMixedExtR16PdcchBlindDetectionCGUE1R16Constraints); err != nil {
			return err
		}
	}

	// 2. pdcch-BlindDetectionCG-UE2-r16: integer
	{
		if err := e.EncodeInteger(ie.Pdcch_BlindDetectionCG_UE2_r16, pDCCHBlindDetectionCGUEMixedExtR16PdcchBlindDetectionCGUE2R16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PDCCH_BlindDetectionCG_UE_MixedExt_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDCCHBlindDetectionCGUEMixedExtR16Constraints)
	_ = seq

	// 1. pdcch-BlindDetectionCG-UE1-r16: integer
	{
		v0, err := d.DecodeInteger(pDCCHBlindDetectionCGUEMixedExtR16PdcchBlindDetectionCGUE1R16Constraints)
		if err != nil {
			return err
		}
		ie.Pdcch_BlindDetectionCG_UE1_r16 = v0
	}

	// 2. pdcch-BlindDetectionCG-UE2-r16: integer
	{
		v1, err := d.DecodeInteger(pDCCHBlindDetectionCGUEMixedExtR16PdcchBlindDetectionCGUE2R16Constraints)
		if err != nil {
			return err
		}
		ie.Pdcch_BlindDetectionCG_UE2_r16 = v1
	}

	return nil
}
