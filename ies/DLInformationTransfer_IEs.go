package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type DLInformationTransfer_IEs struct {
	DedicatedNAS_Message     *DedicatedNAS_Message            `optional`
	LateNonCriticalExtension *[]byte                          `optional`
	NonCriticalExtension     *DLInformationTransfer_v1610_IEs `optional`
}

func (ie *DLInformationTransfer_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.DedicatedNAS_Message != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.DedicatedNAS_Message != nil {
		if err = ie.DedicatedNAS_Message.Encode(w); err != nil {
			return utils.WrapError("Encode DedicatedNAS_Message", err)
		}
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
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

func (ie *DLInformationTransfer_IEs) Decode(r *aper.AperReader) error {
	var err error
	var DedicatedNAS_MessagePresent bool
	if DedicatedNAS_MessagePresent, err = r.ReadBool(); err != nil {
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
	if DedicatedNAS_MessagePresent {
		ie.DedicatedNAS_Message = new(DedicatedNAS_Message)
		if err = ie.DedicatedNAS_Message.Decode(r); err != nil {
			return utils.WrapError("Decode DedicatedNAS_Message", err)
		}
	}
	if LateNonCriticalExtensionPresent {
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = &tmp_os_LateNonCriticalExtension
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(DLInformationTransfer_v1610_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
