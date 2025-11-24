package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UECapabilityEnquiry_IEs struct {
	Ue_CapabilityRAT_RequestList UE_CapabilityRAT_RequestList `madatory`
	LateNonCriticalExtension     *[]byte                      `optional`
	Ue_CapabilityEnquiryExt      *[]byte                      `optional`
}

func (ie *UECapabilityEnquiry_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.LateNonCriticalExtension != nil, ie.Ue_CapabilityEnquiryExt != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Ue_CapabilityRAT_RequestList.Encode(w); err != nil {
		return utils.WrapError("Encode Ue_CapabilityRAT_RequestList", err)
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode LateNonCriticalExtension", err)
		}
	}
	if ie.Ue_CapabilityEnquiryExt != nil {
		if err = w.WriteOctetString(*ie.Ue_CapabilityEnquiryExt, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode Ue_CapabilityEnquiryExt", err)
		}
	}
	return nil
}

func (ie *UECapabilityEnquiry_IEs) Decode(r *uper.UperReader) error {
	var err error
	var LateNonCriticalExtensionPresent bool
	if LateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Ue_CapabilityEnquiryExtPresent bool
	if Ue_CapabilityEnquiryExtPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Ue_CapabilityRAT_RequestList.Decode(r); err != nil {
		return utils.WrapError("Decode Ue_CapabilityRAT_RequestList", err)
	}
	if LateNonCriticalExtensionPresent {
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = &tmp_os_LateNonCriticalExtension
	}
	if Ue_CapabilityEnquiryExtPresent {
		var tmp_os_Ue_CapabilityEnquiryExt []byte
		if tmp_os_Ue_CapabilityEnquiryExt, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode Ue_CapabilityEnquiryExt", err)
		}
		ie.Ue_CapabilityEnquiryExt = &tmp_os_Ue_CapabilityEnquiryExt
	}
	return nil
}
