package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNR_v15g0 struct {
	SimultaneousRxTxInterBandCAPerBandPair *SimultaneousRxTxPerBandPair `optional`
	SimultaneousRxTxSULPerBandPair         *SimultaneousRxTxPerBandPair `optional`
}

func (ie *CA_ParametersNR_v15g0) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.SimultaneousRxTxInterBandCAPerBandPair != nil, ie.SimultaneousRxTxSULPerBandPair != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SimultaneousRxTxInterBandCAPerBandPair != nil {
		if err = ie.SimultaneousRxTxInterBandCAPerBandPair.Encode(w); err != nil {
			return utils.WrapError("Encode SimultaneousRxTxInterBandCAPerBandPair", err)
		}
	}
	if ie.SimultaneousRxTxSULPerBandPair != nil {
		if err = ie.SimultaneousRxTxSULPerBandPair.Encode(w); err != nil {
			return utils.WrapError("Encode SimultaneousRxTxSULPerBandPair", err)
		}
	}
	return nil
}

func (ie *CA_ParametersNR_v15g0) Decode(r *aper.AperReader) error {
	var err error
	var SimultaneousRxTxInterBandCAPerBandPairPresent bool
	if SimultaneousRxTxInterBandCAPerBandPairPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SimultaneousRxTxSULPerBandPairPresent bool
	if SimultaneousRxTxSULPerBandPairPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if SimultaneousRxTxInterBandCAPerBandPairPresent {
		ie.SimultaneousRxTxInterBandCAPerBandPair = new(SimultaneousRxTxPerBandPair)
		if err = ie.SimultaneousRxTxInterBandCAPerBandPair.Decode(r); err != nil {
			return utils.WrapError("Decode SimultaneousRxTxInterBandCAPerBandPair", err)
		}
	}
	if SimultaneousRxTxSULPerBandPairPresent {
		ie.SimultaneousRxTxSULPerBandPair = new(SimultaneousRxTxPerBandPair)
		if err = ie.SimultaneousRxTxSULPerBandPair.Decode(r); err != nil {
			return utils.WrapError("Decode SimultaneousRxTxSULPerBandPair", err)
		}
	}
	return nil
}
