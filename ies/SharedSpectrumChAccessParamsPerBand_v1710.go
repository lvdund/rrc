package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SharedSpectrumChAccessParamsPerBand_v1710 struct {
	Ul_Semi_StaticChAccessDependentConfig_r17   *SharedSpectrumChAccessParamsPerBand_v1710_ul_Semi_StaticChAccessDependentConfig_r17   `optional`
	Ul_Semi_StaticChAccessIndependentConfig_r17 *SharedSpectrumChAccessParamsPerBand_v1710_ul_Semi_StaticChAccessIndependentConfig_r17 `optional`
}

func (ie *SharedSpectrumChAccessParamsPerBand_v1710) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Ul_Semi_StaticChAccessDependentConfig_r17 != nil, ie.Ul_Semi_StaticChAccessIndependentConfig_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ul_Semi_StaticChAccessDependentConfig_r17 != nil {
		if err = ie.Ul_Semi_StaticChAccessDependentConfig_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ul_Semi_StaticChAccessDependentConfig_r17", err)
		}
	}
	if ie.Ul_Semi_StaticChAccessIndependentConfig_r17 != nil {
		if err = ie.Ul_Semi_StaticChAccessIndependentConfig_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ul_Semi_StaticChAccessIndependentConfig_r17", err)
		}
	}
	return nil
}

func (ie *SharedSpectrumChAccessParamsPerBand_v1710) Decode(r *aper.AperReader) error {
	var err error
	var Ul_Semi_StaticChAccessDependentConfig_r17Present bool
	if Ul_Semi_StaticChAccessDependentConfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ul_Semi_StaticChAccessIndependentConfig_r17Present bool
	if Ul_Semi_StaticChAccessIndependentConfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Ul_Semi_StaticChAccessDependentConfig_r17Present {
		ie.Ul_Semi_StaticChAccessDependentConfig_r17 = new(SharedSpectrumChAccessParamsPerBand_v1710_ul_Semi_StaticChAccessDependentConfig_r17)
		if err = ie.Ul_Semi_StaticChAccessDependentConfig_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ul_Semi_StaticChAccessDependentConfig_r17", err)
		}
	}
	if Ul_Semi_StaticChAccessIndependentConfig_r17Present {
		ie.Ul_Semi_StaticChAccessIndependentConfig_r17 = new(SharedSpectrumChAccessParamsPerBand_v1710_ul_Semi_StaticChAccessIndependentConfig_r17)
		if err = ie.Ul_Semi_StaticChAccessIndependentConfig_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ul_Semi_StaticChAccessIndependentConfig_r17", err)
		}
	}
	return nil
}
