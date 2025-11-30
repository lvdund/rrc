package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PDCP_Config_drb_integrityProtection_Enum_enabled aper.Enumerated = 0
)

type PDCP_Config_drb_integrityProtection struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *PDCP_Config_drb_integrityProtection) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode PDCP_Config_drb_integrityProtection", err)
	}
	return nil
}

func (ie *PDCP_Config_drb_integrityProtection) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode PDCP_Config_drb_integrityProtection", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
