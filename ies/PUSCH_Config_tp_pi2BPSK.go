package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PUSCH_Config_tp_pi2BPSK_Enum_enabled aper.Enumerated = 0
)

type PUSCH_Config_tp_pi2BPSK struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *PUSCH_Config_tp_pi2BPSK) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode PUSCH_Config_tp_pi2BPSK", err)
	}
	return nil
}

func (ie *PUSCH_Config_tp_pi2BPSK) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode PUSCH_Config_tp_pi2BPSK", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
