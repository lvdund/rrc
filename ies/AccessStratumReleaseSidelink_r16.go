package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	AccessStratumReleaseSidelink_r16_Enum_rel16  aper.Enumerated = 0
	AccessStratumReleaseSidelink_r16_Enum_rel17  aper.Enumerated = 1
	AccessStratumReleaseSidelink_r16_Enum_spare6 aper.Enumerated = 2
	AccessStratumReleaseSidelink_r16_Enum_spare5 aper.Enumerated = 3
	AccessStratumReleaseSidelink_r16_Enum_spare4 aper.Enumerated = 4
	AccessStratumReleaseSidelink_r16_Enum_spare3 aper.Enumerated = 5
	AccessStratumReleaseSidelink_r16_Enum_spare2 aper.Enumerated = 6
	AccessStratumReleaseSidelink_r16_Enum_spare1 aper.Enumerated = 7
)

type AccessStratumReleaseSidelink_r16 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *AccessStratumReleaseSidelink_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode AccessStratumReleaseSidelink_r16", err)
	}
	return nil
}

func (ie *AccessStratumReleaseSidelink_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode AccessStratumReleaseSidelink_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
