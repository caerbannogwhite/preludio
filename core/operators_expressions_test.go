package preludiocore

import (
	"math"
	"testing"
)

func init() {
	be = new(ByteEater).InitVM()
}

func Test_Expressions(t *testing.T) {
	var err error

	be.RunSource(`true`)
	if err = checkCurrentResult(be, true); err != nil {
		t.Error(err)
	}

	be.RunSource(`false`)
	if err = checkCurrentResult(be, false); err != nil {
		t.Error(err)
	}

	be.RunSource(`true * false`)
	if err = checkCurrentResult(be, int64(0)); err != nil {
		t.Error(err)
	}

	be.RunSource(`true / false`)
	if err = checkCurrentResult(be, math.Inf(1)); err != nil {
		t.Error(err)
	}

	be.RunSource(`true % false`)
	if err = checkCurrentResult(be, math.NaN()); err != nil {
		t.Error(err)
	}

	be.RunSource(`true ** false`)
	if err = checkCurrentResult(be, int64(1)); err != nil {
		t.Error(err)
	}

	be.RunSource(`true + false`)
	if err = checkCurrentResult(be, int64(1)); err != nil {
		t.Error(err)
	}

	be.RunSource(`true - false`)
	if err = checkCurrentResult(be, int64(1)); err != nil {
		t.Error(err)
	}

	be.RunSource(`true and false`)
	if err = checkCurrentResult(be, false); err != nil {
		t.Error(err)
	}

	be.RunSource(`true or false`)
	if err = checkCurrentResult(be, true); err != nil {
		t.Error(err)
	}

	be.RunSource(`not true`)
	if err = checkCurrentResult(be, false); err != nil {
		t.Error(err)
	}

	be.RunSource(`true or (false and true)`)
	if err = checkCurrentResult(be, true); err != nil {
		t.Error(err)
	}

	be.RunSource(`true or not false and true or not true`)
	if err = checkCurrentResult(be, true); err != nil {
		t.Error(err)
	}

	be.RunSource(`1 * 5`)
	if err = checkCurrentResult(be, int64(5)); err != nil {
		t.Error(err)
	}

	be.RunSource(`1 / 3`)
	if err = checkCurrentResult(be, float64(0.3333333333333333)); err != nil {
		t.Error(err)
	}

	be.RunSource(`4682 % 427`)
	if err = checkCurrentResult(be, float64(412)); err != nil {
		t.Error(err)
	}

	be.RunSource(`3 ** 4`)
	if err = checkCurrentResult(be, int64(81)); err != nil {
		t.Error(err)
	}

	be.RunSource(`2 ** (2 + 1 * 2)`)
	if err = checkCurrentResult(be, int64(16)); err != nil {
		t.Error(err)
	}

	be.RunSource(`1 - 2`)
	if err = checkCurrentResult(be, int64(-1)); err != nil {
		t.Error(err)
	}

	be.RunSource(`1 + 2`)
	if err = checkCurrentResult(be, int64(3)); err != nil {
		t.Error(err)
	}

	be.RunSource(`+1 + 2`)
	if err = checkCurrentResult(be, int64(3)); err != nil {
		t.Error(err)
	}

	be.RunSource(`-1`)
	if err = checkCurrentResult(be, int64(-1)); err != nil {
		t.Error(err)
	}

	be.RunSource(`-1 + 2`)
	if err = checkCurrentResult(be, int64(1)); err != nil {
		t.Error(err)
	}

	be.RunSource(`-1.0 - 2`)
	if err = checkCurrentResult(be, float64(-3)); err != nil {
		t.Error(err)
	}

	be.RunSource(`1.325235e-3 * 5`)
	if err = checkCurrentResult(be, float64(0.006626175)); err != nil {
		t.Error(err)
	}

	be.RunSource(`1.325235e-3 / 3`)
	if err = checkCurrentResult(be, float64(0.00044174499999999995)); err != nil {
		t.Error(err)
	}

	be.RunSource(`"hello" + "world"`)
	if err = checkCurrentResult(be, "helloworld"); err != nil {
		t.Error(err)
	}

	be.RunSource(`1 + 2 * 3 - 4 + 5 * 6`)
	if err = checkCurrentResult(be, int64(33)); err != nil {
		t.Error(err)
	}

	be.RunSource(`1 + 2 * 3 - 4 + 5 * 6 % 7 + "hello"`)
	if err = checkCurrentResult(be, "5hello"); err != nil {
		t.Error(err)
	}

	be.RunSource(`3.4 + 2.3 * 3.2 - 4.1 + 5.0 * 6.9`)
	if err = checkCurrentResult(be, float64(41.16)); err != nil {
		t.Error(err)
	}

	be.RunSource(`(1 + 2) * (3 - 4) + 5 * 6`)
	if err = checkCurrentResult(be, int64(27)); err != nil {
		t.Error(err)
	}

	be.RunSource(`(1 + (2 * 3)) - (4 + (5 * (6 % 7 + 8))) / ((9) + (10 * 11 - 12 % 13))`)
	if err = checkCurrentResult(be, float64(6.308411214953271)); err != nil {
		t.Error(err)
	}

	be.RunSource(`(1 + (2 * (3 - (4 + (5 * (6 % (7 + (8 - (9 + (10 * (11 - (12 + (13 * (14 % (15 + (16 - (17 + (18 * (19 - (20 + (21 * (22 % (23 + (24 - (25 + (26 * (27 - (28 + (29 * (30 - (31 + (32 * (33 - (34 + (35 * (36 % (37 + (38 - (39 + (40 * (41 - (42 + (43 * (44 % (45 + (46 - (47 + (48 * (49 - (50 + (51 * (52 % (53 + (54 - (55 + (56 * (57 - (58 + (59 * (60 % (61 + (62 - (63 + (64 * (65 - (66 + (67 * (68 % (69 + (70 - (71 + (72 * (73 - (74 + (75 * (76 % (77 + (78 - (79 + (80 * (81 - (82 + (83 * (84 % (85 + (86 - (87 + (88 * (89 - (90 + (91 * (92 % (93 + (94 - (95 + (96 * (97 - (98 + (99 * (100))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))`)
	if err = checkCurrentResult(be, float64(-61.0)); err != nil {
		t.Error(err)
	}

	be.RunSource(`1e30 / 1.000001 / 1.000002 / 1.000003 / 1.000004 / 1.000005 / 1.000006 / 1.000007 / 1.000008 / 1.000009 / 1.000010 / 1.000011 / 1.000012 / 1.000013 / 1.000014 / 1.000015 / 1.000016 / 1.000017 / 1.000018 / 1.000019 / 1.000020 / 1.000021 / 1.000022 / 1.000023 / 1.000024 / 1.000025 / 1.000026 / 1.000027 / 1.000028 / 1.000029 / 1.000030 / 1.000031 / 1.000032 / 1.000033 / 1.000034 / 1.000035 / 1.000036 / 1.000037 / 1.000038 / 1.000039 / 1.000040 / 1.000041 / 1.000042 / 1.000043 / 1.000044 / 1.000045 / 1.000046 / 1.000047 / 1.000048 / 1.000049 / 1.000050 / 1.000051 / 1.000052 / 1.000053 / 1.000054 / 1.000055 / 1.000056 / 1.000057 / 1.000058 / 1.000059 / 1.000060 / 1.000061 / 1.000062 / 1.000063 / 1.000064 / 1.000065 / 1.000066 / 1.000067 / 1.000068 / 1.000069 / 1.000070 / 1.000071 / 1.000072 / 1.000073 / 1.000074 / 1.000075 / 1.000076 / 1.000077 / 1.000078 / 1.000079 / 1.000080 / 1.000081 / 1.000082 / 1.000083 / 1.000084 / 1.000085 / 1.000086 / 1.000087 / 1.000088 / 1.000089 / 1.000090 / 1.000091 / 1.000092 / 1.000093 / 1.000094 / 1.000095 / 1.000096 / 1.000097 / 1.000098 / 1.000099 / 1.000100`)
	if err = checkCurrentResult(be, float64(9.949628981268441e+29)); err != nil {
		t.Error(err)
	}
}

