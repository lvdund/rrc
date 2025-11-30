package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	WLAN_RTT_r16_rttUnits_r16_Enum_microseconds          aper.Enumerated = 0
	WLAN_RTT_r16_rttUnits_r16_Enum_hundredsofnanoseconds aper.Enumerated = 1
	WLAN_RTT_r16_rttUnits_r16_Enum_tensofnanoseconds     aper.Enumerated = 2
	WLAN_RTT_r16_rttUnits_r16_Enum_nanoseconds           aper.Enumerated = 3
	WLAN_RTT_r16_rttUnits_r16_Enum_tenthsofnanoseconds   aper.Enumerated = 4
)

type WLAN_RTT_r16_rttUnits_r16 struct {
	Value aper.Enumerated `lb:0,ub:4,madatory`
}

func (ie *WLAN_RTT_r16_rttUnits_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Encode WLAN_RTT_r16_rttUnits_r16", err)
	}
	return nil
}

func (ie *WLAN_RTT_r16_rttUnits_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Decode WLAN_RTT_r16_rttUnits_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
