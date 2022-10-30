def read_bytecode(bytes_):
    bytemark = bytes_[0:4]
    table_len = int.from_bytes(bytes_[8:16], byteorder="big")
    
    print("BYTE MARK: ", bytemark)
    print("SYMBOL NUM:", table_len, "\n")

    print("TABLE")
    print("-----")
    offset = 16
    symbols = []
    for _ in range(table_len):
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
