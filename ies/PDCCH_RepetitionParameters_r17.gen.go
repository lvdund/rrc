// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PDCCH-RepetitionParameters-r17 (line 19803).

var pDCCHRepetitionParametersR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "supportedMode-r17"},
		{Name: "limitX-PerCC-r17", Optional: true},
		{Name: "limitX-AcrossCC-r17", Optional: true},
	},
}

const (
	PDCCH_RepetitionParameters_r17_SupportedMode_r17_Intra_Span = 0
	PDCCH_RepetitionParameters_r17_SupportedMode_r17_Inter_Span = 1
	PDCCH_RepetitionParameters_r17_SupportedMode_r17_Both       = 2
)

var pDCCHRepetitionParametersR17SupportedModeR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCCH_RepetitionParameters_r17_SupportedMode_r17_Intra_Span, PDCCH_RepetitionParameters_r17_SupportedMode_r17_Inter_Span, PDCCH_RepetitionParameters_r17_SupportedMode_r17_Both},
}

const (
	PDCCH_RepetitionParameters_r17_LimitX_PerCC_r17_N4      = 0
	PDCCH_RepetitionParameters_r17_LimitX_PerCC_r17_N8      = 1
	PDCCH_RepetitionParameters_r17_LimitX_PerCC_r17_N16     = 2
	PDCCH_RepetitionParameters_r17_LimitX_PerCC_r17_N32     = 3
	PDCCH_RepetitionParameters_r17_LimitX_PerCC_r17_N44     = 4
	PDCCH_RepetitionParameters_r17_LimitX_PerCC_r17_N64     = 5
	PDCCH_RepetitionParameters_r17_LimitX_PerCC_r17_Nolimit = 6
)

var pDCCHRepetitionParametersR17LimitXPerCCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCCH_RepetitionParameters_r17_LimitX_PerCC_r17_N4, PDCCH_RepetitionParameters_r17_LimitX_PerCC_r17_N8, PDCCH_RepetitionParameters_r17_LimitX_PerCC_r17_N16, PDCCH_RepetitionParameters_r17_LimitX_PerCC_r17_N32, PDCCH_RepetitionParameters_r17_LimitX_PerCC_r17_N44, PDCCH_RepetitionParameters_r17_LimitX_PerCC_r17_N64, PDCCH_RepetitionParameters_r17_LimitX_PerCC_r17_Nolimit},
}

const (
	PDCCH_RepetitionParameters_r17_LimitX_AcrossCC_r17_N4      = 0
	PDCCH_RepetitionParameters_r17_LimitX_AcrossCC_r17_N8      = 1
	PDCCH_RepetitionParameters_r17_LimitX_AcrossCC_r17_N16     = 2
	PDCCH_RepetitionParameters_r17_LimitX_AcrossCC_r17_N32     = 3
	PDCCH_RepetitionParameters_r17_LimitX_AcrossCC_r17_N44     = 4
	PDCCH_RepetitionParameters_r17_LimitX_AcrossCC_r17_N64     = 5
	PDCCH_RepetitionParameters_r17_LimitX_AcrossCC_r17_N128    = 6
	PDCCH_RepetitionParameters_r17_LimitX_AcrossCC_r17_N256    = 7
	PDCCH_RepetitionParameters_r17_LimitX_AcrossCC_r17_N512    = 8
	PDCCH_RepetitionParameters_r17_LimitX_AcrossCC_r17_Nolimit = 9
)

var pDCCHRepetitionParametersR17LimitXAcrossCCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCCH_RepetitionParameters_r17_LimitX_AcrossCC_r17_N4, PDCCH_RepetitionParameters_r17_LimitX_AcrossCC_r17_N8, PDCCH_RepetitionParameters_r17_LimitX_AcrossCC_r17_N16, PDCCH_RepetitionParameters_r17_LimitX_AcrossCC_r17_N32, PDCCH_RepetitionParameters_r17_LimitX_AcrossCC_r17_N44, PDCCH_RepetitionParameters_r17_LimitX_AcrossCC_r17_N64, PDCCH_RepetitionParameters_r17_LimitX_AcrossCC_r17_N128, PDCCH_RepetitionParameters_r17_LimitX_AcrossCC_r17_N256, PDCCH_RepetitionParameters_r17_LimitX_AcrossCC_r17_N512, PDCCH_RepetitionParameters_r17_LimitX_AcrossCC_r17_Nolimit},
}

type PDCCH_RepetitionParameters_r17 struct {
	SupportedMode_r17   int64
	LimitX_PerCC_r17    *int64
	LimitX_AcrossCC_r17 *int64
}

func (ie *PDCCH_RepetitionParameters_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDCCHRepetitionParametersR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.LimitX_PerCC_r17 != nil, ie.LimitX_AcrossCC_r17 != nil}); err != nil {
		return err
	}

	// 2. supportedMode-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.SupportedMode_r17, pDCCHRepetitionParametersR17SupportedModeR17Constraints); err != nil {
			return err
		}
	}

	// 3. limitX-PerCC-r17: enumerated
	{
		if ie.LimitX_PerCC_r17 != nil {
			if err := e.EncodeEnumerated(*ie.LimitX_PerCC_r17, pDCCHRepetitionParametersR17LimitXPerCCR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. limitX-AcrossCC-r17: enumerated
	{
		if ie.LimitX_AcrossCC_r17 != nil {
			if err := e.EncodeEnumerated(*ie.LimitX_AcrossCC_r17, pDCCHRepetitionParametersR17LimitXAcrossCCR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PDCCH_RepetitionParameters_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDCCHRepetitionParametersR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. supportedMode-r17: enumerated
	{
		v0, err := d.DecodeEnumerated(pDCCHRepetitionParametersR17SupportedModeR17Constraints)
		if err != nil {
			return err
		}
		ie.SupportedMode_r17 = v0
	}

	// 3. limitX-PerCC-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(pDCCHRepetitionParametersR17LimitXPerCCR17Constraints)
			if err != nil {
				return err
			}
			ie.LimitX_PerCC_r17 = &idx
		}
	}

	// 4. limitX-AcrossCC-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(pDCCHRepetitionParametersR17LimitXAcrossCCR17Constraints)
			if err != nil {
				return err
			}
			ie.LimitX_AcrossCC_r17 = &idx
		}
	}

	return nil
}
