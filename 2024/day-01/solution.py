def parseInput(filename: str) -> tuple[list[int], list[int]]:
    # Build empty tuple to hold values
    result = ([],[])
    file = open(filename, "r")
    lines = file.readlines()

    # Parse out lines to fill each list in the tuple
    for line in lines:
        lr = line.split()
        result[0].append(int(lr[0]))
        result[1].append(int(lr[1]))

    # Sort these values in memory once, ideally.
    result[0].sort()
    result[1].sort()

    return result

def solution1(coordinates: tuple[list[int], list[int]]) -> int:
    # Store parameters for return
    distance = 0

    # Take the absolute value of the difference from each value
    for i in range(len(coordinates[0])):
        distance += abs(coordinates[0][i] - coordinates[1][i])

    return distance

def solution2(coordinates: tuple[list[int], list[int]]) -> int:
    # Store right list for iterable mutation
    repeatList = coordinates[1]

    # Store parameters for return
    similarity = 0
    count = 0

    # This is the 3rd iteration, first iteration simply bruteforced the two lists to find all matches and multiplied accordingly
    # Get the index and coordinate value for the left list, check the right list, once a match is found add to similarity
    # # if the left coordinate value was already found, do not iterate over list, just add to similarity again
    for idx, coordinate in enumerate(coordinates[0]):
        if idx > 0 and coordinate == coordinates[0][idx-1]:
            similarity += coordinate*count
            continue
        count = 0
        # For each index, doordinate that's repeated in the right list, add to the multiplier. If the following value doesn't
        # match the previous value, assume all following values also do not match and break the loop. Trim the list so the next
        # iteration doesn't take as long to count through irrelevant values.
        for jdx, repeat in enumerate(repeatList):
            if coordinate == repeat:
                count += 1
            elif count > 0:
                repeatList = repeatList[jdx:]
                break
        similarity += coordinate*count

    return similarity

if __name__=="__main__":
    input = parseInput("input.txt")
    print(solution1(input))
    print(solution2(input))