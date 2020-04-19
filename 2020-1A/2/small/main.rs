fn solve(n: i64) -> Vec<(i64, i64)> {
	let mut sum = 0i64;
	let mut i = 1i64;
	while sum + i < n {
		sum += i;
		i += 1;
	}
	let mut diff = n - sum;
	assert!(diff > 0 && diff <= i);
	let mut result = Vec::new();
	for i in 1..i {
		if diff > 0 {
			result.push((i, i));
			diff -= 1;
		}
		result.push((i+1, i));
	}
	if diff > 0 {
		result.push((i, i));
		diff -= 1;
	}
	assert_eq!(diff, 0);
	result
}

fn main() {
	let t: usize = {
		let l = read_ints();
		assert_eq!(l.len(), 1);
		l[0]
	};
	for c in 0..t {
		let n: i64 = {
			let l = read_ints();
			assert_eq!(l.len(), 1);
			l[0]
		};
		println!("Case #{}:", c + 1);
		for (r, k) in solve(n) {
			println!("{} {}", r, k);
		}
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
