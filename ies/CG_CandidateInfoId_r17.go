package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CG_CandidateInfoId_r17 struct {
	SsbFrequency_r17 ARFCN_ValueNR `madatory`
	PhysCellId_r17   PhysCellId    `madatory`
}

func (ie *CG_CandidateInfoId_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.SsbFrequency_r17.Encode(w); err != nil {
		return utils.WrapError("Encode SsbFrequency_r17", err)
	}
	if err = ie.PhysCellId_r17.Encode(w); err != nil {
		return utils.WrapError("Encode PhysCellId_r17", err)
	}
	return nil
}

func (ie *CG_CandidateInfoId_r17) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.SsbFrequency_r17.Decode(r); err != nil {
		return utils.WrapError("Decode SsbFrequency_r17", err)
	}
	if err = ie.PhysCellId_r17.Decode(r); err != nil {
		return utils.WrapError("Decode PhysCellId_r17", err)
	}
	return nil
}
