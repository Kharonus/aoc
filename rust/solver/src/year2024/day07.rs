use crate::common::numbers::digits;
use std::collections::HashMap;
use std::error::Error;

pub fn solve_first_star(input: &str) -> Result<String, Box<dyn Error>> {
    let sum: usize = parse_input(input)?
        .iter()
        .filter_map(|(test_value, numbers)| {
            let count_operators = numbers.len() as u32 - 1;

            for combination in build_operator_combinations(count_operators, vec!['+', '*']) {
                let result = numbers.iter().enumerate().fold(0, |acc, (idx, n)| {
                    if idx == 0 {
                        return *n;
                    }

                    let operator = combination[idx - 1];
                    match operator {
                        '+' => acc + n,
                        '*' => acc * n,
                        _ => panic!("Unknown operator: {}", operator),
                    }
                });

                if &result == test_value {
                    return Some(result);
                }
            }

            None
        })
        .sum();

    Ok(format!("{}", sum))
}

pub fn solve_second_star(input: &str) -> Result<String, Box<dyn Error>> {
    let sum: usize = parse_input(input)?
        .iter()
        .filter_map(|(test_value, numbers)| {
            let count_operators = numbers.len() as u32 - 1;

            for combination in build_operator_combinations(count_operators, vec!['+', '*', '|']) {
                let result = numbers.iter().enumerate().fold(0, |acc, (idx, n)| {
                    if idx == 0 {
                        return *n;
                    }

                    let operator = combination[idx - 1];
                    match operator {
                        '+' => acc + n,
                        '*' => acc * n,
                        '|' => combine(acc, *n),
                        _ => panic!("Unknown operator: {}", operator),
                    }
                });

                if &result == test_value {
                    return Some(result);
                }
            }

            None
        })
        .sum();

    Ok(format!("{}", sum))
}

fn combine(a: usize, b: usize) -> usize {
    format!("{}{}", a, b).parse().unwrap()
}

fn build_operator_combinations(
    count: u32,
    operators: Vec<char>,
) -> impl Iterator<Item = Vec<char>> {
    let radix = operators.len() as u32;
    let mut operator_map = HashMap::new();
    for (idx, c) in operators.iter().enumerate() {
        operator_map.insert(std::char::from_digit(idx as u32, radix).unwrap(), c.clone());
    }

    let mut iteration = 0u32;

    std::iter::from_fn(move || {
        if iteration >= radix.pow(count) {
            return None;
        }

        let x = prefix_with_zeros(digits(iteration, radix), count as usize)
            .iter()
            .filter_map(|d| operator_map.get(&d).copied())
            .collect();

        iteration += 1;
        Some(x)
    })
}

fn prefix_with_zeros(digits: Vec<char>, size: usize) -> Vec<char> {
    vec!['0'; size - digits.len()]
        .iter()
        .chain(digits.iter())
        .copied()
        .collect()
}

fn parse_input(input: &str) -> Result<Vec<(usize, Vec<usize>)>, Box<dyn Error>> {
    input
        .lines()
        .map(|line| {
            let result = parse_line(line);
            result
        })
        .collect()
}

fn parse_line(line: &str) -> Result<(usize, Vec<usize>), Box<dyn Error>> {
    let mut test_value = 0;
    let mut numbers = Vec::new();

    for s in line.split_whitespace() {
        if s.ends_with(":") {
            let x = s.split(":").next().unwrap_or(s);
            test_value = x.parse::<usize>()?;
        } else {
            numbers.push(s.parse::<usize>()?);
        }
    }

    Ok((test_value, numbers))
}
