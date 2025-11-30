package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CGI_InfoEUTRA_freqBandIndicatorPriority_Enum_true aper.Enumerated = 0
)

type CGI_InfoEUTRA_freqBandIndicatorPriority struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *CGI_InfoEUTRA_freqBandIndicatorPriority) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode CGI_InfoEUTRA_freqBandIndicatorPriority", err)
	}
	return nil
}

func (ie *CGI_InfoEUTRA_freqBandIndicatorPriority) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode CGI_InfoEUTRA_freqBandIndicatorPriority", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
