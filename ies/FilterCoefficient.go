package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	FilterCoefficient_Enum_fc0    aper.Enumerated = 0
	FilterCoefficient_Enum_fc1    aper.Enumerated = 1
	FilterCoefficient_Enum_fc2    aper.Enumerated = 2
	FilterCoefficient_Enum_fc3    aper.Enumerated = 3
	FilterCoefficient_Enum_fc4    aper.Enumerated = 4
	FilterCoefficient_Enum_fc5    aper.Enumerated = 5
	FilterCoefficient_Enum_fc6    aper.Enumerated = 6
	FilterCoefficient_Enum_fc7    aper.Enumerated = 7
	FilterCoefficient_Enum_fc8    aper.Enumerated = 8
	FilterCoefficient_Enum_fc9    aper.Enumerated = 9
	FilterCoefficient_Enum_fc11   aper.Enumerated = 10
	FilterCoefficient_Enum_fc13   aper.Enumerated = 11
	FilterCoefficient_Enum_fc15   aper.Enumerated = 12
	FilterCoefficient_Enum_fc17   aper.Enumerated = 13
	FilterCoefficient_Enum_fc19   aper.Enumerated = 14
	FilterCoefficient_Enum_spare1 aper.Enumerated = 15
)

type FilterCoefficient struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *FilterCoefficient) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode FilterCoefficient", err)
	}
	return nil
}

func (ie *FilterCoefficient) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode FilterCoefficient", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
