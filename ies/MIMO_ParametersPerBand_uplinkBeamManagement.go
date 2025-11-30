package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_uplinkBeamManagement struct {
	MaxNumberSRS_ResourcePerSet_BM MIMO_ParametersPerBand_uplinkBeamManagement_maxNumberSRS_ResourcePerSet_BM `madatory`
	MaxNumberSRS_ResourceSet       int64                                                                      `lb:1,ub:8,madatory`
}

func (ie *MIMO_ParametersPerBand_uplinkBeamManagement) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.MaxNumberSRS_ResourcePerSet_BM.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberSRS_ResourcePerSet_BM", err)
	}
	if err = w.WriteInteger(ie.MaxNumberSRS_ResourceSet, &aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger MaxNumberSRS_ResourceSet", err)
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_uplinkBeamManagement) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.MaxNumberSRS_ResourcePerSet_BM.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberSRS_ResourcePerSet_BM", err)
	}
	var tmp_int_MaxNumberSRS_ResourceSet int64
	if tmp_int_MaxNumberSRS_ResourceSet, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger MaxNumberSRS_ResourceSet", err)
	}
	ie.MaxNumberSRS_ResourceSet = tmp_int_MaxNumberSRS_ResourceSet
	return nil
}
