use std::error::Error;

type Position = (usize, usize);
type Direction = (isize, isize);

pub fn solve_first_star(input: &str) -> Result<String, Box<dyn Error>> {
    let map = parse_input(input)?;

    let mut visited: Vec<Position> = Vec::new();
    let mut dir = map.start_direction;
    let mut pos = map.start;

    while let Some(is_obstacle) = ahead(&map, pos, dir) {
        if is_obstacle {
            dir = turn_right(dir);
        } else {
            if !visited.contains(&pos) {
                visited.push(pos);
            }
            pos = move_forward(pos, dir)?;
        }
    }

    // Add last position before stepping out the map
    visited.push(pos);

    Ok(format!("{}", visited.len()))
}

pub fn solve_second_star(input: &str) -> Result<String, Box<dyn Error>> {
    let map = parse_input(input)?;

    let mut looping_obstacles: Vec<Position> = Vec::new();
    let mut dir = map.start_direction;
    let mut pos = map.start;

    while let Some(is_obstacle) = ahead(&map, pos, dir) {
        if is_obstacle {
            dir = turn_right(dir);
        } else {
            let forward = move_forward(pos, dir)?;
            let manipulated_map = Map {
                obstacles: add_obstacle(&map.obstacles, forward),
                start: map.start,
                start_direction: map.start_direction,
            };

            if move_guard_without_loop(&manipulated_map).is_none()
                && !looping_obstacles.contains(&forward)
            {
                looping_obstacles.push(forward);
            }

            pos = forward;
        }
    }

    Ok(format!("{}", looping_obstacles.len()))
}

#[derive(Debug)]
struct Map {
    obstacles: Vec<Vec<bool>>,
    start: Position,
    start_direction: Direction,
}

fn move_guard_without_loop(map: &Map) -> Option<()> {
    let mut visited: Vec<(Position, Direction)> = Vec::new();
    let mut dir = map.start_direction;
    let mut pos = map.start;

    while let Some(is_obstacle) = ahead(map, pos, dir) {
        if is_obstacle {
            dir = turn_right(dir);
        } else {
            if !visited.contains(&(pos, dir)) {
                visited.push((pos, dir));
                pos = move_forward(pos, dir).ok()?;
            } else {
                return None;
            }
        }
    }

    Some(())
}

fn add_obstacle(obstacles: &Vec<Vec<bool>>, pos: Position) -> Vec<Vec<bool>> {
    let mut clone = obstacles.clone();
    clone[pos.1][pos.0] = true;
    clone
}

fn ahead(map: &Map, pos: Position, direction: Direction) -> Option<bool> {
    let x = pos.0 as isize + direction.0;
    let y = pos.1 as isize + direction.1;

    if x < 0 || y < 0 {
        return None;
    }

    map.obstacles
        .get(y as usize)
        .and_then(|row| row.get(x as usize).copied())
}

fn turn_right(direction: Direction) -> Direction {
    match direction {
        (0, 1) => (-1, 0),
        (1, 0) => (0, 1),
        (0, -1) => (1, 0),
        (-1, 0) => (0, -1),
        _ => panic!("Invalid direction"),
    }
}

fn move_forward(pos: Position, direction: Direction) -> Result<Position, Box<dyn Error>> {
    pos.0
        .checked_add_signed(direction.0)
        .and_then(|x| pos.1.checked_add_signed(direction.1).map(|y| (x, y)))
        .ok_or_else(|| "Invalid position".into())
}

fn parse_input(input: &str) -> Result<Map, Box<dyn Error>> {
    let mut map = Map {
        obstacles: Vec::new(),
        start: (0, 0),
        start_direction: (0, -1),
    };

    for (idx, line) in input.lines().enumerate() {
        let (obstacles, start) = parse_line(line)?;

        if let Some(pos) = start {
            map.start = (pos, idx);
        }

        map.obstacles.push(obstacles);
    }

    Ok(map)
}

fn parse_line(line: &str) -> Result<(Vec<bool>, Option<usize>), Box<dyn Error>> {
    let mut obstacles: Vec<bool> = Vec::new();
    let mut start: Option<usize> = None;

    for (idx, c) in line.chars().enumerate() {
        match c {
            '.' => obstacles.push(false),
            '#' => obstacles.push(true),
            '^' => {
                obstacles.push(false);
                start = Some(idx);
            }
            _ => return Err("Invalid character".into()),
        }
    }

    Ok((obstacles, start))
}
