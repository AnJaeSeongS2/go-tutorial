package main

import (
	"github.com/anlue/go/tutorial/basics"
	"github.com/anlue/go/tutorial/flowcontrol"
)

func main() {
	basics.TourPackage()
	basics.TourImports()
	basics.TourExportedNames()
	basics.TourFunctions()
	basics.TourFunctionsContinued()
	basics.TourMultipleResults()
	basics.TourNamedResults()
	basics.TourVariables()
	basics.TourVariablesWithInitializers()
	basics.TourShortVariableDeclarations()
	basics.TourBasicTypes()
	basics.TourZero()
	basics.TourTypeConversions()
	basics.TourTypeInference()
	basics.TourConstants()
	basics.TourNumericConstants()

	flowcontrol.TourFor()
	flowcontrol.TourIf()
	flowcontrol.TourIfWithAShortStatement()
	flowcontrol.TourSwitch()
	flowcontrol.TourDefer()
}
