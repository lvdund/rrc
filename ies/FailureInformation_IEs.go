package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FailureInformation_IEs struct {
	FailureInfoRLC_Bearer    *FailureInfoRLC_Bearer        `optional`
	LateNonCriticalExtension *[]byte                       `optional`
	NonCriticalExtension     *FailureInformation_v1610_IEs `optional`
}

func (ie *FailureInformation_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.FailureInfoRLC_Bearer != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.FailureInfoRLC_Bearer != nil {
		if err = ie.FailureInfoRLC_Bearer.Encode(w); err != nil {
			return utils.WrapError("Encode FailureInfoRLC_Bearer", err)
		}
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
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

func (ie *FailureInformation_IEs) Decode(r *uper.UperReader) error {
	var err error
	var FailureInfoRLC_BearerPresent bool
	if FailureInfoRLC_BearerPresent, err = r.ReadBool(); err != nil {
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
	if FailureInfoRLC_BearerPresent {
		ie.FailureInfoRLC_Bearer = new(FailureInfoRLC_Bearer)
		if err = ie.FailureInfoRLC_Bearer.Decode(r); err != nil {
			return utils.WrapError("Decode FailureInfoRLC_Bearer", err)
		}
	}
	if LateNonCriticalExtensionPresent {
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = &tmp_os_LateNonCriticalExtension
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(FailureInformation_v1610_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
