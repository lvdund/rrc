package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SDAP_Config_sdap_HeaderDL_Enum_present aper.Enumerated = 0
	SDAP_Config_sdap_HeaderDL_Enum_absent  aper.Enumerated = 1
)

type SDAP_Config_sdap_HeaderDL struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *SDAP_Config_sdap_HeaderDL) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode SDAP_Config_sdap_HeaderDL", err)
	}
	return nil
}

func (ie *SDAP_Config_sdap_HeaderDL) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode SDAP_Config_sdap_HeaderDL", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
