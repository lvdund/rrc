package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_RS_ResourceMapping_cdm_Type_Enum_noCDM        aper.Enumerated = 0
	CSI_RS_ResourceMapping_cdm_Type_Enum_fd_CDM2      aper.Enumerated = 1
	CSI_RS_ResourceMapping_cdm_Type_Enum_cdm4_FD2_TD2 aper.Enumerated = 2
	CSI_RS_ResourceMapping_cdm_Type_Enum_cdm8_FD2_TD4 aper.Enumerated = 3
)

type CSI_RS_ResourceMapping_cdm_Type struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *CSI_RS_ResourceMapping_cdm_Type) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode CSI_RS_ResourceMapping_cdm_Type", err)
	}
	return nil
}

func (ie *CSI_RS_ResourceMapping_cdm_Type) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode CSI_RS_ResourceMapping_cdm_Type", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
