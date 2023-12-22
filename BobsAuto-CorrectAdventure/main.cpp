#include <iostream>
#include <vector>
#include <sstream>
// hehe changed everything to long
// aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
void findAndPrintCharactersToRemove(const std::string& first, const std::string& second) {
    int length = first.length();
    std::vector<int> indexesToRemove;

    for (int i = 0; i < length; i++) {
        std::string removedString = first.substr(0, i) + first.substr(i + 1);

        if (removedString == second) {
            indexesToRemove.push_back(i + 1);
        }
    }

    std::cout << indexesToRemove.size() << std::endl;
    if (!indexesToRemove.empty()) {
        std::stringstream resultStream;
        for (int i : indexesToRemove) {
            resultStream << i << " ";
        }
        std::cout << resultStream.str() << std::endl;
    }
}

int main() {
    std::string first, second;

    if (std::cin >> first) {
    } else {
        std::cout << "0" << std::endl;
        return 0;
    }

    if (std::cin >> second) {
    } else {
        std::cout << "0" << std::endl;
        return 0;
    }

    findAndPrintCharactersToRemove(first, second);

    return 0;
}