func Test_Lists(t *testing.T) {
	var err error

	be.RunSource(`[true, false, true]`)
	if err = checkCurrentResult(be, []bool{true, false, true}); err != nil {
		t.Error(err)
	}

	be.RunSource(`[1, 2, 3]`)
	if err = checkCurrentResult(be, []int64{1, 2, 3}); err != nil {
		t.Error(err)
	}

	be.RunSource(`[1.0, 2.0, 3.0]`)
	if err = checkCurrentResult(be, []float64{1.0, 2.0, 3.0}); err != nil {
		t.Error(err)
	}

	be.RunSource(`["hello", "world", "this is a string"]`)
	if err = checkCurrentResult(be, []string{"hello", "world", "this is a string"}); err != nil {
		t.Error(err)
	}
}

func Test_Lists_Expressions(t *testing.T) {
	var err error

	// TODO: Fix this https://github.com/caerbannogwhite/preludio/issues/7#issue-1888727166
	// be.RunSource(`not [true, false, true]`)
	//
	// if err = checkCurrentResult(be, []bool{false, true, false}); err != nil {
	// 	t.Error(err)
	// }

	be.RunSource(`false and [false, true, false]`)
	if err = checkCurrentResult(be, []bool{false, false, false}); err != nil {
		t.Error(err)
	}

	be.RunSource(`[false, true, false] - 3.1`)
	if err = checkCurrentResult(be, []float64{-3.1, -2.1, -3.1}); err != nil {
		t.Error(err)
	}

	be.RunSource(`[true, false, true] or [false, true, false]`)
	if err = checkCurrentResult(be, []bool{true, true, true}); err != nil {
		t.Error(err)
	}

	be.RunSource(`[true, false, true] + [true, false, true]`)
	if err = checkCurrentResult(be, []int64{2, 0, 2}); err != nil {
		t.Error(err)
	}

	be.RunSource(`[1, 2, 3] ** 3`)
	if err = checkCurrentResult(be, []int64{1, 8, 27}); err != nil {
		t.Error(err)
	}

	be.RunSource(`[1, 2, 3] * [-1, 2, -3]`)
	if err = checkCurrentResult(be, []int64{-1, 4, -9}); err != nil {
		t.Error(err)
	}

	be.RunSource(`[1.0, 2.0, 3.0] / [1.0, 2.0, 3.0]`)
	if err = checkCurrentResult(be, []float64{1.0, 1.0, 1.0}); err != nil {
		t.Error(err)
	}

	be.RunSource(`["hello", "foo", "this is a string"] + "!"`)
	if err = checkCurrentResult(be, []string{"hello!", "foo!", "this is a string!"}); err != nil {
		t.Error(err)
	}

	be.RunSource(`["hello", "foo", "this is a string"] + [" ", " ", " "] + ["world", "bar", "and this is another string"]`)
	if err = checkCurrentResult(be, []string{"hello world", "foo bar", "this is a string and this is another string"}); err != nil {
		t.Error(err)
	}

	// diff sizes
	be.RunSource(`["hello", "foo", "this is a string"] + [" ", " "]`)
	if be.getLastError() != "binary operator + not supported between String[3] and String[2]" {
		t.Error("Expected error, got", be.getLastError())
	}
}

