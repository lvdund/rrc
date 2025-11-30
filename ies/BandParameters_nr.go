package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BandParameters_nr struct {
	BandNR                 FreqBandIndicatorNR  `madatory`
	Ca_BandwidthClassDL_NR *CA_BandwidthClassNR `optional`
	Ca_BandwidthClassUL_NR *CA_BandwidthClassNR `optional`
}

func (ie *BandParameters_nr) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Ca_BandwidthClassDL_NR != nil, ie.Ca_BandwidthClassUL_NR != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.BandNR.Encode(w); err != nil {
		return utils.WrapError("Encode BandNR", err)
	}
	if ie.Ca_BandwidthClassDL_NR != nil {
		if err = ie.Ca_BandwidthClassDL_NR.Encode(w); err != nil {
			return utils.WrapError("Encode Ca_BandwidthClassDL_NR", err)
		}
	}
	if ie.Ca_BandwidthClassUL_NR != nil {
		if err = ie.Ca_BandwidthClassUL_NR.Encode(w); err != nil {
			return utils.WrapError("Encode Ca_BandwidthClassUL_NR", err)
		}
	}
	return nil
}

func (ie *BandParameters_nr) Decode(r *aper.AperReader) error {
	var err error
	var Ca_BandwidthClassDL_NRPresent bool
	if Ca_BandwidthClassDL_NRPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Ca_BandwidthClassUL_NRPresent bool
	if Ca_BandwidthClassUL_NRPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.BandNR.Decode(r); err != nil {
		return utils.WrapError("Decode BandNR", err)
	}
	if Ca_BandwidthClassDL_NRPresent {
		ie.Ca_BandwidthClassDL_NR = new(CA_BandwidthClassNR)
		if err = ie.Ca_BandwidthClassDL_NR.Decode(r); err != nil {
			return utils.WrapError("Decode Ca_BandwidthClassDL_NR", err)
		}
	}
	if Ca_BandwidthClassUL_NRPresent {
		ie.Ca_BandwidthClassUL_NR = new(CA_BandwidthClassNR)
		if err = ie.Ca_BandwidthClassUL_NR.Decode(r); err != nil {
			return utils.WrapError("Decode Ca_BandwidthClassUL_NR", err)
		}
	}
	return nil
}
