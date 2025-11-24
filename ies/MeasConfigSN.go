package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasConfigSN struct {
	MeasuredFrequenciesSN []NR_FreqInfo `lb:1,ub:maxMeasFreqsSN,optional`
}

func (ie *MeasConfigSN) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.MeasuredFrequenciesSN) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.MeasuredFrequenciesSN) > 0 {
		tmp_MeasuredFrequenciesSN := utils.NewSequence[*NR_FreqInfo]([]*NR_FreqInfo{}, uper.Constraint{Lb: 1, Ub: maxMeasFreqsSN}, false)
		for _, i := range ie.MeasuredFrequenciesSN {
			tmp_MeasuredFrequenciesSN.Value = append(tmp_MeasuredFrequenciesSN.Value, &i)
		}
		if err = tmp_MeasuredFrequenciesSN.Encode(w); err != nil {
			return utils.WrapError("Encode MeasuredFrequenciesSN", err)
		}
	}
	return nil
}

func (ie *MeasConfigSN) Decode(r *uper.UperReader) error {
	var err error
	var MeasuredFrequenciesSNPresent bool
	if MeasuredFrequenciesSNPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if MeasuredFrequenciesSNPresent {
		tmp_MeasuredFrequenciesSN := utils.NewSequence[*NR_FreqInfo]([]*NR_FreqInfo{}, uper.Constraint{Lb: 1, Ub: maxMeasFreqsSN}, false)
		fn_MeasuredFrequenciesSN := func() *NR_FreqInfo {
			return new(NR_FreqInfo)
		}
		if err = tmp_MeasuredFrequenciesSN.Decode(r, fn_MeasuredFrequenciesSN); err != nil {
			return utils.WrapError("Decode MeasuredFrequenciesSN", err)
		}
		ie.MeasuredFrequenciesSN = []NR_FreqInfo{}
		for _, i := range tmp_MeasuredFrequenciesSN.Value {
			ie.MeasuredFrequenciesSN = append(ie.MeasuredFrequenciesSN, *i)
		}
	}
	return nil
}
