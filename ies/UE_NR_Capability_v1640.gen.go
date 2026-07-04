// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-NR-Capability-v1640 (line 25745).

var uENRCapabilityV1640Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "redirectAtResumeByNAS-r16", Optional: true},
		{Name: "phy-ParametersSharedSpectrumChAccess-r16", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	UE_NR_Capability_v1640_RedirectAtResumeByNAS_r16_Supported = 0
)

var uENRCapabilityV1640RedirectAtResumeByNASR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1640_RedirectAtResumeByNAS_r16_Supported},
}

type UE_NR_Capability_v1640 struct {
	RedirectAtResumeByNAS_r16                *int64
	Phy_ParametersSharedSpectrumChAccess_r16 *Phy_ParametersSharedSpectrumChAccess_r16
	NonCriticalExtension                     *UE_NR_Capability_v1650
}

func (ie *UE_NR_Capability_v1640) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uENRCapabilityV1640Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.RedirectAtResumeByNAS_r16 != nil, ie.Phy_ParametersSharedSpectrumChAccess_r16 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. redirectAtResumeByNAS-r16: enumerated
	{
		if ie.RedirectAtResumeByNAS_r16 != nil {
			if err := e.EncodeEnumerated(*ie.RedirectAtResumeByNAS_r16, uENRCapabilityV1640RedirectAtResumeByNASR16Constraints); err != nil {
				return err
			}
		}
	}

	// 3. phy-ParametersSharedSpectrumChAccess-r16: ref
	{
		if ie.Phy_ParametersSharedSpectrumChAccess_r16 != nil {
			if err := ie.Phy_ParametersSharedSpectrumChAccess_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UE_NR_Capability_v1640) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uENRCapabilityV1640Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. redirectAtResumeByNAS-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1640RedirectAtResumeByNASR16Constraints)
			if err != nil {
				return err
			}
			ie.RedirectAtResumeByNAS_r16 = &idx
		}
	}

	// 3. phy-ParametersSharedSpectrumChAccess-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Phy_ParametersSharedSpectrumChAccess_r16 = new(Phy_ParametersSharedSpectrumChAccess_r16)
			if err := ie.Phy_ParametersSharedSpectrumChAccess_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(2) {
			ie.NonCriticalExtension = new(UE_NR_Capability_v1650)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
