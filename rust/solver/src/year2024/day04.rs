use std::error::Error;

pub fn solve_first_star(input: &str) -> Result<String, Box<dyn Error>> {
    let map = parse_input(input);
    let mut sum: usize = 0;

    for i in 0..map.len() {
        for j in 0..map[i].len() {
            for direction in Direction::iter() {
                if is_word(&map, j as i32, i as i32, &direction) {
                    sum += 1;
                }
            }
        }
    }

    Ok(format!("{}", sum))
}

pub fn solve_second_star(input: &str) -> Result<String, Box<dyn Error>> {
    let map = parse_input(input);
    let mut sum: usize = 0;

    for i in 0..map.len() {
        for j in 0..map[i].len() {
            if is_cross(&map, j as i32, i as i32).is_some() {
                sum += 1;
            }
        }
    }

    Ok(format!("{}", sum))
}

fn parse_input(input: &str) -> Vec<Vec<char>> {
    input.lines().map(|line| line.chars().collect()).collect()
}

fn is_word(map: &Vec<Vec<char>>, start_x: i32, start_y: i32, direction: &Direction) -> bool {
    let mut x = start_x;
    let mut y = start_y;

    for c in "XMAS".chars() {
        let character = get_char(map, x, y);

        if character.is_none() || character.unwrap() != c {
            return false;
        }

        (x, y) = mutate_coords(x, y, direction);
    }

    true
}

fn is_cross(map: &Vec<Vec<char>>, x: i32, y: i32) -> Option<()> {
    if 'A' != get_char(map, x, y)? {
        return None;
    }

    let opposite_corners = vec![
        (
            mutate_coords(x, y, &Direction::UpRight),
            mutate_coords(x, y, &Direction::DownLeft),
        ),
        (
            mutate_coords(x, y, &Direction::UpLeft),
            mutate_coords(x, y, &Direction::DownRight),
        ),
    ];

    for corners in opposite_corners {
        let c1 = get_char(map, corners.0.0, corners.0.1)?;
        let c2 = get_char(map, corners.1.0, corners.1.1)?;

        let comparison = format!("{}{}", c1, c2);
        if comparison != "MS" && comparison != "SM" {
            return None;
        }
    }

    Some(())
}

fn mutate_coords(x: i32, y: i32, direction: &Direction) -> (i32, i32) {
    match direction {
        Direction::Up => (x, y - 1),
        Direction::Down => (x, y + 1),
        Direction::Left => (x - 1, y),
        Direction::Right => (x + 1, y),
        Direction::UpLeft => (x - 1, y - 1),
        Direction::UpRight => (x + 1, y - 1),
        Direction::DownLeft => (x - 1, y + 1),
        Direction::DownRight => (x + 1, y + 1),
    }
}

fn get_char(map: &Vec<Vec<char>>, x: i32, y: i32) -> Option<char> {
    if x < 0 || y < 0 {
        return None;
    }

    map.get(y as usize)
        .and_then(|row| row.get(x as usize).copied())
}

enum Direction {
    Up,
    Down,
    Left,
    Right,
    UpLeft,
    UpRight,
    DownLeft,
    DownRight,
}

impl Direction {
    fn iter() -> impl Iterator<Item=Direction> {
        use Direction::*;
        vec![Up, Down, Left, Right, UpLeft, UpRight, DownLeft, DownRight].into_iter()
    }
}
