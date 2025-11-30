package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SBAS_ID_r16_sbas_id_r16_Enum_waas  aper.Enumerated = 0
	SBAS_ID_r16_sbas_id_r16_Enum_egnos aper.Enumerated = 1
	SBAS_ID_r16_sbas_id_r16_Enum_msas  aper.Enumerated = 2
	SBAS_ID_r16_sbas_id_r16_Enum_gagan aper.Enumerated = 3
)

type SBAS_ID_r16_sbas_id_r16 struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *SBAS_ID_r16_sbas_id_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode SBAS_ID_r16_sbas_id_r16", err)
	}
	return nil
}

func (ie *SBAS_ID_r16_sbas_id_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode SBAS_ID_r16_sbas_id_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
