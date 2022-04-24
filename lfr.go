package lfu

type LFUCache struct {
	size      int
	capacity  int
	minFreq   int
	key2Val   map[int]int
	key2Freq  map[int]int
	freq2Keys map[int]*LinkedHashSet
}

func (l *LFUCache) increaseFreq(key int) {
	freq := l.key2Freq[key]

	// update kf
	l.key2Freq[key] = freq + 1

	// update kf
	l.freq2Keys[freq].deleteKey(key)

	// insert key to freq + 1
	if _, ok := l.freq2Keys[freq+1]; !ok {
		l.freq2Keys[freq+1] = createdLinkedHashSet()
	}

	l.freq2Keys[freq+1].add(key)

	if l.freq2Keys[freq].isEmpty() {
		delete(l.freq2Keys, freq)
		if freq == l.minFreq {
			l.minFreq++
		}
	}

}

func (l *LFUCache) removeMinFreqKey() {
	keyList := l.freq2Keys[l.minFreq]
	deletedKey := keyList.cache.removeFirst().key
	delete(keyList.key2node, deletedKey)
	if keyList.isEmpty() {
		delete(l.freq2Keys, l.minFreq)
		// ? update minFreq
		// will set minFreq to 1,
		// who call this func ?
	}

	// update kv
	delete(l.key2Val, deletedKey)

	// update kf
	delete(l.key2Freq, deletedKey)
}

func New(capacity int) *LFUCache {
	key2Val := map[int]int{}
	key2Freq := map[int]int{}
	freq2Keys := map[int]*LinkedHashSet{}
	return &LFUCache{
		size:      0,
		capacity:  capacity,
		minFreq:   0,
		key2Val:   key2Val,
		key2Freq:  key2Freq,
		freq2Keys: freq2Keys,
	}
}

func (l *LFUCache) Get(key int) int {
	if _, ok := l.key2Val[key]; !ok {
		return -1
	}
	l.increaseFreq(key)
	return l.key2Val[key]
}

func (l *LFUCache) Put(key, val int) {
	if l.capacity <= 0 {
		return
	}

	if _, ok := l.key2Val[key]; ok {
		l.key2Val[key] = val
		l.increaseFreq(key)
		return
	}

	if l.capacity <= len(l.key2Val) {
		l.removeMinFreqKey()
	}

	// insert key & val, freq = 1
	l.key2Val[key] = val
	l.key2Freq[key] = 1
	if _, ok := l.freq2Keys[1]; !ok {
		l.freq2Keys[1] = createdLinkedHashSet()
	}
	l.freq2Keys[1].add(key)
	l.minFreq = 1
}
