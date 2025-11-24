package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DedicatedSIBRequest_r16_IEs struct {
	OnDemandSIB_RequestList_r16 *DedicatedSIBRequest_r16_IEs_onDemandSIB_RequestList_r16 `lb:1,ub:maxOnDemandSIB_r16,optional`
	LateNonCriticalExtension    *[]byte                                                  `optional`
	NonCriticalExtension        interface{}                                              `optional`
}

func (ie *DedicatedSIBRequest_r16_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.OnDemandSIB_RequestList_r16 != nil, ie.LateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.OnDemandSIB_RequestList_r16 != nil {
		if err = ie.OnDemandSIB_RequestList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode OnDemandSIB_RequestList_r16", err)
		}
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode LateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *DedicatedSIBRequest_r16_IEs) Decode(r *uper.UperReader) error {
	var err error
	var OnDemandSIB_RequestList_r16Present bool
	if OnDemandSIB_RequestList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var LateNonCriticalExtensionPresent bool
	if LateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if OnDemandSIB_RequestList_r16Present {
		ie.OnDemandSIB_RequestList_r16 = new(DedicatedSIBRequest_r16_IEs_onDemandSIB_RequestList_r16)
		if err = ie.OnDemandSIB_RequestList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode OnDemandSIB_RequestList_r16", err)
		}
	}
	if LateNonCriticalExtensionPresent {
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = &tmp_os_LateNonCriticalExtension
	}
	return nil
}
