#include <stdio.h>


// int main() {
//     char words[18][5] = {"1", "2", "3", "4", "5", "6", "7", "8", "9", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"};
//     printf("starting");
//     for (int i=0; i<18; i++) {
//         char word = words[i];
//         printf("word get\n");
//         printf('%c', word);
//     }
//     printf("done");
// }


int main()
{
  char *arr[] = {"1" "one", "eight"};
  printf("String array Elements are:\n");
   
  for (int i = 0; i < 3; i++) 
  {
    printf("%s\n", arr[i]);
  }
  return 0;
}