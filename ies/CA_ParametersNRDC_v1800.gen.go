// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CA-ParametersNRDC-v1800 (line 18448).

var cAParametersNRDCV1800Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ca-ParametersNR-ForDC-v1800", Optional: true},
		{Name: "pdcch-BlindDetectionNRDC-r18", Optional: true},
	},
}

var cAParametersNRDCV1800PdcchBlindDetectionNRDCR18Constraints = per.SizeRange(1, common.MaxNrofPdcch_BlindDetectionMixed_1_r16)

type CA_ParametersNRDC_v1800 struct {
	Ca_ParametersNR_ForDC_v1800  *CA_ParametersNR_v1800
	Pdcch_BlindDetectionNRDC_r18 []PDCCH_BlindDetectionMixed1_r18
}

func (ie *CA_ParametersNRDC_v1800) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cAParametersNRDCV1800Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ca_ParametersNR_ForDC_v1800 != nil, ie.Pdcch_BlindDetectionNRDC_r18 != nil}); err != nil {
		return err
	}

	// 2. ca-ParametersNR-ForDC-v1800: ref
	{
		if ie.Ca_ParametersNR_ForDC_v1800 != nil {
			if err := ie.Ca_ParametersNR_ForDC_v1800.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. pdcch-BlindDetectionNRDC-r18: sequence-of
	{
		if ie.Pdcch_BlindDetectionNRDC_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(cAParametersNRDCV1800PdcchBlindDetectionNRDCR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Pdcch_BlindDetectionNRDC_r18))); err != nil {
				return err
			}
			for i := range ie.Pdcch_BlindDetectionNRDC_r18 {
				if err := ie.Pdcch_BlindDetectionNRDC_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *CA_ParametersNRDC_v1800) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cAParametersNRDCV1800Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ca-ParametersNR-ForDC-v1800: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Ca_ParametersNR_ForDC_v1800 = new(CA_ParametersNR_v1800)
			if err := ie.Ca_ParametersNR_ForDC_v1800.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. pdcch-BlindDetectionNRDC-r18: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(cAParametersNRDCV1800PdcchBlindDetectionNRDCR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Pdcch_BlindDetectionNRDC_r18 = make([]PDCCH_BlindDetectionMixed1_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Pdcch_BlindDetectionNRDC_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
