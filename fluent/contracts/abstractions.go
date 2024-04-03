package contracts

import (
	"reflect"
	"testing"

	"github.com/sciensoft/fluenttests/fluent"
)

// IComparable[T any] interface for the contracts package. Following
// the Fluent-Interface design pattern to provide chained testing methods.
type IComparable[T any] interface {
	fluent.IComparable[T]
	BeNilOrZero() fluent.IAdditional[T, IComparable[T]]
	NotBeNilOrZero() fluent.IAdditional[T, IComparable[T]]
	BeOfType(typeName reflect.Type) fluent.IAdditional[T, IComparable[T]]
	NotBeOfType(typeName reflect.Type) fluent.IAdditional[T, IComparable[T]]
	HaveMember(memberName string) fluent.IAdditional[T, IComparable[T]]
	HaveField(fieldName string) IAdditionalWith[T]
	HaveFieldWithTag(fieldName string, tagName string) fluent.IAdditional[T, IComparable[T]]
	HaveFieldWithTagPattern(fieldName string, tagPattern string) fluent.IAdditional[T, IComparable[T]]
	HaveAllFieldsWithTag(tagName string) fluent.IAdditional[T, IComparable[T]]
	HaveMethod(methodName string) fluent.IAdditional[T, IComparable[T]]
	HaveAnyOfMembers(membersNames []string) fluent.IAdditional[T, IComparable[T]]
	HaveAllOfMembers(membersNames []string) fluent.IAdditional[T, IComparable[T]]
	// Add more supporting methods here...
}

// IAdditionalWith[T any] interface for the contracts package.
type IAdditionalWith[T any] interface {
	fluent.IAdditional[T, IComparable[T]]
	OfType(ctype reflect.Type) IAdditionalWith[T]
	WithValue(value any) IAdditionalWith[T]
}

// FluentT[T any] is a type returned by calling "contracts.Fluent[any](t)"
// for enabling easy-to-use testing methods.
//
// It carries the "*testing.T" object to be further used in the methods chain.
type FluentT[T any] struct {
	testingT *testing.T
}

// Subject[T any] is a type returned by calling "fluentObj.It(myMockObj)"
// used to encapsulate the "*testing.T" object and testable subject.
type Subject[T any] struct {
	*Testable[T]
}

// Additional[T any] is a type returned by some easy-to-use testing methods,
// providing basic additive test chaining with "And()".
type Additional[T any] struct {
	*Comparable[T]
}

// AdditionalWith[T any] is a type returned by some easy-to-use testing methods,
// providing some extra additive test chaining with "OfType()", and "WithValue()".
type AdditionalWith[T any] struct {
	*Additional[T]
	params []params
}

// Comparable[T any] is a type returned by calling "fluentObj.It(myMockObj).Should()",
// encapsulating a "*Testable[T]" object for further testing, and providing
// all the contracts package's easy-to-use set of methods with signatures declared at
// the "contracts.IComparable[any]" interface.
type Comparable[T any] struct {
	*Testable[T]
}

// Testable[T any] is a type that holds the "*testing.T" object and the underline
// Mock object, supporting further testing using the chain of easy-to-use methods.
// It's used by the "Subject[T any]" and "Comparable[T any]" types.
type Testable[T any] struct {
	testingT *testing.T
	value    T
}

// params is a internal supporting type for the additive operations within
// the contracts package. It records testing values and its type, so further
// additive operation can make usage them to assert an object.
type params struct {
	ptype string
	value any
}
