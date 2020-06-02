package fnlib_test

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/puppetlabs/nebula-sdk/pkg/workflow/spec/convert"
	"github.com/puppetlabs/nebula-sdk/pkg/workflow/spec/fn"
	"github.com/puppetlabs/nebula-sdk/pkg/workflow/spec/fnlib"
	"github.com/stretchr/testify/require"
)

func TestConvertMarkdown(t *testing.T) {
	desc, err := fnlib.Library().Descriptor("convertMarkdown")
	require.NoError(t, err)

	tcs := []struct {
		Name        string
		ConvertType convert.ConvertType
		Markdown    string
		Expected    string
	}{
		{
			Name:        "Sample monitor event",
			ConvertType: convert.ConvertTypeJira,
			Markdown:    "%%% @contact [![imageTitle](imageUrl)](imageRedirect) `data{context} > threshold` Detailed description. - - - [[linkTitle1](link1)] · [[linkTitle2](link2)] %%%",
			Expected:    "@contact \n\n[!imageUrl!|imageRedirect] {code}data{context} > threshold{code} Detailed description.\n----\n[[linkTitle1|link1]] · [[linkTitle2|link2]]",
		},
	}

	for _, test := range tcs {
		t.Run(fmt.Sprintf("%s", test.Name), func(t *testing.T) {
			invoker, err := desc.PositionalInvoker([]interface{}{test.ConvertType.String(), test.Markdown})
			require.NoError(t, err)

			r, err := invoker.Invoke(context.Background())
			require.NoError(t, err)

			require.Equal(t, test.Expected, r)

			invoker, err = desc.KeywordInvoker(map[string]interface{}{
				"to":      test.ConvertType.String(),
				"content": test.Markdown,
			})

			r, err = invoker.Invoke(context.Background())
			require.NoError(t, err)

			require.Equal(t, test.Expected, r)
		})
	}
}

func TestConvertMarkdownFunction(t *testing.T) {
	desc, err := fnlib.Library().Descriptor("convertMarkdown")
	require.NoError(t, err)

	tcs := []struct {
		Name                 string
		Invoker              func() (fn.Invoker, error)
		ExpectedInvokeError  error
		ExpectedInvokerError error
	}{
		{
			Name: "keyword invoker with unsupported convert type",
			Invoker: func() (fn.Invoker, error) {
				return desc.KeywordInvoker(map[string]interface{}{
					"to":      "foo",
					"content": "bar",
				})
			},
			ExpectedInvokerError: convert.ErrConvertTypeNotSupported,
		},
		{
			Name: "keyword invoker with invalid to keyword type",
			Invoker: func() (fn.Invoker, error) {
				return desc.KeywordInvoker(map[string]interface{}{
					"to":      false,
					"content": "bar",
				})
			},
			ExpectedInvokerError: &fn.KeywordArgError{
				Arg: "to",
				Cause: &fn.UnexpectedTypeError{
					Wanted: []reflect.Type{
						reflect.TypeOf(""),
					},
					Got: reflect.TypeOf(false),
				},
			},
		},
		{
			Name: "keyword invoker with invalid content keyword type",
			Invoker: func() (fn.Invoker, error) {
				return desc.KeywordInvoker(map[string]interface{}{
					"to":      "jira",
					"content": false,
				})
			},
			ExpectedInvokerError: &fn.KeywordArgError{
				Arg: "content",
				Cause: &fn.UnexpectedTypeError{
					Wanted: []reflect.Type{
						reflect.TypeOf(""),
					},
					Got: reflect.TypeOf(false),
				},
			},
		},
		{
			Name: "keyword invoker with missing to keyword",
			Invoker: func() (fn.Invoker, error) {
				return desc.KeywordInvoker(map[string]interface{}{
					"content": "bar",
				})
			},
			ExpectedInvokeError: &fn.KeywordArgError{
				Arg:   "to",
				Cause: fn.ErrArgNotFound,
			},
		},
		{
			Name: "keyword invoker with missing content",
			Invoker: func() (fn.Invoker, error) {
				return desc.KeywordInvoker(map[string]interface{}{
					"to": "jira",
				})
			},
			ExpectedInvokeError: &fn.KeywordArgError{
				Arg:   "content",
				Cause: fn.ErrArgNotFound,
			},
		},
	}

	for _, test := range tcs {
		t.Run(fmt.Sprintf("%s", test.Name), func(t *testing.T) {
			invoker, err := test.Invoker()
			if test.ExpectedInvokeError != nil {
				require.Equal(t, test.ExpectedInvokeError, err)
			} else {
				require.NoError(t, err)

				_, err = invoker.Invoke(context.Background())
				if test.ExpectedInvokerError != nil {
					require.Equal(t, test.ExpectedInvokerError, err)
				} else {
					require.NoError(t, err)
				}
			}
		})
	}
}
