package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_NR_Capability_v15c0 struct {
	Nrdc_Parameters_v15c0     *NRDC_Parameters_v15c0                            `optional`
	PartialFR2_FallbackRX_Req *UE_NR_Capability_v15c0_partialFR2_FallbackRX_Req `optional`
	NonCriticalExtension      *UE_NR_Capability_v15g0                           `optional`
}

func (ie *UE_NR_Capability_v15c0) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Nrdc_Parameters_v15c0 != nil, ie.PartialFR2_FallbackRX_Req != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Nrdc_Parameters_v15c0 != nil {
		if err = ie.Nrdc_Parameters_v15c0.Encode(w); err != nil {
			return utils.WrapError("Encode Nrdc_Parameters_v15c0", err)
		}
	}
	if ie.PartialFR2_FallbackRX_Req != nil {
		if err = ie.PartialFR2_FallbackRX_Req.Encode(w); err != nil {
			return utils.WrapError("Encode PartialFR2_FallbackRX_Req", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UE_NR_Capability_v15c0) Decode(r *uper.UperReader) error {
	var err error
	var Nrdc_Parameters_v15c0Present bool
	if Nrdc_Parameters_v15c0Present, err = r.ReadBool(); err != nil {
		return err
	}
	var PartialFR2_FallbackRX_ReqPresent bool
	if PartialFR2_FallbackRX_ReqPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Nrdc_Parameters_v15c0Present {
		ie.Nrdc_Parameters_v15c0 = new(NRDC_Parameters_v15c0)
		if err = ie.Nrdc_Parameters_v15c0.Decode(r); err != nil {
			return utils.WrapError("Decode Nrdc_Parameters_v15c0", err)
		}
	}
	if PartialFR2_FallbackRX_ReqPresent {
		ie.PartialFR2_FallbackRX_Req = new(UE_NR_Capability_v15c0_partialFR2_FallbackRX_Req)
		if err = ie.PartialFR2_FallbackRX_Req.Decode(r); err != nil {
			return utils.WrapError("Decode PartialFR2_FallbackRX_Req", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(UE_NR_Capability_v15g0)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
