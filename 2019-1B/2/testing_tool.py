"""Draupnir interactive judge.
"""

# Usage: `python testing_tool.py test_number`, where the argument test_number is
# either 0 (first test set) or 1 (second test set).
# This can also be run as `python3 testing_tool.py test_number`.

from __future__ import print_function
import sys

# Use raw_input in Python2.
try:
  input = raw_input
except NameError:
  pass

CASES = (((1, 2, 3, 4, 5, 6), (1, 1, 1, 1, 1, 1), (35, 97, 66, 44, 24, 19), (69, 81, 97, 84, 81, 70), (30, 18, 87, 97, 65, 46), (8, 29, 86, 49, 53, 96), (71, 96, 1, 10, 93, 96), (91, 62, 75, 6, 61, 14), (62, 37, 4, 55, 54, 96), (68, 70, 11, 78, 35, 40), (64, 14, 72, 53, 46, 88), (32, 20, 92, 8, 56, 25), (46, 72, 81, 62, 39, 42), (66, 47, 10, 66, 87, 50), (46, 21, 12, 79, 92, 42), (26, 24, 30, 5, 84, 37), (16, 15, 85, 83, 40, 18), (12, 55, 27, 76, 6, 58), (61, 59, 60, 18, 84, 76), (72, 92, 29, 75, 47, 17), (19, 8, 0, 51, 69, 66), (91, 41, 46, 57, 87, 8), (93, 86, 89, 7, 16, 82), (82, 1, 58, 32, 54, 71), (96, 49, 97, 70, 65, 85), (33, 65, 31, 81, 48, 70), (47, 48, 74, 58, 59, 59), (74, 0, 44, 28, 13, 68), (55, 5, 1, 87, 91, 59), (75, 22, 38, 10, 31, 64), (25, 69, 82, 76, 42, 86), (45, 35, 44, 12, 88, 13), (55, 34, 19, 12, 64, 64), (57, 9, 27, 78, 42, 99), (39, 16, 55, 24, 30, 26), (5, 39, 13, 17, 7, 75), (92, 74, 60, 22, 36, 26), (4, 19, 29, 59, 62, 91), (31, 84, 37, 96, 60, 34), (62, 38, 27, 59, 6, 93), (73, 22, 0, 84, 95, 4), (37, 32, 22, 46, 59, 80)),
         ((1, 2, 3, 4, 5, 6), (1, 1, 1, 1, 1, 1), (35, 97, 66, 44, 24, 19), (69, 81, 97, 84, 81, 70), (30, 18, 87, 97, 65, 46), (8, 29, 86, 49, 53, 96), (71, 96, 1, 10, 93, 96), (91, 62, 75, 6, 61, 14), (62, 37, 4, 55, 54, 96), (68, 70, 11, 78, 35, 40), (64, 14, 72, 53, 46, 88), (32, 20, 92, 8, 56, 25), (46, 72, 81, 62, 39, 42), (66, 47, 10, 66, 87, 50), (46, 21, 12, 79, 92, 42))) # set your own cases
WS = (6, 2)

MAX_DAY = 500
WRONG_ANSWER, CORRECT_ANSWER = -1, 1
MOD = 2 ** 63


DAY_OUT_OF_RANGE_ERROR = "Day {} is out of range.".format
EXCEEDED_QUERIES_ERROR = "Exceeded number of queries: {}.".format
INVALID_LINE_ERROR = "Couldn't read a valid line."
NOT_INTEGER_ERROR = "Not an integer: {}".format
RING_AMOUNT_OUT_OF_RANGE_ERROR = "Ring amount {} is out of range.".format
WRONG_NUM_TOKENS_ERROR = "Wrong number of tokens: {}. Expected 1 or 6.".format
WRONG_GUESS_ERROR = "Wrong guess: {}. Expected: {}.".format


def ReadValues(line):
  t = line.split()
  if len(t) != 1 and len(t) != 6:
    return WRONG_NUM_TOKENS_ERROR(len(t))
  r = []
  for s in t:
    try:
      v = int(s)
    except:
      return NOT_INTEGER_ERROR(s if len(s) < 100 else s[:100])
    r.append(v)
  if len(r) == 1:
    if not (1 <= r[0] <= MAX_DAY):
      return DAY_OUT_OF_RANGE_ERROR(r[0])
  else:
    for ri in r:
      if ri < 0:
        return RING_AMOUNT_OUT_OF_RANGE_ERROR(ri)
  return r


def ComputeDay(s, d):
  rings_per_day = list(s)
  for i in range(1, d + 1):
    for j in range(1, 7):
      if i % j == 0:
        rings_per_day[j - 1] += rings_per_day[j - 1]
        rings_per_day[j - 1] %= MOD
  return rings_per_day


def RunCase(w, case, test_input=None):
  outputs = []

  def Input():
    return input()

  def Output(line):
    print(line)
    sys.stdout.flush()

  for ex in range(w + 1):
    try:
      line = Input()
    except:
      Output(WRONG_ANSWER)
      return INVALID_LINE_ERROR, outputs
    v = ReadValues(line)
    if isinstance(v, str):
      Output(WRONG_ANSWER)
      return v, outputs
    if len(v) == 1:
      if ex == w:
        Output(WRONG_ANSWER)
        return EXCEEDED_QUERIES_ERROR(w), outputs
      else:
        Output(sum(ComputeDay(case, v[0])) % MOD)
    else:
      if tuple(v) != tuple(case):
        Output(WRONG_ANSWER)
        return WRONG_GUESS_ERROR(v, case)[:100], outputs
      else:
        Output(CORRECT_ANSWER)
        return None, outputs


def RunCases(w, cases):
  for i, case in enumerate(cases, 1):
    result, _ = RunCase(w, case)
    if result:
      return "Case #{} ({}) failed: {}".format(i, case, result)
  try:
    extra_input = input()
  except EOFError:
    return None
  except Exception:  # pylint: disable=broad-except
    return "Exception raised while reading input after all cases finish."
  return "Additional input after all cases finish: {}".format(extra_input[:100])


def main():
  assert len(sys.argv) == 2
  index = int(sys.argv[1])
  cases = CASES[index]
  w = WS[index]
  print(len(cases), w)
  sys.stdout.flush()
  result = RunCases(w, cases)
  if result:
    print(result, file=sys.stderr)
    sys.stdout.flush()
    sys.exit(1)


if __name__ == "__main__":
  main()
