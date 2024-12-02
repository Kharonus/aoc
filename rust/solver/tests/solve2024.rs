mod common;

mod day01 {
    use super::common;
    use solver::year2024::day01;

    #[test]
    fn first_star() {
        match common::setup("input/2024/d01.txt").and_then(|input| day01::solve_first_star(&input))
        {
            Ok(result) => {
                println!("Result day 01 star 1: {}", result);
                assert_eq!(result.is_empty(), false);
            }
            Err(error) => {
                panic!("Expected success but got an error: {:?}", error);
            }
        }
    }

    #[test]
    fn second_star() {
        match common::setup("input/2024/d01.txt").and_then(|input| day01::solve_second_star(&input))
        {
            Ok(result) => {
                println!("Result day 01 star 2: {}", result);
                assert_eq!(result.is_empty(), false);
            }
            Err(error) => {
                panic!("Expected success but got an error: {:?}", error);
            }
        }
    }
}
