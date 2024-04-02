package contracts

import (
	"reflect"
	"regexp"
	"testing"

	f "github.com/sciensoft/fluenttests/fluent"
)

func beNil[T any](t *testing.T, invert f.AdditiveInverse, value T, messagesf ...string) {
	var vi interface{} = value
	isNil := vi == nil

	if invert {
		isNil = !isNil
	}

	if !isNil {
		message := f.GetMessage("Expected value to be 'nil' but got %q.", messagesf...)
		t.Errorf(message, value)
	}
}

func ofType[T any](t *testing.T, invert f.AdditiveInverse, value T, comparable reflect.Type, messagesf ...string) {
	vtype := reflect.TypeOf(value)
	isOfType := vtype == comparable

	if invert {
		isOfType = !isOfType
	}

	if !isOfType {
		message := f.GetMessage("Expected type %q is not of type %q.", messagesf...)
		t.Errorf(message, vtype, comparable)
	}
}

func haveMembers[T any](t *testing.T, invert f.AdditiveInverse, matchType f.Match, mType f.MemberType, value T, comparable []string, messagesf ...string) {
	vtype := reflect.TypeOf(value)
	hasMember := false
	matchAll := matchType == f.MatchAll

	for _, member := range comparable {
		hasField := testMemberByName(f.MemberTypeField, vtype, member)
		hasMethod := testMemberByName(f.MemberTypeMethod, vtype, member)

		if matchAll || !hasMember {
			switch mType {
			case f.MemberTypeField:
				hasMember = hasField
			case f.MemberTypeMethod:
				hasMember = hasMethod
			default:
				hasMember = hasField || hasMethod
			}
		}

		if matchAll && !hasMember {
			break
		}
	}

	if invert {
		hasMember = !hasMember
	}

	if !hasMember {
		message := f.GetMessage("Object of %q does not have a member called %q.", messagesf...)
		t.Errorf(message, vtype, comparable)
	}
}

func haveFieldWithTag[T any](t *testing.T, value T, comparableFieldName string, comparableTagName string, messagesf ...string) {
	vtype := reflect.TypeOf(value)
	if vtype.Kind() == reflect.Ptr {
		vtype = vtype.Elem()
	}

	field, hasField := vtype.FieldByName(comparableFieldName)

	if hasField {
		_, hasTag := field.Tag.Lookup(comparableTagName)

		if !hasTag {
			regex, _ := regexp.Compile(comparableTagName)
			hasTag = regex.Match([]byte(field.Tag))
		}

		if !hasTag {
			message := f.GetMessage("Object of %q does not have a member %q with tag %q.", messagesf...)
			t.Errorf(message, vtype, comparableFieldName, comparableTagName)
		}
	}
}

func haveAllFieldsWithTag[T any](t *testing.T, value T, comparableTagName string, messagesf ...string) {
	vtype := reflect.TypeOf(value)
	if vtype.Kind() == reflect.Ptr {
		vtype = vtype.Elem()
	}

	numFields := vtype.NumField()

	for i := 0; i < numFields; i++ {
		field := vtype.Field(i)
		_, hasTag := field.Tag.Lookup(comparableTagName)

		if !hasTag {
			message := f.GetMessage("Object of %q does not have all fields with tag %q.", messagesf...)
			t.Errorf(message, vtype, comparableTagName)
		}
	}
}

func testMemberByName(mType f.MemberType, vtype reflect.Type, memberName string) bool {
	defer func() bool {
		recover()
		return false
	}()

	if mType == f.MemberTypeField {
		if vtype.Kind() == reflect.Ptr {
			vtype = vtype.Elem()
		}
		_, hasField := vtype.FieldByName(memberName)
		return hasField
	} else {
		_, hasMethod := vtype.MethodByName(memberName)
		return hasMethod
	}
}
