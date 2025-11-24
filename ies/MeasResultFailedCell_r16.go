package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultFailedCell_r16 struct {
	Cgi_Info       CGI_Info_Logging_r16                    `madatory`
	MeasResult_r16 MeasResultFailedCell_r16_measResult_r16 `madatory`
}

func (ie *MeasResultFailedCell_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.Cgi_Info.Encode(w); err != nil {
		return utils.WrapError("Encode Cgi_Info", err)
	}
	if err = ie.MeasResult_r16.Encode(w); err != nil {
		return utils.WrapError("Encode MeasResult_r16", err)
	}
	return nil
}

func (ie *MeasResultFailedCell_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.Cgi_Info.Decode(r); err != nil {
		return utils.WrapError("Decode Cgi_Info", err)
	}
	if err = ie.MeasResult_r16.Decode(r); err != nil {
		return utils.WrapError("Decode MeasResult_r16", err)
	}
	return nil
}
