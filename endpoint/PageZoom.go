package endpoint

type PageZoom int

// Specifies horizontal alignment. Center alignment is the default value for alignments.

const (

	// Keep unchanged.
	Retain PageZoom = 0

	// Fit page in window.
	FitPage PageZoom = 1

	// Fit width of page in window.
	FitWidth PageZoom = 2

	// Fit height of page in window.
	FitHeight PageZoom = 3

	// Fit all elements of page in window.
	FitBox PageZoom = 4
)
