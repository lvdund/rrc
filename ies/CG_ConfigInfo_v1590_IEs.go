package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_ConfigInfo_v1590_IEs struct {
	ServFrequenciesMN_NR []ARFCN_ValueNR          `lb:1,ub:maxNrofServingCells_1,optional`
	NonCriticalExtension *CG_ConfigInfo_v1610_IEs `optional`
}

func (ie *CG_ConfigInfo_v1590_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.ServFrequenciesMN_NR) > 0, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.ServFrequenciesMN_NR) > 0 {
		tmp_ServFrequenciesMN_NR := utils.NewSequence[*ARFCN_ValueNR]([]*ARFCN_ValueNR{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCells_1}, false)
		for _, i := range ie.ServFrequenciesMN_NR {
			tmp_ServFrequenciesMN_NR.Value = append(tmp_ServFrequenciesMN_NR.Value, &i)
		}
		if err = tmp_ServFrequenciesMN_NR.Encode(w); err != nil {
			return utils.WrapError("Encode ServFrequenciesMN_NR", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *CG_ConfigInfo_v1590_IEs) Decode(r *uper.UperReader) error {
	var err error
	var ServFrequenciesMN_NRPresent bool
	if ServFrequenciesMN_NRPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if ServFrequenciesMN_NRPresent {
		tmp_ServFrequenciesMN_NR := utils.NewSequence[*ARFCN_ValueNR]([]*ARFCN_ValueNR{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCells_1}, false)
		fn_ServFrequenciesMN_NR := func() *ARFCN_ValueNR {
			return new(ARFCN_ValueNR)
		}
		if err = tmp_ServFrequenciesMN_NR.Decode(r, fn_ServFrequenciesMN_NR); err != nil {
			return utils.WrapError("Decode ServFrequenciesMN_NR", err)
		}
		ie.ServFrequenciesMN_NR = []ARFCN_ValueNR{}
		for _, i := range tmp_ServFrequenciesMN_NR.Value {
			ie.ServFrequenciesMN_NR = append(ie.ServFrequenciesMN_NR, *i)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(CG_ConfigInfo_v1610_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
