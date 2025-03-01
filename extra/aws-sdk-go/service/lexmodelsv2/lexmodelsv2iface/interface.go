// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package lexmodelsv2iface provides an interface to enable mocking the Amazon Lex Model Building V2 service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package lexmodelsv2iface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/lexmodelsv2"
)

// LexModelsV2API provides an interface to enable mocking the
// lexmodelsv2.LexModelsV2 service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Lex Model Building V2.
//    func myFunc(svc lexmodelsv2iface.LexModelsV2API) bool {
//        // Make svc.BatchCreateCustomVocabularyItem request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := lexmodelsv2.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockLexModelsV2Client struct {
//        lexmodelsv2iface.LexModelsV2API
//    }
//    func (m *mockLexModelsV2Client) BatchCreateCustomVocabularyItem(input *lexmodelsv2.BatchCreateCustomVocabularyItemInput) (*lexmodelsv2.BatchCreateCustomVocabularyItemOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockLexModelsV2Client{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type LexModelsV2API interface {
	BatchCreateCustomVocabularyItem(*lexmodelsv2.BatchCreateCustomVocabularyItemInput) (*lexmodelsv2.BatchCreateCustomVocabularyItemOutput, error)
	BatchCreateCustomVocabularyItemWithContext(aws.Context, *lexmodelsv2.BatchCreateCustomVocabularyItemInput, ...request.Option) (*lexmodelsv2.BatchCreateCustomVocabularyItemOutput, error)
	BatchCreateCustomVocabularyItemRequest(*lexmodelsv2.BatchCreateCustomVocabularyItemInput) (*request.Request, *lexmodelsv2.BatchCreateCustomVocabularyItemOutput)

	BatchDeleteCustomVocabularyItem(*lexmodelsv2.BatchDeleteCustomVocabularyItemInput) (*lexmodelsv2.BatchDeleteCustomVocabularyItemOutput, error)
	BatchDeleteCustomVocabularyItemWithContext(aws.Context, *lexmodelsv2.BatchDeleteCustomVocabularyItemInput, ...request.Option) (*lexmodelsv2.BatchDeleteCustomVocabularyItemOutput, error)
	BatchDeleteCustomVocabularyItemRequest(*lexmodelsv2.BatchDeleteCustomVocabularyItemInput) (*request.Request, *lexmodelsv2.BatchDeleteCustomVocabularyItemOutput)

	BatchUpdateCustomVocabularyItem(*lexmodelsv2.BatchUpdateCustomVocabularyItemInput) (*lexmodelsv2.BatchUpdateCustomVocabularyItemOutput, error)
	BatchUpdateCustomVocabularyItemWithContext(aws.Context, *lexmodelsv2.BatchUpdateCustomVocabularyItemInput, ...request.Option) (*lexmodelsv2.BatchUpdateCustomVocabularyItemOutput, error)
	BatchUpdateCustomVocabularyItemRequest(*lexmodelsv2.BatchUpdateCustomVocabularyItemInput) (*request.Request, *lexmodelsv2.BatchUpdateCustomVocabularyItemOutput)

	BuildBotLocale(*lexmodelsv2.BuildBotLocaleInput) (*lexmodelsv2.BuildBotLocaleOutput, error)
	BuildBotLocaleWithContext(aws.Context, *lexmodelsv2.BuildBotLocaleInput, ...request.Option) (*lexmodelsv2.BuildBotLocaleOutput, error)
	BuildBotLocaleRequest(*lexmodelsv2.BuildBotLocaleInput) (*request.Request, *lexmodelsv2.BuildBotLocaleOutput)

	CreateBot(*lexmodelsv2.CreateBotInput) (*lexmodelsv2.CreateBotOutput, error)
	CreateBotWithContext(aws.Context, *lexmodelsv2.CreateBotInput, ...request.Option) (*lexmodelsv2.CreateBotOutput, error)
	CreateBotRequest(*lexmodelsv2.CreateBotInput) (*request.Request, *lexmodelsv2.CreateBotOutput)

	CreateBotAlias(*lexmodelsv2.CreateBotAliasInput) (*lexmodelsv2.CreateBotAliasOutput, error)
	CreateBotAliasWithContext(aws.Context, *lexmodelsv2.CreateBotAliasInput, ...request.Option) (*lexmodelsv2.CreateBotAliasOutput, error)
	CreateBotAliasRequest(*lexmodelsv2.CreateBotAliasInput) (*request.Request, *lexmodelsv2.CreateBotAliasOutput)

	CreateBotLocale(*lexmodelsv2.CreateBotLocaleInput) (*lexmodelsv2.CreateBotLocaleOutput, error)
	CreateBotLocaleWithContext(aws.Context, *lexmodelsv2.CreateBotLocaleInput, ...request.Option) (*lexmodelsv2.CreateBotLocaleOutput, error)
	CreateBotLocaleRequest(*lexmodelsv2.CreateBotLocaleInput) (*request.Request, *lexmodelsv2.CreateBotLocaleOutput)

	CreateBotVersion(*lexmodelsv2.CreateBotVersionInput) (*lexmodelsv2.CreateBotVersionOutput, error)
	CreateBotVersionWithContext(aws.Context, *lexmodelsv2.CreateBotVersionInput, ...request.Option) (*lexmodelsv2.CreateBotVersionOutput, error)
	CreateBotVersionRequest(*lexmodelsv2.CreateBotVersionInput) (*request.Request, *lexmodelsv2.CreateBotVersionOutput)

	CreateExport(*lexmodelsv2.CreateExportInput) (*lexmodelsv2.CreateExportOutput, error)
	CreateExportWithContext(aws.Context, *lexmodelsv2.CreateExportInput, ...request.Option) (*lexmodelsv2.CreateExportOutput, error)
	CreateExportRequest(*lexmodelsv2.CreateExportInput) (*request.Request, *lexmodelsv2.CreateExportOutput)

	CreateIntent(*lexmodelsv2.CreateIntentInput) (*lexmodelsv2.CreateIntentOutput, error)
	CreateIntentWithContext(aws.Context, *lexmodelsv2.CreateIntentInput, ...request.Option) (*lexmodelsv2.CreateIntentOutput, error)
	CreateIntentRequest(*lexmodelsv2.CreateIntentInput) (*request.Request, *lexmodelsv2.CreateIntentOutput)

	CreateResourcePolicy(*lexmodelsv2.CreateResourcePolicyInput) (*lexmodelsv2.CreateResourcePolicyOutput, error)
	CreateResourcePolicyWithContext(aws.Context, *lexmodelsv2.CreateResourcePolicyInput, ...request.Option) (*lexmodelsv2.CreateResourcePolicyOutput, error)
	CreateResourcePolicyRequest(*lexmodelsv2.CreateResourcePolicyInput) (*request.Request, *lexmodelsv2.CreateResourcePolicyOutput)

	CreateResourcePolicyStatement(*lexmodelsv2.CreateResourcePolicyStatementInput) (*lexmodelsv2.CreateResourcePolicyStatementOutput, error)
	CreateResourcePolicyStatementWithContext(aws.Context, *lexmodelsv2.CreateResourcePolicyStatementInput, ...request.Option) (*lexmodelsv2.CreateResourcePolicyStatementOutput, error)
	CreateResourcePolicyStatementRequest(*lexmodelsv2.CreateResourcePolicyStatementInput) (*request.Request, *lexmodelsv2.CreateResourcePolicyStatementOutput)

	CreateSlot(*lexmodelsv2.CreateSlotInput) (*lexmodelsv2.CreateSlotOutput, error)
	CreateSlotWithContext(aws.Context, *lexmodelsv2.CreateSlotInput, ...request.Option) (*lexmodelsv2.CreateSlotOutput, error)
	CreateSlotRequest(*lexmodelsv2.CreateSlotInput) (*request.Request, *lexmodelsv2.CreateSlotOutput)

	CreateSlotType(*lexmodelsv2.CreateSlotTypeInput) (*lexmodelsv2.CreateSlotTypeOutput, error)
	CreateSlotTypeWithContext(aws.Context, *lexmodelsv2.CreateSlotTypeInput, ...request.Option) (*lexmodelsv2.CreateSlotTypeOutput, error)
	CreateSlotTypeRequest(*lexmodelsv2.CreateSlotTypeInput) (*request.Request, *lexmodelsv2.CreateSlotTypeOutput)

	CreateUploadUrl(*lexmodelsv2.CreateUploadUrlInput) (*lexmodelsv2.CreateUploadUrlOutput, error)
	CreateUploadUrlWithContext(aws.Context, *lexmodelsv2.CreateUploadUrlInput, ...request.Option) (*lexmodelsv2.CreateUploadUrlOutput, error)
	CreateUploadUrlRequest(*lexmodelsv2.CreateUploadUrlInput) (*request.Request, *lexmodelsv2.CreateUploadUrlOutput)

	DeleteBot(*lexmodelsv2.DeleteBotInput) (*lexmodelsv2.DeleteBotOutput, error)
	DeleteBotWithContext(aws.Context, *lexmodelsv2.DeleteBotInput, ...request.Option) (*lexmodelsv2.DeleteBotOutput, error)
	DeleteBotRequest(*lexmodelsv2.DeleteBotInput) (*request.Request, *lexmodelsv2.DeleteBotOutput)

	DeleteBotAlias(*lexmodelsv2.DeleteBotAliasInput) (*lexmodelsv2.DeleteBotAliasOutput, error)
	DeleteBotAliasWithContext(aws.Context, *lexmodelsv2.DeleteBotAliasInput, ...request.Option) (*lexmodelsv2.DeleteBotAliasOutput, error)
	DeleteBotAliasRequest(*lexmodelsv2.DeleteBotAliasInput) (*request.Request, *lexmodelsv2.DeleteBotAliasOutput)

	DeleteBotLocale(*lexmodelsv2.DeleteBotLocaleInput) (*lexmodelsv2.DeleteBotLocaleOutput, error)
	DeleteBotLocaleWithContext(aws.Context, *lexmodelsv2.DeleteBotLocaleInput, ...request.Option) (*lexmodelsv2.DeleteBotLocaleOutput, error)
	DeleteBotLocaleRequest(*lexmodelsv2.DeleteBotLocaleInput) (*request.Request, *lexmodelsv2.DeleteBotLocaleOutput)

	DeleteBotVersion(*lexmodelsv2.DeleteBotVersionInput) (*lexmodelsv2.DeleteBotVersionOutput, error)
	DeleteBotVersionWithContext(aws.Context, *lexmodelsv2.DeleteBotVersionInput, ...request.Option) (*lexmodelsv2.DeleteBotVersionOutput, error)
	DeleteBotVersionRequest(*lexmodelsv2.DeleteBotVersionInput) (*request.Request, *lexmodelsv2.DeleteBotVersionOutput)

	DeleteCustomVocabulary(*lexmodelsv2.DeleteCustomVocabularyInput) (*lexmodelsv2.DeleteCustomVocabularyOutput, error)
	DeleteCustomVocabularyWithContext(aws.Context, *lexmodelsv2.DeleteCustomVocabularyInput, ...request.Option) (*lexmodelsv2.DeleteCustomVocabularyOutput, error)
	DeleteCustomVocabularyRequest(*lexmodelsv2.DeleteCustomVocabularyInput) (*request.Request, *lexmodelsv2.DeleteCustomVocabularyOutput)

	DeleteExport(*lexmodelsv2.DeleteExportInput) (*lexmodelsv2.DeleteExportOutput, error)
	DeleteExportWithContext(aws.Context, *lexmodelsv2.DeleteExportInput, ...request.Option) (*lexmodelsv2.DeleteExportOutput, error)
	DeleteExportRequest(*lexmodelsv2.DeleteExportInput) (*request.Request, *lexmodelsv2.DeleteExportOutput)

	DeleteImport(*lexmodelsv2.DeleteImportInput) (*lexmodelsv2.DeleteImportOutput, error)
	DeleteImportWithContext(aws.Context, *lexmodelsv2.DeleteImportInput, ...request.Option) (*lexmodelsv2.DeleteImportOutput, error)
	DeleteImportRequest(*lexmodelsv2.DeleteImportInput) (*request.Request, *lexmodelsv2.DeleteImportOutput)

	DeleteIntent(*lexmodelsv2.DeleteIntentInput) (*lexmodelsv2.DeleteIntentOutput, error)
	DeleteIntentWithContext(aws.Context, *lexmodelsv2.DeleteIntentInput, ...request.Option) (*lexmodelsv2.DeleteIntentOutput, error)
	DeleteIntentRequest(*lexmodelsv2.DeleteIntentInput) (*request.Request, *lexmodelsv2.DeleteIntentOutput)

	DeleteResourcePolicy(*lexmodelsv2.DeleteResourcePolicyInput) (*lexmodelsv2.DeleteResourcePolicyOutput, error)
	DeleteResourcePolicyWithContext(aws.Context, *lexmodelsv2.DeleteResourcePolicyInput, ...request.Option) (*lexmodelsv2.DeleteResourcePolicyOutput, error)
	DeleteResourcePolicyRequest(*lexmodelsv2.DeleteResourcePolicyInput) (*request.Request, *lexmodelsv2.DeleteResourcePolicyOutput)

	DeleteResourcePolicyStatement(*lexmodelsv2.DeleteResourcePolicyStatementInput) (*lexmodelsv2.DeleteResourcePolicyStatementOutput, error)
	DeleteResourcePolicyStatementWithContext(aws.Context, *lexmodelsv2.DeleteResourcePolicyStatementInput, ...request.Option) (*lexmodelsv2.DeleteResourcePolicyStatementOutput, error)
	DeleteResourcePolicyStatementRequest(*lexmodelsv2.DeleteResourcePolicyStatementInput) (*request.Request, *lexmodelsv2.DeleteResourcePolicyStatementOutput)

	DeleteSlot(*lexmodelsv2.DeleteSlotInput) (*lexmodelsv2.DeleteSlotOutput, error)
	DeleteSlotWithContext(aws.Context, *lexmodelsv2.DeleteSlotInput, ...request.Option) (*lexmodelsv2.DeleteSlotOutput, error)
	DeleteSlotRequest(*lexmodelsv2.DeleteSlotInput) (*request.Request, *lexmodelsv2.DeleteSlotOutput)

	DeleteSlotType(*lexmodelsv2.DeleteSlotTypeInput) (*lexmodelsv2.DeleteSlotTypeOutput, error)
	DeleteSlotTypeWithContext(aws.Context, *lexmodelsv2.DeleteSlotTypeInput, ...request.Option) (*lexmodelsv2.DeleteSlotTypeOutput, error)
	DeleteSlotTypeRequest(*lexmodelsv2.DeleteSlotTypeInput) (*request.Request, *lexmodelsv2.DeleteSlotTypeOutput)

	DeleteUtterances(*lexmodelsv2.DeleteUtterancesInput) (*lexmodelsv2.DeleteUtterancesOutput, error)
	DeleteUtterancesWithContext(aws.Context, *lexmodelsv2.DeleteUtterancesInput, ...request.Option) (*lexmodelsv2.DeleteUtterancesOutput, error)
	DeleteUtterancesRequest(*lexmodelsv2.DeleteUtterancesInput) (*request.Request, *lexmodelsv2.DeleteUtterancesOutput)

	DescribeBot(*lexmodelsv2.DescribeBotInput) (*lexmodelsv2.DescribeBotOutput, error)
	DescribeBotWithContext(aws.Context, *lexmodelsv2.DescribeBotInput, ...request.Option) (*lexmodelsv2.DescribeBotOutput, error)
	DescribeBotRequest(*lexmodelsv2.DescribeBotInput) (*request.Request, *lexmodelsv2.DescribeBotOutput)

	DescribeBotAlias(*lexmodelsv2.DescribeBotAliasInput) (*lexmodelsv2.DescribeBotAliasOutput, error)
	DescribeBotAliasWithContext(aws.Context, *lexmodelsv2.DescribeBotAliasInput, ...request.Option) (*lexmodelsv2.DescribeBotAliasOutput, error)
	DescribeBotAliasRequest(*lexmodelsv2.DescribeBotAliasInput) (*request.Request, *lexmodelsv2.DescribeBotAliasOutput)

	DescribeBotLocale(*lexmodelsv2.DescribeBotLocaleInput) (*lexmodelsv2.DescribeBotLocaleOutput, error)
	DescribeBotLocaleWithContext(aws.Context, *lexmodelsv2.DescribeBotLocaleInput, ...request.Option) (*lexmodelsv2.DescribeBotLocaleOutput, error)
	DescribeBotLocaleRequest(*lexmodelsv2.DescribeBotLocaleInput) (*request.Request, *lexmodelsv2.DescribeBotLocaleOutput)

	DescribeBotRecommendation(*lexmodelsv2.DescribeBotRecommendationInput) (*lexmodelsv2.DescribeBotRecommendationOutput, error)
	DescribeBotRecommendationWithContext(aws.Context, *lexmodelsv2.DescribeBotRecommendationInput, ...request.Option) (*lexmodelsv2.DescribeBotRecommendationOutput, error)
	DescribeBotRecommendationRequest(*lexmodelsv2.DescribeBotRecommendationInput) (*request.Request, *lexmodelsv2.DescribeBotRecommendationOutput)

	DescribeBotVersion(*lexmodelsv2.DescribeBotVersionInput) (*lexmodelsv2.DescribeBotVersionOutput, error)
	DescribeBotVersionWithContext(aws.Context, *lexmodelsv2.DescribeBotVersionInput, ...request.Option) (*lexmodelsv2.DescribeBotVersionOutput, error)
	DescribeBotVersionRequest(*lexmodelsv2.DescribeBotVersionInput) (*request.Request, *lexmodelsv2.DescribeBotVersionOutput)

	DescribeCustomVocabularyMetadata(*lexmodelsv2.DescribeCustomVocabularyMetadataInput) (*lexmodelsv2.DescribeCustomVocabularyMetadataOutput, error)
	DescribeCustomVocabularyMetadataWithContext(aws.Context, *lexmodelsv2.DescribeCustomVocabularyMetadataInput, ...request.Option) (*lexmodelsv2.DescribeCustomVocabularyMetadataOutput, error)
	DescribeCustomVocabularyMetadataRequest(*lexmodelsv2.DescribeCustomVocabularyMetadataInput) (*request.Request, *lexmodelsv2.DescribeCustomVocabularyMetadataOutput)

	DescribeExport(*lexmodelsv2.DescribeExportInput) (*lexmodelsv2.DescribeExportOutput, error)
	DescribeExportWithContext(aws.Context, *lexmodelsv2.DescribeExportInput, ...request.Option) (*lexmodelsv2.DescribeExportOutput, error)
	DescribeExportRequest(*lexmodelsv2.DescribeExportInput) (*request.Request, *lexmodelsv2.DescribeExportOutput)

	DescribeImport(*lexmodelsv2.DescribeImportInput) (*lexmodelsv2.DescribeImportOutput, error)
	DescribeImportWithContext(aws.Context, *lexmodelsv2.DescribeImportInput, ...request.Option) (*lexmodelsv2.DescribeImportOutput, error)
	DescribeImportRequest(*lexmodelsv2.DescribeImportInput) (*request.Request, *lexmodelsv2.DescribeImportOutput)

	DescribeIntent(*lexmodelsv2.DescribeIntentInput) (*lexmodelsv2.DescribeIntentOutput, error)
	DescribeIntentWithContext(aws.Context, *lexmodelsv2.DescribeIntentInput, ...request.Option) (*lexmodelsv2.DescribeIntentOutput, error)
	DescribeIntentRequest(*lexmodelsv2.DescribeIntentInput) (*request.Request, *lexmodelsv2.DescribeIntentOutput)

	DescribeResourcePolicy(*lexmodelsv2.DescribeResourcePolicyInput) (*lexmodelsv2.DescribeResourcePolicyOutput, error)
	DescribeResourcePolicyWithContext(aws.Context, *lexmodelsv2.DescribeResourcePolicyInput, ...request.Option) (*lexmodelsv2.DescribeResourcePolicyOutput, error)
	DescribeResourcePolicyRequest(*lexmodelsv2.DescribeResourcePolicyInput) (*request.Request, *lexmodelsv2.DescribeResourcePolicyOutput)

	DescribeSlot(*lexmodelsv2.DescribeSlotInput) (*lexmodelsv2.DescribeSlotOutput, error)
	DescribeSlotWithContext(aws.Context, *lexmodelsv2.DescribeSlotInput, ...request.Option) (*lexmodelsv2.DescribeSlotOutput, error)
	DescribeSlotRequest(*lexmodelsv2.DescribeSlotInput) (*request.Request, *lexmodelsv2.DescribeSlotOutput)

	DescribeSlotType(*lexmodelsv2.DescribeSlotTypeInput) (*lexmodelsv2.DescribeSlotTypeOutput, error)
	DescribeSlotTypeWithContext(aws.Context, *lexmodelsv2.DescribeSlotTypeInput, ...request.Option) (*lexmodelsv2.DescribeSlotTypeOutput, error)
	DescribeSlotTypeRequest(*lexmodelsv2.DescribeSlotTypeInput) (*request.Request, *lexmodelsv2.DescribeSlotTypeOutput)

	ListAggregatedUtterances(*lexmodelsv2.ListAggregatedUtterancesInput) (*lexmodelsv2.ListAggregatedUtterancesOutput, error)
	ListAggregatedUtterancesWithContext(aws.Context, *lexmodelsv2.ListAggregatedUtterancesInput, ...request.Option) (*lexmodelsv2.ListAggregatedUtterancesOutput, error)
	ListAggregatedUtterancesRequest(*lexmodelsv2.ListAggregatedUtterancesInput) (*request.Request, *lexmodelsv2.ListAggregatedUtterancesOutput)

	ListAggregatedUtterancesPages(*lexmodelsv2.ListAggregatedUtterancesInput, func(*lexmodelsv2.ListAggregatedUtterancesOutput, bool) bool) error
	ListAggregatedUtterancesPagesWithContext(aws.Context, *lexmodelsv2.ListAggregatedUtterancesInput, func(*lexmodelsv2.ListAggregatedUtterancesOutput, bool) bool, ...request.Option) error

	ListBotAliases(*lexmodelsv2.ListBotAliasesInput) (*lexmodelsv2.ListBotAliasesOutput, error)
	ListBotAliasesWithContext(aws.Context, *lexmodelsv2.ListBotAliasesInput, ...request.Option) (*lexmodelsv2.ListBotAliasesOutput, error)
	ListBotAliasesRequest(*lexmodelsv2.ListBotAliasesInput) (*request.Request, *lexmodelsv2.ListBotAliasesOutput)

	ListBotAliasesPages(*lexmodelsv2.ListBotAliasesInput, func(*lexmodelsv2.ListBotAliasesOutput, bool) bool) error
	ListBotAliasesPagesWithContext(aws.Context, *lexmodelsv2.ListBotAliasesInput, func(*lexmodelsv2.ListBotAliasesOutput, bool) bool, ...request.Option) error

	ListBotLocales(*lexmodelsv2.ListBotLocalesInput) (*lexmodelsv2.ListBotLocalesOutput, error)
	ListBotLocalesWithContext(aws.Context, *lexmodelsv2.ListBotLocalesInput, ...request.Option) (*lexmodelsv2.ListBotLocalesOutput, error)
	ListBotLocalesRequest(*lexmodelsv2.ListBotLocalesInput) (*request.Request, *lexmodelsv2.ListBotLocalesOutput)

	ListBotLocalesPages(*lexmodelsv2.ListBotLocalesInput, func(*lexmodelsv2.ListBotLocalesOutput, bool) bool) error
	ListBotLocalesPagesWithContext(aws.Context, *lexmodelsv2.ListBotLocalesInput, func(*lexmodelsv2.ListBotLocalesOutput, bool) bool, ...request.Option) error

	ListBotRecommendations(*lexmodelsv2.ListBotRecommendationsInput) (*lexmodelsv2.ListBotRecommendationsOutput, error)
	ListBotRecommendationsWithContext(aws.Context, *lexmodelsv2.ListBotRecommendationsInput, ...request.Option) (*lexmodelsv2.ListBotRecommendationsOutput, error)
	ListBotRecommendationsRequest(*lexmodelsv2.ListBotRecommendationsInput) (*request.Request, *lexmodelsv2.ListBotRecommendationsOutput)

	ListBotRecommendationsPages(*lexmodelsv2.ListBotRecommendationsInput, func(*lexmodelsv2.ListBotRecommendationsOutput, bool) bool) error
	ListBotRecommendationsPagesWithContext(aws.Context, *lexmodelsv2.ListBotRecommendationsInput, func(*lexmodelsv2.ListBotRecommendationsOutput, bool) bool, ...request.Option) error

	ListBotVersions(*lexmodelsv2.ListBotVersionsInput) (*lexmodelsv2.ListBotVersionsOutput, error)
	ListBotVersionsWithContext(aws.Context, *lexmodelsv2.ListBotVersionsInput, ...request.Option) (*lexmodelsv2.ListBotVersionsOutput, error)
	ListBotVersionsRequest(*lexmodelsv2.ListBotVersionsInput) (*request.Request, *lexmodelsv2.ListBotVersionsOutput)

	ListBotVersionsPages(*lexmodelsv2.ListBotVersionsInput, func(*lexmodelsv2.ListBotVersionsOutput, bool) bool) error
	ListBotVersionsPagesWithContext(aws.Context, *lexmodelsv2.ListBotVersionsInput, func(*lexmodelsv2.ListBotVersionsOutput, bool) bool, ...request.Option) error

	ListBots(*lexmodelsv2.ListBotsInput) (*lexmodelsv2.ListBotsOutput, error)
	ListBotsWithContext(aws.Context, *lexmodelsv2.ListBotsInput, ...request.Option) (*lexmodelsv2.ListBotsOutput, error)
	ListBotsRequest(*lexmodelsv2.ListBotsInput) (*request.Request, *lexmodelsv2.ListBotsOutput)

	ListBotsPages(*lexmodelsv2.ListBotsInput, func(*lexmodelsv2.ListBotsOutput, bool) bool) error
	ListBotsPagesWithContext(aws.Context, *lexmodelsv2.ListBotsInput, func(*lexmodelsv2.ListBotsOutput, bool) bool, ...request.Option) error

	ListBuiltInIntents(*lexmodelsv2.ListBuiltInIntentsInput) (*lexmodelsv2.ListBuiltInIntentsOutput, error)
	ListBuiltInIntentsWithContext(aws.Context, *lexmodelsv2.ListBuiltInIntentsInput, ...request.Option) (*lexmodelsv2.ListBuiltInIntentsOutput, error)
	ListBuiltInIntentsRequest(*lexmodelsv2.ListBuiltInIntentsInput) (*request.Request, *lexmodelsv2.ListBuiltInIntentsOutput)

	ListBuiltInIntentsPages(*lexmodelsv2.ListBuiltInIntentsInput, func(*lexmodelsv2.ListBuiltInIntentsOutput, bool) bool) error
	ListBuiltInIntentsPagesWithContext(aws.Context, *lexmodelsv2.ListBuiltInIntentsInput, func(*lexmodelsv2.ListBuiltInIntentsOutput, bool) bool, ...request.Option) error

	ListBuiltInSlotTypes(*lexmodelsv2.ListBuiltInSlotTypesInput) (*lexmodelsv2.ListBuiltInSlotTypesOutput, error)
	ListBuiltInSlotTypesWithContext(aws.Context, *lexmodelsv2.ListBuiltInSlotTypesInput, ...request.Option) (*lexmodelsv2.ListBuiltInSlotTypesOutput, error)
	ListBuiltInSlotTypesRequest(*lexmodelsv2.ListBuiltInSlotTypesInput) (*request.Request, *lexmodelsv2.ListBuiltInSlotTypesOutput)

	ListBuiltInSlotTypesPages(*lexmodelsv2.ListBuiltInSlotTypesInput, func(*lexmodelsv2.ListBuiltInSlotTypesOutput, bool) bool) error
	ListBuiltInSlotTypesPagesWithContext(aws.Context, *lexmodelsv2.ListBuiltInSlotTypesInput, func(*lexmodelsv2.ListBuiltInSlotTypesOutput, bool) bool, ...request.Option) error

	ListCustomVocabularyItems(*lexmodelsv2.ListCustomVocabularyItemsInput) (*lexmodelsv2.ListCustomVocabularyItemsOutput, error)
	ListCustomVocabularyItemsWithContext(aws.Context, *lexmodelsv2.ListCustomVocabularyItemsInput, ...request.Option) (*lexmodelsv2.ListCustomVocabularyItemsOutput, error)
	ListCustomVocabularyItemsRequest(*lexmodelsv2.ListCustomVocabularyItemsInput) (*request.Request, *lexmodelsv2.ListCustomVocabularyItemsOutput)

	ListCustomVocabularyItemsPages(*lexmodelsv2.ListCustomVocabularyItemsInput, func(*lexmodelsv2.ListCustomVocabularyItemsOutput, bool) bool) error
	ListCustomVocabularyItemsPagesWithContext(aws.Context, *lexmodelsv2.ListCustomVocabularyItemsInput, func(*lexmodelsv2.ListCustomVocabularyItemsOutput, bool) bool, ...request.Option) error

	ListExports(*lexmodelsv2.ListExportsInput) (*lexmodelsv2.ListExportsOutput, error)
	ListExportsWithContext(aws.Context, *lexmodelsv2.ListExportsInput, ...request.Option) (*lexmodelsv2.ListExportsOutput, error)
	ListExportsRequest(*lexmodelsv2.ListExportsInput) (*request.Request, *lexmodelsv2.ListExportsOutput)

	ListExportsPages(*lexmodelsv2.ListExportsInput, func(*lexmodelsv2.ListExportsOutput, bool) bool) error
	ListExportsPagesWithContext(aws.Context, *lexmodelsv2.ListExportsInput, func(*lexmodelsv2.ListExportsOutput, bool) bool, ...request.Option) error

	ListImports(*lexmodelsv2.ListImportsInput) (*lexmodelsv2.ListImportsOutput, error)
	ListImportsWithContext(aws.Context, *lexmodelsv2.ListImportsInput, ...request.Option) (*lexmodelsv2.ListImportsOutput, error)
	ListImportsRequest(*lexmodelsv2.ListImportsInput) (*request.Request, *lexmodelsv2.ListImportsOutput)

	ListImportsPages(*lexmodelsv2.ListImportsInput, func(*lexmodelsv2.ListImportsOutput, bool) bool) error
	ListImportsPagesWithContext(aws.Context, *lexmodelsv2.ListImportsInput, func(*lexmodelsv2.ListImportsOutput, bool) bool, ...request.Option) error

	ListIntents(*lexmodelsv2.ListIntentsInput) (*lexmodelsv2.ListIntentsOutput, error)
	ListIntentsWithContext(aws.Context, *lexmodelsv2.ListIntentsInput, ...request.Option) (*lexmodelsv2.ListIntentsOutput, error)
	ListIntentsRequest(*lexmodelsv2.ListIntentsInput) (*request.Request, *lexmodelsv2.ListIntentsOutput)

	ListIntentsPages(*lexmodelsv2.ListIntentsInput, func(*lexmodelsv2.ListIntentsOutput, bool) bool) error
	ListIntentsPagesWithContext(aws.Context, *lexmodelsv2.ListIntentsInput, func(*lexmodelsv2.ListIntentsOutput, bool) bool, ...request.Option) error

	ListRecommendedIntents(*lexmodelsv2.ListRecommendedIntentsInput) (*lexmodelsv2.ListRecommendedIntentsOutput, error)
	ListRecommendedIntentsWithContext(aws.Context, *lexmodelsv2.ListRecommendedIntentsInput, ...request.Option) (*lexmodelsv2.ListRecommendedIntentsOutput, error)
	ListRecommendedIntentsRequest(*lexmodelsv2.ListRecommendedIntentsInput) (*request.Request, *lexmodelsv2.ListRecommendedIntentsOutput)

	ListRecommendedIntentsPages(*lexmodelsv2.ListRecommendedIntentsInput, func(*lexmodelsv2.ListRecommendedIntentsOutput, bool) bool) error
	ListRecommendedIntentsPagesWithContext(aws.Context, *lexmodelsv2.ListRecommendedIntentsInput, func(*lexmodelsv2.ListRecommendedIntentsOutput, bool) bool, ...request.Option) error

	ListSlotTypes(*lexmodelsv2.ListSlotTypesInput) (*lexmodelsv2.ListSlotTypesOutput, error)
	ListSlotTypesWithContext(aws.Context, *lexmodelsv2.ListSlotTypesInput, ...request.Option) (*lexmodelsv2.ListSlotTypesOutput, error)
	ListSlotTypesRequest(*lexmodelsv2.ListSlotTypesInput) (*request.Request, *lexmodelsv2.ListSlotTypesOutput)

	ListSlotTypesPages(*lexmodelsv2.ListSlotTypesInput, func(*lexmodelsv2.ListSlotTypesOutput, bool) bool) error
	ListSlotTypesPagesWithContext(aws.Context, *lexmodelsv2.ListSlotTypesInput, func(*lexmodelsv2.ListSlotTypesOutput, bool) bool, ...request.Option) error

	ListSlots(*lexmodelsv2.ListSlotsInput) (*lexmodelsv2.ListSlotsOutput, error)
	ListSlotsWithContext(aws.Context, *lexmodelsv2.ListSlotsInput, ...request.Option) (*lexmodelsv2.ListSlotsOutput, error)
	ListSlotsRequest(*lexmodelsv2.ListSlotsInput) (*request.Request, *lexmodelsv2.ListSlotsOutput)

	ListSlotsPages(*lexmodelsv2.ListSlotsInput, func(*lexmodelsv2.ListSlotsOutput, bool) bool) error
	ListSlotsPagesWithContext(aws.Context, *lexmodelsv2.ListSlotsInput, func(*lexmodelsv2.ListSlotsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*lexmodelsv2.ListTagsForResourceInput) (*lexmodelsv2.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *lexmodelsv2.ListTagsForResourceInput, ...request.Option) (*lexmodelsv2.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*lexmodelsv2.ListTagsForResourceInput) (*request.Request, *lexmodelsv2.ListTagsForResourceOutput)

	SearchAssociatedTranscripts(*lexmodelsv2.SearchAssociatedTranscriptsInput) (*lexmodelsv2.SearchAssociatedTranscriptsOutput, error)
	SearchAssociatedTranscriptsWithContext(aws.Context, *lexmodelsv2.SearchAssociatedTranscriptsInput, ...request.Option) (*lexmodelsv2.SearchAssociatedTranscriptsOutput, error)
	SearchAssociatedTranscriptsRequest(*lexmodelsv2.SearchAssociatedTranscriptsInput) (*request.Request, *lexmodelsv2.SearchAssociatedTranscriptsOutput)

	StartBotRecommendation(*lexmodelsv2.StartBotRecommendationInput) (*lexmodelsv2.StartBotRecommendationOutput, error)
	StartBotRecommendationWithContext(aws.Context, *lexmodelsv2.StartBotRecommendationInput, ...request.Option) (*lexmodelsv2.StartBotRecommendationOutput, error)
	StartBotRecommendationRequest(*lexmodelsv2.StartBotRecommendationInput) (*request.Request, *lexmodelsv2.StartBotRecommendationOutput)

	StartImport(*lexmodelsv2.StartImportInput) (*lexmodelsv2.StartImportOutput, error)
	StartImportWithContext(aws.Context, *lexmodelsv2.StartImportInput, ...request.Option) (*lexmodelsv2.StartImportOutput, error)
	StartImportRequest(*lexmodelsv2.StartImportInput) (*request.Request, *lexmodelsv2.StartImportOutput)

	StopBotRecommendation(*lexmodelsv2.StopBotRecommendationInput) (*lexmodelsv2.StopBotRecommendationOutput, error)
	StopBotRecommendationWithContext(aws.Context, *lexmodelsv2.StopBotRecommendationInput, ...request.Option) (*lexmodelsv2.StopBotRecommendationOutput, error)
	StopBotRecommendationRequest(*lexmodelsv2.StopBotRecommendationInput) (*request.Request, *lexmodelsv2.StopBotRecommendationOutput)

	TagResource(*lexmodelsv2.TagResourceInput) (*lexmodelsv2.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *lexmodelsv2.TagResourceInput, ...request.Option) (*lexmodelsv2.TagResourceOutput, error)
	TagResourceRequest(*lexmodelsv2.TagResourceInput) (*request.Request, *lexmodelsv2.TagResourceOutput)

	UntagResource(*lexmodelsv2.UntagResourceInput) (*lexmodelsv2.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *lexmodelsv2.UntagResourceInput, ...request.Option) (*lexmodelsv2.UntagResourceOutput, error)
	UntagResourceRequest(*lexmodelsv2.UntagResourceInput) (*request.Request, *lexmodelsv2.UntagResourceOutput)

	UpdateBot(*lexmodelsv2.UpdateBotInput) (*lexmodelsv2.UpdateBotOutput, error)
	UpdateBotWithContext(aws.Context, *lexmodelsv2.UpdateBotInput, ...request.Option) (*lexmodelsv2.UpdateBotOutput, error)
	UpdateBotRequest(*lexmodelsv2.UpdateBotInput) (*request.Request, *lexmodelsv2.UpdateBotOutput)

	UpdateBotAlias(*lexmodelsv2.UpdateBotAliasInput) (*lexmodelsv2.UpdateBotAliasOutput, error)
	UpdateBotAliasWithContext(aws.Context, *lexmodelsv2.UpdateBotAliasInput, ...request.Option) (*lexmodelsv2.UpdateBotAliasOutput, error)
	UpdateBotAliasRequest(*lexmodelsv2.UpdateBotAliasInput) (*request.Request, *lexmodelsv2.UpdateBotAliasOutput)

	UpdateBotLocale(*lexmodelsv2.UpdateBotLocaleInput) (*lexmodelsv2.UpdateBotLocaleOutput, error)
	UpdateBotLocaleWithContext(aws.Context, *lexmodelsv2.UpdateBotLocaleInput, ...request.Option) (*lexmodelsv2.UpdateBotLocaleOutput, error)
	UpdateBotLocaleRequest(*lexmodelsv2.UpdateBotLocaleInput) (*request.Request, *lexmodelsv2.UpdateBotLocaleOutput)

	UpdateBotRecommendation(*lexmodelsv2.UpdateBotRecommendationInput) (*lexmodelsv2.UpdateBotRecommendationOutput, error)
	UpdateBotRecommendationWithContext(aws.Context, *lexmodelsv2.UpdateBotRecommendationInput, ...request.Option) (*lexmodelsv2.UpdateBotRecommendationOutput, error)
	UpdateBotRecommendationRequest(*lexmodelsv2.UpdateBotRecommendationInput) (*request.Request, *lexmodelsv2.UpdateBotRecommendationOutput)

	UpdateExport(*lexmodelsv2.UpdateExportInput) (*lexmodelsv2.UpdateExportOutput, error)
	UpdateExportWithContext(aws.Context, *lexmodelsv2.UpdateExportInput, ...request.Option) (*lexmodelsv2.UpdateExportOutput, error)
	UpdateExportRequest(*lexmodelsv2.UpdateExportInput) (*request.Request, *lexmodelsv2.UpdateExportOutput)

	UpdateIntent(*lexmodelsv2.UpdateIntentInput) (*lexmodelsv2.UpdateIntentOutput, error)
	UpdateIntentWithContext(aws.Context, *lexmodelsv2.UpdateIntentInput, ...request.Option) (*lexmodelsv2.UpdateIntentOutput, error)
	UpdateIntentRequest(*lexmodelsv2.UpdateIntentInput) (*request.Request, *lexmodelsv2.UpdateIntentOutput)

	UpdateResourcePolicy(*lexmodelsv2.UpdateResourcePolicyInput) (*lexmodelsv2.UpdateResourcePolicyOutput, error)
	UpdateResourcePolicyWithContext(aws.Context, *lexmodelsv2.UpdateResourcePolicyInput, ...request.Option) (*lexmodelsv2.UpdateResourcePolicyOutput, error)
	UpdateResourcePolicyRequest(*lexmodelsv2.UpdateResourcePolicyInput) (*request.Request, *lexmodelsv2.UpdateResourcePolicyOutput)

	UpdateSlot(*lexmodelsv2.UpdateSlotInput) (*lexmodelsv2.UpdateSlotOutput, error)
	UpdateSlotWithContext(aws.Context, *lexmodelsv2.UpdateSlotInput, ...request.Option) (*lexmodelsv2.UpdateSlotOutput, error)
	UpdateSlotRequest(*lexmodelsv2.UpdateSlotInput) (*request.Request, *lexmodelsv2.UpdateSlotOutput)

	UpdateSlotType(*lexmodelsv2.UpdateSlotTypeInput) (*lexmodelsv2.UpdateSlotTypeOutput, error)
	UpdateSlotTypeWithContext(aws.Context, *lexmodelsv2.UpdateSlotTypeInput, ...request.Option) (*lexmodelsv2.UpdateSlotTypeOutput, error)
	UpdateSlotTypeRequest(*lexmodelsv2.UpdateSlotTypeInput) (*request.Request, *lexmodelsv2.UpdateSlotTypeOutput)

	WaitUntilBotAliasAvailable(*lexmodelsv2.DescribeBotAliasInput) error
	WaitUntilBotAliasAvailableWithContext(aws.Context, *lexmodelsv2.DescribeBotAliasInput, ...request.WaiterOption) error

	WaitUntilBotAvailable(*lexmodelsv2.DescribeBotInput) error
	WaitUntilBotAvailableWithContext(aws.Context, *lexmodelsv2.DescribeBotInput, ...request.WaiterOption) error

	WaitUntilBotExportCompleted(*lexmodelsv2.DescribeExportInput) error
	WaitUntilBotExportCompletedWithContext(aws.Context, *lexmodelsv2.DescribeExportInput, ...request.WaiterOption) error

	WaitUntilBotImportCompleted(*lexmodelsv2.DescribeImportInput) error
	WaitUntilBotImportCompletedWithContext(aws.Context, *lexmodelsv2.DescribeImportInput, ...request.WaiterOption) error

	WaitUntilBotLocaleBuilt(*lexmodelsv2.DescribeBotLocaleInput) error
	WaitUntilBotLocaleBuiltWithContext(aws.Context, *lexmodelsv2.DescribeBotLocaleInput, ...request.WaiterOption) error

	WaitUntilBotLocaleCreated(*lexmodelsv2.DescribeBotLocaleInput) error
	WaitUntilBotLocaleCreatedWithContext(aws.Context, *lexmodelsv2.DescribeBotLocaleInput, ...request.WaiterOption) error

	WaitUntilBotLocaleExpressTestingAvailable(*lexmodelsv2.DescribeBotLocaleInput) error
	WaitUntilBotLocaleExpressTestingAvailableWithContext(aws.Context, *lexmodelsv2.DescribeBotLocaleInput, ...request.WaiterOption) error

	WaitUntilBotVersionAvailable(*lexmodelsv2.DescribeBotVersionInput) error
	WaitUntilBotVersionAvailableWithContext(aws.Context, *lexmodelsv2.DescribeBotVersionInput, ...request.WaiterOption) error
}

var _ LexModelsV2API = (*lexmodelsv2.LexModelsV2)(nil)
