package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MRDC_Parameters_v1590 struct {
	InterBandContiguousMRDC *MRDC_Parameters_v1590_interBandContiguousMRDC `optional`
}

func (ie *MRDC_Parameters_v1590) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.InterBandContiguousMRDC != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.InterBandContiguousMRDC != nil {
		if err = ie.InterBandContiguousMRDC.Encode(w); err != nil {
			return utils.WrapError("Encode InterBandContiguousMRDC", err)
		}
	}
	return nil
}

func (ie *MRDC_Parameters_v1590) Decode(r *aper.AperReader) error {
	var err error
	var InterBandContiguousMRDCPresent bool
	if InterBandContiguousMRDCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if InterBandContiguousMRDCPresent {
		ie.InterBandContiguousMRDC = new(MRDC_Parameters_v1590_interBandContiguousMRDC)
		if err = ie.InterBandContiguousMRDC.Decode(r); err != nil {
			return utils.WrapError("Decode InterBandContiguousMRDC", err)
		}
	}
	return nil
}
