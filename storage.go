package main

type IStorage interface {
	//GetMostUsed() FizzBuzzIn
	Store(FizzBuzzParams)
	GetAll() []FizzBuzzIn
}

func NewInmemory() IStorage {
	return &FizzBuzzStorage{
		Data: make(map[string]FizzBuzzIn),
	}
}

func (f *FizzBuzzStorage) GetAll() []FizzBuzzIn {
	var fizzBuzzList []FizzBuzzIn
	for _, fb := range f.Data {
		fizzBuzzList = append(fizzBuzzList, fb)
	}
	return fizzBuzzList
}

func (f *FizzBuzzStorage) Store(fizzBuzz FizzBuzzParams) {
	fizzBuzzUUID := fizzBuzz.UUID()
	oldFizzBuzz, ok := f.Data[fizzBuzzUUID]
	if ok {
		oldFizzBuzz.Hits += 1
		f.Data[fizzBuzzUUID] = oldFizzBuzz
	} else {
		f.Data[fizzBuzzUUID] = FizzBuzzIn{
			Params: fizzBuzz,
			Hits:   1,
		}
	}
}

/* func (f *FizzBuzzStorage) GetMostUsed() FizzBuzzIn {
	var maxHit int64
	var mostUsed FizzBuzzIn
	for _, fb := range f.Data {
		if fb.Hits > maxHit {
			maxHit = fb.Hits
			mostUsed = FizzBuzzIn{
				maxHit,
				fb.Params,
			}
		}
	}
	return mostUsed
} */
