package ecode

import (
	"fmt"
	"testing"
)

func TestECodeError(t *testing.T) {
	testCases := []struct {
		err error
		str string
	}{
		{err: ECode(0), str: "unknown error (0)"},
		{err: ErrNullPointer, str: "Null reference instance"},
		{err: ErrNoCommand, str: "No command"},
		{err: ErrInvalidAPIParameter, str: "Invalid API parameters"},
		{err: ErrInvalidAPIResponse, str: "Invalid response data from API"},
		{err: ErrNoData, str: "No response data"},
		{err: ErrUnMatchServiceType, str: "Unmatch service type"},
		{err: ECode(7), str: "unknown error (7)"},
	}

	for _, tc := range testCases {
		errStr := tc.err.Error()
		if errStr != tc.str {
			t.Errorf("\"%v\" != \"%v\"", errStr, tc.str)
		}
		fmt.Printf("Info(TestECodeError): %+v\n", tc.err)
	}
}

/* Copyright 2019,2020 Spiegel
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
* 	http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */
