package contracts

import (
	"reflect"

	f "github.com/sciensoft/fluenttests/fluent"
)

// BeOfType asserts a Mock objects as being of the provided type "typeName reflect.Type" argument.
func (c *Comparable[T]) BeOfType(typeName reflect.Type) f.IAdditional[T, IComparable[T]] {
	ofType(c.testingT, f.NegativeDefault, c.value, typeName)
	return c.createAdditional()
}

// NotBeOfType asserts a Mock objects as NOT being of the provided type "typeName reflect.Type" argument.
func (c *Comparable[T]) NotBeOfType(typeName reflect.Type) f.IAdditional[T, IComparable[T]] {
	ofType(c.testingT, f.NegativeInvert, c.value, typeName, "Type [%T] is to not be of type [%v].")
	return c.createAdditional()
}

// HaveMember asserts a Mock objects as having a member, either a method or field, named as per "memberName string" argument.
func (c *Comparable[T]) HaveMember(memberName string) f.IAdditional[T, IComparable[T]] {
	haveMembers(c.testingT, f.NegativeDefault, f.MatchAny, f.MemberTypeAll, c.value, []string{memberName})
	return c.createAdditional()
}

// HaveField asserts a Mock objects as having a field named as per "fieldName string" argument.
func (c *Comparable[T]) HaveField(fieldName string) IAdditionalWith[T] {
	haveMembers(c.testingT, f.NegativeDefault, f.MatchAny, f.MemberTypeField, c.value, []string{fieldName}, "Object of [%T] does not have a field called [%v].")

	return c.createAdditionalWith(params{
		ptype: f.ParamType.Field,
		value: fieldName,
	})
}

// HaveFieldWithTag asserts a Mock objects as having a field named as per "fieldName string" argument
// and with tag name as per "tagName string" argument.
func (c *Comparable[T]) HaveFieldWithTag(fieldName string, tagName string) f.IAdditional[T, IComparable[T]] {
	haveFieldWithTag(c.testingT, c.value, fieldName, tagName, "....")
	return c.createAdditional()
}

// HaveFieldWithTagPattern asserts a Mock objects as having a field named as per "fieldName string" argument
// and with tag patterns as per "tagPattern string" argument.
func (c *Comparable[T]) HaveFieldWithTagPattern(fieldName string, tagPattern string) f.IAdditional[T, IComparable[T]] {
	haveFieldWithTag(c.testingT, c.value, fieldName, tagPattern, "....")
	return c.createAdditional()
}

// HaveAllFieldsWithTag asserts a Mock objects as having a all its fields
// with the tag name as per "tagName string" argument.
func (c *Comparable[T]) HaveAllFieldsWithTag(tagName string) f.IAdditional[T, IComparable[T]] {
	haveAllFieldsWithTag(c.testingT, c.value, tagName, "....")
	return c.createAdditional()
}

// HaveMethod asserts a Mock objects as having a method named as per "methodName string" argument.
func (c *Comparable[T]) HaveMethod(methodName string) f.IAdditional[T, IComparable[T]] {
	haveMembers(c.testingT, f.NegativeDefault, f.MatchAny, f.MemberTypeMethod, c.value, []string{methodName}, "Object of [%T] does not have a method called [%v].")
	return c.createAdditional()
}

// HaveAnyOfMembers asserts a Mock objects as having any of its named members (methods & fields) as per "membersNames []string" argument.
func (c *Comparable[T]) HaveAnyOfMembers(membersNames []string) f.IAdditional[T, IComparable[T]] {
	haveMembers(c.testingT, f.NegativeDefault, f.MatchAny, f.MemberTypeAll, c.value, membersNames, "Object of [%T] does not have any of members [%v].")
	return c.createAdditional()
}

// HaveAllOfMembers asserts a Mock objects as having all of its named members (methods & fields) as per "membersNames []string" argument.
func (c *Comparable[T]) HaveAllOfMembers(membersNames []string) f.IAdditional[T, IComparable[T]] {
	haveMembers(c.testingT, f.NegativeDefault, f.MatchAll, f.MemberTypeAll, c.value, membersNames, "Object of [%T] does not have all of members [%v].")
	return c.createAdditional()
}

// createAdditional supports the creation of chained easy-to-use testing methods.
func (c *Comparable[T]) createAdditional() f.IAdditional[T, IComparable[T]] {
	additional := &Additional[T]{
		Comparable: c,
	}

	return additional
}

// createAdditionalWith supports the creation of chained easy-to-use testing methods.
func (c *Comparable[T]) createAdditionalWith(params ...params) IAdditionalWith[T] {
	additional := &AdditionalWith[T]{
		Additional: &Additional[T]{
			Comparable: c,
		},
		params: params,
	}

	return additional
}
