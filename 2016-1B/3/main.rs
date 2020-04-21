use std::collections::HashSet;
pub fn bipartite_match(m: usize, n: usize, connections: &[(usize, usize)]) -> HashSet<(usize, usize)> {
	debug_assert!(connections.iter().collect::<HashSet<_>>().len() == connections.len());

	#[derive(Clone, Copy, PartialEq, Eq, Hash, Debug)]
	enum Node {
		Left(usize),
		Right(usize)
	}

	impl Node {
		fn is_left(&self) -> bool {
			if let Node::Left(_) = self {
				true
			} else {
				false
			}
		}

		fn is_right(&self) -> bool {
			if let Node::Right(_) = self {
				true
			} else {
				false
			}
		}

		fn index(&self) -> usize {
			match self {
				Node::Left(ref n) => *n,
				Node::Right(ref n) => *n
			}
		}
	}

	let mut ltor = vec![Vec::new(); m];
	let mut rtol = vec![Vec::new(); n];
	for (ref l, ref r) in connections.iter() {
		let l = *l;
		let r = *r;
		if l >= m || r >= n {
			panic!("invalid input.");
		}

		ltor[l].push(r);
		rtol[r].push(l);
	}

	let mut current_matches: HashSet<(usize, usize)> = HashSet::new();

	fn aug_can_go(from: &Node, to: &Node, current_matches: &HashSet<(usize, usize)>) -> bool {
		if from.is_left() && to.is_right() {
			!current_matches.contains(&(from.index(), to.index()))
		} else if from.is_right() && to.is_left() {
			current_matches.contains(&(to.index(), from.index()))
		} else {
			unreachable!()
		}
	}

	loop {
		let mut l_matched: Vec<bool> = vec![false; m];
		let mut r_matched: Vec<bool> = vec![false; n];
		for (ref i, ref j) in current_matches.iter() {
			l_matched[*i] = true;
			r_matched[*j] = true;
		}
		fn dfs(
			n: Node,
			ltor: &[Vec<usize>], rtol: &[Vec<usize>],
			current_match: &HashSet<(usize, usize)>,
			r_matched: &[bool],
			visited: &mut HashSet<Node>
		) -> Option<Vec<Node>> {
			assert!(visited.insert(n));
			if n.is_right() && !r_matched[n.index()] {
				return Some(vec![n]);
			}
			let next_hops: Vec<Node> = match n {
				Node::Left(i) => ltor[i].iter().map(|j| Node::Right(*j)).collect(),
				Node::Right(i) => rtol[i].iter().map(|j| Node::Left(*j)).collect(),
			};
			for j in next_hops {
				if aug_can_go(&n, &j, current_match) && !visited.contains(&j) {
					if let Some(mut path) = dfs(j, ltor, rtol, current_match, r_matched, visited) {
						path.push(n);
						return Some(path);
					}
				}
			}
			None
		}
		let dfs_path = (0..m).into_iter()
			.filter(|i| !l_matched[*i])
			.map(|i| dfs(Node::Left(i), &ltor, &rtol, &current_matches, &r_matched, &mut HashSet::new()))
			.find(|x| x.is_some()).flatten();
		if let Some(mut path) = dfs_path {
			path.reverse();
			debug_assert!(path.len() >= 2);
			for w in path.windows(2) {
				let (i, j) = (w[0], w[1]);
				match i {
					Node::Left(i) => {
						debug_assert!(j.is_right());
						current_matches.insert((i, j.index()));
					},
					Node::Right(i) => {
						debug_assert!(j.is_left());
						assert!(current_matches.remove(&(j.index(), i)));
					}
				}
			}
		} else {
			break;
		}
	}

	current_matches
}

fn test() -> usize {
	let n: usize = read_ints()[0];
	use std::collections::HashMap;
	struct WordToId {
		next: usize,
		hm: HashMap<usize, String>,
		rev: HashMap<String, usize>,
	}
	impl WordToId {
		fn with_capacity(cap: usize) -> Self {
			WordToId {
				next: 0,
				hm: HashMap::with_capacity(cap),
				rev: HashMap::with_capacity(cap)
			}
		}
		fn insert(&mut self, w: String) -> usize {
			if self.rev.contains_key(&w) {
				*self.rev.get(&w).unwrap()
			} else {
				let id = self.next;
				self.next += 1;
				self.hm.insert(id, w.clone());
				self.rev.insert(w, id);
				id
			}
		}
		fn len(&self) -> usize {
			self.next
		}
	}
	let mut left = WordToId::with_capacity(n);
	let mut right = WordToId::with_capacity(n);
	let mut edges: HashSet<(usize, usize)> = HashSet::with_capacity(n);
	for _ in 0..n {
		let l = read_line();
		let parts: Vec<&str> = l.split(' ').collect();
		assert_eq!(parts.len(), 2);
		let w = (left.insert(parts[0].to_owned()), right.insert(parts[1].to_owned()));
		edges.insert(w);
	}
	let base_set = bipartite_match(left.len(), right.len(), &edges.iter().copied().collect::<Vec<_>>());
	let mut left_used: Vec<bool> = vec![false; left.len()];
	let mut right_used: Vec<bool> = vec![false; right.len()];
	for w in base_set.iter().copied() {
		left_used[w.0] = true;
		right_used[w.1] = true;
	}
	let mut nb_not_faked = base_set.len();
	for w in edges {
		if base_set.contains(&w) {
			continue;
		}
		if left_used[w.0] && right_used[w.1] {
			continue;
		}
		nb_not_faked += 1;
		left_used[w.0] = true;
		right_used[w.1] = true;
	}
	n - nb_not_faked
}

fn main() {
	let t: usize = read_ints()[0];
	for i in 1..=t {
		println!("Case #{}: {}", i, test());
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
