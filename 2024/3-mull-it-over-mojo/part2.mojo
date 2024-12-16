
# mojo doesn't have regex yet :'(


fn is_valid_chunk(chunk: String) -> Bool:
    var good_chars: String = "0123456789,"
    var good_count = 0
    for c in chunk:
        for d in good_chars:
            if c == d:
                good_count += 1
    return good_count == len(chunk)


fn process_chunk(chunk: String) raises -> Int:
    if not is_valid_chunk(chunk):
        return 0
    num = chunk.split(",")
    return atol(num[0]) * atol(num[1])


fn process_lines(lines: String) raises -> Int:
    var rem_lines = lines.replace("\n", "")
    var sum = 0
    for i in rem_lines.split("do"):
        if i[][:3] == "n't":
            continue
        for j in i[].split("mul("):
            sum += process_chunk(j[].split(")")[0])

    return sum


fn main():
    try:
        var lines = open("input.txt", "r").read()
        var sum = process_lines(lines)
        print(sum)
    except:
        print("Error reading or processing file")

// 80570939