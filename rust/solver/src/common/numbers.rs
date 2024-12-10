pub fn digits(num: u32, radix: u32) -> Vec<char> {
    if num == 0 {
        return vec!['0'];
    }


    let mut digits = Vec::new();
    let mut n = num;

    loop {
        if n == 0 {
            break;
        } else {
            let digit = n % radix;
            n /= radix;
            digits.push(std::char::from_digit(digit, radix).unwrap());
        }
    }

    digits.reverse();
    digits
}

#[test]
fn test_digits() {
    // assert_eq!(digits(0, 10), vec!['0']);
    assert_eq!(digits(1, 10), vec!['1']);
    assert_eq!(digits(4, 3), vec!['1', '1']);
    assert_eq!(digits(10, 10), vec!['1', '0']);
    assert_eq!(digits(10, 2), vec!['1', '0', '1', '0']);
    assert_eq!(digits(10, 16), vec!['a']);
    assert_eq!(digits(15, 16), vec!['f']);
    assert_eq!(digits(16, 16), vec!['1', '0']);
    assert_eq!(digits(17, 16), vec!['1', '1']);
    assert_eq!(digits(255, 16), vec!['f', 'f']);
    assert_eq!(digits(256, 16), vec!['1', '0', '0']);
    assert_eq!(digits(257, 16), vec!['1', '0', '1']);
    assert_eq!(digits(1000, 16), vec!['3', 'e', '8']);
    assert_eq!(digits(1000, 10), vec!['1', '0', '0', '0']);
    assert_eq!(
        digits(1000, 2),
        vec!['1', '1', '1', '1', '1', '0', '1', '0', '0', '0']
    );
}
