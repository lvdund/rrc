package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SPS_Config_periodicity_Enum_ms10   aper.Enumerated = 0
	SPS_Config_periodicity_Enum_ms20   aper.Enumerated = 1
	SPS_Config_periodicity_Enum_ms32   aper.Enumerated = 2
	SPS_Config_periodicity_Enum_ms40   aper.Enumerated = 3
	SPS_Config_periodicity_Enum_ms64   aper.Enumerated = 4
	SPS_Config_periodicity_Enum_ms80   aper.Enumerated = 5
	SPS_Config_periodicity_Enum_ms128  aper.Enumerated = 6
	SPS_Config_periodicity_Enum_ms160  aper.Enumerated = 7
	SPS_Config_periodicity_Enum_ms320  aper.Enumerated = 8
	SPS_Config_periodicity_Enum_ms640  aper.Enumerated = 9
	SPS_Config_periodicity_Enum_spare6 aper.Enumerated = 10
	SPS_Config_periodicity_Enum_spare5 aper.Enumerated = 11
	SPS_Config_periodicity_Enum_spare4 aper.Enumerated = 12
	SPS_Config_periodicity_Enum_spare3 aper.Enumerated = 13
	SPS_Config_periodicity_Enum_spare2 aper.Enumerated = 14
	SPS_Config_periodicity_Enum_spare1 aper.Enumerated = 15
)

type SPS_Config_periodicity struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *SPS_Config_periodicity) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode SPS_Config_periodicity", err)
	}
	return nil
}

func (ie *SPS_Config_periodicity) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode SPS_Config_periodicity", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
