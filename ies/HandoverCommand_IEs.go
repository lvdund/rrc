package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type HandoverCommand_IEs struct {
	HandoverCommandMessage []byte      `madatory`
	NonCriticalExtension   interface{} `optional`
}

func (ie *HandoverCommand_IEs) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteOctetString(ie.HandoverCommandMessage, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("WriteOctetString HandoverCommandMessage", err)
	}
	return nil
}

func (ie *HandoverCommand_IEs) Decode(r *uper.UperReader) error {
	var err error
	var tmp_os_HandoverCommandMessage []byte
	if tmp_os_HandoverCommandMessage, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("ReadOctetString HandoverCommandMessage", err)
	}
	ie.HandoverCommandMessage = tmp_os_HandoverCommandMessage
	return nil
}
