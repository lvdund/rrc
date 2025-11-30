package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MobilityStateParameters_t_Evaluation_Enum_s30    aper.Enumerated = 0
	MobilityStateParameters_t_Evaluation_Enum_s60    aper.Enumerated = 1
	MobilityStateParameters_t_Evaluation_Enum_s120   aper.Enumerated = 2
	MobilityStateParameters_t_Evaluation_Enum_s180   aper.Enumerated = 3
	MobilityStateParameters_t_Evaluation_Enum_s240   aper.Enumerated = 4
	MobilityStateParameters_t_Evaluation_Enum_spare3 aper.Enumerated = 5
	MobilityStateParameters_t_Evaluation_Enum_spare2 aper.Enumerated = 6
	MobilityStateParameters_t_Evaluation_Enum_spare1 aper.Enumerated = 7
)

type MobilityStateParameters_t_Evaluation struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *MobilityStateParameters_t_Evaluation) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode MobilityStateParameters_t_Evaluation", err)
	}
	return nil
}

func (ie *MobilityStateParameters_t_Evaluation) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode MobilityStateParameters_t_Evaluation", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
