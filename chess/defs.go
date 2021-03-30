package chess

const (
	// PIECE TYPES
	empty = 0

	pawn byte = 0b1
	king byte = 0b10
	knight byte = 0b11
	bishop byte = 0b100
	rook byte = 0b101
	queen byte = 0b110

	wpawn byte = 0b1001
	wking byte = 0b1010
	wknight byte = 0b1011
	wbishop byte = 0b1100
	wrook byte = 0b1101
	wqueen byte = 0b1110

	bpawn byte = 0b10001
	bking byte = 0b10010
	bknight byte = 0b10011
	bbishop byte = 0b10100
	brook byte = 0b10101
	bqueen byte = 0b10110

	// SIDES
	white int8 = 1
	black int8 = -1

	// PIECE MOVE OFFSETS REPRESENTATION:
	up        = -16 // WE RESERVED 8 BITS FOR OFFBOARD DETECTION FUNCTION
	down      = 16
	left      = -1
	right     = 1
	upRight   = -15 // UP + RIGHT
	downRight = 17 // DOWN + RIGHT
	downLeft  = 15 // DOWN + LEFT
	upLeft    = -17 // UP + LEFT
)

var (
	moveVectors = map[byte][]int {
		king: []int{upRight, right, downRight, down, downLeft, left, upLeft, up},
		knight: []int{(up << 1) + right, (right << 1) + up, (right << 1) + down, (down << 1) + right,
			(down << 1) + left, (left << 1) + down, (left << 1) + up, (up << 1) + left},
		bishop: []int{upRight, downRight, downLeft, upLeft},
		rook: []int{right, down, left, up},
		queen: []int{upRight, right, downRight, down, downLeft, left, upLeft, up},
	}
	pawnCaptureVectors = map[byte][]int {
		wpawn: []int{upLeft, upRight},
		bpawn: []int{downLeft, downRight},
	}
	pawnMoveVectors = map[byte][]int {
		wpawn: []int{up, up<<1},
		bpawn: []int{down, down<<1},
	}
)

// SQUARE REPRESENTATION: 1[OFFBOARD] 111[RANK] 1[OFFBOARD] 111[FILE]
func isOnBoard(square int) bool {
	return square&0x88 == 0
}

// CONVERT SQUARE INDEX TO NAME
func index2name(square int) [2]int {
	return [2]int{square % 16 + 'A', 8 - square / 16 + '1'}
}

func name2index(square string) int {
	return int((8 - square[1] - '1') / 16 + (square[0] - 'A') % 16)
}

// PIECE REPRESENTATION: 0000 1[SIDE] 111[TYPE]
func getPieceSide(piece byte) int8 {
	if piece >> 3 == 1 {
		return 1
	}
	return -1
}

// TYPE IS REPRESENTED BY 3 LAST BITS, 6 IS A MAX VALUE 
func getPieceType(piece byte) byte {
	return piece & 0b111 
}

// PIECE GROUP REPRESENTATION: LEAPER = FALSE, RANGE = TRUE
func isRangePiece(piece byte) bool {
	return getPieceType(piece) > 3 // [1-3] LEAPERS, [4-6] RANGE PIECES 
}



