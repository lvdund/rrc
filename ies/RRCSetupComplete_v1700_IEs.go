package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCSetupComplete_v1700_IEs struct {
	OnboardingRequest_r17 *RRCSetupComplete_v1700_IEs_onboardingRequest_r17 `optional`
	NonCriticalExtension  interface{}                                       `optional`
}

func (ie *RRCSetupComplete_v1700_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.OnboardingRequest_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.OnboardingRequest_r17 != nil {
		if err = ie.OnboardingRequest_r17.Encode(w); err != nil {
			return utils.WrapError("Encode OnboardingRequest_r17", err)
		}
	}
	return nil
}

func (ie *RRCSetupComplete_v1700_IEs) Decode(r *uper.UperReader) error {
	var err error
	var OnboardingRequest_r17Present bool
	if OnboardingRequest_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if OnboardingRequest_r17Present {
		ie.OnboardingRequest_r17 = new(RRCSetupComplete_v1700_IEs_onboardingRequest_r17)
		if err = ie.OnboardingRequest_r17.Decode(r); err != nil {
			return utils.WrapError("Decode OnboardingRequest_r17", err)
		}
	}
	return nil
}
