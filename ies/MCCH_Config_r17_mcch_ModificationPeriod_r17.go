package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MCCH_Config_r17_mcch_ModificationPeriod_r17_Enum_rf2     aper.Enumerated = 0
	MCCH_Config_r17_mcch_ModificationPeriod_r17_Enum_rf4     aper.Enumerated = 1
	MCCH_Config_r17_mcch_ModificationPeriod_r17_Enum_rf8     aper.Enumerated = 2
	MCCH_Config_r17_mcch_ModificationPeriod_r17_Enum_rf16    aper.Enumerated = 3
	MCCH_Config_r17_mcch_ModificationPeriod_r17_Enum_rf32    aper.Enumerated = 4
	MCCH_Config_r17_mcch_ModificationPeriod_r17_Enum_rf64    aper.Enumerated = 5
	MCCH_Config_r17_mcch_ModificationPeriod_r17_Enum_rf128   aper.Enumerated = 6
	MCCH_Config_r17_mcch_ModificationPeriod_r17_Enum_rf256   aper.Enumerated = 7
	MCCH_Config_r17_mcch_ModificationPeriod_r17_Enum_rf512   aper.Enumerated = 8
	MCCH_Config_r17_mcch_ModificationPeriod_r17_Enum_rf1024  aper.Enumerated = 9
	MCCH_Config_r17_mcch_ModificationPeriod_r17_Enum_r2048   aper.Enumerated = 10
	MCCH_Config_r17_mcch_ModificationPeriod_r17_Enum_rf4096  aper.Enumerated = 11
	MCCH_Config_r17_mcch_ModificationPeriod_r17_Enum_rf8192  aper.Enumerated = 12
	MCCH_Config_r17_mcch_ModificationPeriod_r17_Enum_rf16384 aper.Enumerated = 13
	MCCH_Config_r17_mcch_ModificationPeriod_r17_Enum_rf32768 aper.Enumerated = 14
	MCCH_Config_r17_mcch_ModificationPeriod_r17_Enum_rf65536 aper.Enumerated = 15
)

type MCCH_Config_r17_mcch_ModificationPeriod_r17 struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *MCCH_Config_r17_mcch_ModificationPeriod_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode MCCH_Config_r17_mcch_ModificationPeriod_r17", err)
	}
	return nil
}

func (ie *MCCH_Config_r17_mcch_ModificationPeriod_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode MCCH_Config_r17_mcch_ModificationPeriod_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
