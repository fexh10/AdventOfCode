use std::{io::{self, BufRead}};

fn read_lines() -> Vec<String> {
    let stdin = io::stdin();
    let lines = stdin.lock().lines();
    lines.filter_map(|l| l.ok()).collect()
}

fn part1(lines: &[String]) {
    let mut numbers: Vec<i64> = lines[0].split_whitespace().filter_map(|s| s.trim().parse::<i64>().ok()).collect();
    let operations: Vec<&str> = lines.last().unwrap().split_whitespace().collect();

    for line in &lines[1..lines.len().saturating_sub(1)] {
        let partial: Vec<i64> = line.split_whitespace().filter_map(|s| s.trim().parse::<i64>().ok()).collect();
        for (i, &value) in partial.iter().enumerate() {
            match operations[i] {
                "*" => numbers[i] *= value,
                "+" => numbers[i] += value,
                _ => {}
            }
        }
    }
    let sum: i64 = numbers.iter().sum();
    println!("{}", sum)
}

fn part2(lines: &[String]) {
    fn parse_numbers_with_positions(line: &str) -> Vec<(String, usize)> {
        let mut results = Vec::new();
        let mut current_num = String::new();
        let mut start_pos = 0;
        
        for (i, ch) in line.chars().enumerate() {
            if ch.is_whitespace() {
                if !current_num.is_empty() {
                    results.push((current_num.clone(), start_pos));
                    current_num.clear();
                }
            } else {
                if current_num.is_empty() {
                    start_pos = i;
                }
                current_num.push(ch);
            }
        }
        
        if !current_num.is_empty() {
            results.push((current_num, start_pos));
        }
        
        results
    }
    
    let operations: Vec<&str> = lines.last().unwrap().split_whitespace().collect();
    let mut col_positions: Vec<Vec<usize>> = vec![Vec::new(); operations.len()];
    
    for line in &lines[..lines.len().saturating_sub(1)] {
        let nums_with_pos = parse_numbers_with_positions(line);
        for (col_idx, (_, pos)) in nums_with_pos.iter().enumerate() {
            if col_idx < col_positions.len() {
                col_positions[col_idx].push(*pos);
            }
        }
    }
    
    let mut col_alignments = Vec::new();
    for positions in &col_positions {
        if positions.is_empty() {
            col_alignments.push(false);
            continue;
        }
       
        let first = positions[0];
        let all_same = positions.iter().all(|&p| p == first);
        col_alignments.push(!all_same);
    }
    
    let mut colums: Vec<Vec<String>> = Vec::new();

    for line in &lines[..lines.len().saturating_sub(1)] {
        let splitted: Vec<String> = line.split_whitespace().map(|s| s.to_string()).collect();

        for (i, string_number) in splitted.iter().enumerate() {
            if colums.len() <= i {
                colums.resize(i + 1, Vec::new());
            }
            let new_string_number = if col_alignments[i] { 
                format!("{:>4}", string_number) 
            } else { 
                format!("{:<4}", string_number) 
            };
            for (j, char) in new_string_number.chars().enumerate() {
                if colums[i].len() <= j {
                    colums[i].resize(j + 1, String::new());
                }
                colums[i][j].push(char);
            }
        }
    }
    
    let mut sum = 0;
    for (i, string_numbers) in colums.iter().enumerate() {
        let numbers: Vec<i64> = string_numbers.iter()
            .filter(|s| !s.trim().is_empty())
            .filter_map(|s| s.trim().parse::<i64>().ok())
            .collect();
        
        if numbers.is_empty() { continue; }
        
        match operations[i] {
            "+" => sum += numbers.iter().sum::<i64>(),
            "*" => sum += numbers.iter().product::<i64>(),
            _ => {},
        }
    }
    println!("{}", sum)
}

fn main() {
    let lines= read_lines();
    part1(&lines);
    part2(&lines);
}