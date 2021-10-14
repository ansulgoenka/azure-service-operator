// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

import (
	"encoding/json"
	"github.com/Azure/azure-service-operator/v2/api/microsoft.documentdb/v1alpha1api20210515storage"
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

func Test_SqlDatabaseThroughputSetting_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlDatabaseThroughputSetting to SqlDatabaseThroughputSetting via AssignPropertiesToSqlDatabaseThroughputSetting & AssignPropertiesFromSqlDatabaseThroughputSetting returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlDatabaseThroughputSetting, SqlDatabaseThroughputSettingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlDatabaseThroughputSetting tests if a specific instance of SqlDatabaseThroughputSetting can be assigned to v1alpha1api20210515storage and back losslessly
func RunPropertyAssignmentTestForSqlDatabaseThroughputSetting(subject SqlDatabaseThroughputSetting) string {
	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20210515storage.SqlDatabaseThroughputSetting
	err := subject.AssignPropertiesToSqlDatabaseThroughputSetting(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlDatabaseThroughputSetting
	err = actual.AssignPropertiesFromSqlDatabaseThroughputSetting(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_SqlDatabaseThroughputSetting_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabaseThroughputSetting via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabaseThroughputSetting, SqlDatabaseThroughputSettingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabaseThroughputSetting runs a test to see if a specific instance of SqlDatabaseThroughputSetting round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabaseThroughputSetting(subject SqlDatabaseThroughputSetting) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabaseThroughputSetting
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

// Generator of SqlDatabaseThroughputSetting instances for property testing - lazily instantiated by
//SqlDatabaseThroughputSettingGenerator()
var sqlDatabaseThroughputSettingGenerator gopter.Gen

// SqlDatabaseThroughputSettingGenerator returns a generator of SqlDatabaseThroughputSetting instances for property testing.
func SqlDatabaseThroughputSettingGenerator() gopter.Gen {
	if sqlDatabaseThroughputSettingGenerator != nil {
		return sqlDatabaseThroughputSettingGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlDatabaseThroughputSetting(generators)
	sqlDatabaseThroughputSettingGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseThroughputSetting{}), generators)

	return sqlDatabaseThroughputSettingGenerator
}

// AddRelatedPropertyGeneratorsForSqlDatabaseThroughputSetting is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlDatabaseThroughputSetting(gens map[string]gopter.Gen) {
	gens["Spec"] = DatabaseAccountsSqlDatabasesThroughputSettingsSpecGenerator()
	gens["Status"] = ThroughputSettingsGetResultsStatusGenerator()
}

func Test_DatabaseAccountsSqlDatabasesThroughputSettings_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DatabaseAccountsSqlDatabasesThroughputSettings_Spec to DatabaseAccountsSqlDatabasesThroughputSettings_Spec via AssignPropertiesToDatabaseAccountsSqlDatabasesThroughputSettingsSpec & AssignPropertiesFromDatabaseAccountsSqlDatabasesThroughputSettingsSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForDatabaseAccountsSqlDatabasesThroughputSettingsSpec, DatabaseAccountsSqlDatabasesThroughputSettingsSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDatabaseAccountsSqlDatabasesThroughputSettingsSpec tests if a specific instance of DatabaseAccountsSqlDatabasesThroughputSettings_Spec can be assigned to v1alpha1api20210515storage and back losslessly
func RunPropertyAssignmentTestForDatabaseAccountsSqlDatabasesThroughputSettingsSpec(subject DatabaseAccountsSqlDatabasesThroughputSettings_Spec) string {
	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20210515storage.DatabaseAccountsSqlDatabasesThroughputSettings_Spec
	err := subject.AssignPropertiesToDatabaseAccountsSqlDatabasesThroughputSettingsSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DatabaseAccountsSqlDatabasesThroughputSettings_Spec
	err = actual.AssignPropertiesFromDatabaseAccountsSqlDatabasesThroughputSettingsSpec(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_DatabaseAccountsSqlDatabasesThroughputSettings_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsSqlDatabasesThroughputSettings_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsSqlDatabasesThroughputSettingsSpec, DatabaseAccountsSqlDatabasesThroughputSettingsSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsSqlDatabasesThroughputSettingsSpec runs a test to see if a specific instance of DatabaseAccountsSqlDatabasesThroughputSettings_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsSqlDatabasesThroughputSettingsSpec(subject DatabaseAccountsSqlDatabasesThroughputSettings_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsSqlDatabasesThroughputSettings_Spec
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

// Generator of DatabaseAccountsSqlDatabasesThroughputSettings_Spec instances for property testing - lazily instantiated
//by DatabaseAccountsSqlDatabasesThroughputSettingsSpecGenerator()
var databaseAccountsSqlDatabasesThroughputSettingsSpecGenerator gopter.Gen

// DatabaseAccountsSqlDatabasesThroughputSettingsSpecGenerator returns a generator of DatabaseAccountsSqlDatabasesThroughputSettings_Spec instances for property testing.
// We first initialize databaseAccountsSqlDatabasesThroughputSettingsSpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsSqlDatabasesThroughputSettingsSpecGenerator() gopter.Gen {
	if databaseAccountsSqlDatabasesThroughputSettingsSpecGenerator != nil {
		return databaseAccountsSqlDatabasesThroughputSettingsSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesThroughputSettingsSpec(generators)
	databaseAccountsSqlDatabasesThroughputSettingsSpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesThroughputSettings_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesThroughputSettingsSpec(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesThroughputSettingsSpec(generators)
	databaseAccountsSqlDatabasesThroughputSettingsSpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesThroughputSettings_Spec{}), generators)

	return databaseAccountsSqlDatabasesThroughputSettingsSpecGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesThroughputSettingsSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesThroughputSettingsSpec(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesThroughputSettingsSpec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesThroughputSettingsSpec(gens map[string]gopter.Gen) {
	gens["Resource"] = ThroughputSettingsResourceGenerator()
}
