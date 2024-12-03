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
                assert_eq!(result, "1882714");
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
                assert_eq!(result, "19437052");
            }
            Err(error) => {
                panic!("Expected success but got an error: {:?}", error);
            }
        }
    }
}

mod day02 {
    use super::common;
    use solver::year2024::day02;

    #[test]
    fn first_star() {
        match common::setup("input/2024/d02.txt").and_then(|input| day02::solve_first_star(&input))
        {
            Ok(result) => {
                println!("Result day 02 star 1: {}", result);
                assert_eq!(result, "534");
            }
            Err(error) => {
                panic!("Expected success but got an error: {:?}", error);
            }
        }
    }

    #[test]
    fn second_star() {
        match common::setup("input/2024/d02.txt").and_then(|input| day02::solve_second_star(&input))
        {
            Ok(result) => {
                println!("Result day 02 star 2: {}", result);
                assert_eq!(result, "577");
            }
            Err(error) => {
                panic!("Expected success but got an error: {:?}", error);
            }
        }
    }
}mod day03 {
    use super::common;
    use solver::year2024::day03;

    #[test]
    fn first_star() {
        match common::setup("input/2024/d03.txt").and_then(|input| day03::solve_first_star(&input))
        {
            Ok(result) => {
                println!("Result day 03 star 1: {}", result);
                assert_eq!(result, "188192787");
            }
            Err(error) => {
                panic!("Expected success but got an error: {:?}", error);
            }
        }
    }

    #[test]
    fn second_star() {
        match common::setup("input/2024/d03.txt").and_then(|input| day03::solve_second_star(&input))
        {
            Ok(result) => {
                println!("Result day 03 star 2: {}", result);
                assert_eq!(result, "113965544");
            }
            Err(error) => {
                panic!("Expected success but got an error: {:?}", error);
            }
        }
    }
}