use std::io::{self, BufRead};
use pcre2::bytes::Regex;

fn read_lines() -> Vec<String> {
    let stdin = io::stdin();
    let lines = stdin.lock().lines();
    lines
        .filter_map(|l| l.ok())
        .collect()
}

fn parse_input(input: String) -> Vec<String> {
    input.split(',')
        .map(|s| s.trim().to_string())
        .collect()
}

fn check_nums(input: &[String], re: Regex) -> i64 { 
    let mut sum: i64 = 0;
    for range in input {
        let splitted: Vec<&str> = range.split('-').collect();
        let start: i64 = splitted[0].parse().unwrap();
        let end: i64 = splitted[1].parse().unwrap();
        for n in start..=end {
            if re.is_match(n.to_string().as_bytes()).unwrap() {
                sum += n;
            }
        }
    }
    sum
}

fn part1(input: &[String]) {
    println!("{}", check_nums(input, Regex::new(r"^(\w+)\1$").unwrap()))    
}

fn part2(input: &[String]) {
    println!("{}", check_nums(input, Regex::new(r"^(\w+)\1+$").unwrap()))
}

fn main() {
    let lines= read_lines();
    let input = parse_input(lines[0].clone());
    part1(&input);
    part2(&input);
}
