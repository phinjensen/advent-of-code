import fileinput

LOSE = 0
DRAW = 3
WIN = 6

class Rock:
    score = 1
    def beats(self, other):
        if type(other) is Paper:
            return LOSE 
        if type(other) is Scissors:
            return WIN
        return DRAW 

class Paper:
    score = 2
    def beats(self, other):
        if type(other) is Scissors:
            return LOSE
        if type(other) is Rock:
            return WIN
        return DRAW 

class Scissors:
    score = 3
    def beats(self, other):
        if type(other) is Rock:
            return LOSE
        if type(other) is Paper:
            return WIN
        return DRAW 

moves = {
        'A': Rock(),
        'B': Paper(),
        'C': Scissors(),
        'X': Rock(),
        'Y': Paper(),
        'Z': Scissors(),
        }
total = 0

for line in fileinput.input():
    enemy, you = [moves[x] for x in line.rstrip().split(' ')]
    total += you.beats(enemy)
    total += you.score
    #print(f"{you.beats(enemy)} + {you.score} = {you.beats(enemy)+you.score}")

print(total)
