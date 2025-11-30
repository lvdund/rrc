package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RLF_Report_r16_nr_RLF_Report_r16_failedPCellId_r16_Choice_nothing uint64 = iota
	RLF_Report_r16_nr_RLF_Report_r16_failedPCellId_r16_Choice_NrFailedPCellId_r16
	RLF_Report_r16_nr_RLF_Report_r16_failedPCellId_r16_Choice_EutraFailedPCellId_r16
)

type RLF_Report_r16_nr_RLF_Report_r16_failedPCellId_r16 struct {
	Choice                 uint64
	NrFailedPCellId_r16    *RLF_Report_r16_nr_RLF_Report_r16_failedPCellId_r16_nrFailedPCellId_r16
	EutraFailedPCellId_r16 *RLF_Report_r16_nr_RLF_Report_r16_failedPCellId_r16_eutraFailedPCellId_r16
}

func (ie *RLF_Report_r16_nr_RLF_Report_r16_failedPCellId_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RLF_Report_r16_nr_RLF_Report_r16_failedPCellId_r16_Choice_NrFailedPCellId_r16:
		if err = ie.NrFailedPCellId_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode NrFailedPCellId_r16", err)
		}
	case RLF_Report_r16_nr_RLF_Report_r16_failedPCellId_r16_Choice_EutraFailedPCellId_r16:
		if err = ie.EutraFailedPCellId_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode EutraFailedPCellId_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RLF_Report_r16_nr_RLF_Report_r16_failedPCellId_r16) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RLF_Report_r16_nr_RLF_Report_r16_failedPCellId_r16_Choice_NrFailedPCellId_r16:
		ie.NrFailedPCellId_r16 = new(RLF_Report_r16_nr_RLF_Report_r16_failedPCellId_r16_nrFailedPCellId_r16)
		if err = ie.NrFailedPCellId_r16.Decode(r); err != nil {
			return utils.WrapError("Decode NrFailedPCellId_r16", err)
		}
	case RLF_Report_r16_nr_RLF_Report_r16_failedPCellId_r16_Choice_EutraFailedPCellId_r16:
		ie.EutraFailedPCellId_r16 = new(RLF_Report_r16_nr_RLF_Report_r16_failedPCellId_r16_eutraFailedPCellId_r16)
		if err = ie.EutraFailedPCellId_r16.Decode(r); err != nil {
			return utils.WrapError("Decode EutraFailedPCellId_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
