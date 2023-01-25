package endpoint

/**
 * Represents the xml response.
 */
type XmlResponse struct {
	JsonResponse

	/* Gets the xml content. */
	Content []string `json:"content,omitempty"`
}
