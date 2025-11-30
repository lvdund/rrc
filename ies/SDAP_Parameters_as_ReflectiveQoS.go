package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SDAP_Parameters_as_ReflectiveQoS_Enum_true aper.Enumerated = 0
)

type SDAP_Parameters_as_ReflectiveQoS struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *SDAP_Parameters_as_ReflectiveQoS) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode SDAP_Parameters_as_ReflectiveQoS", err)
	}
	return nil
}

func (ie *SDAP_Parameters_as_ReflectiveQoS) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode SDAP_Parameters_as_ReflectiveQoS", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
