package element

/**
 * The type of Compaction to encode.
 */
type compaction int

const (
	/*Byte Compaction. */
	Byte compaction = iota

	/*Text Compaction. */
	TextCompaction compaction = iota

	/*Numeric Compaction. */
	Numeric compaction = iota

	/*All Compactions. */
	Auto compaction = iota
)
