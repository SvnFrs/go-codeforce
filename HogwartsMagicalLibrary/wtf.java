package HogwartsMagicalLibrary;
import java.util.Scanner;
public class wtf {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        long t = scanner.nextLong();

        while (t-- > 0) {
            long n = scanner.nextLong();
            long cnt = 0;

            for (long p = 0; p < 2 * n; p++) {
                long x = scanner.nextLong();
                cnt += (x % 2);
            }
            System.out.println(cnt == n ? "Yes" : "No");
        }
    }
}