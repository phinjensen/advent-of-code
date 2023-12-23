use std::error::Error;
use std::fs;

fn main() -> Result<(), Box<dyn Error>> {
    let text = fs::read_to_string("day_2_puzzle_input.txt")?;

    let mut count = 0;

    for line in text.trim_end().split('\n') {
        let mut chunks = line.split(' ');
        let range: Vec<i8>;
        if let Some(range_chunk) = chunks.next() {
            range = range_chunk.split('-').map(|x| x.parse().unwrap()).collect();
        } else {
            panic!("Didnt find a range!");
        }
        let letter = chunks.next().unwrap().strip_suffix(':').unwrap().chars().next().unwrap();
        let password = chunks.next().unwrap();
        let first = password.chars().nth((range[0]-1) as usize).unwrap() == letter;
        let second = password.chars().nth((range[1]-1) as usize).unwrap() == letter;

        if first {
            if !second {
                count += 1;
            }
        } else {
            if second {
                count += 1;
            }
        }
    }

    println!("{}", count);

    Ok(())
}
