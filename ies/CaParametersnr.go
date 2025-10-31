package ies

// CA-ParametersNR ::= SEQUENCE
// Extensible
type CaParametersnr struct {
	Dummy                                    *CaParametersnrDummy
	ParalleltxsrsPucchPusch                  *CaParametersnrParalleltxsrsPucchPusch
	ParalleltxprachSrsPucchPusch             *CaParametersnrParalleltxprachSrsPucchPusch
	Simultaneousrxtxinterbandca              *CaParametersnrSimultaneousrxtxinterbandca
	Simultaneousrxtxsul                      *CaParametersnrSimultaneousrxtxsul
	DiffnumerologyacrosspucchGroup           *CaParametersnrDiffnumerologyacrosspucchGroup
	DiffnumerologywithinpucchGroupsmallerscs *CaParametersnrDiffnumerologywithinpucchGroupsmallerscs
	Supportednumbertag                       *CaParametersnrSupportednumbertag
}
