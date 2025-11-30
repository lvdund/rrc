package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RadioLinkMonitoringRS_purpose_Enum_beamFailure aper.Enumerated = 0
	RadioLinkMonitoringRS_purpose_Enum_rlf         aper.Enumerated = 1
	RadioLinkMonitoringRS_purpose_Enum_both        aper.Enumerated = 2
)

type RadioLinkMonitoringRS_purpose struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *RadioLinkMonitoringRS_purpose) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode RadioLinkMonitoringRS_purpose", err)
	}
	return nil
}

func (ie *RadioLinkMonitoringRS_purpose) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode RadioLinkMonitoringRS_purpose", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
