tenTimes = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]


def Add(a, b):
    return a + b


def Multiply(a, b):
    return a * b


def Divide(a, b):
    return a / b


def IsEven(a):
    if a % 2 == 0:
        answerE = 1
    else:
        answerE = 0
    return answerE


def IsDivisible(a, b):
    if a % b == 0:
        answerD = 1
    else:
        answerD = 0
    return answerD


def ReducingByPercent(a, pUser):
    p = pUser / 100
    percentOfA = p * a
    reduceI = a - percentOfA
    return reduceI


def Sqrt(a):
    return a * a


def ToThePower(a, n):
    scorePI = a
    if n >= 1:
        for i in tenTimes:
            scorePI = scorePI * a
    else:
        scorePI = 1
    return scorePI
