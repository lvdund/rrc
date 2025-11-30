package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_SwitchingTimeNR_switchingTimeDL_Enum_n0us   aper.Enumerated = 0
	SRS_SwitchingTimeNR_switchingTimeDL_Enum_n30us  aper.Enumerated = 1
	SRS_SwitchingTimeNR_switchingTimeDL_Enum_n100us aper.Enumerated = 2
	SRS_SwitchingTimeNR_switchingTimeDL_Enum_n140us aper.Enumerated = 3
	SRS_SwitchingTimeNR_switchingTimeDL_Enum_n200us aper.Enumerated = 4
	SRS_SwitchingTimeNR_switchingTimeDL_Enum_n300us aper.Enumerated = 5
	SRS_SwitchingTimeNR_switchingTimeDL_Enum_n500us aper.Enumerated = 6
	SRS_SwitchingTimeNR_switchingTimeDL_Enum_n900us aper.Enumerated = 7
)

type SRS_SwitchingTimeNR_switchingTimeDL struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *SRS_SwitchingTimeNR_switchingTimeDL) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SRS_SwitchingTimeNR_switchingTimeDL", err)
	}
	return nil
}

func (ie *SRS_SwitchingTimeNR_switchingTimeDL) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SRS_SwitchingTimeNR_switchingTimeDL", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
