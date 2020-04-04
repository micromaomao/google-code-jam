#include <iostream>
#include <string>

using namespace std;

int main(int argc, char const *argv[])
{
	int t;
	cin >> t;
	while (cin.get() != '\n') {}
	for (int i = 0; i < t; i ++) {
		string input, result;
		getline(cin, input);
		int current = 0;
		for (char &c : input) {
			int digit = stoi(string(&c, 1));
			if (digit == current) {
				result.push_back(c);
			} else if (digit > current) {
				while (current < digit) {
					current += 1;
					result.push_back('(');
				}
				result.push_back(c);
			} else if (digit < current) {
				while (current > digit) {
					current -= 1;
					result.push_back(')');
				}
				result.push_back(c);
			}
		}
		while (current > 0) {
			current -= 1;
			result.push_back(')');
		}
		cout << "Case #" << i + 1 << ": " << result << endl;
	}
	return 0;
}
