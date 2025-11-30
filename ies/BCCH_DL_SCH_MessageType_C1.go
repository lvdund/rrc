package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	BCCH_DL_SCH_MessageType_C1_Choice_nothing uint64 = iota
	BCCH_DL_SCH_MessageType_C1_Choice_SystemInformation
	BCCH_DL_SCH_MessageType_C1_Choice_SystemInformationBlockType1
)

type BCCH_DL_SCH_MessageType_C1 struct {
	Choice                      uint64
	SystemInformation           *SystemInformation
	SystemInformationBlockType1 *SIB1
}

func (ie *BCCH_DL_SCH_MessageType_C1) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BCCH_DL_SCH_MessageType_C1_Choice_SystemInformation:
		if err = ie.SystemInformation.Encode(w); err != nil {
			err = utils.WrapError("Encode SystemInformation", err)
		}
	case BCCH_DL_SCH_MessageType_C1_Choice_SystemInformationBlockType1:
		if err = ie.SystemInformationBlockType1.Encode(w); err != nil {
			err = utils.WrapError("Encode SystemInformationBlockType1", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *BCCH_DL_SCH_MessageType_C1) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BCCH_DL_SCH_MessageType_C1_Choice_SystemInformation:
		ie.SystemInformation = new(SystemInformation)
		if err = ie.SystemInformation.Decode(r); err != nil {
			return utils.WrapError("Decode SystemInformation", err)
		}
	case BCCH_DL_SCH_MessageType_C1_Choice_SystemInformationBlockType1:
		ie.SystemInformationBlockType1 = new(SIB1)
		if err = ie.SystemInformationBlockType1.Decode(r); err != nil {
			return utils.WrapError("Decode SystemInformationBlockType1", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
