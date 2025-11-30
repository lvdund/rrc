package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RRCRelease_IEs_deprioritisationReq_deprioritisationType_Enum_frequency aper.Enumerated = 0
	RRCRelease_IEs_deprioritisationReq_deprioritisationType_Enum_nr        aper.Enumerated = 1
)

type RRCRelease_IEs_deprioritisationReq_deprioritisationType struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *RRCRelease_IEs_deprioritisationReq_deprioritisationType) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode RRCRelease_IEs_deprioritisationReq_deprioritisationType", err)
	}
	return nil
}

func (ie *RRCRelease_IEs_deprioritisationReq_deprioritisationType) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode RRCRelease_IEs_deprioritisationReq_deprioritisationType", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
