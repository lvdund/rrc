package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CandidateBeamRS_r16 struct {
	CandidateBeamConfig_r16 CandidateBeamRS_r16_candidateBeamConfig_r16 `madatory`
	ServingCellId           *ServCellIndex                              `optional`
}

func (ie *CandidateBeamRS_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.ServingCellId != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.CandidateBeamConfig_r16.Encode(w); err != nil {
		return utils.WrapError("Encode CandidateBeamConfig_r16", err)
	}
	if ie.ServingCellId != nil {
		if err = ie.ServingCellId.Encode(w); err != nil {
			return utils.WrapError("Encode ServingCellId", err)
		}
	}
	return nil
}

func (ie *CandidateBeamRS_r16) Decode(r *aper.AperReader) error {
	var err error
	var ServingCellIdPresent bool
	if ServingCellIdPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.CandidateBeamConfig_r16.Decode(r); err != nil {
		return utils.WrapError("Decode CandidateBeamConfig_r16", err)
	}
	if ServingCellIdPresent {
		ie.ServingCellId = new(ServCellIndex)
		if err = ie.ServingCellId.Decode(r); err != nil {
			return utils.WrapError("Decode ServingCellId", err)
		}
	}
	return nil
}
