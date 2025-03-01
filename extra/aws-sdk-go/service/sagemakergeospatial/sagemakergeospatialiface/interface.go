// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package sagemakergeospatialiface provides an interface to enable mocking the Amazon SageMaker geospatial capabilities service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package sagemakergeospatialiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sagemakergeospatial"
)

// SageMakerGeospatialAPI provides an interface to enable mocking the
// sagemakergeospatial.SageMakerGeospatial service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon SageMaker geospatial capabilities.
//    func myFunc(svc sagemakergeospatialiface.SageMakerGeospatialAPI) bool {
//        // Make svc.DeleteEarthObservationJob request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := sagemakergeospatial.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockSageMakerGeospatialClient struct {
//        sagemakergeospatialiface.SageMakerGeospatialAPI
//    }
//    func (m *mockSageMakerGeospatialClient) DeleteEarthObservationJob(input *sagemakergeospatial.DeleteEarthObservationJobInput) (*sagemakergeospatial.DeleteEarthObservationJobOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockSageMakerGeospatialClient{}
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
type SageMakerGeospatialAPI interface {
	DeleteEarthObservationJob(*sagemakergeospatial.DeleteEarthObservationJobInput) (*sagemakergeospatial.DeleteEarthObservationJobOutput, error)
	DeleteEarthObservationJobWithContext(aws.Context, *sagemakergeospatial.DeleteEarthObservationJobInput, ...request.Option) (*sagemakergeospatial.DeleteEarthObservationJobOutput, error)
	DeleteEarthObservationJobRequest(*sagemakergeospatial.DeleteEarthObservationJobInput) (*request.Request, *sagemakergeospatial.DeleteEarthObservationJobOutput)

	DeleteVectorEnrichmentJob(*sagemakergeospatial.DeleteVectorEnrichmentJobInput) (*sagemakergeospatial.DeleteVectorEnrichmentJobOutput, error)
	DeleteVectorEnrichmentJobWithContext(aws.Context, *sagemakergeospatial.DeleteVectorEnrichmentJobInput, ...request.Option) (*sagemakergeospatial.DeleteVectorEnrichmentJobOutput, error)
	DeleteVectorEnrichmentJobRequest(*sagemakergeospatial.DeleteVectorEnrichmentJobInput) (*request.Request, *sagemakergeospatial.DeleteVectorEnrichmentJobOutput)

	ExportEarthObservationJob(*sagemakergeospatial.ExportEarthObservationJobInput) (*sagemakergeospatial.ExportEarthObservationJobOutput, error)
	ExportEarthObservationJobWithContext(aws.Context, *sagemakergeospatial.ExportEarthObservationJobInput, ...request.Option) (*sagemakergeospatial.ExportEarthObservationJobOutput, error)
	ExportEarthObservationJobRequest(*sagemakergeospatial.ExportEarthObservationJobInput) (*request.Request, *sagemakergeospatial.ExportEarthObservationJobOutput)

	ExportVectorEnrichmentJob(*sagemakergeospatial.ExportVectorEnrichmentJobInput) (*sagemakergeospatial.ExportVectorEnrichmentJobOutput, error)
	ExportVectorEnrichmentJobWithContext(aws.Context, *sagemakergeospatial.ExportVectorEnrichmentJobInput, ...request.Option) (*sagemakergeospatial.ExportVectorEnrichmentJobOutput, error)
	ExportVectorEnrichmentJobRequest(*sagemakergeospatial.ExportVectorEnrichmentJobInput) (*request.Request, *sagemakergeospatial.ExportVectorEnrichmentJobOutput)

	GetEarthObservationJob(*sagemakergeospatial.GetEarthObservationJobInput) (*sagemakergeospatial.GetEarthObservationJobOutput, error)
	GetEarthObservationJobWithContext(aws.Context, *sagemakergeospatial.GetEarthObservationJobInput, ...request.Option) (*sagemakergeospatial.GetEarthObservationJobOutput, error)
	GetEarthObservationJobRequest(*sagemakergeospatial.GetEarthObservationJobInput) (*request.Request, *sagemakergeospatial.GetEarthObservationJobOutput)

	GetRasterDataCollection(*sagemakergeospatial.GetRasterDataCollectionInput) (*sagemakergeospatial.GetRasterDataCollectionOutput, error)
	GetRasterDataCollectionWithContext(aws.Context, *sagemakergeospatial.GetRasterDataCollectionInput, ...request.Option) (*sagemakergeospatial.GetRasterDataCollectionOutput, error)
	GetRasterDataCollectionRequest(*sagemakergeospatial.GetRasterDataCollectionInput) (*request.Request, *sagemakergeospatial.GetRasterDataCollectionOutput)

	GetTile(*sagemakergeospatial.GetTileInput) (*sagemakergeospatial.GetTileOutput, error)
	GetTileWithContext(aws.Context, *sagemakergeospatial.GetTileInput, ...request.Option) (*sagemakergeospatial.GetTileOutput, error)
	GetTileRequest(*sagemakergeospatial.GetTileInput) (*request.Request, *sagemakergeospatial.GetTileOutput)

	GetVectorEnrichmentJob(*sagemakergeospatial.GetVectorEnrichmentJobInput) (*sagemakergeospatial.GetVectorEnrichmentJobOutput, error)
	GetVectorEnrichmentJobWithContext(aws.Context, *sagemakergeospatial.GetVectorEnrichmentJobInput, ...request.Option) (*sagemakergeospatial.GetVectorEnrichmentJobOutput, error)
	GetVectorEnrichmentJobRequest(*sagemakergeospatial.GetVectorEnrichmentJobInput) (*request.Request, *sagemakergeospatial.GetVectorEnrichmentJobOutput)

	ListEarthObservationJobs(*sagemakergeospatial.ListEarthObservationJobsInput) (*sagemakergeospatial.ListEarthObservationJobsOutput, error)
	ListEarthObservationJobsWithContext(aws.Context, *sagemakergeospatial.ListEarthObservationJobsInput, ...request.Option) (*sagemakergeospatial.ListEarthObservationJobsOutput, error)
	ListEarthObservationJobsRequest(*sagemakergeospatial.ListEarthObservationJobsInput) (*request.Request, *sagemakergeospatial.ListEarthObservationJobsOutput)

	ListEarthObservationJobsPages(*sagemakergeospatial.ListEarthObservationJobsInput, func(*sagemakergeospatial.ListEarthObservationJobsOutput, bool) bool) error
	ListEarthObservationJobsPagesWithContext(aws.Context, *sagemakergeospatial.ListEarthObservationJobsInput, func(*sagemakergeospatial.ListEarthObservationJobsOutput, bool) bool, ...request.Option) error

	ListRasterDataCollections(*sagemakergeospatial.ListRasterDataCollectionsInput) (*sagemakergeospatial.ListRasterDataCollectionsOutput, error)
	ListRasterDataCollectionsWithContext(aws.Context, *sagemakergeospatial.ListRasterDataCollectionsInput, ...request.Option) (*sagemakergeospatial.ListRasterDataCollectionsOutput, error)
	ListRasterDataCollectionsRequest(*sagemakergeospatial.ListRasterDataCollectionsInput) (*request.Request, *sagemakergeospatial.ListRasterDataCollectionsOutput)

	ListRasterDataCollectionsPages(*sagemakergeospatial.ListRasterDataCollectionsInput, func(*sagemakergeospatial.ListRasterDataCollectionsOutput, bool) bool) error
	ListRasterDataCollectionsPagesWithContext(aws.Context, *sagemakergeospatial.ListRasterDataCollectionsInput, func(*sagemakergeospatial.ListRasterDataCollectionsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*sagemakergeospatial.ListTagsForResourceInput) (*sagemakergeospatial.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *sagemakergeospatial.ListTagsForResourceInput, ...request.Option) (*sagemakergeospatial.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*sagemakergeospatial.ListTagsForResourceInput) (*request.Request, *sagemakergeospatial.ListTagsForResourceOutput)

	ListVectorEnrichmentJobs(*sagemakergeospatial.ListVectorEnrichmentJobsInput) (*sagemakergeospatial.ListVectorEnrichmentJobsOutput, error)
	ListVectorEnrichmentJobsWithContext(aws.Context, *sagemakergeospatial.ListVectorEnrichmentJobsInput, ...request.Option) (*sagemakergeospatial.ListVectorEnrichmentJobsOutput, error)
	ListVectorEnrichmentJobsRequest(*sagemakergeospatial.ListVectorEnrichmentJobsInput) (*request.Request, *sagemakergeospatial.ListVectorEnrichmentJobsOutput)

	ListVectorEnrichmentJobsPages(*sagemakergeospatial.ListVectorEnrichmentJobsInput, func(*sagemakergeospatial.ListVectorEnrichmentJobsOutput, bool) bool) error
	ListVectorEnrichmentJobsPagesWithContext(aws.Context, *sagemakergeospatial.ListVectorEnrichmentJobsInput, func(*sagemakergeospatial.ListVectorEnrichmentJobsOutput, bool) bool, ...request.Option) error

	SearchRasterDataCollection(*sagemakergeospatial.SearchRasterDataCollectionInput) (*sagemakergeospatial.SearchRasterDataCollectionOutput, error)
	SearchRasterDataCollectionWithContext(aws.Context, *sagemakergeospatial.SearchRasterDataCollectionInput, ...request.Option) (*sagemakergeospatial.SearchRasterDataCollectionOutput, error)
	SearchRasterDataCollectionRequest(*sagemakergeospatial.SearchRasterDataCollectionInput) (*request.Request, *sagemakergeospatial.SearchRasterDataCollectionOutput)

	SearchRasterDataCollectionPages(*sagemakergeospatial.SearchRasterDataCollectionInput, func(*sagemakergeospatial.SearchRasterDataCollectionOutput, bool) bool) error
	SearchRasterDataCollectionPagesWithContext(aws.Context, *sagemakergeospatial.SearchRasterDataCollectionInput, func(*sagemakergeospatial.SearchRasterDataCollectionOutput, bool) bool, ...request.Option) error

	StartEarthObservationJob(*sagemakergeospatial.StartEarthObservationJobInput) (*sagemakergeospatial.StartEarthObservationJobOutput, error)
	StartEarthObservationJobWithContext(aws.Context, *sagemakergeospatial.StartEarthObservationJobInput, ...request.Option) (*sagemakergeospatial.StartEarthObservationJobOutput, error)
	StartEarthObservationJobRequest(*sagemakergeospatial.StartEarthObservationJobInput) (*request.Request, *sagemakergeospatial.StartEarthObservationJobOutput)

	StartVectorEnrichmentJob(*sagemakergeospatial.StartVectorEnrichmentJobInput) (*sagemakergeospatial.StartVectorEnrichmentJobOutput, error)
	StartVectorEnrichmentJobWithContext(aws.Context, *sagemakergeospatial.StartVectorEnrichmentJobInput, ...request.Option) (*sagemakergeospatial.StartVectorEnrichmentJobOutput, error)
	StartVectorEnrichmentJobRequest(*sagemakergeospatial.StartVectorEnrichmentJobInput) (*request.Request, *sagemakergeospatial.StartVectorEnrichmentJobOutput)

	StopEarthObservationJob(*sagemakergeospatial.StopEarthObservationJobInput) (*sagemakergeospatial.StopEarthObservationJobOutput, error)
	StopEarthObservationJobWithContext(aws.Context, *sagemakergeospatial.StopEarthObservationJobInput, ...request.Option) (*sagemakergeospatial.StopEarthObservationJobOutput, error)
	StopEarthObservationJobRequest(*sagemakergeospatial.StopEarthObservationJobInput) (*request.Request, *sagemakergeospatial.StopEarthObservationJobOutput)

	StopVectorEnrichmentJob(*sagemakergeospatial.StopVectorEnrichmentJobInput) (*sagemakergeospatial.StopVectorEnrichmentJobOutput, error)
	StopVectorEnrichmentJobWithContext(aws.Context, *sagemakergeospatial.StopVectorEnrichmentJobInput, ...request.Option) (*sagemakergeospatial.StopVectorEnrichmentJobOutput, error)
	StopVectorEnrichmentJobRequest(*sagemakergeospatial.StopVectorEnrichmentJobInput) (*request.Request, *sagemakergeospatial.StopVectorEnrichmentJobOutput)

	TagResource(*sagemakergeospatial.TagResourceInput) (*sagemakergeospatial.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *sagemakergeospatial.TagResourceInput, ...request.Option) (*sagemakergeospatial.TagResourceOutput, error)
	TagResourceRequest(*sagemakergeospatial.TagResourceInput) (*request.Request, *sagemakergeospatial.TagResourceOutput)

	UntagResource(*sagemakergeospatial.UntagResourceInput) (*sagemakergeospatial.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *sagemakergeospatial.UntagResourceInput, ...request.Option) (*sagemakergeospatial.UntagResourceOutput, error)
	UntagResourceRequest(*sagemakergeospatial.UntagResourceInput) (*request.Request, *sagemakergeospatial.UntagResourceOutput)
}

var _ SageMakerGeospatialAPI = (*sagemakergeospatial.SageMakerGeospatial)(nil)
