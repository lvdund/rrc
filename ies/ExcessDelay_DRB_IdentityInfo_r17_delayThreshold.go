package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	ExcessDelay_DRB_IdentityInfo_r17_delayThreshold_Enum_ms0dot25 aper.Enumerated = 0
	ExcessDelay_DRB_IdentityInfo_r17_delayThreshold_Enum_ms0dot5  aper.Enumerated = 1
	ExcessDelay_DRB_IdentityInfo_r17_delayThreshold_Enum_ms1      aper.Enumerated = 2
	ExcessDelay_DRB_IdentityInfo_r17_delayThreshold_Enum_ms2      aper.Enumerated = 3
	ExcessDelay_DRB_IdentityInfo_r17_delayThreshold_Enum_ms4      aper.Enumerated = 4
	ExcessDelay_DRB_IdentityInfo_r17_delayThreshold_Enum_ms5      aper.Enumerated = 5
	ExcessDelay_DRB_IdentityInfo_r17_delayThreshold_Enum_ms10     aper.Enumerated = 6
	ExcessDelay_DRB_IdentityInfo_r17_delayThreshold_Enum_ms20     aper.Enumerated = 7
	ExcessDelay_DRB_IdentityInfo_r17_delayThreshold_Enum_ms30     aper.Enumerated = 8
	ExcessDelay_DRB_IdentityInfo_r17_delayThreshold_Enum_ms40     aper.Enumerated = 9
	ExcessDelay_DRB_IdentityInfo_r17_delayThreshold_Enum_ms50     aper.Enumerated = 10
	ExcessDelay_DRB_IdentityInfo_r17_delayThreshold_Enum_ms60     aper.Enumerated = 11
	ExcessDelay_DRB_IdentityInfo_r17_delayThreshold_Enum_ms70     aper.Enumerated = 12
	ExcessDelay_DRB_IdentityInfo_r17_delayThreshold_Enum_ms80     aper.Enumerated = 13
	ExcessDelay_DRB_IdentityInfo_r17_delayThreshold_Enum_ms90     aper.Enumerated = 14
	ExcessDelay_DRB_IdentityInfo_r17_delayThreshold_Enum_ms100    aper.Enumerated = 15
	ExcessDelay_DRB_IdentityInfo_r17_delayThreshold_Enum_ms150    aper.Enumerated = 16
	ExcessDelay_DRB_IdentityInfo_r17_delayThreshold_Enum_ms300    aper.Enumerated = 17
	ExcessDelay_DRB_IdentityInfo_r17_delayThreshold_Enum_ms500    aper.Enumerated = 18
)

type ExcessDelay_DRB_IdentityInfo_r17_delayThreshold struct {
	Value aper.Enumerated `lb:0,ub:18,madatory`
}

func (ie *ExcessDelay_DRB_IdentityInfo_r17_delayThreshold) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 18}, false); err != nil {
		return utils.WrapError("Encode ExcessDelay_DRB_IdentityInfo_r17_delayThreshold", err)
	}
	return nil
}

func (ie *ExcessDelay_DRB_IdentityInfo_r17_delayThreshold) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 18}, false); err != nil {
		return utils.WrapError("Decode ExcessDelay_DRB_IdentityInfo_r17_delayThreshold", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
