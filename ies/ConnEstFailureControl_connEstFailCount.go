package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	ConnEstFailureControl_connEstFailCount_Enum_n1 aper.Enumerated = 0
	ConnEstFailureControl_connEstFailCount_Enum_n2 aper.Enumerated = 1
	ConnEstFailureControl_connEstFailCount_Enum_n3 aper.Enumerated = 2
	ConnEstFailureControl_connEstFailCount_Enum_n4 aper.Enumerated = 3
)

type ConnEstFailureControl_connEstFailCount struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *ConnEstFailureControl_connEstFailCount) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode ConnEstFailureControl_connEstFailCount", err)
	}
	return nil
}

func (ie *ConnEstFailureControl_connEstFailCount) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode ConnEstFailureControl_connEstFailCount", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
