package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	VisitedPSCellInfo_r17_visitedCellId_r17_nr_CellId_r17_Choice_nothing uint64 = iota
	VisitedPSCellInfo_r17_visitedCellId_r17_nr_CellId_r17_Choice_Cgi_Info_r17
	VisitedPSCellInfo_r17_visitedCellId_r17_nr_CellId_r17_Choice_Pci_arfcn_r17
)

type VisitedPSCellInfo_r17_visitedCellId_r17_nr_CellId_r17 struct {
	Choice        uint64
	Cgi_Info_r17  *CGI_Info_Logging_r16
	Pci_arfcn_r17 *PCI_ARFCN_NR_r16
}

func (ie *VisitedPSCellInfo_r17_visitedCellId_r17_nr_CellId_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case VisitedPSCellInfo_r17_visitedCellId_r17_nr_CellId_r17_Choice_Cgi_Info_r17:
		if err = ie.Cgi_Info_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode Cgi_Info_r17", err)
		}
	case VisitedPSCellInfo_r17_visitedCellId_r17_nr_CellId_r17_Choice_Pci_arfcn_r17:
		if err = ie.Pci_arfcn_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode Pci_arfcn_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *VisitedPSCellInfo_r17_visitedCellId_r17_nr_CellId_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case VisitedPSCellInfo_r17_visitedCellId_r17_nr_CellId_r17_Choice_Cgi_Info_r17:
		ie.Cgi_Info_r17 = new(CGI_Info_Logging_r16)
		if err = ie.Cgi_Info_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Cgi_Info_r17", err)
		}
	case VisitedPSCellInfo_r17_visitedCellId_r17_nr_CellId_r17_Choice_Pci_arfcn_r17:
		ie.Pci_arfcn_r17 = new(PCI_ARFCN_NR_r16)
		if err = ie.Pci_arfcn_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Pci_arfcn_r17", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
