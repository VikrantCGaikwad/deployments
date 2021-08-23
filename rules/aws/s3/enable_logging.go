package s3

import (
	"fmt"

	"github.com/aquasecurity/defsec/infra"
	"github.com/aquasecurity/defsec/result"
)

func CheckLoggingIsEnabled(context *infra.Context) []*result.Result {
	var results []*result.Result
	for _, bucket := range context.AWS.S3.Buckets {
		if !bucket.Logging.Enabled.IsTrue() && bucket.ACL.NotEqualTo("log-delivery-write") {
			results = append(results, &result.Result{
				Description: fmt.Sprintf("Resource '%s' does not have logging enabled", bucket.Reference),
				Location:    bucket.Logging.Enabled.Range,
			})
		}
	}
	return results
}