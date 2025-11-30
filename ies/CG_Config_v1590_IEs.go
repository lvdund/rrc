package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CG_Config_v1590_IEs struct {
	ScellFrequenciesSN_NR    []ARFCN_ValueNR      `lb:1,ub:maxNrofServingCells_1,optional`
	ScellFrequenciesSN_EUTRA []ARFCN_ValueEUTRA   `lb:1,ub:maxNrofServingCells_1,optional`
	NonCriticalExtension     *CG_Config_v1610_IEs `optional`
}

func (ie *CG_Config_v1590_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{len(ie.ScellFrequenciesSN_NR) > 0, len(ie.ScellFrequenciesSN_EUTRA) > 0, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.ScellFrequenciesSN_NR) > 0 {
		tmp_ScellFrequenciesSN_NR := utils.NewSequence[*ARFCN_ValueNR]([]*ARFCN_ValueNR{}, aper.Constraint{Lb: 1, Ub: maxNrofServingCells_1}, false)
		for _, i := range ie.ScellFrequenciesSN_NR {
			tmp_ScellFrequenciesSN_NR.Value = append(tmp_ScellFrequenciesSN_NR.Value, &i)
		}
		if err = tmp_ScellFrequenciesSN_NR.Encode(w); err != nil {
			return utils.WrapError("Encode ScellFrequenciesSN_NR", err)
		}
	}
	if len(ie.ScellFrequenciesSN_EUTRA) > 0 {
		tmp_ScellFrequenciesSN_EUTRA := utils.NewSequence[*ARFCN_ValueEUTRA]([]*ARFCN_ValueEUTRA{}, aper.Constraint{Lb: 1, Ub: maxNrofServingCells_1}, false)
		for _, i := range ie.ScellFrequenciesSN_EUTRA {
			tmp_ScellFrequenciesSN_EUTRA.Value = append(tmp_ScellFrequenciesSN_EUTRA.Value, &i)
		}
		if err = tmp_ScellFrequenciesSN_EUTRA.Encode(w); err != nil {
			return utils.WrapError("Encode ScellFrequenciesSN_EUTRA", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *CG_Config_v1590_IEs) Decode(r *aper.AperReader) error {
	var err error
	var ScellFrequenciesSN_NRPresent bool
	if ScellFrequenciesSN_NRPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ScellFrequenciesSN_EUTRAPresent bool
	if ScellFrequenciesSN_EUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if ScellFrequenciesSN_NRPresent {
		tmp_ScellFrequenciesSN_NR := utils.NewSequence[*ARFCN_ValueNR]([]*ARFCN_ValueNR{}, aper.Constraint{Lb: 1, Ub: maxNrofServingCells_1}, false)
		fn_ScellFrequenciesSN_NR := func() *ARFCN_ValueNR {
			return new(ARFCN_ValueNR)
		}
		if err = tmp_ScellFrequenciesSN_NR.Decode(r, fn_ScellFrequenciesSN_NR); err != nil {
			return utils.WrapError("Decode ScellFrequenciesSN_NR", err)
		}
		ie.ScellFrequenciesSN_NR = []ARFCN_ValueNR{}
		for _, i := range tmp_ScellFrequenciesSN_NR.Value {
			ie.ScellFrequenciesSN_NR = append(ie.ScellFrequenciesSN_NR, *i)
		}
	}
	if ScellFrequenciesSN_EUTRAPresent {
		tmp_ScellFrequenciesSN_EUTRA := utils.NewSequence[*ARFCN_ValueEUTRA]([]*ARFCN_ValueEUTRA{}, aper.Constraint{Lb: 1, Ub: maxNrofServingCells_1}, false)
		fn_ScellFrequenciesSN_EUTRA := func() *ARFCN_ValueEUTRA {
			return new(ARFCN_ValueEUTRA)
		}
		if err = tmp_ScellFrequenciesSN_EUTRA.Decode(r, fn_ScellFrequenciesSN_EUTRA); err != nil {
			return utils.WrapError("Decode ScellFrequenciesSN_EUTRA", err)
		}
		ie.ScellFrequenciesSN_EUTRA = []ARFCN_ValueEUTRA{}
		for _, i := range tmp_ScellFrequenciesSN_EUTRA.Value {
			ie.ScellFrequenciesSN_EUTRA = append(ie.ScellFrequenciesSN_EUTRA, *i)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(CG_Config_v1610_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
