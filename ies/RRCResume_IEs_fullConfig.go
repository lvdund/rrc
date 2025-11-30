package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RRCResume_IEs_fullConfig_Enum_true aper.Enumerated = 0
)

type RRCResume_IEs_fullConfig struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *RRCResume_IEs_fullConfig) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode RRCResume_IEs_fullConfig", err)
	}
	return nil
}

func (ie *RRCResume_IEs_fullConfig) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode RRCResume_IEs_fullConfig", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
