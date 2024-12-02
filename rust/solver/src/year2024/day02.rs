use std::error::Error;

pub fn solve_first_star(input: &str) -> Result<String, Box<dyn Error>> {
    let result = parse_input(input)?.iter().fold(0, |acc, row| {
        acc + (if is_safe(row).is_some() { 1 } else { 0 })
    });

    Ok(format!("{}", result))
}

pub fn solve_second_star(input: &str) -> Result<String, Box<dyn Error>> {
    let result = parse_input(input)?.iter().fold(0, |acc, row| {
        let safe = mutate_report(row)
            .iter()
            .any(|mutation| is_safe(mutation).is_some());
        acc + (if safe { 1 } else { 0 })
    });

    Ok(format!("{}", result))
}

fn parse_input(input: &str) -> Result<Vec<Vec<i32>>, Box<dyn Error>> {
    input.lines().map(|line| parse_line(line)).collect()
}

fn parse_line(line: &str) -> Result<Vec<i32>, Box<dyn Error>> {
    line.split_whitespace()
        .map(|s| s.parse::<i32>().map_err(|e| e.into()))
        .collect()
}

fn is_safe(v: &Vec<i32>) -> Option<()> {
    let mut direction = 0;

    for i in 0..v.len() - 1 {
        let this = v[i];
        let next = v[i + 1];
        let diff = (this - next).abs();
        if diff > 3 || diff < 1 {
            return None;
        }

        let dir = (next - this) / diff;
        if direction == 0 {
            direction = dir;
        }
        if dir != direction {
            return None;
        }
    }

    Some(())
}

fn mutate_report(vector: &Vec<i32>) -> Vec<Vec<i32>> {
    vector
        .iter()
        .enumerate()
        .map(|(i, _)| {
            let mut mutation = vector.clone();
            mutation.remove(i);
            mutation
        })
        .collect()
}
