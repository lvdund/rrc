package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UE_MRDC_Capability struct {
	MeasAndMobParametersMRDC     *MeasAndMobParametersMRDC       `optional`
	Phy_ParametersMRDC_v1530     *Phy_ParametersMRDC             `optional`
	Rf_ParametersMRDC            RF_ParametersMRDC               `madatory`
	GeneralParametersMRDC        *GeneralParametersMRDC_XDD_Diff `optional`
	Fdd_Add_UE_MRDC_Capabilities *UE_MRDC_CapabilityAddXDD_Mode  `optional`
	Tdd_Add_UE_MRDC_Capabilities *UE_MRDC_CapabilityAddXDD_Mode  `optional`
	Fr1_Add_UE_MRDC_Capabilities *UE_MRDC_CapabilityAddFRX_Mode  `optional`
	Fr2_Add_UE_MRDC_Capabilities *UE_MRDC_CapabilityAddFRX_Mode  `optional`
	FeatureSetCombinations       []FeatureSetCombination         `lb:1,ub:maxFeatureSetCombinations,optional`
	Pdcp_ParametersMRDC_v1530    *PDCP_ParametersMRDC            `optional`
	LateNonCriticalExtension     *[]byte                         `optional`
	NonCriticalExtension         *UE_MRDC_Capability_v1560       `optional`
}

