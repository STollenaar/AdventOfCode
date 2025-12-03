package day3;

import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;

public class Day3 {
    public static void main(String[] args) {
        long part1 = findPairs(2);
        System.out.println("Part1: " + part1);

        long part2 = findPairs(12);
        System.out.println("Part2: " + part2);
    }

    public static long findPairs(int X) {
        long total = 0L;
        try (BufferedReader reader = new BufferedReader(new FileReader("2025/day3/input.txt"))) {
            String line;
            while ((line = reader.readLine()) != null) {
                if (line.trim().isEmpty()) {
                    continue;
                }

                String digits = line.replaceAll("\\D", "");
                int n = digits.length();
                if (n < X) {
                    continue;
                }

                int removals = n - X; // how many digits we may drop
                StringBuilder stack = new StringBuilder();
                for (int i = 0; i < n; i++) {
                    char c = digits.charAt(i);
                    while (stack.length() > 0 && stack.charAt(stack.length() - 1) < c && removals > 0) {
                        stack.deleteCharAt(stack.length() - 1);
                        removals--;
                    }
                    stack.append(c);
                }

                if (removals > 0) {
                    stack.setLength(stack.length() - removals);
                }

                String best = stack.substring(0, X);
                total += Long.parseLong(best);
            }
        } catch (IOException e) {
            System.err.println("Error reading file: " + e.getMessage());
        }
        return total;
    }
}
