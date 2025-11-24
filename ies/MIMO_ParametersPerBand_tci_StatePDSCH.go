package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_tci_StatePDSCH struct {
	MaxNumberConfiguredTCI_StatesPerCC *MIMO_ParametersPerBand_tci_StatePDSCH_maxNumberConfiguredTCI_StatesPerCC `optional`
	MaxNumberActiveTCI_PerBWP          *MIMO_ParametersPerBand_tci_StatePDSCH_maxNumberActiveTCI_PerBWP          `optional`
}

func (ie *MIMO_ParametersPerBand_tci_StatePDSCH) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.MaxNumberConfiguredTCI_StatesPerCC != nil, ie.MaxNumberActiveTCI_PerBWP != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MaxNumberConfiguredTCI_StatesPerCC != nil {
		if err = ie.MaxNumberConfiguredTCI_StatesPerCC.Encode(w); err != nil {
			return utils.WrapError("Encode MaxNumberConfiguredTCI_StatesPerCC", err)
		}
	}
	if ie.MaxNumberActiveTCI_PerBWP != nil {
		if err = ie.MaxNumberActiveTCI_PerBWP.Encode(w); err != nil {
			return utils.WrapError("Encode MaxNumberActiveTCI_PerBWP", err)
		}
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_tci_StatePDSCH) Decode(r *uper.UperReader) error {
	var err error
	var MaxNumberConfiguredTCI_StatesPerCCPresent bool
	if MaxNumberConfiguredTCI_StatesPerCCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxNumberActiveTCI_PerBWPPresent bool
	if MaxNumberActiveTCI_PerBWPPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if MaxNumberConfiguredTCI_StatesPerCCPresent {
		ie.MaxNumberConfiguredTCI_StatesPerCC = new(MIMO_ParametersPerBand_tci_StatePDSCH_maxNumberConfiguredTCI_StatesPerCC)
		if err = ie.MaxNumberConfiguredTCI_StatesPerCC.Decode(r); err != nil {
			return utils.WrapError("Decode MaxNumberConfiguredTCI_StatesPerCC", err)
		}
	}
	if MaxNumberActiveTCI_PerBWPPresent {
		ie.MaxNumberActiveTCI_PerBWP = new(MIMO_ParametersPerBand_tci_StatePDSCH_maxNumberActiveTCI_PerBWP)
		if err = ie.MaxNumberActiveTCI_PerBWP.Decode(r); err != nil {
			return utils.WrapError("Decode MaxNumberActiveTCI_PerBWP", err)
		}
	}
	return nil
}
