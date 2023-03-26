use std::fs::read_to_string;

fn main() {
    let raw_input = read_to_string("input.txt").unwrap();
    let raw_input: Vec<u32> = raw_input
        .trim()
        .split("\n\n")
        .map(|x| {
            x.split("\n")
                .fold(0, |acc, x| acc + x.parse::<u32>().unwrap())
        })
        .collect();

    let raw_input = raw_input.iter().max().unwrap();

    println!("{:?}", raw_input);
}
