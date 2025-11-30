package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SIB14_r16 struct {
	Sl_V2X_ConfigCommonExt_r16 []byte  `madatory`
	LateNonCriticalExtension   *[]byte `optional`
}

func (ie *SIB14_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.LateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteOctetString(ie.Sl_V2X_ConfigCommonExt_r16, nil, false); err != nil {
		return utils.WrapError("WriteOctetString Sl_V2X_ConfigCommonExt_r16", err)
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, nil, false); err != nil {
			return utils.WrapError("Encode LateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *SIB14_r16) Decode(r *aper.AperReader) error {
	var err error
	var LateNonCriticalExtensionPresent bool
	if LateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_os_Sl_V2X_ConfigCommonExt_r16 []byte
	if tmp_os_Sl_V2X_ConfigCommonExt_r16, err = r.ReadOctetString(nil, false); err != nil {
		return utils.WrapError("ReadOctetString Sl_V2X_ConfigCommonExt_r16", err)
	}
	ie.Sl_V2X_ConfigCommonExt_r16 = tmp_os_Sl_V2X_ConfigCommonExt_r16
	if LateNonCriticalExtensionPresent {
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = &tmp_os_LateNonCriticalExtension
	}
	return nil
}
