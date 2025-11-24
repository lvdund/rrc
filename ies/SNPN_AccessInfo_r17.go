package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SNPN_AccessInfo_r17 struct {
	ExtCH_Supported_r17            *SNPN_AccessInfo_r17_extCH_Supported_r17            `optional`
	ExtCH_WithoutConfigAllowed_r17 *SNPN_AccessInfo_r17_extCH_WithoutConfigAllowed_r17 `optional`
	OnboardingEnabled_r17          *SNPN_AccessInfo_r17_onboardingEnabled_r17          `optional`
	ImsEmergencySupportForSNPN_r17 *SNPN_AccessInfo_r17_imsEmergencySupportForSNPN_r17 `optional`
}

func (ie *SNPN_AccessInfo_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ExtCH_Supported_r17 != nil, ie.ExtCH_WithoutConfigAllowed_r17 != nil, ie.OnboardingEnabled_r17 != nil, ie.ImsEmergencySupportForSNPN_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ExtCH_Supported_r17 != nil {
		if err = ie.ExtCH_Supported_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ExtCH_Supported_r17", err)
		}
	}
	if ie.ExtCH_WithoutConfigAllowed_r17 != nil {
		if err = ie.ExtCH_WithoutConfigAllowed_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ExtCH_WithoutConfigAllowed_r17", err)
		}
	}
	if ie.OnboardingEnabled_r17 != nil {
		if err = ie.OnboardingEnabled_r17.Encode(w); err != nil {
			return utils.WrapError("Encode OnboardingEnabled_r17", err)
		}
	}
	if ie.ImsEmergencySupportForSNPN_r17 != nil {
		if err = ie.ImsEmergencySupportForSNPN_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ImsEmergencySupportForSNPN_r17", err)
		}
	}
	return nil
}

func (ie *SNPN_AccessInfo_r17) Decode(r *uper.UperReader) error {
	var err error
	var ExtCH_Supported_r17Present bool
	if ExtCH_Supported_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ExtCH_WithoutConfigAllowed_r17Present bool
	if ExtCH_WithoutConfigAllowed_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var OnboardingEnabled_r17Present bool
	if OnboardingEnabled_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ImsEmergencySupportForSNPN_r17Present bool
	if ImsEmergencySupportForSNPN_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ExtCH_Supported_r17Present {
		ie.ExtCH_Supported_r17 = new(SNPN_AccessInfo_r17_extCH_Supported_r17)
		if err = ie.ExtCH_Supported_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ExtCH_Supported_r17", err)
		}
	}
	if ExtCH_WithoutConfigAllowed_r17Present {
		ie.ExtCH_WithoutConfigAllowed_r17 = new(SNPN_AccessInfo_r17_extCH_WithoutConfigAllowed_r17)
		if err = ie.ExtCH_WithoutConfigAllowed_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ExtCH_WithoutConfigAllowed_r17", err)
		}
	}
	if OnboardingEnabled_r17Present {
		ie.OnboardingEnabled_r17 = new(SNPN_AccessInfo_r17_onboardingEnabled_r17)
		if err = ie.OnboardingEnabled_r17.Decode(r); err != nil {
			return utils.WrapError("Decode OnboardingEnabled_r17", err)
		}
	}
	if ImsEmergencySupportForSNPN_r17Present {
		ie.ImsEmergencySupportForSNPN_r17 = new(SNPN_AccessInfo_r17_imsEmergencySupportForSNPN_r17)
		if err = ie.ImsEmergencySupportForSNPN_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ImsEmergencySupportForSNPN_r17", err)
		}
	}
	return nil
}
