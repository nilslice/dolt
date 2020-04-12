// Copyright 2020 Liquidata, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package typeinfo

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/liquidata-inc/dolt/go/store/types"
)

func TestDecimalConvertNomsValueToValue(t *testing.T) {
	tests := []struct {
		typ         *decimalType
		input       types.String
		output      string
		expectedErr bool
	}{
		{
			generateDecimalType(t, 1, 0),
			"10",
			"0",
			false,
		},
		{
			generateDecimalType(t, 1, 0),
			"08.5",
			"-2",
			false,
		},
		{
			generateDecimalType(t, 2, 1),
			"08.5",
			"-1.5",
			false,
		},
		{
			generateDecimalType(t, 5, 4),
			"094.2841",
			"-5.7159",
			false,
		},
		{
			generateDecimalType(t, 9, 2),
			"14723245",
			"4723245.00",
			false,
		},
		{
			generateDecimalType(t, 9, 2),
			"100004723245.01",
			"4723245.01",
			false,
		},
		{
			generateDecimalType(t, 1, 0),
			"",
			"",
			true,
		},
		{
			generateDecimalType(t, 9, 2),
			"100014723245.01",
			"",
			true,
		},
		{
			generateDecimalType(t, 5, 4),
			"044.2841",
			"",
			true,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf(`%v %v`, test.typ.String(), test.input), func(t *testing.T) {
			output, err := test.typ.ConvertNomsValueToValue(test.input)
			if test.expectedErr {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, test.output, output)
			}
		})
	}
}

