use std::io::Error;

pub fn solve_first_star(input: &str) -> Result<String, Error> {
    let (mut left, mut right) = parse_input(input);

    left.sort();
    right.sort();

    let distances = left
        .iter()
        .zip(right.iter())
        .fold(0, |acc, (l, r)| acc + (l - r).abs());

    Ok(format!("{}", distances))
}

pub fn solve_second_star(input: &str) -> Result<String, Error> {
    let (left, right) = parse_input(input);

    let similarity = left.iter().fold(0, |acc, &number| {
        let x = right.iter().filter(|&&r| r == number).count() as i32;
        acc + (number * x)
    });

    Ok(format!("{}", similarity))
}

fn parse_input(input: &str) -> (Vec<i32>, Vec<i32>) {
    input.lines().fold(
        (Vec::new(), Vec::new()),
        |mut acc: (Vec<i32>, Vec<i32>), line| {
            let mut s = line.split_whitespace();

            if let Some(left) = s.next() {
                acc.0.push(left.parse().unwrap());
            }

            if let Some(right) = s.next() {
                acc.1.push(right.parse().unwrap());
            }

            acc
        },
    )
}
