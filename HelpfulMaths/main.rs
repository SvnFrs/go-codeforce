use std::io;

fn main() {
    let mut addition = String::new();
    io::stdin().read_line(&mut addition).expect("Failed to read line");
    let addition = addition.trim();

    if addition.len() < 2 {
        println!("{}", addition);
        return;
    }

    let parts: Vec<&str> = addition.split('+').collect();
    let mut numbers: Vec<i32> = Vec::with_capacity(parts.len());

    for part in parts {
        if let Ok(num) = part.parse::<i32>() {
            numbers.push(num);
        } else {
            eprintln!("Error converting {} to integer", part);
            return;
        }
    }

    numbers.sort();
    let final_str = numbers.iter().map(|&num| num.to_string()).collect::<Vec<String>>().join("+");
    println!("{}", final_str);
}
