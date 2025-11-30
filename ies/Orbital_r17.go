package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type Orbital_r17 struct {
	SemiMajorAxis_r17 int64 `lb:0,ub:8589934591,madatory`
	Eccentricity_r17  int64 `lb:0,ub:1048575,madatory`
	Periapsis_r17     int64 `lb:0,ub:268435455,madatory`
	Longitude_r17     int64 `lb:0,ub:268435455,madatory`
	Inclination_r17   int64 `lb:-67108864,ub:67108863,madatory`
	MeanAnomaly_r17   int64 `lb:0,ub:268435455,madatory`
}

func (ie *Orbital_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(ie.SemiMajorAxis_r17, &aper.Constraint{Lb: 0, Ub: 8589934591}, false); err != nil {
		return utils.WrapError("WriteInteger SemiMajorAxis_r17", err)
	}
	if err = w.WriteInteger(ie.Eccentricity_r17, &aper.Constraint{Lb: 0, Ub: 1048575}, false); err != nil {
		return utils.WrapError("WriteInteger Eccentricity_r17", err)
	}
	if err = w.WriteInteger(ie.Periapsis_r17, &aper.Constraint{Lb: 0, Ub: 268435455}, false); err != nil {
		return utils.WrapError("WriteInteger Periapsis_r17", err)
	}
	if err = w.WriteInteger(ie.Longitude_r17, &aper.Constraint{Lb: 0, Ub: 268435455}, false); err != nil {
		return utils.WrapError("WriteInteger Longitude_r17", err)
	}
	if err = w.WriteInteger(ie.Inclination_r17, &aper.Constraint{Lb: -67108864, Ub: 67108863}, false); err != nil {
		return utils.WrapError("WriteInteger Inclination_r17", err)
	}
	if err = w.WriteInteger(ie.MeanAnomaly_r17, &aper.Constraint{Lb: 0, Ub: 268435455}, false); err != nil {
		return utils.WrapError("WriteInteger MeanAnomaly_r17", err)
	}
	return nil
}

func (ie *Orbital_r17) Decode(r *aper.AperReader) error {
	var err error
	var tmp_int_SemiMajorAxis_r17 int64
	if tmp_int_SemiMajorAxis_r17, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 8589934591}, false); err != nil {
		return utils.WrapError("ReadInteger SemiMajorAxis_r17", err)
	}
	ie.SemiMajorAxis_r17 = tmp_int_SemiMajorAxis_r17
	var tmp_int_Eccentricity_r17 int64
	if tmp_int_Eccentricity_r17, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 1048575}, false); err != nil {
		return utils.WrapError("ReadInteger Eccentricity_r17", err)
	}
	ie.Eccentricity_r17 = tmp_int_Eccentricity_r17
	var tmp_int_Periapsis_r17 int64
	if tmp_int_Periapsis_r17, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 268435455}, false); err != nil {
		return utils.WrapError("ReadInteger Periapsis_r17", err)
	}
	ie.Periapsis_r17 = tmp_int_Periapsis_r17
	var tmp_int_Longitude_r17 int64
	if tmp_int_Longitude_r17, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 268435455}, false); err != nil {
		return utils.WrapError("ReadInteger Longitude_r17", err)
	}
	ie.Longitude_r17 = tmp_int_Longitude_r17
	var tmp_int_Inclination_r17 int64
	if tmp_int_Inclination_r17, err = r.ReadInteger(&aper.Constraint{Lb: -67108864, Ub: 67108863}, false); err != nil {
		return utils.WrapError("ReadInteger Inclination_r17", err)
	}
	ie.Inclination_r17 = tmp_int_Inclination_r17
	var tmp_int_MeanAnomaly_r17 int64
	if tmp_int_MeanAnomaly_r17, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 268435455}, false); err != nil {
		return utils.WrapError("ReadInteger MeanAnomaly_r17", err)
	}
	ie.MeanAnomaly_r17 = tmp_int_MeanAnomaly_r17
	return nil
}
