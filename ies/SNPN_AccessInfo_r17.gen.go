// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SNPN-AccessInfo-r17 (line 5514).

var sNPNAccessInfoR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "extCH-Supported-r17", Optional: true},
		{Name: "extCH-WithoutConfigAllowed-r17", Optional: true},
		{Name: "onboardingEnabled-r17", Optional: true},
		{Name: "imsEmergencySupportForSNPN-r17", Optional: true},
	},
}

const (
	SNPN_AccessInfo_r17_ExtCH_Supported_r17_True = 0
)

var sNPNAccessInfoR17ExtCHSupportedR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SNPN_AccessInfo_r17_ExtCH_Supported_r17_True},
}

const (
	SNPN_AccessInfo_r17_ExtCH_WithoutConfigAllowed_r17_True = 0
)

var sNPNAccessInfoR17ExtCHWithoutConfigAllowedR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SNPN_AccessInfo_r17_ExtCH_WithoutConfigAllowed_r17_True},
}

const (
	SNPN_AccessInfo_r17_OnboardingEnabled_r17_True = 0
)

var sNPNAccessInfoR17OnboardingEnabledR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SNPN_AccessInfo_r17_OnboardingEnabled_r17_True},
}

const (
	SNPN_AccessInfo_r17_ImsEmergencySupportForSNPN_r17_True = 0
)

var sNPNAccessInfoR17ImsEmergencySupportForSNPNR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SNPN_AccessInfo_r17_ImsEmergencySupportForSNPN_r17_True},
}

type SNPN_AccessInfo_r17 struct {
	ExtCH_Supported_r17            *int64
	ExtCH_WithoutConfigAllowed_r17 *int64
	OnboardingEnabled_r17          *int64
	ImsEmergencySupportForSNPN_r17 *int64
}

func (ie *SNPN_AccessInfo_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sNPNAccessInfoR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ExtCH_Supported_r17 != nil, ie.ExtCH_WithoutConfigAllowed_r17 != nil, ie.OnboardingEnabled_r17 != nil, ie.ImsEmergencySupportForSNPN_r17 != nil}); err != nil {
		return err
	}

	// 2. extCH-Supported-r17: enumerated
	{
		if ie.ExtCH_Supported_r17 != nil {
			if err := e.EncodeEnumerated(*ie.ExtCH_Supported_r17, sNPNAccessInfoR17ExtCHSupportedR17Constraints); err != nil {
				return err
			}
		}
	}

	// 3. extCH-WithoutConfigAllowed-r17: enumerated
	{
		if ie.ExtCH_WithoutConfigAllowed_r17 != nil {
			if err := e.EncodeEnumerated(*ie.ExtCH_WithoutConfigAllowed_r17, sNPNAccessInfoR17ExtCHWithoutConfigAllowedR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. onboardingEnabled-r17: enumerated
	{
		if ie.OnboardingEnabled_r17 != nil {
			if err := e.EncodeEnumerated(*ie.OnboardingEnabled_r17, sNPNAccessInfoR17OnboardingEnabledR17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. imsEmergencySupportForSNPN-r17: enumerated
	{
		if ie.ImsEmergencySupportForSNPN_r17 != nil {
			if err := e.EncodeEnumerated(*ie.ImsEmergencySupportForSNPN_r17, sNPNAccessInfoR17ImsEmergencySupportForSNPNR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SNPN_AccessInfo_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sNPNAccessInfoR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. extCH-Supported-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(sNPNAccessInfoR17ExtCHSupportedR17Constraints)
			if err != nil {
				return err
			}
			ie.ExtCH_Supported_r17 = &idx
		}
	}

	// 3. extCH-WithoutConfigAllowed-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(sNPNAccessInfoR17ExtCHWithoutConfigAllowedR17Constraints)
			if err != nil {
				return err
			}
			ie.ExtCH_WithoutConfigAllowed_r17 = &idx
		}
	}

	// 4. onboardingEnabled-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(sNPNAccessInfoR17OnboardingEnabledR17Constraints)
			if err != nil {
				return err
			}
			ie.OnboardingEnabled_r17 = &idx
		}
	}

	// 5. imsEmergencySupportForSNPN-r17: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(sNPNAccessInfoR17ImsEmergencySupportForSNPNR17Constraints)
			if err != nil {
				return err
			}
			ie.ImsEmergencySupportForSNPN_r17 = &idx
		}
	}

	return nil
}
