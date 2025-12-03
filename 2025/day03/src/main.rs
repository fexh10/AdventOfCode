use std::io::{self, BufRead};
use itertools::Itertools;

fn read_lines() -> Vec<String> {
    let stdin = io::stdin();
    let lines = stdin.lock().lines();
    lines.filter_map(|l| l.ok()).collect()
}

fn part1(lines: &[String]) {
    let mut sum = 0;
    for line in lines {
        let mut max = 0;
        for (a, b) in line.chars().tuple_combinations() {
            let num : i32 = format!("{}{}", a, b).parse().unwrap();
            if num > max {
                max = num
            } 
        }
        sum += max;
    }
    println!("{}", sum)   
}

fn part2(lines: &[String]) {

    fn max_number(line: &String) -> i64 {
        let chars: Vec<char> = line.chars().collect();
        let mut result = String::new();
        let mut start = 0;

        for n in (0..12).rev() {
            let end = chars.len() - n - 1;
            let mut best_char = '0';
            let mut best_pos = start;

            for i in start..=end {
                if chars[i] > best_char {
                    best_char = chars[i];
                    best_pos = i;
                }
            }
            result.push(best_char);
            start = best_pos + 1;
        }
        result.parse::<i64>().unwrap()
    }

    let mut sum: i64 = 0;
    for line in lines {
        sum += max_number(line);
    }
    println!("{}", sum) 
}

fn main() {
    let lines= read_lines();
    part1(&lines);
    part2(&lines);
}