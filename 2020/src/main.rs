use std::collections::HashMap;
use std::error::Error;
use std::fs;

fn main() -> Result<(), Box<dyn Error>> {
    let text = fs::read_to_string("day_6.txt")?;

    let mut count = 0;
    for group in text.trim_end().split("\n\n") {
        let mut counts = HashMap::new();
        let mut people_count = 0;

        for person in group.split("\n") {
            people_count += 1;
            for letter in person.chars() {
                if let Some(&mut x) = counts.get_mut(&letter) {
                    counts.insert(letter, x + 1);
                } else {
                    counts.insert(letter, 1);
                }
            }
        }

        for occurences in counts.values() {
            if *occurences == people_count {
                count += 1;
            }
        }
    }

    println!("{}", count);

    Ok(())
}
