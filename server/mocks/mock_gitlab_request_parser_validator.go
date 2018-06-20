// Automatically generated by pegomock. DO NOT EDIT!
// Source: github.com/runatlantis/atlantis/server (interfaces: GitlabRequestParserValidator)

package mocks

import (
	http "net/http"
	"reflect"

	pegomock "github.com/petergtz/pegomock"
)

type MockGitlabRequestParserValidator struct {
	fail func(message string, callerSkip ...int)
}

func NewMockGitlabRequestParserValidator() *MockGitlabRequestParserValidator {
	return &MockGitlabRequestParserValidator{fail: pegomock.GlobalFailHandler}
}

func (mock *MockGitlabRequestParserValidator) ParseAndValidate(r *http.Request, secret []byte) (interface{}, error) {
	params := []pegomock.Param{r, secret}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ParseAndValidate", params, []reflect.Type{reflect.TypeOf((*interface{})(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 interface{}
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(interface{})
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockGitlabRequestParserValidator) VerifyWasCalledOnce() *VerifierGitlabRequestParserValidator {
	return &VerifierGitlabRequestParserValidator{mock, pegomock.Times(1), nil}
}

func (mock *MockGitlabRequestParserValidator) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierGitlabRequestParserValidator {
	return &VerifierGitlabRequestParserValidator{mock, invocationCountMatcher, nil}
}

func (mock *MockGitlabRequestParserValidator) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierGitlabRequestParserValidator {
	return &VerifierGitlabRequestParserValidator{mock, invocationCountMatcher, inOrderContext}
}

type VerifierGitlabRequestParserValidator struct {
	mock                   *MockGitlabRequestParserValidator
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
}

func (verifier *VerifierGitlabRequestParserValidator) ParseAndValidate(r *http.Request, secret []byte) *GitlabRequestParserValidator_ParseAndValidate_OngoingVerification {
	params := []pegomock.Param{r, secret}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ParseAndValidate", params)
	return &GitlabRequestParserValidator_ParseAndValidate_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type GitlabRequestParserValidator_ParseAndValidate_OngoingVerification struct {
	mock              *MockGitlabRequestParserValidator
	methodInvocations []pegomock.MethodInvocation
}

func (c *GitlabRequestParserValidator_ParseAndValidate_OngoingVerification) GetCapturedArguments() (*http.Request, []byte) {
	r, secret := c.GetAllCapturedArguments()
	return r[len(r)-1], secret[len(secret)-1]
}

func (c *GitlabRequestParserValidator_ParseAndValidate_OngoingVerification) GetAllCapturedArguments() (_param0 []*http.Request, _param1 [][]byte) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]*http.Request, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(*http.Request)
		}
		_param1 = make([][]byte, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.([]byte)
		}
	}
	return
}
