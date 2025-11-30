package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MIMO_ParametersPerBand_pusch_TransCoherence_Enum_nonCoherent     aper.Enumerated = 0
	MIMO_ParametersPerBand_pusch_TransCoherence_Enum_partialCoherent aper.Enumerated = 1
	MIMO_ParametersPerBand_pusch_TransCoherence_Enum_fullCoherent    aper.Enumerated = 2
)

type MIMO_ParametersPerBand_pusch_TransCoherence struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *MIMO_ParametersPerBand_pusch_TransCoherence) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode MIMO_ParametersPerBand_pusch_TransCoherence", err)
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_pusch_TransCoherence) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode MIMO_ParametersPerBand_pusch_TransCoherence", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
