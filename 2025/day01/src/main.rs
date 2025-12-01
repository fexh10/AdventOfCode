use std::io::{self, BufRead};

fn read_lines() -> Vec<String> {
    let stdin = io::stdin();
    let lines = stdin.lock().lines();
    lines
        .filter_map(|l| l.ok())
        .collect()
}

fn part1(input: &[String]) {
    let numbers: Vec<u32> = (0..100).collect();        
    let len = numbers.len() as i32;
    let mut start = 50;
    let mut result = 0;
    
    for line in input {
        let dir = line.chars().next().unwrap();
        let mut steps: i32 = line[1..].parse().unwrap();
        if dir == 'L' { steps = -steps; }
        start += steps;
        start = ((start % len) + len) % len;
        if numbers[start as usize] == 0 {
            result += 1
        }
    }
    println!("{}", result);
}

fn part2(input: &[String]) {
    let len = 100;
    let mut start = 50;
    let mut result = 0;

    for line in input {
        let dir = line.chars().next().unwrap();
        let mut steps: i32 = line[1..].parse().unwrap();
        if dir == 'L' { steps = -steps; }

        let mut pos = start;
        for _ in 0..steps.abs() {
            pos = (pos + steps.signum()) % len;
            if pos == 0 { result += 1; }
        }
        start += steps;
    }
    println!("{}", result);
}


fn main() {
    let lines= read_lines();
    part1(&lines);
    part2(&lines);
}
