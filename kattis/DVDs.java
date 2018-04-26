import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

public class DVDs {
    public static void main(String[] args) {
        BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(System.in));
        SwapCounter swapCounter = new SwapCounter();

        final int TEST_CASES = IO.getAnInt(bufferedReader);

        for (int testCase = 0; testCase < TEST_CASES; testCase++) {
            int numberOfDvds = IO.getAnInt(bufferedReader);
            int[] dvds = IO.getAnIntArray(bufferedReader, numberOfDvds);
            System.out.println(swapCounter.countSwaps(dvds));
        }
    }
}

class SwapCounter {
    int countSwaps(int[] dvds) {
        if (dvds.length <= 1)
            return 0;
        else {
            int swaps = 0;
            int expected = 1;
            for (int number : dvds)
                if (number != expected) swaps++;
                else expected++;
            return swaps;
        }
    }
}

class IO {
    public static int getAnInt(BufferedReader bufferedReader) {
        try {
            return Integer.parseInt(bufferedReader.readLine());
        } catch (IOException exception) {
            exception.printStackTrace();
            return -1;
        }
    }

    public static int[] getAnIntArray(BufferedReader bufferedReader, int size) {
        try {
            int[] numbers = new int[size];
            String[] words = bufferedReader.readLine().split(" ");
            for (int i = 0; i < size; i++)
                numbers[i] = Integer.parseInt(words[i]);
            return numbers;
        } catch (IOException exception) {
            exception.printStackTrace();
            return new int[]{};
        }
    }
}
