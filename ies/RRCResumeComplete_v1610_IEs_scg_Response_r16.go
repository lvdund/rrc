package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RRCResumeComplete_v1610_IEs_scg_Response_r16_Choice_nothing uint64 = iota
	RRCResumeComplete_v1610_IEs_scg_Response_r16_Choice_Nr_SCG_Response
	RRCResumeComplete_v1610_IEs_scg_Response_r16_Choice_Eutra_SCG_Response
)

type RRCResumeComplete_v1610_IEs_scg_Response_r16 struct {
	Choice             uint64
	Nr_SCG_Response    []byte `madatory`
	Eutra_SCG_Response []byte `madatory`
}

func (ie *RRCResumeComplete_v1610_IEs_scg_Response_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCResumeComplete_v1610_IEs_scg_Response_r16_Choice_Nr_SCG_Response:
		if err = w.WriteOctetString(ie.Nr_SCG_Response, nil, false); err != nil {
			err = utils.WrapError("Encode Nr_SCG_Response", err)
		}
	case RRCResumeComplete_v1610_IEs_scg_Response_r16_Choice_Eutra_SCG_Response:
		if err = w.WriteOctetString(ie.Eutra_SCG_Response, nil, false); err != nil {
			err = utils.WrapError("Encode Eutra_SCG_Response", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RRCResumeComplete_v1610_IEs_scg_Response_r16) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCResumeComplete_v1610_IEs_scg_Response_r16_Choice_Nr_SCG_Response:
		var tmp_os_Nr_SCG_Response []byte
		if tmp_os_Nr_SCG_Response, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode Nr_SCG_Response", err)
		}
		ie.Nr_SCG_Response = tmp_os_Nr_SCG_Response
	case RRCResumeComplete_v1610_IEs_scg_Response_r16_Choice_Eutra_SCG_Response:
		var tmp_os_Eutra_SCG_Response []byte
		if tmp_os_Eutra_SCG_Response, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode Eutra_SCG_Response", err)
		}
		ie.Eutra_SCG_Response = tmp_os_Eutra_SCG_Response
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
