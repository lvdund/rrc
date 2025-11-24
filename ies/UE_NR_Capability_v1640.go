package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_NR_Capability_v1640 struct {
	RedirectAtResumeByNAS_r16                *UE_NR_Capability_v1640_redirectAtResumeByNAS_r16 `optional`
	Phy_ParametersSharedSpectrumChAccess_r16 *Phy_ParametersSharedSpectrumChAccess_r16         `optional`
	NonCriticalExtension                     *UE_NR_Capability_v1650                           `optional`
}

func (ie *UE_NR_Capability_v1640) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.RedirectAtResumeByNAS_r16 != nil, ie.Phy_ParametersSharedSpectrumChAccess_r16 != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.RedirectAtResumeByNAS_r16 != nil {
		if err = ie.RedirectAtResumeByNAS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode RedirectAtResumeByNAS_r16", err)
		}
	}
	if ie.Phy_ParametersSharedSpectrumChAccess_r16 != nil {
		if err = ie.Phy_ParametersSharedSpectrumChAccess_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Phy_ParametersSharedSpectrumChAccess_r16", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UE_NR_Capability_v1640) Decode(r *uper.UperReader) error {
	var err error
	var RedirectAtResumeByNAS_r16Present bool
	if RedirectAtResumeByNAS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Phy_ParametersSharedSpectrumChAccess_r16Present bool
	if Phy_ParametersSharedSpectrumChAccess_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if RedirectAtResumeByNAS_r16Present {
		ie.RedirectAtResumeByNAS_r16 = new(UE_NR_Capability_v1640_redirectAtResumeByNAS_r16)
		if err = ie.RedirectAtResumeByNAS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode RedirectAtResumeByNAS_r16", err)
		}
	}
	if Phy_ParametersSharedSpectrumChAccess_r16Present {
		ie.Phy_ParametersSharedSpectrumChAccess_r16 = new(Phy_ParametersSharedSpectrumChAccess_r16)
		if err = ie.Phy_ParametersSharedSpectrumChAccess_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Phy_ParametersSharedSpectrumChAccess_r16", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(UE_NR_Capability_v1650)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
