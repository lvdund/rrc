package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	EUTRA_ParametersCommon_multiNS_Pmax_EUTRA_Enum_supported aper.Enumerated = 0
)

type EUTRA_ParametersCommon_multiNS_Pmax_EUTRA struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *EUTRA_ParametersCommon_multiNS_Pmax_EUTRA) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode EUTRA_ParametersCommon_multiNS_Pmax_EUTRA", err)
	}
	return nil
}

func (ie *EUTRA_ParametersCommon_multiNS_Pmax_EUTRA) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode EUTRA_ParametersCommon_multiNS_Pmax_EUTRA", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
