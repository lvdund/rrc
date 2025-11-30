package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CA_BandwidthClassEUTRA_Enum_a aper.Enumerated = 0
	CA_BandwidthClassEUTRA_Enum_b aper.Enumerated = 1
	CA_BandwidthClassEUTRA_Enum_c aper.Enumerated = 2
	CA_BandwidthClassEUTRA_Enum_d aper.Enumerated = 3
	CA_BandwidthClassEUTRA_Enum_e aper.Enumerated = 4
	CA_BandwidthClassEUTRA_Enum_f aper.Enumerated = 5
)

type CA_BandwidthClassEUTRA struct {
	Value aper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *CA_BandwidthClassEUTRA) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode CA_BandwidthClassEUTRA", err)
	}
	return nil
}

func (ie *CA_BandwidthClassEUTRA) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode CA_BandwidthClassEUTRA", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
