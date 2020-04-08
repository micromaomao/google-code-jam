fn test(b: usize) -> bool {
	assert!(b % 2 == 0);

	fn make_query(i: usize) -> Option<u8> {
		println!("{}", i + 1);
		let ans = read_line();
		match &ans[..] {
			"0" => Some(0u8),
			"1" => Some(1u8),
			_ => None
		}
	}

	macro_rules! unwrap_or_ret {
		($e:expr) => {{
			if let Some(r) = $e {
				r
			} else {
				return false;
			}
		}};
	}

	let mut arr: Vec<Option<u8>> = vec![None; b];
	let mut first_different_pair: Option<(usize, usize)> = None;
	let mut first_same_pair: Option<(usize, usize)> = None;
	let mut next_query_no = 1usize;
	let mut ptr_l = 0usize;
	let mut ptr_r = b - 1;

	fn invert_arr(arr: &mut [Option<u8>]) {
		for i in 0..arr.len() {
			if let Some(ai) = arr[i] {
				assert!(ai < 2);
				arr[i] = Some(1 - ai);
			}
		}
	}

	macro_rules! query {
		($i:expr) => {{
			next_query_no += 1;
			unwrap_or_ret!(make_query($i))
		}};
	}

	while ptr_l < ptr_r {
		if next_query_no != 1 && next_query_no % 10 == 1 {
			if first_same_pair.is_none() || first_different_pair.is_none() {
				// first_same_pair.is_none(): reverse and invert is the same in this case, hence we can assume either invert or no invert.
				// first_different_pair.is_none(): reverse has no effect. check for invert.
				if query!(0) != arr[0].unwrap() {
					invert_arr(&mut arr);
				}
			} else {
				let same_i = first_same_pair.unwrap().0;
				let same_orig = arr[same_i].unwrap();
				if query!(same_i) != same_orig {
					// reverse can't achieve this. hence must be invert.
					invert_arr(&mut arr);
				}

				let diff_i = first_different_pair.unwrap().0;
				let diff_orig = arr[diff_i].unwrap();
				// invert has been fixed now, so only reverse possible.
				if query!(diff_i) != diff_orig {
					arr.reverse();
				}
			}
		} else {
			let i = ptr_l;
			let j = ptr_r;
			if (next_query_no + 1) % 10 == 1 {
				// to not break invarient, we waste a dummy query.
				query!(0);
				continue;
			}
			assert_eq!(arr[i], None);
			assert_eq!(arr[j], None);
			arr[i] = Some(query!(i));
			arr[j] = Some(query!(j));
			if first_different_pair.is_none() && arr[i].unwrap() != arr[j].unwrap() {
				first_different_pair = Some((i, j));
			}
			if first_same_pair.is_none() && arr[i].unwrap() == arr[j].unwrap() {
				first_same_pair = Some((i, j));
			}
			ptr_l += 1;
			ptr_r -= 1;
		}
	}
	let mut res = String::with_capacity(b);
	for it in arr {
		res.push(match it.unwrap() {
			0 => '0',
			1 => '1',
			_ => unreachable!()
		});
	}
	println!("{}", &res);
	&read_line() == "Y"
}

fn main() {
	let (t, b) = {
		let l = read_ints();
		assert_eq!(l.len(), 2);
		(l[0], l[1])
	};
	for _ in 0..t {
		if !test(b) {
			return;
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
