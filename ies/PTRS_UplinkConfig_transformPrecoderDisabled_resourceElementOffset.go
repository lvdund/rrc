package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PTRS_UplinkConfig_transformPrecoderDisabled_resourceElementOffset_Enum_offset01 aper.Enumerated = 0
	PTRS_UplinkConfig_transformPrecoderDisabled_resourceElementOffset_Enum_offset10 aper.Enumerated = 1
	PTRS_UplinkConfig_transformPrecoderDisabled_resourceElementOffset_Enum_offset11 aper.Enumerated = 2
)

type PTRS_UplinkConfig_transformPrecoderDisabled_resourceElementOffset struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *PTRS_UplinkConfig_transformPrecoderDisabled_resourceElementOffset) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode PTRS_UplinkConfig_transformPrecoderDisabled_resourceElementOffset", err)
	}
	return nil
}

func (ie *PTRS_UplinkConfig_transformPrecoderDisabled_resourceElementOffset) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode PTRS_UplinkConfig_transformPrecoderDisabled_resourceElementOffset", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
