// Build and run with `cargo run`
// Try to build the program before add synchrinzation.

use std::thread;
use std::sync::{Arc, Mutex};

fn main() {
    // TODO: Find out what Arc is and why it is needed?
    // TODO: You need to add a Mutex, should it be Arc<Mutex<i32>> or Mutex<Arc<i32>>?
    let i = Arc::new(Mutex::new(0));

    let i_incrementing = i.clone();
    let i_decrementing = i.clone();

    let join_incrementing = thread::spawn(move || {
        for _ in 0..1_000_000 {
            // TODO: aquire the lock before using i
            let mut i_incrementing = i_incrementing.lock().unwrap();
            *i_incrementing += 1;
            // Do you have to release the mutex here?
   
        }
    });
    
    let join_decrementing = thread::spawn(move || {
        for _ in 0..1_000_000 {
            // TODO: aquire the lock before using i
            let mut i_decrementing = i_decrementing.lock().unwrap();
            *i_decrementing -= 1;
            // Do you have to release the mutex here?
        }

    });

    join_incrementing.join().unwrap();
    join_decrementing.join().unwrap();

    // TODO: aquire the lock before using i
    let i = i.lock().unwrap();
    println!("The number is: {}", *i);
}

