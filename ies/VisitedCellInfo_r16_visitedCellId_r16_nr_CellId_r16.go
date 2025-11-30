package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	VisitedCellInfo_r16_visitedCellId_r16_nr_CellId_r16_Choice_nothing uint64 = iota
	VisitedCellInfo_r16_visitedCellId_r16_nr_CellId_r16_Choice_Cgi_Info
	VisitedCellInfo_r16_visitedCellId_r16_nr_CellId_r16_Choice_Pci_arfcn_r16
)

type VisitedCellInfo_r16_visitedCellId_r16_nr_CellId_r16 struct {
	Choice        uint64
	Cgi_Info      *CGI_Info_Logging_r16
	Pci_arfcn_r16 *PCI_ARFCN_NR_r16
}

func (ie *VisitedCellInfo_r16_visitedCellId_r16_nr_CellId_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case VisitedCellInfo_r16_visitedCellId_r16_nr_CellId_r16_Choice_Cgi_Info:
		if err = ie.Cgi_Info.Encode(w); err != nil {
			err = utils.WrapError("Encode Cgi_Info", err)
		}
	case VisitedCellInfo_r16_visitedCellId_r16_nr_CellId_r16_Choice_Pci_arfcn_r16:
		if err = ie.Pci_arfcn_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode Pci_arfcn_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *VisitedCellInfo_r16_visitedCellId_r16_nr_CellId_r16) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case VisitedCellInfo_r16_visitedCellId_r16_nr_CellId_r16_Choice_Cgi_Info:
		ie.Cgi_Info = new(CGI_Info_Logging_r16)
		if err = ie.Cgi_Info.Decode(r); err != nil {
			return utils.WrapError("Decode Cgi_Info", err)
		}
	case VisitedCellInfo_r16_visitedCellId_r16_nr_CellId_r16_Choice_Pci_arfcn_r16:
		ie.Pci_arfcn_r16 = new(PCI_ARFCN_NR_r16)
		if err = ie.Pci_arfcn_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Pci_arfcn_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
