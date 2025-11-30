package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	UE_CapabilityRequestFilterNR_v1540_srs_SwitchingTimeRequest_Enum_true aper.Enumerated = 0
)

type UE_CapabilityRequestFilterNR_v1540_srs_SwitchingTimeRequest struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *UE_CapabilityRequestFilterNR_v1540_srs_SwitchingTimeRequest) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode UE_CapabilityRequestFilterNR_v1540_srs_SwitchingTimeRequest", err)
	}
	return nil
}

func (ie *UE_CapabilityRequestFilterNR_v1540_srs_SwitchingTimeRequest) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode UE_CapabilityRequestFilterNR_v1540_srs_SwitchingTimeRequest", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
