package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUSCH_CodeBlockGroupTransmission struct {
	MaxCodeBlockGroupsPerTransportBlock PUSCH_CodeBlockGroupTransmission_maxCodeBlockGroupsPerTransportBlock `madatory`
}

func (ie *PUSCH_CodeBlockGroupTransmission) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.MaxCodeBlockGroupsPerTransportBlock.Encode(w); err != nil {
		return utils.WrapError("Encode MaxCodeBlockGroupsPerTransportBlock", err)
	}
	return nil
}

func (ie *PUSCH_CodeBlockGroupTransmission) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.MaxCodeBlockGroupsPerTransportBlock.Decode(r); err != nil {
		return utils.WrapError("Decode MaxCodeBlockGroupsPerTransportBlock", err)
	}
	return nil
}
