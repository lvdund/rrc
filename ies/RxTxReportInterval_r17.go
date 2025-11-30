package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RxTxReportInterval_r17_Enum_ms80   aper.Enumerated = 0
	RxTxReportInterval_r17_Enum_ms120  aper.Enumerated = 1
	RxTxReportInterval_r17_Enum_ms160  aper.Enumerated = 2
	RxTxReportInterval_r17_Enum_ms240  aper.Enumerated = 3
	RxTxReportInterval_r17_Enum_ms320  aper.Enumerated = 4
	RxTxReportInterval_r17_Enum_ms480  aper.Enumerated = 5
	RxTxReportInterval_r17_Enum_ms640  aper.Enumerated = 6
	RxTxReportInterval_r17_Enum_ms1024 aper.Enumerated = 7
	RxTxReportInterval_r17_Enum_ms1280 aper.Enumerated = 8
	RxTxReportInterval_r17_Enum_ms2048 aper.Enumerated = 9
	RxTxReportInterval_r17_Enum_ms2560 aper.Enumerated = 10
	RxTxReportInterval_r17_Enum_ms5120 aper.Enumerated = 11
	RxTxReportInterval_r17_Enum_spare4 aper.Enumerated = 12
	RxTxReportInterval_r17_Enum_spare3 aper.Enumerated = 13
	RxTxReportInterval_r17_Enum_spare2 aper.Enumerated = 14
	RxTxReportInterval_r17_Enum_spare1 aper.Enumerated = 15
)

type RxTxReportInterval_r17 struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *RxTxReportInterval_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode RxTxReportInterval_r17", err)
	}
	return nil
}

func (ie *RxTxReportInterval_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode RxTxReportInterval_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
