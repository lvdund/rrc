package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MRB_PDCP_ConfigBroadcast_r17_t_Reordering_r17_Enum_ms1    aper.Enumerated = 0
	MRB_PDCP_ConfigBroadcast_r17_t_Reordering_r17_Enum_ms10   aper.Enumerated = 1
	MRB_PDCP_ConfigBroadcast_r17_t_Reordering_r17_Enum_ms40   aper.Enumerated = 2
	MRB_PDCP_ConfigBroadcast_r17_t_Reordering_r17_Enum_ms160  aper.Enumerated = 3
	MRB_PDCP_ConfigBroadcast_r17_t_Reordering_r17_Enum_ms500  aper.Enumerated = 4
	MRB_PDCP_ConfigBroadcast_r17_t_Reordering_r17_Enum_ms1000 aper.Enumerated = 5
	MRB_PDCP_ConfigBroadcast_r17_t_Reordering_r17_Enum_ms1250 aper.Enumerated = 6
	MRB_PDCP_ConfigBroadcast_r17_t_Reordering_r17_Enum_ms2750 aper.Enumerated = 7
)

type MRB_PDCP_ConfigBroadcast_r17_t_Reordering_r17 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *MRB_PDCP_ConfigBroadcast_r17_t_Reordering_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode MRB_PDCP_ConfigBroadcast_r17_t_Reordering_r17", err)
	}
	return nil
}

func (ie *MRB_PDCP_ConfigBroadcast_r17_t_Reordering_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode MRB_PDCP_ConfigBroadcast_r17_t_Reordering_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
