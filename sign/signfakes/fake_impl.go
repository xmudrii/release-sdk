/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by counterfeiter. DO NOT EDIT.
package signfakes

import (
	"context"
	"sync"

	"github.com/google/go-containerregistry/pkg/crane"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/sigstore/cosign/cmd/cosign/cli/options"
	"github.com/sigstore/cosign/pkg/oci"
	"github.com/sigstore/cosign/pkg/oci/remote"
	"github.com/sirupsen/logrus"
	"sigs.k8s.io/release-sdk/sign"
)

type FakeImpl struct {
	DigestStub        func(string, ...crane.Option) (string, error)
	digestMutex       sync.RWMutex
	digestArgsForCall []struct {
		arg1 string
		arg2 []crane.Option
	}
	digestReturns struct {
		result1 string
		result2 error
	}
	digestReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	EnvDefaultStub        func(string, string) string
	envDefaultMutex       sync.RWMutex
	envDefaultArgsForCall []struct {
		arg1 string
		arg2 string
	}
	envDefaultReturns struct {
		result1 string
	}
	envDefaultReturnsOnCall map[int]struct {
		result1 string
	}
	FileExistsStub        func(string) bool
	fileExistsMutex       sync.RWMutex
	fileExistsArgsForCall []struct {
		arg1 string
	}
	fileExistsReturns struct {
		result1 bool
	}
	fileExistsReturnsOnCall map[int]struct {
		result1 bool
	}
	FindTLogEntriesByPayloadStub        func(context.Context, signa.KeyOpts, string) ([]string, error)
	findTLogEntriesByPayloadMutex       sync.RWMutex
	findTLogEntriesByPayloadArgsForCall []struct {
		arg1 context.Context
		arg2 signa.KeyOpts
		arg3 string
	}
	findTLogEntriesByPayloadReturns struct {
		result1 []string
		result2 error
	}
	findTLogEntriesByPayloadReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	ParseReferenceStub        func(string, ...name.Option) (name.Reference, error)
	parseReferenceMutex       sync.RWMutex
	parseReferenceArgsForCall []struct {
		arg1 string
		arg2 []name.Option
	}
	parseReferenceReturns struct {
		result1 name.Reference
		result2 error
	}
	parseReferenceReturnsOnCall map[int]struct {
		result1 name.Reference
		result2 error
	}
	SetenvStub        func(string, string) error
	setenvMutex       sync.RWMutex
	setenvArgsForCall []struct {
		arg1 string
		arg2 string
	}
	setenvReturns struct {
		result1 error
	}
	setenvReturnsOnCall map[int]struct {
		result1 error
	}
	SignFileInternalStub        func(options.RootOptions, options.KeyOpts, options.RegistryOptions, string, bool, string, string) error
	signFileInternalMutex       sync.RWMutex
	signFileInternalArgsForCall []struct {
		arg1 options.RootOptions
		arg2 options.KeyOpts
		arg3 options.RegistryOptions
		arg4 string
		arg5 bool
		arg6 string
		arg7 string
	}
	signFileInternalReturns struct {
		result1 error
	}
	signFileInternalReturnsOnCall map[int]struct {
		result1 error
	}
	SignImageInternalStub        func(options.RootOptions, options.KeyOpts, options.RegistryOptions, map[string]interface{}, []string, string, bool, string, string, string, bool, bool, string) error
	signImageInternalMutex       sync.RWMutex
	signImageInternalArgsForCall []struct {
		arg1  options.RootOptions
		arg2  options.KeyOpts
		arg3  options.RegistryOptions
		arg4  map[string]interface{}
		arg5  []string
		arg6  string
		arg7  bool
		arg8  string
		arg9  string
		arg10 string
		arg11 bool
		arg12 bool
		arg13 string
	}
	signImageInternalReturns struct {
		result1 error
	}
	signImageInternalReturnsOnCall map[int]struct {
		result1 error
	}
	SignaturesStub        func(oci.SignedEntity) (oci.Signatures, error)
	signaturesMutex       sync.RWMutex
	signaturesArgsForCall []struct {
		arg1 oci.SignedEntity
	}
	signaturesReturns struct {
		result1 oci.Signatures
		result2 error
	}
	signaturesReturnsOnCall map[int]struct {
		result1 oci.Signatures
		result2 error
	}
	SignaturesListStub        func(oci.Signatures) ([]oci.Signature, error)
	signaturesListMutex       sync.RWMutex
	signaturesListArgsForCall []struct {
		arg1 oci.Signatures
	}
	signaturesListReturns struct {
		result1 []oci.Signature
		result2 error
	}
	signaturesListReturnsOnCall map[int]struct {
		result1 []oci.Signature
		result2 error
	}
	SignedEntityStub        func(name.Reference, ...remote.Option) (oci.SignedEntity, error)
	signedEntityMutex       sync.RWMutex
	signedEntityArgsForCall []struct {
		arg1 name.Reference
		arg2 []remote.Option
	}
	signedEntityReturns struct {
		result1 oci.SignedEntity
		result2 error
	}
	signedEntityReturnsOnCall map[int]struct {
		result1 oci.SignedEntity
		result2 error
	}
	TokenFromProvidersStub        func(context.Context, *logrus.Logger) (string, error)
	tokenFromProvidersMutex       sync.RWMutex
	tokenFromProvidersArgsForCall []struct {
		arg1 context.Context
		arg2 *logrus.Logger
	}
	tokenFromProvidersReturns struct {
		result1 string
		result2 error
	}
	tokenFromProvidersReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	VerifyFileInternalStub        func(context.Context, signa.KeyOpts, string, string, string) error
	verifyFileInternalMutex       sync.RWMutex
	verifyFileInternalArgsForCall []struct {
		arg1 context.Context
		arg2 signa.KeyOpts
		arg3 string
		arg4 string
		arg5 string
	}
	verifyFileInternalReturns struct {
		result1 error
	}
	verifyFileInternalReturnsOnCall map[int]struct {
		result1 error
	}
	VerifyImageInternalStub        func(context.Context, string, []string) (*sign.SignedObject, error)
	verifyImageInternalMutex       sync.RWMutex
	verifyImageInternalArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 []string
	}
	verifyImageInternalReturns struct {
		result1 *sign.SignedObject
		result2 error
	}
	verifyImageInternalReturnsOnCall map[int]struct {
		result1 *sign.SignedObject
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeImpl) Digest(arg1 string, arg2 ...crane.Option) (string, error) {
	fake.digestMutex.Lock()
	ret, specificReturn := fake.digestReturnsOnCall[len(fake.digestArgsForCall)]
	fake.digestArgsForCall = append(fake.digestArgsForCall, struct {
		arg1 string
		arg2 []crane.Option
	}{arg1, arg2})
	stub := fake.DigestStub
	fakeReturns := fake.digestReturns
	fake.recordInvocation("Digest", []interface{}{arg1, arg2})
	fake.digestMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) DigestCallCount() int {
	fake.digestMutex.RLock()
	defer fake.digestMutex.RUnlock()
	return len(fake.digestArgsForCall)
}

