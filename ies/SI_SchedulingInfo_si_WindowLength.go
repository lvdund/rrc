package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SI_SchedulingInfo_si_WindowLength_Enum_s5          aper.Enumerated = 0
	SI_SchedulingInfo_si_WindowLength_Enum_s10         aper.Enumerated = 1
	SI_SchedulingInfo_si_WindowLength_Enum_s20         aper.Enumerated = 2
	SI_SchedulingInfo_si_WindowLength_Enum_s40         aper.Enumerated = 3
	SI_SchedulingInfo_si_WindowLength_Enum_s80         aper.Enumerated = 4
	SI_SchedulingInfo_si_WindowLength_Enum_s160        aper.Enumerated = 5
	SI_SchedulingInfo_si_WindowLength_Enum_s320        aper.Enumerated = 6
	SI_SchedulingInfo_si_WindowLength_Enum_s640        aper.Enumerated = 7
	SI_SchedulingInfo_si_WindowLength_Enum_s1280       aper.Enumerated = 8
	SI_SchedulingInfo_si_WindowLength_Enum_s2560_v1710 aper.Enumerated = 9
	SI_SchedulingInfo_si_WindowLength_Enum_s5120_v1710 aper.Enumerated = 10
)

type SI_SchedulingInfo_si_WindowLength struct {
	Value aper.Enumerated `lb:0,ub:10,madatory`
}

func (ie *SI_SchedulingInfo_si_WindowLength) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 10}, false); err != nil {
		return utils.WrapError("Encode SI_SchedulingInfo_si_WindowLength", err)
	}
	return nil
}

func (ie *SI_SchedulingInfo_si_WindowLength) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 10}, false); err != nil {
		return utils.WrapError("Decode SI_SchedulingInfo_si_WindowLength", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
