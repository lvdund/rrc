package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	ServingCellConfigCommon_ssb_periodicityServingCell_Enum_ms5    aper.Enumerated = 0
	ServingCellConfigCommon_ssb_periodicityServingCell_Enum_ms10   aper.Enumerated = 1
	ServingCellConfigCommon_ssb_periodicityServingCell_Enum_ms20   aper.Enumerated = 2
	ServingCellConfigCommon_ssb_periodicityServingCell_Enum_ms40   aper.Enumerated = 3
	ServingCellConfigCommon_ssb_periodicityServingCell_Enum_ms80   aper.Enumerated = 4
	ServingCellConfigCommon_ssb_periodicityServingCell_Enum_ms160  aper.Enumerated = 5
	ServingCellConfigCommon_ssb_periodicityServingCell_Enum_spare2 aper.Enumerated = 6
	ServingCellConfigCommon_ssb_periodicityServingCell_Enum_spare1 aper.Enumerated = 7
)

type ServingCellConfigCommon_ssb_periodicityServingCell struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *ServingCellConfigCommon_ssb_periodicityServingCell) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode ServingCellConfigCommon_ssb_periodicityServingCell", err)
	}
	return nil
}

func (ie *ServingCellConfigCommon_ssb_periodicityServingCell) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode ServingCellConfigCommon_ssb_periodicityServingCell", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
