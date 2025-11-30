package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SRS_Resources struct {
	MaxNumberAperiodicSRS_PerBWP              SRS_Resources_maxNumberAperiodicSRS_PerBWP      `madatory`
	MaxNumberAperiodicSRS_PerBWP_PerSlot      int64                                           `lb:1,ub:6,madatory`
	MaxNumberPeriodicSRS_PerBWP               SRS_Resources_maxNumberPeriodicSRS_PerBWP       `madatory`
	MaxNumberPeriodicSRS_PerBWP_PerSlot       int64                                           `lb:1,ub:6,madatory`
	MaxNumberSemiPersistentSRS_PerBWP         SRS_Resources_maxNumberSemiPersistentSRS_PerBWP `madatory`
	MaxNumberSemiPersistentSRS_PerBWP_PerSlot int64                                           `lb:1,ub:6,madatory`
	MaxNumberSRS_Ports_PerResource            SRS_Resources_maxNumberSRS_Ports_PerResource    `madatory`
}

func (ie *SRS_Resources) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.MaxNumberAperiodicSRS_PerBWP.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberAperiodicSRS_PerBWP", err)
	}
	if err = w.WriteInteger(ie.MaxNumberAperiodicSRS_PerBWP_PerSlot, &aper.Constraint{Lb: 1, Ub: 6}, false); err != nil {
		return utils.WrapError("WriteInteger MaxNumberAperiodicSRS_PerBWP_PerSlot", err)
	}
	if err = ie.MaxNumberPeriodicSRS_PerBWP.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberPeriodicSRS_PerBWP", err)
	}
	if err = w.WriteInteger(ie.MaxNumberPeriodicSRS_PerBWP_PerSlot, &aper.Constraint{Lb: 1, Ub: 6}, false); err != nil {
		return utils.WrapError("WriteInteger MaxNumberPeriodicSRS_PerBWP_PerSlot", err)
	}
	if err = ie.MaxNumberSemiPersistentSRS_PerBWP.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberSemiPersistentSRS_PerBWP", err)
	}
	if err = w.WriteInteger(ie.MaxNumberSemiPersistentSRS_PerBWP_PerSlot, &aper.Constraint{Lb: 1, Ub: 6}, false); err != nil {
		return utils.WrapError("WriteInteger MaxNumberSemiPersistentSRS_PerBWP_PerSlot", err)
	}
	if err = ie.MaxNumberSRS_Ports_PerResource.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberSRS_Ports_PerResource", err)
	}
	return nil
}

func (ie *SRS_Resources) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.MaxNumberAperiodicSRS_PerBWP.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberAperiodicSRS_PerBWP", err)
	}
	var tmp_int_MaxNumberAperiodicSRS_PerBWP_PerSlot int64
	if tmp_int_MaxNumberAperiodicSRS_PerBWP_PerSlot, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 6}, false); err != nil {
		return utils.WrapError("ReadInteger MaxNumberAperiodicSRS_PerBWP_PerSlot", err)
	}
	ie.MaxNumberAperiodicSRS_PerBWP_PerSlot = tmp_int_MaxNumberAperiodicSRS_PerBWP_PerSlot
	if err = ie.MaxNumberPeriodicSRS_PerBWP.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberPeriodicSRS_PerBWP", err)
	}
	var tmp_int_MaxNumberPeriodicSRS_PerBWP_PerSlot int64
	if tmp_int_MaxNumberPeriodicSRS_PerBWP_PerSlot, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 6}, false); err != nil {
		return utils.WrapError("ReadInteger MaxNumberPeriodicSRS_PerBWP_PerSlot", err)
	}
	ie.MaxNumberPeriodicSRS_PerBWP_PerSlot = tmp_int_MaxNumberPeriodicSRS_PerBWP_PerSlot
	if err = ie.MaxNumberSemiPersistentSRS_PerBWP.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberSemiPersistentSRS_PerBWP", err)
	}
	var tmp_int_MaxNumberSemiPersistentSRS_PerBWP_PerSlot int64
	if tmp_int_MaxNumberSemiPersistentSRS_PerBWP_PerSlot, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 6}, false); err != nil {
		return utils.WrapError("ReadInteger MaxNumberSemiPersistentSRS_PerBWP_PerSlot", err)
	}
	ie.MaxNumberSemiPersistentSRS_PerBWP_PerSlot = tmp_int_MaxNumberSemiPersistentSRS_PerBWP_PerSlot
	if err = ie.MaxNumberSRS_Ports_PerResource.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberSRS_Ports_PerResource", err)
	}
	return nil
}
