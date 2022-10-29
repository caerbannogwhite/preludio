def read_bytecode(bytes_):
    bytemark = bytes_[0:4]
    table_len = int.from_bytes(bytes_[8:16], byteorder="big")
    print("mark:", bytemark, "symbols:", table_len)
    offset = 16
    symbols = []

    for _ in range(table_len):
        l = int.from_bytes(bytes_[offset:offset+8], byteorder="big")
        offset += 8
        symb = bytes_[offset:offset+l]
        offset += l
        symbols.append(symb)
    
    while offset < len(bytes_):
        ops = []
        for _ in range(3):
            ops.append(int.from_bytes(
                bytes_[offset:offset+8], byteorder="big"))
            offset += 8
        print(ops[0], ops[1], ops[2])
    
    return symbols


def get_bytes(path):
    fh = open(path, "rb")
    bytes_ = fh.read()
    fh.close()
    return bytes_
