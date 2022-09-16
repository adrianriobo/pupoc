package aws

// https://github.com/pulumi/automation-api-examples/blob/f5444239378c9891250ee367e9c2a6f26149f375/go/multi_stack_orchestration/main.go#L132

// this is the inline pulumi function for our s3 bucket stack
// func websiteFunc(ctx *pulumi.Context) error {
// 	// similar go git_repo_program, our program defines a s3 website.
// 	// here we create the bucket
// 	siteBucket, err := s3.NewBucket(ctx, "s3-website-bucket", &s3.BucketArgs{
// 		Website: s3.BucketWebsiteArgs{
// 			IndexDocument: pulumi.String("index.html"),
// 		},
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	// Set the access policy for the bucket so all objects are readable.
// 	if _, err := s3.NewBucketPolicy(ctx, "bucketPolicy", &s3.BucketPolicyArgs{
// 		Bucket: siteBucket.ID(), // refer to the bucket created earlier
// 		Policy: pulumi.Any(map[string]interface{}{
// 			"Version": "2012-10-17",
// 			"Statement": []map[string]interface{}{
// 				{
// 					"Effect":    "Allow",
// 					"Principal": "*",
// 					"Action": []interface{}{
// 						"s3:GetObject",
// 					},
// 					"Resource": []interface{}{
// 						pulumi.Sprintf("arn:aws:s3:::%s/*", siteBucket.ID()), // policy refers to bucket name explicitly
// 					},
// 				},
// 			},
// 		}),
// 	}); err != nil {
// 		return err
// 	}

// 	// export the website URL
// 	ctx.Export("websiteUrl", siteBucket.WebsiteEndpoint)
// 	// also export the bucketID for Object stack to refer to
// 	ctx.Export("bucketID", siteBucket.ID())
// 	return nil
// }
