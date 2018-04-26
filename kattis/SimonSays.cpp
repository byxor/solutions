#include <stdlib.h>
#include <iostream>
#include <cstdio>
#include <string>

#define TEST_MODE false

using namespace std;

bool AssertEqual(int expected, int actual) {
  if (expected == actual)
    printf(".");
  else
    printf("F\nExpected: %d\nActual: %d\n", expected, actual);
}

bool AssertEqual(string expected, string actual) {
  if (expected == actual)
    printf(".");
  else
    cout << "F\nExpected: \"" << expected << "\"\nActual: \"" << actual << "\"\n";
}

/* SOURCE */

inline bool ShouldRepeat(string input) {
  string simonString = "Simon says";
  return input.compare(0, simonString.length(), simonString) == 0;
}

inline string GetAction(string input) {
  return input.substr(10, input.length() - 10);
}

int Run() {
  std::string input;
  getline(cin, input);
  int N = atoi(input.c_str());
  
  for (int i = 0; i < N; i++) {
    string whatSimonSaid;
    getline(cin, whatSimonSaid);
    if (ShouldRepeat(whatSimonSaid))
      cout << GetAction(whatSimonSaid) << endl;
  }
  
  return 0;
}

/* TESTS */

void TestShouldRepeat() {
  printf("TestShouldRepeat\n");
  AssertEqual(false, ShouldRepeat(""));
  AssertEqual(true, ShouldRepeat("Simon says"));
  AssertEqual(false, ShouldRepeat("Sauce"));
  AssertEqual(false, ShouldRepeat("Simon say"));
  AssertEqual(false, ShouldRepeat("simon says"));
  AssertEqual(true, ShouldRepeat("Simon says hi"));
  printf("\n");
}

void TestGetAction() {
  printf("\nTestGetAction\n");
  AssertEqual("", GetAction("Simon says"));
  AssertEqual(" ", GetAction("Simon says "));
  AssertEqual(" smile!", GetAction("Simon says smile!"));
  AssertEqual(" go to the park!", GetAction("Simon says go to the park!"));
  printf("\n");
}

int RunTests() {
  TestShouldRepeat();
  TestGetAction();
  printf("\n");
  return 0;
}

int main() {
  return TEST_MODE ? RunTests() : Run();
}
