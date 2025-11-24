package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RRCReconfigurationComplete_v1560_IEs_scg_Response_Choice_nothing uint64 = iota
	RRCReconfigurationComplete_v1560_IEs_scg_Response_Choice_Nr_SCG_Response
	RRCReconfigurationComplete_v1560_IEs_scg_Response_Choice_Eutra_SCG_Response
)

type RRCReconfigurationComplete_v1560_IEs_scg_Response struct {
	Choice             uint64
	Nr_SCG_Response    []byte `madatory`
	Eutra_SCG_Response []byte `madatory`
}

func (ie *RRCReconfigurationComplete_v1560_IEs_scg_Response) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCReconfigurationComplete_v1560_IEs_scg_Response_Choice_Nr_SCG_Response:
		if err = w.WriteOctetString(ie.Nr_SCG_Response, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			err = utils.WrapError("Encode Nr_SCG_Response", err)
		}
	case RRCReconfigurationComplete_v1560_IEs_scg_Response_Choice_Eutra_SCG_Response:
		if err = w.WriteOctetString(ie.Eutra_SCG_Response, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			err = utils.WrapError("Encode Eutra_SCG_Response", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RRCReconfigurationComplete_v1560_IEs_scg_Response) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCReconfigurationComplete_v1560_IEs_scg_Response_Choice_Nr_SCG_Response:
		var tmp_os_Nr_SCG_Response []byte
		if tmp_os_Nr_SCG_Response, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode Nr_SCG_Response", err)
		}
		ie.Nr_SCG_Response = tmp_os_Nr_SCG_Response
	case RRCReconfigurationComplete_v1560_IEs_scg_Response_Choice_Eutra_SCG_Response:
		var tmp_os_Eutra_SCG_Response []byte
		if tmp_os_Eutra_SCG_Response, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode Eutra_SCG_Response", err)
		}
		ie.Eutra_SCG_Response = tmp_os_Eutra_SCG_Response
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
