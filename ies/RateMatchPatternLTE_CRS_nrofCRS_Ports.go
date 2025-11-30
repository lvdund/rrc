package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RateMatchPatternLTE_CRS_nrofCRS_Ports_Enum_n1 aper.Enumerated = 0
	RateMatchPatternLTE_CRS_nrofCRS_Ports_Enum_n2 aper.Enumerated = 1
	RateMatchPatternLTE_CRS_nrofCRS_Ports_Enum_n4 aper.Enumerated = 2
)

type RateMatchPatternLTE_CRS_nrofCRS_Ports struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *RateMatchPatternLTE_CRS_nrofCRS_Ports) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode RateMatchPatternLTE_CRS_nrofCRS_Ports", err)
	}
	return nil
}

func (ie *RateMatchPatternLTE_CRS_nrofCRS_Ports) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode RateMatchPatternLTE_CRS_nrofCRS_Ports", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
