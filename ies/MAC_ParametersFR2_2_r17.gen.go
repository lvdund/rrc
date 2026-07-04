// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MAC-ParametersFR2-2-r17 (line 21043).

var mACParametersFR22R17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "directMCG-SCellActivation-r17", Optional: true},
		{Name: "directMCG-SCellActivationResume-r17", Optional: true},
		{Name: "directSCG-SCellActivation-r17", Optional: true},
		{Name: "directSCG-SCellActivationResume-r17", Optional: true},
		{Name: "drx-Adaptation-r17", Optional: true},
	},
}

const (
	MAC_ParametersFR2_2_r17_DirectMCG_SCellActivation_r17_Supported = 0
)

var mACParametersFR22R17DirectMCGSCellActivationR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersFR2_2_r17_DirectMCG_SCellActivation_r17_Supported},
}

const (
	MAC_ParametersFR2_2_r17_DirectMCG_SCellActivationResume_r17_Supported = 0
)

var mACParametersFR22R17DirectMCGSCellActivationResumeR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersFR2_2_r17_DirectMCG_SCellActivationResume_r17_Supported},
}

const (
	MAC_ParametersFR2_2_r17_DirectSCG_SCellActivation_r17_Supported = 0
)

var mACParametersFR22R17DirectSCGSCellActivationR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersFR2_2_r17_DirectSCG_SCellActivation_r17_Supported},
}

const (
	MAC_ParametersFR2_2_r17_DirectSCG_SCellActivationResume_r17_Supported = 0
)

var mACParametersFR22R17DirectSCGSCellActivationResumeR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersFR2_2_r17_DirectSCG_SCellActivationResume_r17_Supported},
}

var mACParametersFR22R17DrxAdaptationR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "non-SharedSpectrumChAccess-r17", Optional: true},
		{Name: "sharedSpectrumChAccess-r17", Optional: true},
	},
}

type MAC_ParametersFR2_2_r17 struct {
	DirectMCG_SCellActivation_r17       *int64
	DirectMCG_SCellActivationResume_r17 *int64
	DirectSCG_SCellActivation_r17       *int64
	DirectSCG_SCellActivationResume_r17 *int64
	Drx_Adaptation_r17                  *struct {
		Non_SharedSpectrumChAccess_r17 *MinTimeGapFR2_2_r17
		SharedSpectrumChAccess_r17     *MinTimeGapFR2_2_r17
	}
}

func (ie *MAC_ParametersFR2_2_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mACParametersFR22R17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.DirectMCG_SCellActivation_r17 != nil, ie.DirectMCG_SCellActivationResume_r17 != nil, ie.DirectSCG_SCellActivation_r17 != nil, ie.DirectSCG_SCellActivationResume_r17 != nil, ie.Drx_Adaptation_r17 != nil}); err != nil {
		return err
	}

	// 3. directMCG-SCellActivation-r17: enumerated
	{
		if ie.DirectMCG_SCellActivation_r17 != nil {
			if err := e.EncodeEnumerated(*ie.DirectMCG_SCellActivation_r17, mACParametersFR22R17DirectMCGSCellActivationR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. directMCG-SCellActivationResume-r17: enumerated
	{
		if ie.DirectMCG_SCellActivationResume_r17 != nil {
			if err := e.EncodeEnumerated(*ie.DirectMCG_SCellActivationResume_r17, mACParametersFR22R17DirectMCGSCellActivationResumeR17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. directSCG-SCellActivation-r17: enumerated
	{
		if ie.DirectSCG_SCellActivation_r17 != nil {
			if err := e.EncodeEnumerated(*ie.DirectSCG_SCellActivation_r17, mACParametersFR22R17DirectSCGSCellActivationR17Constraints); err != nil {
				return err
			}
		}
	}

	// 6. directSCG-SCellActivationResume-r17: enumerated
	{
		if ie.DirectSCG_SCellActivationResume_r17 != nil {
			if err := e.EncodeEnumerated(*ie.DirectSCG_SCellActivationResume_r17, mACParametersFR22R17DirectSCGSCellActivationResumeR17Constraints); err != nil {
				return err
			}
		}
	}

	// 7. drx-Adaptation-r17: sequence
	{
		if ie.Drx_Adaptation_r17 != nil {
			c := ie.Drx_Adaptation_r17
			mACParametersFR22R17DrxAdaptationR17Seq := e.NewSequenceEncoder(mACParametersFR22R17DrxAdaptationR17Constraints)
			if err := mACParametersFR22R17DrxAdaptationR17Seq.EncodePreamble([]bool{c.Non_SharedSpectrumChAccess_r17 != nil, c.SharedSpectrumChAccess_r17 != nil}); err != nil {
				return err
			}
			if c.Non_SharedSpectrumChAccess_r17 != nil {
				if err := c.Non_SharedSpectrumChAccess_r17.Encode(e); err != nil {
					return err
				}
			}
			if c.SharedSpectrumChAccess_r17 != nil {
				if err := c.SharedSpectrumChAccess_r17.Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *MAC_ParametersFR2_2_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mACParametersFR22R17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. directMCG-SCellActivation-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(mACParametersFR22R17DirectMCGSCellActivationR17Constraints)
			if err != nil {
				return err
			}
			ie.DirectMCG_SCellActivation_r17 = &idx
		}
	}

	// 4. directMCG-SCellActivationResume-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(mACParametersFR22R17DirectMCGSCellActivationResumeR17Constraints)
			if err != nil {
				return err
			}
			ie.DirectMCG_SCellActivationResume_r17 = &idx
		}
	}

	// 5. directSCG-SCellActivation-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(mACParametersFR22R17DirectSCGSCellActivationR17Constraints)
			if err != nil {
				return err
			}
			ie.DirectSCG_SCellActivation_r17 = &idx
		}
	}

	// 6. directSCG-SCellActivationResume-r17: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(mACParametersFR22R17DirectSCGSCellActivationResumeR17Constraints)
			if err != nil {
				return err
			}
			ie.DirectSCG_SCellActivationResume_r17 = &idx
		}
	}

	// 7. drx-Adaptation-r17: sequence
	{
		if seq.IsComponentPresent(4) {
			ie.Drx_Adaptation_r17 = &struct {
				Non_SharedSpectrumChAccess_r17 *MinTimeGapFR2_2_r17
				SharedSpectrumChAccess_r17     *MinTimeGapFR2_2_r17
			}{}
			c := ie.Drx_Adaptation_r17
			mACParametersFR22R17DrxAdaptationR17Seq := d.NewSequenceDecoder(mACParametersFR22R17DrxAdaptationR17Constraints)
			if err := mACParametersFR22R17DrxAdaptationR17Seq.DecodePreamble(); err != nil {
				return err
			}
			if mACParametersFR22R17DrxAdaptationR17Seq.IsComponentPresent(0) {
				c.Non_SharedSpectrumChAccess_r17 = new(MinTimeGapFR2_2_r17)
				if err := (*c.Non_SharedSpectrumChAccess_r17).Decode(d); err != nil {
					return err
				}
			}
			if mACParametersFR22R17DrxAdaptationR17Seq.IsComponentPresent(1) {
				c.SharedSpectrumChAccess_r17 = new(MinTimeGapFR2_2_r17)
				if err := (*c.SharedSpectrumChAccess_r17).Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
