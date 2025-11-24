package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_NR_Capability struct {
	AccessStratumRelease       AccessStratumRelease         `madatory`
	Pdcp_Parameters            PDCP_Parameters              `madatory`
	Rlc_Parameters             *RLC_Parameters              `optional`
	Mac_Parameters             *MAC_Parameters              `optional`
	Phy_Parameters             Phy_Parameters               `madatory`
	Rf_Parameters              RF_Parameters                `madatory`
	MeasAndMobParameters       *MeasAndMobParameters        `optional`
	Fdd_Add_UE_NR_Capabilities *UE_NR_CapabilityAddXDD_Mode `optional`
	Tdd_Add_UE_NR_Capabilities *UE_NR_CapabilityAddXDD_Mode `optional`
	Fr1_Add_UE_NR_Capabilities *UE_NR_CapabilityAddFRX_Mode `optional`
	Fr2_Add_UE_NR_Capabilities *UE_NR_CapabilityAddFRX_Mode `optional`
	FeatureSets                *FeatureSets                 `optional`
	FeatureSetCombinations     []FeatureSetCombination      `lb:1,ub:maxFeatureSetCombinations,optional`
	LateNonCriticalExtension   *[]byte                      `optional`
	NonCriticalExtension       *UE_NR_Capability_v1530      `optional`
}

