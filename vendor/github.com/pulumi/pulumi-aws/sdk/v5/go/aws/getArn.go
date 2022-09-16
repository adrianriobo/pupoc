// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Parses an Amazon Resource Name (ARN) into its constituent parts.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := aws.GetArn(ctx, &GetArnArgs{
//				Arn: "arn:aws:rds:eu-west-1:123456789012:db:mysql-db",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetArn(ctx *pulumi.Context, args *GetArnArgs, opts ...pulumi.InvokeOption) (*GetArnResult, error) {
	var rv GetArnResult
	err := ctx.Invoke("aws:index/getArn:getArn", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getArn.
type GetArnArgs struct {
	// The ARN to parse.
	Arn string `pulumi:"arn"`
}

// A collection of values returned by getArn.
type GetArnResult struct {
	// The [ID](https://docs.aws.amazon.com/general/latest/gr/acct-identifiers.html) of the AWS account that owns the resource, without the hyphens.
	Account string `pulumi:"account"`
	Arn     string `pulumi:"arn"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The partition that the resource is in.
	Partition string `pulumi:"partition"`
	// The region the resource resides in.
	// Note that the ARNs for some resources do not require a region, so this component might be omitted.
	Region string `pulumi:"region"`
	// The content of this part of the ARN varies by service.
	// It often includes an indicator of the type of resource—for example, an IAM user or Amazon RDS database —followed by a slash (/) or a colon (:), followed by the resource name itself.
	Resource string `pulumi:"resource"`
	// The [service namespace](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#genref-aws-service-namespaces) that identifies the AWS product.
	Service string `pulumi:"service"`
}

func GetArnOutput(ctx *pulumi.Context, args GetArnOutputArgs, opts ...pulumi.InvokeOption) GetArnResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetArnResult, error) {
			args := v.(GetArnArgs)
			r, err := GetArn(ctx, &args, opts...)
			var s GetArnResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetArnResultOutput)
}

// A collection of arguments for invoking getArn.
type GetArnOutputArgs struct {
	// The ARN to parse.
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (GetArnOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetArnArgs)(nil)).Elem()
}

// A collection of values returned by getArn.
type GetArnResultOutput struct{ *pulumi.OutputState }

func (GetArnResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetArnResult)(nil)).Elem()
}

func (o GetArnResultOutput) ToGetArnResultOutput() GetArnResultOutput {
	return o
}

func (o GetArnResultOutput) ToGetArnResultOutputWithContext(ctx context.Context) GetArnResultOutput {
	return o
}

// The [ID](https://docs.aws.amazon.com/general/latest/gr/acct-identifiers.html) of the AWS account that owns the resource, without the hyphens.
func (o GetArnResultOutput) Account() pulumi.StringOutput {
	return o.ApplyT(func(v GetArnResult) string { return v.Account }).(pulumi.StringOutput)
}

func (o GetArnResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v GetArnResult) string { return v.Arn }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetArnResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetArnResult) string { return v.Id }).(pulumi.StringOutput)
}

// The partition that the resource is in.
func (o GetArnResultOutput) Partition() pulumi.StringOutput {
	return o.ApplyT(func(v GetArnResult) string { return v.Partition }).(pulumi.StringOutput)
}

// The region the resource resides in.
// Note that the ARNs for some resources do not require a region, so this component might be omitted.
func (o GetArnResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetArnResult) string { return v.Region }).(pulumi.StringOutput)
}

// The content of this part of the ARN varies by service.
// It often includes an indicator of the type of resource—for example, an IAM user or Amazon RDS database —followed by a slash (/) or a colon (:), followed by the resource name itself.
func (o GetArnResultOutput) Resource() pulumi.StringOutput {
	return o.ApplyT(func(v GetArnResult) string { return v.Resource }).(pulumi.StringOutput)
}

// The [service namespace](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#genref-aws-service-namespaces) that identifies the AWS product.
func (o GetArnResultOutput) Service() pulumi.StringOutput {
	return o.ApplyT(func(v GetArnResult) string { return v.Service }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetArnResultOutput{})
}
