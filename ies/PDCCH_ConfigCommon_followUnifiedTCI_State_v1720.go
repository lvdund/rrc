package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PDCCH_ConfigCommon_followUnifiedTCI_State_v1720_Enum_enabled aper.Enumerated = 0
)

type PDCCH_ConfigCommon_followUnifiedTCI_State_v1720 struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *PDCCH_ConfigCommon_followUnifiedTCI_State_v1720) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode PDCCH_ConfigCommon_followUnifiedTCI_State_v1720", err)
	}
	return nil
}

func (ie *PDCCH_ConfigCommon_followUnifiedTCI_State_v1720) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode PDCCH_ConfigCommon_followUnifiedTCI_State_v1720", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
