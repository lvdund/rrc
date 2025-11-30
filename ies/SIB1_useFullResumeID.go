package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SIB1_useFullResumeID_Enum_true aper.Enumerated = 0
)

type SIB1_useFullResumeID struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *SIB1_useFullResumeID) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode SIB1_useFullResumeID", err)
	}
	return nil
}

func (ie *SIB1_useFullResumeID) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode SIB1_useFullResumeID", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
