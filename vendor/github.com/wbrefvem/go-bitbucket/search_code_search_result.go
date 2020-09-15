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

type SearchCodeSearchResult struct {

	Type_ string `json:"type,omitempty"`

	ContentMatchCount int64 `json:"content_match_count,omitempty"`

	ContentMatches []SearchContentMatch `json:"content_matches,omitempty"`

	PathMatches []SearchSegment `json:"path_matches,omitempty"`

	File *CommitFile `json:"file,omitempty"`
}
