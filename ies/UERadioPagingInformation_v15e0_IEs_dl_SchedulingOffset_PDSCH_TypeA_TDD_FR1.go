package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	UERadioPagingInformation_v15e0_IEs_dl_SchedulingOffset_PDSCH_TypeA_TDD_FR1_Enum_supported aper.Enumerated = 0
)

type UERadioPagingInformation_v15e0_IEs_dl_SchedulingOffset_PDSCH_TypeA_TDD_FR1 struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *UERadioPagingInformation_v15e0_IEs_dl_SchedulingOffset_PDSCH_TypeA_TDD_FR1) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode UERadioPagingInformation_v15e0_IEs_dl_SchedulingOffset_PDSCH_TypeA_TDD_FR1", err)
	}
	return nil
}

func (ie *UERadioPagingInformation_v15e0_IEs_dl_SchedulingOffset_PDSCH_TypeA_TDD_FR1) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode UERadioPagingInformation_v15e0_IEs_dl_SchedulingOffset_PDSCH_TypeA_TDD_FR1", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
