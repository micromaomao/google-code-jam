fn solve(_patterns: Vec<String>) -> String {
	let mut patterns: Vec<Vec<PatternItem>> = Vec::with_capacity(_patterns.len());
	#[derive(Debug, Clone, Copy, PartialEq, Eq)]
	enum PatternItem {
		EndStar,
		Item(bool, char)
	}
	for pattern in _patterns.into_iter() {
		let mut pat = Vec::with_capacity(pattern.len());
		let chars: Vec<char> = pattern.chars().collect();
		let mut i = 0usize;
		while i < chars.len() {
			let c = chars[i];
			if c == '*' {
				if i + 1 < chars.len() && chars[i + 1] == '*' {
					i += 1;
					continue;
				}
				if i == chars.len() - 1 {
					pat.push(PatternItem::EndStar);
					i += 1;
				} else {
					pat.push(PatternItem::Item(true, chars[i + 1]));
					i += 2;
				}
			} else {
				pat.push(PatternItem::Item(false, chars[i]));
				i += 1;
			}
		}
		patterns.push(pat);
	}

	use std::collections::HashSet;
	#[derive(Hash, PartialEq, Eq, Clone, Debug)]
	struct State {
		prefix: String,
		at_pattern_positions: Vec<usize>
	}

	let mut current_states = HashSet::new();
	current_states.insert(State{
		prefix: String::new(),
		at_pattern_positions: vec![0usize; patterns.len()]
	});


	loop {
		if let Some(end_state) = current_states.iter().find(|s| {
			for (i, pos) in s.at_pattern_positions.iter().enumerate() {
				let pos = *pos;
				let pat = &patterns[i];
				if pos >= pat.len() {
					continue;
				}
				if pat[pos] != PatternItem::EndStar {
					return false;
				}
			}
			true
		}) {
			return end_state.prefix.to_owned();
		}

		let mut possible_next_letters = HashSet::new();
		for s in current_states.iter() {
			for (i, pos) in s.at_pattern_positions.iter().enumerate() {
				let pos = *pos;
				let pat = &patterns[i];
				if pos == pat.len() {
					continue;
				}
				match pat[pos] {
					PatternItem::Item(_, c) => {
						possible_next_letters.insert(c);
					},
					PatternItem::EndStar => {}
				};
			}
		}

		let mut new_states = HashSet::new();

		's: for s in current_states.iter() {
			for advance_char in possible_next_letters.iter().map(|x| *x) {
				let mut allow_star_patterns: Vec<&Vec<PatternItem>> = Vec::new();
				for (i, pos) in s.at_pattern_positions.iter().enumerate() {
					let pos = *pos;
					let pat = &patterns[i];
					if pos >= pat.len() {
						continue 's;
					}
					match pat[pos] {
						PatternItem::Item(allow_star, _) => {
							if allow_star {
								allow_star_patterns.push(pat);
							}
						},
						PatternItem::EndStar => {}
					}
				}

				let mut new_prefix = s.prefix.clone();
				new_prefix.push(advance_char);

				'a: for pred in 0..(1u64 << (allow_star_patterns.len() as u64)) {
					let mut new_poss = Vec::with_capacity(s.at_pattern_positions.len());
					let mut every_pattern_stared = true;
					for (i, pos) in s.at_pattern_positions.iter().enumerate() {
						let pos = *pos;
						let pat = &patterns[i];
						if pos >= pat.len() {
							continue;
						}
						match pat[pos] {
							PatternItem::Item(allow_star, c) => {
								if allow_star {
									let idx = allow_star_patterns.iter().position(|x| x == &pat).unwrap();
									if (pred >> idx) % 2 == 1 {
										// follow star
										new_poss.push(pos)
									} else if advance_char == c {
										new_poss.push(pos + 1);
										every_pattern_stared = false;
									} else {
										continue 'a;
									}
								} else if advance_char == c {
									new_poss.push(pos + 1);
									every_pattern_stared = false;
								} else {
									continue 'a;
								}
							},
							PatternItem::EndStar => {
								new_poss.push(pos);
							}
						}
					}

					assert_eq!(new_poss.len(), patterns.len());

					if !every_pattern_stared {
						new_states.insert(State{
							prefix: new_prefix.clone(),
							at_pattern_positions: new_poss
						});
					}
				}
			}
		}

		current_states = new_states;
		if current_states.len() == 0 {
			return "*".to_owned();
		}
	}
}

fn main() {
	let t: usize = {
		let l = read_ints();
		assert_eq!(l.len(), 1);
		l[0]
	};
	for c in 0..t {
		let n: usize = {
			let l = read_ints();
			assert_eq!(l.len(), 1);
			l[0]
		};
		let mut patterns: Vec<String> = Vec::with_capacity(n);
		for _ in 0..n {
			patterns.push(read_line());
		}
		println!("Case #{}: {}", c + 1, solve(patterns));
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
