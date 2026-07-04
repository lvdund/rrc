// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MAC-ParametersFRX-Diff-r16 (line 21030).

var mACParametersFRXDiffR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "directMCG-SCellActivation-r16", Optional: true},
		{Name: "directMCG-SCellActivationResume-r16", Optional: true},
		{Name: "directSCG-SCellActivation-r16", Optional: true},
		{Name: "directSCG-SCellActivationResume-r16", Optional: true},
		{Name: "drx-Adaptation-r16", Optional: true},
	},
}

const (
	MAC_ParametersFRX_Diff_r16_DirectMCG_SCellActivation_r16_Supported = 0
)

var mACParametersFRXDiffR16DirectMCGSCellActivationR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersFRX_Diff_r16_DirectMCG_SCellActivation_r16_Supported},
}

const (
	MAC_ParametersFRX_Diff_r16_DirectMCG_SCellActivationResume_r16_Supported = 0
)

var mACParametersFRXDiffR16DirectMCGSCellActivationResumeR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersFRX_Diff_r16_DirectMCG_SCellActivationResume_r16_Supported},
}

const (
	MAC_ParametersFRX_Diff_r16_DirectSCG_SCellActivation_r16_Supported = 0
)

var mACParametersFRXDiffR16DirectSCGSCellActivationR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersFRX_Diff_r16_DirectSCG_SCellActivation_r16_Supported},
}

const (
	MAC_ParametersFRX_Diff_r16_DirectSCG_SCellActivationResume_r16_Supported = 0
)

var mACParametersFRXDiffR16DirectSCGSCellActivationResumeR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersFRX_Diff_r16_DirectSCG_SCellActivationResume_r16_Supported},
}

var mACParametersFRXDiffR16DrxAdaptationR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "non-SharedSpectrumChAccess-r16", Optional: true},
		{Name: "sharedSpectrumChAccess-r16", Optional: true},
	},
}

type MAC_ParametersFRX_Diff_r16 struct {
	DirectMCG_SCellActivation_r16       *int64
	DirectMCG_SCellActivationResume_r16 *int64
	DirectSCG_SCellActivation_r16       *int64
	DirectSCG_SCellActivationResume_r16 *int64
	Drx_Adaptation_r16                  *struct {
		Non_SharedSpectrumChAccess_r16 *MinTimeGap_r16
		SharedSpectrumChAccess_r16     *MinTimeGap_r16
	}
}

func (ie *MAC_ParametersFRX_Diff_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mACParametersFRXDiffR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.DirectMCG_SCellActivation_r16 != nil, ie.DirectMCG_SCellActivationResume_r16 != nil, ie.DirectSCG_SCellActivation_r16 != nil, ie.DirectSCG_SCellActivationResume_r16 != nil, ie.Drx_Adaptation_r16 != nil}); err != nil {
		return err
	}

	// 3. directMCG-SCellActivation-r16: enumerated
	{
		if ie.DirectMCG_SCellActivation_r16 != nil {
			if err := e.EncodeEnumerated(*ie.DirectMCG_SCellActivation_r16, mACParametersFRXDiffR16DirectMCGSCellActivationR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. directMCG-SCellActivationResume-r16: enumerated
	{
		if ie.DirectMCG_SCellActivationResume_r16 != nil {
			if err := e.EncodeEnumerated(*ie.DirectMCG_SCellActivationResume_r16, mACParametersFRXDiffR16DirectMCGSCellActivationResumeR16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. directSCG-SCellActivation-r16: enumerated
	{
		if ie.DirectSCG_SCellActivation_r16 != nil {
			if err := e.EncodeEnumerated(*ie.DirectSCG_SCellActivation_r16, mACParametersFRXDiffR16DirectSCGSCellActivationR16Constraints); err != nil {
				return err
			}
		}
	}

	// 6. directSCG-SCellActivationResume-r16: enumerated
	{
		if ie.DirectSCG_SCellActivationResume_r16 != nil {
			if err := e.EncodeEnumerated(*ie.DirectSCG_SCellActivationResume_r16, mACParametersFRXDiffR16DirectSCGSCellActivationResumeR16Constraints); err != nil {
				return err
			}
		}
	}

	// 7. drx-Adaptation-r16: sequence
	{
		if ie.Drx_Adaptation_r16 != nil {
			c := ie.Drx_Adaptation_r16
			mACParametersFRXDiffR16DrxAdaptationR16Seq := e.NewSequenceEncoder(mACParametersFRXDiffR16DrxAdaptationR16Constraints)
			if err := mACParametersFRXDiffR16DrxAdaptationR16Seq.EncodePreamble([]bool{c.Non_SharedSpectrumChAccess_r16 != nil, c.SharedSpectrumChAccess_r16 != nil}); err != nil {
				return err
			}
			if c.Non_SharedSpectrumChAccess_r16 != nil {
				if err := c.Non_SharedSpectrumChAccess_r16.Encode(e); err != nil {
					return err
				}
			}
			if c.SharedSpectrumChAccess_r16 != nil {
				if err := c.SharedSpectrumChAccess_r16.Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *MAC_ParametersFRX_Diff_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mACParametersFRXDiffR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. directMCG-SCellActivation-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(mACParametersFRXDiffR16DirectMCGSCellActivationR16Constraints)
			if err != nil {
				return err
			}
			ie.DirectMCG_SCellActivation_r16 = &idx
		}
	}

	// 4. directMCG-SCellActivationResume-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(mACParametersFRXDiffR16DirectMCGSCellActivationResumeR16Constraints)
			if err != nil {
				return err
			}
			ie.DirectMCG_SCellActivationResume_r16 = &idx
		}
	}

	// 5. directSCG-SCellActivation-r16: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(mACParametersFRXDiffR16DirectSCGSCellActivationR16Constraints)
			if err != nil {
				return err
			}
			ie.DirectSCG_SCellActivation_r16 = &idx
		}
	}

	// 6. directSCG-SCellActivationResume-r16: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(mACParametersFRXDiffR16DirectSCGSCellActivationResumeR16Constraints)
			if err != nil {
				return err
			}
			ie.DirectSCG_SCellActivationResume_r16 = &idx
		}
	}

	// 7. drx-Adaptation-r16: sequence
	{
		if seq.IsComponentPresent(4) {
			ie.Drx_Adaptation_r16 = &struct {
				Non_SharedSpectrumChAccess_r16 *MinTimeGap_r16
				SharedSpectrumChAccess_r16     *MinTimeGap_r16
			}{}
			c := ie.Drx_Adaptation_r16
			mACParametersFRXDiffR16DrxAdaptationR16Seq := d.NewSequenceDecoder(mACParametersFRXDiffR16DrxAdaptationR16Constraints)
			if err := mACParametersFRXDiffR16DrxAdaptationR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if mACParametersFRXDiffR16DrxAdaptationR16Seq.IsComponentPresent(0) {
				c.Non_SharedSpectrumChAccess_r16 = new(MinTimeGap_r16)
				if err := (*c.Non_SharedSpectrumChAccess_r16).Decode(d); err != nil {
					return err
				}
			}
			if mACParametersFRXDiffR16DrxAdaptationR16Seq.IsComponentPresent(1) {
				c.SharedSpectrumChAccess_r16 = new(MinTimeGap_r16)
				if err := (*c.SharedSpectrumChAccess_r16).Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