func TestDecimalConvertValueToNomsValue(t *testing.T) {
	tests := []struct {
		typ         *decimalType
		input       interface{}
		output      types.String
		expectedErr bool
	}{
		{
			generateDecimalType(t, 1, 0),
			7,
			"17",
			false,
		},
		{
			generateDecimalType(t, 5, 1),
			-4.5,
			"09995.5",
			false,
		},
		{
			generateDecimalType(t, 10, 0),
			"77",
			"10000000077",
			false,
		},
		{
			generateDecimalType(t, 5, 0),
			"dog",
			"",
			true,
		},
		{
			generateDecimalType(t, 15, 7),
			true,
			"",
			true,
		},
		{
			generateDecimalType(t, 20, 5),
			time.Unix(137849, 0),
			"",
			true,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf(`%v %v`, test.typ.String(), test.input), func(t *testing.T) {
			output, err := test.typ.ConvertValueToNomsValue(test.input)
			if !test.expectedErr {
				require.NoError(t, err)
				assert.Equal(t, test.output, output)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestDecimalFormatValue(t *testing.T) {
	tests := []struct {
		typ         *decimalType
		input       types.String
		output      string
		expectedErr bool
	}{
		{
			generateDecimalType(t, 1, 0),
			"10",
			"0",
			false,
		},
		{
			generateDecimalType(t, 1, 0),
			"08.5",
			"-2",
			false,
		},
		{
			generateDecimalType(t, 2, 1),
			"08.5",
			"-1.5",
			false,
		},
		{
			generateDecimalType(t, 5, 4),
			"094.2841",
			"-5.7159",
			false,
		},
		{
			generateDecimalType(t, 9, 2),
			"14723245",
			"4723245.00",
			false,
		},
		{
			generateDecimalType(t, 9, 2),
			"100004723245.01",
			"4723245.01",
			false,
		},
		{
			generateDecimalType(t, 1, 0),
			"",
			"",
			true,
		},
		{
			generateDecimalType(t, 9, 2),
			"100014723245.01",
			"",
			true,
		},
		{
			generateDecimalType(t, 5, 4),
			"044.2841",
			"",
			true,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf(`%v %v`, test.typ.String(), test.input), func(t *testing.T) {
			output, err := test.typ.FormatValue(test.input)
			if test.expectedErr {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, test.output, *output)
			}
		})
	}
}

func TestDecimalParseValue(t *testing.T) {
	tests := []struct {
		typ         *decimalType
		input       string
		output      types.String
		expectedErr bool
	}{
		{
			generateDecimalType(t, 1, 0),
			"0",
			"10",
			false,
		},
		{
			generateDecimalType(t, 1, 0),
			"-1.5",
			"08",
			false,
		},
		{
			generateDecimalType(t, 2, 1),
			"-1.5",
			"08.5",
			false,
		},
		{
			generateDecimalType(t, 5, 4),
			"-5.7159",
			"04.2841",
			false,
		},
		{
			generateDecimalType(t, 9, 2),
			"4723245.00",
			"14723245.00",
			false,
		},
		{
			generateDecimalType(t, 13, 2),
			"4723245.01",
			"100004723245.01",
			false,
		},
		{
			generateDecimalType(t, 9, 2),
			"24723245.01",
			"",
			true,
		},
		{
			generateDecimalType(t, 5, 4),
			"-44.2841",
			"",
			true,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf(`%v %v`, test.typ.String(), test.input), func(t *testing.T) {
			output, err := test.typ.ParseValue(&test.input)
			if !test.expectedErr {
				require.NoError(t, err)
				assert.Equal(t, test.output, output)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestDecimalRoundTrip(t *testing.T) {
	tests := []struct {
		typ         *decimalType
		input       string
		output      string
		expectedErr bool
	}{
		{
			generateDecimalType(t, 1, 0),
			"0",
			"0",
			false,
		},
		{
			generateDecimalType(t, 4, 1),
			"0",
			"0.0",
			false,
		},
		{
			generateDecimalType(t, 9, 4),
			"0",
			"0.0000",
			false,
		},
		{
			generateDecimalType(t, 26, 0),
			"0",
			"0",
			false,
		},
		{
			generateDecimalType(t, 48, 22),
			"0",
			"0.0000000000000000000000",
			false,
		},
		{
			generateDecimalType(t, 65, 30),
			"0",
			"0.000000000000000000000000000000",
			false,
		},
		{
			generateDecimalType(t, 1, 0),
			"-1.5",
			"-2",
			false,
		},
		{
			generateDecimalType(t, 4, 1),
			"-1.5",
			"-1.5",
			false,
		},
		{
			generateDecimalType(t, 9, 4),
			"-1.5",
			"-1.5000",
			false,
		},
		{
			generateDecimalType(t, 26, 0),
			"-1.5",
			"-2",
			false,
		},
		{
			generateDecimalType(t, 48, 22),
			"-1.5",
			"-1.5000000000000000000000",
			false,
		},
		{
			generateDecimalType(t, 65, 30),
			"-1.5",
			"-1.500000000000000000000000000000",
			false,
		},
		{
			generateDecimalType(t, 1, 0),
			"9351580",
			"",
			true,
		},
		{
			generateDecimalType(t, 4, 1),
			"9351580",
			"",
			true,
		},
		{
			generateDecimalType(t, 9, 4),
			"9351580",
			"",
			true,
		},
		{
			generateDecimalType(t, 26, 0),
			"9351580",
			"9351580",
			false,
		},
		{
			generateDecimalType(t, 48, 22),
			"9351580",
			"9351580.0000000000000000000000",
			false,
		},
		{
			generateDecimalType(t, 65, 30),
			"9351580",
			"9351580.000000000000000000000000000000",
			false,
		},
		{
			generateDecimalType(t, 1, 0),
			"-1076416.875",
			"",
			true,
		},
		{
			generateDecimalType(t, 4, 1),
			"-1076416.875",
			"",
			true,
		},
		{
			generateDecimalType(t, 9, 4),
			"-1076416.875",
			"",
			true,
		},
		{
			generateDecimalType(t, 26, 0),
			"-1076416.875",
			"-1076417",
			false,
		},
		{
			generateDecimalType(t, 48, 22),
			"-1076416.875",
			"-1076416.8750000000000000000000",
			false,
		},
		{
			generateDecimalType(t, 65, 30),
			"-1076416.875",
			"-1076416.875000000000000000000000000000",
			false,
		},
		{
			generateDecimalType(t, 1, 0),
			"198728394234798423466321.27349757",
			"",
			true,
		},
		{
			generateDecimalType(t, 4, 1),
			"198728394234798423466321.27349757",
			"",
			true,
		},
		{
			generateDecimalType(t, 9, 4),
			"198728394234798423466321.27349757",
			"",
			true,
		},
		{
			generateDecimalType(t, 26, 0),
			"198728394234798423466321.27349757",
			"198728394234798423466321",
			false,
		},
		{
			generateDecimalType(t, 48, 22),
			"198728394234798423466321.27349757",
			"198728394234798423466321.2734975700000000000000",
			false,
		},
		{
			generateDecimalType(t, 65, 30),
			"198728394234798423466321.27349757",
			"198728394234798423466321.273497570000000000000000000000",
			false,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf(`%v %v %v`, test.typ.String(), test.input, test.output), func(t *testing.T) {
			parsed, err := test.typ.ConvertValueToNomsValue(test.input)
			if !test.expectedErr {
				require.NoError(t, err)
				output, err := test.typ.ConvertNomsValueToValue(parsed)
				require.NoError(t, err)
				assert.Equal(t, test.output, output)
				parsed2, err := test.typ.ParseValue(&test.input)
				require.NoError(t, err)
				assert.Equal(t, parsed, parsed2)
				output2, err := test.typ.FormatValue(parsed2)
				require.NoError(t, err)
				assert.Equal(t, test.output, *output2)
			} else {
				assert.Error(t, err)
				_, err = test.typ.ParseValue(&test.input)
				assert.Error(t, err)
			}
		})
	}
}
