package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BeamManagementSSB_CSI_RS struct {
	MaxNumberSSB_CSI_RS_ResourceOneTx BeamManagementSSB_CSI_RS_maxNumberSSB_CSI_RS_ResourceOneTx `madatory`
	MaxNumberCSI_RS_Resource          BeamManagementSSB_CSI_RS_maxNumberCSI_RS_Resource          `madatory`
	MaxNumberCSI_RS_ResourceTwoTx     BeamManagementSSB_CSI_RS_maxNumberCSI_RS_ResourceTwoTx     `madatory`
	SupportedCSI_RS_Density           *BeamManagementSSB_CSI_RS_supportedCSI_RS_Density          `optional`
	MaxNumberAperiodicCSI_RS_Resource BeamManagementSSB_CSI_RS_maxNumberAperiodicCSI_RS_Resource `madatory`
}

func (ie *BeamManagementSSB_CSI_RS) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.SupportedCSI_RS_Density != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.MaxNumberSSB_CSI_RS_ResourceOneTx.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberSSB_CSI_RS_ResourceOneTx", err)
	}
	if err = ie.MaxNumberCSI_RS_Resource.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberCSI_RS_Resource", err)
	}
	if err = ie.MaxNumberCSI_RS_ResourceTwoTx.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberCSI_RS_ResourceTwoTx", err)
	}
	if ie.SupportedCSI_RS_Density != nil {
		if err = ie.SupportedCSI_RS_Density.Encode(w); err != nil {
			return utils.WrapError("Encode SupportedCSI_RS_Density", err)
		}
	}
	if err = ie.MaxNumberAperiodicCSI_RS_Resource.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberAperiodicCSI_RS_Resource", err)
	}
	return nil
}

func (ie *BeamManagementSSB_CSI_RS) Decode(r *aper.AperReader) error {
	var err error
	var SupportedCSI_RS_DensityPresent bool
	if SupportedCSI_RS_DensityPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.MaxNumberSSB_CSI_RS_ResourceOneTx.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberSSB_CSI_RS_ResourceOneTx", err)
	}
	if err = ie.MaxNumberCSI_RS_Resource.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberCSI_RS_Resource", err)
	}
	if err = ie.MaxNumberCSI_RS_ResourceTwoTx.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberCSI_RS_ResourceTwoTx", err)
	}
	if SupportedCSI_RS_DensityPresent {
		ie.SupportedCSI_RS_Density = new(BeamManagementSSB_CSI_RS_supportedCSI_RS_Density)
		if err = ie.SupportedCSI_RS_Density.Decode(r); err != nil {
			return utils.WrapError("Decode SupportedCSI_RS_Density", err)
		}
	}
	if err = ie.MaxNumberAperiodicCSI_RS_Resource.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberAperiodicCSI_RS_Resource", err)
	}
	return nil
}
