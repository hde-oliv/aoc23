#include <ctype.h>
#include <fcntl.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

int issymbol(char c) {
	if (!isdigit(c) && c != '.' && c != '\n' && c != '\0') {
		return 1;
	}
	return 0;
}

int check_line(char *line, int index, int offset) {
	int result = 0;

	if (line == NULL) {
		return 0;
	}

	for (int i = -1; i < offset + 1; i++) {
		if (index == 0 && i == -1) {
			continue;
		}

		if (issymbol(line[index + i])) {
			result = 1;
		}
	}

	return result;
}

int parse_line(char *p_line, char *a_line, char *n_line) {
	int result = 0;

	for (int i = 0; i < strlen(a_line); i++) {
		if (isdigit(a_line[i])) {
			char *digit = calloc(4, sizeof(char));
			digit[0]	= a_line[i];

			int offset = 1;
			while (isdigit(a_line[i + offset])) {
				digit[offset] = a_line[i + offset];
				offset++;
			}

			int p = check_line(p_line, i, offset);
			int a = check_line(a_line, i, offset);
			int n = check_line(n_line, i, offset);

			if (p || a || n) {
				result += atoi(digit);
			}
			free(digit);
			i += offset - 1;
		}
	}
	return result;
}

int validate(char **lines) {
	int sum = 0;

	sum += parse_line(NULL, lines[0], lines[1]);

	for (int i = 1; i < 140; i++) {
		sum += parse_line(lines[i - 1], lines[i], lines[i + 1]);
	}

	return sum;
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
