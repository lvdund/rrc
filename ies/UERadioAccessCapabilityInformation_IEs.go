package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UERadioAccessCapabilityInformation_IEs struct {
	Ue_RadioAccessCapabilityInfo []byte      `madatory`
	NonCriticalExtension         interface{} `optional`
}

func (ie *UERadioAccessCapabilityInformation_IEs) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteOctetString(ie.Ue_RadioAccessCapabilityInfo, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("WriteOctetString Ue_RadioAccessCapabilityInfo", err)
	}
	return nil
}

func (ie *UERadioAccessCapabilityInformation_IEs) Decode(r *uper.UperReader) error {
	var err error
	var tmp_os_Ue_RadioAccessCapabilityInfo []byte
	if tmp_os_Ue_RadioAccessCapabilityInfo, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("ReadOctetString Ue_RadioAccessCapabilityInfo", err)
	}
	ie.Ue_RadioAccessCapabilityInfo = tmp_os_Ue_RadioAccessCapabilityInfo
	return nil
}
