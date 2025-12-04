use std::io::{self, BufRead};

fn read_lines() -> Vec<String> {
    let stdin = io::stdin();
    let lines = stdin.lock().lines();
    lines.filter_map(|l| l.ok()).collect()
}

fn count_adjacent(lines: &[String], i: usize, j: usize) -> i32 {
        let directions: [[i32; 2]; 8] = [[-1, -1], [-1, 0], [-1, 1], [0, -1], [0, 1], [1, -1], [1, 0], [1, 1]]; 
        let mut adjacents = 0;
        for [di, dj] in directions {
            let new_i = i as i32 + di;
            let new_j = j as i32 + dj;
            if new_i >= 0 && new_i < lines.len() as i32 && new_j >= 0 
                && lines.get(new_i as usize)
                    .and_then(|line| line.chars().nth(new_j as usize))
                    .map_or(false, |c| c == '@') {
                adjacents += 1;
            }
        }
        adjacents
    }

fn part1(lines: &[String]) {
    let mut sum = 0;
    for (i, line) in lines.iter().enumerate() {
        for (j, char) in line.chars().enumerate() {
            if char == '@' && count_adjacent(lines, i, j) < 4{ 
                sum += 1;
            }
        }
    }
    println!("{}", sum)
}

fn part2(lines: &[String]) {
    fn remove(lines: &[String], sum: i32) -> i32 {
        let mut sum1 = 0;
        let mut removed = false;
        let mut new_lines = lines.to_vec();

        for (i, line) in lines.iter().enumerate() {
            for (j, char) in line.chars().enumerate() {
                if char == '@' && count_adjacent(lines, i, j) < 4{ 
                    sum1 += 1;
                    new_lines[i].replace_range(j..j+1, "x");
                    removed = true;
                }
            }
        }
        if removed {
            remove(&new_lines, sum + sum1)
        } else {
            sum
        }
    }
    let sum = remove(lines, 0);
    println!("{}", sum)
}

fn main() {
    let lines= read_lines();
    part1(&lines);
    part2(&lines);
}