use std::io::{self, Read};

fn main() {
    let mut input = String::new();
    io::stdin().read_to_string(&mut input).unwrap();
    let mut lines = input.lines();

    let t: usize = lines.next().unwrap().trim().parse().unwrap();

    for _ in 0..t {
        solve_testcase(&mut lines);
    }
}

fn solve_testcase(lines: &mut std::str::Lines) {
    let mut read_input = || -> (usize, usize, Vec<Vec<i32>>) {
        let line = lines.next().unwrap();
        let mut iter = line.split_whitespace();
        let n: usize = iter.next().unwrap().parse().unwrap();
        let m: usize = iter.next().unwrap().parse().unwrap();

        let mut grid = Vec::with_capacity(n);
        for _ in 0..n {
            let row: Vec<i32> = lines
                .next()
                .unwrap()
                .split_whitespace()
                .map(|x| x.parse().unwrap())
                .collect();
            grid.push(row);
        }

        (n, m, grid)
    };

    let (n, m, mut grid) = read_input();
    let mut ans = 0;

    for i in 0..n {
        for j in 0..m {
            if grid[i][j] != 0 {
                ans = ans.max(explore_region(&mut grid, i, j));
            }
        }
    }

    println!("{}", ans);
}

fn explore_region(grid: &mut Vec<Vec<i32>>, i: usize, j: usize) -> i32 {
    if i >= grid.len() || j >= grid[0].len() || grid[i][j] == 0 {
        return 0;
    }

    let mut size = grid[i][j];
    grid[i][j] = 0;

    size += explore_region(grid, i + 1, j);
    size += explore_region(grid, i - 1, j);
    size += explore_region(grid, i, j + 1);
    size += explore_region(grid, i, j - 1);

    size
}
