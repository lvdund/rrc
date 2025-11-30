package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UE_CapabilityRequestFilterCommon_mrdc_Request struct {
	OmitEN_DC    *UE_CapabilityRequestFilterCommon_mrdc_Request_omitEN_DC    `optional`
	IncludeNR_DC *UE_CapabilityRequestFilterCommon_mrdc_Request_includeNR_DC `optional`
	IncludeNE_DC *UE_CapabilityRequestFilterCommon_mrdc_Request_includeNE_DC `optional`
}

func (ie *UE_CapabilityRequestFilterCommon_mrdc_Request) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.OmitEN_DC != nil, ie.IncludeNR_DC != nil, ie.IncludeNE_DC != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.OmitEN_DC != nil {
		if err = ie.OmitEN_DC.Encode(w); err != nil {
			return utils.WrapError("Encode OmitEN_DC", err)
		}
	}
	if ie.IncludeNR_DC != nil {
		if err = ie.IncludeNR_DC.Encode(w); err != nil {
			return utils.WrapError("Encode IncludeNR_DC", err)
		}
	}
	if ie.IncludeNE_DC != nil {
		if err = ie.IncludeNE_DC.Encode(w); err != nil {
			return utils.WrapError("Encode IncludeNE_DC", err)
		}
	}
	return nil
}

func (ie *UE_CapabilityRequestFilterCommon_mrdc_Request) Decode(r *aper.AperReader) error {
	var err error
	var OmitEN_DCPresent bool
	if OmitEN_DCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var IncludeNR_DCPresent bool
	if IncludeNR_DCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var IncludeNE_DCPresent bool
	if IncludeNE_DCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if OmitEN_DCPresent {
		ie.OmitEN_DC = new(UE_CapabilityRequestFilterCommon_mrdc_Request_omitEN_DC)
		if err = ie.OmitEN_DC.Decode(r); err != nil {
			return utils.WrapError("Decode OmitEN_DC", err)
		}
	}
	if IncludeNR_DCPresent {
		ie.IncludeNR_DC = new(UE_CapabilityRequestFilterCommon_mrdc_Request_includeNR_DC)
		if err = ie.IncludeNR_DC.Decode(r); err != nil {
			return utils.WrapError("Decode IncludeNR_DC", err)
		}
	}
	if IncludeNE_DCPresent {
		ie.IncludeNE_DC = new(UE_CapabilityRequestFilterCommon_mrdc_Request_includeNE_DC)
		if err = ie.IncludeNE_DC.Decode(r); err != nil {
			return utils.WrapError("Decode IncludeNE_DC", err)
		}
	}
	return nil
}
