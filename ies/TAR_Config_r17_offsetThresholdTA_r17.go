package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	TAR_Config_r17_offsetThresholdTA_r17_Enum_ms0dot5 aper.Enumerated = 0
	TAR_Config_r17_offsetThresholdTA_r17_Enum_ms1     aper.Enumerated = 1
	TAR_Config_r17_offsetThresholdTA_r17_Enum_ms2     aper.Enumerated = 2
	TAR_Config_r17_offsetThresholdTA_r17_Enum_ms3     aper.Enumerated = 3
	TAR_Config_r17_offsetThresholdTA_r17_Enum_ms4     aper.Enumerated = 4
	TAR_Config_r17_offsetThresholdTA_r17_Enum_ms5     aper.Enumerated = 5
	TAR_Config_r17_offsetThresholdTA_r17_Enum_ms6     aper.Enumerated = 6
	TAR_Config_r17_offsetThresholdTA_r17_Enum_ms7     aper.Enumerated = 7
	TAR_Config_r17_offsetThresholdTA_r17_Enum_ms8     aper.Enumerated = 8
	TAR_Config_r17_offsetThresholdTA_r17_Enum_ms9     aper.Enumerated = 9
	TAR_Config_r17_offsetThresholdTA_r17_Enum_ms10    aper.Enumerated = 10
	TAR_Config_r17_offsetThresholdTA_r17_Enum_ms11    aper.Enumerated = 11
	TAR_Config_r17_offsetThresholdTA_r17_Enum_ms12    aper.Enumerated = 12
	TAR_Config_r17_offsetThresholdTA_r17_Enum_ms13    aper.Enumerated = 13
	TAR_Config_r17_offsetThresholdTA_r17_Enum_ms14    aper.Enumerated = 14
	TAR_Config_r17_offsetThresholdTA_r17_Enum_ms15    aper.Enumerated = 15
	TAR_Config_r17_offsetThresholdTA_r17_Enum_spare13 aper.Enumerated = 16
	TAR_Config_r17_offsetThresholdTA_r17_Enum_spare12 aper.Enumerated = 17
	TAR_Config_r17_offsetThresholdTA_r17_Enum_spare11 aper.Enumerated = 18
	TAR_Config_r17_offsetThresholdTA_r17_Enum_spare10 aper.Enumerated = 19
	TAR_Config_r17_offsetThresholdTA_r17_Enum_spare9  aper.Enumerated = 20
	TAR_Config_r17_offsetThresholdTA_r17_Enum_spare8  aper.Enumerated = 21
	TAR_Config_r17_offsetThresholdTA_r17_Enum_spare7  aper.Enumerated = 22
	TAR_Config_r17_offsetThresholdTA_r17_Enum_spare6  aper.Enumerated = 23
	TAR_Config_r17_offsetThresholdTA_r17_Enum_spare5  aper.Enumerated = 24
	TAR_Config_r17_offsetThresholdTA_r17_Enum_spare4  aper.Enumerated = 25
	TAR_Config_r17_offsetThresholdTA_r17_Enum_spare3  aper.Enumerated = 26
	TAR_Config_r17_offsetThresholdTA_r17_Enum_spare2  aper.Enumerated = 27
	TAR_Config_r17_offsetThresholdTA_r17_Enum_spare1  aper.Enumerated = 28
)

type TAR_Config_r17_offsetThresholdTA_r17 struct {
	Value aper.Enumerated `lb:0,ub:28,madatory`
}

func (ie *TAR_Config_r17_offsetThresholdTA_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 28}, false); err != nil {
		return utils.WrapError("Encode TAR_Config_r17_offsetThresholdTA_r17", err)
	}
	return nil
}

func (ie *TAR_Config_r17_offsetThresholdTA_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 28}, false); err != nil {
		return utils.WrapError("Decode TAR_Config_r17_offsetThresholdTA_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
