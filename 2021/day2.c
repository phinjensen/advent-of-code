#include <stdio.h>
#include <string.h>
#include <stdlib.h>

int main() {
    char command[8];
    int val, depth = 0, distance = 0, aim = 0;
    while (scanf("%s %d\n", command, &val) != EOF) {
        if (strcmp(command, "forward") == 0) {
            distance += val;
            depth += aim * val;
        } else if (strcmp(command, "down") == 0) {
            aim += val;
        } else if (strcmp(command, "up") == 0) {
            aim -= val;
        } else {
            fprintf(stderr, "Invalid command\n");
            exit(1);
        }
    }
    printf("Depth: %d, distance: %d, together: %d\n", depth, distance, depth*distance);
}

/* Part 1:
int main() {
    char command[8];
    int val, depth = 0, distance = 0;
    while (scanf("%s %d\n", command, &val) != EOF) {
        if (strcmp(command, "forward") == 0) {
            distance += val;
        } else if (strcmp(command, "down") == 0) {
            depth += val;
        } else if (strcmp(command, "up") == 0) {
            depth -= val;
        } else {
            fprintf(stderr, "Invalid command\n");
            exit(1);
        }
    }
    printf("Depth: %d, distance: %d, together: %d\n", depth, distance, depth*distance);
}
*/
