#include <iostream>
#include <climits>

int main() {
    int testCases;
    std::cin >> testCases;

    for (int i = 0; i < testCases; ++i) {
        int boxes;
        std::cin >> boxes;

        int totalReductions = 0;
        int minValue = INT_MAX;

        for (int j = 0; j < boxes; ++j) {
            int marble;
            std::cin >> marble;

            totalReductions += marble;
            if (marble < minValue) {
                minValue = marble;
            }
        }
        totalReductions -= minValue * boxes;

        std::cout << totalReductions << std::endl;
    }

    return 0;
}
