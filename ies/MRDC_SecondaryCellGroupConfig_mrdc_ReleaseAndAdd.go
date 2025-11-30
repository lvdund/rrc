package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MRDC_SecondaryCellGroupConfig_mrdc_ReleaseAndAdd_Enum_true aper.Enumerated = 0
)

type MRDC_SecondaryCellGroupConfig_mrdc_ReleaseAndAdd struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *MRDC_SecondaryCellGroupConfig_mrdc_ReleaseAndAdd) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode MRDC_SecondaryCellGroupConfig_mrdc_ReleaseAndAdd", err)
	}
	return nil
}

func (ie *MRDC_SecondaryCellGroupConfig_mrdc_ReleaseAndAdd) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode MRDC_SecondaryCellGroupConfig_mrdc_ReleaseAndAdd", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
