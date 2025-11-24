package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasObjectToAddMod_measObject_Choice_nothing uint64 = iota
	MeasObjectToAddMod_measObject_Choice_MeasObjectNR
	MeasObjectToAddMod_measObject_Choice_MeasObjectEUTRA
	MeasObjectToAddMod_measObject_Choice_MeasObjectUTRA_FDD_r16
	MeasObjectToAddMod_measObject_Choice_MeasObjectNR_SL_r16
	MeasObjectToAddMod_measObject_Choice_MeasObjectCLI_r16
	MeasObjectToAddMod_measObject_Choice_MeasObjectRxTxDiff_r17
	MeasObjectToAddMod_measObject_Choice_MeasObjectRelay_r17
)

type MeasObjectToAddMod_measObject struct {
	Choice                 uint64
	MeasObjectNR           *MeasObjectNR
	MeasObjectEUTRA        *MeasObjectEUTRA
	MeasObjectUTRA_FDD_r16 *MeasObjectUTRA_FDD_r16
	MeasObjectNR_SL_r16    *MeasObjectNR_SL_r16
	MeasObjectCLI_r16      *MeasObjectCLI_r16
	MeasObjectRxTxDiff_r17 *MeasObjectRxTxDiff_r17
	MeasObjectRelay_r17    *SL_MeasObject_r16
}

func (ie *MeasObjectToAddMod_measObject) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 7, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasObjectToAddMod_measObject_Choice_MeasObjectNR:
		if err = ie.MeasObjectNR.Encode(w); err != nil {
			err = utils.WrapError("Encode MeasObjectNR", err)
		}
	case MeasObjectToAddMod_measObject_Choice_MeasObjectEUTRA:
		if err = ie.MeasObjectEUTRA.Encode(w); err != nil {
			err = utils.WrapError("Encode MeasObjectEUTRA", err)
		}
	case MeasObjectToAddMod_measObject_Choice_MeasObjectUTRA_FDD_r16:
		if err = ie.MeasObjectUTRA_FDD_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode MeasObjectUTRA_FDD_r16", err)
		}
	case MeasObjectToAddMod_measObject_Choice_MeasObjectNR_SL_r16:
		if err = ie.MeasObjectNR_SL_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode MeasObjectNR_SL_r16", err)
		}
	case MeasObjectToAddMod_measObject_Choice_MeasObjectCLI_r16:
		if err = ie.MeasObjectCLI_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode MeasObjectCLI_r16", err)
		}
	case MeasObjectToAddMod_measObject_Choice_MeasObjectRxTxDiff_r17:
		if err = ie.MeasObjectRxTxDiff_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode MeasObjectRxTxDiff_r17", err)
		}
	case MeasObjectToAddMod_measObject_Choice_MeasObjectRelay_r17:
		if err = ie.MeasObjectRelay_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode MeasObjectRelay_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *MeasObjectToAddMod_measObject) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(7, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasObjectToAddMod_measObject_Choice_MeasObjectNR:
		ie.MeasObjectNR = new(MeasObjectNR)
		if err = ie.MeasObjectNR.Decode(r); err != nil {
			return utils.WrapError("Decode MeasObjectNR", err)
		}
	case MeasObjectToAddMod_measObject_Choice_MeasObjectEUTRA:
		ie.MeasObjectEUTRA = new(MeasObjectEUTRA)
		if err = ie.MeasObjectEUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode MeasObjectEUTRA", err)
		}
	case MeasObjectToAddMod_measObject_Choice_MeasObjectUTRA_FDD_r16:
		ie.MeasObjectUTRA_FDD_r16 = new(MeasObjectUTRA_FDD_r16)
		if err = ie.MeasObjectUTRA_FDD_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MeasObjectUTRA_FDD_r16", err)
		}
	case MeasObjectToAddMod_measObject_Choice_MeasObjectNR_SL_r16:
		ie.MeasObjectNR_SL_r16 = new(MeasObjectNR_SL_r16)
		if err = ie.MeasObjectNR_SL_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MeasObjectNR_SL_r16", err)
		}
	case MeasObjectToAddMod_measObject_Choice_MeasObjectCLI_r16:
		ie.MeasObjectCLI_r16 = new(MeasObjectCLI_r16)
		if err = ie.MeasObjectCLI_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MeasObjectCLI_r16", err)
		}
	case MeasObjectToAddMod_measObject_Choice_MeasObjectRxTxDiff_r17:
		ie.MeasObjectRxTxDiff_r17 = new(MeasObjectRxTxDiff_r17)
		if err = ie.MeasObjectRxTxDiff_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MeasObjectRxTxDiff_r17", err)
		}
	case MeasObjectToAddMod_measObject_Choice_MeasObjectRelay_r17:
		ie.MeasObjectRelay_r17 = new(SL_MeasObject_r16)
		if err = ie.MeasObjectRelay_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MeasObjectRelay_r17", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
