package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	ServingCellConfig_sCellDeactivationTimer_Enum_ms20   aper.Enumerated = 0
	ServingCellConfig_sCellDeactivationTimer_Enum_ms40   aper.Enumerated = 1
	ServingCellConfig_sCellDeactivationTimer_Enum_ms80   aper.Enumerated = 2
	ServingCellConfig_sCellDeactivationTimer_Enum_ms160  aper.Enumerated = 3
	ServingCellConfig_sCellDeactivationTimer_Enum_ms200  aper.Enumerated = 4
	ServingCellConfig_sCellDeactivationTimer_Enum_ms240  aper.Enumerated = 5
	ServingCellConfig_sCellDeactivationTimer_Enum_ms320  aper.Enumerated = 6
	ServingCellConfig_sCellDeactivationTimer_Enum_ms400  aper.Enumerated = 7
	ServingCellConfig_sCellDeactivationTimer_Enum_ms480  aper.Enumerated = 8
	ServingCellConfig_sCellDeactivationTimer_Enum_ms520  aper.Enumerated = 9
	ServingCellConfig_sCellDeactivationTimer_Enum_ms640  aper.Enumerated = 10
	ServingCellConfig_sCellDeactivationTimer_Enum_ms720  aper.Enumerated = 11
	ServingCellConfig_sCellDeactivationTimer_Enum_ms840  aper.Enumerated = 12
	ServingCellConfig_sCellDeactivationTimer_Enum_ms1280 aper.Enumerated = 13
	ServingCellConfig_sCellDeactivationTimer_Enum_spare2 aper.Enumerated = 14
	ServingCellConfig_sCellDeactivationTimer_Enum_spare1 aper.Enumerated = 15
)

type ServingCellConfig_sCellDeactivationTimer struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *ServingCellConfig_sCellDeactivationTimer) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode ServingCellConfig_sCellDeactivationTimer", err)
	}
	return nil
}

func (ie *ServingCellConfig_sCellDeactivationTimer) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode ServingCellConfig_sCellDeactivationTimer", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
