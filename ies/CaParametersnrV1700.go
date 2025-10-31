package ies

// CA-ParametersNR-v1700 ::= SEQUENCE
type CaParametersnrV1700 struct {
	Codebookparametersfetype2perbcR17         *Codebookparametersfetype2perbcR17
	DemodulationenhancementcaR17              *CaParametersnrV1700DemodulationenhancementcaR17
	MaxuplinkdutycycleInterbandcaPc2R17       *CaParametersnrV1700MaxuplinkdutycycleInterbandcaPc2R17
	MaxuplinkdutycycleSulcombinationPc2R17    *CaParametersnrV1700MaxuplinkdutycycleSulcombinationPc2R17
	BeammanagementtypeCbmR17                  *CaParametersnrV1700BeammanagementtypeCbmR17
	ParalleltxpucchPuschR17                   *CaParametersnrV1700ParalleltxpucchPuschR17
	CodebookcomboparametermixedtypeperbcR17   *CodebookcomboparametermixedtypeperbcR17
	MtrpCsiEnhancementperbcR17                *CaParametersnrV1700MtrpCsiEnhancementperbcR17
	CodebookcomboparametermultitrpPerbcR17    *CodebookcomboparametermultitrpPerbcR17
	Maxcc32DlHarqProcessfr22R17               *CaParametersnrV1700Maxcc32DlHarqProcessfr22R17
	Maxcc32UlHarqProcessfr22R17               *CaParametersnrV1700Maxcc32UlHarqProcessfr22R17
	CrosscarrierschedulingscellSpcelltypebR17 *CrosscarrierschedulingscellSpcellR17
	CrosscarrierschedulingscellSpcelltypeaR17 *CrosscarrierschedulingscellSpcellR17
	DciFormatspcellpscellussSetsR17           *CaParametersnrV1700DciFormatspcellpscellussSetsR17
	DisablingscalingfactordeactscellR17       *CaParametersnrV1700DisablingscalingfactordeactscellR17
	DisablingscalingfactordormantscellR17     *CaParametersnrV1700DisablingscalingfactordormantscellR17
	NonAlignedframeboundariesR17              *CaParametersnrV1700NonAlignedframeboundariesR17
}
