package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MRDC_Parameters_v15g0 struct {
	SimultaneousRxTxInterBandENDCPerBandPair *SimultaneousRxTxPerBandPair `optional`
}

func (ie *MRDC_Parameters_v15g0) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.SimultaneousRxTxInterBandENDCPerBandPair != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SimultaneousRxTxInterBandENDCPerBandPair != nil {
		if err = ie.SimultaneousRxTxInterBandENDCPerBandPair.Encode(w); err != nil {
			return utils.WrapError("Encode SimultaneousRxTxInterBandENDCPerBandPair", err)
		}
	}
	return nil
}

func (ie *MRDC_Parameters_v15g0) Decode(r *uper.UperReader) error {
	var err error
	var SimultaneousRxTxInterBandENDCPerBandPairPresent bool
	if SimultaneousRxTxInterBandENDCPerBandPairPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if SimultaneousRxTxInterBandENDCPerBandPairPresent {
		ie.SimultaneousRxTxInterBandENDCPerBandPair = new(SimultaneousRxTxPerBandPair)
		if err = ie.SimultaneousRxTxInterBandENDCPerBandPair.Decode(r); err != nil {
			return utils.WrapError("Decode SimultaneousRxTxInterBandENDCPerBandPair", err)
		}
	}
	return nil
}
