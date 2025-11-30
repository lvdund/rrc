package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PTRS_UplinkConfig_transformPrecoderDisabled_ptrs_Power_Enum_p00 aper.Enumerated = 0
	PTRS_UplinkConfig_transformPrecoderDisabled_ptrs_Power_Enum_p01 aper.Enumerated = 1
	PTRS_UplinkConfig_transformPrecoderDisabled_ptrs_Power_Enum_p10 aper.Enumerated = 2
	PTRS_UplinkConfig_transformPrecoderDisabled_ptrs_Power_Enum_p11 aper.Enumerated = 3
)

type PTRS_UplinkConfig_transformPrecoderDisabled_ptrs_Power struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *PTRS_UplinkConfig_transformPrecoderDisabled_ptrs_Power) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode PTRS_UplinkConfig_transformPrecoderDisabled_ptrs_Power", err)
	}
	return nil
}

func (ie *PTRS_UplinkConfig_transformPrecoderDisabled_ptrs_Power) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode PTRS_UplinkConfig_transformPrecoderDisabled_ptrs_Power", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
