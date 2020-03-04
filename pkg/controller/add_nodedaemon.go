package controller

import (
	"github.com/VaishnaviHire/timeapp-operator/pkg/controller/nodedaemon"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, nodedaemon.Add)
}
