#include <stdio.h>

#define SIZE 101

static char input[SIZE];

int main(void) {
//  printf("Hi\n");
  scanf("%s", input);
  for (int i = 0; i < SIZE; i++) {
    char c = input[i];

    if (c == 0)
      break;

    if (c >= 65 && c <= 90) {
      printf("%c", c);
    }
  }
  printf("\n");
  return 0;
}
