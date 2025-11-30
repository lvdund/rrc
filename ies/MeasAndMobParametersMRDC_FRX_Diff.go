package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasAndMobParametersMRDC_FRX_Diff struct {
	SimultaneousRxDataSSB_DiffNumerology *MeasAndMobParametersMRDC_FRX_Diff_simultaneousRxDataSSB_DiffNumerology `optional`
}

func (ie *MeasAndMobParametersMRDC_FRX_Diff) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.SimultaneousRxDataSSB_DiffNumerology != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SimultaneousRxDataSSB_DiffNumerology != nil {
		if err = ie.SimultaneousRxDataSSB_DiffNumerology.Encode(w); err != nil {
			return utils.WrapError("Encode SimultaneousRxDataSSB_DiffNumerology", err)
		}
	}
	return nil
}

func (ie *MeasAndMobParametersMRDC_FRX_Diff) Decode(r *aper.AperReader) error {
	var err error
	var SimultaneousRxDataSSB_DiffNumerologyPresent bool
	if SimultaneousRxDataSSB_DiffNumerologyPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if SimultaneousRxDataSSB_DiffNumerologyPresent {
		ie.SimultaneousRxDataSSB_DiffNumerology = new(MeasAndMobParametersMRDC_FRX_Diff_simultaneousRxDataSSB_DiffNumerology)
		if err = ie.SimultaneousRxDataSSB_DiffNumerology.Decode(r); err != nil {
			return utils.WrapError("Decode SimultaneousRxDataSSB_DiffNumerology", err)
		}
	}
	return nil
}
