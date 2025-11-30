package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	DataInactivityTimer_Enum_s1   aper.Enumerated = 0
	DataInactivityTimer_Enum_s2   aper.Enumerated = 1
	DataInactivityTimer_Enum_s3   aper.Enumerated = 2
	DataInactivityTimer_Enum_s5   aper.Enumerated = 3
	DataInactivityTimer_Enum_s7   aper.Enumerated = 4
	DataInactivityTimer_Enum_s10  aper.Enumerated = 5
	DataInactivityTimer_Enum_s15  aper.Enumerated = 6
	DataInactivityTimer_Enum_s20  aper.Enumerated = 7
	DataInactivityTimer_Enum_s40  aper.Enumerated = 8
	DataInactivityTimer_Enum_s50  aper.Enumerated = 9
	DataInactivityTimer_Enum_s60  aper.Enumerated = 10
	DataInactivityTimer_Enum_s80  aper.Enumerated = 11
	DataInactivityTimer_Enum_s100 aper.Enumerated = 12
	DataInactivityTimer_Enum_s120 aper.Enumerated = 13
	DataInactivityTimer_Enum_s150 aper.Enumerated = 14
	DataInactivityTimer_Enum_s180 aper.Enumerated = 15
)

type DataInactivityTimer struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *DataInactivityTimer) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode DataInactivityTimer", err)
	}
	return nil
}

func (ie *DataInactivityTimer) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode DataInactivityTimer", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
