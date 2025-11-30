package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_PathSwitchConfig_r17_t420_r17_Enum_ms50    aper.Enumerated = 0
	SL_PathSwitchConfig_r17_t420_r17_Enum_ms100   aper.Enumerated = 1
	SL_PathSwitchConfig_r17_t420_r17_Enum_ms150   aper.Enumerated = 2
	SL_PathSwitchConfig_r17_t420_r17_Enum_ms200   aper.Enumerated = 3
	SL_PathSwitchConfig_r17_t420_r17_Enum_ms500   aper.Enumerated = 4
	SL_PathSwitchConfig_r17_t420_r17_Enum_ms1000  aper.Enumerated = 5
	SL_PathSwitchConfig_r17_t420_r17_Enum_ms2000  aper.Enumerated = 6
	SL_PathSwitchConfig_r17_t420_r17_Enum_ms10000 aper.Enumerated = 7
)

type SL_PathSwitchConfig_r17_t420_r17 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *SL_PathSwitchConfig_r17_t420_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SL_PathSwitchConfig_r17_t420_r17", err)
	}
	return nil
}

func (ie *SL_PathSwitchConfig_r17_t420_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SL_PathSwitchConfig_r17_t420_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
