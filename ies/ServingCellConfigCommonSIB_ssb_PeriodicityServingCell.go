package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	ServingCellConfigCommonSIB_ssb_PeriodicityServingCell_Enum_ms5   aper.Enumerated = 0
	ServingCellConfigCommonSIB_ssb_PeriodicityServingCell_Enum_ms10  aper.Enumerated = 1
	ServingCellConfigCommonSIB_ssb_PeriodicityServingCell_Enum_ms20  aper.Enumerated = 2
	ServingCellConfigCommonSIB_ssb_PeriodicityServingCell_Enum_ms40  aper.Enumerated = 3
	ServingCellConfigCommonSIB_ssb_PeriodicityServingCell_Enum_ms80  aper.Enumerated = 4
	ServingCellConfigCommonSIB_ssb_PeriodicityServingCell_Enum_ms160 aper.Enumerated = 5
)

type ServingCellConfigCommonSIB_ssb_PeriodicityServingCell struct {
	Value aper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *ServingCellConfigCommonSIB_ssb_PeriodicityServingCell) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode ServingCellConfigCommonSIB_ssb_PeriodicityServingCell", err)
	}
	return nil
}

func (ie *ServingCellConfigCommonSIB_ssb_PeriodicityServingCell) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode ServingCellConfigCommonSIB_ssb_PeriodicityServingCell", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
