#include <iostream>
#include "mylib.h"

int main() {
    int a = 3;
    int b = 5;

    int sum = Add(a, b);
    std::cout << "The sum of " << a << " and " << b << " is " << sum << std::endl;

    auto p = RetPtr();

    return 0;
}
