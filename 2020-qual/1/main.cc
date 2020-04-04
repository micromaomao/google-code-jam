#include <iostream>
#include <vector>
#include <iterator>
#include <unordered_set>

using namespace std;

int main(int argc, char const *argv[])
{
	cin.sync_with_stdio(false);
	int t;
	cin >> t;
	for (int i = 0; i < t; i ++) {
		int n;
		cin >> n;
		vector<int> mat;
		mat.reserve(n*n);
		for (int i = 0; i < n*n; i ++) {
			int r;
			cin >> r;
			mat.push_back(r);
		}
		int trace = 0, nb_repeating_row = 0, nb_repeating_col = 0;
		for (int i = 0; i < n; i ++) trace += mat.at(i*n+i);
		for (int x = 0; x < n; x ++) {
			unordered_set<int> s;
			for (int y = 0; y < n; y ++) {
				int it = mat.at(y*n+x);
				if (it >= 1 && it <= n && s.find(it) == s.end()) {
					s.insert(it);
				} else {
					nb_repeating_col += 1;
					break;
				}
			}
		}
		for (int y = 0; y < n; y ++) {
			unordered_set<int> s;
			for (int x = 0; x < n; x ++) {
				int it = mat.at(y*n+x);
				if (it >= 1 && it <= n && s.find(it) == s.end()) {
					s.insert(it);
				} else {
					nb_repeating_row += 1;
					break;
				}
			}
		}
		cout << "Case #" << i + 1 << ": " << trace << " " << nb_repeating_row << " " << nb_repeating_col << endl;
	}
	return 0;
}
