/*
 * Bitbucket API
 *
 * Code against the Bitbucket API to automate simple tasks, embed Bitbucket data into your own site, build mobile or desktop apps, or even add custom UI add-ons into Bitbucket itself using the Connect framework.
 *
 * API version: 2.0
 * Contact: support@bitbucket.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package bitbucket

// The definition of a Docker image that can be used for a Bitbucket Pipelines step execution context.
type PipelineImage struct {

	// The name of the image. If the image is hosted on DockerHub the short name can be used, otherwise the fully qualified name is required here.
	Name string `json:"name,omitempty"`

	// The username needed to authenticate with the Docker registry. Only required when using a private Docker image.
	Username string `json:"username,omitempty"`

	// The password needed to authenticate with the Docker registry. Only required when using a private Docker image.
	Password string `json:"password,omitempty"`

	// The email needed to authenticate with the Docker registry. Only required when using a private Docker image.
	Email string `json:"email,omitempty"`
}
