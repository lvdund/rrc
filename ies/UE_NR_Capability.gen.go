// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: UE-NR-Capability (line 25641).

var uENRCapabilityConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "accessStratumRelease"},
		{Name: "pdcp-Parameters"},
		{Name: "rlc-Parameters", Optional: true},
		{Name: "mac-Parameters", Optional: true},
		{Name: "phy-Parameters"},
		{Name: "rf-Parameters"},
		{Name: "measAndMobParameters", Optional: true},
		{Name: "fdd-Add-UE-NR-Capabilities", Optional: true},
		{Name: "tdd-Add-UE-NR-Capabilities", Optional: true},
		{Name: "fr1-Add-UE-NR-Capabilities", Optional: true},
		{Name: "fr2-Add-UE-NR-Capabilities", Optional: true},
		{Name: "featureSets", Optional: true},
		{Name: "featureSetCombinations", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var uENRCapabilityFeatureSetCombinationsConstraints = per.SizeRange(1, common.MaxFeatureSetCombinations)

var uENRCapabilityLateNonCriticalExtensionConstraints = per.SizeConstraints{}

type UE_NR_Capability struct {
	AccessStratumRelease       AccessStratumRelease
	Pdcp_Parameters            PDCP_Parameters
	Rlc_Parameters             *RLC_Parameters
	Mac_Parameters             *MAC_Parameters
	Phy_Parameters             Phy_Parameters
	Rf_Parameters              RF_Parameters
	MeasAndMobParameters       *MeasAndMobParameters
	Fdd_Add_UE_NR_Capabilities *UE_NR_CapabilityAddXDD_Mode
	Tdd_Add_UE_NR_Capabilities *UE_NR_CapabilityAddXDD_Mode
	Fr1_Add_UE_NR_Capabilities *UE_NR_CapabilityAddFRX_Mode
	Fr2_Add_UE_NR_Capabilities *UE_NR_CapabilityAddFRX_Mode
	FeatureSets                *FeatureSets
	FeatureSetCombinations     []FeatureSetCombination
	LateNonCriticalExtension   []byte
	NonCriticalExtension       *UE_NR_Capability_v1530
}

func (ie *UE_NR_Capability) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uENRCapabilityConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Rlc_Parameters != nil, ie.Mac_Parameters != nil, ie.MeasAndMobParameters != nil, ie.Fdd_Add_UE_NR_Capabilities != nil, ie.Tdd_Add_UE_NR_Capabilities != nil, ie.Fr1_Add_UE_NR_Capabilities != nil, ie.Fr2_Add_UE_NR_Capabilities != nil, ie.FeatureSets != nil, ie.FeatureSetCombinations != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. accessStratumRelease: ref
	{
		if err := ie.AccessStratumRelease.Encode(e); err != nil {
			return err
		}
	}

	// 3. pdcp-Parameters: ref
	{
		if err := ie.Pdcp_Parameters.Encode(e); err != nil {
			return err
		}
	}

	// 4. rlc-Parameters: ref
	{
		if ie.Rlc_Parameters != nil {
			if err := ie.Rlc_Parameters.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. mac-Parameters: ref
	{
		if ie.Mac_Parameters != nil {
			if err := ie.Mac_Parameters.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. phy-Parameters: ref
	{
		if err := ie.Phy_Parameters.Encode(e); err != nil {
			return err
		}
	}

	// 7. rf-Parameters: ref
	{
		if err := ie.Rf_Parameters.Encode(e); err != nil {
			return err
		}
	}

	// 8. measAndMobParameters: ref
	{
		if ie.MeasAndMobParameters != nil {
			if err := ie.MeasAndMobParameters.Encode(e); err != nil {
				return err
			}
		}
	}

	// 9. fdd-Add-UE-NR-Capabilities: ref
	{
		if ie.Fdd_Add_UE_NR_Capabilities != nil {
			if err := ie.Fdd_Add_UE_NR_Capabilities.Encode(e); err != nil {
				return err
			}
		}
	}

	// 10. tdd-Add-UE-NR-Capabilities: ref
	{
		if ie.Tdd_Add_UE_NR_Capabilities != nil {
			if err := ie.Tdd_Add_UE_NR_Capabilities.Encode(e); err != nil {
				return err
			}
		}
	}

	// 11. fr1-Add-UE-NR-Capabilities: ref
	{
		if ie.Fr1_Add_UE_NR_Capabilities != nil {
			if err := ie.Fr1_Add_UE_NR_Capabilities.Encode(e); err != nil {
				return err
			}
		}
	}

	// 12. fr2-Add-UE-NR-Capabilities: ref
	{
		if ie.Fr2_Add_UE_NR_Capabilities != nil {
			if err := ie.Fr2_Add_UE_NR_Capabilities.Encode(e); err != nil {
				return err
			}
		}
	}

	// 13. featureSets: ref
	{
		if ie.FeatureSets != nil {
			if err := ie.FeatureSets.Encode(e); err != nil {
				return err
			}
		}
	}

	// 14. featureSetCombinations: sequence-of
	{
		if ie.FeatureSetCombinations != nil {
			seqOf := e.NewSequenceOfEncoder(uENRCapabilityFeatureSetCombinationsConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.FeatureSetCombinations))); err != nil {
				return err
			}
			for i := range ie.FeatureSetCombinations {
				if err := ie.FeatureSetCombinations[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 15. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, uENRCapabilityLateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	// 16. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UE_NR_Capability) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uENRCapabilityConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. accessStratumRelease: ref
	{
		if err := ie.AccessStratumRelease.Decode(d); err != nil {
			return err
		}
	}

	// 3. pdcp-Parameters: ref
	{
		if err := ie.Pdcp_Parameters.Decode(d); err != nil {
			return err
		}
	}

	// 4. rlc-Parameters: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Rlc_Parameters = new(RLC_Parameters)
			if err := ie.Rlc_Parameters.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. mac-Parameters: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Mac_Parameters = new(MAC_Parameters)
			if err := ie.Mac_Parameters.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. phy-Parameters: ref
	{
		if err := ie.Phy_Parameters.Decode(d); err != nil {
			return err
		}
	}

	// 7. rf-Parameters: ref
	{
		if err := ie.Rf_Parameters.Decode(d); err != nil {
			return err
		}
	}

	// 8. measAndMobParameters: ref
	{
		if seq.IsComponentPresent(6) {
			ie.MeasAndMobParameters = new(MeasAndMobParameters)
			if err := ie.MeasAndMobParameters.Decode(d); err != nil {
				return err
			}
		}
	}

	// 9. fdd-Add-UE-NR-Capabilities: ref
	{
		if seq.IsComponentPresent(7) {
			ie.Fdd_Add_UE_NR_Capabilities = new(UE_NR_CapabilityAddXDD_Mode)
			if err := ie.Fdd_Add_UE_NR_Capabilities.Decode(d); err != nil {
				return err
			}
		}
	}

	// 10. tdd-Add-UE-NR-Capabilities: ref
	{
		if seq.IsComponentPresent(8) {
			ie.Tdd_Add_UE_NR_Capabilities = new(UE_NR_CapabilityAddXDD_Mode)
			if err := ie.Tdd_Add_UE_NR_Capabilities.Decode(d); err != nil {
				return err
			}
		}
	}

	// 11. fr1-Add-UE-NR-Capabilities: ref
	{
		if seq.IsComponentPresent(9) {
			ie.Fr1_Add_UE_NR_Capabilities = new(UE_NR_CapabilityAddFRX_Mode)
			if err := ie.Fr1_Add_UE_NR_Capabilities.Decode(d); err != nil {
				return err
			}
		}
	}

	// 12. fr2-Add-UE-NR-Capabilities: ref
	{
		if seq.IsComponentPresent(10) {
			ie.Fr2_Add_UE_NR_Capabilities = new(UE_NR_CapabilityAddFRX_Mode)
			if err := ie.Fr2_Add_UE_NR_Capabilities.Decode(d); err != nil {
				return err
			}
		}
	}

	// 13. featureSets: ref
	{
		if seq.IsComponentPresent(11) {
			ie.FeatureSets = new(FeatureSets)
			if err := ie.FeatureSets.Decode(d); err != nil {
				return err
			}
		}
	}

	// 14. featureSetCombinations: sequence-of
	{
		if seq.IsComponentPresent(12) {
			seqOf := d.NewSequenceOfDecoder(uENRCapabilityFeatureSetCombinationsConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeatureSetCombinations = make([]FeatureSetCombination, n)
			for i := int64(0); i < n; i++ {
				if err := ie.FeatureSetCombinations[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 15. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(13) {
			v, err := d.DecodeOctetString(uENRCapabilityLateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	// 16. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(14) {
			ie.NonCriticalExtension = new(UE_NR_Capability_v1530)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
