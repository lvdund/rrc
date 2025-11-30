package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MobilityFromNRCommand_IEs struct {
	TargetRAT_Type             MobilityFromNRCommand_IEs_targetRAT_Type `madatory`
	TargetRAT_MessageContainer []byte                                   `madatory`
	Nas_SecurityParamFromNR    *[]byte                                  `optional`
	LateNonCriticalExtension   *[]byte                                  `optional`
	NonCriticalExtension       *MobilityFromNRCommand_v1610_IEs         `optional`
}

func (ie *MobilityFromNRCommand_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Nas_SecurityParamFromNR != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.TargetRAT_Type.Encode(w); err != nil {
		return utils.WrapError("Encode TargetRAT_Type", err)
	}
	if err = w.WriteOctetString(ie.TargetRAT_MessageContainer, nil, false); err != nil {
		return utils.WrapError("WriteOctetString TargetRAT_MessageContainer", err)
	}
	if ie.Nas_SecurityParamFromNR != nil {
		if err = w.WriteOctetString(*ie.Nas_SecurityParamFromNR, nil, false); err != nil {
			return utils.WrapError("Encode Nas_SecurityParamFromNR", err)
		}
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, nil, false); err != nil {
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

func (ie *MobilityFromNRCommand_IEs) Decode(r *aper.AperReader) error {
	var err error
	var Nas_SecurityParamFromNRPresent bool
	if Nas_SecurityParamFromNRPresent, err = r.ReadBool(); err != nil {
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
	if err = ie.TargetRAT_Type.Decode(r); err != nil {
		return utils.WrapError("Decode TargetRAT_Type", err)
	}
	var tmp_os_TargetRAT_MessageContainer []byte
	if tmp_os_TargetRAT_MessageContainer, err = r.ReadOctetString(nil, false); err != nil {
		return utils.WrapError("ReadOctetString TargetRAT_MessageContainer", err)
	}
	ie.TargetRAT_MessageContainer = tmp_os_TargetRAT_MessageContainer
	if Nas_SecurityParamFromNRPresent {
		var tmp_os_Nas_SecurityParamFromNR []byte
		if tmp_os_Nas_SecurityParamFromNR, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode Nas_SecurityParamFromNR", err)
		}
		ie.Nas_SecurityParamFromNR = &tmp_os_Nas_SecurityParamFromNR
	}
	if LateNonCriticalExtensionPresent {
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = &tmp_os_LateNonCriticalExtension
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(MobilityFromNRCommand_v1610_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
