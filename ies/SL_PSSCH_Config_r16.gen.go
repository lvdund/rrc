// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-PSSCH-Config-r16 (line 28035).

var sLPSSCHConfigR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-PSSCH-DMRS-TimePatternList-r16", Optional: true},
		{Name: "sl-BetaOffsets2ndSCI-r16", Optional: true},
		{Name: "sl-Scaling-r16", Optional: true},
	},
}

var sLPSSCHConfigR16SlPSSCHDMRSTimePatternListR16Constraints = per.SizeRange(1, 3)

var sLPSSCHConfigR16SlBetaOffsets2ndSCIR16Constraints = per.FixedSize(4)

const (
	SL_PSSCH_Config_r16_Sl_Scaling_r16_F0p5  = 0
	SL_PSSCH_Config_r16_Sl_Scaling_r16_F0p65 = 1
	SL_PSSCH_Config_r16_Sl_Scaling_r16_F0p8  = 2
	SL_PSSCH_Config_r16_Sl_Scaling_r16_F1    = 3
)

var sLPSSCHConfigR16SlScalingR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_PSSCH_Config_r16_Sl_Scaling_r16_F0p5, SL_PSSCH_Config_r16_Sl_Scaling_r16_F0p65, SL_PSSCH_Config_r16_Sl_Scaling_r16_F0p8, SL_PSSCH_Config_r16_Sl_Scaling_r16_F1},
}

type SL_PSSCH_Config_r16 struct {
	Sl_PSSCH_DMRS_TimePatternList_r16 []int64
	Sl_BetaOffsets2ndSCI_r16          []SL_BetaOffsets_r16
	Sl_Scaling_r16                    *int64
}

func (ie *SL_PSSCH_Config_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLPSSCHConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_PSSCH_DMRS_TimePatternList_r16 != nil, ie.Sl_BetaOffsets2ndSCI_r16 != nil, ie.Sl_Scaling_r16 != nil}); err != nil {
		return err
	}

	// 3. sl-PSSCH-DMRS-TimePatternList-r16: sequence-of
	{
		if ie.Sl_PSSCH_DMRS_TimePatternList_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(sLPSSCHConfigR16SlPSSCHDMRSTimePatternListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_PSSCH_DMRS_TimePatternList_r16))); err != nil {
				return err
			}
			for i := range ie.Sl_PSSCH_DMRS_TimePatternList_r16 {
				if err := e.EncodeInteger(ie.Sl_PSSCH_DMRS_TimePatternList_r16[i], per.Constrained(2, 4)); err != nil {
					return err
				}
			}
		}
	}

	// 4. sl-BetaOffsets2ndSCI-r16: sequence-of
	{
		if ie.Sl_BetaOffsets2ndSCI_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(sLPSSCHConfigR16SlBetaOffsets2ndSCIR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_BetaOffsets2ndSCI_r16))); err != nil {
				return err
			}
			for i := range ie.Sl_BetaOffsets2ndSCI_r16 {
				if err := ie.Sl_BetaOffsets2ndSCI_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. sl-Scaling-r16: enumerated
	{
		if ie.Sl_Scaling_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_Scaling_r16, sLPSSCHConfigR16SlScalingR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_PSSCH_Config_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLPSSCHConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-PSSCH-DMRS-TimePatternList-r16: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(sLPSSCHConfigR16SlPSSCHDMRSTimePatternListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_PSSCH_DMRS_TimePatternList_r16 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(2, 4))
				if err != nil {
					return err
				}
				ie.Sl_PSSCH_DMRS_TimePatternList_r16[i] = v
			}
		}
	}

	// 4. sl-BetaOffsets2ndSCI-r16: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(sLPSSCHConfigR16SlBetaOffsets2ndSCIR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_BetaOffsets2ndSCI_r16 = make([]SL_BetaOffsets_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_BetaOffsets2ndSCI_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. sl-Scaling-r16: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(sLPSSCHConfigR16SlScalingR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_Scaling_r16 = &idx
		}
	}

	return nil
}
