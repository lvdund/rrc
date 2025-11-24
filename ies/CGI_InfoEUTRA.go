package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CGI_InfoEUTRA struct {
	Cgi_info_EPC              *CGI_InfoEUTRA_cgi_info_EPC              `lb:1,ub:maxPLMN,optional`
	Cgi_info_5GC              []CellAccessRelatedInfo_EUTRA_5GC        `lb:1,ub:maxPLMN,optional`
	FreqBandIndicator         FreqBandIndicatorEUTRA                   `madatory`
	MultiBandInfoList         *MultiBandInfoListEUTRA                  `optional`
	FreqBandIndicatorPriority *CGI_InfoEUTRA_freqBandIndicatorPriority `optional`
}

func (ie *CGI_InfoEUTRA) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Cgi_info_EPC != nil, len(ie.Cgi_info_5GC) > 0, ie.MultiBandInfoList != nil, ie.FreqBandIndicatorPriority != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Cgi_info_EPC != nil {
		if err = ie.Cgi_info_EPC.Encode(w); err != nil {
			return utils.WrapError("Encode Cgi_info_EPC", err)
		}
	}
	if len(ie.Cgi_info_5GC) > 0 {
		tmp_Cgi_info_5GC := utils.NewSequence[*CellAccessRelatedInfo_EUTRA_5GC]([]*CellAccessRelatedInfo_EUTRA_5GC{}, uper.Constraint{Lb: 1, Ub: maxPLMN}, false)
		for _, i := range ie.Cgi_info_5GC {
			tmp_Cgi_info_5GC.Value = append(tmp_Cgi_info_5GC.Value, &i)
		}
		if err = tmp_Cgi_info_5GC.Encode(w); err != nil {
			return utils.WrapError("Encode Cgi_info_5GC", err)
		}
	}
	if err = ie.FreqBandIndicator.Encode(w); err != nil {
		return utils.WrapError("Encode FreqBandIndicator", err)
	}
	if ie.MultiBandInfoList != nil {
		if err = ie.MultiBandInfoList.Encode(w); err != nil {
			return utils.WrapError("Encode MultiBandInfoList", err)
		}
	}
	if ie.FreqBandIndicatorPriority != nil {
		if err = ie.FreqBandIndicatorPriority.Encode(w); err != nil {
			return utils.WrapError("Encode FreqBandIndicatorPriority", err)
		}
	}
	return nil
}

func (ie *CGI_InfoEUTRA) Decode(r *uper.UperReader) error {
	var err error
	var Cgi_info_EPCPresent bool
	if Cgi_info_EPCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Cgi_info_5GCPresent bool
	if Cgi_info_5GCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MultiBandInfoListPresent bool
	if MultiBandInfoListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var FreqBandIndicatorPriorityPresent bool
	if FreqBandIndicatorPriorityPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Cgi_info_EPCPresent {
		ie.Cgi_info_EPC = new(CGI_InfoEUTRA_cgi_info_EPC)
		if err = ie.Cgi_info_EPC.Decode(r); err != nil {
			return utils.WrapError("Decode Cgi_info_EPC", err)
		}
	}
	if Cgi_info_5GCPresent {
		tmp_Cgi_info_5GC := utils.NewSequence[*CellAccessRelatedInfo_EUTRA_5GC]([]*CellAccessRelatedInfo_EUTRA_5GC{}, uper.Constraint{Lb: 1, Ub: maxPLMN}, false)
		fn_Cgi_info_5GC := func() *CellAccessRelatedInfo_EUTRA_5GC {
			return new(CellAccessRelatedInfo_EUTRA_5GC)
		}
		if err = tmp_Cgi_info_5GC.Decode(r, fn_Cgi_info_5GC); err != nil {
			return utils.WrapError("Decode Cgi_info_5GC", err)
		}
		ie.Cgi_info_5GC = []CellAccessRelatedInfo_EUTRA_5GC{}
		for _, i := range tmp_Cgi_info_5GC.Value {
			ie.Cgi_info_5GC = append(ie.Cgi_info_5GC, *i)
		}
	}
	if err = ie.FreqBandIndicator.Decode(r); err != nil {
		return utils.WrapError("Decode FreqBandIndicator", err)
	}
	if MultiBandInfoListPresent {
		ie.MultiBandInfoList = new(MultiBandInfoListEUTRA)
		if err = ie.MultiBandInfoList.Decode(r); err != nil {
			return utils.WrapError("Decode MultiBandInfoList", err)
		}
	}
	if FreqBandIndicatorPriorityPresent {
		ie.FreqBandIndicatorPriority = new(CGI_InfoEUTRA_freqBandIndicatorPriority)
		if err = ie.FreqBandIndicatorPriority.Decode(r); err != nil {
			return utils.WrapError("Decode FreqBandIndicatorPriority", err)
		}
	}
	return nil
}