func (fake *FakeImpl) DigestCalls(stub func(string, ...crane.Option) (string, error)) {
	fake.digestMutex.Lock()
	defer fake.digestMutex.Unlock()
	fake.DigestStub = stub
}

func (fake *FakeImpl) DigestArgsForCall(i int) (string, []crane.Option) {
	fake.digestMutex.RLock()
	defer fake.digestMutex.RUnlock()
	argsForCall := fake.digestArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImpl) DigestReturns(result1 string, result2 error) {
	fake.digestMutex.Lock()
	defer fake.digestMutex.Unlock()
	fake.DigestStub = nil
	fake.digestReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) DigestReturnsOnCall(i int, result1 string, result2 error) {
	fake.digestMutex.Lock()
	defer fake.digestMutex.Unlock()
	fake.DigestStub = nil
	if fake.digestReturnsOnCall == nil {
		fake.digestReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.digestReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) EnvDefault(arg1 string, arg2 string) string {
	fake.envDefaultMutex.Lock()
	ret, specificReturn := fake.envDefaultReturnsOnCall[len(fake.envDefaultArgsForCall)]
	fake.envDefaultArgsForCall = append(fake.envDefaultArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.EnvDefaultStub
	fakeReturns := fake.envDefaultReturns
	fake.recordInvocation("EnvDefault", []interface{}{arg1, arg2})
	fake.envDefaultMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) EnvDefaultCallCount() int {
	fake.envDefaultMutex.RLock()
	defer fake.envDefaultMutex.RUnlock()
	return len(fake.envDefaultArgsForCall)
}

func (fake *FakeImpl) EnvDefaultCalls(stub func(string, string) string) {
	fake.envDefaultMutex.Lock()
	defer fake.envDefaultMutex.Unlock()
	fake.EnvDefaultStub = stub
}

func (fake *FakeImpl) EnvDefaultArgsForCall(i int) (string, string) {
	fake.envDefaultMutex.RLock()
	defer fake.envDefaultMutex.RUnlock()
	argsForCall := fake.envDefaultArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImpl) EnvDefaultReturns(result1 string) {
	fake.envDefaultMutex.Lock()
	defer fake.envDefaultMutex.Unlock()
	fake.EnvDefaultStub = nil
	fake.envDefaultReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeImpl) EnvDefaultReturnsOnCall(i int, result1 string) {
	fake.envDefaultMutex.Lock()
	defer fake.envDefaultMutex.Unlock()
	fake.EnvDefaultStub = nil
	if fake.envDefaultReturnsOnCall == nil {
		fake.envDefaultReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.envDefaultReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeImpl) FileExists(arg1 string) bool {
	fake.fileExistsMutex.Lock()
	ret, specificReturn := fake.fileExistsReturnsOnCall[len(fake.fileExistsArgsForCall)]
	fake.fileExistsArgsForCall = append(fake.fileExistsArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.FileExistsStub
	fakeReturns := fake.fileExistsReturns
	fake.recordInvocation("FileExists", []interface{}{arg1})
	fake.fileExistsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) FileExistsCallCount() int {
	fake.fileExistsMutex.RLock()
	defer fake.fileExistsMutex.RUnlock()
	return len(fake.fileExistsArgsForCall)
}

func (fake *FakeImpl) FileExistsCalls(stub func(string) bool) {
	fake.fileExistsMutex.Lock()
	defer fake.fileExistsMutex.Unlock()
	fake.FileExistsStub = stub
}

func (fake *FakeImpl) FileExistsArgsForCall(i int) string {
	fake.fileExistsMutex.RLock()
	defer fake.fileExistsMutex.RUnlock()
	argsForCall := fake.fileExistsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) FileExistsReturns(result1 bool) {
	fake.fileExistsMutex.Lock()
	defer fake.fileExistsMutex.Unlock()
	fake.FileExistsStub = nil
	fake.fileExistsReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeImpl) FileExistsReturnsOnCall(i int, result1 bool) {
	fake.fileExistsMutex.Lock()
	defer fake.fileExistsMutex.Unlock()
	fake.FileExistsStub = nil
	if fake.fileExistsReturnsOnCall == nil {
		fake.fileExistsReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.fileExistsReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeImpl) FindTLogEntriesByPayload(arg1 context.Context, arg2 signa.KeyOpts, arg3 string) ([]string, error) {
	fake.findTLogEntriesByPayloadMutex.Lock()
	ret, specificReturn := fake.findTLogEntriesByPayloadReturnsOnCall[len(fake.findTLogEntriesByPayloadArgsForCall)]
	fake.findTLogEntriesByPayloadArgsForCall = append(fake.findTLogEntriesByPayloadArgsForCall, struct {
		arg1 context.Context
		arg2 signa.KeyOpts
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.FindTLogEntriesByPayloadStub
	fakeReturns := fake.findTLogEntriesByPayloadReturns
	fake.recordInvocation("FindTLogEntriesByPayload", []interface{}{arg1, arg2, arg3})
	fake.findTLogEntriesByPayloadMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) FindTLogEntriesByPayloadCallCount() int {
	fake.findTLogEntriesByPayloadMutex.RLock()
	defer fake.findTLogEntriesByPayloadMutex.RUnlock()
	return len(fake.findTLogEntriesByPayloadArgsForCall)
}

func (fake *FakeImpl) FindTLogEntriesByPayloadCalls(stub func(context.Context, signa.KeyOpts, string) ([]string, error)) {
	fake.findTLogEntriesByPayloadMutex.Lock()
	defer fake.findTLogEntriesByPayloadMutex.Unlock()
	fake.FindTLogEntriesByPayloadStub = stub
}

func (fake *FakeImpl) FindTLogEntriesByPayloadArgsForCall(i int) (context.Context, signa.KeyOpts, string) {
	fake.findTLogEntriesByPayloadMutex.RLock()
	defer fake.findTLogEntriesByPayloadMutex.RUnlock()
	argsForCall := fake.findTLogEntriesByPayloadArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeImpl) FindTLogEntriesByPayloadReturns(result1 []string, result2 error) {
	fake.findTLogEntriesByPayloadMutex.Lock()
	defer fake.findTLogEntriesByPayloadMutex.Unlock()
	fake.FindTLogEntriesByPayloadStub = nil
	fake.findTLogEntriesByPayloadReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) FindTLogEntriesByPayloadReturnsOnCall(i int, result1 []string, result2 error) {
	fake.findTLogEntriesByPayloadMutex.Lock()
	defer fake.findTLogEntriesByPayloadMutex.Unlock()
	fake.FindTLogEntriesByPayloadStub = nil
	if fake.findTLogEntriesByPayloadReturnsOnCall == nil {
		fake.findTLogEntriesByPayloadReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.findTLogEntriesByPayloadReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) ParseReference(arg1 string, arg2 ...name.Option) (name.Reference, error) {
	fake.parseReferenceMutex.Lock()
	ret, specificReturn := fake.parseReferenceReturnsOnCall[len(fake.parseReferenceArgsForCall)]
	fake.parseReferenceArgsForCall = append(fake.parseReferenceArgsForCall, struct {
		arg1 string
		arg2 []name.Option
	}{arg1, arg2})
	stub := fake.ParseReferenceStub
	fakeReturns := fake.parseReferenceReturns
	fake.recordInvocation("ParseReference", []interface{}{arg1, arg2})
	fake.parseReferenceMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) ParseReferenceCallCount() int {
	fake.parseReferenceMutex.RLock()
	defer fake.parseReferenceMutex.RUnlock()
	return len(fake.parseReferenceArgsForCall)
}

func (fake *FakeImpl) ParseReferenceCalls(stub func(string, ...name.Option) (name.Reference, error)) {
	fake.parseReferenceMutex.Lock()
	defer fake.parseReferenceMutex.Unlock()
	fake.ParseReferenceStub = stub
}

func (fake *FakeImpl) ParseReferenceArgsForCall(i int) (string, []name.Option) {
	fake.parseReferenceMutex.RLock()
	defer fake.parseReferenceMutex.RUnlock()
	argsForCall := fake.parseReferenceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImpl) ParseReferenceReturns(result1 name.Reference, result2 error) {
	fake.parseReferenceMutex.Lock()
	defer fake.parseReferenceMutex.Unlock()
	fake.ParseReferenceStub = nil
	fake.parseReferenceReturns = struct {
		result1 name.Reference
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) ParseReferenceReturnsOnCall(i int, result1 name.Reference, result2 error) {
	fake.parseReferenceMutex.Lock()
	defer fake.parseReferenceMutex.Unlock()
	fake.ParseReferenceStub = nil
	if fake.parseReferenceReturnsOnCall == nil {
		fake.parseReferenceReturnsOnCall = make(map[int]struct {
			result1 name.Reference
			result2 error
		})
	}
	fake.parseReferenceReturnsOnCall[i] = struct {
		result1 name.Reference
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) Setenv(arg1 string, arg2 string) error {
	fake.setenvMutex.Lock()
	ret, specificReturn := fake.setenvReturnsOnCall[len(fake.setenvArgsForCall)]
	fake.setenvArgsForCall = append(fake.setenvArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.SetenvStub
	fakeReturns := fake.setenvReturns
	fake.recordInvocation("Setenv", []interface{}{arg1, arg2})
	fake.setenvMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) SetenvCallCount() int {
	fake.setenvMutex.RLock()
	defer fake.setenvMutex.RUnlock()
	return len(fake.setenvArgsForCall)
}

func (fake *FakeImpl) SetenvCalls(stub func(string, string) error) {
	fake.setenvMutex.Lock()
	defer fake.setenvMutex.Unlock()
	fake.SetenvStub = stub
}

func (fake *FakeImpl) SetenvArgsForCall(i int) (string, string) {
	fake.setenvMutex.RLock()
	defer fake.setenvMutex.RUnlock()
	argsForCall := fake.setenvArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImpl) SetenvReturns(result1 error) {
	fake.setenvMutex.Lock()
	defer fake.setenvMutex.Unlock()
	fake.SetenvStub = nil
	fake.setenvReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) SetenvReturnsOnCall(i int, result1 error) {
	fake.setenvMutex.Lock()
	defer fake.setenvMutex.Unlock()
	fake.SetenvStub = nil
	if fake.setenvReturnsOnCall == nil {
		fake.setenvReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setenvReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) SignFileInternal(arg1 options.RootOptions, arg2 options.KeyOpts, arg3 options.RegistryOptions, arg4 string, arg5 bool, arg6 string, arg7 string) error {
	fake.signFileInternalMutex.Lock()
	ret, specificReturn := fake.signFileInternalReturnsOnCall[len(fake.signFileInternalArgsForCall)]
	fake.signFileInternalArgsForCall = append(fake.signFileInternalArgsForCall, struct {
		arg1 options.RootOptions
		arg2 options.KeyOpts
		arg3 options.RegistryOptions
		arg4 string
		arg5 bool
		arg6 string
		arg7 string
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	stub := fake.SignFileInternalStub
	fakeReturns := fake.signFileInternalReturns
	fake.recordInvocation("SignFileInternal", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	fake.signFileInternalMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) SignFileInternalCallCount() int {
	fake.signFileInternalMutex.RLock()
	defer fake.signFileInternalMutex.RUnlock()
	return len(fake.signFileInternalArgsForCall)
}

func (fake *FakeImpl) SignFileInternalCalls(stub func(options.RootOptions, options.KeyOpts, options.RegistryOptions, string, bool, string, string) error) {
	fake.signFileInternalMutex.Lock()
	defer fake.signFileInternalMutex.Unlock()
	fake.SignFileInternalStub = stub
}

func (fake *FakeImpl) SignFileInternalArgsForCall(i int) (options.RootOptions, options.KeyOpts, options.RegistryOptions, string, bool, string, string) {
	fake.signFileInternalMutex.RLock()
	defer fake.signFileInternalMutex.RUnlock()
	argsForCall := fake.signFileInternalArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6, argsForCall.arg7
}

func (fake *FakeImpl) SignFileInternalReturns(result1 error) {
	fake.signFileInternalMutex.Lock()
	defer fake.signFileInternalMutex.Unlock()
	fake.SignFileInternalStub = nil
	fake.signFileInternalReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) SignFileInternalReturnsOnCall(i int, result1 error) {
	fake.signFileInternalMutex.Lock()
	defer fake.signFileInternalMutex.Unlock()
	fake.SignFileInternalStub = nil
	if fake.signFileInternalReturnsOnCall == nil {
		fake.signFileInternalReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.signFileInternalReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) SignImageInternal(arg1 options.RootOptions, arg2 options.KeyOpts, arg3 options.RegistryOptions, arg4 map[string]interface{}, arg5 []string, arg6 string, arg7 bool, arg8 string, arg9 string, arg10 string, arg11 bool, arg12 bool, arg13 string) error {
	var arg5Copy []string
	if arg5 != nil {
		arg5Copy = make([]string, len(arg5))
		copy(arg5Copy, arg5)
	}
	fake.signImageInternalMutex.Lock()
	ret, specificReturn := fake.signImageInternalReturnsOnCall[len(fake.signImageInternalArgsForCall)]
	fake.signImageInternalArgsForCall = append(fake.signImageInternalArgsForCall, struct {
		arg1  options.RootOptions
		arg2  options.KeyOpts
		arg3  options.RegistryOptions
		arg4  map[string]interface{}
		arg5  []string
		arg6  string
		arg7  bool
		arg8  string
		arg9  string
		arg10 string
		arg11 bool
		arg12 bool
		arg13 string
	}{arg1, arg2, arg3, arg4, arg5Copy, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13})
	stub := fake.SignImageInternalStub
	fakeReturns := fake.signImageInternalReturns
	fake.recordInvocation("SignImageInternal", []interface{}{arg1, arg2, arg3, arg4, arg5Copy, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13})
	fake.signImageInternalMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) SignImageInternalCallCount() int {
	fake.signImageInternalMutex.RLock()
	defer fake.signImageInternalMutex.RUnlock()
	return len(fake.signImageInternalArgsForCall)
}

func (fake *FakeImpl) SignImageInternalCalls(stub func(options.RootOptions, options.KeyOpts, options.RegistryOptions, map[string]interface{}, []string, string, bool, string, string, string, bool, bool, string) error) {
	fake.signImageInternalMutex.Lock()
	defer fake.signImageInternalMutex.Unlock()
	fake.SignImageInternalStub = stub
}

func (fake *FakeImpl) SignImageInternalArgsForCall(i int) (options.RootOptions, options.KeyOpts, options.RegistryOptions, map[string]interface{}, []string, string, bool, string, string, string, bool, bool, string) {
	fake.signImageInternalMutex.RLock()
	defer fake.signImageInternalMutex.RUnlock()
	argsForCall := fake.signImageInternalArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6, argsForCall.arg7, argsForCall.arg8, argsForCall.arg9, argsForCall.arg10, argsForCall.arg11, argsForCall.arg12, argsForCall.arg13
}

func (fake *FakeImpl) SignImageInternalReturns(result1 error) {
	fake.signImageInternalMutex.Lock()
	defer fake.signImageInternalMutex.Unlock()
	fake.SignImageInternalStub = nil
	fake.signImageInternalReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) SignImageInternalReturnsOnCall(i int, result1 error) {
	fake.signImageInternalMutex.Lock()
	defer fake.signImageInternalMutex.Unlock()
	fake.SignImageInternalStub = nil
	if fake.signImageInternalReturnsOnCall == nil {
		fake.signImageInternalReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.signImageInternalReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) Signatures(arg1 oci.SignedEntity) (oci.Signatures, error) {
	fake.signaturesMutex.Lock()
	ret, specificReturn := fake.signaturesReturnsOnCall[len(fake.signaturesArgsForCall)]
	fake.signaturesArgsForCall = append(fake.signaturesArgsForCall, struct {
		arg1 oci.SignedEntity
	}{arg1})
	stub := fake.SignaturesStub
	fakeReturns := fake.signaturesReturns
	fake.recordInvocation("Signatures", []interface{}{arg1})
	fake.signaturesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) SignaturesCallCount() int {
	fake.signaturesMutex.RLock()
	defer fake.signaturesMutex.RUnlock()
	return len(fake.signaturesArgsForCall)
}

func (fake *FakeImpl) SignaturesCalls(stub func(oci.SignedEntity) (oci.Signatures, error)) {
	fake.signaturesMutex.Lock()
	defer fake.signaturesMutex.Unlock()
	fake.SignaturesStub = stub
}

func (fake *FakeImpl) SignaturesArgsForCall(i int) oci.SignedEntity {
	fake.signaturesMutex.RLock()
	defer fake.signaturesMutex.RUnlock()
	argsForCall := fake.signaturesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) SignaturesReturns(result1 oci.Signatures, result2 error) {
	fake.signaturesMutex.Lock()
	defer fake.signaturesMutex.Unlock()
	fake.SignaturesStub = nil
	fake.signaturesReturns = struct {
		result1 oci.Signatures
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) SignaturesReturnsOnCall(i int, result1 oci.Signatures, result2 error) {
	fake.signaturesMutex.Lock()
	defer fake.signaturesMutex.Unlock()
	fake.SignaturesStub = nil
	if fake.signaturesReturnsOnCall == nil {
		fake.signaturesReturnsOnCall = make(map[int]struct {
			result1 oci.Signatures
			result2 error
		})
	}
	fake.signaturesReturnsOnCall[i] = struct {
		result1 oci.Signatures
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) SignaturesList(arg1 oci.Signatures) ([]oci.Signature, error) {
	fake.signaturesListMutex.Lock()
	ret, specificReturn := fake.signaturesListReturnsOnCall[len(fake.signaturesListArgsForCall)]
	fake.signaturesListArgsForCall = append(fake.signaturesListArgsForCall, struct {
		arg1 oci.Signatures
	}{arg1})
	stub := fake.SignaturesListStub
	fakeReturns := fake.signaturesListReturns
	fake.recordInvocation("SignaturesList", []interface{}{arg1})
	fake.signaturesListMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) SignaturesListCallCount() int {
	fake.signaturesListMutex.RLock()
	defer fake.signaturesListMutex.RUnlock()
	return len(fake.signaturesListArgsForCall)
}

func (fake *FakeImpl) SignaturesListCalls(stub func(oci.Signatures) ([]oci.Signature, error)) {
	fake.signaturesListMutex.Lock()
	defer fake.signaturesListMutex.Unlock()
	fake.SignaturesListStub = stub
}

func (fake *FakeImpl) SignaturesListArgsForCall(i int) oci.Signatures {
	fake.signaturesListMutex.RLock()
	defer fake.signaturesListMutex.RUnlock()
	argsForCall := fake.signaturesListArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) SignaturesListReturns(result1 []oci.Signature, result2 error) {
	fake.signaturesListMutex.Lock()
	defer fake.signaturesListMutex.Unlock()
	fake.SignaturesListStub = nil
	fake.signaturesListReturns = struct {
		result1 []oci.Signature
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) SignaturesListReturnsOnCall(i int, result1 []oci.Signature, result2 error) {
	fake.signaturesListMutex.Lock()
	defer fake.signaturesListMutex.Unlock()
	fake.SignaturesListStub = nil
	if fake.signaturesListReturnsOnCall == nil {
		fake.signaturesListReturnsOnCall = make(map[int]struct {
			result1 []oci.Signature
			result2 error
		})
	}
	fake.signaturesListReturnsOnCall[i] = struct {
		result1 []oci.Signature
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) SignedEntity(arg1 name.Reference, arg2 ...remote.Option) (oci.SignedEntity, error) {
	fake.signedEntityMutex.Lock()
	ret, specificReturn := fake.signedEntityReturnsOnCall[len(fake.signedEntityArgsForCall)]
	fake.signedEntityArgsForCall = append(fake.signedEntityArgsForCall, struct {
		arg1 name.Reference
		arg2 []remote.Option
	}{arg1, arg2})
	stub := fake.SignedEntityStub
	fakeReturns := fake.signedEntityReturns
	fake.recordInvocation("SignedEntity", []interface{}{arg1, arg2})
	fake.signedEntityMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) SignedEntityCallCount() int {
	fake.signedEntityMutex.RLock()
	defer fake.signedEntityMutex.RUnlock()
	return len(fake.signedEntityArgsForCall)
}

func (fake *FakeImpl) SignedEntityCalls(stub func(name.Reference, ...remote.Option) (oci.SignedEntity, error)) {
	fake.signedEntityMutex.Lock()
	defer fake.signedEntityMutex.Unlock()
	fake.SignedEntityStub = stub
}

func (fake *FakeImpl) SignedEntityArgsForCall(i int) (name.Reference, []remote.Option) {
	fake.signedEntityMutex.RLock()
	defer fake.signedEntityMutex.RUnlock()
	argsForCall := fake.signedEntityArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImpl) SignedEntityReturns(result1 oci.SignedEntity, result2 error) {
	fake.signedEntityMutex.Lock()
	defer fake.signedEntityMutex.Unlock()
	fake.SignedEntityStub = nil
	fake.signedEntityReturns = struct {
		result1 oci.SignedEntity
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) SignedEntityReturnsOnCall(i int, result1 oci.SignedEntity, result2 error) {
	fake.signedEntityMutex.Lock()
	defer fake.signedEntityMutex.Unlock()
	fake.SignedEntityStub = nil
	if fake.signedEntityReturnsOnCall == nil {
		fake.signedEntityReturnsOnCall = make(map[int]struct {
			result1 oci.SignedEntity
			result2 error
		})
	}
	fake.signedEntityReturnsOnCall[i] = struct {
		result1 oci.SignedEntity
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) TokenFromProviders(arg1 context.Context, arg2 *logrus.Logger) (string, error) {
	fake.tokenFromProvidersMutex.Lock()
	ret, specificReturn := fake.tokenFromProvidersReturnsOnCall[len(fake.tokenFromProvidersArgsForCall)]
	fake.tokenFromProvidersArgsForCall = append(fake.tokenFromProvidersArgsForCall, struct {
		arg1 context.Context
		arg2 *logrus.Logger
	}{arg1, arg2})
	stub := fake.TokenFromProvidersStub
	fakeReturns := fake.tokenFromProvidersReturns
	fake.recordInvocation("TokenFromProviders", []interface{}{arg1, arg2})
	fake.tokenFromProvidersMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) TokenFromProvidersCallCount() int {
	fake.tokenFromProvidersMutex.RLock()
	defer fake.tokenFromProvidersMutex.RUnlock()
	return len(fake.tokenFromProvidersArgsForCall)
}

func (fake *FakeImpl) TokenFromProvidersCalls(stub func(context.Context, *logrus.Logger) (string, error)) {
	fake.tokenFromProvidersMutex.Lock()
	defer fake.tokenFromProvidersMutex.Unlock()
	fake.TokenFromProvidersStub = stub
}

func (fake *FakeImpl) TokenFromProvidersArgsForCall(i int) (context.Context, *logrus.Logger) {
	fake.tokenFromProvidersMutex.RLock()
	defer fake.tokenFromProvidersMutex.RUnlock()
	argsForCall := fake.tokenFromProvidersArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImpl) TokenFromProvidersReturns(result1 string, result2 error) {
	fake.tokenFromProvidersMutex.Lock()
	defer fake.tokenFromProvidersMutex.Unlock()
	fake.TokenFromProvidersStub = nil
	fake.tokenFromProvidersReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) TokenFromProvidersReturnsOnCall(i int, result1 string, result2 error) {
	fake.tokenFromProvidersMutex.Lock()
	defer fake.tokenFromProvidersMutex.Unlock()
	fake.TokenFromProvidersStub = nil
	if fake.tokenFromProvidersReturnsOnCall == nil {
		fake.tokenFromProvidersReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.tokenFromProvidersReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) VerifyFileInternal(arg1 context.Context, arg2 signa.KeyOpts, arg3 string, arg4 string, arg5 string) error {
	fake.verifyFileInternalMutex.Lock()
	ret, specificReturn := fake.verifyFileInternalReturnsOnCall[len(fake.verifyFileInternalArgsForCall)]
	fake.verifyFileInternalArgsForCall = append(fake.verifyFileInternalArgsForCall, struct {
		arg1 context.Context
		arg2 signa.KeyOpts
		arg3 string
		arg4 string
		arg5 string
	}{arg1, arg2, arg3, arg4, arg5})
	stub := fake.VerifyFileInternalStub
	fakeReturns := fake.verifyFileInternalReturns
	fake.recordInvocation("VerifyFileInternal", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.verifyFileInternalMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) VerifyFileInternalCallCount() int {
	fake.verifyFileInternalMutex.RLock()
	defer fake.verifyFileInternalMutex.RUnlock()
	return len(fake.verifyFileInternalArgsForCall)
}

func (fake *FakeImpl) VerifyFileInternalCalls(stub func(context.Context, signa.KeyOpts, string, string, string) error) {
	fake.verifyFileInternalMutex.Lock()
	defer fake.verifyFileInternalMutex.Unlock()
	fake.VerifyFileInternalStub = stub
}

func (fake *FakeImpl) VerifyFileInternalArgsForCall(i int) (context.Context, signa.KeyOpts, string, string, string) {
	fake.verifyFileInternalMutex.RLock()
	defer fake.verifyFileInternalMutex.RUnlock()
	argsForCall := fake.verifyFileInternalArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeImpl) VerifyFileInternalReturns(result1 error) {
	fake.verifyFileInternalMutex.Lock()
	defer fake.verifyFileInternalMutex.Unlock()
	fake.VerifyFileInternalStub = nil
	fake.verifyFileInternalReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) VerifyFileInternalReturnsOnCall(i int, result1 error) {
	fake.verifyFileInternalMutex.Lock()
	defer fake.verifyFileInternalMutex.Unlock()
	fake.VerifyFileInternalStub = nil
	if fake.verifyFileInternalReturnsOnCall == nil {
		fake.verifyFileInternalReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.verifyFileInternalReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) VerifyImageInternal(arg1 context.Context, arg2 string, arg3 []string) (*sign.SignedObject, error) {
	var arg3Copy []string
	if arg3 != nil {
		arg3Copy = make([]string, len(arg3))
		copy(arg3Copy, arg3)
	}
	fake.verifyImageInternalMutex.Lock()
	ret, specificReturn := fake.verifyImageInternalReturnsOnCall[len(fake.verifyImageInternalArgsForCall)]
	fake.verifyImageInternalArgsForCall = append(fake.verifyImageInternalArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 []string
	}{arg1, arg2, arg3Copy})
	stub := fake.VerifyImageInternalStub
	fakeReturns := fake.verifyImageInternalReturns
	fake.recordInvocation("VerifyImageInternal", []interface{}{arg1, arg2, arg3Copy})
	fake.verifyImageInternalMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) VerifyImageInternalCallCount() int {
	fake.verifyImageInternalMutex.RLock()
	defer fake.verifyImageInternalMutex.RUnlock()
	return len(fake.verifyImageInternalArgsForCall)
}

func (fake *FakeImpl) VerifyImageInternalCalls(stub func(context.Context, string, []string) (*sign.SignedObject, error)) {
	fake.verifyImageInternalMutex.Lock()
	defer fake.verifyImageInternalMutex.Unlock()
	fake.VerifyImageInternalStub = stub
}

func (fake *FakeImpl) VerifyImageInternalArgsForCall(i int) (context.Context, string, []string) {
	fake.verifyImageInternalMutex.RLock()
	defer fake.verifyImageInternalMutex.RUnlock()
	argsForCall := fake.verifyImageInternalArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeImpl) VerifyImageInternalReturns(result1 *sign.SignedObject, result2 error) {
	fake.verifyImageInternalMutex.Lock()
	defer fake.verifyImageInternalMutex.Unlock()
	fake.VerifyImageInternalStub = nil
	fake.verifyImageInternalReturns = struct {
		result1 *sign.SignedObject
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) VerifyImageInternalReturnsOnCall(i int, result1 *sign.SignedObject, result2 error) {
	fake.verifyImageInternalMutex.Lock()
	defer fake.verifyImageInternalMutex.Unlock()
	fake.VerifyImageInternalStub = nil
	if fake.verifyImageInternalReturnsOnCall == nil {
		fake.verifyImageInternalReturnsOnCall = make(map[int]struct {
			result1 *sign.SignedObject
			result2 error
		})
	}
	fake.verifyImageInternalReturnsOnCall[i] = struct {
		result1 *sign.SignedObject
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.digestMutex.RLock()
	defer fake.digestMutex.RUnlock()
	fake.envDefaultMutex.RLock()
	defer fake.envDefaultMutex.RUnlock()
	fake.fileExistsMutex.RLock()
	defer fake.fileExistsMutex.RUnlock()
	fake.findTLogEntriesByPayloadMutex.RLock()
	defer fake.findTLogEntriesByPayloadMutex.RUnlock()
	fake.parseReferenceMutex.RLock()
	defer fake.parseReferenceMutex.RUnlock()
	fake.setenvMutex.RLock()
	defer fake.setenvMutex.RUnlock()
	fake.signFileInternalMutex.RLock()
	defer fake.signFileInternalMutex.RUnlock()
	fake.signImageInternalMutex.RLock()
	defer fake.signImageInternalMutex.RUnlock()
	fake.signaturesMutex.RLock()
	defer fake.signaturesMutex.RUnlock()
	fake.signaturesListMutex.RLock()
	defer fake.signaturesListMutex.RUnlock()
	fake.signedEntityMutex.RLock()
	defer fake.signedEntityMutex.RUnlock()
	fake.tokenFromProvidersMutex.RLock()
	defer fake.tokenFromProvidersMutex.RUnlock()
	fake.verifyFileInternalMutex.RLock()
	defer fake.verifyFileInternalMutex.RUnlock()
	fake.verifyImageInternalMutex.RLock()
	defer fake.verifyImageInternalMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeImpl) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
