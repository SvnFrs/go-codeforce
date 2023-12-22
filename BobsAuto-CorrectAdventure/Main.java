import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

public class Main {

    public static void findAndPrintCharactersToRemove(String first, String second) {
        int length = first.length();
        List<Integer> indexesToRemove = new ArrayList<>();
        StringBuilder sb = new StringBuilder();

        for (int i = 0; i < length; i++) {
            sb.setLength(0);
            sb.append(first, 0, i).append(first, i + 1, length);

            if (sb.toString().equals(second)) {
                indexesToRemove.add(i + 1);
            }
        }

        System.out.println(indexesToRemove.size());
        if (!indexesToRemove.isEmpty()) {
            String result = String.join(" ", indexesToRemove.stream().map(Object::toString).toArray(String[]::new));
            System.out.println(result);
        }
    }

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);

        String first, second;

        if (scanner.hasNext()) {
            first = scanner.next();
        } else {
            System.out.println("0");
            return;
        }

        if (scanner.hasNext()) {
            second = scanner.next();
        } else {
            System.out.println("0");
            return;
        }

        findAndPrintCharactersToRemove(first, second);
    }
}
