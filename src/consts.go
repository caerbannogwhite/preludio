package preludiocore

type UserDefinedFunction func(*ByteEater)

type pSymbol string

type pList []pIntern

type pInternTagType uint8

const (
	// PRELUDIO_INTERNAL_TAG_ERROR       pInternTagType = 0
	PRELUDIO_INTERNAL_TAG_EXPRESSION  pInternTagType = 1
	PRELUDIO_INTERNAL_TAG_NAMED_PARAM pInternTagType = 2
	PRELUDIO_INTERNAL_TAG_ASSIGNMENT  pInternTagType = 3
	PRELUDIO_INTERNAL_TAG_BEGIN_FRAME pInternTagType = 4
)

// VM parameters
const DEFAULT_OUTPUT_SNIPPET_LENGTH = 10
