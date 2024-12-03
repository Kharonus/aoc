extern crate regex;
use regex::Regex;
use std::error::Error;

pub fn solve_first_star(input: &str) -> Result<String, Box<dyn Error>> {
    let result = parse_input_without_do(input)?
        .iter()
        .fold(0, |acc, tuple| acc + tuple.0 * tuple.1);

    Ok(format!("{}", result))
}

pub fn solve_second_star(input: &str) -> Result<String, Box<dyn Error>> {
    let result = parse_input_with_do(input)?
        .iter()
        .fold(0, |acc, tuple| {
            acc + tuple.0 * tuple.1 * tuple.2
        });

    Ok(format!("{}", result))
}

fn parse_input_without_do(input: &str) -> Result<Vec<(i32, i32)>, Box<dyn Error>> {
    Ok(Regex::new(r"mul\((\d{1,3}),(\d{1,3})\)")?
        .captures_iter(input)
        .map(|cap| {
            (
                cap[1].parse::<i32>().unwrap(),
                cap[2].parse::<i32>().unwrap(),
            )
        })
        .collect())
}

fn parse_input_with_do(input: &str) -> Result<Vec<(i32, i32, i32)>, Box<dyn Error>> {
    let mut factor = 1;

    Ok(
        Regex::new(r"mul\((\d{1,3}),(\d{1,3})\)|(do\(\)|don't\(\))")?
            .captures_iter(input)
            .map(|cap| {
                match cap.get(0).unwrap().as_str() {
                    "do()" => {
                        factor = 1;
                        (0, 0, 0)
                    }
                    "don't()" => {
                        factor = 0;
                        (0, 0, 0)
                    }
                    _ => (
                        factor,
                        cap[1].parse::<i32>().unwrap(),
                        cap[2].parse::<i32>().unwrap(),
                    ),
                }
            })
            .collect(),
    )
}
