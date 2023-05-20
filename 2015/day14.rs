use regex::Regex;
use std::collections::HashMap;

#[derive(Clone)]
struct Reindeer {
    name: String,
    speed: u64,
    run_duration: u64,
    rest_duration: u64,
}

impl Reindeer {
    fn distance_after(&self, secs: u64) -> u64 {
        let cycle_distance = self.speed * self.run_duration;
        let cycle_duration = self.run_duration + self.rest_duration;
        let cycle_count = secs / cycle_duration;
        let prev_cycles_distance = cycle_distance * cycle_count;
        let remaining_secs = std::cmp::min(secs % cycle_duration, self.run_duration);
        prev_cycles_distance + remaining_secs * self.speed
    }
}

#[aoc_generator(day14)]
fn parse_input(input: &str) -> Vec<Reindeer> {
    let mut deers: Vec<Reindeer> = Vec::new();
    let re = Regex::new(
        r"(\w+) can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds.",
    )
    .unwrap();
    for l in input.lines() {
        if let Some(matches) = re.captures(l) {
            let name: String = matches.get(1).unwrap().as_str().to_string();
            let speed: u64 = matches.get(2).unwrap().as_str().parse().expect("uint64");
            let run_duration: u64 = matches.get(3).unwrap().as_str().parse().expect("uint64");
            let rest_duration: u64 = matches.get(4).unwrap().as_str().parse().expect("uint64");
            deers.push(Reindeer {
                name,
                speed,
                run_duration,
                rest_duration,
            });
        } else {
            panic!("NO");
        }
    }
    return deers.to_vec();
}

fn best_distance(deers: &[Reindeer], time: u64) -> u64 {
    deers
        .iter()
        .map(|deer| deer.distance_after(time))
        .max()
        .unwrap()
}

/// Part 1
fn part1_(input: &[Reindeer], time: u64) -> u64 {
    best_distance(input, time)
}

#[aoc(day14, part1)]
fn part1(input: &[Reindeer]) -> u64 {
    part1_(input, 2503)
}

/// Part 2
fn part2_(deers: &[Reindeer], time: u64) -> u64 {
    let mut points: HashMap<String, u64> = HashMap::new();
    for deer in deers {
        points.insert(deer.name.clone(), 0);
    }
    for t in 1..=time {
        let best_dist = best_distance(&deers, t);
        for deer in deers {
            if deer.distance_after(t) == best_dist {
                *points.get_mut(&deer.name).unwrap() += 1;
            }
        }
    }
    *points.values().max().unwrap()
}

#[aoc(day14, part2)]
fn part2(deers: &[Reindeer]) -> u64 {
    part2_(deers, 2503)
}

#[cfg(test)]
mod tests {
    use super::*;

    const EXAMPLE: &str = "Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.\nDancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.";

    #[test]
    fn part1_examples() {
        assert_eq!(1120, part1_(&parse_input(EXAMPLE), 1000));
    }

    #[test]
    fn part2_examples() {
        assert_eq!(689, part2_(&parse_input(EXAMPLE), 1000));
    }
}
