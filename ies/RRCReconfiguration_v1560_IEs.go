package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCReconfiguration_v1560_IEs struct {
	Mrdc_SecondaryCellGroupConfig *MRDC_SecondaryCellGroupConfig `optional,setuprelease`
	RadioBearerConfig2            *[]byte                        `optional`
	Sk_Counter                    *SK_Counter                    `optional`
	NonCriticalExtension          *RRCReconfiguration_v1610_IEs  `optional`
}

func (ie *RRCReconfiguration_v1560_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Mrdc_SecondaryCellGroupConfig != nil, ie.RadioBearerConfig2 != nil, ie.Sk_Counter != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Mrdc_SecondaryCellGroupConfig != nil {
		tmp_Mrdc_SecondaryCellGroupConfig := utils.SetupRelease[*MRDC_SecondaryCellGroupConfig]{
			Setup: ie.Mrdc_SecondaryCellGroupConfig,
		}
		if err = tmp_Mrdc_SecondaryCellGroupConfig.Encode(w); err != nil {
			return utils.WrapError("Encode Mrdc_SecondaryCellGroupConfig", err)
		}
	}
	if ie.RadioBearerConfig2 != nil {
		if err = w.WriteOctetString(*ie.RadioBearerConfig2, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode RadioBearerConfig2", err)
		}
	}
	if ie.Sk_Counter != nil {
		if err = ie.Sk_Counter.Encode(w); err != nil {
			return utils.WrapError("Encode Sk_Counter", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *RRCReconfiguration_v1560_IEs) Decode(r *uper.UperReader) error {
	var err error
	var Mrdc_SecondaryCellGroupConfigPresent bool
	if Mrdc_SecondaryCellGroupConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var RadioBearerConfig2Present bool
	if RadioBearerConfig2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sk_CounterPresent bool
	if Sk_CounterPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Mrdc_SecondaryCellGroupConfigPresent {
		tmp_Mrdc_SecondaryCellGroupConfig := utils.SetupRelease[*MRDC_SecondaryCellGroupConfig]{}
		if err = tmp_Mrdc_SecondaryCellGroupConfig.Decode(r); err != nil {
			return utils.WrapError("Decode Mrdc_SecondaryCellGroupConfig", err)
		}
		ie.Mrdc_SecondaryCellGroupConfig = tmp_Mrdc_SecondaryCellGroupConfig.Setup
	}
	if RadioBearerConfig2Present {
		var tmp_os_RadioBearerConfig2 []byte
		if tmp_os_RadioBearerConfig2, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode RadioBearerConfig2", err)
		}
		ie.RadioBearerConfig2 = &tmp_os_RadioBearerConfig2
	}
	if Sk_CounterPresent {
		ie.Sk_Counter = new(SK_Counter)
		if err = ie.Sk_Counter.Decode(r); err != nil {
			return utils.WrapError("Decode Sk_Counter", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(RRCReconfiguration_v1610_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
