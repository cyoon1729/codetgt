package crdt

import (
	"strings"
)

type BlockId struct {
	uid   int
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
	clocks map[int]int
}

func initDoc() Doc {
	empty := Doc{
		blocks: make(Blocks, 0),
		clocks: make(map[int]int),
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
	return (id == BlockId{} || id.clock == 0)
}

func (blocks Blocks) findTruePosition(idx int) int {
	pos := 0
	for ; pos < len(blocks); pos++ {
		if blocks[pos].isDeleted {
			continue
		}
		if idx == 0 {
			return pos
		}
		idx -= 1
	}

	return pos
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
	leftIdx := -1
	if !isNullBlock(newBlock.leftOrigin) {
		leftIdx = blocks.findBlockIdx(newBlock.leftOrigin, idxHint-1)
	}
	rightIdx := len(blocks)
	if !isNullBlock(newBlock.rightOrigin) {
		rightIdx = blocks.findBlockIdx(newBlock.rightOrigin, idxHint)
	}
	destIdx := leftIdx + 1
	scanning := false

	for i := destIdx; i <= len(blocks); i++ {
		if !scanning {
			destIdx = i
		}
		if i == rightIdx || i == len(blocks) {
			break
		}

		other := blocks[i]
		oleft := blocks.findBlockIdx(other.leftOrigin, -1)
		oright := blocks.findBlockIdx(other.rightOrigin, -1)

		if oleft < leftIdx {
			break
		} else if oleft == leftIdx {
			if oright == rightIdx && newBlock.id.uid <= other.id.uid {
				break
			}
			scanning = newBlock.id.uid <= other.id.uid
		} else {
			continue
		}
	}

	return destIdx
}

func integrate(doc Doc, newBlock Block, idxHint int) Doc {
	blocks := doc.blocks
	destIdx := blocks.findInsertIdx(newBlock, idxHint)

	integrated := make(Blocks, len(blocks))
	copy(integrated, blocks)

	if destIdx >= 0 && destIdx < len(integrated) {
		integrated = append(integrated[:destIdx+1], integrated[destIdx:]...)
		integrated[destIdx] = newBlock
	} else {
		integrated = append(integrated, newBlock)
	}

	return Doc{integrated, doc.clocks}
}

func insert(doc Doc, uid int, idx int, value string) Doc {
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
	inserted := integrate(doc, newBlock, truePos)
	inserted.clocks[uid] += 1
	return inserted
}

func delete(doc Doc, idx int) Doc {
	tombstoned := make(Blocks, len(doc.blocks))
	copy(tombstoned, doc.blocks)
	truePos := tombstoned.findTruePosition(idx)
	tombstoned[truePos].isDeleted = true

	return Doc{tombstoned, doc.clocks}
}

func canInsertNow(doc Doc, seen map[BlockId]bool, block Block) bool {
	ok := seen[block.id]
	leftOk := seen[block.leftOrigin]
	rightOk := seen[block.rightOrigin]
	return (!ok && leftOk && rightOk)
}

func merge(src Doc, dst Doc) Doc {
	deletes := make(map[BlockId]bool)
	isDeleted := func(b Block) bool { return b.isDeleted }
	for _, b := range filterBlocks(isDeleted, src.blocks) {
		deletes[b.id] = true
	}

	for i, b := range dst.blocks {
		if deletes[b.id] {
			dst.blocks[i].isDeleted = true
		}
	}

	seen := make(map[BlockId]bool)
	for _, b := range dst.blocks {
		seen[b.id] = true
	}

	shouldMerge := func(b Block) bool { return !seen[b.id] }
	needsIntegrate := filterBlocks(shouldMerge, src.blocks)
	numToIntegrate := len(needsIntegrate)

	for numToIntegrate > 0 {
		for _, b := range needsIntegrate {
			if canInsertNow(dst, seen, b) {
				seen[b.id] = true
				dst = integrate(dst, b, -1)
				numToIntegrate -= 1
			}
		}
	}

	return dst
}
