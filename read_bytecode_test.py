
def read_bytecode(bytes_):
    bytemark = bytes_[0:4]
    string_symbol_num = int.from_bytes(bytes_[8:16], byteorder="big")

    print("BYTES:             ", len(bytes_))
    print("BYTE MARK:         ", bytemark)
    print("NUMERIC SYMBOL NUM:", string_symbol_num, "\n")

    # print("NUMERIC SYMBOLS")
    # print("---------------")
    offset = 16
    # symbols = []
    # for _ in range(numeric_symbol_num):
    #     num = int.from_bytes(bytes_[offset:offset+8], byteorder="big")
    #     offset += 8
    #     symbols.append(num)
    #     print(num)
    # print()

    # string_symbol_num = int.from_bytes(
    #     bytes_[offset:offset+8], byteorder="big")
    # offset += 8

    # print("STRING SYMBOL NUM:", string_symbol_num, "\n")

    print("STRING SYMBOLS")
    print("--------------")
    symbols = []
    for _ in range(string_symbol_num):
        l = int.from_bytes(bytes_[offset:offset+8], byteorder="big")
        offset += 8
        symb = bytes_[offset:offset+l]
        offset += l
        symbols.append(symb)
        print(symb)
    print()

    print("OPERATIONS")
    print("----------")
    operations = []
    while offset < len(bytes_):
        ops = []
        for _ in range(3):
            ops.append(int.from_bytes(
                bytes_[offset:offset+8], byteorder="big"))
            offset += 8
        operations.append(ops)
        print(f"{ops[0]:4} {ops[1]:4} {ops[2]:4}")

    return symbols, operations


def get_bytes(path):
    fh = open(path, "rb")
    bytes_ = fh.read()
    fh.close()
    return bytes_
