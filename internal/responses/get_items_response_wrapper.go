package responses

import "github.com/aditya109/library-system/internal/models"

// GetItemsResponseWrapper is a list of items returned in the response
// swagger:response GetItemsResponse
type GetItemsResponseWrapper struct {
	// All items in the system
	// in: body
	Body []models.Item
}
