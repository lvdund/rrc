// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: UE-MRDC-Capability (line 25541).

var uEMRDCCapabilityConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "measAndMobParametersMRDC", Optional: true},
		{Name: "phy-ParametersMRDC-v1530", Optional: true},
		{Name: "rf-ParametersMRDC"},
		{Name: "generalParametersMRDC", Optional: true},
		{Name: "fdd-Add-UE-MRDC-Capabilities", Optional: true},
		{Name: "tdd-Add-UE-MRDC-Capabilities", Optional: true},
		{Name: "fr1-Add-UE-MRDC-Capabilities", Optional: true},
		{Name: "fr2-Add-UE-MRDC-Capabilities", Optional: true},
		{Name: "featureSetCombinations", Optional: true},
		{Name: "pdcp-ParametersMRDC-v1530", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var uEMRDCCapabilityFeatureSetCombinationsConstraints = per.SizeRange(1, common.MaxFeatureSetCombinations)

var uEMRDCCapabilityLateNonCriticalExtensionConstraints = per.SizeConstraints{}

type UE_MRDC_Capability struct {
	MeasAndMobParametersMRDC     *MeasAndMobParametersMRDC
	Phy_ParametersMRDC_v1530     *Phy_ParametersMRDC
	Rf_ParametersMRDC            RF_ParametersMRDC
	GeneralParametersMRDC        *GeneralParametersMRDC_XDD_Diff
	Fdd_Add_UE_MRDC_Capabilities *UE_MRDC_CapabilityAddXDD_Mode
	Tdd_Add_UE_MRDC_Capabilities *UE_MRDC_CapabilityAddXDD_Mode
	Fr1_Add_UE_MRDC_Capabilities *UE_MRDC_CapabilityAddFRX_Mode
	Fr2_Add_UE_MRDC_Capabilities *UE_MRDC_CapabilityAddFRX_Mode
	FeatureSetCombinations       []FeatureSetCombination
	Pdcp_ParametersMRDC_v1530    *PDCP_ParametersMRDC
	LateNonCriticalExtension     []byte
	NonCriticalExtension         *UE_MRDC_Capability_v1560
}

func (ie *UE_MRDC_Capability) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uEMRDCCapabilityConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MeasAndMobParametersMRDC != nil, ie.Phy_ParametersMRDC_v1530 != nil, ie.GeneralParametersMRDC != nil, ie.Fdd_Add_UE_MRDC_Capabilities != nil, ie.Tdd_Add_UE_MRDC_Capabilities != nil, ie.Fr1_Add_UE_MRDC_Capabilities != nil, ie.Fr2_Add_UE_MRDC_Capabilities != nil, ie.FeatureSetCombinations != nil, ie.Pdcp_ParametersMRDC_v1530 != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. measAndMobParametersMRDC: ref
	{
		if ie.MeasAndMobParametersMRDC != nil {
			if err := ie.MeasAndMobParametersMRDC.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. phy-ParametersMRDC-v1530: ref
	{
		if ie.Phy_ParametersMRDC_v1530 != nil {
			if err := ie.Phy_ParametersMRDC_v1530.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. rf-ParametersMRDC: ref
	{
		if err := ie.Rf_ParametersMRDC.Encode(e); err != nil {
			return err
		}
	}

	// 5. generalParametersMRDC: ref
	{
		if ie.GeneralParametersMRDC != nil {
			if err := ie.GeneralParametersMRDC.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. fdd-Add-UE-MRDC-Capabilities: ref
	{
		if ie.Fdd_Add_UE_MRDC_Capabilities != nil {
			if err := ie.Fdd_Add_UE_MRDC_Capabilities.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. tdd-Add-UE-MRDC-Capabilities: ref
	{
		if ie.Tdd_Add_UE_MRDC_Capabilities != nil {
			if err := ie.Tdd_Add_UE_MRDC_Capabilities.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. fr1-Add-UE-MRDC-Capabilities: ref
	{
		if ie.Fr1_Add_UE_MRDC_Capabilities != nil {
			if err := ie.Fr1_Add_UE_MRDC_Capabilities.Encode(e); err != nil {
				return err
			}
		}
	}

	// 9. fr2-Add-UE-MRDC-Capabilities: ref
	{
		if ie.Fr2_Add_UE_MRDC_Capabilities != nil {
			if err := ie.Fr2_Add_UE_MRDC_Capabilities.Encode(e); err != nil {
				return err
			}
		}
	}

	// 10. featureSetCombinations: sequence-of
	{
		if ie.FeatureSetCombinations != nil {
			seqOf := e.NewSequenceOfEncoder(uEMRDCCapabilityFeatureSetCombinationsConstraints)
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

	// 11. pdcp-ParametersMRDC-v1530: ref
	{
		if ie.Pdcp_ParametersMRDC_v1530 != nil {
			if err := ie.Pdcp_ParametersMRDC_v1530.Encode(e); err != nil {
				return err
			}
		}
	}

	// 12. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, uEMRDCCapabilityLateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	// 13. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UE_MRDC_Capability) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uEMRDCCapabilityConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. measAndMobParametersMRDC: ref
	{
		if seq.IsComponentPresent(0) {
			ie.MeasAndMobParametersMRDC = new(MeasAndMobParametersMRDC)
			if err := ie.MeasAndMobParametersMRDC.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. phy-ParametersMRDC-v1530: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Phy_ParametersMRDC_v1530 = new(Phy_ParametersMRDC)
			if err := ie.Phy_ParametersMRDC_v1530.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. rf-ParametersMRDC: ref
	{
		if err := ie.Rf_ParametersMRDC.Decode(d); err != nil {
			return err
		}
	}

	// 5. generalParametersMRDC: ref
	{
		if seq.IsComponentPresent(3) {
			ie.GeneralParametersMRDC = new(GeneralParametersMRDC_XDD_Diff)
			if err := ie.GeneralParametersMRDC.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. fdd-Add-UE-MRDC-Capabilities: ref
	{
		if seq.IsComponentPresent(4) {
			ie.Fdd_Add_UE_MRDC_Capabilities = new(UE_MRDC_CapabilityAddXDD_Mode)
			if err := ie.Fdd_Add_UE_MRDC_Capabilities.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. tdd-Add-UE-MRDC-Capabilities: ref
	{
		if seq.IsComponentPresent(5) {
			ie.Tdd_Add_UE_MRDC_Capabilities = new(UE_MRDC_CapabilityAddXDD_Mode)
			if err := ie.Tdd_Add_UE_MRDC_Capabilities.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. fr1-Add-UE-MRDC-Capabilities: ref
	{
		if seq.IsComponentPresent(6) {
			ie.Fr1_Add_UE_MRDC_Capabilities = new(UE_MRDC_CapabilityAddFRX_Mode)
			if err := ie.Fr1_Add_UE_MRDC_Capabilities.Decode(d); err != nil {
				return err
			}
		}
	}

	// 9. fr2-Add-UE-MRDC-Capabilities: ref
	{
		if seq.IsComponentPresent(7) {
			ie.Fr2_Add_UE_MRDC_Capabilities = new(UE_MRDC_CapabilityAddFRX_Mode)
			if err := ie.Fr2_Add_UE_MRDC_Capabilities.Decode(d); err != nil {
				return err
			}
		}
	}

	// 10. featureSetCombinations: sequence-of
	{
		if seq.IsComponentPresent(8) {
			seqOf := d.NewSequenceOfDecoder(uEMRDCCapabilityFeatureSetCombinationsConstraints)
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

	// 11. pdcp-ParametersMRDC-v1530: ref
	{
		if seq.IsComponentPresent(9) {
			ie.Pdcp_ParametersMRDC_v1530 = new(PDCP_ParametersMRDC)
			if err := ie.Pdcp_ParametersMRDC_v1530.Decode(d); err != nil {
				return err
			}
		}
	}

	// 12. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(10) {
			v, err := d.DecodeOctetString(uEMRDCCapabilityLateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	// 13. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(11) {
			ie.NonCriticalExtension = new(UE_MRDC_Capability_v1560)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
