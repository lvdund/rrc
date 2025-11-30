package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type NR_FreqInfo struct {
	MeasuredFrequency *ARFCN_ValueNR `optional`
}

func (ie *NR_FreqInfo) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.MeasuredFrequency != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MeasuredFrequency != nil {
		if err = ie.MeasuredFrequency.Encode(w); err != nil {
			return utils.WrapError("Encode MeasuredFrequency", err)
		}
	}
	return nil
}

func (ie *NR_FreqInfo) Decode(r *aper.AperReader) error {
	var err error
	var MeasuredFrequencyPresent bool
	if MeasuredFrequencyPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if MeasuredFrequencyPresent {
		ie.MeasuredFrequency = new(ARFCN_ValueNR)
		if err = ie.MeasuredFrequency.Decode(r); err != nil {
			return utils.WrapError("Decode MeasuredFrequency", err)
		}
	}
	return nil
}
