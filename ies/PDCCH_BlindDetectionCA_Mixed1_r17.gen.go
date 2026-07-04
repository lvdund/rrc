// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PDCCH-BlindDetectionCA-Mixed1-r17 (line 18298).

var pDCCHBlindDetectionCAMixed1R17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pdcch-BlindDetectionCA1-r17", Optional: true},
		{Name: "pdcch-BlindDetectionCA2-r17", Optional: true},
		{Name: "pdcch-BlindDetectionCA3-r17", Optional: true},
	},
}

var pDCCHBlindDetectionCAMixed1R17PdcchBlindDetectionCA1R17Constraints = per.Constrained(1, 15)

var pDCCHBlindDetectionCAMixed1R17PdcchBlindDetectionCA2R17Constraints = per.Constrained(1, 15)

var pDCCHBlindDetectionCAMixed1R17PdcchBlindDetectionCA3R17Constraints = per.Constrained(1, 15)

type PDCCH_BlindDetectionCA_Mixed1_r17 struct {
	Pdcch_BlindDetectionCA1_r17 *int64
	Pdcch_BlindDetectionCA2_r17 *int64
	Pdcch_BlindDetectionCA3_r17 *int64
}

func (ie *PDCCH_BlindDetectionCA_Mixed1_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDCCHBlindDetectionCAMixed1R17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Pdcch_BlindDetectionCA1_r17 != nil, ie.Pdcch_BlindDetectionCA2_r17 != nil, ie.Pdcch_BlindDetectionCA3_r17 != nil}); err != nil {
		return err
	}

	// 2. pdcch-BlindDetectionCA1-r17: integer
	{
		if ie.Pdcch_BlindDetectionCA1_r17 != nil {
			if err := e.EncodeInteger(*ie.Pdcch_BlindDetectionCA1_r17, pDCCHBlindDetectionCAMixed1R17PdcchBlindDetectionCA1R17Constraints); err != nil {
				return err
			}
		}
	}

	// 3. pdcch-BlindDetectionCA2-r17: integer
	{
		if ie.Pdcch_BlindDetectionCA2_r17 != nil {
			if err := e.EncodeInteger(*ie.Pdcch_BlindDetectionCA2_r17, pDCCHBlindDetectionCAMixed1R17PdcchBlindDetectionCA2R17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. pdcch-BlindDetectionCA3-r17: integer
	{
		if ie.Pdcch_BlindDetectionCA3_r17 != nil {
			if err := e.EncodeInteger(*ie.Pdcch_BlindDetectionCA3_r17, pDCCHBlindDetectionCAMixed1R17PdcchBlindDetectionCA3R17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PDCCH_BlindDetectionCA_Mixed1_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDCCHBlindDetectionCAMixed1R17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. pdcch-BlindDetectionCA1-r17: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(pDCCHBlindDetectionCAMixed1R17PdcchBlindDetectionCA1R17Constraints)
			if err != nil {
				return err
			}
			ie.Pdcch_BlindDetectionCA1_r17 = &v
		}
	}

	// 3. pdcch-BlindDetectionCA2-r17: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(pDCCHBlindDetectionCAMixed1R17PdcchBlindDetectionCA2R17Constraints)
			if err != nil {
				return err
			}
			ie.Pdcch_BlindDetectionCA2_r17 = &v
		}
	}

	// 4. pdcch-BlindDetectionCA3-r17: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(pDCCHBlindDetectionCAMixed1R17PdcchBlindDetectionCA3R17Constraints)
			if err != nil {
				return err
			}
			ie.Pdcch_BlindDetectionCA3_r17 = &v
		}
	}

	return nil
}
