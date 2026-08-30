package preludiocore

type UserDefinedFunction func(*ByteEater)

type __p_symbol__ string

type __p_list__ []__p_intern__

type __p_intern_tag__ uint8

const (
	// PRELUDIO_INTERNAL_TAG_ERROR       __p_intern_tag__ = 0
	PRELUDIO_INTERNAL_TAG_EXPRESSION  __p_intern_tag__ = 1
	PRELUDIO_INTERNAL_TAG_NAMED_PARAM __p_intern_tag__ = 2
	PRELUDIO_INTERNAL_TAG_ASSIGNMENT  __p_intern_tag__ = 3
	PRELUDIO_INTERNAL_TAG_BEGIN_FRAME __p_intern_tag__ = 4
)

// VM parameters
const DEFAULT_OUTPUT_SNIPPET_LENGTH = 10
