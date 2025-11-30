package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RRCRelease_IEs_deprioritisationReq_deprioritisationTimer_Enum_min5  aper.Enumerated = 0
	RRCRelease_IEs_deprioritisationReq_deprioritisationTimer_Enum_min10 aper.Enumerated = 1
	RRCRelease_IEs_deprioritisationReq_deprioritisationTimer_Enum_min15 aper.Enumerated = 2
	RRCRelease_IEs_deprioritisationReq_deprioritisationTimer_Enum_min30 aper.Enumerated = 3
)

type RRCRelease_IEs_deprioritisationReq_deprioritisationTimer struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *RRCRelease_IEs_deprioritisationReq_deprioritisationTimer) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode RRCRelease_IEs_deprioritisationReq_deprioritisationTimer", err)
	}
	return nil
}

func (ie *RRCRelease_IEs_deprioritisationReq_deprioritisationTimer) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode RRCRelease_IEs_deprioritisationReq_deprioritisationTimer", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
