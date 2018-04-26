import java.util.Arrays;
import java.util.Scanner;

public class ColoringSocks {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int numberOfSocks = scanner.nextInt();
        int capacity = scanner.nextInt();
        int threshold = scanner.nextInt();

        int[] socks = new int[numberOfSocks];
        for (int i = 0; i < numberOfSocks; i++)
            socks[i] = scanner.nextInt();

        SockSorter sockSorter = new SockSorter();
        sockSorter.sort(socks);

        LaundryMachineCounter laundryMachineCounter = new LaundryMachineCounter(capacity, threshold);
        laundryMachineCounter.load(socks);

        int laundryMachineCount = laundryMachineCounter.getCount();

        System.out.println(laundryMachineCount);
    }
}


class LaundryMachine {
    private final int capacity;
    private final int threshold;
    private int sockCount = 0;
    private int maxColour = -1;

    public LaundryMachine(int capacity, int threshold) {
        this.capacity = capacity;
        this.threshold = threshold;
    }

    public void load(int sock) throws LaundryMachineException{
        if (sockCount >= capacity)
            throw new LaundryMachineException();
        else {
            if (maxColour == -1) maxColour = sock + threshold;
            if (sock > maxColour) throw new LaundryMachineException();
            else sockCount++;
        }
    }

    public int count() {
        return sockCount;
    }
}

class LaundryMachineException extends Exception {}

class SockSorter {
    public void sort(int[] socks) {
        Arrays.sort(socks);
    }
}

class LaundryMachineCounter {
    private final int capacity;
    private final int threshold;
    private LaundryMachine laundryMachine;
    private int count = 0;

    public LaundryMachineCounter(int capacity, int threshold) {
        this.capacity = capacity;
        this.threshold = threshold;
    }

    public void load(int[] socks) {
        if (socks.length > 0) {
            if (laundryMachine == null)
                createNewMachine();
            for (int sock : socks)
                attemptToLoad(sock);
        }
    }

    private void createNewMachine() {
        laundryMachine = new LaundryMachine(capacity, threshold);
        count++;
    }

    private void attemptToLoad(int sock) {
        try {
            laundryMachine.load(sock);
        } catch (LaundryMachineException laundryMachineException) {
            createNewMachine();
            attemptToLoad(sock);
        }
    }

    public int getCount() {
        return count;
    }
}
