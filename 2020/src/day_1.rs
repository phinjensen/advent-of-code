use std::fs::File;
use std::io;
use std::io::prelude::*;

fn main() -> io::Result<()> {
    let file = File::open("day_1_puzzle_1_input.txt")?;
    let mut values: Vec<i32> = Vec::new();

    for line in io::BufReader::new(file).lines() {
        let num: i32 = line.unwrap().parse().unwrap();
        values.push(num);
    }

    'outer: for i in &values {
        for j in &values {
            if i + j < 2020 {
                for k in &values {
                    if i + j + k == 2020 {
                        println!("{} + {} + {} = {}", i, j, k, 2020);
                        println!("{} * {} * {} = {}", i, j, k, i * j * k);
                        break 'outer;
                    }
                }
            }
        }
    }

    Ok(())
}

//fn main() -> io::Result<()> {
//    let f = File::open("day_1_puzzle_1_input.txt")?;
//    let mut reader = BufReader::new(f);
//
//    let mut line = String::new();
//    Ok(())
//}