func (ie *UE_MRDC_Capability) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.MeasAndMobParametersMRDC != nil, ie.Phy_ParametersMRDC_v1530 != nil, ie.GeneralParametersMRDC != nil, ie.Fdd_Add_UE_MRDC_Capabilities != nil, ie.Tdd_Add_UE_MRDC_Capabilities != nil, ie.Fr1_Add_UE_MRDC_Capabilities != nil, ie.Fr2_Add_UE_MRDC_Capabilities != nil, len(ie.FeatureSetCombinations) > 0, ie.Pdcp_ParametersMRDC_v1530 != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MeasAndMobParametersMRDC != nil {
		if err = ie.MeasAndMobParametersMRDC.Encode(w); err != nil {
			return utils.WrapError("Encode MeasAndMobParametersMRDC", err)
		}
	}
	if ie.Phy_ParametersMRDC_v1530 != nil {
		if err = ie.Phy_ParametersMRDC_v1530.Encode(w); err != nil {
			return utils.WrapError("Encode Phy_ParametersMRDC_v1530", err)
		}
	}
	if err = ie.Rf_ParametersMRDC.Encode(w); err != nil {
		return utils.WrapError("Encode Rf_ParametersMRDC", err)
	}
	if ie.GeneralParametersMRDC != nil {
		if err = ie.GeneralParametersMRDC.Encode(w); err != nil {
			return utils.WrapError("Encode GeneralParametersMRDC", err)
		}
	}
	if ie.Fdd_Add_UE_MRDC_Capabilities != nil {
		if err = ie.Fdd_Add_UE_MRDC_Capabilities.Encode(w); err != nil {
			return utils.WrapError("Encode Fdd_Add_UE_MRDC_Capabilities", err)
		}
	}
	if ie.Tdd_Add_UE_MRDC_Capabilities != nil {
		if err = ie.Tdd_Add_UE_MRDC_Capabilities.Encode(w); err != nil {
			return utils.WrapError("Encode Tdd_Add_UE_MRDC_Capabilities", err)
		}
	}
	if ie.Fr1_Add_UE_MRDC_Capabilities != nil {
		if err = ie.Fr1_Add_UE_MRDC_Capabilities.Encode(w); err != nil {
			return utils.WrapError("Encode Fr1_Add_UE_MRDC_Capabilities", err)
		}
	}
	if ie.Fr2_Add_UE_MRDC_Capabilities != nil {
		if err = ie.Fr2_Add_UE_MRDC_Capabilities.Encode(w); err != nil {
			return utils.WrapError("Encode Fr2_Add_UE_MRDC_Capabilities", err)
		}
	}
	if len(ie.FeatureSetCombinations) > 0 {
		tmp_FeatureSetCombinations := utils.NewSequence[*FeatureSetCombination]([]*FeatureSetCombination{}, aper.Constraint{Lb: 1, Ub: maxFeatureSetCombinations}, false)
		for _, i := range ie.FeatureSetCombinations {
			tmp_FeatureSetCombinations.Value = append(tmp_FeatureSetCombinations.Value, &i)
		}
		if err = tmp_FeatureSetCombinations.Encode(w); err != nil {
			return utils.WrapError("Encode FeatureSetCombinations", err)
		}
	}
	if ie.Pdcp_ParametersMRDC_v1530 != nil {
		if err = ie.Pdcp_ParametersMRDC_v1530.Encode(w); err != nil {
			return utils.WrapError("Encode Pdcp_ParametersMRDC_v1530", err)
		}
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, nil, false); err != nil {
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

func (ie *UE_MRDC_Capability) Decode(r *aper.AperReader) error {
	var err error
	var MeasAndMobParametersMRDCPresent bool
	if MeasAndMobParametersMRDCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Phy_ParametersMRDC_v1530Present bool
	if Phy_ParametersMRDC_v1530Present, err = r.ReadBool(); err != nil {
		return err
	}
	var GeneralParametersMRDCPresent bool
	if GeneralParametersMRDCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Fdd_Add_UE_MRDC_CapabilitiesPresent bool
	if Fdd_Add_UE_MRDC_CapabilitiesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Tdd_Add_UE_MRDC_CapabilitiesPresent bool
	if Tdd_Add_UE_MRDC_CapabilitiesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Fr1_Add_UE_MRDC_CapabilitiesPresent bool
	if Fr1_Add_UE_MRDC_CapabilitiesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Fr2_Add_UE_MRDC_CapabilitiesPresent bool
	if Fr2_Add_UE_MRDC_CapabilitiesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var FeatureSetCombinationsPresent bool
	if FeatureSetCombinationsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdcp_ParametersMRDC_v1530Present bool
	if Pdcp_ParametersMRDC_v1530Present, err = r.ReadBool(); err != nil {
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
	if MeasAndMobParametersMRDCPresent {
		ie.MeasAndMobParametersMRDC = new(MeasAndMobParametersMRDC)
		if err = ie.MeasAndMobParametersMRDC.Decode(r); err != nil {
			return utils.WrapError("Decode MeasAndMobParametersMRDC", err)
		}
	}
	if Phy_ParametersMRDC_v1530Present {
		ie.Phy_ParametersMRDC_v1530 = new(Phy_ParametersMRDC)
		if err = ie.Phy_ParametersMRDC_v1530.Decode(r); err != nil {
			return utils.WrapError("Decode Phy_ParametersMRDC_v1530", err)
		}
	}
	if err = ie.Rf_ParametersMRDC.Decode(r); err != nil {
		return utils.WrapError("Decode Rf_ParametersMRDC", err)
	}
	if GeneralParametersMRDCPresent {
		ie.GeneralParametersMRDC = new(GeneralParametersMRDC_XDD_Diff)
		if err = ie.GeneralParametersMRDC.Decode(r); err != nil {
			return utils.WrapError("Decode GeneralParametersMRDC", err)
		}
	}
	if Fdd_Add_UE_MRDC_CapabilitiesPresent {
		ie.Fdd_Add_UE_MRDC_Capabilities = new(UE_MRDC_CapabilityAddXDD_Mode)
		if err = ie.Fdd_Add_UE_MRDC_Capabilities.Decode(r); err != nil {
			return utils.WrapError("Decode Fdd_Add_UE_MRDC_Capabilities", err)
		}
	}
	if Tdd_Add_UE_MRDC_CapabilitiesPresent {
		ie.Tdd_Add_UE_MRDC_Capabilities = new(UE_MRDC_CapabilityAddXDD_Mode)
		if err = ie.Tdd_Add_UE_MRDC_Capabilities.Decode(r); err != nil {
			return utils.WrapError("Decode Tdd_Add_UE_MRDC_Capabilities", err)
		}
	}
	if Fr1_Add_UE_MRDC_CapabilitiesPresent {
		ie.Fr1_Add_UE_MRDC_Capabilities = new(UE_MRDC_CapabilityAddFRX_Mode)
		if err = ie.Fr1_Add_UE_MRDC_Capabilities.Decode(r); err != nil {
			return utils.WrapError("Decode Fr1_Add_UE_MRDC_Capabilities", err)
		}
	}
	if Fr2_Add_UE_MRDC_CapabilitiesPresent {
		ie.Fr2_Add_UE_MRDC_Capabilities = new(UE_MRDC_CapabilityAddFRX_Mode)
		if err = ie.Fr2_Add_UE_MRDC_Capabilities.Decode(r); err != nil {
			return utils.WrapError("Decode Fr2_Add_UE_MRDC_Capabilities", err)
		}
	}
	if FeatureSetCombinationsPresent {
		tmp_FeatureSetCombinations := utils.NewSequence[*FeatureSetCombination]([]*FeatureSetCombination{}, aper.Constraint{Lb: 1, Ub: maxFeatureSetCombinations}, false)
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
	if Pdcp_ParametersMRDC_v1530Present {
		ie.Pdcp_ParametersMRDC_v1530 = new(PDCP_ParametersMRDC)
		if err = ie.Pdcp_ParametersMRDC_v1530.Decode(r); err != nil {
			return utils.WrapError("Decode Pdcp_ParametersMRDC_v1530", err)
		}
	}
	if LateNonCriticalExtensionPresent {
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = &tmp_os_LateNonCriticalExtension
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(UE_MRDC_Capability_v1560)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
