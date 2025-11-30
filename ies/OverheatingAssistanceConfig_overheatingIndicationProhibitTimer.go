package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	OverheatingAssistanceConfig_overheatingIndicationProhibitTimer_Enum_s0     aper.Enumerated = 0
	OverheatingAssistanceConfig_overheatingIndicationProhibitTimer_Enum_s0dot5 aper.Enumerated = 1
	OverheatingAssistanceConfig_overheatingIndicationProhibitTimer_Enum_s1     aper.Enumerated = 2
	OverheatingAssistanceConfig_overheatingIndicationProhibitTimer_Enum_s2     aper.Enumerated = 3
	OverheatingAssistanceConfig_overheatingIndicationProhibitTimer_Enum_s5     aper.Enumerated = 4
	OverheatingAssistanceConfig_overheatingIndicationProhibitTimer_Enum_s10    aper.Enumerated = 5
	OverheatingAssistanceConfig_overheatingIndicationProhibitTimer_Enum_s20    aper.Enumerated = 6
	OverheatingAssistanceConfig_overheatingIndicationProhibitTimer_Enum_s30    aper.Enumerated = 7
	OverheatingAssistanceConfig_overheatingIndicationProhibitTimer_Enum_s60    aper.Enumerated = 8
	OverheatingAssistanceConfig_overheatingIndicationProhibitTimer_Enum_s90    aper.Enumerated = 9
	OverheatingAssistanceConfig_overheatingIndicationProhibitTimer_Enum_s120   aper.Enumerated = 10
	OverheatingAssistanceConfig_overheatingIndicationProhibitTimer_Enum_s300   aper.Enumerated = 11
	OverheatingAssistanceConfig_overheatingIndicationProhibitTimer_Enum_s600   aper.Enumerated = 12
	OverheatingAssistanceConfig_overheatingIndicationProhibitTimer_Enum_spare3 aper.Enumerated = 13
	OverheatingAssistanceConfig_overheatingIndicationProhibitTimer_Enum_spare2 aper.Enumerated = 14
	OverheatingAssistanceConfig_overheatingIndicationProhibitTimer_Enum_spare1 aper.Enumerated = 15
)

type OverheatingAssistanceConfig_overheatingIndicationProhibitTimer struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *OverheatingAssistanceConfig_overheatingIndicationProhibitTimer) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode OverheatingAssistanceConfig_overheatingIndicationProhibitTimer", err)
	}
	return nil
}

func (ie *OverheatingAssistanceConfig_overheatingIndicationProhibitTimer) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode OverheatingAssistanceConfig_overheatingIndicationProhibitTimer", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
