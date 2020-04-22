fn case(i: usize, mut r: usize, s: usize) {
    let mut res: Vec<(usize, usize)> = Vec::new();

    while r >= 2 {
        let mut cs = s - 1;
        while cs > 0 {
            let a = r * (s - 1) - (s - 1 - cs);
            let b = r - 1;
            res.push((a, b));
            cs -= 1;
        }
        r -= 1;
    }

    println!("Case #{}: {}", i, res.len());
    for l in res {
        println!("{} {}", l.0, l.1);
    }
}

fn main() {
    let t: usize = read_ints()[0];
    for i in 0..t {
        let (r, s): (usize, usize) = {
            let l = read_ints();
            (l[0], l[1])
        };
        case(i + 1, r, s);
    }
}

// ---- start boilerplate ----

fn read_line() -> String {
    use std::io::stdin;
    let mut line = String::new();
    stdin().read_line(&mut line).expect("read stdin");
    while line.ends_with('\n') {
        line.pop();
    }
    line
}

trait IntT: std::str::FromStr {}
impl IntT for i32 {}
impl IntT for i64 {}
impl IntT for usize {}

fn read_ints<T: IntT>() -> Vec<T> where <T as std::str::FromStr>::Err: std::fmt::Debug {
    read_line().split(' ').map(|part| T::from_str(part).expect("parse number")).collect()
}
