use std::error::Error;
use std::fs;

fn main() -> Result<(), Box<dyn Error>> {
    let text = fs::read_to_string("day_4.txt")?;

    let fields = ["byl", "cid", "ecl", "eyr", "hcl", "hgt", "iyr", "pid"];

    for line in text.trim_end().split("\n\n") {
        let chunks: Vec<&str> = line.split(' ').map(|x| &x[..3]).collect();
        chunks.sort();
        println!("{:?}", chunks);
        if chunks.len() >= 7 {
            
        }
    }

    Ok(())
}
