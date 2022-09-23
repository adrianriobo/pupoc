// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an Traffic mirror filter.\
// Read [limits and considerations](https://docs.aws.amazon.com/vpc/latest/mirroring/traffic-mirroring-considerations.html) for traffic mirroring
//
// ## Example Usage
//
// # To create a basic traffic mirror filter
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
//			_, err := ec2.NewTrafficMirrorFilter(ctx, "foo", &ec2.TrafficMirrorFilterArgs{
//				Description: pulumi.String("traffic mirror filter - example"),
//				NetworkServices: pulumi.StringArray{
//					pulumi.String("amazon-dns"),
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
// Traffic mirror filter can be imported using the `id`, e.g.,
//
// ```sh
//
//	$ pulumi import aws:ec2/trafficMirrorFilter:TrafficMirrorFilter foo tmf-0fbb93ddf38198f64
//
// ```
type TrafficMirrorFilter struct {
	pulumi.CustomResourceState

	// The ARN of the traffic mirror filter.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A description of the filter.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// List of amazon network services that should be mirrored. Valid values: `amazon-dns`.
	NetworkServices pulumi.StringArrayOutput `pulumi:"networkServices"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewTrafficMirrorFilter registers a new resource with the given unique name, arguments, and options.
func NewTrafficMirrorFilter(ctx *pulumi.Context,
	name string, args *TrafficMirrorFilterArgs, opts ...pulumi.ResourceOption) (*TrafficMirrorFilter, error) {
	if args == nil {
		args = &TrafficMirrorFilterArgs{}
	}

	var resource TrafficMirrorFilter
	err := ctx.RegisterResource("aws:ec2/trafficMirrorFilter:TrafficMirrorFilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrafficMirrorFilter gets an existing TrafficMirrorFilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrafficMirrorFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrafficMirrorFilterState, opts ...pulumi.ResourceOption) (*TrafficMirrorFilter, error) {
	var resource TrafficMirrorFilter
	err := ctx.ReadResource("aws:ec2/trafficMirrorFilter:TrafficMirrorFilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TrafficMirrorFilter resources.
type trafficMirrorFilterState struct {
	// The ARN of the traffic mirror filter.
	Arn *string `pulumi:"arn"`
	// A description of the filter.
	Description *string `pulumi:"description"`
	// List of amazon network services that should be mirrored. Valid values: `amazon-dns`.
	NetworkServices []string `pulumi:"networkServices"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type TrafficMirrorFilterState struct {
	// The ARN of the traffic mirror filter.
	Arn pulumi.StringPtrInput
	// A description of the filter.
	Description pulumi.StringPtrInput
	// List of amazon network services that should be mirrored. Valid values: `amazon-dns`.
	NetworkServices pulumi.StringArrayInput
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
}

func (TrafficMirrorFilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficMirrorFilterState)(nil)).Elem()
}

