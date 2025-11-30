package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	EUTRA_AllowedMeasBandwidth_Enum_mbw6   aper.Enumerated = 0
	EUTRA_AllowedMeasBandwidth_Enum_mbw15  aper.Enumerated = 1
	EUTRA_AllowedMeasBandwidth_Enum_mbw25  aper.Enumerated = 2
	EUTRA_AllowedMeasBandwidth_Enum_mbw50  aper.Enumerated = 3
	EUTRA_AllowedMeasBandwidth_Enum_mbw75  aper.Enumerated = 4
	EUTRA_AllowedMeasBandwidth_Enum_mbw100 aper.Enumerated = 5
)

type EUTRA_AllowedMeasBandwidth struct {
	Value aper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *EUTRA_AllowedMeasBandwidth) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode EUTRA_AllowedMeasBandwidth", err)
	}
	return nil
}

func (ie *EUTRA_AllowedMeasBandwidth) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode EUTRA_AllowedMeasBandwidth", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
