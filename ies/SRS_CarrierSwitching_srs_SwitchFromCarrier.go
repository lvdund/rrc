package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_CarrierSwitching_srs_SwitchFromCarrier_Enum_sUL aper.Enumerated = 0
	SRS_CarrierSwitching_srs_SwitchFromCarrier_Enum_nUL aper.Enumerated = 1
)

type SRS_CarrierSwitching_srs_SwitchFromCarrier struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *SRS_CarrierSwitching_srs_SwitchFromCarrier) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode SRS_CarrierSwitching_srs_SwitchFromCarrier", err)
	}
	return nil
}

func (ie *SRS_CarrierSwitching_srs_SwitchFromCarrier) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode SRS_CarrierSwitching_srs_SwitchFromCarrier", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
