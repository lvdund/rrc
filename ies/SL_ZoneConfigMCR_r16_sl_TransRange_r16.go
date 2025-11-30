package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m20    aper.Enumerated = 0
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m50    aper.Enumerated = 1
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m80    aper.Enumerated = 2
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m100   aper.Enumerated = 3
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m120   aper.Enumerated = 4
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m150   aper.Enumerated = 5
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m180   aper.Enumerated = 6
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m200   aper.Enumerated = 7
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m220   aper.Enumerated = 8
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m250   aper.Enumerated = 9
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m270   aper.Enumerated = 10
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m300   aper.Enumerated = 11
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m350   aper.Enumerated = 12
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m370   aper.Enumerated = 13
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m400   aper.Enumerated = 14
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m420   aper.Enumerated = 15
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m450   aper.Enumerated = 16
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m480   aper.Enumerated = 17
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m500   aper.Enumerated = 18
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m550   aper.Enumerated = 19
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m600   aper.Enumerated = 20
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m700   aper.Enumerated = 21
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m1000  aper.Enumerated = 22
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_spare9 aper.Enumerated = 23
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_spare8 aper.Enumerated = 24
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_spare7 aper.Enumerated = 25
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_spare6 aper.Enumerated = 26
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_spare5 aper.Enumerated = 27
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_spare4 aper.Enumerated = 28
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_spare3 aper.Enumerated = 29
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_spare2 aper.Enumerated = 30
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_spare1 aper.Enumerated = 31
)

type SL_ZoneConfigMCR_r16_sl_TransRange_r16 struct {
	Value aper.Enumerated `lb:0,ub:31,madatory`
}

func (ie *SL_ZoneConfigMCR_r16_sl_TransRange_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("Encode SL_ZoneConfigMCR_r16_sl_TransRange_r16", err)
	}
	return nil
}

func (ie *SL_ZoneConfigMCR_r16_sl_TransRange_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("Decode SL_ZoneConfigMCR_r16_sl_TransRange_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
