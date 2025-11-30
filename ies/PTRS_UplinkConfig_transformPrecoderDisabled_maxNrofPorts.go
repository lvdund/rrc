package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PTRS_UplinkConfig_transformPrecoderDisabled_maxNrofPorts_Enum_n1 aper.Enumerated = 0
	PTRS_UplinkConfig_transformPrecoderDisabled_maxNrofPorts_Enum_n2 aper.Enumerated = 1
)

type PTRS_UplinkConfig_transformPrecoderDisabled_maxNrofPorts struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *PTRS_UplinkConfig_transformPrecoderDisabled_maxNrofPorts) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode PTRS_UplinkConfig_transformPrecoderDisabled_maxNrofPorts", err)
	}
	return nil
}

func (ie *PTRS_UplinkConfig_transformPrecoderDisabled_maxNrofPorts) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode PTRS_UplinkConfig_transformPrecoderDisabled_maxNrofPorts", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
