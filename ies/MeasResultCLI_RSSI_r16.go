package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultCLI_RSSI_r16 struct {
	Rssi_ResourceId_r16 RSSI_ResourceId_r16 `madatory`
	Cli_RSSI_Result_r16 CLI_RSSI_Range_r16  `madatory`
}

func (ie *MeasResultCLI_RSSI_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.Rssi_ResourceId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Rssi_ResourceId_r16", err)
	}
	if err = ie.Cli_RSSI_Result_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Cli_RSSI_Result_r16", err)
	}
	return nil
}

func (ie *MeasResultCLI_RSSI_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.Rssi_ResourceId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Rssi_ResourceId_r16", err)
	}
	if err = ie.Cli_RSSI_Result_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Cli_RSSI_Result_r16", err)
	}
	return nil
}
