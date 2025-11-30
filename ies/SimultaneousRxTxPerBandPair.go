package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SimultaneousRxTxPerBandPair struct {
	Value aper.BitString `lb:3,ub:496,madatory`
}

func (ie *SimultaneousRxTxPerBandPair) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteBitString(ie.Value.Bytes, uint(ie.Value.NumBits), &aper.Constraint{Lb: 3, Ub: 496}, false); err != nil {
		return utils.WrapError("Encode SimultaneousRxTxPerBandPair", err)
	}
	return nil
}

func (ie *SimultaneousRxTxPerBandPair) Decode(r *aper.AperReader) error {
	var err error
	var v []byte
	var n uint
	if v, n, err = r.ReadBitString(&aper.Constraint{Lb: 3, Ub: 496}, false); err != nil {
		return utils.WrapError("Decode SimultaneousRxTxPerBandPair", err)
	}
	ie.Value = aper.BitString{
		Bytes:   v,
		NumBits: uint64(n),
	}
	return nil
}
