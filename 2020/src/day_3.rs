use std::error::Error;
use std::fs;

fn main() -> Result<(), Box<dyn Error>> {
    let text = fs::read_to_string("day_3_puzzle_input.txt")?;
    let mut trees: Vec<Vec<bool>> = Vec::new();

    for line in text.trim_end().split('\n') {
        let mut tree_line: Vec<bool> = Vec::new();
        for c in line.chars() {
            if c == '.' {
                tree_line.push(false);
            } else {
                tree_line.push(true);
            }
        }
        trees.push(tree_line);
    }

    let slopes_down = [1, 1, 1, 1, 2];
    let slopes_right = [1, 3, 5, 7, 1];
    let mut product: i128 = 1;

    for index in (0..slopes_down.len()).rev() {
        let mut line_number = slopes_down[index];
        let mut x = slopes_right[index];
        let mut count = 0;
        while line_number < trees.len() {
            let tree_line = &trees[line_number];
            if tree_line[x] {
                count += 1;
            }
            line_number += slopes_down[index];
            x = (x + slopes_right[index]) % tree_line.len();
        }
        product *= count;
    }
    println!("{}", product);

    Ok(())
}
