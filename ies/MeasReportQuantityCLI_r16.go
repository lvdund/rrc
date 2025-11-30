package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasReportQuantityCLI_r16_Enum_srs_rsrp aper.Enumerated = 0
	MeasReportQuantityCLI_r16_Enum_cli_rssi aper.Enumerated = 1
)

type MeasReportQuantityCLI_r16 struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *MeasReportQuantityCLI_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode MeasReportQuantityCLI_r16", err)
	}
	return nil
}

func (ie *MeasReportQuantityCLI_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode MeasReportQuantityCLI_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
