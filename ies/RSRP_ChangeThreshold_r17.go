package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RSRP_ChangeThreshold_r17_Enum_dB4    aper.Enumerated = 0
	RSRP_ChangeThreshold_r17_Enum_dB6    aper.Enumerated = 1
	RSRP_ChangeThreshold_r17_Enum_dB8    aper.Enumerated = 2
	RSRP_ChangeThreshold_r17_Enum_dB10   aper.Enumerated = 3
	RSRP_ChangeThreshold_r17_Enum_dB14   aper.Enumerated = 4
	RSRP_ChangeThreshold_r17_Enum_dB18   aper.Enumerated = 5
	RSRP_ChangeThreshold_r17_Enum_dB22   aper.Enumerated = 6
	RSRP_ChangeThreshold_r17_Enum_dB26   aper.Enumerated = 7
	RSRP_ChangeThreshold_r17_Enum_dB30   aper.Enumerated = 8
	RSRP_ChangeThreshold_r17_Enum_dB34   aper.Enumerated = 9
	RSRP_ChangeThreshold_r17_Enum_spare6 aper.Enumerated = 10
	RSRP_ChangeThreshold_r17_Enum_spare5 aper.Enumerated = 11
	RSRP_ChangeThreshold_r17_Enum_spare4 aper.Enumerated = 12
	RSRP_ChangeThreshold_r17_Enum_spare3 aper.Enumerated = 13
	RSRP_ChangeThreshold_r17_Enum_spare2 aper.Enumerated = 14
	RSRP_ChangeThreshold_r17_Enum_spare1 aper.Enumerated = 15
)

type RSRP_ChangeThreshold_r17 struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *RSRP_ChangeThreshold_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode RSRP_ChangeThreshold_r17", err)
	}
	return nil
}

func (ie *RSRP_ChangeThreshold_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode RSRP_ChangeThreshold_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
