package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UECapabilityInformation_IEs struct {
	Ue_CapabilityRAT_ContainerList *UE_CapabilityRAT_ContainerList `optional`
	LateNonCriticalExtension       *[]byte                         `optional`
	NonCriticalExtension           interface{}                     `optional`
}

func (ie *UECapabilityInformation_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Ue_CapabilityRAT_ContainerList != nil, ie.LateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ue_CapabilityRAT_ContainerList != nil {
		if err = ie.Ue_CapabilityRAT_ContainerList.Encode(w); err != nil {
			return utils.WrapError("Encode Ue_CapabilityRAT_ContainerList", err)
		}
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, nil, false); err != nil {
			return utils.WrapError("Encode LateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UECapabilityInformation_IEs) Decode(r *aper.AperReader) error {
	var err error
	var Ue_CapabilityRAT_ContainerListPresent bool
	if Ue_CapabilityRAT_ContainerListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var LateNonCriticalExtensionPresent bool
	if LateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Ue_CapabilityRAT_ContainerListPresent {
		ie.Ue_CapabilityRAT_ContainerList = new(UE_CapabilityRAT_ContainerList)
		if err = ie.Ue_CapabilityRAT_ContainerList.Decode(r); err != nil {
			return utils.WrapError("Decode Ue_CapabilityRAT_ContainerList", err)
		}
	}
	if LateNonCriticalExtensionPresent {
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = &tmp_os_LateNonCriticalExtension
	}
	return nil
}
