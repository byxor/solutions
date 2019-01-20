#include <stdio.h>

int factorials[] = {1, 1, 2, 6, 4, 0, 0, 0, 0, 0, 0};

int main(void) {
  int tests, n;
  scanf("%d", &tests);

  for (int i = 0; i < tests; i++) {
    scanf("%d", &n);
    printf("%d\n", factorials[n]);
  }

  return 0;
}
