package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandParameters_v1540 struct {
	Srs_CarrierSwitch *BandParameters_v1540_srs_CarrierSwitch `lb:1,ub:maxSimultaneousBands,optional`
	Srs_TxSwitch      *BandParameters_v1540_srs_TxSwitch      `lb:1,ub:32,optional`
}

func (ie *BandParameters_v1540) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Srs_CarrierSwitch != nil, ie.Srs_TxSwitch != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Srs_CarrierSwitch != nil {
		if err = ie.Srs_CarrierSwitch.Encode(w); err != nil {
			return utils.WrapError("Encode Srs_CarrierSwitch", err)
		}
	}
	if ie.Srs_TxSwitch != nil {
		if err = ie.Srs_TxSwitch.Encode(w); err != nil {
			return utils.WrapError("Encode Srs_TxSwitch", err)
		}
	}
	return nil
}

func (ie *BandParameters_v1540) Decode(r *uper.UperReader) error {
	var err error
	var Srs_CarrierSwitchPresent bool
	if Srs_CarrierSwitchPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Srs_TxSwitchPresent bool
	if Srs_TxSwitchPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Srs_CarrierSwitchPresent {
		ie.Srs_CarrierSwitch = new(BandParameters_v1540_srs_CarrierSwitch)
		if err = ie.Srs_CarrierSwitch.Decode(r); err != nil {
			return utils.WrapError("Decode Srs_CarrierSwitch", err)
		}
	}
	if Srs_TxSwitchPresent {
		ie.Srs_TxSwitch = new(BandParameters_v1540_srs_TxSwitch)
		if err = ie.Srs_TxSwitch.Decode(r); err != nil {
			return utils.WrapError("Decode Srs_TxSwitch", err)
		}
	}
	return nil
}
