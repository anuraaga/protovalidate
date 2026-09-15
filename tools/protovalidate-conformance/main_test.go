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

package main

import (
	"testing"

	"github.com/bufbuild/protovalidate/tools/internal/gen/buf/validate/conformance/harness"
	"github.com/bufbuild/protovalidate/tools/protovalidate-conformance/internal/cases"
	"github.com/bufbuild/protovalidate/tools/protovalidate-conformance/internal/suites"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// compilationErrorsSuiteName is the one suite whose cases are expected to fail
// compilation.
const compilationErrorsSuiteName = "compilation_errors"

func TestProcessAllSuites(t *testing.T) {
	t.Parallel()
	// Conformance test cases can be invalid: For example, they can define a bad
	// field path, or a message that fails to marshal.
	err := cases.GlobalSuites().Range(nil, func(suiteName string, suite suites.Suite) error {
		_, err := suite.ToRequestProto(nil)
		if err != nil {
			return err
		}
		_, err = suite.ProcessResults(suiteName, nil, nil, nil, nil)
		return err
	})
	if err != nil {
		t.Error(err)
	}
}

// TestCompilationErrorsAreIsolated asserts that only the compilation_errors
// suite expects compilation errors, and that no descriptor it sends leaks into
// another suite's file descriptor set. An implementation can therefore eagerly
// compile the rules of every descriptor in the sets of all other suites.
func TestCompilationErrorsAreIsolated(t *testing.T) {
	t.Parallel()
	failing := map[string]struct{}{}
	err := cases.GlobalSuites().Range(nil, func(suiteName string, suite suites.Suite) error {
		return suite.Range(nil, func(caseName string, testCase suites.Case) error {
			if _, ok := testCase.Expected.ToProto().GetResult().(*harness.TestResult_CompilationError); !ok {
				return nil
			}
			assert.Equal(t, compilationErrorsSuiteName, suiteName,
				"case %q expects a compilation error, so it belongs in the %q suite",
				caseName, compilationErrorsSuiteName)
			failing[string(testCase.Message.ProtoReflect().Descriptor().FullName())] = struct{}{}
			return nil
		})
	})
	require.NoError(t, err)
	require.NotEmpty(t, failing, "no compilation error cases found")

	err = cases.GlobalSuites().Range(nil, func(suiteName string, suite suites.Suite) error {
		if suiteName == compilationErrorsSuiteName {
			return nil
		}
		req, err := suite.ToRequestProto(nil)
		require.NoError(t, err)
		files, err := protodesc.NewFiles(req.GetFdset())
		require.NoError(t, err, "%s: file descriptor set is not self-contained", suiteName)
		files.RangeFiles(func(file protoreflect.FileDescriptor) bool {
			rangeMessages(file.Messages(), func(desc protoreflect.MessageDescriptor) {
				assert.NotContains(t, failing, string(desc.FullName()),
					"%s: descriptor expected to fail compilation is in a suite whose rules must compile",
					suiteName)
			})
			return true
		})
		return nil
	})
	require.NoError(t, err)
}

func rangeMessages(descs protoreflect.MessageDescriptors, fn func(protoreflect.MessageDescriptor)) {
	for i := range descs.Len() {
		desc := descs.Get(i)
		fn(desc)
		rangeMessages(desc.Messages(), fn)
	}
}
