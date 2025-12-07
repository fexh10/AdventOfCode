use std::{collections::{HashSet, VecDeque}, io::{self, BufRead}};

fn read_lines() -> Vec<String> {
    let stdin = io::stdin();
    let lines = stdin.lock().lines();
    lines.filter_map(|l| l.ok()).collect()
}

fn part1(lines: &[String]) {
    fn count_splits(lines: &[String], s: &[(usize, usize)], split_positions: &mut HashSet<(usize, usize)>) -> usize {
        match s {
            [] => return split_positions.len(),
            _ => {
                let mut new_s: Vec<(usize, usize)> = Vec::new();
                let mut deque: VecDeque<(usize, usize)> = s.to_vec().into();
                let rows: usize = lines.len() - 1;
                let cols: usize = lines[0].len() - 1;

                let process_beam = |pos: (usize, usize), new_s: &mut Vec<(usize, usize)>, split_positions: &mut HashSet<(usize, usize)>| {
                    let new_row = pos.0 + 1;
                    if new_row <= rows {
                        if lines[new_row].as_bytes()[pos.1] == b'^' {
                            split_positions.insert((new_row, pos.1));
                            let pos1 = if pos.1 > 0 {
                                Some((new_row, pos.1 - 1))
                            } else {
                                None
                            };
                            let pos2 = if pos.1 < cols {
                                Some((new_row, pos.1 + 1))
                            } else {
                                None
                            };
                            if let Some(p) = pos1 {
                                if !new_s.contains(&p) { new_s.push(p) }
                            }
                            if let Some(p) = pos2 {
                                if !new_s.contains(&p) { new_s.push(p) }
                            }
                        } else {
                            new_s.push((new_row, pos.1));
                        }
                    }
                };
                
                while !deque.is_empty() {
                    process_beam(deque.pop_front().unwrap(), &mut new_s, split_positions);
                    if !deque.is_empty() {
                        process_beam(deque.pop_front().unwrap(), &mut new_s, split_positions);
                    }
                }
                count_splits(lines, &new_s, split_positions)
            }
        }
    }

    let s: Vec<(usize, usize)> = vec![(0, lines[0].find("S").unwrap())];
    let mut split_positions = HashSet::new();
    println!("{}", count_splits(&lines[1..], &s, &mut split_positions));
}

/*fn part2(lines: &[String]) {
    
}*/

fn main() {
    let lines= read_lines();
    part1(&lines);
    //part2(&lines);
}