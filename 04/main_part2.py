def count_cards(card_map, card_list):
    for card in card_list:
        num = card["card"]
        count = card["count"]

        while count != 0:
            card_list.append(card_map[num + count])
            count -= 1

    return len(card_list)


def parse_card(line: str):
    splitted = line.split(": ")

    card_number = splitted[0].split(" ")[-1]

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
        "card": int(card_number),
        "count": len(contained),
    }

    return summary


def get_lines():
    with open("./input.txt", "r") as f:
        lines = f.readlines()
        return list(map(str.strip, lines))


def main():
    lines = get_lines()

    card_list = [parse_card(line) for line in lines]
    card_map = dict(zip(range(1, len(card_list) + 1), card_list))

    count = count_cards(card_map, card_list)

    print(count)


if __name__ == "__main__":
    main()
