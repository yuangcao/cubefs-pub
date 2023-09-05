// Copyright 2023 The CubeFS Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package util_test

import (
	"testing"

	"github.com/cubefs/cubefs/util"
)

func TestRandomString(t *testing.T) {
	first := util.RandomString(256, util.UpperLetter|util.LowerLetter|util.Numeric)
	second := util.RandomString(256, util.UpperLetter|util.LowerLetter|util.Numeric)
	if first == second {
		t.Errorf("first should not equal with second")
		return
	}
}

func TestSubStr(t *testing.T) {
	str := "abcd"
	if util.SubString(str, 1, 2) != "b" {
		t.Errorf("SubString should returns \"b\" not got \"%v\"", util.SubString(str, 1, 2))
	}
}