package ies

// CA-ParametersNRDC ::= SEQUENCE
type CaParametersnrdc struct {
	CaParametersnrFordc      *CaParametersnr
	CaParametersnrFordcV1540 *CaParametersnrV1540
	CaParametersnrFordcV1550 *CaParametersnrV1550
	CaParametersnrFordcV1560 *CaParametersnrV1560
	Featuresetcombinationdc  *Featuresetcombinationid
}
