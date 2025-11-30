package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern0_subcarrierLocation_p0_Enum_s0  aper.Enumerated = 0
	CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern0_subcarrierLocation_p0_Enum_s2  aper.Enumerated = 1
	CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern0_subcarrierLocation_p0_Enum_s4  aper.Enumerated = 2
	CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern0_subcarrierLocation_p0_Enum_s6  aper.Enumerated = 3
	CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern0_subcarrierLocation_p0_Enum_s8  aper.Enumerated = 4
	CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern0_subcarrierLocation_p0_Enum_s10 aper.Enumerated = 5
)

type CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern0_subcarrierLocation_p0 struct {
	Value aper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern0_subcarrierLocation_p0) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern0_subcarrierLocation_p0", err)
	}
	return nil
}

func (ie *CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern0_subcarrierLocation_p0) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern0_subcarrierLocation_p0", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
