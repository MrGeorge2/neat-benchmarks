
use rustneat::Environment;


pub struct XOREnv;

impl Environment for XOREnv {
    fn test(&self, organism: &mut rustneat::Organism) -> f64 {
        let mut output = vec![0f64];
        let mut distance: f64;

        organism.activate(&vec![0f64, 0f64], &mut output);
        distance = (0f64 - output[0]).abs();
        organism.activate(&vec![0f64, 1f64], &mut output);
        distance += (1f64 - output[0]).abs();
        organism.activate(&vec![1f64, 0f64], &mut output);
        distance += (1f64 - output[0]).abs();
        organism.activate(&vec![1f64, 1f64], &mut output);
        distance += (0f64 - output[0]).abs();
        (4f64 - distance).powi(2)
    }
}