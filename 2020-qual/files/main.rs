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
      if let &Node::Left(_) = self {
        true
      } else {
        false
      }
    }

    fn is_right(&self) -> bool {
      if let &Node::Right(_) = self {
        true
      } else {
        false
      }
    }

    fn index(&self) -> usize {
      match *self {
        Node::Left(n) => n,
        Node::Right(n) => n
      }
    }
  }

  let mut ltor = vec![Vec::new(); m];
  let mut rtol = vec![Vec::new(); n];
  for &(l, r) in connections.iter() {
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
    for &(i, j) in current_matches.iter() {
      l_matched[i] = true;
      r_matched[j] = true;
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

fn decompose_abc(n: i32, k: i32) -> Option<(i32, i32, i32)> {
  for a in 1..(n+1) {
    for b in 1..(n+1) {
      let c = k - b - a*(n-2);
      if c < 1 || c > n {
        continue;
      }
      if (a == b && a != c) || (a == c && a != b) {
        continue;
      }
      if n == 3 && b == c && a != b {
        continue;
      }
      debug_assert_eq!(a*(n-2)+b+c, k);
      return Some((a, b, c))
    }
  }
  None
}

fn solve(n: i32, k: i32) -> Option<Vec<i32>> {
  if n == 1 {
    return if k == 1 {
      Some(vec![1])
    } else {
      None
    };
  }
  if n == 2 {
    return match k {
      2 => Some(vec![1,2,2,1]),
      4 => Some(vec![2,1,1,2]),
      _ => None
    };
  }

  match decompose_abc(n, k) {
    None => None,
    Some((a, b, c)) => {
      let mut mat = vec![0i32; (n*n) as usize];
      mat[0] = b;
      mat[n as usize+1] = c;
      for x in 2..n as usize {
        mat[x*n as usize+x] = a;
      }

      for y in 0..n as usize {
        let mut connections = Vec::with_capacity((n*n) as usize);
        for x in 0..n as usize {
          if mat[y*n as usize+x] != 0 {
            connections.push((x, mat[y*n as usize+x] as usize));
          } else {
            let mut allowed = vec![true; (n+1) as usize];
            allowed[0] = false;
            for y in 0..n as usize {
              allowed[mat[y*n as usize+x] as usize] = false;
            }
            for a in 1..(n+1) as usize {
              if allowed[a] {
                connections.push((x, a));
              }
            }
          }
        }
        let m = bipartite_match(n as usize, (n + 1) as usize, &connections);
        assert_eq!(m.len(), n as usize);
        for (x, nb) in m {
          let nb = nb as i32;
          mat[y * n as usize + x] = nb;
        }
      }

      Some(mat)
    }
  }
}

fn main() {
  use std::fmt::Write;

  let t = read_ints()[0];
  for i in 0..t {
    let (n, k) = {
      let l = read_ints();
      (l[0], l[1])
    };
    let mat = solve(n, k);
    let possible = mat.is_some();
    println!("Case #{}: {}", i + 1, if possible { "POSSIBLE" } else { "IMPOSSIBLE" });
    if possible {
      let mat = mat.unwrap();
      for y in 0..n {
        let mut line = String::new();
        for x in 0..n {
          if x != 0 {
            write!(line, " ");
          }
          write!(line, "{}", mat[(y*n+x) as usize]);
        }
        println!("{}", line);
      }
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
