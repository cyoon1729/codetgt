package crdt

import (
	"strings"
)

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

type Doc struct {
	blocks Blocks
	clocks map[string]int
}

func initDoc() Doc {
	empty := Doc{
		blocks: make(Blocks, 0),
		clocks: make(map[string]int),
	}
	return empty
}

func (doc Doc) getContents() string {
	var sb strings.Builder

	for _, b := range doc.blocks {
		if !b.isDeleted {
			sb.WriteString(b.value)
		}
	}

	return sb.String()
}

func isNullBlock(id BlockId) bool {
	return id != BlockId{}
}

func (blocks Blocks) findTruePosition(idx int) int {
	truePos := 0
	iter := 0
	for iter < idx {
		if !blocks[iter].isDeleted {
			iter = iter + 1
		}
		truePos += 1
	}

	return truePos
}

func (blocks Blocks) safelyGetBlockId(idx int) BlockId {
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

	for i := destIdx; i < len(blocks); i++ {
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

func (doc Doc) integrate(newBlock Block, idxHint int) Doc {
	blocks := doc.blocks
	destIdx := blocks.findInsertIdx(newBlock, idxHint)

	if destIdx > 0 && destIdx < len(blocks) {
		blocks = append(blocks[:destIdx+1], blocks[destIdx:]...)
		blocks[destIdx] = newBlock
	} else {
		blocks = append(blocks, newBlock)
	}

	return Doc{blocks, doc.clocks}
}

func (doc Doc) insert(uid string, idx int, value string) Doc {
	blocks := doc.blocks
	clocks := doc.clocks

	truePos := blocks.findTruePosition(idx)
	newBlock := Block{
		id:          BlockId{uid, clocks[uid] + 1},
		leftOrigin:  blocks.safelyGetBlockId(truePos - 1),
		rightOrigin: blocks.safelyGetBlockId(truePos),
		isDeleted:   false,
		value:       value,
	}
	inserted := doc.integrate(newBlock, truePos)
	return inserted
}

func (doc Doc) delete(idx int) Doc {
	tombstoned := doc.blocks
	truePos := tombstoned.findTruePosition(idx)
	tombstoned[truePos].isDeleted = true

	return Doc{tombstoned, doc.clocks}
}
