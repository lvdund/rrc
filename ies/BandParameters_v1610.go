package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BandParameters_v1610 struct {
	Srs_TxSwitch_v1610 *BandParameters_v1610_srs_TxSwitch_v1610 `optional`
}

func (ie *BandParameters_v1610) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Srs_TxSwitch_v1610 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Srs_TxSwitch_v1610 != nil {
		if err = ie.Srs_TxSwitch_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode Srs_TxSwitch_v1610", err)
		}
	}
	return nil
}

func (ie *BandParameters_v1610) Decode(r *aper.AperReader) error {
	var err error
	var Srs_TxSwitch_v1610Present bool
	if Srs_TxSwitch_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Srs_TxSwitch_v1610Present {
		ie.Srs_TxSwitch_v1610 = new(BandParameters_v1610_srs_TxSwitch_v1610)
		if err = ie.Srs_TxSwitch_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode Srs_TxSwitch_v1610", err)
		}
	}
	return nil
}
