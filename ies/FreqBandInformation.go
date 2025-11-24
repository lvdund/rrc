package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	FreqBandInformation_Choice_nothing uint64 = iota
	FreqBandInformation_Choice_BandInformationEUTRA
	FreqBandInformation_Choice_BandInformationNR
)

type FreqBandInformation struct {
	Choice               uint64
	BandInformationEUTRA *FreqBandInformationEUTRA
	BandInformationNR    *FreqBandInformationNR
}

func (ie *FreqBandInformation) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case FreqBandInformation_Choice_BandInformationEUTRA:
		if err = ie.BandInformationEUTRA.Encode(w); err != nil {
			err = utils.WrapError("Encode BandInformationEUTRA", err)
		}
	case FreqBandInformation_Choice_BandInformationNR:
		if err = ie.BandInformationNR.Encode(w); err != nil {
			err = utils.WrapError("Encode BandInformationNR", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *FreqBandInformation) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case FreqBandInformation_Choice_BandInformationEUTRA:
		ie.BandInformationEUTRA = new(FreqBandInformationEUTRA)
		if err = ie.BandInformationEUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode BandInformationEUTRA", err)
		}
	case FreqBandInformation_Choice_BandInformationNR:
		ie.BandInformationNR = new(FreqBandInformationNR)
		if err = ie.BandInformationNR.Decode(r); err != nil {
			return utils.WrapError("Decode BandInformationNR", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
