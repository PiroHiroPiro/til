# -*- coding: utf-8 -*-

from logic import Board
import numpy as np


class OthelloGame():
    """
    This class specifies the base Game class.
    To define your own game, subclass this class and implement the functions below.
    This works when the game is two-player, adversarial and turn-based.
    Use 1 for player1 and -1 for player2.
    See othello/OthelloGame.py for an example implementation.
    """

    def __init__(self, n=8):
        self.N = n

    def getInitBoard(self):
        """
        Returns:
            startBoard: a representation of the board (ideally this is the form that will be the input to your neural network)
        """

        b = Board(self.N)
        return np.array(b.pieces)

    def getBoardSize(self):
        """
        Returns:
            (x,y):  a tuple of board dimensions
        """

        return (self.N, self.N)

    def getActionSize(self):
        """
        Returns:
            actionSize: number of all possible actions
        """

        return self.N * self.N + 1

    def getNextState(self, board, player, action):
        """
        Input:
            board:  current board
            player: current player (1 or -1)
            action: action taken by current player
        Returns:
            nextBoard:  board after applying action
            nextPlayer: player who plays in the next turn (should be -player)
        """

        if action == self.N * self.N:
            return (board, -player)
        b = Board(self.N)
        b.pieces = np.copy(board)
        move = (int(action / self.N), action % self.N)
        b.execute_move(move, player)
        return (b.pieces, -player)

    def getValidMoves(self, board, player):
        """
        Input:
            board:  current board
            player: current player
        Returns:
            validMoves: a binary vector of length self.getActionSize(), 1 for moves that are valid from the current board and player, 0 for invalid moves
        """

        valids = [0] * self.getActionSize()
        b = Board(self.N)
        b.pieces = np.copy(board)
        legalMoves = b.get_legal_moves(player)
        if len(legalMoves) == 0:
            valids[-1] = 1
            return np.array(valids)
        for x, y in legalMoves:
            valids[self.N * x + y]=1
        return np.array(valids)

    def getGameEnded(self, board, player):
        """
        Input:
            board:  current board
            player: current player (1 or -1)
        Returns:
            r:  0 if game has not ended.
                1 if player won, -1 if player lost, small non-zero value for draw.
        """

        b = Board(self.N)
        b.pieces = np.copy(board)
        if b.has_legal_moves(player):
            return 0
        if b.has_legal_moves(-player):
            return 0
        if b.countDiff(player) > 0:
            return 1
        return -1

    def getCanonicalForm(self, board, player):
        """
        Input:
            board:  current board
            player: current player (1 or -1)
        Returns:
            canonicalBoard: returns canonical form of board.
                            The canonical form should be independent of player.
                            For e.g. in chess, the canonical form can be chosen to be from the pov of white.
                            When the player is white, we can return board as is.
                            When the player is black, we can invert the colors and return the board.
        """

        # return state if player==1, else return -state if player==-1
        return player * board

    def getSymmetries(self, board, pi):
        """
        Input:
            board:  current board
            pi:     policy vector of size self.getActionSize()
        Returns:
            symmForms:  a list of [(board,pi)] where each tuple is a symmetrical form of the board and the corresponding pi vector.
                        This is used when training the neural network from examples.
        """

        # mirror, rotational
        # 1 for pass
        assert(len(pi) == self.N ** 2 + 1)
        pi_board = np.reshape(pi[:-1], (self.N, self.N))
        l = []

        for i in range(1, 5):
            for j in [True, False]:
                newB = np.rot90(board, i)
                newPi = np.rot90(pi_board, i)
                if j:
                    newB = np.fliplr(newB)
                    newPi = np.fliplr(newPi)
                l += [(newB, list(newPi.ravel()) + [pi[-1]])]
        return l

    def stringRepresentation(self, board):
        """
        Input:
            board:  current board
        Returns:
            boardString:    a quick conversion of board to a string format.
                            Required by MCTS for hashing.
        """

        # 8x8 numpy array (canonical board)
        return board.tostring()

    def getScore(self, board, player):
        b = Board(self.N)
        b.pieces = np.copy(board)
        return b.countDiff(player)

def display(board):
    n = board.shape[0]

    for y in range(n):
        print (y, "|", end="")
    print("")
    print(" -----------------------")
    for y in range(n):
        # print the row
        print(y, "|", end="")
        for x in range(n):
            # get the piece to print
            piece = board[y][x]
            if piece == -1:
                print("b ", end="")
            elif piece == 1:
                print("W ", end="")
            else:
                if x==n:
                    print("-", end="")
                else:
                    print("- ", end="")
        print("|")

    print(" -----------------------")
