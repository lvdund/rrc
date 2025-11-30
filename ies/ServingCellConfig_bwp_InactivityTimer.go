package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	ServingCellConfig_bwp_InactivityTimer_Enum_ms2     aper.Enumerated = 0
	ServingCellConfig_bwp_InactivityTimer_Enum_ms3     aper.Enumerated = 1
	ServingCellConfig_bwp_InactivityTimer_Enum_ms4     aper.Enumerated = 2
	ServingCellConfig_bwp_InactivityTimer_Enum_ms5     aper.Enumerated = 3
	ServingCellConfig_bwp_InactivityTimer_Enum_ms6     aper.Enumerated = 4
	ServingCellConfig_bwp_InactivityTimer_Enum_ms8     aper.Enumerated = 5
	ServingCellConfig_bwp_InactivityTimer_Enum_ms10    aper.Enumerated = 6
	ServingCellConfig_bwp_InactivityTimer_Enum_ms20    aper.Enumerated = 7
	ServingCellConfig_bwp_InactivityTimer_Enum_ms30    aper.Enumerated = 8
	ServingCellConfig_bwp_InactivityTimer_Enum_ms40    aper.Enumerated = 9
	ServingCellConfig_bwp_InactivityTimer_Enum_ms50    aper.Enumerated = 10
	ServingCellConfig_bwp_InactivityTimer_Enum_ms60    aper.Enumerated = 11
	ServingCellConfig_bwp_InactivityTimer_Enum_ms80    aper.Enumerated = 12
	ServingCellConfig_bwp_InactivityTimer_Enum_ms100   aper.Enumerated = 13
	ServingCellConfig_bwp_InactivityTimer_Enum_ms200   aper.Enumerated = 14
	ServingCellConfig_bwp_InactivityTimer_Enum_ms300   aper.Enumerated = 15
	ServingCellConfig_bwp_InactivityTimer_Enum_ms500   aper.Enumerated = 16
	ServingCellConfig_bwp_InactivityTimer_Enum_ms750   aper.Enumerated = 17
	ServingCellConfig_bwp_InactivityTimer_Enum_ms1280  aper.Enumerated = 18
	ServingCellConfig_bwp_InactivityTimer_Enum_ms1920  aper.Enumerated = 19
	ServingCellConfig_bwp_InactivityTimer_Enum_ms2560  aper.Enumerated = 20
	ServingCellConfig_bwp_InactivityTimer_Enum_spare10 aper.Enumerated = 21
	ServingCellConfig_bwp_InactivityTimer_Enum_spare9  aper.Enumerated = 22
	ServingCellConfig_bwp_InactivityTimer_Enum_spare8  aper.Enumerated = 23
	ServingCellConfig_bwp_InactivityTimer_Enum_spare7  aper.Enumerated = 24
	ServingCellConfig_bwp_InactivityTimer_Enum_spare6  aper.Enumerated = 25
	ServingCellConfig_bwp_InactivityTimer_Enum_spare5  aper.Enumerated = 26
	ServingCellConfig_bwp_InactivityTimer_Enum_spare4  aper.Enumerated = 27
	ServingCellConfig_bwp_InactivityTimer_Enum_spare3  aper.Enumerated = 28
	ServingCellConfig_bwp_InactivityTimer_Enum_spare2  aper.Enumerated = 29
	ServingCellConfig_bwp_InactivityTimer_Enum_spare1  aper.Enumerated = 30
)

type ServingCellConfig_bwp_InactivityTimer struct {
	Value aper.Enumerated `lb:0,ub:30,madatory`
}

func (ie *ServingCellConfig_bwp_InactivityTimer) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 30}, false); err != nil {
		return utils.WrapError("Encode ServingCellConfig_bwp_InactivityTimer", err)
	}
	return nil
}

func (ie *ServingCellConfig_bwp_InactivityTimer) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 30}, false); err != nil {
		return utils.WrapError("Decode ServingCellConfig_bwp_InactivityTimer", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
