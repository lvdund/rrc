package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandParameters_v1540_srs_CarrierSwitch_Choice_nothing uint64 = iota
	BandParameters_v1540_srs_CarrierSwitch_Choice_Nr
	BandParameters_v1540_srs_CarrierSwitch_Choice_Eutra
)

type BandParameters_v1540_srs_CarrierSwitch struct {
	Choice uint64
	Nr     *BandParameters_v1540_srs_CarrierSwitch_nr
	Eutra  *BandParameters_v1540_srs_CarrierSwitch_eutra
}

func (ie *BandParameters_v1540_srs_CarrierSwitch) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BandParameters_v1540_srs_CarrierSwitch_Choice_Nr:
		if err = ie.Nr.Encode(w); err != nil {
			err = utils.WrapError("Encode Nr", err)
		}
	case BandParameters_v1540_srs_CarrierSwitch_Choice_Eutra:
		if err = ie.Eutra.Encode(w); err != nil {
			err = utils.WrapError("Encode Eutra", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *BandParameters_v1540_srs_CarrierSwitch) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BandParameters_v1540_srs_CarrierSwitch_Choice_Nr:
		ie.Nr = new(BandParameters_v1540_srs_CarrierSwitch_nr)
		if err = ie.Nr.Decode(r); err != nil {
			return utils.WrapError("Decode Nr", err)
		}
	case BandParameters_v1540_srs_CarrierSwitch_Choice_Eutra:
		ie.Eutra = new(BandParameters_v1540_srs_CarrierSwitch_eutra)
		if err = ie.Eutra.Decode(r); err != nil {
			return utils.WrapError("Decode Eutra", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
