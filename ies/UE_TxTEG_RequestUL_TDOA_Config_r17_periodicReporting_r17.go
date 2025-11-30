package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	UE_TxTEG_RequestUL_TDOA_Config_r17_periodicReporting_r17_Enum_ms160    aper.Enumerated = 0
	UE_TxTEG_RequestUL_TDOA_Config_r17_periodicReporting_r17_Enum_ms320    aper.Enumerated = 1
	UE_TxTEG_RequestUL_TDOA_Config_r17_periodicReporting_r17_Enum_ms1280   aper.Enumerated = 2
	UE_TxTEG_RequestUL_TDOA_Config_r17_periodicReporting_r17_Enum_ms2560   aper.Enumerated = 3
	UE_TxTEG_RequestUL_TDOA_Config_r17_periodicReporting_r17_Enum_ms61440  aper.Enumerated = 4
	UE_TxTEG_RequestUL_TDOA_Config_r17_periodicReporting_r17_Enum_ms81920  aper.Enumerated = 5
	UE_TxTEG_RequestUL_TDOA_Config_r17_periodicReporting_r17_Enum_ms368640 aper.Enumerated = 6
	UE_TxTEG_RequestUL_TDOA_Config_r17_periodicReporting_r17_Enum_ms737280 aper.Enumerated = 7
)

type UE_TxTEG_RequestUL_TDOA_Config_r17_periodicReporting_r17 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *UE_TxTEG_RequestUL_TDOA_Config_r17_periodicReporting_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode UE_TxTEG_RequestUL_TDOA_Config_r17_periodicReporting_r17", err)
	}
	return nil
}

func (ie *UE_TxTEG_RequestUL_TDOA_Config_r17_periodicReporting_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode UE_TxTEG_RequestUL_TDOA_Config_r17_periodicReporting_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
