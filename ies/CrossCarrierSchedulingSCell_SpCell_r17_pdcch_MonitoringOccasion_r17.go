package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CrossCarrierSchedulingSCell_SpCell_r17_pdcch_MonitoringOccasion_r17_Enum_val1 aper.Enumerated = 0
	CrossCarrierSchedulingSCell_SpCell_r17_pdcch_MonitoringOccasion_r17_Enum_val2 aper.Enumerated = 1
)

type CrossCarrierSchedulingSCell_SpCell_r17_pdcch_MonitoringOccasion_r17 struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *CrossCarrierSchedulingSCell_SpCell_r17_pdcch_MonitoringOccasion_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode CrossCarrierSchedulingSCell_SpCell_r17_pdcch_MonitoringOccasion_r17", err)
	}
	return nil
}

func (ie *CrossCarrierSchedulingSCell_SpCell_r17_pdcch_MonitoringOccasion_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode CrossCarrierSchedulingSCell_SpCell_r17_pdcch_MonitoringOccasion_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
