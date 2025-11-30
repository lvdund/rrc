package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RRM_Config_ue_InactiveTime_Enum_s1            aper.Enumerated = 0
	RRM_Config_ue_InactiveTime_Enum_s2            aper.Enumerated = 1
	RRM_Config_ue_InactiveTime_Enum_s3            aper.Enumerated = 2
	RRM_Config_ue_InactiveTime_Enum_s5            aper.Enumerated = 3
	RRM_Config_ue_InactiveTime_Enum_s7            aper.Enumerated = 4
	RRM_Config_ue_InactiveTime_Enum_s10           aper.Enumerated = 5
	RRM_Config_ue_InactiveTime_Enum_s15           aper.Enumerated = 6
	RRM_Config_ue_InactiveTime_Enum_s20           aper.Enumerated = 7
	RRM_Config_ue_InactiveTime_Enum_s25           aper.Enumerated = 8
	RRM_Config_ue_InactiveTime_Enum_s30           aper.Enumerated = 9
	RRM_Config_ue_InactiveTime_Enum_s40           aper.Enumerated = 10
	RRM_Config_ue_InactiveTime_Enum_s50           aper.Enumerated = 11
	RRM_Config_ue_InactiveTime_Enum_min1          aper.Enumerated = 12
	RRM_Config_ue_InactiveTime_Enum_min1s20       aper.Enumerated = 13
	RRM_Config_ue_InactiveTime_Enum_min1s40       aper.Enumerated = 14
	RRM_Config_ue_InactiveTime_Enum_min2          aper.Enumerated = 15
	RRM_Config_ue_InactiveTime_Enum_min2s30       aper.Enumerated = 16
	RRM_Config_ue_InactiveTime_Enum_min3          aper.Enumerated = 17
	RRM_Config_ue_InactiveTime_Enum_min3s30       aper.Enumerated = 18
	RRM_Config_ue_InactiveTime_Enum_min4          aper.Enumerated = 19
	RRM_Config_ue_InactiveTime_Enum_min5          aper.Enumerated = 20
	RRM_Config_ue_InactiveTime_Enum_min6          aper.Enumerated = 21
	RRM_Config_ue_InactiveTime_Enum_min7          aper.Enumerated = 22
	RRM_Config_ue_InactiveTime_Enum_min8          aper.Enumerated = 23
	RRM_Config_ue_InactiveTime_Enum_min9          aper.Enumerated = 24
	RRM_Config_ue_InactiveTime_Enum_min10         aper.Enumerated = 25
	RRM_Config_ue_InactiveTime_Enum_min12         aper.Enumerated = 26
	RRM_Config_ue_InactiveTime_Enum_min14         aper.Enumerated = 27
	RRM_Config_ue_InactiveTime_Enum_min17         aper.Enumerated = 28
	RRM_Config_ue_InactiveTime_Enum_min20         aper.Enumerated = 29
	RRM_Config_ue_InactiveTime_Enum_min24         aper.Enumerated = 30
	RRM_Config_ue_InactiveTime_Enum_min28         aper.Enumerated = 31
	RRM_Config_ue_InactiveTime_Enum_min33         aper.Enumerated = 32
	RRM_Config_ue_InactiveTime_Enum_min38         aper.Enumerated = 33
	RRM_Config_ue_InactiveTime_Enum_min44         aper.Enumerated = 34
	RRM_Config_ue_InactiveTime_Enum_min50         aper.Enumerated = 35
	RRM_Config_ue_InactiveTime_Enum_hr1           aper.Enumerated = 36
	RRM_Config_ue_InactiveTime_Enum_hr1min30      aper.Enumerated = 37
	RRM_Config_ue_InactiveTime_Enum_hr2           aper.Enumerated = 38
	RRM_Config_ue_InactiveTime_Enum_hr2min30      aper.Enumerated = 39
	RRM_Config_ue_InactiveTime_Enum_hr3           aper.Enumerated = 40
	RRM_Config_ue_InactiveTime_Enum_hr3min30      aper.Enumerated = 41
	RRM_Config_ue_InactiveTime_Enum_hr4           aper.Enumerated = 42
	RRM_Config_ue_InactiveTime_Enum_hr5           aper.Enumerated = 43
	RRM_Config_ue_InactiveTime_Enum_hr6           aper.Enumerated = 44
	RRM_Config_ue_InactiveTime_Enum_hr8           aper.Enumerated = 45
	RRM_Config_ue_InactiveTime_Enum_hr10          aper.Enumerated = 46
	RRM_Config_ue_InactiveTime_Enum_hr13          aper.Enumerated = 47
	RRM_Config_ue_InactiveTime_Enum_hr16          aper.Enumerated = 48
	RRM_Config_ue_InactiveTime_Enum_hr20          aper.Enumerated = 49
	RRM_Config_ue_InactiveTime_Enum_day1          aper.Enumerated = 50
	RRM_Config_ue_InactiveTime_Enum_day1hr12      aper.Enumerated = 51
	RRM_Config_ue_InactiveTime_Enum_day2          aper.Enumerated = 52
	RRM_Config_ue_InactiveTime_Enum_day2hr12      aper.Enumerated = 53
	RRM_Config_ue_InactiveTime_Enum_day3          aper.Enumerated = 54
	RRM_Config_ue_InactiveTime_Enum_day4          aper.Enumerated = 55
	RRM_Config_ue_InactiveTime_Enum_day5          aper.Enumerated = 56
	RRM_Config_ue_InactiveTime_Enum_day7          aper.Enumerated = 57
	RRM_Config_ue_InactiveTime_Enum_day10         aper.Enumerated = 58
	RRM_Config_ue_InactiveTime_Enum_day14         aper.Enumerated = 59
	RRM_Config_ue_InactiveTime_Enum_day19         aper.Enumerated = 60
	RRM_Config_ue_InactiveTime_Enum_day24         aper.Enumerated = 61
	RRM_Config_ue_InactiveTime_Enum_day30         aper.Enumerated = 62
	RRM_Config_ue_InactiveTime_Enum_dayMoreThan30 aper.Enumerated = 63
)

type RRM_Config_ue_InactiveTime struct {
	Value aper.Enumerated `lb:0,ub:63,madatory`
}

func (ie *RRM_Config_ue_InactiveTime) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("Encode RRM_Config_ue_InactiveTime", err)
	}
	return nil
}

func (ie *RRM_Config_ue_InactiveTime) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("Decode RRM_Config_ue_InactiveTime", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
