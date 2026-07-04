// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-PTRS-Config-r16 (line 28051).

var sLPTRSConfigR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-PTRS-FreqDensity-r16", Optional: true},
		{Name: "sl-PTRS-TimeDensity-r16", Optional: true},
		{Name: "sl-PTRS-RE-Offset-r16", Optional: true},
	},
}

var sLPTRSConfigR16SlPTRSFreqDensityR16Constraints = per.FixedSize(2)

var sLPTRSConfigR16SlPTRSTimeDensityR16Constraints = per.FixedSize(3)

const (
	SL_PTRS_Config_r16_Sl_PTRS_RE_Offset_r16_Offset01 = 0
	SL_PTRS_Config_r16_Sl_PTRS_RE_Offset_r16_Offset10 = 1
	SL_PTRS_Config_r16_Sl_PTRS_RE_Offset_r16_Offset11 = 2
)

var sLPTRSConfigR16SlPTRSREOffsetR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_PTRS_Config_r16_Sl_PTRS_RE_Offset_r16_Offset01, SL_PTRS_Config_r16_Sl_PTRS_RE_Offset_r16_Offset10, SL_PTRS_Config_r16_Sl_PTRS_RE_Offset_r16_Offset11},
}

type SL_PTRS_Config_r16 struct {
	Sl_PTRS_FreqDensity_r16 []int64
	Sl_PTRS_TimeDensity_r16 []int64
	Sl_PTRS_RE_Offset_r16   *int64
}

func (ie *SL_PTRS_Config_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLPTRSConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_PTRS_FreqDensity_r16 != nil, ie.Sl_PTRS_TimeDensity_r16 != nil, ie.Sl_PTRS_RE_Offset_r16 != nil}); err != nil {
		return err
	}

	// 3. sl-PTRS-FreqDensity-r16: sequence-of
	{
		if ie.Sl_PTRS_FreqDensity_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(sLPTRSConfigR16SlPTRSFreqDensityR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_PTRS_FreqDensity_r16))); err != nil {
				return err
			}
			for i := range ie.Sl_PTRS_FreqDensity_r16 {
				if err := e.EncodeInteger(ie.Sl_PTRS_FreqDensity_r16[i], per.Constrained(1, 276)); err != nil {
					return err
				}
			}
		}
	}

	// 4. sl-PTRS-TimeDensity-r16: sequence-of
	{
		if ie.Sl_PTRS_TimeDensity_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(sLPTRSConfigR16SlPTRSTimeDensityR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_PTRS_TimeDensity_r16))); err != nil {
				return err
			}
			for i := range ie.Sl_PTRS_TimeDensity_r16 {
				if err := e.EncodeInteger(ie.Sl_PTRS_TimeDensity_r16[i], per.Constrained(0, 29)); err != nil {
					return err
				}
			}
		}
	}

	// 5. sl-PTRS-RE-Offset-r16: enumerated
	{
		if ie.Sl_PTRS_RE_Offset_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_PTRS_RE_Offset_r16, sLPTRSConfigR16SlPTRSREOffsetR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_PTRS_Config_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLPTRSConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-PTRS-FreqDensity-r16: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(sLPTRSConfigR16SlPTRSFreqDensityR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_PTRS_FreqDensity_r16 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(1, 276))
				if err != nil {
					return err
				}
				ie.Sl_PTRS_FreqDensity_r16[i] = v
			}
		}
	}

	// 4. sl-PTRS-TimeDensity-r16: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(sLPTRSConfigR16SlPTRSTimeDensityR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_PTRS_TimeDensity_r16 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, 29))
				if err != nil {
					return err
				}
				ie.Sl_PTRS_TimeDensity_r16[i] = v
			}
		}
	}

	// 5. sl-PTRS-RE-Offset-r16: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(sLPTRSConfigR16SlPTRSREOffsetR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PTRS_RE_Offset_r16 = &idx
		}
	}

	return nil
}
