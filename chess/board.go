package chess

type PiecesMap map[int]byte

type Board struct {
	Pieces        PiecesMap
	Kings         [2]byte
	MovesNext     byte
	CastlePerm    byte
	EnPassant     byte
	HalfmoveClock byte
}

func (b *Board) forwardMove(move [2]int) byte { // [0]FROM [1]TO
	capturedPiece := b.Pieces[move[1]]
	b.Pieces[move[1]] = b.Pieces[move[0]]
	b.Pieces[move[0]] = empty
	b.MovesNext = -1 * b.MovesNext
	return capturedPiece 
}

func (b *Board) undoMove(move [2]int, capturedPiece byte) { // [0]FROM [1]TO
	b.Pieces[move[0]] = b.Pieces[move[1]]
	b.MovesNext = -1 * b.MovesNext
	if capturedPiece == empty {
		delete(b.Pieces, move[1])
	} else {
		b.Pieces[move[1]] = capturedPiece
	}
}

func (b *Board) isAttacked(square int) bool {
	for pieceType, offsetSlice := range moveVectors {
		for _, offset := range offsetSlice {
			targetSquare := square + offset
			for isOnBoard(targetSquare) && b.Pieces[targetSquare] == empty {
				targetSquare += offset
			}
			if isOnBoard(targetSquare) && 
			getPieceSide(b.Pieces[targetSquare]) == b.MovesNext &&
			getPieceType(b.Pieces[targetSquare]) == pieceType {
				return true
			}
		}
	}
	opponentPawns := bpawn
	if b.MovesNext == black {
		opponentPawns = wpawn
	}
	for _, offset := range pawnCaptureVectors[opponentPawns] {
		targetSquare := square + offset
		if isOnBoard(targetSquare) && 
		getPieceSide(b.Pieces[targetSquare]) == b.MovesNext {
			return true
		}
	}
	return false
}

func (b *Board) GenAllowedMoves() {
	movesWithoutCapture := make([][2]int, 0, 200)
	movesWithCapture := make([][2]int, 0, 100)
	for square, piece := range b.Pieces {
		if getPieceSide(piece) != b.MovesNext {
			continue
		}
		// FIND ALL PSEUDO-VALID MOVES AND CAPTURES FOR ALL PIECE TYPES EXCEPT PAWN
		for index, offset := range moveVectors[getPieceType(piece)] {
			targetSquare := square + offset
			maxDistance := 1
			if isRangePiece(piece) {
				maxDistance = 7
			}
			
			for isOnBoard(targetSquare) && maxDistance > 0 {
				if b.Pieces[targetSquare] == empty {
					movesWithoutCapture = append(movesWithoutCapture, [2]int{square, targetSquare})
				} else if getPieceSide(b.Pieces[targetSquare]) != b.MovesNext {
					movesWithCapture = append(movesWithCapture, [2]int{square, targetSquare})
					break;
				} else {
					break;
				}
				targetSquare += offset
				maxDistance--;
			}
		}
		// FIND ALL PSEUDO-VALID PAWN CAPTURES
		for _, offset := range pawnCaptureVectors[getPieceType(piece)] {
			targetSquare := square + offset
			if b.Pieces[targetSquare] != empty && getPieceSide(b.Pieces[targetSquare]) =! b.MovesNext {
				movesWithCapture = append(movesWithCapture, [2]int{square, targetSquare})
			}
		}
		// FIND ALL PSEUDO-VALID PAWN MOVES
		for _, offset := range pawnMoveVectors[getPieceType(piece)] {
			targetSquare := square + offset
			if b.Pieces[targetSquare] == empty {
				movesWithoutCapture = append(movesWithoutCapture, [2]int{square, targetSquare})
			} else {
				break
			}
		}
	}
}

