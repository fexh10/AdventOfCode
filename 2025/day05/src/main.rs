use std::{io::{self, BufRead}};

fn read_lines() -> (Vec<String>, Vec<u64>) {
    let stdin = io::stdin();
    let lines: Vec<String> = stdin.lock().lines().filter_map(|l| l.ok()).collect();
    
    let mut fresh_ingredients = Vec::new();
    let mut available_ingredients = Vec::new();
    let mut found_separator = false;
    
    for line in lines {
        if line.trim().is_empty() {
            found_separator = true;
            continue;
        }
        
        if found_separator {
            let num: u64 = line.parse().unwrap();
            available_ingredients.push(num);
        } else {
            fresh_ingredients.push(line);
        }
    }
    
    (fresh_ingredients, available_ingredients)
}

fn part1(fresh_ingredients: &[String], available_ingredients: &[u64]) {
    let mut sum = 0;
    for n in available_ingredients {
        for ingr in fresh_ingredients {
            let parts: Vec<&str> = ingr.split('-').collect();
            let start: u64 = parts[0].parse().unwrap();
            let end: u64 = parts[1].parse().unwrap();
            if n >= &start && n <= &end {
                sum += 1;
                break;
            }
        }
    }
    println!("{}", sum)
}

fn part2(fresh_ingredients: &[String]) {
    let mut ranges: Vec<(u64, u64)> = Vec::new();
    for ingr in fresh_ingredients {
        let parts: Vec<&str> = ingr.split('-').collect();
        let start: u64 = parts[0].parse().unwrap();
        let end: u64 = parts[1].parse().unwrap();
        ranges.push((start, end));
    }
    ranges.sort_by_key(|r| r.0);
    
    let mut merged: Vec<(u64, u64)> = Vec::new();
    for (start, end) in ranges {
        if let Some(last) = merged.last_mut() {
            if start <= last.1 {
                last.1 = last.1.max(end);
            } else {
                merged.push((start, end));
            }
        } else {
            merged.push((start, end));
        }
    }
    let sum: u64 = merged.iter().map(|(start, end)| end - start + 1).sum();
    println!("{}", sum)
}

fn main() {
    let (fresh_ingredients, available_ingredients) = read_lines();
    part1(&fresh_ingredients, &available_ingredients);
    part2(&fresh_ingredients);
}
