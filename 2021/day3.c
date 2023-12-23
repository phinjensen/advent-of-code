#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#define LENGTH 5
#define LINES 12

int main() {
    char lines[LINES][LENGTH + 1];
    for (int i = 0; scanf("%s\n", lines + i) != EOF; i++);
    for (int bit = 0; bit < LENGTH; bit++) {
        int bit_count = 0;
        for (int line = 0; line < LINES; line++) {
            if (lines[line] == NULL) continue;
            if (lines[line][bit] == '0') bit_count--;
            else bit_count++;
        }
        for (int line = 0; line < LINES; line++) {
            if (lines[line][bit] == '0' && bit_count > 0) lines[line] = NULL;
            else if (lines[line][bit] == '1' && bit_count < 0) lines[line] = NULL;
        }
    }
}

/* Part 1: 
int main() {
    char line[LENGTH + 1];
    int counts[LENGTH] = { 0 };
    while (scanf("%s\n", line) != EOF) {
        for (int i = 0; i < LENGTH; i++) {
            if (line[i] == '0') {
                counts[i]--;
            } else if (line[i] == '1') {
                counts[i]++;
            } else {
                fprintf(stderr, "Invalid bit\n");
                exit(2);
            }
        }
    }
    unsigned int gamma = 0, epsilon = 0;
    for (int i = 0; i < LENGTH; i++) {
        if (i != 0) {
            gamma = gamma << 1;
            epsilon = epsilon << 1;
        }
        if (counts[i] > 0) gamma |= 1;
        if (counts[i] < 0) epsilon |= 1;
    }
    printf("gamma: %d, epsilon: %u, together: %u\n", gamma, epsilon, gamma * epsilon);
}
*/
