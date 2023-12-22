use std::io;

fn main() {
    let mut input = String::new();
    io::stdin().read_line(&mut input).expect("Failed to read input");
    let test_cases: i64 = input.trim().parse().expect("Invalid input");

    for _ in 0..test_cases {
        input.clear();
        io::stdin().read_line(&mut input).expect("Failed to read input");
        let mut iter = input.split_whitespace();
        let computers: i64 = iter.next().expect("Expected computers").parse().expect("Invalid input");
        let cables: i64 = iter.next().expect("Expected cables").parse().expect("Invalid input");

        let mut current_computer = 1;
        let mut hours = 0;

        while current_computer < computers && current_computer <= cables {
            current_computer *= 2;
            hours += 1;
        }

        if current_computer < computers {
            hours += (computers - current_computer + cables - 1) / cables;
        }

        println!("{}", hours);
    }
}
