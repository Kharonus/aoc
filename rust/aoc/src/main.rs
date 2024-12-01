mod year2024;

use std::env;
use std::fs::File;
use std::io::{self, Read};

fn main() {
    let args: Vec<String> = env::args().collect();
    println!("Command line arguments: {:?}", args);

    match read_file_to_string("input/2024/d01.txt") {
        Ok(contents) => {
            // let result = year2024::day1::solve_first_star(&contents);
            let result = year2024::day1::solve_second_star(&contents);
            println!("Result: {}", result);
        }
        Err(e) => println!("Error reading file: {}", e),
    }
}

fn read_file_to_string(file_path: &str) -> io::Result<String> {
    let mut file = File::open(file_path)?;
    let mut contents = String::new();
    file.read_to_string(&mut contents)?;
    Ok(contents)
}
