// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-PSFCH-Config-r16 (line 28042).

var sLPSFCHConfigR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-PSFCH-Period-r16", Optional: true},
		{Name: "sl-PSFCH-RB-Set-r16", Optional: true},
		{Name: "sl-NumMuxCS-Pair-r16", Optional: true},
		{Name: "sl-MinTimeGapPSFCH-r16", Optional: true},
		{Name: "sl-PSFCH-HopID-r16", Optional: true},
		{Name: "sl-PSFCH-CandidateResourceType-r16", Optional: true},
	},
}

const (
	SL_PSFCH_Config_r16_Sl_PSFCH_Period_r16_Sl0 = 0
	SL_PSFCH_Config_r16_Sl_PSFCH_Period_r16_Sl1 = 1
	SL_PSFCH_Config_r16_Sl_PSFCH_Period_r16_Sl2 = 2
	SL_PSFCH_Config_r16_Sl_PSFCH_Period_r16_Sl4 = 3
)

var sLPSFCHConfigR16SlPSFCHPeriodR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_PSFCH_Config_r16_Sl_PSFCH_Period_r16_Sl0, SL_PSFCH_Config_r16_Sl_PSFCH_Period_r16_Sl1, SL_PSFCH_Config_r16_Sl_PSFCH_Period_r16_Sl2, SL_PSFCH_Config_r16_Sl_PSFCH_Period_r16_Sl4},
}

var sLPSFCHConfigR16SlPSFCHRBSetR16Constraints = per.SizeRange(10, 275)

const (
	SL_PSFCH_Config_r16_Sl_NumMuxCS_Pair_r16_N1 = 0
	SL_PSFCH_Config_r16_Sl_NumMuxCS_Pair_r16_N2 = 1
	SL_PSFCH_Config_r16_Sl_NumMuxCS_Pair_r16_N3 = 2
	SL_PSFCH_Config_r16_Sl_NumMuxCS_Pair_r16_N6 = 3
)

var sLPSFCHConfigR16SlNumMuxCSPairR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_PSFCH_Config_r16_Sl_NumMuxCS_Pair_r16_N1, SL_PSFCH_Config_r16_Sl_NumMuxCS_Pair_r16_N2, SL_PSFCH_Config_r16_Sl_NumMuxCS_Pair_r16_N3, SL_PSFCH_Config_r16_Sl_NumMuxCS_Pair_r16_N6},
}

const (
	SL_PSFCH_Config_r16_Sl_MinTimeGapPSFCH_r16_Sl2 = 0
	SL_PSFCH_Config_r16_Sl_MinTimeGapPSFCH_r16_Sl3 = 1
)

var sLPSFCHConfigR16SlMinTimeGapPSFCHR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_PSFCH_Config_r16_Sl_MinTimeGapPSFCH_r16_Sl2, SL_PSFCH_Config_r16_Sl_MinTimeGapPSFCH_r16_Sl3},
}

var sLPSFCHConfigR16SlPSFCHHopIDR16Constraints = per.Constrained(0, 1023)

const (
	SL_PSFCH_Config_r16_Sl_PSFCH_CandidateResourceType_r16_StartSubCH = 0
	SL_PSFCH_Config_r16_Sl_PSFCH_CandidateResourceType_r16_AllocSubCH = 1
)

var sLPSFCHConfigR16SlPSFCHCandidateResourceTypeR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_PSFCH_Config_r16_Sl_PSFCH_CandidateResourceType_r16_StartSubCH, SL_PSFCH_Config_r16_Sl_PSFCH_CandidateResourceType_r16_AllocSubCH},
}

type SL_PSFCH_Config_r16 struct {
	Sl_PSFCH_Period_r16                *int64
	Sl_PSFCH_RB_Set_r16                *per.BitString
	Sl_NumMuxCS_Pair_r16               *int64
	Sl_MinTimeGapPSFCH_r16             *int64
	Sl_PSFCH_HopID_r16                 *int64
	Sl_PSFCH_CandidateResourceType_r16 *int64
}

func (ie *SL_PSFCH_Config_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLPSFCHConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_PSFCH_Period_r16 != nil, ie.Sl_PSFCH_RB_Set_r16 != nil, ie.Sl_NumMuxCS_Pair_r16 != nil, ie.Sl_MinTimeGapPSFCH_r16 != nil, ie.Sl_PSFCH_HopID_r16 != nil, ie.Sl_PSFCH_CandidateResourceType_r16 != nil}); err != nil {
		return err
	}

	// 3. sl-PSFCH-Period-r16: enumerated
	{
		if ie.Sl_PSFCH_Period_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_PSFCH_Period_r16, sLPSFCHConfigR16SlPSFCHPeriodR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. sl-PSFCH-RB-Set-r16: bit-string
	{
		if ie.Sl_PSFCH_RB_Set_r16 != nil {
			if err := e.EncodeBitString(*ie.Sl_PSFCH_RB_Set_r16, sLPSFCHConfigR16SlPSFCHRBSetR16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. sl-NumMuxCS-Pair-r16: enumerated
	{
		if ie.Sl_NumMuxCS_Pair_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_NumMuxCS_Pair_r16, sLPSFCHConfigR16SlNumMuxCSPairR16Constraints); err != nil {
				return err
			}
		}
	}

	// 6. sl-MinTimeGapPSFCH-r16: enumerated
	{
		if ie.Sl_MinTimeGapPSFCH_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_MinTimeGapPSFCH_r16, sLPSFCHConfigR16SlMinTimeGapPSFCHR16Constraints); err != nil {
				return err
			}
		}
	}

	// 7. sl-PSFCH-HopID-r16: integer
	{
		if ie.Sl_PSFCH_HopID_r16 != nil {
			if err := e.EncodeInteger(*ie.Sl_PSFCH_HopID_r16, sLPSFCHConfigR16SlPSFCHHopIDR16Constraints); err != nil {
				return err
			}
		}
	}

	// 8. sl-PSFCH-CandidateResourceType-r16: enumerated
	{
		if ie.Sl_PSFCH_CandidateResourceType_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_PSFCH_CandidateResourceType_r16, sLPSFCHConfigR16SlPSFCHCandidateResourceTypeR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_PSFCH_Config_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLPSFCHConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-PSFCH-Period-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(sLPSFCHConfigR16SlPSFCHPeriodR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PSFCH_Period_r16 = &idx
		}
	}

	// 4. sl-PSFCH-RB-Set-r16: bit-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeBitString(sLPSFCHConfigR16SlPSFCHRBSetR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PSFCH_RB_Set_r16 = &v
		}
	}

	// 5. sl-NumMuxCS-Pair-r16: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(sLPSFCHConfigR16SlNumMuxCSPairR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_NumMuxCS_Pair_r16 = &idx
		}
	}

	// 6. sl-MinTimeGapPSFCH-r16: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(sLPSFCHConfigR16SlMinTimeGapPSFCHR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_MinTimeGapPSFCH_r16 = &idx
		}
	}

	// 7. sl-PSFCH-HopID-r16: integer
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeInteger(sLPSFCHConfigR16SlPSFCHHopIDR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PSFCH_HopID_r16 = &v
		}
	}

	// 8. sl-PSFCH-CandidateResourceType-r16: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(sLPSFCHConfigR16SlPSFCHCandidateResourceTypeR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PSFCH_CandidateResourceType_r16 = &idx
		}
	}

	return nil
}
