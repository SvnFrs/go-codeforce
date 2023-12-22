#include <stdio.h>
#include <stdlib.h>
#include <string.h>
// final
void find_and_print_characters_to_remove(const char *first, const char *second) {
    int length = strlen(first);
    int *indexes_to_remove = NULL;
    int count = 0;

    char *removed_string = (char *)malloc(length + 1); // +1 for null terminator
    if (!removed_string) {
        perror("Memory allocation failed");
        exit(EXIT_FAILURE);
    }

    for (int i = 0; i < length; i++) {
        snprintf(removed_string, i + 1, "%.*s", i, first);
        snprintf(removed_string + i, length - i, "%s", first + i + 1);

        if (strcmp(removed_string, second) == 0) {
            indexes_to_remove = (int *)realloc(indexes_to_remove, (count + 1) * sizeof(int));
            if (!indexes_to_remove) {
                perror("Memory reallocation failed");
                exit(EXIT_FAILURE);
            }
            indexes_to_remove[count++] = i;
        }
    }

    free(removed_string);

    printf("%d\n", count);
    for (int i = 0; i < count; i++) {
        printf("%d ", indexes_to_remove[i]);
    }
    printf("\n");

    free(indexes_to_remove);
}

int main() {
    char input_line[1024];
    char first[1024];

    if (fgets(input_line, sizeof(input_line), stdin) == NULL) {
        printf("0\n");
        return 0;
    }

    sscanf(input_line, "%[^\n]", first);

    if (fgets(input_line, sizeof(input_line), stdin) == NULL) {
        printf("0\n");
        return 0;
    }

    char second[1024];
    sscanf(input_line, "%[^\n]", second);

    find_and_print_characters_to_remove(first, second);

    return 0;
}
