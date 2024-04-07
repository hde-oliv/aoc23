#include <ctype.h>
#include <fcntl.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

typedef struct s_gear {
	int x;
	int y;
} t_gear;

typedef struct s_number {
	int value;
	int length;
	int x1;
	int x2;
	int x3;
	int y1;
	int y2;
	int y3;
} t_number;

int count_gears(char **lines) {
	int result = 0;

	for (int i = 0; i < 141; i++) {
		for (int j = 0; j < 142; j++) {
			if (lines[i][j] == '*') {
				result++;
			}
		}
	}

	return result;
}

void allocate_gears(t_gear **array, char **lines) {
	int idx = 0;

	for (int i = 0; i < 141; i++) {
		for (int j = 0; j < 142; j++) {
			if (lines[i][j] == '*') {
				t_gear *gear = calloc(1, sizeof(t_gear));
				gear->x		 = j;
				gear->y		 = i;
				array[idx]	 = gear;

				idx++;
			}
		}
	}
}

int count_numbers(char **lines) {
	int result = 0;

	for (int i = 0; i < 141; i++) {
		for (int j = 0; j < 142; j++) {
			if (isdigit(lines[i][j])) {
				int offset = 1;
				while (isdigit(lines[i][j + offset])) {
					offset++;
				}
				j += offset - 1;

				result++;
			}
		}
	}

	return result;
}

void allocate_numbers(t_number **numbers, char **lines) {
	int idx = 0;

	for (int i = 0; i < 141; i++) {
		for (int j = 0; j < 142; j++) {
			if (isdigit(lines[i][j])) {
				char *digit = calloc(4, sizeof(char));
				digit[0]	= lines[i][j];

				int offset = 1;
				while (isdigit(lines[i][j + offset])) {
					digit[offset] = lines[i][j + offset];
					offset++;
				}

				t_number *number = calloc(1, sizeof(t_number));
				number->value	 = atoi(digit);
				number->length	 = strlen(digit);
				number->x1		 = j;
				number->y1		 = i;

				if (digit[1] != '\0') {
					number->x2 = j + 1;
					number->y2 = i;
				}
				if (digit[2] != '\0') {
					number->x3 = j + 2;
					number->y3 = i;
				}

				j += offset - 1;
				free(digit);

				numbers[idx] = number;
				idx++;
			}
		}
	}
}

int validate_number(t_number *n, int x, int y) {
	if (n->x1 == x && n->y1 == y) {
		return 1;
	}

	if (n->length >= 2 && n->x2 == x && n->y2 == y) {
		return 1;
	}

	if (n->length >= 3 && n->x3 == x && n->y3 == y) {
		return 1;
	}

	return 0;
}

int validate(char **lines) {
	int result = 0;

	int		 gear_count = count_gears(lines);
	t_gear **gear_array = calloc(gear_count + 1, sizeof(t_gear *));
	allocate_gears(gear_array, lines);

	int		   number_count = count_numbers(lines);
	t_number **number_array = calloc(number_count + 1, sizeof(t_number *));
	allocate_numbers(number_array, lines);

	for (int i = 0; i < gear_count; i++) {
		t_gear	 *tmp = gear_array[i];
		t_number *n_tmp[8];
		int		  n = 0;

		for (int j = 0; j < number_count; j++) {
			int a = validate_number(number_array[j], tmp->x - 1, tmp->y + 1);
			int b = validate_number(number_array[j], tmp->x, tmp->y + 1);
			int c = validate_number(number_array[j], tmp->x + 1, tmp->y + 1);

			int d = validate_number(number_array[j], tmp->x - 1, tmp->y);
			int e = validate_number(number_array[j], tmp->x + 1, tmp->y);

			int f = validate_number(number_array[j], tmp->x - 1, tmp->y - 1);
			int g = validate_number(number_array[j], tmp->x, tmp->y - 1);
			int h = validate_number(number_array[j], tmp->x + 1, tmp->y - 1);

			if (a || b || c || d || e || f || g || h) {
				n_tmp[n] = number_array[j];
				n++;
			}
		}

		if (n == 2) {
			t_number *one = n_tmp[0];
			t_number *two = n_tmp[1];

			int t = one->value * two->value;

			result += t;
		}
	}

	for (int i = 0; i < gear_count; i++) {
		free(gear_array[i]);
	}
	free(gear_array);

	for (int i = 0; i < number_count; i++) {
		free(number_array[i]);
	}
	free(number_array);

	return result;
}

char **get_lines(int fd) {
	char **arr = calloc(141, sizeof(char *));

	for (int i = 0; i < 141; i++) {
		char *ptr = calloc(142, sizeof(char));

		if (read(fd, ptr, 141) == -1) {
			printf("Error on read - get_lines\n");
			exit(1);
		}

		arr[i] = ptr;
	}

	return arr;
}

// NOTE: lol, got it on first try
int main(void) {
	int fd = open("./input.txt", O_RDONLY);

	char **lines = get_lines(fd);

	int result = validate(lines);

	printf("result = %d\n", result);

	for (int i = 0; i < 141; i++) {
		free(lines[i]);
	}

	free(lines);
}
