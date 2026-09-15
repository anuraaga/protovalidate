// Copyright 2023-2026 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cases

import (
	"github.com/bufbuild/protovalidate/tools/internal/gen/buf/validate/conformance/cases"
	"github.com/bufbuild/protovalidate/tools/internal/gen/buf/validate/conformance/cases/custom_rules"
	"github.com/bufbuild/protovalidate/tools/protovalidate-conformance/internal/results"
	"github.com/bufbuild/protovalidate/tools/protovalidate-conformance/internal/suites"
)

// compilationErrorsSuite collects every case whose rules are expected to fail
// compilation. They are kept in a suite of their own, backed by proto files
// that hold nothing else, so that the file descriptor set of every other suite
// can be compiled in full. An implementation is free to compile the rules of
// each descriptor it is sent eagerly, up front.
//
// A case belongs here, and only here, if its expected result is a compilation
// error.
func compilationErrorsSuite() suites.Suite {
	return suites.Suite{
		"float/wrong_type": {
			Message:  &cases.FloatIncorrectType{Val: 123},
			Expected: results.CompilationError("double rules on float field"),
		},
		"double/wrong_type": {
			Message:  &cases.DoubleIncorrectType{Val: 123},
			Expected: results.CompilationError("double rules on float field"),
		},
		"int32/wrong_type": {
			Message:  &cases.Int32IncorrectType{Val: 123},
			Expected: results.CompilationError("double rules on float field"),
		},
		"int64/wrong_type": {
			Message:  &cases.Int64IncorrectType{Val: 123},
			Expected: results.CompilationError("double rules on float field"),
		},
		"uint32/wrong_type": {
			Message:  &cases.UInt32IncorrectType{Val: 123},
			Expected: results.CompilationError("double rules on float field"),
		},
		"uint64/wrong_type": {
			Message:  &cases.UInt64IncorrectType{Val: 123},
			Expected: results.CompilationError("double rules on float field"),
		},
		"sint32/wrong_type": {
			Message:  &cases.SInt32IncorrectType{Val: 123},
			Expected: results.CompilationError("double rules on float field"),
		},
		"sint64/wrong_type": {
			Message:  &cases.SInt64IncorrectType{Val: 123},
			Expected: results.CompilationError("double rules on float field"),
		},
		"fixed32/wrong_type": {
			Message:  &cases.Fixed32IncorrectType{Val: 123},
			Expected: results.CompilationError("double rules on float field"),
		},
		"fixed64/wrong_type": {
			Message:  &cases.Fixed64IncorrectType{Val: 123},
			Expected: results.CompilationError("double rules on float field"),
		},
		"message/oneof/unknown-field": {
			Message:  &cases.MessageOneofUnknownFieldName{},
			Expected: results.CompilationError("field xxx not found in message buf.validate.conformance.cases.MessageOneofUnknownFieldName"),
		},
		"message/oneof/duplicate-field": {
			Message:  &cases.MessageOneofDuplicateField{},
			Expected: results.CompilationError("duplicate str_field in oneof rule for the message buf.validate.conformance.cases.MessageOneofDuplicateField"),
		},
		"message/oneof/zero-fields": {
			Message:  &cases.MessageOneofZeroFields{},
			Expected: results.CompilationError("at least one field must be specified in oneof rule for the message buf.validate.conformance.cases.MessageOneofZeroFields"),
		},
		"any/wrong_type/scalar": {
			Message:  &cases.AnyWrongTypeScalar{},
			Expected: results.CompilationError("mismatched rule type and field type"),
		},
		"any/wrong_type/message": {
			Message:  &cases.AnyWrongTypeMessage{},
			Expected: results.CompilationError("mismatched rule type and field type"),
		},
		"any/wrong_type/wrapper": {
			Message:  &cases.AnyWrongTypeWrapper{},
			Expected: results.CompilationError("mismatched rule type and field type"),
		},
		"any/wrong_type/wkt": {
			Message:  &cases.AnyWrongTypeWKT{},
			Expected: results.CompilationError("mismatched rule type and field type"),
		},
		"duration/wrong_type/scalar": {
			Message:  &cases.DurationWrongTypeScalar{},
			Expected: results.CompilationError("mismatched rule type and field type"),
		},
		"duration/wrong_type/message": {
			Message:  &cases.DurationWrongTypeMessage{},
			Expected: results.CompilationError("mismatched rule type and field type"),
		},
		"duration/wrong_type/wrapper": {
			Message:  &cases.DurationWrongTypeWrapper{},
			Expected: results.CompilationError("mismatched rule type and field type"),
		},
		"duration/wrong_type/wkt": {
			Message:  &cases.DurationWrongTypeWKT{},
			Expected: results.CompilationError("mismatched rule type and field type"),
		},
		"timestamp/wrong_type/scalar": {
			Message:  &cases.TimestampWrongTypeScalar{},
			Expected: results.CompilationError("mismatched rule type and field type"),
		},
		"timestamp/wrong_type/message": {
			Message:  &cases.TimestampWrongTypeMessage{},
			Expected: results.CompilationError("mismatched rule type and field type"),
		},
		"timestamp/wrong_type/wrapper": {
			Message:  &cases.TimestampWrongTypeWrapper{},
			Expected: results.CompilationError("mismatched rule type and field type"),
		},
		"timestamp/wrong_type/wkt": {
			Message:  &cases.TimestampWrongTypeWKT{},
			Expected: results.CompilationError("mismatched rule type and field type"),
		},
		"custom_rules/missing_field": {
			Message: &custom_rules.MissingField{A: 123},
			Expected: results.CompilationError(
				"expression references a non-existent field b"),
		},
		"custom_rules/incorrect_type": {
			Message: &custom_rules.IncorrectType{A: 123},
			Expected: results.CompilationError(
				"expression incorrectly treats an int32 field as a string"),
		},
	}
}