type trafficMirrorFilterArgs struct {
	// A description of the filter.
	Description *string `pulumi:"description"`
	// List of amazon network services that should be mirrored. Valid values: `amazon-dns`.
	NetworkServices []string `pulumi:"networkServices"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a TrafficMirrorFilter resource.
type TrafficMirrorFilterArgs struct {
	// A description of the filter.
	Description pulumi.StringPtrInput
	// List of amazon network services that should be mirrored. Valid values: `amazon-dns`.
	NetworkServices pulumi.StringArrayInput
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (TrafficMirrorFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficMirrorFilterArgs)(nil)).Elem()
}

type TrafficMirrorFilterInput interface {
	pulumi.Input

	ToTrafficMirrorFilterOutput() TrafficMirrorFilterOutput
	ToTrafficMirrorFilterOutputWithContext(ctx context.Context) TrafficMirrorFilterOutput
}

func (*TrafficMirrorFilter) ElementType() reflect.Type {
	return reflect.TypeOf((**TrafficMirrorFilter)(nil)).Elem()
}

func (i *TrafficMirrorFilter) ToTrafficMirrorFilterOutput() TrafficMirrorFilterOutput {
	return i.ToTrafficMirrorFilterOutputWithContext(context.Background())
}

func (i *TrafficMirrorFilter) ToTrafficMirrorFilterOutputWithContext(ctx context.Context) TrafficMirrorFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficMirrorFilterOutput)
}

// TrafficMirrorFilterArrayInput is an input type that accepts TrafficMirrorFilterArray and TrafficMirrorFilterArrayOutput values.
// You can construct a concrete instance of `TrafficMirrorFilterArrayInput` via:
//
//	TrafficMirrorFilterArray{ TrafficMirrorFilterArgs{...} }
type TrafficMirrorFilterArrayInput interface {
	pulumi.Input

	ToTrafficMirrorFilterArrayOutput() TrafficMirrorFilterArrayOutput
	ToTrafficMirrorFilterArrayOutputWithContext(context.Context) TrafficMirrorFilterArrayOutput
}

type TrafficMirrorFilterArray []TrafficMirrorFilterInput

func (TrafficMirrorFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TrafficMirrorFilter)(nil)).Elem()
}

func (i TrafficMirrorFilterArray) ToTrafficMirrorFilterArrayOutput() TrafficMirrorFilterArrayOutput {
	return i.ToTrafficMirrorFilterArrayOutputWithContext(context.Background())
}

func (i TrafficMirrorFilterArray) ToTrafficMirrorFilterArrayOutputWithContext(ctx context.Context) TrafficMirrorFilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficMirrorFilterArrayOutput)
}

// TrafficMirrorFilterMapInput is an input type that accepts TrafficMirrorFilterMap and TrafficMirrorFilterMapOutput values.
// You can construct a concrete instance of `TrafficMirrorFilterMapInput` via:
//
//	TrafficMirrorFilterMap{ "key": TrafficMirrorFilterArgs{...} }
type TrafficMirrorFilterMapInput interface {
	pulumi.Input

	ToTrafficMirrorFilterMapOutput() TrafficMirrorFilterMapOutput
	ToTrafficMirrorFilterMapOutputWithContext(context.Context) TrafficMirrorFilterMapOutput
}

type TrafficMirrorFilterMap map[string]TrafficMirrorFilterInput

func (TrafficMirrorFilterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TrafficMirrorFilter)(nil)).Elem()
}

func (i TrafficMirrorFilterMap) ToTrafficMirrorFilterMapOutput() TrafficMirrorFilterMapOutput {
	return i.ToTrafficMirrorFilterMapOutputWithContext(context.Background())
}

func (i TrafficMirrorFilterMap) ToTrafficMirrorFilterMapOutputWithContext(ctx context.Context) TrafficMirrorFilterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficMirrorFilterMapOutput)
}

type TrafficMirrorFilterOutput struct{ *pulumi.OutputState }

func (TrafficMirrorFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrafficMirrorFilter)(nil)).Elem()
}

func (o TrafficMirrorFilterOutput) ToTrafficMirrorFilterOutput() TrafficMirrorFilterOutput {
	return o
}

func (o TrafficMirrorFilterOutput) ToTrafficMirrorFilterOutputWithContext(ctx context.Context) TrafficMirrorFilterOutput {
	return o
}

// The ARN of the traffic mirror filter.
func (o TrafficMirrorFilterOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficMirrorFilter) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// A description of the filter.
func (o TrafficMirrorFilterOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrafficMirrorFilter) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// List of amazon network services that should be mirrored. Valid values: `amazon-dns`.
func (o TrafficMirrorFilterOutput) NetworkServices() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TrafficMirrorFilter) pulumi.StringArrayOutput { return v.NetworkServices }).(pulumi.StringArrayOutput)
}

// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o TrafficMirrorFilterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *TrafficMirrorFilter) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o TrafficMirrorFilterOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *TrafficMirrorFilter) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type TrafficMirrorFilterArrayOutput struct{ *pulumi.OutputState }

func (TrafficMirrorFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TrafficMirrorFilter)(nil)).Elem()
}

func (o TrafficMirrorFilterArrayOutput) ToTrafficMirrorFilterArrayOutput() TrafficMirrorFilterArrayOutput {
	return o
}

func (o TrafficMirrorFilterArrayOutput) ToTrafficMirrorFilterArrayOutputWithContext(ctx context.Context) TrafficMirrorFilterArrayOutput {
	return o
}

func (o TrafficMirrorFilterArrayOutput) Index(i pulumi.IntInput) TrafficMirrorFilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TrafficMirrorFilter {
		return vs[0].([]*TrafficMirrorFilter)[vs[1].(int)]
	}).(TrafficMirrorFilterOutput)
}

type TrafficMirrorFilterMapOutput struct{ *pulumi.OutputState }

func (TrafficMirrorFilterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TrafficMirrorFilter)(nil)).Elem()
}

func (o TrafficMirrorFilterMapOutput) ToTrafficMirrorFilterMapOutput() TrafficMirrorFilterMapOutput {
	return o
}

func (o TrafficMirrorFilterMapOutput) ToTrafficMirrorFilterMapOutputWithContext(ctx context.Context) TrafficMirrorFilterMapOutput {
	return o
}

func (o TrafficMirrorFilterMapOutput) MapIndex(k pulumi.StringInput) TrafficMirrorFilterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TrafficMirrorFilter {
		return vs[0].(map[string]*TrafficMirrorFilter)[vs[1].(string)]
	}).(TrafficMirrorFilterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficMirrorFilterInput)(nil)).Elem(), &TrafficMirrorFilter{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficMirrorFilterArrayInput)(nil)).Elem(), TrafficMirrorFilterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficMirrorFilterMapInput)(nil)).Elem(), TrafficMirrorFilterMap{})
	pulumi.RegisterOutputType(TrafficMirrorFilterOutput{})
	pulumi.RegisterOutputType(TrafficMirrorFilterArrayOutput{})
	pulumi.RegisterOutputType(TrafficMirrorFilterMapOutput{})
}
