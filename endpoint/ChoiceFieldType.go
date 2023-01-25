package endpoint

/** Specifies type of the Choice Field.*/
type ChoiceFieldType int

const (
	/** Represents a List box. */
	ListBox ChoiceFieldType = 1
	/** Represents a Combo box. */
	ComboBox ChoiceFieldType = 2
)
