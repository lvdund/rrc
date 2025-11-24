package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_NR_Capability_v1540 struct {
	Sdap_Parameters                  *SDAP_Parameters                       `optional`
	OverheatingInd                   *UE_NR_Capability_v1540_overheatingInd `optional`
	Ims_Parameters                   *IMS_Parameters                        `optional`
	Fr1_Add_UE_NR_Capabilities_v1540 *UE_NR_CapabilityAddFRX_Mode_v1540     `optional`
	Fr2_Add_UE_NR_Capabilities_v1540 *UE_NR_CapabilityAddFRX_Mode_v1540     `optional`
	Fr1_fr2_Add_UE_NR_Capabilities   *UE_NR_CapabilityAddFRX_Mode           `optional`
	NonCriticalExtension             *UE_NR_Capability_v1550                `optional`
}

func (ie *UE_NR_Capability_v1540) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Sdap_Parameters != nil, ie.OverheatingInd != nil, ie.Ims_Parameters != nil, ie.Fr1_Add_UE_NR_Capabilities_v1540 != nil, ie.Fr2_Add_UE_NR_Capabilities_v1540 != nil, ie.Fr1_fr2_Add_UE_NR_Capabilities != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sdap_Parameters != nil {
		if err = ie.Sdap_Parameters.Encode(w); err != nil {
			return utils.WrapError("Encode Sdap_Parameters", err)
		}
	}
	if ie.OverheatingInd != nil {
		if err = ie.OverheatingInd.Encode(w); err != nil {
			return utils.WrapError("Encode OverheatingInd", err)
		}
	}
	if ie.Ims_Parameters != nil {
		if err = ie.Ims_Parameters.Encode(w); err != nil {
			return utils.WrapError("Encode Ims_Parameters", err)
		}
	}
	if ie.Fr1_Add_UE_NR_Capabilities_v1540 != nil {
		if err = ie.Fr1_Add_UE_NR_Capabilities_v1540.Encode(w); err != nil {
			return utils.WrapError("Encode Fr1_Add_UE_NR_Capabilities_v1540", err)
		}
	}
	if ie.Fr2_Add_UE_NR_Capabilities_v1540 != nil {
		if err = ie.Fr2_Add_UE_NR_Capabilities_v1540.Encode(w); err != nil {
			return utils.WrapError("Encode Fr2_Add_UE_NR_Capabilities_v1540", err)
		}
	}
	if ie.Fr1_fr2_Add_UE_NR_Capabilities != nil {
		if err = ie.Fr1_fr2_Add_UE_NR_Capabilities.Encode(w); err != nil {
			return utils.WrapError("Encode Fr1_fr2_Add_UE_NR_Capabilities", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UE_NR_Capability_v1540) Decode(r *uper.UperReader) error {
	var err error
	var Sdap_ParametersPresent bool
	if Sdap_ParametersPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var OverheatingIndPresent bool
	if OverheatingIndPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Ims_ParametersPresent bool
	if Ims_ParametersPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Fr1_Add_UE_NR_Capabilities_v1540Present bool
	if Fr1_Add_UE_NR_Capabilities_v1540Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Fr2_Add_UE_NR_Capabilities_v1540Present bool
	if Fr2_Add_UE_NR_Capabilities_v1540Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Fr1_fr2_Add_UE_NR_CapabilitiesPresent bool
	if Fr1_fr2_Add_UE_NR_CapabilitiesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Sdap_ParametersPresent {
		ie.Sdap_Parameters = new(SDAP_Parameters)
		if err = ie.Sdap_Parameters.Decode(r); err != nil {
			return utils.WrapError("Decode Sdap_Parameters", err)
		}
	}
	if OverheatingIndPresent {
		ie.OverheatingInd = new(UE_NR_Capability_v1540_overheatingInd)
		if err = ie.OverheatingInd.Decode(r); err != nil {
			return utils.WrapError("Decode OverheatingInd", err)
		}
	}
	if Ims_ParametersPresent {
		ie.Ims_Parameters = new(IMS_Parameters)
		if err = ie.Ims_Parameters.Decode(r); err != nil {
			return utils.WrapError("Decode Ims_Parameters", err)
		}
	}
	if Fr1_Add_UE_NR_Capabilities_v1540Present {
		ie.Fr1_Add_UE_NR_Capabilities_v1540 = new(UE_NR_CapabilityAddFRX_Mode_v1540)
		if err = ie.Fr1_Add_UE_NR_Capabilities_v1540.Decode(r); err != nil {
			return utils.WrapError("Decode Fr1_Add_UE_NR_Capabilities_v1540", err)
		}
	}
	if Fr2_Add_UE_NR_Capabilities_v1540Present {
		ie.Fr2_Add_UE_NR_Capabilities_v1540 = new(UE_NR_CapabilityAddFRX_Mode_v1540)
		if err = ie.Fr2_Add_UE_NR_Capabilities_v1540.Decode(r); err != nil {
			return utils.WrapError("Decode Fr2_Add_UE_NR_Capabilities_v1540", err)
		}
	}
	if Fr1_fr2_Add_UE_NR_CapabilitiesPresent {
		ie.Fr1_fr2_Add_UE_NR_Capabilities = new(UE_NR_CapabilityAddFRX_Mode)
		if err = ie.Fr1_fr2_Add_UE_NR_Capabilities.Decode(r); err != nil {
			return utils.WrapError("Decode Fr1_fr2_Add_UE_NR_Capabilities", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(UE_NR_Capability_v1550)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
