package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_SwitchingTimeEUTRA_switchingTimeDL_Enum_n0     aper.Enumerated = 0
	SRS_SwitchingTimeEUTRA_switchingTimeDL_Enum_n0dot5 aper.Enumerated = 1
	SRS_SwitchingTimeEUTRA_switchingTimeDL_Enum_n1     aper.Enumerated = 2
	SRS_SwitchingTimeEUTRA_switchingTimeDL_Enum_n1dot5 aper.Enumerated = 3
	SRS_SwitchingTimeEUTRA_switchingTimeDL_Enum_n2     aper.Enumerated = 4
	SRS_SwitchingTimeEUTRA_switchingTimeDL_Enum_n2dot5 aper.Enumerated = 5
	SRS_SwitchingTimeEUTRA_switchingTimeDL_Enum_n3     aper.Enumerated = 6
	SRS_SwitchingTimeEUTRA_switchingTimeDL_Enum_n3dot5 aper.Enumerated = 7
	SRS_SwitchingTimeEUTRA_switchingTimeDL_Enum_n4     aper.Enumerated = 8
	SRS_SwitchingTimeEUTRA_switchingTimeDL_Enum_n4dot5 aper.Enumerated = 9
	SRS_SwitchingTimeEUTRA_switchingTimeDL_Enum_n5     aper.Enumerated = 10
	SRS_SwitchingTimeEUTRA_switchingTimeDL_Enum_n5dot5 aper.Enumerated = 11
	SRS_SwitchingTimeEUTRA_switchingTimeDL_Enum_n6     aper.Enumerated = 12
	SRS_SwitchingTimeEUTRA_switchingTimeDL_Enum_n6dot5 aper.Enumerated = 13
	SRS_SwitchingTimeEUTRA_switchingTimeDL_Enum_n7     aper.Enumerated = 14
)

type SRS_SwitchingTimeEUTRA_switchingTimeDL struct {
	Value aper.Enumerated `lb:0,ub:14,madatory`
}

func (ie *SRS_SwitchingTimeEUTRA_switchingTimeDL) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 14}, false); err != nil {
		return utils.WrapError("Encode SRS_SwitchingTimeEUTRA_switchingTimeDL", err)
	}
	return nil
}

func (ie *SRS_SwitchingTimeEUTRA_switchingTimeDL) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 14}, false); err != nil {
		return utils.WrapError("Decode SRS_SwitchingTimeEUTRA_switchingTimeDL", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
