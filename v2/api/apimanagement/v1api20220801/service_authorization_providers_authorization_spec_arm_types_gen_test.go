// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220801

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_Service_AuthorizationProviders_Authorization_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Service_AuthorizationProviders_Authorization_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForService_AuthorizationProviders_Authorization_Spec_ARM, Service_AuthorizationProviders_Authorization_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForService_AuthorizationProviders_Authorization_Spec_ARM runs a test to see if a specific instance of Service_AuthorizationProviders_Authorization_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForService_AuthorizationProviders_Authorization_Spec_ARM(subject Service_AuthorizationProviders_Authorization_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Service_AuthorizationProviders_Authorization_Spec_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Service_AuthorizationProviders_Authorization_Spec_ARM instances for property testing - lazily
// instantiated by Service_AuthorizationProviders_Authorization_Spec_ARMGenerator()
var service_AuthorizationProviders_Authorization_Spec_ARMGenerator gopter.Gen

// Service_AuthorizationProviders_Authorization_Spec_ARMGenerator returns a generator of Service_AuthorizationProviders_Authorization_Spec_ARM instances for property testing.
// We first initialize service_AuthorizationProviders_Authorization_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Service_AuthorizationProviders_Authorization_Spec_ARMGenerator() gopter.Gen {
	if service_AuthorizationProviders_Authorization_Spec_ARMGenerator != nil {
		return service_AuthorizationProviders_Authorization_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForService_AuthorizationProviders_Authorization_Spec_ARM(generators)
	service_AuthorizationProviders_Authorization_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Service_AuthorizationProviders_Authorization_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForService_AuthorizationProviders_Authorization_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForService_AuthorizationProviders_Authorization_Spec_ARM(generators)
	service_AuthorizationProviders_Authorization_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Service_AuthorizationProviders_Authorization_Spec_ARM{}), generators)

	return service_AuthorizationProviders_Authorization_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForService_AuthorizationProviders_Authorization_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForService_AuthorizationProviders_Authorization_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForService_AuthorizationProviders_Authorization_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForService_AuthorizationProviders_Authorization_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(AuthorizationContractProperties_ARMGenerator())
}

func Test_AuthorizationContractProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AuthorizationContractProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAuthorizationContractProperties_ARM, AuthorizationContractProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAuthorizationContractProperties_ARM runs a test to see if a specific instance of AuthorizationContractProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAuthorizationContractProperties_ARM(subject AuthorizationContractProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AuthorizationContractProperties_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of AuthorizationContractProperties_ARM instances for property testing - lazily instantiated by
// AuthorizationContractProperties_ARMGenerator()
var authorizationContractProperties_ARMGenerator gopter.Gen

// AuthorizationContractProperties_ARMGenerator returns a generator of AuthorizationContractProperties_ARM instances for property testing.
// We first initialize authorizationContractProperties_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AuthorizationContractProperties_ARMGenerator() gopter.Gen {
	if authorizationContractProperties_ARMGenerator != nil {
		return authorizationContractProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAuthorizationContractProperties_ARM(generators)
	authorizationContractProperties_ARMGenerator = gen.Struct(reflect.TypeOf(AuthorizationContractProperties_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAuthorizationContractProperties_ARM(generators)
	AddRelatedPropertyGeneratorsForAuthorizationContractProperties_ARM(generators)
	authorizationContractProperties_ARMGenerator = gen.Struct(reflect.TypeOf(AuthorizationContractProperties_ARM{}), generators)

	return authorizationContractProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForAuthorizationContractProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAuthorizationContractProperties_ARM(gens map[string]gopter.Gen) {
	gens["AuthorizationType"] = gen.PtrOf(gen.OneConstOf(AuthorizationContractProperties_AuthorizationType_OAuth2))
	gens["Oauth2GrantType"] = gen.PtrOf(gen.OneConstOf(AuthorizationContractProperties_Oauth2GrantType_AuthorizationCode, AuthorizationContractProperties_Oauth2GrantType_ClientCredentials))
	gens["Parameters"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Status"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForAuthorizationContractProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAuthorizationContractProperties_ARM(gens map[string]gopter.Gen) {
	gens["Error"] = gen.PtrOf(AuthorizationError_ARMGenerator())
}

func Test_AuthorizationError_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AuthorizationError_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAuthorizationError_ARM, AuthorizationError_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAuthorizationError_ARM runs a test to see if a specific instance of AuthorizationError_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAuthorizationError_ARM(subject AuthorizationError_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AuthorizationError_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of AuthorizationError_ARM instances for property testing - lazily instantiated by
// AuthorizationError_ARMGenerator()
var authorizationError_ARMGenerator gopter.Gen

// AuthorizationError_ARMGenerator returns a generator of AuthorizationError_ARM instances for property testing.
func AuthorizationError_ARMGenerator() gopter.Gen {
	if authorizationError_ARMGenerator != nil {
		return authorizationError_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAuthorizationError_ARM(generators)
	authorizationError_ARMGenerator = gen.Struct(reflect.TypeOf(AuthorizationError_ARM{}), generators)

	return authorizationError_ARMGenerator
}

// AddIndependentPropertyGeneratorsForAuthorizationError_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAuthorizationError_ARM(gens map[string]gopter.Gen) {
	gens["Code"] = gen.PtrOf(gen.AlphaString())
	gens["Message"] = gen.PtrOf(gen.AlphaString())
}
