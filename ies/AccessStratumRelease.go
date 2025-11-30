package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	AccessStratumRelease_Enum_rel15  aper.Enumerated = 0
	AccessStratumRelease_Enum_rel16  aper.Enumerated = 1
	AccessStratumRelease_Enum_rel17  aper.Enumerated = 2
	AccessStratumRelease_Enum_spare5 aper.Enumerated = 3
	AccessStratumRelease_Enum_spare4 aper.Enumerated = 4
	AccessStratumRelease_Enum_spare3 aper.Enumerated = 5
	AccessStratumRelease_Enum_spare2 aper.Enumerated = 6
	AccessStratumRelease_Enum_spare1 aper.Enumerated = 7
)

type AccessStratumRelease struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *AccessStratumRelease) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode AccessStratumRelease", err)
	}
	return nil
}

func (ie *AccessStratumRelease) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode AccessStratumRelease", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
