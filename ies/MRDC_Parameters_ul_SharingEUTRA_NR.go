package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MRDC_Parameters_ul_SharingEUTRA_NR_Enum_tdm  aper.Enumerated = 0
	MRDC_Parameters_ul_SharingEUTRA_NR_Enum_fdm  aper.Enumerated = 1
	MRDC_Parameters_ul_SharingEUTRA_NR_Enum_both aper.Enumerated = 2
)

type MRDC_Parameters_ul_SharingEUTRA_NR struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *MRDC_Parameters_ul_SharingEUTRA_NR) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode MRDC_Parameters_ul_SharingEUTRA_NR", err)
	}
	return nil
}

func (ie *MRDC_Parameters_ul_SharingEUTRA_NR) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode MRDC_Parameters_ul_SharingEUTRA_NR", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
