package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination_v1590 struct {
	SupportedBandwidthCombinationSetIntraENDC *aper.BitString       `lb:1,ub:32,optional`
	Mrdc_Parameters_v1590                     MRDC_Parameters_v1590 `madatory`
}

func (ie *BandCombination_v1590) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.SupportedBandwidthCombinationSetIntraENDC != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SupportedBandwidthCombinationSetIntraENDC != nil {
		if err = w.WriteBitString(ie.SupportedBandwidthCombinationSetIntraENDC.Bytes, uint(ie.SupportedBandwidthCombinationSetIntraENDC.NumBits), &aper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Encode SupportedBandwidthCombinationSetIntraENDC", err)
		}
	}
	if err = ie.Mrdc_Parameters_v1590.Encode(w); err != nil {
		return utils.WrapError("Encode Mrdc_Parameters_v1590", err)
	}
	return nil
}

func (ie *BandCombination_v1590) Decode(r *aper.AperReader) error {
	var err error
	var SupportedBandwidthCombinationSetIntraENDCPresent bool
	if SupportedBandwidthCombinationSetIntraENDCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if SupportedBandwidthCombinationSetIntraENDCPresent {
		var tmp_bs_SupportedBandwidthCombinationSetIntraENDC []byte
		var n_SupportedBandwidthCombinationSetIntraENDC uint
		if tmp_bs_SupportedBandwidthCombinationSetIntraENDC, n_SupportedBandwidthCombinationSetIntraENDC, err = r.ReadBitString(&aper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode SupportedBandwidthCombinationSetIntraENDC", err)
		}
		tmp_bitstring := aper.BitString{
			Bytes:   tmp_bs_SupportedBandwidthCombinationSetIntraENDC,
			NumBits: uint64(n_SupportedBandwidthCombinationSetIntraENDC),
		}
		ie.SupportedBandwidthCombinationSetIntraENDC = &tmp_bitstring
	}
	if err = ie.Mrdc_Parameters_v1590.Decode(r); err != nil {
		return utils.WrapError("Decode Mrdc_Parameters_v1590", err)
	}
	return nil
}
