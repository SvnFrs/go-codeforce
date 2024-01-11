#include <stdio.h>

int main() {
    int numCases, i, start, multiple, result;

    scanf("%d", &numCases);

    for (i = 0; i < numCases; i++) {
        scanf("%d%d", &start, &multiple);

        // Calculate the result using a formula
        result = (start % multiple == 0) ? 0 : multiple - (start % multiple);

        printf("%d\n", result);
    }

    return 0;
}
