package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_RS_IM_ReceptionForFeedback_maxConfigNumberCSI_IM_PerCC_Enum_n1  aper.Enumerated = 0
	CSI_RS_IM_ReceptionForFeedback_maxConfigNumberCSI_IM_PerCC_Enum_n2  aper.Enumerated = 1
	CSI_RS_IM_ReceptionForFeedback_maxConfigNumberCSI_IM_PerCC_Enum_n4  aper.Enumerated = 2
	CSI_RS_IM_ReceptionForFeedback_maxConfigNumberCSI_IM_PerCC_Enum_n8  aper.Enumerated = 3
	CSI_RS_IM_ReceptionForFeedback_maxConfigNumberCSI_IM_PerCC_Enum_n16 aper.Enumerated = 4
	CSI_RS_IM_ReceptionForFeedback_maxConfigNumberCSI_IM_PerCC_Enum_n32 aper.Enumerated = 5
)

type CSI_RS_IM_ReceptionForFeedback_maxConfigNumberCSI_IM_PerCC struct {
	Value aper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *CSI_RS_IM_ReceptionForFeedback_maxConfigNumberCSI_IM_PerCC) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode CSI_RS_IM_ReceptionForFeedback_maxConfigNumberCSI_IM_PerCC", err)
	}
	return nil
}

func (ie *CSI_RS_IM_ReceptionForFeedback_maxConfigNumberCSI_IM_PerCC) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode CSI_RS_IM_ReceptionForFeedback_maxConfigNumberCSI_IM_PerCC", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
