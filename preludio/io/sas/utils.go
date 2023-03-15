package sasio

import (
	"encoding/binary"
	"fmt"
	"math"
	"regexp"
)

// Python divmod build-in for Go.
func DivMod(n, d int64) (q, r int64) {
	q = n / d
	r = n % d

	// the numerator or the denominator is negarive (but not both)
	if r != 0 && n*d < 0 {
		q--
		r += d
	}
	return
}

// SAS float numbers.
//
// SAS supports 27 special missing values, allowing the categorization of
// missing data by tagging or labeling missing values using the letters A to Z
// or an underscore.
type SasFloat []byte

var SpecialMissingValueRegex = regexp.MustCompile("^[A-Z_]\x00\x00\x00\x00\x00\x00\x00$")

// Check if float is a SpecialMissingValue from an XPORT-format bytestring.
func IsIbmSpecialMissingValue(ieee float64) bool {
	ulong := math.Float64bits(ieee)

	buff := make([]byte, 8)
	binary.BigEndian.PutUint64(buff, ulong)
	if ((buff[0] >= 'A' && buff[0] <= 'Z') || buff[0] == '_') && (ulong&0x00ffffffffffffff == 0) {
		return true
	}
	return false
}

// Convert IBM-format floating point (bytes) to IEEE 754 64-bit (float).
func (sf *SasFloat) ToIeee() (float64, error) {

	// IBM mainframe:    sign * 0.mantissa * 16 ** (exponent - 64)
	// Python uses IEEE: sign * 1.mantissa * 2 ** (exponent - 1023)

	// Pad-out to 8 bytes if necessary. We expect 2 to 8 bytes, but
	// there's no need to check; bizarre sizes will cause a struct
	// module unpack error.
	ibm := *sf
	for len(ibm) < 8 {
		ibm = append(ibm, 0)
	}

	// parse the 64 bits of IBM float as one 8-byte unsigned long long
	ulong := binary.BigEndian.Uint64(ibm)

	// IBM: 1-bit sign, 7-bits exponent, 56-bits mantissa
	sign := int64(ulong & 0x8000000000000000)
	exponent := int64(ulong&0x7f00000000000000) >> 56
	mantissa := ulong & 0x00ffffffffffffff

	if mantissa == 0 {
		if ibm[0] == 0 || ibm[0] == '\x80' {
			return 0.0, nil
		} else if ibm[0] == '.' {
			return math.NaN(), nil
		} else if (ibm[0] >= 'A' && ibm[0] <= 'Z') || ibm[0] == '_' {
			return math.Float64frombits(ulong & 0xffff000000000000), nil
		} else {
			return 0.0, fmt.Errorf("neither \"true\" zero nor NaN: %s", ibm)
		}
	}

	// IBM-format exponent is base 16, so the mantissa can have up to 3
	// leading zero-bits in the binary mantissa. IEEE format exponent
	// is base 2, so we don't need any leading zero-bits and will shift
	// accordingly. This is one of the criticisms of IBM-format, its
	// wobbling precision.
	shift := int64(0)
	if ulong&0x0080000000000000 != 0 {
		shift = 3
	} else if ulong&0x0040000000000000 != 0 {
		shift = 2
	} else if ulong&0x0020000000000000 != 0 {
		shift = 1
	}

	mantissa >>= shift

	// clear the 1 bit to the left of the binary point
	// this is implicit in IEEE specification
	mantissa &= 0xffefffffffffffff

	// IBM exponent is excess 64, but we subtract 65, because of the
	// implicit 1 left of the radix point for the IEEE mantissa
	exponent -= 65
	// IBM exponent is base 16, IEEE is base 2, so we multiply by 4
	exponent <<= 2
	// IEEE exponent is excess 1023, but we also increment for each
	// right-shift when aligning the mantissa's first 1-bit
	exponent += shift + 1023

	// IEEE: 1-bit sign, 11-bits exponent, 52-bits mantissa
	// We didn't shift the sign bit, so it's already in the right spot
	return math.Float64frombits(uint64(sign|exponent<<52) | mantissa), nil
}

