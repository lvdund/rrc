package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	LTE_NeighCellsCRS_AssistInfo_r17_neighCarrierBandwidthDL_r17_Enum_n6     aper.Enumerated = 0
	LTE_NeighCellsCRS_AssistInfo_r17_neighCarrierBandwidthDL_r17_Enum_n15    aper.Enumerated = 1
	LTE_NeighCellsCRS_AssistInfo_r17_neighCarrierBandwidthDL_r17_Enum_n25    aper.Enumerated = 2
	LTE_NeighCellsCRS_AssistInfo_r17_neighCarrierBandwidthDL_r17_Enum_n50    aper.Enumerated = 3
	LTE_NeighCellsCRS_AssistInfo_r17_neighCarrierBandwidthDL_r17_Enum_n75    aper.Enumerated = 4
	LTE_NeighCellsCRS_AssistInfo_r17_neighCarrierBandwidthDL_r17_Enum_n100   aper.Enumerated = 5
	LTE_NeighCellsCRS_AssistInfo_r17_neighCarrierBandwidthDL_r17_Enum_spare2 aper.Enumerated = 6
	LTE_NeighCellsCRS_AssistInfo_r17_neighCarrierBandwidthDL_r17_Enum_spare1 aper.Enumerated = 7
)

type LTE_NeighCellsCRS_AssistInfo_r17_neighCarrierBandwidthDL_r17 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *LTE_NeighCellsCRS_AssistInfo_r17_neighCarrierBandwidthDL_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode LTE_NeighCellsCRS_AssistInfo_r17_neighCarrierBandwidthDL_r17", err)
	}
	return nil
}

func (ie *LTE_NeighCellsCRS_AssistInfo_r17_neighCarrierBandwidthDL_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode LTE_NeighCellsCRS_AssistInfo_r17_neighCarrierBandwidthDL_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
