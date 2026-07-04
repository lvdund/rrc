// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCSetupComplete-v1700-IEs (line 1744).

var rRCSetupCompleteV1700IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "onboardingRequest-r17", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	RRCSetupComplete_v1700_IEs_OnboardingRequest_r17_True = 0
)

var rRCSetupCompleteV1700IEsOnboardingRequestR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCSetupComplete_v1700_IEs_OnboardingRequest_r17_True},
}

type RRCSetupComplete_v1700_IEs struct {
	OnboardingRequest_r17 *int64
	NonCriticalExtension  *RRCSetupComplete_v1800_IEs
}

func (ie *RRCSetupComplete_v1700_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCSetupCompleteV1700IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.OnboardingRequest_r17 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. onboardingRequest-r17: enumerated
	{
		if ie.OnboardingRequest_r17 != nil {
			if err := e.EncodeEnumerated(*ie.OnboardingRequest_r17, rRCSetupCompleteV1700IEsOnboardingRequestR17Constraints); err != nil {
				return err
			}
		}
	}

	// 3. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *RRCSetupComplete_v1700_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCSetupCompleteV1700IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. onboardingRequest-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(rRCSetupCompleteV1700IEsOnboardingRequestR17Constraints)
			if err != nil {
				return err
			}
			ie.OnboardingRequest_r17 = &idx
		}
	}

	// 3. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = new(RRCSetupComplete_v1800_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
