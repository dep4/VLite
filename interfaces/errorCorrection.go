package interfaces

type ErrorCorrectionFacility interface {
	AddShard(id int, data []byte) (done bool, encoded []byte)
	Reconstruct() [][]byte

	AddData(data []byte) (id int, wrappedData []byte)
	ConstructReconstructShard() (id int, encoded []byte, more bool)

	MaxShardYieldRemaining() int
}

type RSParityShardSum struct {
	ParityLookupTable []int
}
