package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DummyG struct {
	MaxNumberSSB_CSI_RS_ResourceOneTx DummyG_maxNumberSSB_CSI_RS_ResourceOneTx `madatory`
	MaxNumberSSB_CSI_RS_ResourceTwoTx DummyG_maxNumberSSB_CSI_RS_ResourceTwoTx `madatory`
	SupportedCSI_RS_Density           DummyG_supportedCSI_RS_Density           `madatory`
}

func (ie *DummyG) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.MaxNumberSSB_CSI_RS_ResourceOneTx.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberSSB_CSI_RS_ResourceOneTx", err)
	}
	if err = ie.MaxNumberSSB_CSI_RS_ResourceTwoTx.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberSSB_CSI_RS_ResourceTwoTx", err)
	}
	if err = ie.SupportedCSI_RS_Density.Encode(w); err != nil {
		return utils.WrapError("Encode SupportedCSI_RS_Density", err)
	}
	return nil
}

func (ie *DummyG) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.MaxNumberSSB_CSI_RS_ResourceOneTx.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberSSB_CSI_RS_ResourceOneTx", err)
	}
	if err = ie.MaxNumberSSB_CSI_RS_ResourceTwoTx.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberSSB_CSI_RS_ResourceTwoTx", err)
	}
	if err = ie.SupportedCSI_RS_Density.Decode(r); err != nil {
		return utils.WrapError("Decode SupportedCSI_RS_Density", err)
	}
	return nil
}
