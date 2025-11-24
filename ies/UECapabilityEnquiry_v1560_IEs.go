package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UECapabilityEnquiry_v1560_IEs struct {
	CapabilityRequestFilterCommon *UE_CapabilityRequestFilterCommon `optional`
	NonCriticalExtension          *UECapabilityEnquiry_v1610_IEs    `optional`
}

func (ie *UECapabilityEnquiry_v1560_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.CapabilityRequestFilterCommon != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.CapabilityRequestFilterCommon != nil {
		if err = ie.CapabilityRequestFilterCommon.Encode(w); err != nil {
			return utils.WrapError("Encode CapabilityRequestFilterCommon", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UECapabilityEnquiry_v1560_IEs) Decode(r *uper.UperReader) error {
	var err error
	var CapabilityRequestFilterCommonPresent bool
	if CapabilityRequestFilterCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if CapabilityRequestFilterCommonPresent {
		ie.CapabilityRequestFilterCommon = new(UE_CapabilityRequestFilterCommon)
		if err = ie.CapabilityRequestFilterCommon.Decode(r); err != nil {
			return utils.WrapError("Decode CapabilityRequestFilterCommon", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(UECapabilityEnquiry_v1610_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
