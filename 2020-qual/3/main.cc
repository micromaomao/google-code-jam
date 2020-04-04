#include <iostream>
#include <vector>
#include <utility>
#include <algorithm>

using namespace std;

typedef struct Activity {
	int start;
	int end;
	char assigned = '\0';
};

int main(int argc, char const *argv[])
{
	int t;
	cin >> t;
	for (int i = 0; i < t; i ++) {
		int nb_activities;
		cin >> nb_activities;
		vector<Activity*> activities;
		activities.reserve(nb_activities);
		for (int i = 0; i < nb_activities; i ++) {
			int start, end;
			cin >> start >> end;
			Activity* act = new Activity{.start = start, .end = end};
			activities.push_back(act);
		}
		vector<Activity*> sorted_activities = activities;
		sort(sorted_activities.begin(), sorted_activities.end(), [](Activity* &a, Activity* &b) {
			return a->start < b->start;
		});

		int c_busy_until = 0;
		int j_busy_until = 0;

		bool impossible = false;

		for (Activity* &act : sorted_activities) {
			int start = act->start;
			if (c_busy_until <= start) {
				act->assigned = 'C';
				c_busy_until = act->end;
			} else if (j_busy_until <= start) {
				act->assigned = 'J';
				j_busy_until = act->end;
			} else {
				impossible = true;
				break;
			}
		}

		cout << "Case #" << i + 1 << ": ";
		if (impossible) {
			cout << "IMPOSSIBLE" << endl;
		} else {
			for (Activity* &act : activities) {
				cout << act->assigned;
				delete act;
			}
			cout << endl;
		}
	}
	return 0;
}
