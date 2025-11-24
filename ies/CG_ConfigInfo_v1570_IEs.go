package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_ConfigInfo_v1570_IEs struct {
	SftdFrequencyList_NR    *SFTD_FrequencyList_NR    `optional`
	SftdFrequencyList_EUTRA *SFTD_FrequencyList_EUTRA `optional`
	NonCriticalExtension    *CG_ConfigInfo_v1590_IEs  `optional`
}

func (ie *CG_ConfigInfo_v1570_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.SftdFrequencyList_NR != nil, ie.SftdFrequencyList_EUTRA != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SftdFrequencyList_NR != nil {
		if err = ie.SftdFrequencyList_NR.Encode(w); err != nil {
			return utils.WrapError("Encode SftdFrequencyList_NR", err)
		}
	}
	if ie.SftdFrequencyList_EUTRA != nil {
		if err = ie.SftdFrequencyList_EUTRA.Encode(w); err != nil {
			return utils.WrapError("Encode SftdFrequencyList_EUTRA", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *CG_ConfigInfo_v1570_IEs) Decode(r *uper.UperReader) error {
	var err error
	var SftdFrequencyList_NRPresent bool
	if SftdFrequencyList_NRPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SftdFrequencyList_EUTRAPresent bool
	if SftdFrequencyList_EUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if SftdFrequencyList_NRPresent {
		ie.SftdFrequencyList_NR = new(SFTD_FrequencyList_NR)
		if err = ie.SftdFrequencyList_NR.Decode(r); err != nil {
			return utils.WrapError("Decode SftdFrequencyList_NR", err)
		}
	}
	if SftdFrequencyList_EUTRAPresent {
		ie.SftdFrequencyList_EUTRA = new(SFTD_FrequencyList_EUTRA)
		if err = ie.SftdFrequencyList_EUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode SftdFrequencyList_EUTRA", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(CG_ConfigInfo_v1590_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
