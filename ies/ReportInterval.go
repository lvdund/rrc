package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	ReportInterval_Enum_ms120   aper.Enumerated = 0
	ReportInterval_Enum_ms240   aper.Enumerated = 1
	ReportInterval_Enum_ms480   aper.Enumerated = 2
	ReportInterval_Enum_ms640   aper.Enumerated = 3
	ReportInterval_Enum_ms1024  aper.Enumerated = 4
	ReportInterval_Enum_ms2048  aper.Enumerated = 5
	ReportInterval_Enum_ms5120  aper.Enumerated = 6
	ReportInterval_Enum_ms10240 aper.Enumerated = 7
	ReportInterval_Enum_ms20480 aper.Enumerated = 8
	ReportInterval_Enum_ms40960 aper.Enumerated = 9
	ReportInterval_Enum_min1    aper.Enumerated = 10
	ReportInterval_Enum_min6    aper.Enumerated = 11
	ReportInterval_Enum_min12   aper.Enumerated = 12
	ReportInterval_Enum_min30   aper.Enumerated = 13
)

type ReportInterval struct {
	Value aper.Enumerated `lb:0,ub:13,madatory`
}

func (ie *ReportInterval) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 13}, false); err != nil {
		return utils.WrapError("Encode ReportInterval", err)
	}
	return nil
}

func (ie *ReportInterval) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 13}, false); err != nil {
		return utils.WrapError("Decode ReportInterval", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
