mod xor;
use rustneat::Population;

pub fn neat_benchmark(){
    let mut population = Population::create_population(200);
    let mut environment = xor::XOREnv;

    let mut generation_counter = 1;

    for _ in 0..100 {
        population.evolve();
        population.evaluate_in(&mut environment);

        println!("fitness {:?}", generation_counter);

        generation_counter += 1;
    }
}

fn main() {
    neat_benchmark()
}
