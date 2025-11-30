package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CrossCarrierSchedulingProcessing_DiffSCS_r16_scs_15kHz_30kHz_r16_Enum_n2 aper.Enumerated = 0
)

type CrossCarrierSchedulingProcessing_DiffSCS_r16_scs_15kHz_30kHz_r16 struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *CrossCarrierSchedulingProcessing_DiffSCS_r16_scs_15kHz_30kHz_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode CrossCarrierSchedulingProcessing_DiffSCS_r16_scs_15kHz_30kHz_r16", err)
	}
	return nil
}

func (ie *CrossCarrierSchedulingProcessing_DiffSCS_r16_scs_15kHz_30kHz_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode CrossCarrierSchedulingProcessing_DiffSCS_r16_scs_15kHz_30kHz_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
