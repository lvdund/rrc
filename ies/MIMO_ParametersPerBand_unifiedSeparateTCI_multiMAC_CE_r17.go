package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17 struct {
	MinBeamApplicationTime_r17  MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17_minBeamApplicationTime_r17 `madatory`
	MaxActivatedDL_TCIPerCC_r17 int64                                                                                `lb:2,ub:8,madatory`
	MaxActivatedUL_TCIPerCC_r17 int64                                                                                `lb:2,ub:8,madatory`
}

func (ie *MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.MinBeamApplicationTime_r17.Encode(w); err != nil {
		return utils.WrapError("Encode MinBeamApplicationTime_r17", err)
	}
	if err = w.WriteInteger(ie.MaxActivatedDL_TCIPerCC_r17, &aper.Constraint{Lb: 2, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger MaxActivatedDL_TCIPerCC_r17", err)
	}
	if err = w.WriteInteger(ie.MaxActivatedUL_TCIPerCC_r17, &aper.Constraint{Lb: 2, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger MaxActivatedUL_TCIPerCC_r17", err)
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.MinBeamApplicationTime_r17.Decode(r); err != nil {
		return utils.WrapError("Decode MinBeamApplicationTime_r17", err)
	}
	var tmp_int_MaxActivatedDL_TCIPerCC_r17 int64
	if tmp_int_MaxActivatedDL_TCIPerCC_r17, err = r.ReadInteger(&aper.Constraint{Lb: 2, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger MaxActivatedDL_TCIPerCC_r17", err)
	}
	ie.MaxActivatedDL_TCIPerCC_r17 = tmp_int_MaxActivatedDL_TCIPerCC_r17
	var tmp_int_MaxActivatedUL_TCIPerCC_r17 int64
	if tmp_int_MaxActivatedUL_TCIPerCC_r17, err = r.ReadInteger(&aper.Constraint{Lb: 2, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger MaxActivatedUL_TCIPerCC_r17", err)
	}
	ie.MaxActivatedUL_TCIPerCC_r17 = tmp_int_MaxActivatedUL_TCIPerCC_r17
	return nil
}
