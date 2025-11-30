package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type DummyD struct {
	MaxNumberTxPortsPerResource    DummyD_maxNumberTxPortsPerResource `madatory`
	MaxNumberResources             int64                              `lb:1,ub:64,madatory`
	TotalNumberTxPorts             int64                              `lb:2,ub:256,madatory`
	ParameterLx                    int64                              `lb:2,ub:4,madatory`
	AmplitudeScalingType           DummyD_amplitudeScalingType        `madatory`
	AmplitudeSubsetRestriction     *DummyD_amplitudeSubsetRestriction `optional`
	MaxNumberCSI_RS_PerResourceSet int64                              `lb:1,ub:8,madatory`
}

func (ie *DummyD) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.AmplitudeSubsetRestriction != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.MaxNumberTxPortsPerResource.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberTxPortsPerResource", err)
	}
	if err = w.WriteInteger(ie.MaxNumberResources, &aper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return utils.WrapError("WriteInteger MaxNumberResources", err)
	}
	if err = w.WriteInteger(ie.TotalNumberTxPorts, &aper.Constraint{Lb: 2, Ub: 256}, false); err != nil {
		return utils.WrapError("WriteInteger TotalNumberTxPorts", err)
	}
	if err = w.WriteInteger(ie.ParameterLx, &aper.Constraint{Lb: 2, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger ParameterLx", err)
	}
	if err = ie.AmplitudeScalingType.Encode(w); err != nil {
		return utils.WrapError("Encode AmplitudeScalingType", err)
	}
	if ie.AmplitudeSubsetRestriction != nil {
		if err = ie.AmplitudeSubsetRestriction.Encode(w); err != nil {
			return utils.WrapError("Encode AmplitudeSubsetRestriction", err)
		}
	}
	if err = w.WriteInteger(ie.MaxNumberCSI_RS_PerResourceSet, &aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger MaxNumberCSI_RS_PerResourceSet", err)
	}
	return nil
}

func (ie *DummyD) Decode(r *aper.AperReader) error {
	var err error
	var AmplitudeSubsetRestrictionPresent bool
	if AmplitudeSubsetRestrictionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.MaxNumberTxPortsPerResource.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberTxPortsPerResource", err)
	}
	var tmp_int_MaxNumberResources int64
	if tmp_int_MaxNumberResources, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return utils.WrapError("ReadInteger MaxNumberResources", err)
	}
	ie.MaxNumberResources = tmp_int_MaxNumberResources
	var tmp_int_TotalNumberTxPorts int64
	if tmp_int_TotalNumberTxPorts, err = r.ReadInteger(&aper.Constraint{Lb: 2, Ub: 256}, false); err != nil {
		return utils.WrapError("ReadInteger TotalNumberTxPorts", err)
	}
	ie.TotalNumberTxPorts = tmp_int_TotalNumberTxPorts
	var tmp_int_ParameterLx int64
	if tmp_int_ParameterLx, err = r.ReadInteger(&aper.Constraint{Lb: 2, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger ParameterLx", err)
	}
	ie.ParameterLx = tmp_int_ParameterLx
	if err = ie.AmplitudeScalingType.Decode(r); err != nil {
		return utils.WrapError("Decode AmplitudeScalingType", err)
	}
	if AmplitudeSubsetRestrictionPresent {
		ie.AmplitudeSubsetRestriction = new(DummyD_amplitudeSubsetRestriction)
		if err = ie.AmplitudeSubsetRestriction.Decode(r); err != nil {
			return utils.WrapError("Decode AmplitudeSubsetRestriction", err)
		}
	}
	var tmp_int_MaxNumberCSI_RS_PerResourceSet int64
	if tmp_int_MaxNumberCSI_RS_PerResourceSet, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger MaxNumberCSI_RS_PerResourceSet", err)
	}
	ie.MaxNumberCSI_RS_PerResourceSet = tmp_int_MaxNumberCSI_RS_PerResourceSet
	return nil
}