func (ie *UE_NR_Capability) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Rlc_Parameters != nil, ie.Mac_Parameters != nil, ie.MeasAndMobParameters != nil, ie.Fdd_Add_UE_NR_Capabilities != nil, ie.Tdd_Add_UE_NR_Capabilities != nil, ie.Fr1_Add_UE_NR_Capabilities != nil, ie.Fr2_Add_UE_NR_Capabilities != nil, ie.FeatureSets != nil, len(ie.FeatureSetCombinations) > 0, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.AccessStratumRelease.Encode(w); err != nil {
		return utils.WrapError("Encode AccessStratumRelease", err)
	}
	if err = ie.Pdcp_Parameters.Encode(w); err != nil {
		return utils.WrapError("Encode Pdcp_Parameters", err)
	}
	if ie.Rlc_Parameters != nil {
		if err = ie.Rlc_Parameters.Encode(w); err != nil {
			return utils.WrapError("Encode Rlc_Parameters", err)
		}
	}
	if ie.Mac_Parameters != nil {
		if err = ie.Mac_Parameters.Encode(w); err != nil {
			return utils.WrapError("Encode Mac_Parameters", err)
		}
	}
	if err = ie.Phy_Parameters.Encode(w); err != nil {
		return utils.WrapError("Encode Phy_Parameters", err)
	}
	if err = ie.Rf_Parameters.Encode(w); err != nil {
		return utils.WrapError("Encode Rf_Parameters", err)
	}
	if ie.MeasAndMobParameters != nil {
		if err = ie.MeasAndMobParameters.Encode(w); err != nil {
			return utils.WrapError("Encode MeasAndMobParameters", err)
		}
	}
	if ie.Fdd_Add_UE_NR_Capabilities != nil {
		if err = ie.Fdd_Add_UE_NR_Capabilities.Encode(w); err != nil {
			return utils.WrapError("Encode Fdd_Add_UE_NR_Capabilities", err)
		}
	}
	if ie.Tdd_Add_UE_NR_Capabilities != nil {
		if err = ie.Tdd_Add_UE_NR_Capabilities.Encode(w); err != nil {
			return utils.WrapError("Encode Tdd_Add_UE_NR_Capabilities", err)
		}
	}
	if ie.Fr1_Add_UE_NR_Capabilities != nil {
		if err = ie.Fr1_Add_UE_NR_Capabilities.Encode(w); err != nil {
			return utils.WrapError("Encode Fr1_Add_UE_NR_Capabilities", err)
		}
	}
	if ie.Fr2_Add_UE_NR_Capabilities != nil {
		if err = ie.Fr2_Add_UE_NR_Capabilities.Encode(w); err != nil {
			return utils.WrapError("Encode Fr2_Add_UE_NR_Capabilities", err)
		}
	}
	if ie.FeatureSets != nil {
		if err = ie.FeatureSets.Encode(w); err != nil {
			return utils.WrapError("Encode FeatureSets", err)
		}
	}
	if len(ie.FeatureSetCombinations) > 0 {
		tmp_FeatureSetCombinations := utils.NewSequence[*FeatureSetCombination]([]*FeatureSetCombination{}, uper.Constraint{Lb: 1, Ub: maxFeatureSetCombinations}, false)
		for _, i := range ie.FeatureSetCombinations {
			tmp_FeatureSetCombinations.Value = append(tmp_FeatureSetCombinations.Value, &i)
		}
		if err = tmp_FeatureSetCombinations.Encode(w); err != nil {
			return utils.WrapError("Encode FeatureSetCombinations", err)
		}
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode LateNonCriticalExtension", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UE_NR_Capability) Decode(r *uper.UperReader) error {
	var err error
	var Rlc_ParametersPresent bool
	if Rlc_ParametersPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Mac_ParametersPresent bool
	if Mac_ParametersPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasAndMobParametersPresent bool
	if MeasAndMobParametersPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Fdd_Add_UE_NR_CapabilitiesPresent bool
	if Fdd_Add_UE_NR_CapabilitiesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Tdd_Add_UE_NR_CapabilitiesPresent bool
	if Tdd_Add_UE_NR_CapabilitiesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Fr1_Add_UE_NR_CapabilitiesPresent bool
	if Fr1_Add_UE_NR_CapabilitiesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Fr2_Add_UE_NR_CapabilitiesPresent bool
	if Fr2_Add_UE_NR_CapabilitiesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var FeatureSetsPresent bool
	if FeatureSetsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var FeatureSetCombinationsPresent bool
	if FeatureSetCombinationsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var LateNonCriticalExtensionPresent bool
	if LateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.AccessStratumRelease.Decode(r); err != nil {
		return utils.WrapError("Decode AccessStratumRelease", err)
	}
	if err = ie.Pdcp_Parameters.Decode(r); err != nil {
		return utils.WrapError("Decode Pdcp_Parameters", err)
	}
	if Rlc_ParametersPresent {
		ie.Rlc_Parameters = new(RLC_Parameters)
		if err = ie.Rlc_Parameters.Decode(r); err != nil {
			return utils.WrapError("Decode Rlc_Parameters", err)
		}
	}
	if Mac_ParametersPresent {
		ie.Mac_Parameters = new(MAC_Parameters)
		if err = ie.Mac_Parameters.Decode(r); err != nil {
			return utils.WrapError("Decode Mac_Parameters", err)
		}
	}
	if err = ie.Phy_Parameters.Decode(r); err != nil {
		return utils.WrapError("Decode Phy_Parameters", err)
	}
	if err = ie.Rf_Parameters.Decode(r); err != nil {
		return utils.WrapError("Decode Rf_Parameters", err)
	}
	if MeasAndMobParametersPresent {
		ie.MeasAndMobParameters = new(MeasAndMobParameters)
		if err = ie.MeasAndMobParameters.Decode(r); err != nil {
			return utils.WrapError("Decode MeasAndMobParameters", err)
		}
	}
	if Fdd_Add_UE_NR_CapabilitiesPresent {
		ie.Fdd_Add_UE_NR_Capabilities = new(UE_NR_CapabilityAddXDD_Mode)
		if err = ie.Fdd_Add_UE_NR_Capabilities.Decode(r); err != nil {
			return utils.WrapError("Decode Fdd_Add_UE_NR_Capabilities", err)
		}
	}
	if Tdd_Add_UE_NR_CapabilitiesPresent {
		ie.Tdd_Add_UE_NR_Capabilities = new(UE_NR_CapabilityAddXDD_Mode)
		if err = ie.Tdd_Add_UE_NR_Capabilities.Decode(r); err != nil {
			return utils.WrapError("Decode Tdd_Add_UE_NR_Capabilities", err)
		}
	}
	if Fr1_Add_UE_NR_CapabilitiesPresent {
		ie.Fr1_Add_UE_NR_Capabilities = new(UE_NR_CapabilityAddFRX_Mode)
		if err = ie.Fr1_Add_UE_NR_Capabilities.Decode(r); err != nil {
			return utils.WrapError("Decode Fr1_Add_UE_NR_Capabilities", err)
		}
	}
	if Fr2_Add_UE_NR_CapabilitiesPresent {
		ie.Fr2_Add_UE_NR_Capabilities = new(UE_NR_CapabilityAddFRX_Mode)
		if err = ie.Fr2_Add_UE_NR_Capabilities.Decode(r); err != nil {
			return utils.WrapError("Decode Fr2_Add_UE_NR_Capabilities", err)
		}
	}
	if FeatureSetsPresent {
		ie.FeatureSets = new(FeatureSets)
		if err = ie.FeatureSets.Decode(r); err != nil {
			return utils.WrapError("Decode FeatureSets", err)
		}
	}
	if FeatureSetCombinationsPresent {
		tmp_FeatureSetCombinations := utils.NewSequence[*FeatureSetCombination]([]*FeatureSetCombination{}, uper.Constraint{Lb: 1, Ub: maxFeatureSetCombinations}, false)
		fn_FeatureSetCombinations := func() *FeatureSetCombination {
			return new(FeatureSetCombination)
		}
		if err = tmp_FeatureSetCombinations.Decode(r, fn_FeatureSetCombinations); err != nil {
			return utils.WrapError("Decode FeatureSetCombinations", err)
		}
		ie.FeatureSetCombinations = []FeatureSetCombination{}
		for _, i := range tmp_FeatureSetCombinations.Value {
			ie.FeatureSetCombinations = append(ie.FeatureSetCombinations, *i)
		}
	}
	if LateNonCriticalExtensionPresent {
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = &tmp_os_LateNonCriticalExtension
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(UE_NR_Capability_v1530)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
