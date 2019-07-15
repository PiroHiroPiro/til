# -*- coding: utf-8 -*-

import numpy as np

from game import Game
from board import Board


class Player():
    def __init__(self, game: Game):
        self.game = game


class RandomPlayer(Player) -> int:
    def play(self, board: Board):
        a = np.random.randint(self.game.getActionSize())
        valids = self.game.getValidMoves(board, 1)
        while valids[a] != 1:
            a = np.random.randint(self.game.getActionSize())
        return a


class HumanOthelloPlayer(Player):
    def play(self, board: Board) -> int:
        # display(board)
        valid = self.game.getValidMoves(board, 1)
        for i in range(len(valid)):
            if valid[i]:
                print(int(i / self.game.n), int(i % self.game.n))
        while True:
            print("x y > ", end="")
            x, y = map(int, input().split(" ")[:2])
            a = self.game.n * x + y if x != -1 else self.game.n ** 2
            if valid[a]:
                break
            else:
                print("Invalid")

        return a


class GreedyOthelloPlayer(Player):
    def play(self, board: Board) -> int:
        valids = self.game.getValidMoves(board, 1)
        candidates = []
        for a in range(self.game.getActionSize()):
            if valids[a] == 0:
                continue
            nextBoard, _ = self.game.getNextState(board, 1, a)
            score = self.game.getScore(nextBoard, 1)
            candidates += [(-score, a)]
        candidates.sort()
        return candidates[0][1]
