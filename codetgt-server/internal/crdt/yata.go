package crdt

type BlockId struct {
	uid   string
	clock int
}

type Block struct {
	id          BlockId
	leftOrigin  BlockId
	rightOrigin BlockId
	isDeleted   bool
	value       string
}

type Blocks []Block

type Yata struct {
	blocks     Blocks
	lastClocks map[string]int
}

func isNullBlock(id BlockId) bool {
	return id != BlockId{}
}

func (blocks Blocks) findTruePosition(idx int) int {
	truePos := 0
	iter := 0
	for iter < idx {
		if blocks[iter].isDeleted {
			iter = iter + 1
		}
		truePos += 1
	}

	return truePos
}

func (blocks Blocks) safeGetId(idx int) BlockId {
	if idx < 0 || idx >= len(blocks) {
		return BlockId{}
	}
	return blocks[idx].id
}

func (blocks Blocks) findBlockIdx(targetId BlockId, idxHint int) int {
	if idxHint >= 0 && idxHint < len(blocks) {
		hint_item := blocks[idxHint]
		if hint_item.id == targetId {
			return idxHint
		}
	}

	idx := -1
	for i, b := range blocks {
		if b.id == targetId {
			idx = i
			break
		}
	}

	return idx
}

func (blocks Blocks) findInsertIdx(newBlock Block, idxHint int) int {
	leftIdx := blocks.findBlockIdx(newBlock.leftOrigin, idxHint-1)
	destIdx := leftIdx + 1
	rightIdx := len(blocks)
	if !isNullBlock(newBlock.leftOrigin) {
		blocks.findBlockIdx(newBlock.rightOrigin, idxHint)
	}
	scanning := false

	for i := destIdx; i < len(blocks); {
		if !scanning {
			destIdx = i
		}
		if i == rightIdx {
			break
		}

		other := blocks[i]
		oleft := blocks.findBlockIdx(other.leftOrigin, idxHint-1)
		oright := blocks.findBlockIdx(other.rightOrigin, idxHint)

		if oleft < leftIdx {
			break
		} else if oleft == leftIdx {
			if oright < rightIdx {
				scanning = true
				continue
			} else if oright == rightIdx {
				if newBlock.id.uid < other.id.uid {
					break
				} else {
					scanning = false
					continue
				}
			} else { // oright > rightIdx
				scanning = false
				continue
			}
		} else { // oleft > leftIdx
			continue
		}
	}

	return destIdx
}

func (yata Yata) integrate(newBlock Block, idxHint int) Yata {
	blocks := yata.blocks
	destIdx := blocks.findInsertIdx(newBlock, idxHint)
	blocks = append(blocks[:destIdx+1], blocks[destIdx:]...)
	blocks[destIdx] = newBlock

	return Yata{blocks, yata.lastClocks}
}

func (yata Yata) insertYata(uid string, idx int, value string) Yata {
	blocks := yata.blocks
	lastClocks := yata.lastClocks

	truePos := blocks.findTruePosition(idx)
	newBlock := Block{
		id:          BlockId{uid, lastClocks[uid] + 1},
		leftOrigin:  blocks.safeGetId(truePos - 1),
		rightOrigin: blocks.safeGetId(truePos),
		isDeleted:   false,
		value:       value,
	}
	inserted := yata.integrate(newBlock, truePos)
	return inserted
}

/* Synchronous insert */
func (yata Yata) syncInsert(uid string, idx int, value string) Yata {
	blocks := yata.blocks
	lastClocks := yata.lastClocks

	truePos := blocks.findTruePosition(idx)
	leftOrg := blocks.safeGetId(truePos - 1)
	rightOrg := blocks.safeGetId(truePos)
	clock := lastClocks[uid] + 1
	newBlock := Block{
		id:          BlockId{uid, clock},
		leftOrigin:  leftOrg,
		rightOrigin: rightOrg,
		isDeleted:   false,
		value:       value,
	}

	lastClocks[uid] = clock
	inserted := append(blocks[:truePos+1], blocks[truePos:]...)
	inserted[truePos] = newBlock

	return Yata{inserted, lastClocks}
}

func (yata Yata) deleteBlock(idx int) Yata {
	tombstoned := yata.blocks
	truePos := tombstoned.findTruePosition(idx)
	tombstoned[truePos].isDeleted = true

	return Yata{tombstoned, yata.lastClocks}
}
