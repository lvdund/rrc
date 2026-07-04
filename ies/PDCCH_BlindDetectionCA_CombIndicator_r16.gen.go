// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PDCCH-BlindDetectionCA-CombIndicator-r16 (line 11740).

var pDCCHBlindDetectionCACombIndicatorR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pdcch-BlindDetectionCA1-r16"},
		{Name: "pdcch-BlindDetectionCA2-r16"},
	},
}

var pDCCHBlindDetectionCACombIndicatorR16PdcchBlindDetectionCA1R16Constraints = per.Constrained(1, 15)

var pDCCHBlindDetectionCACombIndicatorR16PdcchBlindDetectionCA2R16Constraints = per.Constrained(1, 15)

type PDCCH_BlindDetectionCA_CombIndicator_r16 struct {
	Pdcch_BlindDetectionCA1_r16 int64
	Pdcch_BlindDetectionCA2_r16 int64
}

func (ie *PDCCH_BlindDetectionCA_CombIndicator_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDCCHBlindDetectionCACombIndicatorR16Constraints)
	_ = seq

	// 1. pdcch-BlindDetectionCA1-r16: integer
	{
		if err := e.EncodeInteger(ie.Pdcch_BlindDetectionCA1_r16, pDCCHBlindDetectionCACombIndicatorR16PdcchBlindDetectionCA1R16Constraints); err != nil {
			return err
		}
	}

	// 2. pdcch-BlindDetectionCA2-r16: integer
	{
		if err := e.EncodeInteger(ie.Pdcch_BlindDetectionCA2_r16, pDCCHBlindDetectionCACombIndicatorR16PdcchBlindDetectionCA2R16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PDCCH_BlindDetectionCA_CombIndicator_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDCCHBlindDetectionCACombIndicatorR16Constraints)
	_ = seq

	// 1. pdcch-BlindDetectionCA1-r16: integer
	{
		v0, err := d.DecodeInteger(pDCCHBlindDetectionCACombIndicatorR16PdcchBlindDetectionCA1R16Constraints)
		if err != nil {
			return err
		}
		ie.Pdcch_BlindDetectionCA1_r16 = v0
	}

	// 2. pdcch-BlindDetectionCA2-r16: integer
	{
		v1, err := d.DecodeInteger(pDCCHBlindDetectionCACombIndicatorR16PdcchBlindDetectionCA2R16Constraints)
		if err != nil {
			return err
		}
		ie.Pdcch_BlindDetectionCA2_r16 = v1
	}

	return nil
}
