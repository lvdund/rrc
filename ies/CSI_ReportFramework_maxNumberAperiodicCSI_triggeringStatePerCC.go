package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_ReportFramework_maxNumberAperiodicCSI_triggeringStatePerCC_Enum_n3   aper.Enumerated = 0
	CSI_ReportFramework_maxNumberAperiodicCSI_triggeringStatePerCC_Enum_n7   aper.Enumerated = 1
	CSI_ReportFramework_maxNumberAperiodicCSI_triggeringStatePerCC_Enum_n15  aper.Enumerated = 2
	CSI_ReportFramework_maxNumberAperiodicCSI_triggeringStatePerCC_Enum_n31  aper.Enumerated = 3
	CSI_ReportFramework_maxNumberAperiodicCSI_triggeringStatePerCC_Enum_n63  aper.Enumerated = 4
	CSI_ReportFramework_maxNumberAperiodicCSI_triggeringStatePerCC_Enum_n128 aper.Enumerated = 5
)

type CSI_ReportFramework_maxNumberAperiodicCSI_triggeringStatePerCC struct {
	Value aper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *CSI_ReportFramework_maxNumberAperiodicCSI_triggeringStatePerCC) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode CSI_ReportFramework_maxNumberAperiodicCSI_triggeringStatePerCC", err)
	}
	return nil
}

func (ie *CSI_ReportFramework_maxNumberAperiodicCSI_triggeringStatePerCC) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode CSI_ReportFramework_maxNumberAperiodicCSI_triggeringStatePerCC", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
