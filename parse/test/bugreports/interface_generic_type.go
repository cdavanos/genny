package bugreports

import "github.com/cdavanos/genny/generic"

type GenericType generic.Type

type InterfaceGenericType interface {
	DoSomthingGenericType()
}

// Call calls a method on an instance of generic interface.
// Targets github issue 49
func CallWithGenericType(i InterfaceGenericType) {
	i.DoSomthingGenericType()
}
