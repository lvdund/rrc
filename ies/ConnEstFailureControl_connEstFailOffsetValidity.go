package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	ConnEstFailureControl_connEstFailOffsetValidity_Enum_s30  aper.Enumerated = 0
	ConnEstFailureControl_connEstFailOffsetValidity_Enum_s60  aper.Enumerated = 1
	ConnEstFailureControl_connEstFailOffsetValidity_Enum_s120 aper.Enumerated = 2
	ConnEstFailureControl_connEstFailOffsetValidity_Enum_s240 aper.Enumerated = 3
	ConnEstFailureControl_connEstFailOffsetValidity_Enum_s300 aper.Enumerated = 4
	ConnEstFailureControl_connEstFailOffsetValidity_Enum_s420 aper.Enumerated = 5
	ConnEstFailureControl_connEstFailOffsetValidity_Enum_s600 aper.Enumerated = 6
	ConnEstFailureControl_connEstFailOffsetValidity_Enum_s900 aper.Enumerated = 7
)

type ConnEstFailureControl_connEstFailOffsetValidity struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *ConnEstFailureControl_connEstFailOffsetValidity) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode ConnEstFailureControl_connEstFailOffsetValidity", err)
	}
	return nil
}

func (ie *ConnEstFailureControl_connEstFailOffsetValidity) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode ConnEstFailureControl_connEstFailOffsetValidity", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
