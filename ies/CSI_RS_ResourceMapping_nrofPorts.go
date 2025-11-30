package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_RS_ResourceMapping_nrofPorts_Enum_p1  aper.Enumerated = 0
	CSI_RS_ResourceMapping_nrofPorts_Enum_p2  aper.Enumerated = 1
	CSI_RS_ResourceMapping_nrofPorts_Enum_p4  aper.Enumerated = 2
	CSI_RS_ResourceMapping_nrofPorts_Enum_p8  aper.Enumerated = 3
	CSI_RS_ResourceMapping_nrofPorts_Enum_p12 aper.Enumerated = 4
	CSI_RS_ResourceMapping_nrofPorts_Enum_p16 aper.Enumerated = 5
	CSI_RS_ResourceMapping_nrofPorts_Enum_p24 aper.Enumerated = 6
	CSI_RS_ResourceMapping_nrofPorts_Enum_p32 aper.Enumerated = 7
)

type CSI_RS_ResourceMapping_nrofPorts struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *CSI_RS_ResourceMapping_nrofPorts) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode CSI_RS_ResourceMapping_nrofPorts", err)
	}
	return nil
}

func (ie *CSI_RS_ResourceMapping_nrofPorts) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode CSI_RS_ResourceMapping_nrofPorts", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
