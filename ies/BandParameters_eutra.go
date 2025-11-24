package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandParameters_eutra struct {
	BandEUTRA                 FreqBandIndicatorEUTRA  `madatory`
	Ca_BandwidthClassDL_EUTRA *CA_BandwidthClassEUTRA `optional`
	Ca_BandwidthClassUL_EUTRA *CA_BandwidthClassEUTRA `optional`
}

func (ie *BandParameters_eutra) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Ca_BandwidthClassDL_EUTRA != nil, ie.Ca_BandwidthClassUL_EUTRA != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.BandEUTRA.Encode(w); err != nil {
		return utils.WrapError("Encode BandEUTRA", err)
	}
	if ie.Ca_BandwidthClassDL_EUTRA != nil {
		if err = ie.Ca_BandwidthClassDL_EUTRA.Encode(w); err != nil {
			return utils.WrapError("Encode Ca_BandwidthClassDL_EUTRA", err)
		}
	}
	if ie.Ca_BandwidthClassUL_EUTRA != nil {
		if err = ie.Ca_BandwidthClassUL_EUTRA.Encode(w); err != nil {
			return utils.WrapError("Encode Ca_BandwidthClassUL_EUTRA", err)
		}
	}
	return nil
}

func (ie *BandParameters_eutra) Decode(r *uper.UperReader) error {
	var err error
	var Ca_BandwidthClassDL_EUTRAPresent bool
	if Ca_BandwidthClassDL_EUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Ca_BandwidthClassUL_EUTRAPresent bool
	if Ca_BandwidthClassUL_EUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.BandEUTRA.Decode(r); err != nil {
		return utils.WrapError("Decode BandEUTRA", err)
	}
	if Ca_BandwidthClassDL_EUTRAPresent {
		ie.Ca_BandwidthClassDL_EUTRA = new(CA_BandwidthClassEUTRA)
		if err = ie.Ca_BandwidthClassDL_EUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode Ca_BandwidthClassDL_EUTRA", err)
		}
	}
	if Ca_BandwidthClassUL_EUTRAPresent {
		ie.Ca_BandwidthClassUL_EUTRA = new(CA_BandwidthClassEUTRA)
		if err = ie.Ca_BandwidthClassUL_EUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode Ca_BandwidthClassUL_EUTRA", err)
		}
	}
	return nil
}
