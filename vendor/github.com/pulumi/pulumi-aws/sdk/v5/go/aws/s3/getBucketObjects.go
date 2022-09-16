// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// > **NOTE:** The `s3.getBucketObjects` data source is DEPRECATED and will be removed in a future version! Use `s3.getObjects` instead, where new features and fixes will be added.
//
// > **NOTE on `maxKeys`:** Retrieving very large numbers of keys can adversely affect this provider's performance.
//
// The objects data source returns keys (i.e., file names) and other metadata about objects in an S3 bucket.
func GetBucketObjects(ctx *pulumi.Context, args *GetBucketObjectsArgs, opts ...pulumi.InvokeOption) (*GetBucketObjectsResult, error) {
	var rv GetBucketObjectsResult
	err := ctx.Invoke("aws:s3/getBucketObjects:getBucketObjects", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBucketObjects.
type GetBucketObjectsArgs struct {
	// Lists object keys in this S3 bucket. Alternatively, an [S3 access point](https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html) ARN can be specified
	//
	// Deprecated: Use the aws_s3_objects data source instead
	Bucket string `pulumi:"bucket"`
	// A character used to group keys (Default: none)
	Delimiter *string `pulumi:"delimiter"`
	// Encodes keys using this method (Default: none; besides none, only "url" can be used)
	EncodingType *string `pulumi:"encodingType"`
	// Boolean specifying whether to populate the owner list (Default: false)
	FetchOwner *bool `pulumi:"fetchOwner"`
	// Maximum object keys to return (Default: 1000)
	MaxKeys *int `pulumi:"maxKeys"`
	// Limits results to object keys with this prefix (Default: none)
	Prefix *string `pulumi:"prefix"`
	// Returns key names lexicographically after a specific object key in your bucket (Default: none; S3 lists object keys in UTF-8 character encoding in lexicographical order)
	StartAfter *string `pulumi:"startAfter"`
}

// A collection of values returned by getBucketObjects.
type GetBucketObjectsResult struct {
	// Deprecated: Use the aws_s3_objects data source instead
	Bucket string `pulumi:"bucket"`
	// List of any keys between `prefix` and the next occurrence of `delimiter` (i.e., similar to subdirectories of the `prefix` "directory"); the list is only returned when you specify `delimiter`
	CommonPrefixes []string `pulumi:"commonPrefixes"`
	Delimiter      *string  `pulumi:"delimiter"`
	EncodingType   *string  `pulumi:"encodingType"`
	FetchOwner     *bool    `pulumi:"fetchOwner"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of strings representing object keys
	Keys    []string `pulumi:"keys"`
	MaxKeys *int     `pulumi:"maxKeys"`
	// List of strings representing object owner IDs (see `fetchOwner` above)
	Owners     []string `pulumi:"owners"`
	Prefix     *string  `pulumi:"prefix"`
	StartAfter *string  `pulumi:"startAfter"`
}

func GetBucketObjectsOutput(ctx *pulumi.Context, args GetBucketObjectsOutputArgs, opts ...pulumi.InvokeOption) GetBucketObjectsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetBucketObjectsResult, error) {
			args := v.(GetBucketObjectsArgs)
			r, err := GetBucketObjects(ctx, &args, opts...)
			var s GetBucketObjectsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetBucketObjectsResultOutput)
}

// A collection of arguments for invoking getBucketObjects.
type GetBucketObjectsOutputArgs struct {
	// Lists object keys in this S3 bucket. Alternatively, an [S3 access point](https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html) ARN can be specified
	//
	// Deprecated: Use the aws_s3_objects data source instead
	Bucket pulumi.StringInput `pulumi:"bucket"`
	// A character used to group keys (Default: none)
	Delimiter pulumi.StringPtrInput `pulumi:"delimiter"`
	// Encodes keys using this method (Default: none; besides none, only "url" can be used)
	EncodingType pulumi.StringPtrInput `pulumi:"encodingType"`
	// Boolean specifying whether to populate the owner list (Default: false)
	FetchOwner pulumi.BoolPtrInput `pulumi:"fetchOwner"`
	// Maximum object keys to return (Default: 1000)
	MaxKeys pulumi.IntPtrInput `pulumi:"maxKeys"`
	// Limits results to object keys with this prefix (Default: none)
	Prefix pulumi.StringPtrInput `pulumi:"prefix"`
	// Returns key names lexicographically after a specific object key in your bucket (Default: none; S3 lists object keys in UTF-8 character encoding in lexicographical order)
	StartAfter pulumi.StringPtrInput `pulumi:"startAfter"`
}

func (GetBucketObjectsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBucketObjectsArgs)(nil)).Elem()
}

// A collection of values returned by getBucketObjects.
type GetBucketObjectsResultOutput struct{ *pulumi.OutputState }

func (GetBucketObjectsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBucketObjectsResult)(nil)).Elem()
}

func (o GetBucketObjectsResultOutput) ToGetBucketObjectsResultOutput() GetBucketObjectsResultOutput {
	return o
}

func (o GetBucketObjectsResultOutput) ToGetBucketObjectsResultOutputWithContext(ctx context.Context) GetBucketObjectsResultOutput {
	return o
}

// Deprecated: Use the aws_s3_objects data source instead
func (o GetBucketObjectsResultOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v GetBucketObjectsResult) string { return v.Bucket }).(pulumi.StringOutput)
}

// List of any keys between `prefix` and the next occurrence of `delimiter` (i.e., similar to subdirectories of the `prefix` "directory"); the list is only returned when you specify `delimiter`
func (o GetBucketObjectsResultOutput) CommonPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetBucketObjectsResult) []string { return v.CommonPrefixes }).(pulumi.StringArrayOutput)
}

func (o GetBucketObjectsResultOutput) Delimiter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBucketObjectsResult) *string { return v.Delimiter }).(pulumi.StringPtrOutput)
}

func (o GetBucketObjectsResultOutput) EncodingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBucketObjectsResult) *string { return v.EncodingType }).(pulumi.StringPtrOutput)
}

func (o GetBucketObjectsResultOutput) FetchOwner() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetBucketObjectsResult) *bool { return v.FetchOwner }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetBucketObjectsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetBucketObjectsResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of strings representing object keys
func (o GetBucketObjectsResultOutput) Keys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetBucketObjectsResult) []string { return v.Keys }).(pulumi.StringArrayOutput)
}

func (o GetBucketObjectsResultOutput) MaxKeys() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetBucketObjectsResult) *int { return v.MaxKeys }).(pulumi.IntPtrOutput)
}

// List of strings representing object owner IDs (see `fetchOwner` above)
func (o GetBucketObjectsResultOutput) Owners() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetBucketObjectsResult) []string { return v.Owners }).(pulumi.StringArrayOutput)
}

func (o GetBucketObjectsResultOutput) Prefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBucketObjectsResult) *string { return v.Prefix }).(pulumi.StringPtrOutput)
}

func (o GetBucketObjectsResultOutput) StartAfter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBucketObjectsResult) *string { return v.StartAfter }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBucketObjectsResultOutput{})
}
