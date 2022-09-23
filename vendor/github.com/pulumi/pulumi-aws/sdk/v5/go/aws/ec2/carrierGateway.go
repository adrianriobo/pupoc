// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an EC2 Carrier Gateway. See the AWS [documentation](https://docs.aws.amazon.com/vpc/latest/userguide/Carrier_Gateway.html) for more information.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ec2.NewCarrierGateway(ctx, "example", &ec2.CarrierGatewayArgs{
//				VpcId: pulumi.Any(aws_vpc.Example.Id),
//				Tags: pulumi.StringMap{
//					"Name": pulumi.String("example-carrier-gateway"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// `aws_ec2_carrier_gateway` can be imported using the carrier gateway's ID, e.g.,
//
// ```sh
//
//	$ pulumi import aws:ec2/carrierGateway:CarrierGateway example cgw-12345
//
// ```
type CarrierGateway struct {
	pulumi.CustomResourceState

	// The ARN of the carrier gateway.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The AWS account ID of the owner of the carrier gateway.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The ID of the VPC to associate with the carrier gateway.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewCarrierGateway registers a new resource with the given unique name, arguments, and options.
func NewCarrierGateway(ctx *pulumi.Context,
	name string, args *CarrierGatewayArgs, opts ...pulumi.ResourceOption) (*CarrierGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	var resource CarrierGateway
	err := ctx.RegisterResource("aws:ec2/carrierGateway:CarrierGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCarrierGateway gets an existing CarrierGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCarrierGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CarrierGatewayState, opts ...pulumi.ResourceOption) (*CarrierGateway, error) {
	var resource CarrierGateway
	err := ctx.ReadResource("aws:ec2/carrierGateway:CarrierGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CarrierGateway resources.
type carrierGatewayState struct {
	// The ARN of the carrier gateway.
	Arn *string `pulumi:"arn"`
	// The AWS account ID of the owner of the carrier gateway.
	OwnerId *string `pulumi:"ownerId"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The ID of the VPC to associate with the carrier gateway.
	VpcId *string `pulumi:"vpcId"`
}

type CarrierGatewayState struct {
	// The ARN of the carrier gateway.
	Arn pulumi.StringPtrInput
	// The AWS account ID of the owner of the carrier gateway.
	OwnerId pulumi.StringPtrInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// The ID of the VPC to associate with the carrier gateway.
	VpcId pulumi.StringPtrInput
}

func (CarrierGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*carrierGatewayState)(nil)).Elem()
}

type carrierGatewayArgs struct {
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// The ID of the VPC to associate with the carrier gateway.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a CarrierGateway resource.
type CarrierGatewayArgs struct {
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// The ID of the VPC to associate with the carrier gateway.
	VpcId pulumi.StringInput
}

func (CarrierGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*carrierGatewayArgs)(nil)).Elem()
}

type CarrierGatewayInput interface {
	pulumi.Input

	ToCarrierGatewayOutput() CarrierGatewayOutput
	ToCarrierGatewayOutputWithContext(ctx context.Context) CarrierGatewayOutput
}

func (*CarrierGateway) ElementType() reflect.Type {
	return reflect.TypeOf((**CarrierGateway)(nil)).Elem()
}

func (i *CarrierGateway) ToCarrierGatewayOutput() CarrierGatewayOutput {
	return i.ToCarrierGatewayOutputWithContext(context.Background())
}

func (i *CarrierGateway) ToCarrierGatewayOutputWithContext(ctx context.Context) CarrierGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CarrierGatewayOutput)
}

// CarrierGatewayArrayInput is an input type that accepts CarrierGatewayArray and CarrierGatewayArrayOutput values.
// You can construct a concrete instance of `CarrierGatewayArrayInput` via:
//
//	CarrierGatewayArray{ CarrierGatewayArgs{...} }
type CarrierGatewayArrayInput interface {
	pulumi.Input

	ToCarrierGatewayArrayOutput() CarrierGatewayArrayOutput
	ToCarrierGatewayArrayOutputWithContext(context.Context) CarrierGatewayArrayOutput
}

type CarrierGatewayArray []CarrierGatewayInput

func (CarrierGatewayArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CarrierGateway)(nil)).Elem()
}

func (i CarrierGatewayArray) ToCarrierGatewayArrayOutput() CarrierGatewayArrayOutput {
	return i.ToCarrierGatewayArrayOutputWithContext(context.Background())
}

func (i CarrierGatewayArray) ToCarrierGatewayArrayOutputWithContext(ctx context.Context) CarrierGatewayArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CarrierGatewayArrayOutput)
}

// CarrierGatewayMapInput is an input type that accepts CarrierGatewayMap and CarrierGatewayMapOutput values.
// You can construct a concrete instance of `CarrierGatewayMapInput` via:
//
//	CarrierGatewayMap{ "key": CarrierGatewayArgs{...} }
type CarrierGatewayMapInput interface {
	pulumi.Input

	ToCarrierGatewayMapOutput() CarrierGatewayMapOutput
	ToCarrierGatewayMapOutputWithContext(context.Context) CarrierGatewayMapOutput
}

type CarrierGatewayMap map[string]CarrierGatewayInput

func (CarrierGatewayMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CarrierGateway)(nil)).Elem()
}

func (i CarrierGatewayMap) ToCarrierGatewayMapOutput() CarrierGatewayMapOutput {
	return i.ToCarrierGatewayMapOutputWithContext(context.Background())
}

func (i CarrierGatewayMap) ToCarrierGatewayMapOutputWithContext(ctx context.Context) CarrierGatewayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CarrierGatewayMapOutput)
}

type CarrierGatewayOutput struct{ *pulumi.OutputState }

func (CarrierGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CarrierGateway)(nil)).Elem()
}

func (o CarrierGatewayOutput) ToCarrierGatewayOutput() CarrierGatewayOutput {
	return o
}

func (o CarrierGatewayOutput) ToCarrierGatewayOutputWithContext(ctx context.Context) CarrierGatewayOutput {
	return o
}

// The ARN of the carrier gateway.
func (o CarrierGatewayOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *CarrierGateway) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The AWS account ID of the owner of the carrier gateway.
func (o CarrierGatewayOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *CarrierGateway) pulumi.StringOutput { return v.OwnerId }).(pulumi.StringOutput)
}

// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o CarrierGatewayOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CarrierGateway) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o CarrierGatewayOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CarrierGateway) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// The ID of the VPC to associate with the carrier gateway.
func (o CarrierGatewayOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *CarrierGateway) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type CarrierGatewayArrayOutput struct{ *pulumi.OutputState }

func (CarrierGatewayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CarrierGateway)(nil)).Elem()
}

func (o CarrierGatewayArrayOutput) ToCarrierGatewayArrayOutput() CarrierGatewayArrayOutput {
	return o
}

func (o CarrierGatewayArrayOutput) ToCarrierGatewayArrayOutputWithContext(ctx context.Context) CarrierGatewayArrayOutput {
	return o
}

func (o CarrierGatewayArrayOutput) Index(i pulumi.IntInput) CarrierGatewayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CarrierGateway {
		return vs[0].([]*CarrierGateway)[vs[1].(int)]
	}).(CarrierGatewayOutput)
}

type CarrierGatewayMapOutput struct{ *pulumi.OutputState }

func (CarrierGatewayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CarrierGateway)(nil)).Elem()
}

func (o CarrierGatewayMapOutput) ToCarrierGatewayMapOutput() CarrierGatewayMapOutput {
	return o
}

func (o CarrierGatewayMapOutput) ToCarrierGatewayMapOutputWithContext(ctx context.Context) CarrierGatewayMapOutput {
	return o
}

func (o CarrierGatewayMapOutput) MapIndex(k pulumi.StringInput) CarrierGatewayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CarrierGateway {
		return vs[0].(map[string]*CarrierGateway)[vs[1].(string)]
	}).(CarrierGatewayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CarrierGatewayInput)(nil)).Elem(), &CarrierGateway{})
	pulumi.RegisterInputType(reflect.TypeOf((*CarrierGatewayArrayInput)(nil)).Elem(), CarrierGatewayArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CarrierGatewayMapInput)(nil)).Elem(), CarrierGatewayMap{})
	pulumi.RegisterOutputType(CarrierGatewayOutput{})
	pulumi.RegisterOutputType(CarrierGatewayArrayOutput{})
	pulumi.RegisterOutputType(CarrierGatewayMapOutput{})
}
