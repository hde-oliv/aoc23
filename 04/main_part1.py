import math

# { "winning": [1, 2, 3], "my_numbers": [3, 2, 4], value: 2 }


def parse_card(line: str):
    splitted = line.split(": ")

    card_numbers = splitted[-1].split(" | ")

    winning = card_numbers[0].split(" ")
    my_numbers = card_numbers[1].split(" ")

    winning = [n for n in winning if n]
    my_numbers = [n for n in my_numbers if n]

    winning = list(map(int, winning))
    my_numbers = list(map(int, my_numbers))

    win_test = lambda x: True if x in winning else False

    contained = list(filter(win_test, my_numbers))

    summary = {
        "winning": winning,
        "my_numbers": my_numbers,
        "value": math.floor(2 ** (len(contained) - 1)),
    }

    print(summary)
    return summary


def get_lines():
    with open("./input.txt", "r") as f:
        lines = f.readlines()
        return list(map(str.strip, lines))


def main():
    lines = get_lines()

    cards = [parse_card(line) for line in lines]

    r = 0
    for card in cards:
        r += card["value"]

    print(r)


if __name__ == "__main__":
    main()
