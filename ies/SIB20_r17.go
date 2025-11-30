package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SIB20_r17 struct {
	Mcch_Config_r17          MCCH_Config_r17          `madatory`
	Cfr_ConfigMCCH_MTCH_r17  *CFR_ConfigMCCH_MTCH_r17 `optional`
	LateNonCriticalExtension *[]byte                  `optional`
}

func (ie *SIB20_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Cfr_ConfigMCCH_MTCH_r17 != nil, ie.LateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Mcch_Config_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Mcch_Config_r17", err)
	}
	if ie.Cfr_ConfigMCCH_MTCH_r17 != nil {
		if err = ie.Cfr_ConfigMCCH_MTCH_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Cfr_ConfigMCCH_MTCH_r17", err)
		}
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode LateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *SIB20_r17) Decode(r *aper.AperReader) error {
	var err error
	var Cfr_ConfigMCCH_MTCH_r17Present bool
	if Cfr_ConfigMCCH_MTCH_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var LateNonCriticalExtensionPresent bool
	if LateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Mcch_Config_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Mcch_Config_r17", err)
	}
	if Cfr_ConfigMCCH_MTCH_r17Present {
		ie.Cfr_ConfigMCCH_MTCH_r17 = new(CFR_ConfigMCCH_MTCH_r17)
		if err = ie.Cfr_ConfigMCCH_MTCH_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Cfr_ConfigMCCH_MTCH_r17", err)
		}
	}
	if LateNonCriticalExtensionPresent {
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = &tmp_os_LateNonCriticalExtension
	}
	return nil
}
