// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PDCCH-BlindDetectionCA-MixedExt-r16 (line 18252).

var pDCCHBlindDetectionCAMixedExtR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pdcch-BlindDetectionCA1-r16"},
		{Name: "pdcch-BlindDetectionCA2-r16"},
	},
}

var pDCCHBlindDetectionCAMixedExtR16PdcchBlindDetectionCA1R16Constraints = per.Constrained(1, 15)

var pDCCHBlindDetectionCAMixedExtR16PdcchBlindDetectionCA2R16Constraints = per.Constrained(1, 15)

type PDCCH_BlindDetectionCA_MixedExt_r16 struct {
	Pdcch_BlindDetectionCA1_r16 int64
	Pdcch_BlindDetectionCA2_r16 int64
}

func (ie *PDCCH_BlindDetectionCA_MixedExt_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDCCHBlindDetectionCAMixedExtR16Constraints)
	_ = seq

	// 1. pdcch-BlindDetectionCA1-r16: integer
	{
		if err := e.EncodeInteger(ie.Pdcch_BlindDetectionCA1_r16, pDCCHBlindDetectionCAMixedExtR16PdcchBlindDetectionCA1R16Constraints); err != nil {
			return err
		}
	}

	// 2. pdcch-BlindDetectionCA2-r16: integer
	{
		if err := e.EncodeInteger(ie.Pdcch_BlindDetectionCA2_r16, pDCCHBlindDetectionCAMixedExtR16PdcchBlindDetectionCA2R16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PDCCH_BlindDetectionCA_MixedExt_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDCCHBlindDetectionCAMixedExtR16Constraints)
	_ = seq

	// 1. pdcch-BlindDetectionCA1-r16: integer
	{
		v0, err := d.DecodeInteger(pDCCHBlindDetectionCAMixedExtR16PdcchBlindDetectionCA1R16Constraints)
		if err != nil {
			return err
		}
		ie.Pdcch_BlindDetectionCA1_r16 = v0
	}

	// 2. pdcch-BlindDetectionCA2-r16: integer
	{
		v1, err := d.DecodeInteger(pDCCHBlindDetectionCAMixedExtR16PdcchBlindDetectionCA2R16Constraints)
		if err != nil {
			return err
		}
		ie.Pdcch_BlindDetectionCA2_r16 = v1
	}

	return nil
}