func Test_Assignements(t *testing.T) {
	var source string
	var err error

	source = `
	let a = 1
	let b = 2
	let c = 3
	c = 6
	a + b * 3 - 4 + 5 * c
	`

	be.RunSource(source)
	if err = checkCurrentResult(be, int64(33)); err != nil {
		t.Error(err)
	}

	source = `
	let d = (1.000000000001 * (1.000000000002 * (1.000000000003 * (1.000000000004 * (1.000000000005 * (1.000000000006 * (1.000000000007 * (1.000000000008 * (1.000000000009 * (1.000000000010 * (1.000000000011 * (1.000000000012 * (1.000000000013 * (1.000000000014 * (1.000000000015 * (1.000000000016 * (1.000000000017 * (1.000000000018 * (1.000000000019 * (1.000000000020 * (1.000000000021 * (1.000000000022 * (1.000000000023 * (1.000000000024 * (1.000000000025 * (1.000000000026 * (1.000000000027 * (1.000000000028 * (1.000000000029 * (1.000000000030 * (1.000000000031 * (1.000000000032 * (1.000000000033 * (1.000000000034 * (1.000000000035 * (1.000000000036 * (1.000000000037 * (1.000000000038 * (1.000000000039 * (1.000000000040 * (1.000000000041 * (1.000000000042 * (1.000000000043 * (1.000000000044 * (1.000000000045 * (1.000000000046 * (1.000000000047 * (1.000000000048 * (1.000000000049 * (1.000000000050))))))))))))))))))))))))))))))))))))))))))))))))))
	let e = 1.000000000051 * (1.000000000052 * (1.000000000053 * (1.000000000054 * (1.000000000055 * (1.000000000056 * (1.000000000057 * (1.000000000058 * (1.000000000059 * (1.000000000060 * (1.000000000061 * (1.000000000062 * (1.000000000063 * (1.000000000064 * (1.000000000065 * (1.000000000066 * (1.000000000067 * (1.000000000068 * (1.000000000069 * (1.000000000070 * (1.000000000071 * (1.000000000072 * (1.000000000073 * (1.000000000074 * (1.000000000075 * (1.000000000076 * (1.000000000077 * (1.000000000078 * (1.000000000079 * (1.000000000080 * (1.000000000081 * (1.000000000082 * (1.000000000083 * (1.000000000084 * (1.000000000085 * (1.000000000086 * (1.000000000087 * (1.000000000088 * (1.000000000089 * (1.000000000090 * (1.000000000091 * (1.000000000092 * (1.000000000093 * (1.000000000094 * (1.000000000095 * (1.000000000096 * (1.000000000097 * (1.000000000098 * (1.000000000099 * (1.000000000100)))))))))))))))))))))))))))))))))))))))))))))))))
	let f = 1.000000000101 * (1.000000000102 * (1.000000000103 * (1.000000000104 * (1.000000000105 * (1.000000000106 * (1.000000000107 * (1.000000000108 * (1.000000000109 * (1.000000000110 * (1.000000000111 * (1.000000000112 * (1.000000000113 * (1.000000000114 * (1.000000000115 * (1.000000000116 * (1.000000000117 * (1.000000000118 * (1.000000000119 * (1.000000000120 * (1.000000000121 * (1.000000000122 * (1.000000000123 * (1.000000000124 * (1.000000000125 * (1.000000000126 * (1.000000000127 * (1.000000000128 * (1.000000000129 * (1.000000000130 * (1.000000000131 * (1.000000000132 * (1.000000000133 * (1.000000000134 * (1.000000000135 * (1.000000000136 * (1.000000000137 * (1.000000000138 * (1.000000000139 * (1.000000000140 * (1.000000000141 * (1.000000000142 * (1.000000000143 * (1.000000000144 * (1.000000000145 * (1.000000000146 * (1.000000000147 * (1.000000000148 * (1.000000000149 * (1.000000000150)))))))))))))))))))))))))))))))))))))))))))))))))
	let g = 1.000000000151 * (1.000000000152 * (1.000000000153 * (1.000000000154 * (1.000000000155 * (1.000000000156 * (1.000000000157 * (1.000000000158 * (1.000000000159 * (1.000000000160 * (1.000000000161 * (1.000000000162 * (1.000000000163 * (1.000000000164 * (1.000000000165 * (1.000000000166 * (1.000000000167 * (1.000000000168 * (1.000000000169 * (1.000000000170 * (1.000000000171 * (1.000000000172 * (1.000000000173 * (1.000000000174 * (1.000000000175 * (1.000000000176 * (1.000000000177 * (1.000000000178 * (1.000000000179 * (1.000000000180 * (1.000000000181 * (1.000000000182 * (1.000000000183 * (1.000000000184 * (1.000000000185 * (1.000000000186 * (1.000000000187 * (1.000000000188 * (1.000000000189 * (1.000000000190 * (1.000000000191 * (1.000000000192 * (1.000000000193 * (1.000000000194 * (1.000000000195 * (1.000000000196 * (1.000000000197 * (1.000000000198 * (1.000000000199 * (1.000000000200)))))))))))))))))))))))))))))))))))))))))))))))))
	d * (e * (f * (g * (d * (e * (f * (g)))))))
	`

	be.RunSource(source)
	if err = checkCurrentResult(be, 1.000000040200004); err != nil {
		t.Error(err)
	}

	be.RunSource(`let a = 4`)
	if be.__currentResult != nil {
		t.Error("Expected nil, got", be.__currentResult)
	}

	source = `
	let l = [1.1, 2.02, 3.003, 4.0004, 5.00005, 6.000006, 7.00000007, 8.000000008, 9.0000000009, 10.00000000010]
	l + l * l + l * l + l * l + l * l + l * l + l * l + l * l + l * l + l * l
	`

	be.RunSource(source)
	if err = checkCurrentResult(be, []float64{11.990000000000004, 38.7436, 84.16508100000001, 148.02920144, 230.00455002249993, 330.00065400032395, 448.0000088899999, 584.0000011600001, 738.0000001467001, 910.0000000180999}); err != nil {
		t.Error(err)
	}
}
