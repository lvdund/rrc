package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	NTN_Config_r17_ntn_UlSyncValidityDuration_r17_Enum_s5   aper.Enumerated = 0
	NTN_Config_r17_ntn_UlSyncValidityDuration_r17_Enum_s10  aper.Enumerated = 1
	NTN_Config_r17_ntn_UlSyncValidityDuration_r17_Enum_s15  aper.Enumerated = 2
	NTN_Config_r17_ntn_UlSyncValidityDuration_r17_Enum_s20  aper.Enumerated = 3
	NTN_Config_r17_ntn_UlSyncValidityDuration_r17_Enum_s25  aper.Enumerated = 4
	NTN_Config_r17_ntn_UlSyncValidityDuration_r17_Enum_s30  aper.Enumerated = 5
	NTN_Config_r17_ntn_UlSyncValidityDuration_r17_Enum_s35  aper.Enumerated = 6
	NTN_Config_r17_ntn_UlSyncValidityDuration_r17_Enum_s40  aper.Enumerated = 7
	NTN_Config_r17_ntn_UlSyncValidityDuration_r17_Enum_s45  aper.Enumerated = 8
	NTN_Config_r17_ntn_UlSyncValidityDuration_r17_Enum_s50  aper.Enumerated = 9
	NTN_Config_r17_ntn_UlSyncValidityDuration_r17_Enum_s55  aper.Enumerated = 10
	NTN_Config_r17_ntn_UlSyncValidityDuration_r17_Enum_s60  aper.Enumerated = 11
	NTN_Config_r17_ntn_UlSyncValidityDuration_r17_Enum_s120 aper.Enumerated = 12
	NTN_Config_r17_ntn_UlSyncValidityDuration_r17_Enum_s180 aper.Enumerated = 13
	NTN_Config_r17_ntn_UlSyncValidityDuration_r17_Enum_s240 aper.Enumerated = 14
	NTN_Config_r17_ntn_UlSyncValidityDuration_r17_Enum_s900 aper.Enumerated = 15
)

type NTN_Config_r17_ntn_UlSyncValidityDuration_r17 struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *NTN_Config_r17_ntn_UlSyncValidityDuration_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode NTN_Config_r17_ntn_UlSyncValidityDuration_r17", err)
	}
	return nil
}

func (ie *NTN_Config_r17_ntn_UlSyncValidityDuration_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode NTN_Config_r17_ntn_UlSyncValidityDuration_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
