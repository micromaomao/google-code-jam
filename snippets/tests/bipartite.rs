extern crate snippets;
extern crate rand;
use snippets::bipartite::bipartite_match;

use std::collections::HashSet;

#[test]
fn basic_test() {
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

  fn it(m: usize, n: usize, connections: &[(usize, usize)], expected_len: usize) {
    let r = bipartite_match(m, n, connections);
    check_matches_validity(&r, m, n);
    assert_eq!(r.len(), expected_len);
    use std::iter::FromIterator;
    assert!(HashSet::from_iter(connections.into_iter().map(|x| *x)).is_superset(&r));

    assert_eq!(brute_force_max(m, n, connections), expected_len);
  }

  it(0, 0, &[], 0);

  it(1, 3, &[(0, 0), (0, 1), (0, 2)], 1);
  it(3, 1, &[(0, 0), (1, 0), (2, 0)], 1);

  it(2, 2, &[], 0);
  it(2, 2, &[(0, 1)], 1);
  it(2, 2, &[(0, 0), (0, 1), (1, 0)], 2);

  it(4, 4, &[(0, 3), (1, 1), (1, 2), (2, 2)], 3);
  it(3, 3, &[(0, 1), (0, 2), (2, 1)], 2);
  it(3, 3, &[(0, 1), (0, 2), (2, 1), (2, 0)], 2);

  it(4, 3, &[(0, 0), (0, 1), (0, 2), (1, 0), (2, 2)], 3);

  it(3,2,&[(2,1),(0,0),(1,1)], 2);
}

fn brute_force_max(m: usize, n: usize, connections: &[(usize, usize)]) -> usize {
  let mut ans = 0usize;
  'c: for choice_set in 0u64..(1u64 << connections.len() as u64) {
    let mut a_used = vec![false; m];
    let mut b_used = vec![false; n];
    for (i, c) in connections.iter().enumerate() {
      if choice_set & (1u64 << i as u64) > 0 {
        let a = c.0;
        let b = c.1;
        if a_used[a] || b_used[b] {
          continue 'c;
        }
        a_used[a] = true;
        b_used[b] = true;
      }
    }
    let new_ans = a_used.iter().filter(|x| **x).count();
    assert_eq!(new_ans, b_used.iter().filter(|x| **x).count());
    ans = usize::max(ans, new_ans);
  }
  ans
}

#[test]
fn generated_tests() {
  use rand::{Rng, SeedableRng};
  let mut rng = rand::rngs::SmallRng::seed_from_u64(0);

  fn generate<R: Rng>(rng: &mut R) -> (usize, usize, Vec<(usize, usize)>) {
    let m: usize = rng.gen_range(1,10);
    let n: usize = rng.gen_range(1, 10);
    let len: usize = rng.gen_range(0, usize::min(m*n, 20));
    let mut conns = HashSet::with_capacity(len);
    for _ in 0..len {
      loop {
        let conn: (usize, usize) = (rng.gen_range(0, m), rng.gen_range(0, n));
        if conns.contains(&conn) {
          continue;
        }
        conns.insert(conn);
        break;
      }
    }
    (m, n, conns.into_iter().collect())
  }

  #[cfg(debug_assertions)]
  const TEST_TIMES: usize = 10;
  #[cfg(not(debug_assertions))]
  const TEST_TIMES: usize = 1000;

  for _ in 0..TEST_TIMES {
    let (m, n, conns) = generate(&mut rng);
    let ans = brute_force_max(m, n, &conns);
    assert_eq!(bipartite_match(m, n, &conns).len(), ans);
  }
}
