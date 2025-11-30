package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasAndMobParametersCommon_nr_CGI_Reporting_NEDC_Enum_supported aper.Enumerated = 0
)

type MeasAndMobParametersCommon_nr_CGI_Reporting_NEDC struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *MeasAndMobParametersCommon_nr_CGI_Reporting_NEDC) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode MeasAndMobParametersCommon_nr_CGI_Reporting_NEDC", err)
	}
	return nil
}

func (ie *MeasAndMobParametersCommon_nr_CGI_Reporting_NEDC) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode MeasAndMobParametersCommon_nr_CGI_Reporting_NEDC", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
