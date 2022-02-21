/*
 * TextMagic API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package TextMagic

type GetUserDedicatedNumbersPaginatedResponse struct {
	Page int32 `json:"page"`
	// The total number of pages.
	PageCount int32 `json:"pageCount"`
	// The number of results per page.
	Limit int32 `json:"limit"`
	Resources []UsersInbound `json:"resources"`
}
