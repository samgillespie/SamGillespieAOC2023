#include <string.h>
#include <stdio.h>
#include <time.h>


int char_to_int(char* input) {
    if (strlen(input) == 1) {
        int solution = *input - '0';
        return solution;
    }
    if (strcmp(input, "one") == 0) {
        return 1;
    }
    if (strcmp(input, "two") == 0) {
        return 2;
    }
    if (strcmp(input, "three") == 0) {
        return 3;
    }
    if (strcmp(input, "four") == 0) {
        return 4;
    }
    if (strcmp(input, "five") == 0) {
        return 5;
    }
    if (strcmp(input, "six") == 0) {
        return 6;
    }
    if (strcmp(input, "seven") == 0) {
        return 7;
    }
    if (strcmp(input, "eight") == 0) {
        return 8;
    }
    if (strcmp(input, "nine") == 0) {
        return 9;
    }
    printf("PANIC %s\n", input);
    return 0;
}

int parse_digits(char *input) {
    char *digits[] = {"1", "2", "3", "4", "5", "6", "7", "8", "9"};
    int min_position = 9999;
    int max_position = -1;
    char* min_value = NULL;
    char* max_value = NULL;
    for (int i=0; i < 9; i++) {
        for (int j=0; j< strlen(input); j++) {
            if (input[j] != *digits[i]) {
                continue;
            }
            if (j < min_position) {
                min_position = j;
                min_value = digits[i];
            }
            if (j > max_position) {
                max_position = j;
                max_value = digits[i];
            }
        }
    }
    int tens = char_to_int(min_value);
    int ones = char_to_int(max_value);
    return tens * 10 + ones;
}

char *get_substring(int pos, int len, char *string) {
    char substring[100];
    strncpy(substring, string + pos, len); 
    substring[len] = '\0';
    return &substring;
}

int parse_numbers(char input[]) {
    char *words[] = {"1", "2", "3", "4", "5", "6", "7", "8", "9", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"};
    int min_position = 9999;
    int max_position = -1;
    char* min_value = NULL;
    char* max_value = NULL;
    for (int i=0; i < 18; i++) {
        int word_length = strlen(words[i]);
        for (int j=0; j< strlen(input); j++) {
            char substring[100];
            strncpy(substring, input + j, word_length); 
            substring[word_length] = '\0';         
            if (strcmp(substring, words[i]) != 0) {
                continue;
            }
            if (j < min_position) {
                min_position = j;
                min_value = words[i];
            }
            if (j > max_position) {
                max_position = j;
                max_value = words[i];
            }
        }
    }
    
    int tens = char_to_int(min_value);
    int ones = char_to_int(max_value);
    return tens * 10 + ones;
}

void day1() {
    int rows = 1000;
    FILE *fptr;
    fptr = fopen("inputs/q1.txt", "r");
    if (fptr == NULL) {
        printf("failed to read contents\n");
        return;
    }
    char instruction[100];
    int part_a = 0;
    int part_b = 0;
    
    while(fgets(instruction, 100, fptr)) {
        instruction[strcspn(instruction, "\n")] = 0;
        int digit_value = parse_digits(&instruction);
        part_a += digit_value;
        int number_value = parse_numbers(instruction);
        part_b += number_value;
    }
    // printf("Part a: %d\n", part_a);
    // printf("Part b: %d\n", part_b);
    fclose(fptr);
}


int main()
{
    clock_t start, end;
    start = clock();
    for (int i=0; i<1000; i++)
    day1();
    end = clock();
    double cpu_time_used = ((double) (end - start)) / CLOCKS_PER_SEC;
    printf("fun() took %f seconds to execute \n", cpu_time_used); 
    return 0;
}