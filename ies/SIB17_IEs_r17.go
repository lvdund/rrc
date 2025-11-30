package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SIB17_IEs_r17 struct {
	Trs_ResourceSetConfig_r17 []TRS_ResourceSet_r17               `lb:1,ub:maxNrofTRS_ResourceSets_r17,madatory`
	ValidityDuration_r17      *SIB17_IEs_r17_validityDuration_r17 `optional`
	LateNonCriticalExtension  *[]byte                             `optional`
}

func (ie *SIB17_IEs_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.ValidityDuration_r17 != nil, ie.LateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	tmp_Trs_ResourceSetConfig_r17 := utils.NewSequence[*TRS_ResourceSet_r17]([]*TRS_ResourceSet_r17{}, aper.Constraint{Lb: 1, Ub: maxNrofTRS_ResourceSets_r17}, false)
	for _, i := range ie.Trs_ResourceSetConfig_r17 {
		tmp_Trs_ResourceSetConfig_r17.Value = append(tmp_Trs_ResourceSetConfig_r17.Value, &i)
	}
	if err = tmp_Trs_ResourceSetConfig_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Trs_ResourceSetConfig_r17", err)
	}
	if ie.ValidityDuration_r17 != nil {
		if err = ie.ValidityDuration_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ValidityDuration_r17", err)
		}
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, nil, false); err != nil {
			return utils.WrapError("Encode LateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *SIB17_IEs_r17) Decode(r *aper.AperReader) error {
	var err error
	var ValidityDuration_r17Present bool
	if ValidityDuration_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var LateNonCriticalExtensionPresent bool
	if LateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	tmp_Trs_ResourceSetConfig_r17 := utils.NewSequence[*TRS_ResourceSet_r17]([]*TRS_ResourceSet_r17{}, aper.Constraint{Lb: 1, Ub: maxNrofTRS_ResourceSets_r17}, false)
	fn_Trs_ResourceSetConfig_r17 := func() *TRS_ResourceSet_r17 {
		return new(TRS_ResourceSet_r17)
	}
	if err = tmp_Trs_ResourceSetConfig_r17.Decode(r, fn_Trs_ResourceSetConfig_r17); err != nil {
		return utils.WrapError("Decode Trs_ResourceSetConfig_r17", err)
	}
	ie.Trs_ResourceSetConfig_r17 = []TRS_ResourceSet_r17{}
	for _, i := range tmp_Trs_ResourceSetConfig_r17.Value {
		ie.Trs_ResourceSetConfig_r17 = append(ie.Trs_ResourceSetConfig_r17, *i)
	}
	if ValidityDuration_r17Present {
		ie.ValidityDuration_r17 = new(SIB17_IEs_r17_validityDuration_r17)
		if err = ie.ValidityDuration_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ValidityDuration_r17", err)
		}
	}
	if LateNonCriticalExtensionPresent {
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = &tmp_os_LateNonCriticalExtension
	}
	return nil
}
