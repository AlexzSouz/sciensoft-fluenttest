package tdd

import (
	"reflect"
	"testing"

	"github.com/sciensoft/fluenttests/fluent/contracts"
)

type Cat struct{}

func (c Cat) MakeSound() string {
	return "Meow"
}

func TestFluentContractsShouldBeNil(t *testing.T) {
	// Arrange
	fluent := contracts.Fluent[any](t)
	var cat Cat
	// var animal interface{} = cat

	// Act
	// ... Noop

	// Assert
	fluent.It(cat).Should().BeNilOrZero()
}

func TestFluentContractsShouldBeNotNil(t *testing.T) {
	// Arrange
	fluent := contracts.Fluent[any](t)
	var obj any = struct {
		A int
	}{
		A: 1,
	}

	// Act
	// ... Noop

	// Assert
	fluent.It(obj).Should().NotBeNilOrZero()
}

func TestFluentContractsShouldBeOfType(t *testing.T) {
	// Arrange
	fluent := contracts.Fluent[any](t)
	obj := struct {
		A int
		B any
	}{
		A: 1,
		B: struct {
			B1 string
		}{
			B1: "Hello world!",
		},
	}
	objType := reflect.TypeOf(struct {
		A int
		B any
	}{})
	fieldType := reflect.TypeOf(int(0))

	// Act
	// ... Noop

	// Assert
	fluent.It(obj).
		Should().BeOfType(objType).
		And().HaveField("A").OfType(fieldType).WithValue(1).
		And().HaveField("B").WithValue(struct {
		B1 string
	}{
		B1: "Hello world!",
	})
}

func TestFluentContractsShouldNotBeOfType(t *testing.T) {
	// Arrange
	fluent := contracts.Fluent[any](t)
	obj1 := struct{ A int }{A: 1}
	obj2 := struct{ B int }{B: 1}
	rtype := reflect.TypeOf(obj2)

	// Act
	// ... Noop

	// Assert
	fluent.It(obj1).
		Should().NotBeOfType(rtype)
}

func TestFluentContractsShouldHaveMember(t *testing.T) {
	// Arrange
	fluent := contracts.Fluent[any](t)
	obj := struct{ Name string }{Name: "Robert Griesemer"}

	// Act
	// ... Noop

	// Assert
	fluent.It(obj).
		Should().HaveMember("Name")
}

func TestFluentContractsShouldHaveField(t *testing.T) {
	// Arrange
	fluent := contracts.Fluent[any](t)
	obj := struct{ Value float32 }{Value: 105.4}

	// Act
	// ... Noop

	// Assert
	fluent.It(obj).
		Should().HaveField("Value")
}

func TestFluentContractsShouldHaveFieldWithTag(t *testing.T) {
	// Arrange
	fluent := contracts.Fluent[any](t)
	obj := struct {
		Value float32 `json:"value"`
	}{
		Value: 105.4,
	}

	// Act
	// ... Noop

	// Assert
	fluent.It(obj).
		Should().HaveFieldWithTag("Value", "json").
		And().HaveFieldWithTagPattern("Value", `json:"value"`)
}

func TestFluentContractsShouldHaveAllFieldsWithTag(t *testing.T) {
	// Arrange
	fluent := contracts.Fluent[any](t)
	obj := struct {
		Title string  `json:"title"`
		Value float32 `json:"value"`
	}{
		Title: "Book Xyz",
		Value: 105.4,
	}

	// Act
	// ... Noop

	// Assert
	fluent.It(obj).
		Should().HaveAllFieldsWithTag("json")
}

func TestFluentContractsShouldHaveMethod(t *testing.T) {
	// Arrange
	fluent := contracts.Fluent[any](t)

	// Act
	// ... Noop

	// Assert
	fluent.It(fluent).
		Should().HaveMethod("It")
}

func TestFluentContractsShouldHaveAnyOfMembers(t *testing.T) {
	// Arrange
	fluent := contracts.Fluent[any](t)
	obj := struct {
		Title     string
		Author    string
		Publisher string
	}{
		Title:     "A Concurrent Window System",
		Author:    "Rob Pike",
		Publisher: "AT&T Bell Laboratories",
	}

	// Act
	// ... Noop

	// Assert
	fluent.It(obj).
		Should().
		HaveAnyOfMembers([]string{
			"Title",
			"Description",
			"PublicationDate",
		})
}

func TestFluentContractsShouldHaveAllOfMembers(t *testing.T) {
	// Arrange
	fluent := contracts.Fluent[any](t)
	obj := struct {
		Title     string
		Author    string
		Publisher string
	}{
		Title:     "A Concurrent Window System",
		Author:    "Rob Pike",
		Publisher: "AT&T Bell Laboratories",
	}

	// Act
	// ... Noop

	// Assert
	fluent.It(obj).
		Should().
		HaveAllOfMembers([]string{
			"Title",
			"Author",
			"Publisher",
		})
}
