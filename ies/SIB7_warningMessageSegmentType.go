package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SIB7_warningMessageSegmentType_Enum_notLastSegment aper.Enumerated = 0
	SIB7_warningMessageSegmentType_Enum_lastSegment    aper.Enumerated = 1
)

type SIB7_warningMessageSegmentType struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *SIB7_warningMessageSegmentType) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode SIB7_warningMessageSegmentType", err)
	}
	return nil
}

func (ie *SIB7_warningMessageSegmentType) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode SIB7_warningMessageSegmentType", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
