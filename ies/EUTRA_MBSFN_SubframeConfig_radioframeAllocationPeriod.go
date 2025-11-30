package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	EUTRA_MBSFN_SubframeConfig_radioframeAllocationPeriod_Enum_n1  aper.Enumerated = 0
	EUTRA_MBSFN_SubframeConfig_radioframeAllocationPeriod_Enum_n2  aper.Enumerated = 1
	EUTRA_MBSFN_SubframeConfig_radioframeAllocationPeriod_Enum_n4  aper.Enumerated = 2
	EUTRA_MBSFN_SubframeConfig_radioframeAllocationPeriod_Enum_n8  aper.Enumerated = 3
	EUTRA_MBSFN_SubframeConfig_radioframeAllocationPeriod_Enum_n16 aper.Enumerated = 4
	EUTRA_MBSFN_SubframeConfig_radioframeAllocationPeriod_Enum_n32 aper.Enumerated = 5
)

type EUTRA_MBSFN_SubframeConfig_radioframeAllocationPeriod struct {
	Value aper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *EUTRA_MBSFN_SubframeConfig_radioframeAllocationPeriod) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode EUTRA_MBSFN_SubframeConfig_radioframeAllocationPeriod", err)
	}
	return nil
}

func (ie *EUTRA_MBSFN_SubframeConfig_radioframeAllocationPeriod) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode EUTRA_MBSFN_SubframeConfig_radioframeAllocationPeriod", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
