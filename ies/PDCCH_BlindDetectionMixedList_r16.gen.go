// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PDCCH-BlindDetectionMixedList-r16 (line 18241).

var pDCCHBlindDetectionMixedListR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pdcch-BlindDetectionCA-MixedExt-r16", Optional: true},
		{Name: "pdcch-BlindDetectionCG-UE-MixedExt-r16", Optional: true},
	},
}

var pDCCH_BlindDetectionMixedList_r16PdcchBlindDetectionCAMixedExtR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "pdcch-BlindDetectionCA-Mixed-v16a0"},
		{Name: "pdcch-BlindDetectionCA-Mixed-NonAlignedSpan-v16a0"},
	},
}

const (
	PDCCH_BlindDetectionMixedList_r16_Pdcch_BlindDetectionCA_MixedExt_r16_Pdcch_BlindDetectionCA_Mixed_V16a0                = 0
	PDCCH_BlindDetectionMixedList_r16_Pdcch_BlindDetectionCA_MixedExt_r16_Pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_V16a0 = 1
)

type PDCCH_BlindDetectionMixedList_r16 struct {
	Pdcch_BlindDetectionCA_MixedExt_r16 *struct {
		Choice                                            int
		Pdcch_BlindDetectionCA_Mixed_V16a0                *PDCCH_BlindDetectionCA_MixedExt_r16
		Pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_V16a0 *PDCCH_BlindDetectionCA_MixedExt_r16
	}
	Pdcch_BlindDetectionCG_UE_MixedExt_r16 *struct {
		Pdcch_BlindDetectionMCG_UE_Mixed_V16a0 PDCCH_BlindDetectionCG_UE_MixedExt_r16
		Pdcch_BlindDetectionSCG_UE_Mixed_V16a0 PDCCH_BlindDetectionCG_UE_MixedExt_r16
	}
}

func (ie *PDCCH_BlindDetectionMixedList_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDCCHBlindDetectionMixedListR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Pdcch_BlindDetectionCA_MixedExt_r16 != nil, ie.Pdcch_BlindDetectionCG_UE_MixedExt_r16 != nil}); err != nil {
		return err
	}

	// 2. pdcch-BlindDetectionCA-MixedExt-r16: choice
	{
		if ie.Pdcch_BlindDetectionCA_MixedExt_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(pDCCH_BlindDetectionMixedList_r16PdcchBlindDetectionCAMixedExtR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Pdcch_BlindDetectionCA_MixedExt_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Pdcch_BlindDetectionCA_MixedExt_r16).Choice {
			case PDCCH_BlindDetectionMixedList_r16_Pdcch_BlindDetectionCA_MixedExt_r16_Pdcch_BlindDetectionCA_Mixed_V16a0:
				if err := (*ie.Pdcch_BlindDetectionCA_MixedExt_r16).Pdcch_BlindDetectionCA_Mixed_V16a0.Encode(e); err != nil {
					return err
				}
			case PDCCH_BlindDetectionMixedList_r16_Pdcch_BlindDetectionCA_MixedExt_r16_Pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_V16a0:
				if err := (*ie.Pdcch_BlindDetectionCA_MixedExt_r16).Pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_V16a0.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Pdcch_BlindDetectionCA_MixedExt_r16).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 3. pdcch-BlindDetectionCG-UE-MixedExt-r16: sequence
	{
		if ie.Pdcch_BlindDetectionCG_UE_MixedExt_r16 != nil {
			c := ie.Pdcch_BlindDetectionCG_UE_MixedExt_r16
			if err := c.Pdcch_BlindDetectionMCG_UE_Mixed_V16a0.Encode(e); err != nil {
				return err
			}
			if err := c.Pdcch_BlindDetectionSCG_UE_Mixed_V16a0.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PDCCH_BlindDetectionMixedList_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDCCHBlindDetectionMixedListR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. pdcch-BlindDetectionCA-MixedExt-r16: choice
	{
		if seq.IsComponentPresent(0) {
			ie.Pdcch_BlindDetectionCA_MixedExt_r16 = &struct {
				Choice                                            int
				Pdcch_BlindDetectionCA_Mixed_V16a0                *PDCCH_BlindDetectionCA_MixedExt_r16
				Pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_V16a0 *PDCCH_BlindDetectionCA_MixedExt_r16
			}{}
			choiceDec := d.NewChoiceDecoder(pDCCH_BlindDetectionMixedList_r16PdcchBlindDetectionCAMixedExtR16Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Pdcch_BlindDetectionCA_MixedExt_r16).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case PDCCH_BlindDetectionMixedList_r16_Pdcch_BlindDetectionCA_MixedExt_r16_Pdcch_BlindDetectionCA_Mixed_V16a0:
				(*ie.Pdcch_BlindDetectionCA_MixedExt_r16).Pdcch_BlindDetectionCA_Mixed_V16a0 = new(PDCCH_BlindDetectionCA_MixedExt_r16)
				if err := (*ie.Pdcch_BlindDetectionCA_MixedExt_r16).Pdcch_BlindDetectionCA_Mixed_V16a0.Decode(d); err != nil {
					return err
				}
			case PDCCH_BlindDetectionMixedList_r16_Pdcch_BlindDetectionCA_MixedExt_r16_Pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_V16a0:
				(*ie.Pdcch_BlindDetectionCA_MixedExt_r16).Pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_V16a0 = new(PDCCH_BlindDetectionCA_MixedExt_r16)
				if err := (*ie.Pdcch_BlindDetectionCA_MixedExt_r16).Pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_V16a0.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. pdcch-BlindDetectionCG-UE-MixedExt-r16: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.Pdcch_BlindDetectionCG_UE_MixedExt_r16 = &struct {
				Pdcch_BlindDetectionMCG_UE_Mixed_V16a0 PDCCH_BlindDetectionCG_UE_MixedExt_r16
				Pdcch_BlindDetectionSCG_UE_Mixed_V16a0 PDCCH_BlindDetectionCG_UE_MixedExt_r16
			}{}
			c := ie.Pdcch_BlindDetectionCG_UE_MixedExt_r16
			{
				if err := c.Pdcch_BlindDetectionMCG_UE_Mixed_V16a0.Decode(d); err != nil {
					return err
				}
			}
			{
				if err := c.Pdcch_BlindDetectionSCG_UE_Mixed_V16a0.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
