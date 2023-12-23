import fileinput

LOSE = 0
DRAW = 3
WIN = 6

class Rock:
    score = 1
    def complement(self, outcome):
        if outcome == LOSE:
            return Scissors()
        elif outcome == WIN:
            return Paper()
        return Rock()

class Paper:
    score = 2
    def complement(self, outcome):
        if outcome == LOSE:
            return Rock()
        elif outcome == WIN:
            return Scissors()
        return Paper()

class Scissors:
    score = 3
    def complement(self, outcome):
        if outcome == LOSE:
            return Paper()
        elif outcome == WIN:
            return Rock()
        return Scissors()

moves = {
        'A': Rock(),
        'B': Paper(),
        'C': Scissors(),
        'X': LOSE,
        'Y': DRAW,
        'Z': WIN,
        }
total = 0

for line in fileinput.input():
    enemy, outcome = [moves[x] for x in line.rstrip().split(' ')]
    total += enemy.complement(outcome).score
    total += outcome
    #print(f"{you.beats(enemy)} + {you.score} = {you.beats(enemy)+you.score}")

print(total)
