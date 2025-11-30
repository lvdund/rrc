package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PeriodicalReportConfigInterRAT_reportAmount_Enum_r1       aper.Enumerated = 0
	PeriodicalReportConfigInterRAT_reportAmount_Enum_r2       aper.Enumerated = 1
	PeriodicalReportConfigInterRAT_reportAmount_Enum_r4       aper.Enumerated = 2
	PeriodicalReportConfigInterRAT_reportAmount_Enum_r8       aper.Enumerated = 3
	PeriodicalReportConfigInterRAT_reportAmount_Enum_r16      aper.Enumerated = 4
	PeriodicalReportConfigInterRAT_reportAmount_Enum_r32      aper.Enumerated = 5
	PeriodicalReportConfigInterRAT_reportAmount_Enum_r64      aper.Enumerated = 6
	PeriodicalReportConfigInterRAT_reportAmount_Enum_infinity aper.Enumerated = 7
)

type PeriodicalReportConfigInterRAT_reportAmount struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *PeriodicalReportConfigInterRAT_reportAmount) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode PeriodicalReportConfigInterRAT_reportAmount", err)
	}
	return nil
}

func (ie *PeriodicalReportConfigInterRAT_reportAmount) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode PeriodicalReportConfigInterRAT_reportAmount", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