// Convert Python floating point numbers to IBM-format (bytes).
func SasFloatFromIeee(ieee float64) (*SasFloat, error) {
	// Python uses IEEE: sign * 1.mantissa * 2 ** (exponent - 1023)
	// IBM mainframe:    sign * 0.mantissa * 16 ** (exponent - 64)

	if ieee == 0.0 {
		val := SasFloat([]byte{0, 0, 0, 0, 0, 0, 0, 0})
		return &val, nil
	}

	// The IBM hexadecimal floating point (HFP) format represents the number
	// zero with all zero bits. All zero bits is the "true zero" or normalized
	// form of zero. Any values for the sign and exponent can be used if the
	// mantissa portion of the encoding is all zero bits, but an IBM machine
	// might lose precision when performing arithmetic with alternative zero
	// representations. With that in mind, and because this format was not
	// defined with a mechanism for not-a-number (NaN) values, SAS uses
	// alternative zero encodings to represent NaN. By default, a SAS missing
	// value is encoded with an ASCII-encoded period (".") as the first byte.

	if math.IsNaN(ieee) {
		val := SasFloat([]byte{'.', 0, 0, 0, 0, 0, 0, 0})
		return &val, nil
	}
	if math.IsInf(ieee, 0) {
		return nil, fmt.Errorf("cannot convert infinity")
	}

	ulong := math.Float64bits(ieee)

	sign := int64((ulong & (1 << 63)) >> 63)          // 1-bit     sign
	exponent := int64((ulong&(0x7ff<<52))>>52) - 1023 // 11-bits   exponent
	mantissa := int64(ulong & 0x000fffffffffffff)     // 52-bits   mantissa/significand

	// Special Missing Values
	buff := make([]byte, 8)
	binary.BigEndian.PutUint64(buff, ulong)
	if ((buff[0] >= 'A' && buff[0] <= 'Z') || buff[0] == '_') && (ulong&0x00ffffffffffffff == 0) {
		val := SasFloat([]byte{buff[0], 0, 0, 0, 0, 0, 0, 0})
		return &val, nil
	}

	if exponent > 248 {
		return nil, fmt.Errorf("cannot store magnitude more than ~ 16 ** 63 as IBM-format")
	}
	if exponent < -260 {
		return nil, fmt.Errorf("cannot store magnitude less than ~ 16 ** -65 as IBM-format")
	}

	// IEEE mantissa has an implicit 1 left of the radix:    1.significand
	// IBM mantissa has an implicit 0 left of the radix:     0.significand
	// We must bitwise-or the implicit 1.mmm into the mantissa
	// later we will increment the exponent to account for this change
	mantissa = 0x0010000000000000 | mantissa

	// IEEE exponents are for base 2:    mantissa * 2 ** exponent
	// IBM exponents are for base 16:    mantissa * 16 ** exponent
	// We must divide the exponent by 4, since 16 ** x == 2 ** (4 * x)
	q, remainder := DivMod(exponent, 4)
	exponent = q

	// We don't want to lose information;
	// the remainder from the divided exponent adjusts the mantissa
	mantissa <<= remainder

	// Increment exponent, because of earlier adjustment to mantissa
	// this corresponds to the 1.mantissa vs 0.mantissa implicit bit
	exponent += 1

	// IBM exponents are excess 64
	exponent += 64

	// IBM has 1-bit sign, 7-bits exponent, and 56-bits mantissa.
	// We must shift the sign and exponent into their places.
	sign <<= 63
	exponent <<= 56

	// We lose some precision, but who said floats were perfect?
	buff = make([]byte, 8)
	binary.BigEndian.PutUint64(buff, uint64(sign|exponent|mantissa))
	sf := SasFloat(buff)
	return &sf, nil
}
