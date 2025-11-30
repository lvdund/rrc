package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CrossCarrierSchedulingConfig_ccs_BlindDetectionSplit_r17_Enum_oneSeventh      aper.Enumerated = 0
	CrossCarrierSchedulingConfig_ccs_BlindDetectionSplit_r17_Enum_threeFourteenth aper.Enumerated = 1
	CrossCarrierSchedulingConfig_ccs_BlindDetectionSplit_r17_Enum_twoSeventh      aper.Enumerated = 2
	CrossCarrierSchedulingConfig_ccs_BlindDetectionSplit_r17_Enum_threeSeventh    aper.Enumerated = 3
	CrossCarrierSchedulingConfig_ccs_BlindDetectionSplit_r17_Enum_oneHalf         aper.Enumerated = 4
	CrossCarrierSchedulingConfig_ccs_BlindDetectionSplit_r17_Enum_fourSeventh     aper.Enumerated = 5
	CrossCarrierSchedulingConfig_ccs_BlindDetectionSplit_r17_Enum_fiveSeventh     aper.Enumerated = 6
	CrossCarrierSchedulingConfig_ccs_BlindDetectionSplit_r17_Enum_spare1          aper.Enumerated = 7
)

type CrossCarrierSchedulingConfig_ccs_BlindDetectionSplit_r17 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *CrossCarrierSchedulingConfig_ccs_BlindDetectionSplit_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode CrossCarrierSchedulingConfig_ccs_BlindDetectionSplit_r17", err)
	}
	return nil
}

func (ie *CrossCarrierSchedulingConfig_ccs_BlindDetectionSplit_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode CrossCarrierSchedulingConfig_ccs_BlindDetectionSplit_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
