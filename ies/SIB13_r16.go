package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SIB13_r16 struct {
	Sl_V2X_ConfigCommon_r16  []byte  `madatory`
	Dummy                    []byte  `madatory`
	Tdd_Config_r16           []byte  `madatory`
	LateNonCriticalExtension *[]byte `optional`
}

func (ie *SIB13_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.LateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteOctetString(ie.Sl_V2X_ConfigCommon_r16, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("WriteOctetString Sl_V2X_ConfigCommon_r16", err)
	}
	if err = w.WriteOctetString(ie.Dummy, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("WriteOctetString Dummy", err)
	}
	if err = w.WriteOctetString(ie.Tdd_Config_r16, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("WriteOctetString Tdd_Config_r16", err)
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode LateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *SIB13_r16) Decode(r *aper.AperReader) error {
	var err error
	var LateNonCriticalExtensionPresent bool
	if LateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_os_Sl_V2X_ConfigCommon_r16 []byte
	if tmp_os_Sl_V2X_ConfigCommon_r16, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("ReadOctetString Sl_V2X_ConfigCommon_r16", err)
	}
	ie.Sl_V2X_ConfigCommon_r16 = tmp_os_Sl_V2X_ConfigCommon_r16
	var tmp_os_Dummy []byte
	if tmp_os_Dummy, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("ReadOctetString Dummy", err)
	}
	ie.Dummy = tmp_os_Dummy
	var tmp_os_Tdd_Config_r16 []byte
	if tmp_os_Tdd_Config_r16, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("ReadOctetString Tdd_Config_r16", err)
	}
	ie.Tdd_Config_r16 = tmp_os_Tdd_Config_r16
	if LateNonCriticalExtensionPresent {
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = &tmp_os_LateNonCriticalExtension
	}
	return nil
}
