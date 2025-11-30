package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MRDC_Parameters_ul_SwitchingTimeEUTRA_NR_Enum_type1 aper.Enumerated = 0
	MRDC_Parameters_ul_SwitchingTimeEUTRA_NR_Enum_type2 aper.Enumerated = 1
)

type MRDC_Parameters_ul_SwitchingTimeEUTRA_NR struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *MRDC_Parameters_ul_SwitchingTimeEUTRA_NR) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode MRDC_Parameters_ul_SwitchingTimeEUTRA_NR", err)
	}
	return nil
}

func (ie *MRDC_Parameters_ul_SwitchingTimeEUTRA_NR) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode MRDC_Parameters_ul_SwitchingTimeEUTRA_NR", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
