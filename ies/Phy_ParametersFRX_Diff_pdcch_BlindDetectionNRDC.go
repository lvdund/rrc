package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type Phy_ParametersFRX_Diff_pdcch_BlindDetectionNRDC struct {
	Pdcch_BlindDetectionMCG_UE int64 `lb:1,ub:15,madatory`
	Pdcch_BlindDetectionSCG_UE int64 `lb:1,ub:15,madatory`
}

func (ie *Phy_ParametersFRX_Diff_pdcch_BlindDetectionNRDC) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(ie.Pdcch_BlindDetectionMCG_UE, &aper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
		return utils.WrapError("WriteInteger Pdcch_BlindDetectionMCG_UE", err)
	}
	if err = w.WriteInteger(ie.Pdcch_BlindDetectionSCG_UE, &aper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
		return utils.WrapError("WriteInteger Pdcch_BlindDetectionSCG_UE", err)
	}
	return nil
}

func (ie *Phy_ParametersFRX_Diff_pdcch_BlindDetectionNRDC) Decode(r *aper.AperReader) error {
	var err error
	var tmp_int_Pdcch_BlindDetectionMCG_UE int64
	if tmp_int_Pdcch_BlindDetectionMCG_UE, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
		return utils.WrapError("ReadInteger Pdcch_BlindDetectionMCG_UE", err)
	}
	ie.Pdcch_BlindDetectionMCG_UE = tmp_int_Pdcch_BlindDetectionMCG_UE
	var tmp_int_Pdcch_BlindDetectionSCG_UE int64
	if tmp_int_Pdcch_BlindDetectionSCG_UE, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
		return utils.WrapError("ReadInteger Pdcch_BlindDetectionSCG_UE", err)
	}
	ie.Pdcch_BlindDetectionSCG_UE = tmp_int_Pdcch_BlindDetectionSCG_UE
	return nil
}
