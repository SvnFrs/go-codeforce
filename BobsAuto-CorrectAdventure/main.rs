use std::io::{self, BufRead};

fn find_and_print_characters_to_remove(first: &str, second: &str) {
    let length = first.len();
    let mut indexes_to_remove = Vec::new();

    for i in 0..length {
        let mut removed_string = String::with_capacity(length - 1);
        removed_string.push_str(&first[..i]);
        removed_string.push_str(&first[i + 1..]);

        if removed_string == second {
            indexes_to_remove.push(i + 1);
        }
    }

    println!("{}", indexes_to_remove.len());
    if !indexes_to_remove.is_empty() {
        let result_str: String = indexes_to_remove
            .iter()
            .map(|&i| i.to_string() + " ")
            .collect();
        println!("{}", result_str);
    }
}
// yen tam, code bang tay nha
fn main() {
    let mut input_line = String::new();

    if io::stdin().read_line(&mut input_line).is_err() {
        println!("0");
        return;
    }

    let mut first = input_line.trim().to_string();
    input_line.clear();

    if io::stdin().read_line(&mut input_line).is_err() {
        println!("0");
        return;
    }

    let second = input_line.trim();

    first.shrink_to_fit();

    find_and_print_characters_to_remove(&first, second);
}
