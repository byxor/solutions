#include <stdio.h>

static const char *HISS = "hiss\n";
static const char *NO_HISS = "no hiss\n";

#define SIZE 31
static char input[SIZE];

int main(void) {
  scanf("%s", input);

  int s = 0;

  for (int i = 0; i < SIZE; i++) {
    char c = input[i];
    switch(c) {
      case 0:
        break;
      case 's':
        s += 1;
        break;
      default:
        s = 0;
        break;
    }

    if (s >= 2) {
      printf(HISS);
      return 0;
    }
  }

  printf(NO_HISS);
  return 0;
}
