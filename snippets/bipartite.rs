use std::collections::HashSet;
// FIXME: may be wrong
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
			.find(|x| x.is_some()).map(|x| x.unwrap());
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

#[test]
fn test() {
	fn check_matches_validity(matches: &HashSet<(usize, usize)>, m: usize, n: usize) {
		let mut match_ltor: Vec<Option<usize>> = vec![None; m];
		let mut match_rtol: Vec<Option<usize>> = vec![None; n];
		for (ref i, ref j) in matches.iter() {
			assert!(*i < m && *j < n);
			if match_ltor[*i].is_some() {
				panic!("L{} already matched to R{}, but got ({}, {})", *i, match_ltor[*i].unwrap(), *i, *j);
			} else {
				match_ltor[*i] = Some(*j);
			}
			if match_rtol[*j].is_some() {
				panic!("R{} already matched to L{}, but got ({}, {})", *j, match_rtol[*j].unwrap(), *i, *j);
			} else {
				match_rtol[*j] = Some(*i);
			}
		}
	}

	fn it(m: usize, n: usize, connections: &Vec<(usize, usize)>, expected_len: usize) {
		let r = bipartite_match(m, n, connections);
		check_matches_validity(&r, m, n);
		assert_eq!(r.len(), expected_len);
		use std::iter::FromIterator;
		assert!(HashSet::from_iter(connections.into_iter().map(|x| *x)).is_superset(&r));
	}

	it(1, 3, &vec![(0, 0), (0, 1), (0, 2)], 1);
	it(3, 1, &vec![(0, 0), (1, 0), (2, 0)], 1);

	it(2, 2, &vec![], 0);
	it(2, 2, &vec![(0, 1)], 1);
	it(2, 2, &vec![(0, 0), (0, 1), (1, 0)], 2);

	it(4, 4, &vec![(0, 3), (1, 1), (1, 2), (2, 2)], 3);
	it(3, 3, &vec![(0, 1), (0, 2), (2, 1)], 2);
	it(3, 3, &vec![(0, 1), (0, 2), (2, 1), (2, 0)], 2);

	it(4, 3, &vec![(0, 0), (0, 1), (0, 2), (1, 0), (2, 2)], 3);

	it(3,2,&vec![(2,1),(0,0),(1,1)], 2);
}
