package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SecurityAlgorithmConfig struct {
	CipheringAlgorithm     CipheringAlgorithm      `madatory`
	IntegrityProtAlgorithm *IntegrityProtAlgorithm `optional`
}

func (ie *SecurityAlgorithmConfig) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.IntegrityProtAlgorithm != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.CipheringAlgorithm.Encode(w); err != nil {
		return utils.WrapError("Encode CipheringAlgorithm", err)
	}
	if ie.IntegrityProtAlgorithm != nil {
		if err = ie.IntegrityProtAlgorithm.Encode(w); err != nil {
			return utils.WrapError("Encode IntegrityProtAlgorithm", err)
		}
	}
	return nil
}

func (ie *SecurityAlgorithmConfig) Decode(r *aper.AperReader) error {
	var err error
	var IntegrityProtAlgorithmPresent bool
	if IntegrityProtAlgorithmPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.CipheringAlgorithm.Decode(r); err != nil {
		return utils.WrapError("Decode CipheringAlgorithm", err)
	}
	if IntegrityProtAlgorithmPresent {
		ie.IntegrityProtAlgorithm = new(IntegrityProtAlgorithm)
		if err = ie.IntegrityProtAlgorithm.Decode(r); err != nil {
			return utils.WrapError("Decode IntegrityProtAlgorithm", err)
		}
	}
	return nil
}
