package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PDCP_Config_drb_pdcp_SN_SizeUL_Enum_len12bits aper.Enumerated = 0
	PDCP_Config_drb_pdcp_SN_SizeUL_Enum_len18bits aper.Enumerated = 1
)

type PDCP_Config_drb_pdcp_SN_SizeUL struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *PDCP_Config_drb_pdcp_SN_SizeUL) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode PDCP_Config_drb_pdcp_SN_SizeUL", err)
	}
	return nil
}

func (ie *PDCP_Config_drb_pdcp_SN_SizeUL) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode PDCP_Config_drb_pdcp_SN_SizeUL", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
