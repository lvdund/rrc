package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	FailureInfoRLC_Bearer_failureType_Enum_rlc_failure aper.Enumerated = 0
	FailureInfoRLC_Bearer_failureType_Enum_spare3      aper.Enumerated = 1
	FailureInfoRLC_Bearer_failureType_Enum_spare2      aper.Enumerated = 2
	FailureInfoRLC_Bearer_failureType_Enum_spare1      aper.Enumerated = 3
)

type FailureInfoRLC_Bearer_failureType struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *FailureInfoRLC_Bearer_failureType) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode FailureInfoRLC_Bearer_failureType", err)
	}
	return nil
}

func (ie *FailureInfoRLC_Bearer_failureType) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode FailureInfoRLC_Bearer_failureType", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
